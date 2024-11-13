package handlers

import (
    "encoding/json"
    "net/http"
    "time"
    "github.com/gorilla/mux"
    "go-rest-api/internal/models"
    "go-rest-api/internal/utils"
)

var orders = make(map[string]models.Order)

func CreateOrder(w http.ResponseWriter, r *http.Request) {
    var order models.Order
    if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
        utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
        return
    }

    // Validate items and calculate total
    var total float64
    for i, item := range order.Items {
        product, exists := products[item.ProductID]
        if !exists {
            utils.RespondWithError(w, http.StatusBadRequest, "Invalid product ID")
            return
        }
        
        if product.Stock < item.Quantity {
            utils.RespondWithError(w, http.StatusBadRequest, "Insufficient stock")
            return
        }

        // Update item price and subtotal
        order.Items[i].Price = product.Price
        order.Items[i].Subtotal = product.Price * float64(item.Quantity)
        total += order.Items[i].Subtotal

        // Update product stock
        product.Stock -= item.Quantity
        products[item.ProductID] = product
    }

    order.ID = utils.GenerateUUID()
    order.TotalAmount = total
    order.Status = models.OrderStatusPending
    order.CreatedAt = time.Now()
    order.UpdatedAt = time.Now()

    orders[order.ID] = order

    utils.RespondWithJSON(w, http.StatusCreated, order)
}

func UpdateOrderStatus(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id := params["id"]

    order, exists := orders[id]
    if !exists {
        utils.RespondWithError(w, http.StatusNotFound, "Order not found")
        return
    }

    var req struct {
        Status models.OrderStatus `json:"status"`
    }
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
        return
    }

    order.Status = req.Status
    order.UpdatedAt = time.Now()
    orders[id] = order

    utils.RespondWithJSON(w, http.StatusOK, order)
}