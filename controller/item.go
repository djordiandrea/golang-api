package controller

import (
	"fmt"
	"log"
	"net/http"
	"rest-api/config"
	"rest-api/model"

	"github.com/gin-gonic/gin"
)

func RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ini root handler",
	})
}

func ItemParam(c *gin.Context) {
	id := c.Param("id")

	c.JSON(http.StatusOK, gin.H{
		"message": "ini item param with id " + id,
	})
}

func ItemQuery(c *gin.Context) {
	id := c.Query("id")
	title := c.Query("title")

	c.JSON(http.StatusOK, gin.H{
		"message": "ini item query with id " + id + " and title = " + title,
	})

}

func ItemPost(c *gin.Context) {
	var items model.Item

	err := c.ShouldBindJSON(&items)
	if err != nil {
		fmt.Println("Data Error : ", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   200,
		"id":       items.ID,
		"itemName": items.ItemName,
	})

}

// func GetAllUser(c *gin.Context) {
// 	var user []model.User

// 	// "mus_id", "mus_email", "mus_fullname", "mus_username"
// 	config.DB.Table("mst_user").Select("*").Scan(&user)

// 	c.JSON(http.StatusOK, gin.H{
// 		"result": user,
// 	})

// }

// func GetAllVehicle2(c *gin.Context) {
// 	var vehicle []model.Vehicle

// 	// "mus_id", "mus_email", "mus_fullname", "mus_username"
// 	config.DB.Table("mst_list_vehicle").Select("*").Scan(&vehicle)

// 	c.JSON(http.StatusOK, gin.H{
// 		"result": vehicle,
// 	})

// }

func GetUserLogin(c *gin.Context) {
	var user model.User
	var result []model.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		fmt.Println("Data Error : ", err)
		return
	}

	// rows, err := config.DB2.Query(
	// 	"select mus_id, mus_fullname, mus_username, mus_email from mst_user where mus_email = '" + user.Mus_email +
	// 		"' and mus_password = '" + user.Mus_password + "' ")

	rows, err := config.DB2.Query("CALL getUserLogin('" + user.Mus_email + "', '" + user.Mus_password + "')")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(
			&user.Mus_id,
			&user.Mus_fullname,
			&user.Mus_username,
			&user.Mus_email,
		); err != nil {
			panic(err)
		}

		result = append(result, user)
	}

	if len(result) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"result": result,
			"errMsg": "Data tidak ditemukan",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"result": result,
			"errMsg": "",
		})
	}

}

func GetAllUser(c *gin.Context) {

	var users []model.User

	rows, err := config.DB2.Query("select mus_id, mus_fullname, mus_username, mus_email from mst_user")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer rows.Close()

	for rows.Next() {
		var user model.User
		if err := rows.Scan(
			&user.Mus_id,
			&user.Mus_fullname,
			&user.Mus_username,
			&user.Mus_email,
		); err != nil {
			panic(err)
		}

		users = append(users, user)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": users,
	})

}

func GetAllVehicle(c *gin.Context) {
	var vehicles []model.Vehicle
	rows, err := config.DB2.Query("CALL getAllKendaraan();")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer rows.Close()

	for rows.Next() {
		var v model.Vehicle
		if err := rows.Scan(
			&v.Mlv_ID,
			&v.Mlv_vehicleName,
			&v.Mlv_vehicleCode,
			&v.Mlv_vehicleBrand,
			&v.Mlv_vehicleModel,
			&v.Mlv_vehicleNumber,
			&v.Mlv_vehicleLongRun,
			&v.Mlv_BuyingDate,
			&v.Mlv_vehiclePicture,
		); err != nil {
			log.Fatal(err)
		}
		vehicles = append(vehicles, v)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": vehicles,
	})

}
