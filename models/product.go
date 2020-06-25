package models

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
type Products struct {
	gorm.Model
	//	Id          uint   `gorm:"AUTO_INCREMENT" json:"product_id"`
	Name        string `gorm:"type:varchar(100)" json:"product_name"`
	Description string `gorm:"type:varchar(100)" json:"product_desc"`
	Status      bool   `json:"status"`
	Image       string `json:"image"`
}

// Function that creates product by entering details
func CreateProduct(c *gin.Context) {
	// if os.Getenv("app_status") != "true" {
	// 	c.JSON(http.StatusInternalServerError, "Product service is temporarily down")
	// 	return
	// }
	db := DbConn()
	var prod Products
	err := c.BindJSON(&prod)
	fmt.Println("Error binding json: ", err)
	fmt.Println("prod: ", prod)
	if err != nil {
		fmt.Println("Error in req format")
		c.JSON(http.StatusBadRequest, "Error in req format")
	}
	err = db.Create(&prod).Error
	if err != nil {
		fmt.Println("Error in inserting in database")
		c.String(http.StatusServiceUnavailable, "Error in inserting in database")
	}

	c.String(http.StatusOK, "Success")

}

//Function that updates product by id.
func UpdateProduct(c *gin.Context) {

	id := c.Param("id")
	//id := c.Query("id")
	fmt.Println("id is: ", id)
	db := DbConn()
	var prod Products
	dbc := db.Debug().Where("id = ?", id).First(&prod)
	if dbc.Error != nil {
		fmt.Println("Error while updating: ", dbc.Error)

		c.String(http.StatusBadRequest, "Error while updating")
	}
	if e := c.BindJSON(&prod); e != nil {
		fmt.Println("Error in req format", e)
		c.JSON(http.StatusBadRequest, "Error in req format")
	}
	db.Save(&prod)
	fmt.Println("updated successfully")
	c.String(http.StatusOK, "Success")
}

//Function that deletes product by id.
func DeleteProduct(c *gin.Context) {
	ID := c.Param("id")

	//ID := c.Query("id")
	fmt.Println("id is: ", ID)
	log.Println("id: ", ID)
	db := DbConn()
	var prod Products
	dbc := db.Where("id = ?", ID).Delete(&prod).Error
	if dbc != nil {
		fmt.Println("Error while deleting: ", dbc)
		c.String(http.StatusBadRequest, "Error while deleting")
	}
	c.String(http.StatusOK, "Success")
}

//Function to view list of all products
func ViewAllProducts(c *gin.Context) {
	var prod []Products
	db := DbConn()
	err := db.Find(&prod).Error

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusServiceUnavailable, err)
	}

	c.JSON(http.StatusOK, prod)
}

//Function to ViewProductById
func ViewProductById(c *gin.Context) {

	id := c.Param("id")

	fmt.Println("id is: ", id)
	var prod Products
	db := DbConn()
	defer db.Close()
	err := db.Where("id = ?", id).First(&prod).Error
	if err != nil {
		fmt.Println(err)
		c.String(http.StatusServiceUnavailable, "Error While Searching ID:")
	}
	c.JSON(http.StatusOK, prod)
}
