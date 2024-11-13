package models

import (
    "time"
)

type OrderStatus string

const (
    OrderStatusPending   OrderStatus = "pending"
    OrderStatusPaid     OrderStatus = "paid"
    OrderStatusShipped  OrderStatus = "shipped"
    OrderStatusDelivered OrderStatus = "delivered"
    OrderStatusCanceled OrderStatus = "canceled"
)

type OrderItem struct {
    ProductID string  `json:"product_id"`
    Quantity  int     `json:"quantity"`
    Price     float64 `json:"price"`
    Subtotal  float64 `json:"subtotal"`
}

type Order struct {
    ID            string      `json:"id"`
    UserID        string      `json:"user_id"`
    Items         []OrderItem `json:"items"`
    TotalAmount   float64     `json:"total_amount"`
    Status        OrderStatus `json:"status"`
    ShippingAddr  Address     `json:"shipping_address"`
    BillingAddr   Address     `json:"billing_address"`
    PaymentMethod string      `json:"payment_method"`
    CreatedAt     time.Time   `json:"created_at"`
    UpdatedAt     time.Time   `json:"updated_at"`
}

type Address struct {
    Street     string `json:"street"`
    City       string `json:"city"`
    State      string `json:"state"`
    PostalCode string `json:"postal_code"`
    Country    string `json:"country"`
}