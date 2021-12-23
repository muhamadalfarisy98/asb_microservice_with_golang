package middleware

import (
	"fmt"
	"net/http"
)

// MIDDLEWARE : BASIC AUTH
func Auth(w http.ResponseWriter, r *http.Request) bool {

	username, password, ok := r.BasicAuth()
	stat := false
	// tidak masuk uname pass
	if !ok {
		w.Write([]byte("username dan password tidak boleh kosong"))
		stat = true
		return stat
	}
	// as admin
	if username == "admin" && password == "admin" {
		// next.ServeHTTP(w, r)
		return stat
	}
	// as editor
	if username == "editor" && password == "secret" {
		// next.ServeHTTP(w, r)
		return stat
	}
	// as trainer
	if username == "trainer" && password == "rahasia" {
		// next.ServeHTTP(w, r)
		return stat
	}
	stat = true
	// uname/ pass salah
	w.Write([]byte("username atau password salah \n"))
	fmt.Fprintf(w, "Ini dari middleware Log....\n")
	return stat
}
