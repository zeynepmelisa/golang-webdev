package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

type Hello struct {
	Sale string
	Time string
}

func main() {
	/* initialize the data fields contained within the struct */
	hello := Hello{"Sale Commences Now", time.Now().Format(time.Stamp)}

	/* Must() -> used to test the validity of the template during parsing */
	outline := template.Must(template.ParseFiles("outline/outline.html"))

	/* configure a route that serves static files from the "steady" directory */
	/* when a user accesses a URL with a path starting with /steady/, this handler
	removes the /steady/ prefix and serves files from the "steady" directory */
	http.Handle("/steady/", http.StripPrefix("/steady/", http.FileServer(http.Dir("steady"))))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if sale := r.FormValue("sale"); sale != "" {
			hello.Sale = sale
		}

		/* it stops applying the templatr if there is any error */
		/* it is taking into consideration  the template that has been described in
		outline.html */
		if err := outline.ExecuteTemplate(w, "outline.html", hello); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println(http.ListenAndServe(":8081", nil))

}
