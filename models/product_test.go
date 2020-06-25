package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestCreateProductSuccess(t *testing.T) {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	// Setup your router, just like you did in your main function, and
	// register your routes
	r := gin.Default()
	r.POST("/Admin/products/create", CreateProduct)

	pro := Products{
		Name:        "lamborghini",
		Description: "7000bhp",
		Status:      true,
	}

	reqBodyBytes := new(bytes.Buffer)

	json.NewEncoder(reqBodyBytes).Encode(pro)

	re, err := json.Marshal(pro)
	if err != nil {
		fmt.Println("Error in forming req body")
	}

	body := bytes.NewReader(re)

	// Create the mock request you'd like to test. Make sure the second argument
	// here is the same as one of the routes you defined in the router setup
	// block!
	req, err := http.NewRequest(http.MethodPost, "/Admin/products/create", body)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	// Create a response recorder so you can inspect the response
	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)

	// Check to see if the response was what you expected
	if w.Code != http.StatusOK {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
	}

	if w.Body.String() != "Success" {
		t.Errorf("Response body didnt match. actual %s, expected: %s.", w.Body.String(), "Success")
	}
}

func TestCreateProductFail(t *testing.T) {

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/Admin/products/create", CreateProduct)

	req, err := http.NewRequest(http.MethodPost, "/Admin/products/create", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}
	// Create a response recorder so you can inspect the response
	w := httptest.NewRecorder()
	// Perform the request
	r.ServeHTTP(w, req)
	// Check to see if the response was what you expected
	if w.Code != http.StatusBadRequest {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusBadRequest, w.Code)
	}
	// if w.Body.String() != "Error in inserting in database" {
	// 	t.Errorf("Response body didnt match. actual %s, expected: %s.", w.Body.String(), "Error in inserting in database")
	// }
}

func TestUpdateProductSuccess(t *testing.T) {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	// Setup your router, just like you did in your main function, and
	// register your routes
	r := gin.Default()
	r.PUT("/Admin/products/update", UpdateProduct)

	pro := Products{
		Name:        "maritu",
		Description: "700bhp",
		Status:      true,
	}

	reqBodyBytes := new(bytes.Buffer)

	json.NewEncoder(reqBodyBytes).Encode(pro)

	re, err := json.Marshal(pro)
	if err != nil {
		fmt.Println("Error in forming req body")
	}

	body := bytes.NewReader(re)

	// Create the mock request you'd like to test. Make sure the second argument
	// here is the same as one of the routes you defined in the router setup
	// block!
	req, err := http.NewRequest(http.MethodPut, "/Admin/products/update", body)

	q := req.URL.Query()
	q.Add("id", "6")
	req.URL.RawQuery = q.Encode()

	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	// Create a response recorder so you can inspect the response
	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)

	// Check to see if the response was what you expected
	if w.Code != http.StatusOK {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
	}

	if w.Body.String() != "Success" {
		t.Errorf("Response body didnt match. actual %s, expected: %s.", w.Body.String(), "Success")
	}
}

func TestUpdateProductFail(t *testing.T) {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)
	// Setup your router, just like you did in your main function, and
	// register your routes
	r := gin.Default()
	r.PUT("/Admin/products/update", UpdateProduct)

	// Create the mock request you'd like to test. Make sure the second argument
	// here is the same as one of the routes you defined in the router setup
	// block!
	req, err := http.NewRequest(http.MethodPut, "/Admin/products/update", nil)

	q := req.URL.Query()
	q.Add("id", "4")
	req.URL.RawQuery = q.Encode()

	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	// Create a response recorder so you can inspect the response
	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)

	// Check to see if the response was what you expected
	if w.Code != http.StatusBadRequest {

		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusBadRequest, w.Code)
	}
}

func TestDeleteProductSuccess(t *testing.T) {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	// Setup your router, just like you did in your main function, and
	// register your routes
	r := gin.Default()
	r.DELETE("/Admin/products/delete", DeleteProduct)

	// Create the mock request you'd like to test. Make sure the second argument
	// here is the same as one of the routes you defined in the router setup
	// block!
	req, err := http.NewRequest(http.MethodDelete, "/Admin/products/delete", nil)

	q := req.URL.Query()
	q.Add("id", "1")
	req.URL.RawQuery = q.Encode()

	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	// Create a response recorder so you can inspect the response
	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)

	// Check to see if the response was what you expected

	if w.Code != http.StatusOK {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
	}

	if w.Body.String() != "Success" {
		t.Errorf("Response body didnt match. actual %s, expected: %s.", w.Body.String(), "Success")
	}
}

func TestDeleteProductFail(t *testing.T) {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	// Setup your router, just like you did in your main function, and
	// register your routes
	r := gin.Default()
	r.DELETE("/Admin/products/delete", DeleteProduct)

	// Create the mock request you'd like to test. Make sure the second argument
	// here is the same as one of the routes you defined in the router setup
	// block!
	req, err := http.NewRequest(http.MethodDelete, "/Admin/products/delete", nil)

	q := req.URL.Query()
	q.Add("id", "89")
	req.URL.RawQuery = q.Encode()

	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	// Create a response recorder so you can inspect the response
	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)

	// Check to see if the response was what you expected

	if w.Code == http.StatusBadRequest {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusBadRequest, w.Code)
	}
}
func TestViewAllProducts(t *testing.T) {

	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	// Setup your router, just like you did in your main function, and
	// register your routes
	r := gin.Default()
	r.GET("/Admin/products/list", ViewAllProducts)

	// Create the mock request you'd like to test. Make sure the second argument
	// here is the same as one of the routes you defined in the router setup
	// block!
	req, err := http.NewRequest(http.MethodGet, "/Admin/products/list", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	// Create a response recorder so you can inspect the response
	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)

	// Check to see if the response was what you expected
	if w.Code != http.StatusOK {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
	}
}

func TestViewProductByIdSuccess(t *testing.T) {

	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	// Setup your router, just like you did in your main function, and
	// register your routes
	r := gin.Default()
	r.GET("/Admin/products/list", ViewProductById)

	// Create the mock request you'd like to test. Make sure the second argument
	// here is the same as one of the routes you defined in the router setup
	// block!
	req, err := http.NewRequest(http.MethodGet, "/Admin/products/list", nil)

	q := req.URL.Query()
	q.Add("id", "3")
	req.URL.RawQuery = q.Encode()

	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	// Create a response recorder so you can inspect the response
	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)

	// Check to see if the response was what you expected

	if w.Code != http.StatusOK {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
	}

	if w.Body.String() != "Success" {
		t.Errorf("Response body didnt match. actual %s, expected: %s.", w.Body.String(), "Success")
	}
}

func TestViewProductByIdFail(t *testing.T) {

	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	// Setup your router, just like you did in your main function, and
	// register your routes
	r := gin.Default()
	r.GET("/Admin/products/list", ViewProductById)

	// Create the mock request you'd like to test. Make sure the second argument
	// here is the same as one of the routes you defined in the router setup
	// block!
	req, err := http.NewRequest(http.MethodGet, "/Admin/products/list", nil)

	q := req.URL.Query()
	q.Add("id", "1")
	req.URL.RawQuery = q.Encode()

	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	// Create a response recorder so you can inspect the response
	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)

	// Check to see if the response was what you expected
	if w.Code != http.StatusServiceUnavailable {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusServiceUnavailable, w.Code)
	}
}
