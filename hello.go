package staticbug

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func listDir(name string) string {
	files, err := ioutil.ReadDir(name)
	if err == nil {
		fns := []string{}
		for _, f := range files {
			fns = append(fns, f.Name())
		}
		return strings.Join(fns, " ")
	}
	return fmt.Sprintf("%s: %v", name, err)
}

func init() {
	http.HandleFunc("/ls", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, listDir("."))
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		f, err := os.Open("static/index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			defer f.Close()
			io.Copy(w, f)
		}
	})
}
