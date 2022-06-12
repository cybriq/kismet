package chain

import (
	"github.com/cybriq/kismet/pkg/block"
	"github.com/cybriq/kismet/pkg/hash"
	"sync"
)

type Index struct {
	loaded map[hash.Hash]bool
	blocks map[hash.Hash]*block.Block
	head   hash.Hash
	mx     sync.Mutex
}

// Load retrieves a block if it is not loaded
func (idx *Index) Load(h hash.Hash) (found bool) {
	loaded, ok := idx.loaded[h]
	if ok {
		if !loaded {
			// load it from database
		} else {
			found = true
		}
	}
	return
}

func (idx *Index) Add(b *block.Block) (err error) {
	var h hash.Hash
	h, err = b.IndexHash()
	if err != nil {
		return
	}
	idx.loaded[h] = true
	idx.blocks[h] = b
	// save to database
	return
}

func (idx *Index) GetByHash(h hash.Hash) (b *block.Block, found bool) {

	b, found = idx.blocks[h]
	return
}
