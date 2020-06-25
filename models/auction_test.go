package models

/*
func TestDeleteAuctionPositive(t *testing.T) {
	// router := gin.Default()
	router := IntializeRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/Bidding/auction/5", nil)
	router.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("Response code didnt match. actual %d, expected: %d.", w.Code, 200)
	}

	if w.Body.String() != "Deleted auction successfully" {
		t.Errorf("Response body didnt match. actual %s, expected: %s.", w.Body.String(), "Deleted auction successfully")
	}
}

func TestDeleteAuctionNegative(t *testing.T) {
	// router := gin.Default()
	router := IntializeRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/Bidding/auction", nil)
	router.ServeHTTP(w, req)

	if w.Code != 404 {
		t.Errorf("Response code didnt match. actual %d, expected: %d.", w.Code, 404)
	}

	if w.Body.String() != "404 page not found" {
		t.Errorf("Response body didnt match. actual %s, expected: %s.", w.Body.String(), "404 page not found")
	}

}

func TestCreateAuctionPositive(t *testing.T) {
	// router := gin.Default()
	router := IntializeRoutes()
	w := httptest.NewRecorder()

	endtime, _ := time.Parse(time.RFC3339, "2019-12-12T16:40:36.48527+05:30")
	starttime, _ := time.Parse(time.RFC3339, "2019-12-12T16:40:36.48527+05:30")

	r := Auction{
		ProductID:   1,
		Status:      "ongoing",

		StartTime:   starttime,
		StopTime:    endtime,
	}

	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(r)

	re, err := json.Marshal(r)
	if err != nil {
		fmt.Println("Error in forming req body")
	}

	body := bytes.NewReader(re)

	req, _ := http.NewRequest("POST", "/Bidding/auction", body)
	router.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("Response code didnt match. actual %d, expected: %d.", w.Code, 200)
	}

	if w.Body.String() != "Created auction successfully" {
		t.Errorf("Response body didnt match. actual %s, expected: %s.", w.Body.String(), "Created auction successfully")
	}
}

func TestCreateAuctionNegative(t *testing.T) {
	// router := gin.Default()
	router := IntializeRoutes()
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/Bidding/auction", nil)
	router.ServeHTTP(w, req)

	if w.Code != 400 {
		t.Errorf("Response code didnt match. actual %d, expected: %d.", w.Code, 400)
	}

	if w.Body.String() != "Error in request body format" {
		t.Errorf("Response body didnt match. actual %s, expected: %s.", w.Body.String(), "Error in request body format")
	}
}

func TestUpdateAuctionPositive(t *testing.T) {
	// router := gin.Default()
	router := IntializeRoutes()
	w := httptest.NewRecorder()

	endtime, _ := time.Parse(time.RFC3339, "2019-12-12T16:40:36.48527+05:30")
	starttime, _ := time.Parse(time.RFC3339, "2019-12-12T16:40:36.48527+05:30")

	r := Auction{
		ProductID:   1,
		Status:      "completed",
		StartTime:   starttime,
		StopTime:    endtime,
	}

	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(r)

	re, err := json.Marshal(r)
	if err != nil {
		fmt.Println("Error in forming req body")
	}

	body := bytes.NewReader(re)

	req, _ := http.NewRequest("PUT", "/Bidding/auction/2", body)
	router.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("Response code didnt match. actual %d, expected: %d.", w.Code, 200)
	}

	if w.Body.String() != "Updated auction successfully" {
		t.Errorf("Response body didnt match. actual %s, expected: %s.", w.Body.String(), "Updated auction successfully")
	}
}

func TestUpdateNegative(t *testing.T) {
	// router := gin.Default()
	router := IntializeRoutes()
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("PUT", "/Bidding/auction/2", nil)
	router.ServeHTTP(w, req)

	if w.Code != 400 {
		t.Errorf("Response code didnt match. actual %d, expected: %d.", w.Code, 400)
	}

	if w.Body.String() != "Error in request body format" {
		t.Errorf("Response body didnt match. actual %s, expected: %s.", w.Body.String(), "Error in request body format")
	}
}

func TestViewAuctionByIdPositive(t *testing.T) {
	// router := gin.Default()
	router := IntializeRoutes()
	w := httptest.NewRecorder()

	endtime, _ := time.Parse(time.RFC3339, "2020-12-12T16:40:36.48527+05:30")
	starttime, _ := time.Parse(time.RFC3339, "2020-12-12T16:40:36.48527+05:30")

	expected := Auction{
		ProductID:   1,
		Status:      "upcoming",
		StartTime:   starttime,
		StopTime:    endtime,
	}

	req, _ := http.NewRequest("GET", "/Bidding/auction/list/details/2", nil)
	router.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("Response code didnt match. actual %d, expected: %d.", w.Code, 200)
	}

	// str := w.Body.String()
	actual := Auction{}
	json.Unmarshal([]byte(w.Body.String()), &actual)

	if actual.ProductID != expected.ProductID {
		t.Errorf("Product ID didnot match. actual %d, expected: %d.", actual.ProductID, expected.ProductID)
	}

	if actual.Status != expected.Status {
		t.Errorf("Status didnot match. actual %s, expected: %s.", actual.Status, expected.Status)
	}

	if actual.StartTime != expected.StartTime {
		t.Errorf("Start Time didnot match. actual %s, expected: %s.", actual.StartTime, expected.StartTime)
	}

	if actual.StopTime != expected.StopTime {
		t.Errorf("Stop Time didnot match. actual %s, expected: %s.", actual.StopTime, expected.StopTime)
	}
}

func TestViewAuctionByIdNegative(t *testing.T) {
	// router := gin.Default()
	router := IntializeRoutes()
	w := httptest.NewRecorder()

	endtime, _ := time.Parse(time.RFC3339, "2020-12-12T16:44:36.48527+05:30")
	starttime, _ := time.Parse(time.RFC3339, "2020-12-12T16:20:36.48527+05:30")

	expected := Auction{
		ProductID:   2,
		Status:      "completed",
		StartTime:   starttime,
		StopTime:    endtime,
	}

	req, _ := http.NewRequest("GET", "/Bidding/auction/list/details/2", nil)
	router.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("Response code didnt match. actual %d, expected: %d.", w.Code, 200)
	}

	// str := w.Body.String()
	actual := Auction{}
	json.Unmarshal([]byte(w.Body.String()), &actual)

	if actual.ProductID == expected.ProductID {
		t.Errorf("Product ID didnot match. actual %d, expected: %d.", actual.ProductID, expected.ProductID)
	}

	if actual.Status == expected.Status {
		t.Errorf("Status didnot match. actual %s, expected: %s.", actual.Status, expected.Status)
	}

	if actual.StartTime == expected.StartTime {
		t.Errorf("Start Time didnot match. actual %s, expected: %s.", actual.StartTime, expected.StartTime)
	}

	if actual.StopTime == expected.StopTime {
		t.Errorf("Stop Time didnot match. actual %s, expected: %s.", actual.StopTime, expected.StopTime)
	}
}

func TestViewAllAuctionsPositive(t *testing.T) {
	// router := gin.Default()
	router := IntializeRoutes()
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/Bidding/auction/list", nil)
	router.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("Response code didnt match. actual %d, expected: %d.", w.Code, 200)
	}
}

func TestViewAllAuctionsNegative(t *testing.T) {
	// router := gin.Default()
	router := IntializeRoutes()
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/Bidding/auction/lis", nil)
	router.ServeHTTP(w, req)

	if w.Code != 404 {
		t.Errorf("Response code didnt match. actual %d, expected: %d.", w.Code, 404)
	}
}

func TestViewUpcomingAuctionsPositive(t *testing.T) {
	// router := gin.Default()
	router := IntializeRoutes()
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/Bidding/auction/list/upcoming", nil)
	router.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("Response code didnt match. actual %d, expected: %d.", w.Code, 200)
	}
}

func TestViewUpcomingAuctionsNegative(t *testing.T) {
	// router := gin.Default()
	router := IntializeRoutes()
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/Bidding/auction/list/upcomi", nil)
	router.ServeHTTP(w, req)

	if w.Code != 404 {
		t.Errorf("Response code didnt match. actual %d, expected: %d.", w.Code, 404)
	}
}

func TestViewOngoingAuctionsPositive(t *testing.T) {
	// router := gin.Default()
	router := IntializeRoutes()
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/Bidding/auction/list/ongoing", nil)
	router.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("Response code didnt match. actual %d, expected: %d.", w.Code, 200)
	}
}

func TestViewOngoingAuctionsNegative(t *testing.T) {
	// router := gin.Default()
	router := IntializeRoutes()
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/Bidding/auction/list/ongoi", nil)
	router.ServeHTTP(w, req)

	if w.Code != 404 {
		t.Errorf("Response code didnt match. actual %d, expected: %d.", w.Code, 404)
	}
}

func TestViewCompletedAuctionsPositive(t *testing.T) {
	// router := gin.Default()
	router := IntializeRoutes()
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/Bidding/auction/list/completed", nil)
	router.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("Response code didnt match. actual %d, expected: %d.", w.Code, 200)
	}
}

func TestViewCompletedAuctionsNegative(t *testing.T) {
	// router := gin.Default()
	router := IntializeRoutes()
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/Bidding/auction/list/complet", nil)
	router.ServeHTTP(w, req)

	if w.Code != 404 {
		t.Errorf("Response code didnt match. actual %d, expected: %d.", w.Code, 404)
	}
}
*/
