package chain

import (
	"fmt"
	"github.com/cybriq/kismet/pkg/block"
	"github.com/cybriq/kismet/pkg/hash"
	"github.com/cybriq/qu"
	"github.com/dgraph-io/badger/v3"
	"sort"
	"sync"
	"time"
)

type AccessTrackedBlock struct {
	hash.Hash
	time.Time
	*block.Block
}

type SortedAccess []AccessTrackedBlock

func (a SortedAccess) Len() int {
	return len(a)
}

func (a SortedAccess) Less(i, j int) bool {
	return a[i].Time.Before(a[j].Time)
}

func (a SortedAccess) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

type Blocks map[hash.Hash]AccessTrackedBlock

type Index struct {
	db      *badger.DB
	cache   Blocks
	cacheMx sync.Mutex
}

type IndexBlock struct {
	*Index
	*block.Block
}

func (ib *IndexBlock) Previous() (b *IndexBlock, err error) {

	var byHash *block.Block
	if byHash, err = ib.GetByHash(ib.Block.Previous); log.E.Chk(err) {
		return
	}

	b = &IndexBlock{Index: ib.Index, Block: byHash}
	return
}

type Chain []IndexBlock

// New creates a new block index. maxToCache is the maximum we will cache as
// recently used, and lowWaterMark is the number below maxToCache that a cache
// purge will reduce the cache size to when it hits the max.
func New(path string, maxToCache, lowWaterMark int, stop qu.C) (idx *Index, err error) {

	idx = &Index{cache: make(Blocks)}

	if idx.db, err = badger.Open(badger.DefaultOptions(path)); log.E.Chk(err) {
		return
	}

	// here we need a GC to expire the cache periodically if it gets huge
	go func() {

		timer := time.NewTicker(time.Minute)
		var sa SortedAccess
	out:
		for {
			select {
			case <-timer.C:

				idx.cacheMx.Lock()
				if len(idx.cache) > maxToCache {
					sa = make(SortedAccess, len(idx.cache))
					var counter int
					for i := range idx.cache {
						sa[counter] = idx.cache[i]
						counter++
					}
				} else {

					// unlock since we aren't accessing the cache
					idx.cacheMx.Unlock()
					break
				}

				// no need to hold lock while we sort our list
				idx.cacheMx.Unlock()

				sort.Sort(sa)
				sal := len(sa)

				idx.cacheMx.Lock()
				for i := lowWaterMark; i < sal; i++ {
					delete(idx.cache, sa[i].Hash)
				}
				idx.cacheMx.Unlock()

				// hint to GC we don't need this temporary list anymore
				sa = SortedAccess{}

			case <-stop.Wait():

				err = idx.db.Close()
				log.E.Chk(err)
				timer.Stop()
				break out
			}
		}
	}()

	return
}

func (idx *Index) Add(b *block.Block) (err error) {

	if b == nil {
		err = fmt.Errorf("cannot add nil block")
		log.E.Ln(err)
		return
	}

	// get the index hash to check we aren't re-adding
	var h hash.Hash
	if h, err = b.IndexHash(); log.E.Chk(err) {
		return
	}

	// check we are not adding the same block again
	idx.cacheMx.Lock()
	if _, found := idx.cache[h]; found {

		err = fmt.Errorf(
			"block already exists in cache, not adding as" +
				" it has been seen/added already",
		)
		idx.cacheMx.Unlock()
		log.E.Ln(err)
		return
	}
	idx.cacheMx.Unlock()

	// next, check if the block has already been stored in db but not recently
	// accessed.
	if err = idx.db.View(
		func(txn *badger.Txn) (err error) {

			if _, err = txn.Get(h[:]); err == nil {

				return fmt.Errorf("block already exists, not adding again")
			}
			return
		},
	); log.E.Chk(err) {

		idx.cacheMx.Lock()
		idx.cache[h] = AccessTrackedBlock{Time: time.Now(), Block: b}
		idx.cacheMx.Unlock()
		return
	}

	// add the block to the cache for fast retrieval
	idx.cacheMx.Lock()
	idx.cache[h] = AccessTrackedBlock{Time: time.Now(), Block: b}
	idx.cacheMx.Unlock()

	var blk block.WireBlock
	if blk, err = b.Marshal(); log.E.Chk(err) {
		return
	}

	if err = idx.db.Update(
		func(txn *badger.Txn) error {
			return txn.Set(h[:], blk[:])
		},
	); log.E.Chk(err) {
		return err
	}

	return
}

func (idx *Index) Delete(h hash.Hash) (err error) {

	return
}

func (idx *Index) GetByHash(h hash.Hash) (b *block.Block, err error) {

	var blk []byte

	key := h[:]
	if err = idx.db.View(
		func(txn *badger.Txn) (err error) {
			var item *badger.Item
			if item, err = txn.Get(key); log.E.Chk(err) {
				if blk, err = item.ValueCopy(nil); log.E.Chk(err) {
					return
				}
			}
			return
		},
	); log.E.Chk(err) {
		return
	}

	var wireBlock block.WireBlock
	if wireBlock, err = block.ToWireBlock(blk); log.E.Chk(err) {
		return nil, err
	}
	rb := wireBlock.Unmarshal()
	b = &rb
	return
}

func (idx *Index) GetByHeight(h int) (b *block.Block) {

	return
}

func (idx *Index) GetBestChain() (bestChain Chain) {

	return
}

func (idx *Index) Head() (b *block.Block) {

	return
}
