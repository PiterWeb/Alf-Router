package templates

import (
	"html/template"

	alf "github.com/PiterWeb/Alf-Router"
)

func Templates(templatesFolder string, extension string) func(ctx *alf.Ctx, templateName string, data any) error {

	tpl := template.Must(template.ParseGlob(templatesFolder + "/*"+ extension))

	return func(ctx *alf.Ctx, templateName string, data any) error {
		ctx.SetContentType("text/html")
		return tpl.ExecuteTemplate(ctx, templateName + extension, data)
	}

}
