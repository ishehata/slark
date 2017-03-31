package slark

import (
	"html/template"
	"log"
	"net/http"
	"path"
	"runtime"

	"fmt"

	"github.com/emostafa/garson"
	"github.com/jmoiron/sqlx"
)

func handleError(err error, w http.ResponseWriter) {
	if err != nil {
		log.Print(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
}

func index(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func list(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		params := ctx.Value("route_params").(garson.Params)
		modelName := params["model"]
		model := models[modelName]
		q := fmt.Sprintf("SELECT * FROM %s LIMIT 20", model.TableName)
		rows, err := db.Queryx(q)
		handleError(err, w)
		// cols, err := rows.Columns()
		// handleError(err, w)

		results := []interface{}{}
		for rows.Next() {
			result := make(map[string]interface{})
			// for _, col := range cols {
			// 	result[col] = new(interface{})
			// }
			rows.MapScan(result)

			results = append(results, result)
		}

		_, filename, _, ok := runtime.Caller(0)
		if !ok {
			log.Fatal("No caller information")
		}
		p := path.Join(path.Dir(filename), "templates/list.html")
		t, err := template.ParseFiles(p)
		handleError(err, w)
		t.Execute(w, results)
	}
}

func read(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
func getCreate(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
func postCreate(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
func getUpdate(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
func postUpdate(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
func delete(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
