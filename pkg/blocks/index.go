package blocks

import (
	"fmt"

	"github.com/cybriq/kismet/pkg/block"
	"github.com/cybriq/kismet/pkg/hash"
	"github.com/cybriq/qu"
	"github.com/dgraph-io/badger/v3"
)

type Index struct {
	*badger.DB
}

// New creates a new block index.
func New(path string, stop qu.C) (idx *Index, err error) {

	idx = &Index{}

	if idx.DB, err = badger.Open(badger.DefaultOptions(path)); log.E.Chk(err) {
		return
	}

	// when the stop channel is closed close the database
	go func() {
	out:
		select {
		case <-stop.Wait():

			err = idx.DB.Close()
			log.E.Chk(err)
			break out
		}
	}()

	return
}

// Add a new block to the index
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

	if err = idx.DB.View(
		func(txn *badger.Txn) (err error) {

			if _, err = txn.Get(h[:]); err == nil {
				return fmt.Errorf("block already exists, not adding again")
			}
			return
		},
	); log.E.Chk(err) {

		return
	}

	var blk []byte
	if blk, err = b.Marshal(); log.E.Chk(err) {
		return
	}

	err = idx.DB.Update(func(txn *badger.Txn) error {
		return txn.Set(h[:], blk[:])
	})
	log.E.Chk(err)

	return
}

// Delete a block from the database (for the case of as yet not defined pruning regime)
func (idx *Index) Delete(h hash.Hash) (err error) {

	err = idx.DB.Update(func(txn *badger.Txn) (err error) {
		return txn.Delete(h[:])
	})
	log.E.Chk(err)
	return
}

// GetByHash returns a block given its IndexHash
func (idx *Index) GetByHash(h hash.Hash) (b *block.Block, err error) {

	var blk []byte

	key := h[:]
	if err = idx.DB.View(
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

	err = b.Unmarshal(blk)
	log.E.Chk(err)
	return
}
