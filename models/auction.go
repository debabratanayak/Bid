package models

import (
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"

	"github.com/gin-gonic/gin"
)

//Auction struct for holding auction details
type Auction struct {
	gorm.Model
	ProductID int    `json:"product_id"`
	Status    string `json:"status"`
	StartTime string `json:"start_time"`
	StopTime  string `json:"stop_time"`
}

var db *gorm.DB

//CreateAuction handler to insert auction details in database
func CreateAuction(c *gin.Context) {
	db := DbConn()
	var auct Auction
	err := c.BindJSON(&auct)

	if err != nil {
		fmt.Println(err)
		c.String(http.StatusBadRequest, "Error in request body format")
		return
	}

	err = db.Create(&auct).Error
	if err != nil {
		fmt.Println(err)
		c.String(http.StatusInternalServerError, "Error while inserting into database")
		return
	}

	c.String(http.StatusOK, "Created auction successfully")

}

//UpdateAuction handler to update auction details in database
func UpdateAuction(c *gin.Context) {
	db := DbConn()
	var auct Auction
	ID := c.Params.ByName("id")

	if err := db.Where("id = ?", ID).First(&auct).Error; err != nil {
		fmt.Println(err)
		c.String(http.StatusInternalServerError, "Error while updating database")
		return
	}

	err := c.BindJSON(&auct)
	if err != nil {
		fmt.Println(err)
		c.String(http.StatusBadRequest, "Error in request body format")
		return
	}

	db.Save(&auct)
	c.String(http.StatusOK, "Updated auction successfully")

}

//DeleteAuction handler to delete auction details in database
func DeleteAuction(c *gin.Context) {
	db := DbConn()
	var auct Auction
	ID := c.Params.ByName("id")

	if err := db.Where("id = ?", ID).Delete(&auct).Error; err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, "Error while deleting database")
		return
	}

	c.String(http.StatusOK, "Deleted auction successfully")

}

//ViewAllAuctions handler to list all auction details
func ViewAllAuctions(c *gin.Context) {
	db := DbConn()
	var auct []Auction

	if err := db.Order("created_at desc").Find(&auct).Error; err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, "Error while retrieving all auctions from database")
		return
	}

	c.JSON(http.StatusOK, auct)

}

//ViewOngoingAuctions handler to list all ongoing auctions
func ViewOngoingAuctions(c *gin.Context) {
	db := DbConn()
	var auct []Auction

	if err := db.Where("status = ?", "ongoing").Find(&auct).Error; err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, "Error while retrieving ongoing auctions from database")
		return
	}

	c.JSON(http.StatusOK, auct)

}

//ViewCompletedAuctions handler to list all completed auctions
func ViewCompletedAuctions(c *gin.Context) {
	db := DbConn()
	var auct []Auction

	if err := db.Where("status = ?", "completed").Find(&auct).Error; err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, "Error while retrieving completed auctions from database")
		return
	}

	c.JSON(http.StatusOK, auct)

}

//ViewUpcomingAuctions handler to list all upcoming auctions
func ViewUpcomingAuctions(c *gin.Context) {
	db := DbConn()
	var auct []Auction

	if err := db.Where("status = ?", "upcoming").Find(&auct).Error; err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, "Error while retrieving upcoming auctions from database")
		return
	}

	c.JSON(http.StatusOK, auct)

}

//ViewAuctionByID handler to list specific auction details
func ViewAuctionByID(c *gin.Context) {
	db := DbConn()
	var auct Auction
	ID := c.Params.ByName("id")

	if err := db.Where("id = ?", ID).Find(&auct).Error; err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusOK, "record not found")
			return
		}
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, "Error while retrieving from database")
		return
	}

	c.JSON(http.StatusOK, auct)
}
