package render

import "html/template"

func RenderTemplate(path string) *template.Template {

	render := template.Must(template.ParseFiles(path))

	return render

}
