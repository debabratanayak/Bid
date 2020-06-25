package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"test1/models"
	"text/template"

	"github.com/gin-gonic/gin"
)

//Product add controller
func AddProduct(c *gin.Context) {
	fmt.Println("Inside Add product")
	name := c.PostForm("prod_name")
	desc := c.PostForm("prod_desc")
	file, er := c.FormFile("file")
	if er != nil {
		fmt.Println("Upload err:", er)
	}
	chk := false
	check := c.PostForm("check")
	fmt.Println("Checkbox :", check)
	if check == "check" {
		chk = true
	}
	fmt.Println("Value of Chk", chk)
	apiUrl := "http://localhost:8080/admin/products/create"
	u, _ := url.ParseRequestURI(apiUrl)
	urlStr := u.String()
	fmt.Println("Name", name)
	fmt.Println("Desc", desc)
	// filename := filepath.Base(file.Filename)
	// fmt.Println(filename)
	path := "C:/Users/Debabratan/Desktop/images"

	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, os.ModePerm)
	}

	image := path + "/" + file.Filename
	if err := c.SaveUploadedFile(file, image); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
		return
	}

	if name == "" && desc == "" {
		c.HTML(http.StatusOK, "admin-addproduct.html", gin.H{
			"msg": "Name ,Desc & File Required..",
		})
		return
	} else {
		jsonData := map[string]interface{}{
			"product_name": name,
			"product_desc": desc,
			"image":        image,
			"status":       chk,
		}
		//jsonData := tmpStr{name, desc}
		jsonValue, err := json.Marshal(jsonData)
		if err != nil {
			log.Println("Error marshalling :", err)
		} //convert string to json by marshalling
		r, err := http.NewRequest("POST", urlStr, bytes.NewBuffer(jsonValue)) // URL-encoded payload
		if err != nil {
			fmt.Println("Error while doing post request: ", err)
		}
		r.Header.Add("Content-Type", "application/json")

		client := &http.Client{}
		resp, _ := client.Do(r)
		fmt.Println("Status:", resp.Status)

		if resp.StatusCode == http.StatusOK {
			c.Redirect(http.StatusMovedPermanently, "/admin/manageproduct")
			// c.Redirect(http.StatusOK, "manage-product.html")
			// c.HTML(http.StatusOK, "manage-product.html", gin.H{
			// 	"msg": "Success",
			// })
		} else {
			//fmt.Println("login else")
			c.HTML(http.StatusBadRequest, "admin-addproduct.html", gin.H{
				"msg": "Some error occured while saving",
			})
		}
	}
}

//Update Product Controller
func UpdateProduct(c *gin.Context) {
	fmt.Println("Inside Update product")
	name := c.PostForm("prod_name")
	desc := c.PostForm("prod_desc")
	file, er := c.FormFile("file")
	if er != nil {
		fmt.Println("Upload err:", er)
	}
	id := c.PostForm("id")
	fmt.Println("id: ", id)
	apiUrl := "http://10.10.34.163:8080/admin/products/update/" + id
	u, _ := url.ParseRequestURI(apiUrl)
	urlStr := u.String()
	fmt.Println("urlStr:", urlStr)
	fmt.Println("Name", name)
	fmt.Println("Desc", desc)
	// filename := filepath.Base(file.Filename)
	// fmt.Println(filename)
	path := "C:/Users/Debabratan/Desktop/images"

	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, os.ModePerm)
	}

	image := path + "/" + file.Filename
	if err := c.SaveUploadedFile(file, image); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
		return
	}

	if name == "" && desc == "" {
		c.HTML(http.StatusOK, "admin-addproduct.html", gin.H{
			"msg": "Name ,Desc & File Required..",
		})
		return
	} else {
		jsonData := map[string]interface{}{
			"product_name": name,
			"product_desc": desc,
			"image":        image,
		}
		//jsonData := tmpStr{name, desc}
		jsonValue, err := json.Marshal(jsonData)
		if err != nil {
			log.Println("Error marshalling :", err)
		} //convert string to json by marshalling
		r, err := http.NewRequest("PUT", urlStr, bytes.NewBuffer(jsonValue)) // URL-encoded payload
		if err != nil {
			fmt.Println("Error while doing post request: ", err)
		}
		r.Header.Add("Content-Type", "application/json")

		client := &http.Client{}
		resp, _ := client.Do(r)
		fmt.Println("Status:", resp.Status)

		if resp.StatusCode == http.StatusOK {
			c.Redirect(http.StatusMovedPermanently, "/admin/manageproduct")
			// c.Redirect(http.StatusOK, "manage-product.html")
			// c.HTML(http.StatusOK, "manage-product.html", gin.H{
			// 	"msg": "Success",
			// })
		} else {
			//fmt.Println("login else")
			c.HTML(http.StatusBadRequest, "admin-addproduct.html", gin.H{
				"msg": "Some error occured while saving",
			})
		}
	}
}

//List Product Controller
func ListProductView(c *gin.Context) {
	apiUrl := "http://localhost:8080/admin/products/list"
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

	tmpl, err := template.ParseFiles("./templates/includes/manage-product.html")
	if err != nil {
		fmt.Println("Error in Template: ", err)
	}
	var obj []models.Products
	if err := json.Unmarshal(body, &obj); err != nil {
		panic(err)
	}

	tmpl.Execute(c.Writer, obj)

	if resp.StatusCode == http.StatusOK {
		// c.Redirect(http.StatusMovedPermanently, "/admin/manageproduct")
		// c.Redirect(http.StatusOK, "manage-product.html")
		c.HTML(http.StatusOK, "manage-product.html", gin.H{
			"Product": obj,
		})
	} else {
		//fmt.Println("login else")
		c.HTML(http.StatusBadRequest, "admin-addproduct.html", gin.H{
			"msg": "Some error occured while saving",
		})
	}
}

//Edit Product controller
func EditProductController(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("id :=", id)

	var prod models.Products
	apiUrl := "http://10.10.34.163:8080/admin/products/list/" + id

	u, _ := url.ParseRequestURI(apiUrl)
	urlStr := u.String()
	r, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		fmt.Println("Error while doing get request: ", err)
	}

	client := &http.Client{}
	resp, err := client.Do(r)
	if err != nil {
		fmt.Println("Error while reading body: ", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error while reading body: ", err)
	}
	fmt.Println("body:", string(body))

	if err := json.Unmarshal(body, &prod); err != nil {
		fmt.Println("Error while unmarshalling body: ", err)
	}
	fmt.Println("Product details:", prod)
	// db := models.DbConn()
	// defer db.Close()

	// err := db.Where("id = ?", id).First(&prod).Error
	// if err != nil {
	// 	fmt.Println(err)
	// 	//c.String(http.StatusServiceUnavailable, "Error While Searching ID:")
	// }

	c.HTML(200, "admin-editproduct.html", gin.H{
		"title":       "Edit Product",
		"id":          prod.ID,
		"name":        prod.Name,
		"description": prod.Description,
	})
}

//Delete product controller
func DeleteProductController(c *gin.Context) {
	fmt.Println("Inside delete product controller")
	id := c.Param("id")
	fmt.Println("ID :", id)
	apiUrl := "http://localhost:8080/admin/products/delete/" + id
	u, _ := url.ParseRequestURI(apiUrl)
	urlStr := u.String()
	r, err := http.NewRequest("DELETE", urlStr, nil) // URL-encoded payload
	if err != nil {
		fmt.Println("Error while doing post request: ", err)
	}
	client := &http.Client{}
	resp, _ := client.Do(r)
	fmt.Println("Status:", resp.Status)

	if resp.StatusCode == http.StatusOK {
		c.Redirect(http.StatusMovedPermanently, "/admin/manageproduct")
		// c.Redirect(http.StatusOK, "manage-product.html")
		// c.HTML(http.StatusOK, "manage-product.html", gin.H{
		// 	"msg": "Success",
		// })
	} else {
		//fmt.Println("login else")
		c.HTML(http.StatusBadRequest, "admin-addproduct.html", gin.H{
			"msg": "Some error occured while saving",
		})
	}
}
