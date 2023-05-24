package model

type Item struct {
	ID       int
	ItemName string
}

// key yg ada di struct, itu harus nama kolomnya, huruf depannya juga harus huruf besar
type User struct {
	Mus_id       int    //`json:"mus_id"`
	Mus_email    string //`json:"mus_email", binding :"required"`
	Mus_fullname string //`json:"mus_fullname"`
	Mus_username string //`json:"mus_username"`
	Mus_password string `json:"mus_password"`
}

type Vehicle struct {
	Mlv_ID             int    `json:"id"`
	Mlv_vehicleName    string `json:"vehicleName"`
	Mlv_vehicleCode    string `json:"vehicleCode"`
	Mlv_vehicleBrand   string `json:"vehicleBrand"`
	Mlv_vehicleModel   string `json:"vehicleModel"`
	Mlv_vehicleNumber  string `json:"vehicleNumber"`
	Mlv_vehicleLongRun string `json:"vehicleLongRun"`
	Mlv_BuyingDate     string `json:"buyingDate"`
	Mlv_vehiclePicture string `json:"vehiclePicture"`
}
