package main

import (
	"html/template"
	"io/ioutil"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	ls, err := ioutil.ReadDir(".")
	if err != nil {
		http.Error(w, "Cannot read directory", http.StatusInternalServerError)
		return
	}
	t := template.Must(template.New("index").Parse(displayPage))
	t.Execute(w, ls)
}

func main() {
	http.Handle("/", http.HandlerFunc(IndexHandler))
	http.Handle("/imgs/", http.StripPrefix("/imgs/", http.FileServer(http.Dir("."))))
	http.ListenAndServe(":8080", nil)
}

var displayPage = `<html>
<body>
    {{ range . }}
    <figure><img src="/imgs/{{ .Name }}" /></figure>
    {{ end }}
</body>
</html>`
