package routes

import (
	//"path/filepath"
	//"test1"

	"path/filepath"
	"test1/controller"
	"test1/models"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

func IntializeRoutes(router *gin.Engine) {

	router.StaticFile("assets/css/bootstrap.min.css", "assets/css/bootstrap.min.css")
	router.StaticFile("assets/js/jquery-3.3.1.slim.min.js", "assets/js/jquery-3.3.1.slim.min.js")
	router.StaticFile("assets/js/popper.min.js", "assets/js/popper.min.js")
	router.StaticFile("assets/js/bootstrap.min.js", "assets/js/bootstrap.min.js")
	router.StaticFile("assets/css/signin.css", "assets/css/signin.css")

	router.HTMLRender = LoadTemplates("./templates")
	router.GET("/admin/login", func(c *gin.Context) {
		c.HTML(200, "admin-login.html", gin.H{
			"title": "Admin Login",
		})
	})
	router.GET("/admin/forgot", func(c *gin.Context) {
		c.HTML(200, "admin-forgotpassword.html", gin.H{
			"title": "Admin Forgot Password",
		})
	})
	router.GET("/user/login", func(c *gin.Context) {
		c.HTML(200, "user-login.html", gin.H{
			"title": "User Login",
		})
	})
	router.GET("/user/register", func(c *gin.Context) {
		c.HTML(200, "user-register.html", gin.H{
			"title": "User Register",
		})
	})
	router.GET("/admin/dashboard/user-listing", func(c *gin.Context) {
		c.HTML(200, "user-listing.html", gin.H{
			"title": "User Listing",
		})
	})

	router.GET("/admin/manageproduct", controller.ListProductView)

	router.GET("/admin/addproduct", func(c *gin.Context) {
		c.HTML(200, "admin-addproduct.html", gin.H{
			"title": "Product Admin Dashboard",
		})
	})

	router.GET("/admin/products/save", func(c *gin.Context) {
		c.HTML(200, "manage-product.html", gin.H{
			"title": "Add Product",
		})
	})

	router.GET("/admin/updateproduct/:id", controller.EditProductController)
	router.GET("/admin/deleteproduct/:id", controller.DeleteProductController)
	// router.GET("/admin/products/ccc", func(c *gin.Context) {
	// 	c.HTML(200, "manage-product.html", gin.H{
	// 		"title": "Add Product",
	// 	})
	// })

	// APIs for CRUD operations
	api := router.Group("/admin")
	api.POST("/products/create", models.CreateProduct)
	api.POST("/products/update", controller.UpdateProduct)
	api.PUT("/products/update/:id", models.UpdateProduct)
	api.DELETE("/products/delete/:id", models.DeleteProduct)
	api.GET("/products/list", models.ViewAllProducts)
	api.GET("products/list/:id", models.ViewProductById)

	api.POST("/products/save", controller.AddProduct)
	api.GET("/products/ccc", controller.ListProductView)

	bid := router.Group("/Bidding")
	bid.POST("/auction", models.CreateAuction)
	bid.PUT("/auction/update/:id", models.UpdateAuction)
	bid.DELETE("/auction/delete/:id", models.DeleteAuction)
	bid.GET("/auction/list", models.ViewAllAuctions)
	bid.GET("/auction/list/details/:id", models.ViewAuctionByID)
	bid.GET("/auction/list/ongoing", models.ViewOngoingAuctions)
	bid.GET("/auction/list/completed", models.ViewCompletedAuctions)
	bid.GET("/auction/list/upcoming", models.ViewUpcomingAuctions)

	router.GET("/admin/auction/create", controller.AddAuctionPage)
	router.GET("/admin/auction/list", controller.ViewAuctionPage)
	router.POST("/admin/auctions/save", controller.AddAuction)
	router.GET("/admin/auction/update/:id", controller.EditAuctionController)
	router.POST("/admin/auctions/update", controller.UpdateAuction)
	router.GET("/admin/auction/delete/:id", controller.DeleteAuction)

	// bid.POST("/bid", CreateBid)
	// bid.GET("/bid/list", ViewAllBid)
	// bid.GET("bid/list/:id", ViewBidByID)
}

func LoadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	layouts, err := filepath.Glob(templatesDir + "/layouts/*.html")
	if err != nil {
		panic(err.Error())
	}

	includes, err := filepath.Glob(templatesDir + "/includes/*.html")
	if err != nil {
		panic(err.Error())
	}

	// Generate our templates map from our layouts/ and includes/ directories
	for _, include := range includes {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, include)
		r.AddFromFiles(filepath.Base(include), files...)
	}
	return r
}
