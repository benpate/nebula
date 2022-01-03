package transaction

import (
	"github.com/benpate/datatype"
	"github.com/benpate/derp"
	"github.com/benpate/nebula"
)

type NilTransaction datatype.Map

func (txn NilTransaction) Execute(library *nebula.Library, container *nebula.Container) (int, error) {
	return 0, derp.New(500, "content.NilTransaction", "Unrecognized Transaction Type", txn)
}
