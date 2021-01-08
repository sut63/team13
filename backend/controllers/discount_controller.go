package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/team13/app/ent"
	"github.com/team13/app/ent/discount"
	"github.com/gin-gonic/gin"
)

// DiscountController defines the struct for the discount controller
type DiscountController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateDiscount handles POST requests for adding discount entities
// @Summary Create discount
// @Description Create discount
// @ID create-discount
// @Accept   json
// @Produce  json
// @Param discount body ent.Discount true "Discount entity"
// @Success 200 {object} ent.Discount
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /discounts [post]
func (ctl *DiscountController) CreateDiscount(c *gin.Context) {
	obj := ent.Discount{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "discount binding failed",
		})
		return
	}

	u, err := ctl.client.Discount.
		Create().
		SetSale(obj.Sale).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, u)
}

// GetDiscount handles GET requests to retrieve a discount entity
// @Summary Get a discount entity by ID
// @Description get discount by ID
// @ID get-discount
// @Produce  json
// @Param id path int true "Discount ID"
// @Success 200 {object} ent.Discount
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /discounts/{id} [get]
func (ctl *DiscountController) GetDiscount(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	p, err := ctl.client.Discount.
		Query().
		Where(discount.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, p)
}

// ListDiscount handles request to get a list of discount entities
// @Summary List discount entities
// @Description list discount entities
// @ID list-discount
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Discount
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /discounts [get]
func (ctl *DiscountController) ListDiscount(c *gin.Context) {
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

	discounts, err := ctl.client.Discount.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, discounts)
}

// DeleteDiscount handles DELETE requests to delete a discount entity
// @Summary Delete a discount entity by ID
// @Description get discount by ID
// @ID delete-discount
// @Produce  json
// @Param id path int true "Discount ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /discounts/{id} [delete]
func (ctl *DiscountController) DeleteDiscount(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Discount.
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

// UpdateDiscount handles PUT requests to update a discount entity
// @Summary Update a discount entity by ID
// @Description update discount by ID
// @ID update-discount
// @Accept   json
// @Produce  json
// @Param id path int true "Discount ID"
// @Param discount body ent.Discount true "Discount entity"
// @Success 200 {object} ent.Discount
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /discounts/{id} [put]
func (ctl *DiscountController) UpdateDiscount(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Discount{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "discount binding failed",
		})
		return
	}
	obj.ID = int(id)
	fmt.Println(obj.ID)
	u, err := ctl.client.Discount.
		UpdateOneID(int(id)).
		SetSale(obj.Sale).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, u)
}

// NewDiscountController creates and registers handles for the discount controller
func NewDiscountController(router gin.IRouter, client *ent.Client) *DiscountController {
	p := &DiscountController{
		client: client,
		router: router,
	}
	p.register()
	return p
}

// InitDiscountController registers routes to the main engine
func (ctl *DiscountController) register() {
	discounts := ctl.router.Group("/discounts")

	discounts.GET("", ctl.ListDiscount)

	// CRUD
	discounts.POST("", ctl.CreateDiscount)
	discounts.GET(":id", ctl.GetDiscount)
	discounts.PUT(":id", ctl.UpdateDiscount)
	discounts.DELETE(":id", ctl.DeleteDiscount)
}
