package content

// Getter wraps the "GetContent" method, which allows an object
// to safely expose it's underlying content data
type Getter interface {
	GetContent() Content
}

// Setter wraps the "SetContent" method, which allows an object
// to safely expose it's underlying content data
type Setter interface {
	SetContent(Content)
}

// GetterSetter wraps both "GetContent" and "SetContent" methods,
// allowing read/write access to an object's underlying content data.
type GetterSetter interface {
	Getter
	Setter
}
