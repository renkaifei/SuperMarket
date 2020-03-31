package main

import (
	"log"
	"net/http"
	"superMarket/controllers"
	"superMarket/middlewares"
)

func main() {
	http.Handle("/wxtxt/", http.StripPrefix("/wxtxt/", http.FileServer(http.Dir("wxtxt"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.Handle("/views/", http.StripPrefix("/views/", http.FileServer(http.Dir("views"))))

	goodsController := &controllers.GoodsController{}

	http.HandleFunc("/goods/create", MyHandlerFunc(goodsController.Create))
	http.HandleFunc("/goods/update", MyHandlerFunc(goodsController.Update))
	http.HandleFunc("/goods/delete", MyHandlerFunc(goodsController.Delete))
	http.HandleFunc("/goods/selectById", MyHandlerFunc(goodsController.SelectById))
	http.HandleFunc("/goods/selectInMerchant", MyHandlerFunc(goodsController.SelectInMerchant))

	wxController := &controllers.WxController{}
	http.HandleFunc("/wx/ListenMessage", wxController.ListenMessage)

	merchanterController := &controllers.MerchanterController{}
	http.HandleFunc("/login", merchanterController.Login)

	merchantController := &controllers.MerchantController{}
	http.HandleFunc("/merchant/selectById", MyHandlerFunc(merchantController.SelectById))

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func MyHandlerFunc(handleFunc func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	initMiddleware := &middlewares.InitMiddleware{}
	sessionMiddleware := &middlewares.SessionMiddleware{}
	initMiddleware.NextTask = sessionMiddleware.HandleFunc
	sessionMiddleware.NextTask = handleFunc
	return initMiddleware.HandleFunc
}
