package controllers

import (
	"context"
	"fmt"
	"strconv"
	"time"
	"github.com/team13/app/ent"
	"github.com/team13/app/ent/discount"
	"github.com/team13/app/ent/giveaway"
	"github.com/team13/app/ent/promotion"
	"github.com/team13/app/ent/product"
	"github.com/gin-gonic/gin"
)

// PromotionController defines the struct for the promotion controller
type PromotionController struct {
	client *ent.Client
	router gin.IRouter
}

// Promotion defines the struct for the promotion
type Promotion struct {
	Discount       		int
	Giveaway       		int
	Product		   		int
	PromotionName  		string
	StartPromotion      string
	EndPromotion        string
	Price		   		float64
}

// CreatePromotion handles POST requests for adding promotion entities
// @Summary Create promotion
// @Description Create promotion
// @ID create-promotion
// @Accept   json
// @Produce  json
// @Param promotion body Promotion true "Promotion entity"
// @Success 200 {object} Promotion
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /promotions [post]
func (ctl *PromotionController) CreatePromotion(c *gin.Context) {
	obj := Promotion{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "promotion binding failed",
		})
		return
	}

	d, err := ctl.client.Discount.
		Query().
		Where(discount.IDEQ(int(obj.Discount))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "discount not found",
		})
		return
	}

	g, err := ctl.client.Giveaway.
		Query().
		Where(giveaway.IDEQ(int(obj.Giveaway))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Giveaway diagnostic  not found",
		})
		return
	}

	p, err := ctl.client.Product.
		Query().
		Where(product.IDEQ(int(obj.Product))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Product diagnostic  not found",
		})
		return
	}
	print(obj.StartPromotion)
	if obj.StartPromotion == ":00+07:00" || obj.EndPromotion == ":00+07:00" {
		c.JSON(400, gin.H{
			"status" : false,
			"error": "timepromotion",
		})
		return
	}
	Stimes, err := time.Parse(time.RFC3339, obj.StartPromotion)
	Etimes, err := time.Parse(time.RFC3339, obj.EndPromotion)

	po, err := ctl.client.Promotion.
		Create().
		SetPromotionName(obj.PromotionName).
		SetProduct(p).
		SetGive(g).
		SetSale(d).
		SetStartPromotion(Stimes).
		SetEndPromotion(Etimes).
		SetPrice(obj.Price).
	
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
		"data": po,
	})
}

// GetPromotion handles GET requests to retrieve a promotion entity
// @Summary Get a promotion entity 
// @Description get promotion 
// @ID get-promotion
// @Produce  json
// @Param product  query string false "Product"
// @Success 200 {array} ent.Promotion
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /promo [get]
func (ctl *PromotionController) GetPromotion(c *gin.Context) {
	t1 := c.Query("product")

	pa, err := ctl.client.Promotion.
		Query().
		WithGive().
		WithProduct().
		WithSale().
		Where(promotion.HasProductWith(product.NameProductContains(t1))).
		All(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, pa)
 }

// ListPromotion handles request to get a list of promotion entities
// @Summary List promotion entities
// @Description list promotion entities
// @ID list-promotion
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Promotion
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /promotions [get]
func (ctl *PromotionController) ListPromotion(c *gin.Context) {
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

	promotions, err := ctl.client.Promotion.
		Query().
		WithSale().
		WithGive().
		WithProduct().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, promotions)
}

// UpdatePromotion handles PUT requests to update a promotion entity
// @Summary Update a promotion entity by ID
// @Description update promotion by ID
// @ID update-promotion
// @Accept   json
// @Produce  json
// @Param id path int true "Promotion ID"
// @Param promotion body ent.Promotion true "Promotion entity"
// @Success 200 {object} ent.Promotion
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /promotions/{id} [put]

func (ctl *PromotionController) UpdatePromotion(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
 
	if err != nil {
 
		c.JSON(400, gin.H{
 
			"error": err.Error(),
 
		})
 
		return
 
	}
 
 
	obj := ent.Promotion{}
 
	if err := c.ShouldBind(&obj); err != nil {
 
		c.JSON(400, gin.H{
 
			"error": "promotion binding failed",
 
		})
 
		return
 
	}
 
	obj.ID = int(id)
 
	u, err := ctl.client.Promotion.
 
		UpdateOne(&obj).
 
		Save(context.Background())
 
	if err != nil {
 
		c.JSON(400, gin.H{"error": "update failed",})
 
		return
 
	}
 
 
	c.JSON(200, u)
 
 }

// DeletePromotion handles DELETE requests to delete a promotion entity
// @Summary Delete a promotion entity by ID
// @Description get promotion by ID
// @ID delete-promotion
// @Produce  json
// @Param id path int true "Promotion ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /promotions/{id} [delete]
func (ctl *PromotionController) DeletePromotion(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Promotion.
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

// NewPromotionController creates and registers handles for the promotion controller
func NewPromotionController(router gin.IRouter, client *ent.Client) *PromotionController {
	drc := &PromotionController{
		client: client,
		router: router,
	}
	drc.register()
	return drc
}

// InitUserController registers routes to the main engine
func (ctl *PromotionController) register() {
	promotions := ctl.router.Group("/promotions")
	pro := ctl.router.Group("/promo")
	promotions.GET("", ctl.ListPromotion)
	// CRUD
	promotions.POST("", ctl.CreatePromotion)
	pro.GET("", ctl.GetPromotion)
	promotions.PUT(":id", ctl.UpdatePromotion)
	promotions.DELETE(":id", ctl.DeletePromotion)
}