package utils

import (
	"encoding/json"
	"net/http"
)

/*
	menampilkan data dengan bentuk json di browser
	fungsi ini akan dipanggildd di main.go nanti
*/
func ResponseJSON(w http.ResponseWriter, p interface{}, status int) {
	ubahkeByte, err := json.Marshal(p)
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		http.Error(w, "error om", http.StatusBadRequest)
	}

	// buat statusnya
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(ubahkeByte))
}
