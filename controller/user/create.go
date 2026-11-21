package user

import (
	"fmt"
	"net/http"
)

func CreateHandlerFunc(w http.ResponseWriter, r *http.Request) {
	name, lastname := r.FormValue("name"), r.FormValue("lastname")
	if name == "" {
		fmt.Fprintf(w, "The field \"name\" can not be empty.")
		return
	}
	if lastname == "" {
		fmt.Fprintf(w, "The field \"lastname\" can not be empty.")
		return
	}
	fmt.Fprintf(w, "User with full-name {%s %s} has been created.", r.FormValue("name"), r.FormValue("lastname"))
}
