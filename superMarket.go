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
	http.HandleFunc("/goods/selectOnePage", MyHandlerFunc(goodsController.SelectOnePage))
	http.HandleFunc("/goods/selectByBarCode", MyHandlerFunc(goodsController.SelectByBarCode))

	wxController := &controllers.WxController{}
	http.HandleFunc("/wx/ListenMessage", wxController.ListenMessage)
	http.HandleFunc("/wx/QueryConfig", MyHandlerFunc2(wxController.QueryConfig))

	merchanterController := &controllers.MerchanterController{}
	http.HandleFunc("/login", merchanterController.Login)

	merchantController := &controllers.MerchantController{}
	http.HandleFunc("/merchant/selectById", MyHandlerFunc(merchantController.SelectById))

	goodsCategory := &controllers.GoodsCategoryController{}
	http.HandleFunc("/goodsCategory/selectByName", MyHandlerFunc(goodsCategory.SelectByName))

	merchantGoods := &controllers.MerchantGoodsController{}
	http.HandleFunc("/merchantGoods/selectByMerchantIdAndGoodsId", MyHandlerFunc(merchantGoods.SelectByMerchantIdAndGoodsId))
	http.HandleFunc("/merchantGoods/create", MyHandlerFunc(merchantGoods.Create))
	http.HandleFunc("/merchantGoods/update", MyHandlerFunc(merchantGoods.Update))
	http.HandleFunc("/merchantGoods/delete", MyHandlerFunc(merchantGoods.Delete))
	http.HandleFunc("/merchantGoods/selectPageByMerchantId", MyHandlerFunc(merchantGoods.SelectPageByMerchantId))
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

func MyHandlerFunc2(handleFunc func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	initMiddleware := &middlewares.InitMiddleware{}
	initMiddleware.NextTask = handleFunc
	return initMiddleware.HandleFunc
}
