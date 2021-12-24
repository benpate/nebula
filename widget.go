package content

type Widget interface {
	Draw(Content) string
}

type EditorWidget interface {
	ItemTypes() []ItemType
	Widget
}

type ItemType struct {
	Code        string
	Label       string
	Description string
}
