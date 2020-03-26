package main

import (
	"log"
	"net/http"
	"superMarket/controllers"
)

func main() {
	http.Handle("/wxtxt/", http.StripPrefix("/wxtxt/", http.FileServer(http.Dir("wxtxt"))))

	goodsController := &controllers.GoodsController{}
	http.HandleFunc("/goods/create", goodsController.Create)
	http.HandleFunc("/goods/update", goodsController.Update)
	http.HandleFunc("/goods/delete", goodsController.Delete)
	http.HandleFunc("/goods/selectById", goodsController.SelectById)

	wxController := &controllers.WxController{}
	http.HandleFunc("/wx/ListenMessage", wxController.ListenMessage)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}
