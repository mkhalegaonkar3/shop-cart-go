package main

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"

	list "github.com/KaustubhLonkar/shop-cart-go/controller/list"
	login "github.com/KaustubhLonkar/shop-cart-go/controller/login"
	"github.com/KaustubhLonkar/shop-cart-go/controller/products"
	"github.com/KaustubhLonkar/shop-cart-go/model/DB"

	//order "github.com/KaustubhLonkar/shop-cart-go/order"
	//products "github.com/KaustubhLonkar/shop-cart-go/products"

	"github.com/gin-gonic/gin"
)

// used to match which service is being called
const (
	LOGIN          = "/login"
	VERIFYUSERNAME = "/verifyUsername"
	SIGNUP         = "/signup"
	CHANGEPASS     = "/changepass"
	RESETPWD       = "/resetpwd"
	CREATELIST     = "/createList"
	ADDITEMS       = "/addItems"
	SHARELIST      = "/shareList"
	DELETEITEM     = "/deleteItem"
	DELETELIST     = "/deleteList"
	ADDPRODUCT     = "/addProduct"
	// GETPRODUCTS = "/getProducts"
	// PLACEORDER = "/placeOrder"
)

func main() {

	rand.Seed(time.Now().UTC().UnixNano())
	router := initRouter()
	router.Run(":8888")
}

func initRouter() *gin.Engine {

	r := gin.New()
	//p.Use(r)
	r.Use(gin.Recovery(), plainLoggerWithWriter(gin.DefaultWriter))
	r.GET("/status", statusCheck)
	r.POST("/signup", requestRouter)
	r.POST("/login", requestRouter)
	r.POST("/verifyUsername", requestRouter)
	r.POST("/changepass", requestRouter)
	//r.GET("/resetpwd", login.ResetPasswordPage())
	r.GET("/sessions", requestRouter)
	r.POST("/createList", requestRouter)
	r.POST("/addItems", requestRouter)
	r.POST("/shareList", requestRouter)
	r.POST("/deleteItem", requestRouter)
	r.POST("/deleteList", requestRouter)
	// r.GET("/getProducts", requestRouter)
	// r.POST("/placeOrder", requestRouter)
	return r
}

// PlainLoggerWithWriter mimics the Gin LoggerWithWriter without the colors
func plainLoggerWithWriter(out io.Writer) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()

		// Process request
		c.Next()

		// Stop timer
		end := time.Now()
		latency := end.Sub(start)

		if c.Request.URL.Path != "/status" {
			fmt.Fprintf(out, "%s [%s] %s [%v] \"%s %s %s\" %d %d %v %s %s %s \"%s\"\n",
				c.ClientIP(),
				c.Request.UserAgent(),
				c.Request.Header.Get(gin.AuthUserKey),
				end.Format("02/Jan/2006:15:04:05 -0700"),
				c.Request.Method,
				c.Request.URL.Path,
				c.Request.Proto,
				c.Writer.Status(),
				c.Writer.Size(),
				fmt.Sprintf("%.4f", latency.Seconds()),
				c.Request.Header.Get("RequestType"),
				c.Request.Header.Get("ResponseSource"),
				c.Request.Form.Encode(),
				c.Request.Header.Get("ResponseBody"),
			)
		}
	}
}

// statusCheck returns a 200/OK when called if we can contact the be env
func statusCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}

func exception(c *gin.Context) {
	c.JSON(500, gin.H{"success": false, "error": "Unable to process order"})
}

func requestRouter(c *gin.Context) {
	db, err := DB.Start()

	if err != nil {
		fmt.Printf(err.Error())
	}

	defer db.Close()
	path := c.Request.URL.Path
	fmt.Println("The obtained path is:- ", path)
	switch path {
	case ADDPRODUCT:
		products.AddProduct(db, c)
	// case GETPRODUCTS:
	// products.GetProducts(c)
	// case PLACEORDER:
	// order.PlaceOrder(c)
	case LOGIN:
		login.LoginPost(db, c)
	case VERIFYUSERNAME:
		login.VerifyUsername(db, c)
	case SIGNUP:
		login.RegistrationPost(db, c)
	case CHANGEPASS:
		login.PasswordReset(db, c)
		// case RESETPWD:
		// 	login.ResetPasswordPage(db)
	case CREATELIST:
		list.CreateList(db, c)
	case ADDITEMS:
		list.AddItemsList(db, c)
	// case SHARELIST:

	case DELETEITEM:
		list.DeleteItemList(db, c)
	case DELETELIST:
		list.DeleteList(db, c)

	}

}
