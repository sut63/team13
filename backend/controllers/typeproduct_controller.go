package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/team13/app/ent"
	"github.com/team13/app/ent/typeproduct"
)

// TypeproductController defines the struct for the typeproduct controller
type TypeproductController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateTypeproduct handles POST requests for adding typeproduct entities
// @Summary Create typeproduct
// @Description Create typeproduct
// @ID create-typeproduct
// @Accept   json
// @Produce  json
// @Param typeproduct body ent.Typeproduct true "Typeproduct entity"
// @Success 200 {object} ent.Typeproduct
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /typeproducts [post]
func (ctl *TypeproductController) CreateTypeproduct(c *gin.Context) {
	obj := ent.Typeproduct{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "typeproduct binding failed",
		})
		return
	}

	u, err := ctl.client.Typeproduct.
		Create().
		SetTypeproduct(obj.Typeproduct).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, u)
}

// GetTypeproduct handles GET requests to retrieve a typeproduct entity
// @Summary Get a typeproduct entity by ID
// @Description get typeproduct by ID
// @ID get-typeproduct
// @Produce  json
// @Param id path int true "Typeproduct ID"
// @Success 200 {object} ent.Typeproduct
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /typeproducts/{id} [get]
func (ctl *TypeproductController) GetTypeproduct(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	u, err := ctl.client.Typeproduct.
		Query().
		Where(typeproduct.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, u)
}

// ListTypeproduct handles request to get a list of typeproduct entities
// @Summary List typeproduct entities
// @Description list typeproduct entities
// @ID list-typeproduct
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Typeproduct
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /typeproducts [get]
func (ctl *TypeproductController) ListTypeproduct(c *gin.Context) {
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

	typeproducts, err := ctl.client.Typeproduct.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, typeproducts)
}

// DeleteTypeproduct handles DELETE requests to delete a typeproduct entity
// @Summary Delete a typeproduct entity by ID
// @Description get typeproduct by ID
// @ID delete-typeproduct
// @Produce  json
// @Param id path int true "Typeproduct ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /typeproducts/{id} [delete]
func (ctl *TypeproductController) DeleteTypeproduct(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Typeproduct.
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

// UpdateTypeproduct handles PUT requests to update a typeproduct entity
// @Summary Update a typeproduct entity by ID
// @Description update typeproduct by ID
// @ID update-typeproduct
// @Accept   json
// @Produce  json
// @Param id path int true "Typeproduct ID"
// @Param typeproduct body ent.Typeproduct true "Typeproduct entity"
// @Success 200 {object} ent.Typeproduct
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /typeproducts/{id} [put]
func (ctl *TypeproductController) UpdateTypeproduct(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Typeproduct{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "typeproduct binding failed",
		})
		return
	}
	obj.ID = int(id)
	u, err := ctl.client.Typeproduct.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, u)
}

// NewTypeproductController creates and registers handles for the typeproduct controller
func NewTypeproductController(router gin.IRouter, client *ent.Client) *TypeproductController {
	uc := &TypeproductController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
}

// InitTypeproductController registers routes to the main engine
func (ctl *TypeproductController) register() {
	typeproducts := ctl.router.Group("/typeproducts")

	typeproducts.GET("", ctl.ListTypeproduct)

	// CRUD
	typeproducts.POST("", ctl.CreateTypeproduct)
	typeproducts.GET(":id", ctl.GetTypeproduct)
	typeproducts.PUT(":id", ctl.UpdateTypeproduct)
	typeproducts.DELETE(":id", ctl.DeleteTypeproduct)
}