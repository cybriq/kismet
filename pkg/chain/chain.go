package chain

import "github.com/dgraph-io/badger/v3"

type Chain struct {
	*badger.DB
}
