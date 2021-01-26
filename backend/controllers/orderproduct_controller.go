package controllers

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/team13/app/ent"
	"github.com/team13/app/ent/typeproduct"
	"github.com/team13/app/ent/company"
	"github.com/team13/app/ent/manager"
	"github.com/team13/app/ent/product"
	"github.com/team13/app/ent/orderproduct"
	
)

// OrderproductController defines the struct for the orderproduct controller
type OrderproductController struct {
	client *ent.Client
	router gin.IRouter
}

// Orderproduct defines the struct for the orderproduct
type Orderproduct struct {
	ManagerID       int
	TypeproductID   int
	ProductID  		int
	CompanyID       int
	Stock			int
	Addedtime       string
	Shipment      string
	Detail        string
}

// CreateOrderproduct handles POST requests for adding orderproduct entities
// @Summary Create orderproduct
// @Description Create orderproduct
// @ID create-orderproduct
// @Accept   json
// @Produce  json
// @Param orderproduct body Orderproduct true "Orderproduct entity"
// @Success 200 {object} Orderproduct
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /orderproducts [post]
func (ctl *OrderproductController) CreateOrderproduct(c *gin.Context) {
	obj := Orderproduct{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "orderproduct binding failed",
		})
		return
	}

	mn, err := ctl.client.Manager.
		Query().
		Where(manager.IDEQ(int(obj.ManagerID))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "manager not found",
		})
		return
	}

	tp, err := ctl.client.Typeproduct.
		Query().
		Where(typeproduct.IDEQ(int(obj.TypeproductID))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "typeproduct  not found",
		})
		return
	}        

	pr, err := ctl.client.Product.
		Query().
		Where(product.IDEQ(int(obj.ProductID))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "product   not found",
		})
		return
	}

	cp, err := ctl.client.Company.
		Query().
		Where(company.IDEQ(int(obj.CompanyID))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "company   not found",
		})
		return
	}


	settime := time.Now().Format("2006-01-02T15:04:05Z07:00")
	time, err := time.Parse(time.RFC3339, settime)

	
	sa, err := ctl.client.Orderproduct.
		Create().
		SetManagers(mn).
		SetTypeproduct(tp).
		SetProduct(pr).
		SetCompany(cp).
		SetStock(obj.Stock).
		SetAddedtime(time).
		SetShipment(obj.Shipment).
		SetDetail(obj.Detail).
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
			"data": sa,
		})
}

// GetOrderproduct handles GET requests to retrieve a orderproduct entity
// @Summary Get a orderproduct entity by ID
// @Description get orderproduct by ID
// @ID get-orderproduct
// @Produce  json
// @Param id path int true "Orderproduct ID"
// @Success 200 {object} ent.Orderproduct
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /orderproducts/{id} [get]
func (ctl *OrderproductController) GetOrderproduct(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	u, err := ctl.client.Orderproduct.
		Query().
		WithProduct().
		Where(orderproduct.HasProductWith(product.IDEQ(int(id)))).
		All(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, u)
}

// ListOrderproduct handles request to get a list of orderproduct entities
// @Summary List orderproduct entities
// @Description list orderproduct entities
// @ID list-orderproduct
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Orderproduct
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /orderproducts [get]
func (ctl *OrderproductController) ListOrderproduct(c *gin.Context) {
	limitQuery := c.Query("limit")
	limit := 10
	if limitQuery != "" {
		limit64, err := strconv.ParseInt(limitQuery, 10, 64)
		if err == nil {
			limit = int(limit64)
		}
	}

	offsetQuery := c.Query("offset")
	offset := 0
	if offsetQuery != "" {
		offset64, err := strconv.ParseInt(offsetQuery, 10, 64)
		if err == nil {
			offset = int(offset64)
		}
	}

	orderproducts, err := ctl.client.Orderproduct.
		Query().
		WithManagers().
		WithTypeproduct().
		WithProduct().
		WithCompany().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, orderproducts)
}

// DeleteOrderproduct handles DELETE requests to delete a orderproduct entity
// @Summary Delete a orderproduct entity by ID
// @Description get orderproduct by ID
// @ID delete-orderproduct
// @Produce  json
// @Param id path int true "Orderproduct ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /orderproducts/{id} [delete]
func (ctl *OrderproductController) DeleteOrderproduct(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Orderproduct.
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

// UpdateOrderproduct handles PUT requests to update a orderproduct entity
// @Summary Update a orderproduct entity by ID
// @Description update orderproduct by ID
// @ID update-orderproduct
// @Accept   json
// @Produce  json
// @Param id path int true "Orderproduct ID"
// @Param orderproduct body ent.Orderproduct true "Orderproduct entity"
// @Success 200 {object} ent.Orderproduct
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /orderproducts/{id} [put]
func (ctl *OrderproductController) UpdateOrderproduct(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Orderproduct{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "orderproduct binding failed",
		})
		return
	}
	obj.ID = int(id)
	u, err := ctl.client.Orderproduct.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, u)
}

// NewOrderproductController creates and registers handles for the orderproduct controller
func NewOrderproductController(router gin.IRouter, client *ent.Client) *OrderproductController {
	uc := &OrderproductController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
}

// InitOrderproductController registers routes to the main engine
func (ctl *OrderproductController) register() {
	orderproducts := ctl.router.Group("/orderproducts")

	orderproducts.GET("", ctl.ListOrderproduct)

	// CRUD
	orderproducts.POST("", ctl.CreateOrderproduct)
	orderproducts.GET(":id", ctl.GetOrderproduct)
	orderproducts.PUT(":id", ctl.UpdateOrderproduct)
	orderproducts.DELETE(":id", ctl.DeleteOrderproduct)
}