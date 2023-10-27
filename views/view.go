package views

import (
	"html/template"
	"path/filepath"
)

var (
	LayoutDir   string = "views/layouts/"
	TemplateExt string = ".html"
)

type View struct {
	Template *template.Template
	Layout   string
}

func NewView(Layout string, files ...string) *View {
	files = append(files, layoutFiles()...)

	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}
	return &View{
		Template: t,
		Layout:   Layout,
	}
}

// layoutFiles return a slice of string representing
// the layout files used in our web app
func layoutFiles() []string {
	files, err := filepath.Glob(LayoutDir + "*" + ".html")
	if err != nil {
		panic(err)
	}
	return files
}
