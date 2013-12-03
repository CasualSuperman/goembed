/* Serving resource from zip file appended to Go executable (this enables on file deploy).

Making it happen:
	1. Add code to serve resources (see example below)
	2. Compile your executable
	3. Run the `goembed` script from https://github.com/CasualSuperman/goembed/goembed
	4. Deploy

Example code:

	package main

	import (
		"fmt"
		"net/http"
		"os"

		"github.com/CasualSuperman/goembed"
	)

	type params struct {
		Number  int
	}

	func indexHandler(w http.ResponseWriter, req *http.Request) {
		t, err := goembed.LoadTemplates(nil, "t.html")
		if err != nil {
			http.NotFound(w, req)
		}
		if err = t.Execute(w, params{7}); err != nil {
			http.NotFound(w, req)
		}
	}

	func main() {
		goembed.Handle("/static/")
		http.HandleFunc("/", indexHandler)
		if err := http.ListenAndServe(":8080", nil); err != nil {
			fmt.Fprintf(os.Stderr, "error: %s\n", err)
			os.Exit(1)
		}
	}
*/
package goembed
