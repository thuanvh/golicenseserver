package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CustomerHandler struct{}

func CreateCustomerHandler(r *gin.RouterGroup) {
	// Create an account group
	// Create a handler (which will later have injected services)
	h := &CustomerHandler{} // currently has no properties

	// Create a group, or base url for all routes
	g := r.Group("/customers")

	g.GET("/", h.List)
	g.POST("/", h.NewLicense)
	g.GET("/:id", h.GetLicense)
	g.POST("/signout", h.UpdateLicense)
	g.POST("/tokens", h.DeleteLicense)
}

// sample::handler::all_licenses,
// sample::handler::create_license,
// sample::handler::get_license,
// sample::handler::update_license,
// sample::handler::delete_license
// ],
// )
// .mount("/customers",
// routes![
// sample::handler::all_customers,
// sample::handler::create_customer,
// sample::handler::get_customer,
// sample::handler::update_customer,
// sample::handler::delete_customer
// ],
// )
// .mount("/customer_licenses",
// routes![
// sample::handler::all_customer_licenses,
// sample::handler::create_customer_license,
// sample::handler::get_customer_license,
// sample::handler::get_customer_license_by_customer_id,
// sample::handler::get_customer_license_by_license_id,
// sample::handler::update_customer_license,
// sample::handler::delete_customer_license
func (h *CustomerHandler) List(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"list": "list ok",
	})
}

func (h *CustomerHandler) NewLicense(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"list": "list ok",
	})
}

func (h *CustomerHandler) UpdateLicense(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"list": "list ok",
	})
}

func (h *CustomerHandler) GetLicense(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"list": "list ok",
	})
}

func (h *CustomerHandler) DeleteLicense(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"list": "list ok",
	})
}
