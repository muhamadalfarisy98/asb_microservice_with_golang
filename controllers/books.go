package controllers

import (
	"asb_microservice_with_golang/book"
	"asb_microservice_with_golang/middleware"
	"asb_microservice_with_golang/models"
	"asb_microservice_with_golang/utils"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
)

// --controller : CRUD book --
// Read
// GetBook
func GetBooks(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// parsing ParamsQuery
	title := r.FormValue("title")
	minYear := r.FormValue("minYear")
	maxYear := r.FormValue("maxYear")
	minPage := r.FormValue("minPage")
	maxPage := r.FormValue("maxPage")
	sortByTitle := r.FormValue("sortByTitle")

	fmt.Println(title)
	data := map[string]string{
		"title":       title,
		"minyear":     minYear,
		"maxyear":     maxYear,
		"minpage":     minPage,
		"maxpage":     maxPage,
		"sortbytitle": sortByTitle,
	}
	nilai, err := book.GetAll(ctx, data)

	if err != nil {
		fmt.Println(err)
	}

	utils.ResponseJSON(w, nilai, http.StatusOK)
}

func GetBooksByIDCategory(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// parsing ParamsQuery
	title := r.FormValue("title")
	minYear := r.FormValue("minYear")
	maxYear := r.FormValue("maxYear")
	minPage := r.FormValue("minPage")
	maxPage := r.FormValue("maxPage")
	sortByTitle := r.FormValue("sortByTitle")

	data := map[string]string{
		"title":       title,
		"minyear":     minYear,
		"maxyear":     maxYear,
		"minpage":     minPage,
		"maxpage":     maxPage,
		"sortbytitle": sortByTitle,
	}
	// fmt.Println(data)
	var idCateg = ps.ByName("id")
	nilai, err := book.GetBookFilterByCategory(ctx, idCateg, data)

	if err != nil {
		fmt.Println(err)
	}

	utils.ResponseJSON(w, nilai, http.StatusOK)
}

// Create
// PostBooks
func PostBooks(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}
	// MIDDLEWARE : BASIC AUTH
	auth := middleware.Auth(w, r)
	if auth == true {
		return
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var buku models.Book
	if err := json.NewDecoder(r.Body).Decode(&buku); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	// VALIDASI
	is_release_year := false
	is_image_url := false
	// validasi release year
	if (buku.ReleaseYear >= 1980) && (buku.ReleaseYear <= 2021) {
		is_release_year = true
	}
	// validasi image_url
	/*
		ASUMSI

		var url1,url2,url3 string
		url1 = "/image.png" // not accepted
		url2 = "http://abc/image.jpg"  // accepted
		url3 = "https://abc/image.jpg" // accepted
	*/
	if strings.HasPrefix(buku.ImageUrl, "http://") || strings.HasPrefix(buku.ImageUrl, "https://") {
		// Valid URL
		is_image_url = true
	}
	// cek validasi
	if is_release_year == false {
		if is_image_url == false {
			http.Error(w, "url tidak mengandung prefix http:// ataupun https://, dan Release year tidak berada diantara tahun 1980 dan 2021", http.StatusBadRequest)
			return
		}
		http.Error(w, "Release year tidak berada diantara tahun 1980 dan 2021", http.StatusBadRequest)
		return
	} else if is_image_url == false {
		http.Error(w, "url tidak mengandung prefix http:// ataupun https://", http.StatusBadRequest)
		return
	}

	// konversi thickness
	if buku.TotalPage <= 100 {
		buku.Thickness = "tipis"
	} else if (buku.TotalPage > 100) && (buku.TotalPage <= 200) {
		buku.Thickness = "sedang"
	} else if buku.TotalPage > 200 {
		buku.Thickness = "tebal"
	}

	if err := book.Insert(ctx, buku); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}
	res := map[string]string{
		"status": "Data berhasil ditambahkan",
	}
	utils.ResponseJSON(w, res, http.StatusCreated)

}

// Update
// UpdateBooks
func UpdateBooks(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}
	// MIDDLEWARE : BASIC AUTH
	auth := middleware.Auth(w, r)
	if auth == true {
		return
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var buku models.Book

	if err := json.NewDecoder(r.Body).Decode(&buku); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	// VALIDASI
	is_release_year := false
	is_image_url := false
	// validasi release year
	if (buku.ReleaseYear >= 1980) && (buku.ReleaseYear <= 2021) {
		is_release_year = true
	}
	// validasi image_url
	/*
		ASUMSI

		var url1,url2,url3 string
		url1 = "/image.png" // not accepted
		url2 = "http://abc/image.jpg"  // accepted
		url3 = "https://abc/image.jpg" // accepted
	*/
	if strings.HasPrefix(buku.ImageUrl, "http://") || strings.HasPrefix(buku.ImageUrl, "https://") {
		// Valid URL
		is_image_url = true
	}
	// cek validasi
	if is_release_year == false {
		if is_image_url == false {
			http.Error(w, "url tidak mengandung prefix http:// ataupun https://, dan Release year tidak berada diantara tahun 1980 dan 2021", http.StatusBadRequest)
			return
		}
		http.Error(w, "Release year tidak berada diantara tahun 1980 dan 2021", http.StatusBadRequest)
		return
	} else if is_image_url == false {
		http.Error(w, "url tidak mengandung prefix http:// ataupun https://", http.StatusBadRequest)
		return
	}

	var idBuku = ps.ByName("id")

	if err := book.Update(ctx, buku, idBuku); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Edit data berhasil",
	}

	utils.ResponseJSON(w, res, http.StatusCreated)
}

// Delete
// DeleteBooks
func DeleteBooks(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	// MIDDLEWARE : BASIC AUTH
	auth := middleware.Auth(w, r)
	if auth == true {
		return
	}
	defer cancel()
	var idBook = ps.ByName("id")
	if err := book.Delete(ctx, idBook); err != nil {
		kesalahan := map[string]string{
			"error": fmt.Sprintf("%v", err),
		}
		utils.ResponseJSON(w, kesalahan, http.StatusInternalServerError)
		return
	}
	res := map[string]string{
		"status": "data Berhasil dihapus",
	}
	utils.ResponseJSON(w, res, http.StatusOK)
}
