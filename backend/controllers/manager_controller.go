package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/team13/app/ent"
	"github.com/team13/app/ent/manager"
)

// ManagerController defines the struct for the manager controller
type ManagerController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateManager handles POST requests for adding manager entities
// @Summary Create manager
// @Description Create manager
// @ID create-manager
// @Accept   json
// @Produce  json
// @Param manager body ent.Manager true "Manager entity"
// @Success 200 {object} ent.Manager
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /managers [post]
func (ctl *ManagerController) CreateManager(c *gin.Context) {
	obj := ent.Manager{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "manager binding failed",
		})
		return
	}

	u, err := ctl.client.Manager.
		Create().
		SetName(obj.Name).
		SetEmail(obj.Email).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, u)
}

// GetManager handles GET requests to retrieve a manager entity
// @Summary Get a manager entity by ID
// @Description get manager by ID
// @ID get-manager
// @Produce  json
// @Param id path int true "Manager ID"
// @Success 200 {object} ent.Manager
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /managers/{id} [get]
func (ctl *ManagerController) GetManager(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	u, err := ctl.client.Manager.
		Query().
		Where(manager.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, u)
}

// ListManager handles request to get a list of manager entities
// @Summary List manager entities
// @Description list manager entities
// @ID list-manager
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Manager
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /managers [get]
func (ctl *ManagerController) ListManager(c *gin.Context) {
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

	managers, err := ctl.client.Manager.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, managers)
}

// DeleteManager handles DELETE requests to delete a manager entity
// @Summary Delete a manager entity by ID
// @Description get manager by ID
// @ID delete-manager
// @Produce  json
// @Param id path int true "Manager ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /managers/{id} [delete]
func (ctl *ManagerController) DeleteManager(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Manager.
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

// UpdateManager handles PUT requests to update a manager entity
// @Summary Update a manager entity by ID
// @Description update manager by ID
// @ID update-manager
// @Accept   json
// @Produce  json
// @Param id path int true "Manager ID"
// @Param manager body ent.Manager true "Manager entity"
// @Success 200 {object} ent.Manager
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /managers/{id} [put]
func (ctl *ManagerController) UpdateManager(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Manager{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "manager binding failed",
		})
		return
	}
	obj.ID = int(id)
	u, err := ctl.client.Manager.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, u)
}

// NewManagerController creates and registers handles for the manager controller
func NewManagerController(router gin.IRouter, client *ent.Client) *ManagerController {
	uc := &ManagerController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
}

// InitManagerController registers routes to the main engine
func (ctl *ManagerController) register() {
	managers := ctl.router.Group("/managers")

	managers.GET("", ctl.ListManager)

	// CRUD
	managers.POST("", ctl.CreateManager)
	managers.GET(":id", ctl.GetManager)
	managers.PUT(":id", ctl.UpdateManager)
	managers.DELETE(":id", ctl.DeleteManager)
}