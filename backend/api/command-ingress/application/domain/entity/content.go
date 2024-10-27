package entity

type Content struct {
	ID          string
	Name        string
	Description string
	Type        string
	Children    []Content
}
