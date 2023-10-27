package views

import "html/template"

type View struct {
	Template *template.Template
	Layout   string
}

func NewView(Layout string, files ...string) *View {
	files = append(files,
		"views/layouts/footer.html",
		"views/layouts/bootstrap.html",
		"views/layouts/navbar.html",
	)

	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}
	return &View{
		Template: t,
		Layout:   Layout,
	}
}
