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
