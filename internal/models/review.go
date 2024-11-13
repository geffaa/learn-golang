package models

import (
    "time"
)

type Review struct {
    ID        string    `json:"id"`
    ProductID string    `json:"product_id"`
    UserID    string    `json:"user_id"`
    Rating    int       `json:"rating"`
    Comment   string    `json:"comment"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}