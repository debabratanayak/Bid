package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"test1/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Products struct {
	gorm.Model
	//	Id          uint   `gorm:"AUTO_INCREMENT" json:"product_id"`
	Name        string `gorm:"type:varchar(100)" json:"product_name"`
	Description string `gorm:"type:varchar(100)" json:"product_desc"`
	Status      bool   `json:"status"`
	Image       string `json:"image"`
}

func AddAuctionPage(c *gin.Context) {
	api := "http://10.10.34.163:8080/admin/products/list"

	r, err := http.NewRequest("GET", api, nil)
	if err != nil {
		fmt.Println(err)
		c.HTML(http.StatusBadRequest, "admin-addauction.html", gin.H{
			"msg": "Error in getting products",
		})
		return
	}

	client := &http.Client{}
	resp, _ := client.Do(r)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error while reading body: ", err)
	}
	tmpl, err := template.ParseFiles("./templates/includes/admin-addauction.html")
	if err != nil {
		fmt.Println("Error in Template: ", err)
	}

	var obj []Products
	if err := json.Unmarshal(body, &obj); err != nil {
		fmt.Println(err)
	}

	tmpl.Execute(c.Writer, obj)

	if resp.StatusCode == http.StatusOK {

		// c.Redirect(http.StatusMovedPermanently, "/admin/manageproduct")
		// c.Redirect(http.StatusOK, "manage-product.html")
		c.HTML(http.StatusOK, "admin-addauction.html", gin.H{
			"Product": obj,
		})
	} else {
		//fmt.Println("login else")
		c.HTML(http.StatusBadRequest, "admin-addauction.html", gin.H{
			"msg": "Some error occured while saving",
		})
	}

}

func AddAuction(c *gin.Context) {
	// layout := "2006-12-12 15:37:31.847349+05:30"

	product := c.PostForm("product_id")
	startTime := c.PostForm("start_time")
	stopTime := c.PostForm("stop_time")
	status := "upcoming"
	fmt.Println(product)
	fmt.Println(startTime)
	fmt.Println(stopTime)

	api := "http://localhost:8080/Bidding/auction"
	u, _ := url.ParseRequestURI(api)
	urlStr := u.String()

	if startTime == "" || stopTime == "" || product == "" {
		fmt.Println("91")
		c.HTML(http.StatusOK, "admin-addauction.html", gin.H{
			"msg": "All fields are required",
		})
		return
	}

	productID, err := strconv.Atoi(product)
	if err != nil {
		fmt.Println(err)
		c.HTML(http.StatusBadRequest, "admin-addauction.html", gin.H{
			"msg": "Product Id",
		})
		fmt.Println("103")
		return
	}

	// start, err := time.Parse(layout, startTime)
	// if err != nil {
	// 	fmt.Println(err)
	// 	c.HTML(http.StatusBadRequest, "admin-addauction.html", gin.H{
	// 		"msg": "Bad Request Start Time",
	// 	})
	// 	return
	// }

	// stop, err := time.Parse(layout, stopTime)
	// if err != nil {
	// 	fmt.Println(err)
	// 	c.HTML(http.StatusBadRequest, "admin-addauction.html", gin.H{
	// 		"msg": "Bad Request Stop Time",
	// 	})
	// 	return
	// }

	auct := models.Auction{
		ProductID: productID,
		StartTime: startTime,
		StopTime:  stopTime,
		Status:    status,
	}

	jsonValue, err := json.Marshal(auct)
	if err != nil {
		fmt.Println(err)
		c.HTML(http.StatusBadRequest, "admin-addauction.html", gin.H{
			"msg": "Error in json marshalling",
		})
		return
	}

	r, err := http.NewRequest("POST", urlStr, bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Println(err)
		c.HTML(http.StatusBadRequest, "admin-addauction.html", gin.H{
			"msg": "Error in sending post request",
		})
		return
	}

	client := &http.Client{}
	resp, _ := client.Do(r)
	fmt.Println("Status:", resp.Status)

	if resp.StatusCode != http.StatusOK {
		c.HTML(http.StatusBadRequest, "admin-addproduct.html", gin.H{
			"msg": "Status code didnt match",
		})
		return
	}

	// c.HTML(http.StatusOK, "/admin/auction/list", gin.H{
	// 	"msg": "Success",
	// })

	c.Redirect(http.StatusMovedPermanently, "/admin/auction/list")

}
func ViewAuctionPage(c *gin.Context) {
	apiUrl := "http://localhost:8080/Bidding/auction/list"
	u, _ := url.ParseRequestURI(apiUrl)
	urlStr := u.String()
	r, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		fmt.Println("Error while doing get request: ", err)
	}

	client := &http.Client{}
	resp, _ := client.Do(r)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error while reading body: ", err)
	}

	tmpl, err := template.ParseFiles("./templates/includes/manage-auctions.html")
	if err != nil {
		fmt.Println("Error in Template: ", err)
	}
	var obj []models.Auction
	if err := json.Unmarshal(body, &obj); err != nil {
		panic(err)
	}

	tmpl.Execute(c.Writer, obj)

	if resp.StatusCode == http.StatusOK {
		// c.Redirect(http.StatusMovedPermanently, "/admin/manageproduct")
		// c.Redirect(http.StatusOK, "manage-product.html")
		c.HTML(http.StatusOK, "manage-auctions.html", gin.H{
			"Auctions": obj,
		})
	} else {
		//fmt.Println("login else")
		c.HTML(http.StatusBadRequest, "manage-auctions.html", gin.H{
			"msg": "Some error occured while saving",
		})
	}
}

func UpdateAuction(c *gin.Context) {

	product := c.PostForm("product_id")
	startTime := c.PostForm("start_time")
	stopTime := c.PostForm("stop_time")
	id := c.PostForm("id")

	fmt.Println("pro_id:", product)
	fmt.Println("starttime:", startTime)
	fmt.Println("stoptime:", stopTime)
	fmt.Println("id:", id)

	productID, err := strconv.Atoi(product)
	if err != nil {
		fmt.Println(err)
		c.HTML(http.StatusBadRequest, "admin-addauction.html", gin.H{
			"msg": "Product Id",
		})
		fmt.Println("103")
		return
	}

	fmt.Println("233")

	apiUrl := "http://localhost:8080/Bidding/auction/update/" + id
	fmt.Println(apiUrl)
	u, _ := url.ParseRequestURI(apiUrl)
	urlStr := u.String()

	if product == "" || startTime == "" || stopTime == "" {
		c.HTML(http.StatusOK, "admin-editauction.html", gin.H{
			"msg": "All fields are required",
		})
		return
	}

	jsonData := map[string]interface{}{
		"product_id": productID,
		"start_time": startTime,
		"stop_time":  stopTime,
		"status":     "ongoing",
	}
	//jsonData := tmpStr{name, desc}

	jsonValue, err := json.Marshal(jsonData)
	if err != nil {
		fmt.Println(err)
		log.Println("Error marshalling :", err)

	} //convert string to json by marshalling

	r, err := http.NewRequest("PUT", urlStr, bytes.NewBuffer(jsonValue)) // URL-encoded payload
	if err != nil {
		fmt.Println("Error while doing post request: ", err)
	}
	r.Header.Add("Content-Type", "application/json")
	fmt.Println(r)
	client := &http.Client{}
	fmt.Println(client)
	resp, err := client.Do(r)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	fmt.Println(resp.Status)
	if resp.StatusCode == http.StatusOK {
		c.Redirect(http.StatusMovedPermanently, "/admin/auction/list")
		// c.Redirect(http.StatusOK, "manage-product.html")
		// c.HTML(http.StatusOK, "manage-product.html", gin.H{
		// 	"msg": "Success",
		// })
	} else {
		//fmt.Println("login else")
		c.HTML(http.StatusBadRequest, "manage-auctions.html", gin.H{
			"msg": "Some error occured while saving",
		})
	}

}

func EditAuctionController(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("id :=", id)

	var auct models.Auction
	api := "http://10.10.34.163:8080/Bidding/auction/list/details/" + id

	r, err := http.NewRequest("GET", api, nil)
	if err != nil {
		fmt.Println(err)
		c.HTML(http.StatusBadRequest, "admin-addauction.html", gin.H{
			"msg": "Error in getting products",
		})
		return
	}

	client := &http.Client{}
	resp, err := client.Do(r)
	if err != nil {
		fmt.Println(err)
		c.HTML(http.StatusBadRequest, "admin-addauction.html", gin.H{
			"msg": "Error in getting products",
		})
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error while reading body: ", err)
	}

	if err := json.Unmarshal(body, &auct); err != nil {
		fmt.Println(err)
	}

	api = "http://10.10.34.163:8080/admin/products/list"

	r, err = http.NewRequest("GET", api, nil)
	if err != nil {
		fmt.Println(err)
		c.HTML(http.StatusBadRequest, "admin-addauction.html", gin.H{
			"msg": "Error in getting products",
		})
		return
	}

	client = &http.Client{}
	resp, err = client.Do(r)
	if err != nil {
		fmt.Println(err)
		c.HTML(http.StatusBadRequest, "admin-addauction.html", gin.H{
			"msg": "Error in getting products",
		})
		return
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error while reading body: ", err)
	}
	tmpl, err := template.ParseFiles("./templates/includes/admin-addauction.html")
	if err != nil {
		fmt.Println("Error in Template: ", err)
	}

	var obj []Products
	if err := json.Unmarshal(body, &obj); err != nil {
		fmt.Println(err)
	}

	tmpl.Execute(c.Writer, obj)

	pr := strconv.Itoa(auct.ProductID)

	api = "http://10.10.34.163:8080/admin/products/list/" + pr
	fmt.Println(api)
	r, err = http.NewRequest("GET", api, nil)
	if err != nil {
		fmt.Println(err)
		c.HTML(http.StatusBadRequest, "admin-addauction.html", gin.H{
			"msg": "Error in getting products",
		})
		return
	}

	client = &http.Client{}
	resp, err = client.Do(r)
	if err != nil {
		fmt.Println(err)
		c.HTML(http.StatusBadRequest, "admin-addauction.html", gin.H{
			"msg": "Error in getting products",
		})
		return
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error while reading body: ", err)
	}

	var productdetails Products
	if err := json.Unmarshal(body, &productdetails); err != nil {
		fmt.Println(err)
	}

	fmt.Println("283", auct.ProductID)
	c.HTML(200, "admin-editauction.html", gin.H{
		"title":        "Edit Product",
		"id":           id,
		"product_id":   auct.ProductID,
		"product_name": productdetails.Name,
		"all_product":  obj,
		"start_time":   auct.StartTime,
		"stop_time":    auct.StopTime,
		"status":       auct.Status,
	})
}

func DeleteAuction(c *gin.Context) {

	id := c.Param("id")
	fmt.Println("id:", id)

	apiUrl := "http://localhost:8080/Bidding/auction/delete/" + id
	fmt.Println(apiUrl)
	u, _ := url.ParseRequestURI(apiUrl)
	urlStr := u.String()

	r, err := http.NewRequest("DELETE", urlStr, nil) // URL-encoded payload
	if err != nil {
		fmt.Println("Error while doing post request: ", err)
	}

	client := &http.Client{}

	resp, err := client.Do(r)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	fmt.Println(resp.Status)
	if resp.StatusCode == http.StatusOK {
		c.Redirect(http.StatusMovedPermanently, "/admin/auction/list")
		// c.Redirect(http.StatusOK, "manage-product.html")
		// c.HTML(http.StatusOK, "manage-product.html", gin.H{
		// 	"msg": "Success",
		// })
	} else {
		//fmt.Println("login else")
		c.HTML(http.StatusBadRequest, "manage-auctions.html", gin.H{
			"msg": "Some error occured while saving",
		})
	}

}
