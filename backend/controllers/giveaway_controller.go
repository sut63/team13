package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/team13/app/ent"
	"github.com/team13/app/ent/discount"
	"github.com/gin-gonic/gin"
)

// GiveawayController defines the struct for the discount controller
type GiveawayController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateGiveaway handles POST requests for adding discount entities
// @Summary Create discount
// @Description Create discount
// @ID create-discount
// @Accept   json
// @Produce  json
// @Param discount body ent.Giveaway true "Giveaway entity"
// @Success 200 {object} ent.Giveaway
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /giveaway [post]
func (ctl *GiveawayController) CreateGiveaway(c *gin.Context) {
	obj := ent.Giveaway{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "discount binding failed",
		})
		return
	}

	u, err := ctl.client.Giveaway.
		Create().
		SetGiveawayName(obj.GiveawayName).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, u)
}

// GetGiveaway handles GET requests to retrieve a discount entity
// @Summary Get a discount entity by ID
// @Description get discount by ID
// @ID get-discount
// @Produce  json
// @Param id path int true "Giveaway ID"
// @Success 200 {object} ent.Giveaway
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /giveaway/{id} [get]
func (ctl *GiveawayController) GetGiveaway(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	p, err := ctl.client.Giveaway.
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

// ListGiveaway handles request to get a list of discount entities
// @Summary List discount entities
// @Description list discount entities
// @ID list-discount
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Giveaway
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /giveaway [get]
func (ctl *GiveawayController) ListGiveaway(c *gin.Context) {
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

	giveaway, err := ctl.client.Giveaway.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, giveaway)
}

// DeleteGiveaway handles DELETE requests to delete a discount entity
// @Summary Delete a discount entity by ID
// @Description get discount by ID
// @ID delete-discount
// @Produce  json
// @Param id path int true "Giveaway ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /giveaway/{id} [delete]
func (ctl *GiveawayController) DeleteGiveaway(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Giveaway.
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

// UpdateGiveaway handles PUT requests to update a discount entity
// @Summary Update a discount entity by ID
// @Description update discount by ID
// @ID update-discount
// @Accept   json
// @Produce  json
// @Param id path int true "Giveaway ID"
// @Param discount body ent.Giveaway true "Giveaway entity"
// @Success 200 {object} ent.Giveaway
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /giveaway/{id} [put]
func (ctl *GiveawayController) UpdateGiveaway(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Giveaway{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "discount binding failed",
		})
		return
	}
	obj.ID = int(id)
	fmt.Println(obj.ID)
	u, err := ctl.client.Giveaway.
		UpdateOneID(int(id)).
		SetGiveawayName(obj.GiveawayName).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, u)
}

// NewGiveawayController creates and registers handles for the discount controller
func NewGiveawayController(router gin.IRouter, client *ent.Client) *GiveawayController {
	p := &GiveawayController{
		client: client,
		router: router,
	}
	p.register()
	return p
}

// InitGiveawayController registers routes to the main engine
func (ctl *GiveawayController) register() {
	giveaway := ctl.router.Group("/giveaway")

	giveaway.GET("", ctl.ListGiveaway)

	// CRUD
	giveaway.POST("", ctl.CreateGiveaway)
	giveaway.GET(":id", ctl.GetGiveaway)
	giveaway.PUT(":id", ctl.UpdateGiveaway)
	giveaway.DELETE(":id", ctl.DeleteGiveaway)
}
