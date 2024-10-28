package command

type CreateContentCommand struct {
	Name        string
	Description string
	Type        string
	ParentID    string
}
