package orders

import (
	"net/http"

	"github.com/qor/render"
)

// Controller products controller
type Controller struct {
	View *render.Render
}

// Cart shopping cart
func (ctrl Controller) Cart(w http.ResponseWriter, req *http.Request) {
	ctrl.View.Execute("cart", map[string]interface{}{}, req, w)
}

// UpdateCart update shopping cart
func (ctrl Controller) UpdateCart(w http.ResponseWriter, req *http.Request) {
}

// Checkout checkout shopping cart
func (ctrl Controller) Checkout(w http.ResponseWriter, req *http.Request) {
	ctrl.View.Execute("checkout", map[string]interface{}{}, req, w)
}

// Complete complete shopping cart
func (ctrl Controller) Complete(w http.ResponseWriter, req *http.Request) {
}

// CheckoutSuccess checkout success page
func (ctrl Controller) CheckoutSuccess(w http.ResponseWriter, req *http.Request) {
	ctrl.View.Execute("success", map[string]interface{}{}, req, w)
}
