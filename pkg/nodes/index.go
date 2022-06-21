package nodes

import "github.com/dgraph-io/badger/v3"

type Index struct {
	*badger.DB
}
