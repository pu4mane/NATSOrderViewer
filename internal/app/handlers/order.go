package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/pu4mane/NATSOrderViewer/internal/app/cache"
	"github.com/pu4mane/NATSOrderViewer/internal/app/model"
)

func HandlerOrder(w http.ResponseWriter, r *http.Request) {
	orderIDStr := r.URL.Query().Get("id")
	orderID, err := strconv.Atoi(orderIDStr)
	if err != nil {
		http.Error(w, "Invalid order ID", http.StatusBadRequest)
		return
	}

	orderItem, err := cache.Get(fmt.Sprintf("order:%d", orderID))
	if err != nil {
		http.Error(w, "Order not found", http.StatusNotFound)
		return
	}

	order := model.Order{}
	if err := json.Unmarshal(orderItem, &order); err != nil {
		http.Error(w, "Error unmarshalling order", http.StatusInternalServerError)
		return
	}

	delivery := order.Delivery
	payment := order.Payment
	items := order.Items

	for i := range items {
		items[i].ID = i + 1
	}

	tmpl, err := template.ParseFiles("../orderpage.html")
	if err != nil {
		log.Fatal("Error parsing template:", err)
	}

	data := model.OrderData{
		Order:    order,
		Delivery: delivery,
		Payment:  payment,
		Items:    items,
	}

	if err := tmpl.Execute(w, data); err != nil {
		log.Fatal("Error executing template:", err)
	}
}

func HandlerIndex(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("../indexpage.html")
	if err != nil {
		log.Fatal("Error parsing template:", err)
	}

	if err := tmpl.Execute(w, nil); err != nil {
		log.Fatal("Error executing template:", err)
	}
}
