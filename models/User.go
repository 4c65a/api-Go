package models


type Task struct{

	Title string  `gorm:"not null;unique_index" json:"title"`
	Description string  `json:"description"`
	Done bool  `gorm:"default:false" json:"done"`
	UserId uint  `json:"user_id"`
	
}