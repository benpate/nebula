package nebula

// Getter wraps the "GetContainer" method, which allows an object
// to safely expose it's underlying container data
type Getter interface {
	GetContainer() Container
}

// Setter wraps the "SetContent" method, which allows an object
// to safely expose it's underlying container data
type Setter interface {
	SetContainer(Container)
}

// GetterSetter wraps both "GetContainer" and "SetContainer" methods,
// allowing read/write access to an object's underlying container data.
type GetterSetter interface {
	Getter
	Setter
}

// Action interface wraps the "Execute" method, which takes some
// non-reversible, non-idempotent action on the container.
type Action interface {
	Get(*Library, *Container) string

	// Execute performs an update to the content data.  It returns
	// the ID of the element to be re-rendered on the client, along
	// with an error (if present).  Implementations can use this to
	// selectively re-render portions of the content structure without
	// reloading the entire page.
	Post(*Library, *Container) (int, error)

	// TODO: ExecuteMongo() -- generates an efficient update statement for a mongodb collection.
}
