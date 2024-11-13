package models

type Category struct {
    ID          string `json:"id"`
    Name        string `json:"name"`
    Description string `json:"description"`
    ParentID    string `json:"parent_id,omitempty"`
}