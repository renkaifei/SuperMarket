package main

import (
	"log"
	"net/http"
	"superMarket/controllers"
)

func main() {
	goodsController := &controllers.GoodsController{}
	http.HandleFunc("/goods/create", goodsController.Create)
	http.HandleFunc("/goods/update", goodsController.Update)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}
