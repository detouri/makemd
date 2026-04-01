package markdown

type Section struct {
	Level   int
	Title   string
	Body    string
	Hidden  bool
	Comment string
}

func NewSection(title, body string) Section {
	return Section{Level: 2, Title: title, Body: body}
}
