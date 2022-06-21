package tree

import (
	"math/big"

	"github.com/cybriq/kismet/pkg/block"
	"github.com/cybriq/kismet/pkg/blocks"
	"github.com/cybriq/kismet/pkg/hash"
	"github.com/cybriq/qu"
	"github.com/dgraph-io/badger/v3"
)

type NodeIndex struct {
	Number int
	Hash   hash.Hash
	Weight big.Int
}

type Tree struct {
	*badger.DB

	stop qu.C
	// Index is the database of all known blocks
	*blocks.Index
	// Nodes maps Chain index values back to the Index of blocks
	Nodes []NodeIndex
	// Chain is the mapping between the NodeIndex.Number of a node to its
	// previous
	Chain []int
	// Head is the block that forms the heaviest chain, ie, has the lowest
	// NodeIndex.Weight
	Head int
}

// New starts up a new Tree
func New(path string, stop qu.C) (tr *Tree, err error) {

	tr = &Tree{}

	if tr.DB, err = badger.Open(badger.DefaultOptions(path)); log.E.Chk(err) {
		return
	}

	tr.stop = stop

	// when the stop channel is closed close the database
	go func() {
	out:
		select {
		case <-tr.stop.Wait():

			err = tr.DB.Close()
			log.E.Chk(err)
			break out
		}
	}()

	return
}

// Load constructs the in memory tables from what is stored in the DB
func (tr *Tree) Load() (err error) {

	return
}

func (tr *Tree) AddBlock(b *block.Block) (err error) {

	return
}
