package nebula

import (
	"github.com/benpate/datatype"
	"github.com/benpate/derp"
)

type NilAction datatype.Map

func (txn NilAction) Execute(library *Library, container *Container) (int, error) {
	return 0, derp.New(500, "content.NilAction", "Unrecognized Action Type", txn)
}
