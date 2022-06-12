package chain

import (
	"fmt"
	"github.com/cybriq/kismet/pkg/block"
	"github.com/cybriq/kismet/pkg/hash"
	"github.com/cybriq/qu"
	"github.com/dgraph-io/badger/v3"
	"time"
)

type AccessTrackedBlock struct {
	time.Time
	*block.Block
}

type Blocks map[hash.Hash]AccessTrackedBlock

type Index struct {
	db    *badger.DB
	cache Blocks
}

type IndexBlock struct {
	Index
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

func New(path string, maxToCache int, stop qu.C) (idx *Index, err error) {

	idx = &Index{cache: make(Blocks)}

	if idx.db, err = badger.Open(badger.DefaultOptions(path)); log.E.Chk(err) {
		return
	}

	// here we need a GC to expire the cache periodically if it gets huge
	go func() {

		timer := time.NewTicker(time.Minute)

	out:
		for {
			select {
			case <-timer.C:
				// do gc here
			case <-stop.Wait():
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
	if _, found := idx.cache[h]; found {

		err = fmt.Errorf(
			"block already exists in cache, not adding as" +
				" it has been seen/added already",
		)
		log.E.Ln(err)
		return
	}

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

		idx.cache[h] = AccessTrackedBlock{Time: time.Now(), Block: b}
		return
	}

	// add the block to the cache for fast retrieval
	idx.cache[h] = AccessTrackedBlock{Time: time.Now(), Block: b}

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
