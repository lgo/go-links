package endpoint

import (
	"net/http"

	"github.com/aymerick/raymond"
	"github.com/xLegoz/go-links/backend"
)

//
// var template *raymond.Template
//
// func init() {
// 	// parse template
// 	var err error
// 	template, err = raymond.Parse("templates/admin.html.mustache")
// 	if err != nil {
// 		panic(err)
// 	}
// }

/**
 * AdminDashboardHandler
 * @type {String}
 */
func AdminDashboardHandler(w http.ResponseWriter, r *http.Request) {
	links, err := backend.ActiveBackend.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	ctx := map[string]interface{}{
		"links": links,
	}

	template, err := raymond.ParseFile("templates/admin.html.mustache")
	if err != nil {
		panic(err)
	}

	// render template
	result, err := template.Exec(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Write([]byte(result))
}
