package controllers
import (
   "context"
   "fmt"
   "strconv"
   "time"
   "github.com/gin-gonic/gin"
   "github.com/team13/app/ent"
   "github.com/team13/app/ent/orderonline"
   "github.com/team13/app/ent/paymentchannel"
   "github.com/team13/app/ent/customer"
   "github.com/team13/app/ent/product"
   "github.com/team13/app/ent/typeproduct"

) 
// OrderonlineController defines the struct for the orderonline controller
type OrderonlineController struct {
   client *ent.Client
   router gin.IRouter
}

type Orderonline struct{
	Customerid    	   int
	Paymentchannelid   int
	Productid          int
	Typeproductid      int
	Stock			   int
	Accountnumber	   string
	Cvv				   string
	Addedtime       string
}

// CreateOrderonline handles POST requests for adding orderonline entities
// @Summary Create orderonline  
// @Description Create orderonline  
// @ID create-orderonline  
// @Accept   json
// @Produce  json
// @Param orderonline body Orderonline true "Orderonline entity"
// @Success 200 {object} Orderonline  
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /orderonlines [post]
func (ctl *OrderonlineController) CreateOrderonline(c *gin.Context) {
	obj := Orderonline{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "orderonline binding failed",
		})
		return
	}
  
	u, err := ctl.client.Customer.
		Query().
		Where(customer.IDEQ(int(obj.Customerid))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "customer not found",
		})
		return
	}

	pay, err := ctl.client.Paymentchannel.
		Query().
		Where(paymentchannel.IDEQ(int(obj.Paymentchannelid))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "paymentchannel not found",
		})
		return
	}

	p, err := ctl.client.Product.
		Query().
		Where(product.IDEQ(int(obj.Productid))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "product not found",
		})
		return
	}

	tp, err := ctl.client.Typeproduct.
		Query().
		Where(typeproduct.IDEQ(int(obj.Typeproductid))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Typeproduct not found",
		})
		return
	}

	times, err := time.Parse(time.RFC3339, obj.Addedtime)

	ol, err := ctl.client.Orderonline.
		Create().
		SetCustomer(u).
		SetPaymentchannel(pay).
		SetProduct(p).
		SetTypeproduct(tp).
		SetStock(obj.Stock).
		SetAddedtime(times).
		SetAccountnumber(obj.Accountnumber).
		SetCvv(obj.Cvv).
		Save(context.Background())
		
		if err != nil {
			c.JSON(400, gin.H{
				"status" : false,
				"error": err,
			})
			return
		}
	
		c.JSON(200, gin.H{
			"status" : true,
			"data": ol,
		})
	}
// GetOrderonline handles GET requests to retrieve a orderonline entity
// @Summary Get a orderonline entity by ID
// @Description get orderonline by ID
// @ID get-orderonline
// @Produce  json
// @Param id path int true "Orderonline ID"
// @Success 200 {array} ent.Orderonline
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /orderonlines/{id} [get]
func (ctl *OrderonlineController) GetOrderonline(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	pa, err := ctl.client.Orderonline.
		Query().			
		WithCustomer().
		WithPaymentchannel().
		WithTypeproduct().
		WithProduct().
		Where(orderonline.HasProductWith(product.IDEQ(int(id)))).
		All(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, pa)
 }

// SearchOrderonlines handles request to get a search of orderonline entities
// @Summary Search orderonline entities
// @Description Search orderonline entities
// @ID Search-orderonline
// @Produce json
// @Param name  query string false "Name"
// @Param userid query int false "Userid"
// @Success 200 {array} ent.Orderonline
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /order [get]
func (ctl *OrderonlineController) SearchOrderonlines(c *gin.Context) {
	t1 := c.Query("name")
	useridQuery := c.Query("userid")
	uid := 0
	if useridQuery != "" {
		limit64, err := strconv.ParseInt(useridQuery, 10, 64)
		if err == nil {uid = int(limit64)}
	}
	order, err := ctl.client.Orderonline.
		Query().
		Where(orderonline.And(
			orderonline.HasProductWith(product.NameProductContains(t1)),
			orderonline.HasCustomerWith(customer.IDEQ(int(uid))),
			)).
		WithCustomer().
		WithPaymentchannel().
		WithTypeproduct().
		WithProduct().
		All(context.Background())
		if err != nil {
		c.JSON(400, gin.H{"error": err.Error(),})
		return
	}
	c.JSON(200, order)
 }

// DeleteOrderonline handles DELETE requests to delete a orderonline entity
// @Summary Delete a orderonline entity by ID
// @Description get orderonline by ID
// @ID delete-orderonline
// @Produce  json
// @Param id path int true "Orderonline ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /orderonlines/{id} [delete]
func (ctl *OrderonlineController) DeleteOrderonline(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = ctl.client.Orderonline.
		DeleteOneID(int(id)).
		Exec(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{"result": fmt.Sprintf("ok deleted %v", id)})
 }
// UpdateOrderonline handles PUT requests to update a orderonline entity
// @Summary Update a orderonline entity by ID
// @Description update orderonline by ID
// @ID update-orderonline
// @Accept   json
// @Produce  json
// @Param id path int true "Orderonline ID"
// @Param orderonline body ent.Orderonline true "Orderonline entity"
// @Success 200 {object} ent.Orderonline
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /orderonlines/{id} [put]

func (ctl *OrderonlineController) UpdateOrderonline(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
 
	if err != nil {
 
		c.JSON(400, gin.H{
 
			"error": err.Error(),
 
		})
 
		return
 
	}
 
 
	obj := ent.Orderonline{}
 
	if err := c.ShouldBind(&obj); err != nil {
 
		c.JSON(400, gin.H{
 
			"error": "orderonline binding failed",
 
		})
 
		return
 
	}
 
	obj.ID = int(id)
 
	u, err := ctl.client.Orderonline.
 
		UpdateOne(&obj).
 
		Save(context.Background())
 
	if err != nil {
 
		c.JSON(400, gin.H{"error": "update failed",})
 
		return
 
	}
 
 
	c.JSON(200, u)
 
 }
// NewOrderonlineController creates and registers handles for the orderonline controller
func NewOrderonlineController(router gin.IRouter, client *ent.Client) *OrderonlineController {
	uc := &OrderonlineController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
 }
// InitOrderonlineController registers routes to the main engine
 func (ctl *OrderonlineController) register() {
	orderonlines := ctl.router.Group("/orderonlines")
	order := ctl.router.Group("/order")
	order.GET("", ctl.SearchOrderonlines)
	// CRUD
	orderonlines.POST("", ctl.CreateOrderonline)
	orderonlines.GET(":id", ctl.GetOrderonline)
	orderonlines.PUT(":id", ctl.UpdateOrderonline)
	orderonlines.DELETE(":id", ctl.DeleteOrderonline)
 }