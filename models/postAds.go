package models

type Advert struct{
	 
	Id uint `json:"id"`
	Title string `json:"title"`
	Desc  string `json:"desc"`
	UserID uint  `json:"userid"`
	User User `json:"user"`
	 
	 	 
}