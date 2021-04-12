package models

type entry struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	Artist    string `json:"artist"`
	Title     string `json:"title"`
  VendorID  string `json:"vendorid"`
  Cover     string `json:"cover"`
  Category  string `json:"category"`
}
