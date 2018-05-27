package render

import (
	"fmt"
	"html/template"
	"net/http"
)

func HTMLRender (f string, w http.ResponseWriter, data interface{}) error {
	t, err := template.ParseFiles(fmt.Sprintf("./tmpls/layout/%s", f))
	if err != nil {
		return err
	}

	return t.Execute(w, data)
}