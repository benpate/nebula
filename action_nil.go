package nebula

import (
	"github.com/benpate/datatype"
	"github.com/benpate/derp"
)

type NilAction datatype.Map

func (txn NilAction) Get(library *Library, container *Container, endpoint string) string {
	return ""
}

func (txn NilAction) Post(library *Library, container *Container) (int, error) {
	return 0, derp.New(500, "content.NilAction", "Unrecognized Action Type", txn)
}
