package controllers

import (
	"asb_microservice_with_golang/utils"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// go routine
func HitungPrice(price, discount, tax int, ch chan int, hitung string) {
	if hitung == "pajak" {
		ch <- price * (100 + tax) / 100
	} else if hitung == "discount" {
		ch <- price * (100 - tax) / 100
	}
}

func QueryHitungPrice(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	price_str := r.FormValue("price")
	price, _ := strconv.Atoi(price_str)

	tax_str := r.FormValue("tax")
	tax, _ := strconv.Atoi(tax_str)

	discount_str := r.FormValue("discount")
	discount, _ := strconv.Atoi(discount_str)

	var hitung = r.FormValue("hitung")

	var ch6 = make(chan int)
	cal := 0
	go HitungPrice(price, discount, tax, ch6, hitung)
	if hitung == "pajak" {
		cal = <-ch6
	} else if hitung == "discount" {
		cal = <-ch6
	}
	res := cal
	utils.ResponseJSON(w, res, http.StatusOK)
}
