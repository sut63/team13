package controllers
import (
   "context"
   "fmt"
   "strconv"
   "github.com/team13/app/ent"
   "github.com/team13/app/ent/paymentchannel"
   "github.com/gin-gonic/gin"
)
// PaymentchannelController defines the struct for the paymentchannel controller
type PaymentchannelController struct {
   client *ent.Client
   router gin.IRouter
}
// CreatePaymentchannel handles POST requests for adding paymentchannel entities
// @Summary Create paymentchannel
// @Description Create paymentchannel
// @ID create-paymentchannel  
// @Accept   json
// @Produce  json
// @Param paymentchannel   body ent.Paymentchannel true "Paymentchannel entity"
// @Success 200 {object} ent.Paymentchannel  
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /paymentchannels [post]
func (ctl *PaymentchannelController) CreatePaymentchannel(c *gin.Context) {
	obj := ent.Paymentchannel{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "paymentchannel binding failed",
		})
		return
	}
  
	u, err := ctl.client.Paymentchannel.
		Create().
		SetPaymentchannelname(obj.Paymentchannelname).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}
  
	c.JSON(200, u)
 }
// GetPaymentchannel handles GET requests to retrieve a paymentchannel entity
// @Summary Get a paymentchannel entity by ID
// @Description get paymentchannel by ID
// @ID get-paymentchannel
// @Produce  json
// @Param id path int true "Paymentchannel ID"
// @Success 200 {object} ent.Paymentchannel
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /paymentchannels/{id} [get]
func (ctl *PaymentchannelController) GetPaymentchannel(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	u, err := ctl.client.Paymentchannel.
		Query().
		Where(paymentchannel.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	c.JSON(200, u)
 }
 // ListPaymentchannel handles request to get a list of paymentchannel entities
// @Summary List paymentchannel entities
// @Description list paymentchannel entities
// @ID list-paymentchannel
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Paymentchannel
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /paymentchannels [get]
func (ctl *PaymentchannelController) ListPaymentchannel(c *gin.Context) {
	limitQuery := c.Query("limit")
	limit := 10
	if limitQuery != "" {
		limit64, err := strconv.ParseInt(limitQuery, 10, 64)
		if err == nil {limit = int(limit64)}
	}
  
	offsetQuery := c.Query("offset")
	offset := 0
	if offsetQuery != "" {
		offset64, err := strconv.ParseInt(offsetQuery, 10, 64)
		if err == nil {offset = int(offset64)}
	}
  
	paymentchannels, err := ctl.client.Paymentchannel.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
		if err != nil {
		c.JSON(400, gin.H{"error": err.Error(),})
		return
	}
  
	c.JSON(200, paymentchannels)
 }
 // DeletePaymentchannel handles DELETE requests to delete a paymentchannel entity
// @Summary Delete a paymentchannel entity by ID
// @Description get paymentchannel by ID
// @ID delete-paymentchannel
// @Produce  json
// @Param id path int true "Paymentchannel ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /paymentchannels/{id} [delete]

func (ctl *PaymentchannelController) DeletePaymentchannel(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
 
	if err != nil {
 
		c.JSON(400, gin.H{
 
			"error": err.Error(),
 
		})
 
		return
 
	}
 
 
	err = ctl.client.Paymentchannel.
 
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
// UpdatePaymentchannel handles PUT requests to update a paymentchannel entity
// @Summary Update a paymentchannel entity by ID
// @Description update paymentchannel by ID
// @ID update-paymentchannel
// @Accept   json
// @Produce  json
// @Param id path int true "Paymentchannel ID"
// @Param paymentchannel body ent.Paymentchannel true "Paymentchannel entity"
// @Success 200 {object} ent.Paymentchannel
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /paymentchannels/{id} [put]

func (ctl *PaymentchannelController) UpdatePaymentchannel(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
 
	if err != nil {
 
		c.JSON(400, gin.H{
 
			"error": err.Error(),
 
		})
 
		return
 
	}
 
 
	obj := ent.Paymentchannel{}
 
	if err := c.ShouldBind(&obj); err != nil {
 
		c.JSON(400, gin.H{
 
			"error": "paymentchannel binding failed",
 
		})
 
		return
 
	}
 
	obj.ID = int(id)
 
	u, err := ctl.client.Paymentchannel.
 
		UpdateOne(&obj).
 
		Save(context.Background())
 
	if err != nil {
 
		c.JSON(400, gin.H{"error": "update failed",})
 
		return
 
	}
 
 
	c.JSON(200, u)
 
 }
 // NewPaymentchannelController creates and registers handles for the paymentchannel controller
func NewPaymentchannelController(router gin.IRouter, client *ent.Client) *PaymentchannelController {
	uc := &PaymentchannelController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
 }
  
 // InitPaymentchannelController registers routes to the main engine
 func (ctl *PaymentchannelController) register() {
	paymentchannels := ctl.router.Group("/paymentchannels")
	paymentchannels.GET("", ctl.ListPaymentchannel)
	
	// CRUD
	paymentchannels.POST("", ctl.CreatePaymentchannel)
	paymentchannels.GET(":id", ctl.GetPaymentchannel)
	paymentchannels.PUT(":id", ctl.UpdatePaymentchannel)
	paymentchannels.DELETE(":id", ctl.DeletePaymentchannel)
 }