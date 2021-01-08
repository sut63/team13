package controllers

import (
	"context"
	"fmt"
	"strconv"

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
	Discount       int
	Giveaway       int
	Product		   int
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


	po, err := ctl.client.Promotion.
		Create().
		SetDiscount(d).
		SetGiveaway(g).
		SetProduct(p).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, po)
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
		WithDiscount().
		WithGiveaway().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, promotions)
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

	promotions.GET("", ctl.ListPromotion)

	// CRUD
	promotions.POST("", ctl.CreatePromotion)
	promotions.DELETE(":id", ctl.DeletePromotion)
}