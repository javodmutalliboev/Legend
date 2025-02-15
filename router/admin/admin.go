package admin

import (
	"Legend/middleware"
	"Legend/response"
	"Legend/shared"

	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func AdminRouter() *mux.Router {
	router := mux.NewRouter().PathPrefix("/api/v1/admin").Subrouter()

	router.HandleFunc("", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s: someone is requesting me", r.URL.Path)
		json.NewEncoder(w).Encode(response.Response{Status: "success", Code: http.StatusOK, Data: "admin route"})
	}).Methods("GET")

	router.HandleFunc("/login", Login()).Methods("POST")
	router.HandleFunc("/logout", middleware.Chain(Logout(), middleware.Auth())).Methods("GET")

	// POST /menu will create a new menu
	router.HandleFunc("/menu", middleware.Chain(CreateMenu(), middleware.Auth())).Methods("POST")

	// POST /menu/sub will create a new sub menu
	router.HandleFunc("/menu/sub", middleware.Chain(CreateSubMenu(), middleware.Auth())).Methods("POST")

	// GET /menus/{type:[0-9]+} will return all menus of a type
	router.HandleFunc("/menus/{type:[0-9]+}", middleware.Chain(shared.GetMenus(), middleware.Auth())).Methods("GET")

	// PATCH /menu will update a menu
	router.HandleFunc("/menu", middleware.Chain(UpdateMenu(), middleware.Auth())).Methods("PATCH")

	// DELETE /menu/{id:[0-9]+} will delete a menu
	router.HandleFunc("/menu/{id:[0-9]+}", middleware.Chain(DeleteMenu(), middleware.Auth())).Methods("DELETE")

	// GET /menu/types will return all menu types
	router.HandleFunc("/menu/types", middleware.Chain(GetMenuTypes(), middleware.Auth())).Methods("GET")

	// POST /goods will create a new goods
	router.HandleFunc("/goods", middleware.Chain(CreateGoods(), middleware.Auth())).Methods("POST")

	// GET /goods/{menu_id:[0-9]+} will return all goods of a menu
	router.HandleFunc("/goods/{menu_id:[0-9]+}", middleware.Chain(shared.GetGoods(), middleware.Auth())).Methods("GET")

	// PATCH /goods will update a goods
	router.HandleFunc("/goods", middleware.Chain(UpdateGoods(), middleware.Auth())).Methods("PATCH")

	// DELETE /goods/{id:[0-9]+} will delete a goods
	router.HandleFunc("/goods/{id:[0-9]+}", middleware.Chain(DeleteGoods(), middleware.Auth())).Methods("DELETE")

	// GET /goods/photo/{id:[0-9]+} will return a photo of a goods
	router.HandleFunc("/goods/photo/{id:[0-9]+}", middleware.Chain(shared.GetGoodsPhoto(), middleware.Auth())).Methods("GET")

	// DELETE /goods/photo/{id:[0-9]+} will delete a photo of a goods
	router.HandleFunc("/goods/photo/{id:[0-9]+}", middleware.Chain(DeleteGoodsPhoto(), middleware.Auth())).Methods("DELETE")

	// GET /goods/{id:[0-9]+} will return a goods
	router.HandleFunc("/goods/id/{id:[0-9]+}", middleware.Chain(shared.GetGoodsByID(), middleware.Auth())).Methods("GET")

	// POST /goods/photos will upload photos of a goods
	router.HandleFunc("/goods/photos", middleware.Chain(UploadGoodsPhotos(), middleware.Auth())).Methods("POST")

	// POST /general_discount will create a new general discount
	router.HandleFunc("/general_discount", middleware.Chain(CreateGeneralDiscount(), middleware.Auth())).Methods("POST")

	// GET /general_discount/{menu_type:[0-9]+} will return a general discount
	router.HandleFunc("/general_discount/{menu_type:[0-9]+}", middleware.Chain(shared.GetGeneralDiscount(), middleware.Auth())).Methods("GET")

	// PATCH /general_discount will update a general discount
	router.HandleFunc("/general_discount", middleware.Chain(UpdateGeneralDiscount(), middleware.Auth())).Methods("PATCH")

	// DELETE /general_discount/{menu_type:[0-9]+} will delete a general discount
	router.HandleFunc("/general_discount/{menu_type:[0-9]+}", middleware.Chain(DeleteGeneralDiscount(), middleware.Auth())).Methods("DELETE")

	// POST /legend_information will create the legend information
	router.HandleFunc("/legend_information", middleware.Chain(CreateLegendInformation(), middleware.Auth())).Methods("POST")

	// GET /legend_information will get the legend information
	router.HandleFunc("/legend_information", middleware.Chain(shared.GetLegendInformation(), middleware.Auth())).Methods("GET")

	// PATCH /legend_information will update the legend information
	router.HandleFunc("/legend_information", middleware.Chain(UpdateLegendInformation(), middleware.Auth())).Methods("PATCH")

	// DELETE /legend_information//{id:[0-9]+} will delete the legend information
	router.HandleFunc("/legend_information/{id:[0-9]+}", middleware.Chain(DeleteLegendInformation(), middleware.Auth())).Methods("DELETE")

	// POST /ctw_information will create the ctw information
	router.HandleFunc("/ctw_information", middleware.Chain(CreateCTWInformation(), middleware.Auth())).Methods("POST")

	// GET /ctw_information will get the ctw information
	router.HandleFunc("/ctw_information", middleware.Chain(shared.GetCTWInformation(), middleware.Auth())).Methods("GET")

	// PATCH /ctw_information will update the ctw information
	router.HandleFunc("/ctw_information", middleware.Chain(UpdateCTWInformation(), middleware.Auth())).Methods("PATCH")

	// DELETE /ctw_information//{id:[0-9]+} will delete the ctw information
	router.HandleFunc("/ctw_information/{id:[0-9]+}", middleware.Chain(DeleteCTWInformation(), middleware.Auth())).Methods("DELETE")

	// GET /goods/search/{menu_type:[0-9]+} will search goods by keyword
	router.HandleFunc("/goods/search/{menu_type:[0-9]+}", middleware.Chain(shared.SearchGoods(), middleware.Auth())).Methods("GET")

	// GET /goods/with_discount/{menu_type:[0-9]+} will return goods with discount
	router.HandleFunc("/goods/with_discount/{menu_type:[0-9]+}", middleware.Chain(shared.GetGoodsWithDiscount(), middleware.Auth())).Methods("GET")

	// GET /orders?page={page:[0-9]+}&limit={limit:[0-9]+} will return orders
	router.HandleFunc("/orders", middleware.Chain(GetOrders(), middleware.Auth())).Methods("GET")

	// PATCH /order/canceled will update the canceled status of an order
	router.HandleFunc("/order/canceled", middleware.Chain(UpdateOrderCanceled(), middleware.Auth())).Methods("PATCH")

	// PATCH /order/delivered will update the delivered status of an order
	router.HandleFunc("/order/delivered", middleware.Chain(UpdateOrderDelivered(), middleware.Auth())).Methods("PATCH")

	// POST /payment_method will create a new payment method
	router.HandleFunc("/payment_method", middleware.Chain(CreatePaymentMethod(), middleware.Auth())).Methods("POST")

	// GET /payment_method/list will return all payment methods
	router.HandleFunc("/payment_method/list", middleware.Chain(shared.GetPaymentMethods(), middleware.Auth())).Methods("GET")

	// GET /payment_method/{id:[0-9]+} will return a payment method
	router.HandleFunc("/payment_method/{id:[0-9]+}", middleware.Chain(GetPaymentMethod(), middleware.Auth())).Methods("GET")

	// GET /payment_method/{id:[0-9]+}/logo will return the logo of a payment method
	router.HandleFunc("/payment_method/{id:[0-9]+}/logo", middleware.Chain(shared.GetPaymentMethodLogo(), middleware.Auth())).Methods("GET")

	// PATCH /payment_method will update a payment method
	router.HandleFunc("/payment_method", middleware.Chain(UpdatePaymentMethod(), middleware.Auth())).Methods("PATCH")

	// DELETE /payment_method/{id:[0-9]+} will delete a payment method
	router.HandleFunc("/payment_method/{id:[0-9]+}", middleware.Chain(DeletePaymentMethod(), middleware.Auth())).Methods("DELETE")

	// POST /social_network will create a new social network
	router.HandleFunc("/social_network", middleware.Chain(CreateSocialNetwork(), middleware.Auth())).Methods("POST")

	// GET /social_network/list will return all social networks
	router.HandleFunc("/social_network/list", middleware.Chain(shared.GetSocialNetworks(), middleware.Auth())).Methods("GET")

	// GET /social_network/{id:[0-9]+} will return a social network
	router.HandleFunc("/social_network/{id:[0-9]+}", middleware.Chain(GetSocialNetwork(), middleware.Auth())).Methods("GET")

	// GET /social_network/{id:[0-9]+}/icon will return the icon of a social network
	router.HandleFunc("/social_network/{id:[0-9]+}/icon", middleware.Chain(shared.GetSocialNetworkIcon(), middleware.Auth())).Methods("GET")

	// PATCH /social_network will update a social network
	router.HandleFunc("/social_network", middleware.Chain(UpdateSocialNetwork(), middleware.Auth())).Methods("PATCH")

	// DELETE /social_network/{id:[0-9]+} will delete a social network
	router.HandleFunc("/social_network/{id:[0-9]+}", middleware.Chain(DeleteSocialNetwork(), middleware.Auth())).Methods("DELETE")

	return router
}
