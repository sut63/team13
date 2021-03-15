package controllers
 
import (
   "context"
   "fmt"
   "strconv"
   "github.com/team13/app/ent"
   "github.com/team13/app/ent/endwork"
   "github.com/gin-gonic/gin"
)
 
// EndWorkController defines the struct for the endwork controller
type EndWorkController struct {
   client *ent.Client
   router gin.IRouter
}

// CreateEndWork handles POST requests for adding endwork entities
// @Summary Create endwork
// @Description Create endwork
// @ID create-endwork
// @Accept   json
// @Produce  json
// @Param endwork body ent.EndWork true "EndWork entity"
// @Success 200 {object} ent.EndWork
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /endworks [post]
func (ctl *EndWorkController) CreateEndWork(c *gin.Context) {
	obj := ent.EndWork{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "endwork binding failed",
		})
		return
	}
  
	u, err := ctl.client.EndWork.
		Create().
		SetEndWork(obj.EndWork).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}
  
	c.JSON(200, u)
 }
 
 // GetEndWork handles GET requests to retrieve a endwork entity
// @Summary Get a endwork entity by ID
// @Description get endwork by ID
// @ID get-endwork
// @Produce  json
// @Param id path int true "EndWork ID"
// @Success 200 {object} ent.EndWork
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /endworks/{id} [get]
func (ctl *EndWorkController) GetEndWork(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	u, err := ctl.client.EndWork.
		Query().
		Where(endwork.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	c.JSON(200, u)
 }

 // ListEndWork handles request to get a list of endwork entities
// @Summary List endwork entities
// @Description list endwork entities
// @ID list-endwork
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.EndWork
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /endworks [get]
func (ctl *EndWorkController) ListEndWork(c *gin.Context) {
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
  
	endworks, err := ctl.client.EndWork.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
		if err != nil {
		c.JSON(400, gin.H{"error": err.Error(),})
		return
	}
  
	c.JSON(200, endworks)
 }

 // DeleteEndWork handles DELETE requests to delete a endwork entity
// @Summary Delete a endwork entity by ID
// @Description get endwork by ID
// @ID delete-endwork
// @Produce  json
// @Param id path int true "EndWork ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /endworks/{id} [delete]
func (ctl *EndWorkController) DeleteEndWork(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	err = ctl.client.EndWork.
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

 // UpdateEndWork handles PUT requests to update a endwork entity
// @Summary Update a endwork entity by ID
// @Description update endwork by ID
// @ID update-endwork
// @Accept   json
// @Produce  json
// @Param id path int true "EndWork ID"
// @Param endwork body ent.EndWork true "EndWork entity"
// @Success 200 {object} ent.EndWork
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /endworks/{id} [put]
func (ctl *EndWorkController) UpdateEndWork(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	obj := ent.EndWork{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "endwork binding failed",
		})
		return
	}
	obj.ID = int(id)
	u, err := ctl.client.EndWork.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed",})
		return
	}
  
	c.JSON(200, u)
 }

 // NewEndWorkController creates and registers handles for the endwork controller
func NewEndWorkController(router gin.IRouter, client *ent.Client) *EndWorkController {
	uc := &EndWorkController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
 }
  
 // InitEndWorkController registers routes to the main engine
 func (ctl *EndWorkController) register() {
	endworks := ctl.router.Group("/endworks")
  
	endworks.GET("", ctl.ListEndWork)
  
	// CRUD
	endworks.POST("", ctl.CreateEndWork)
	endworks.GET(":id", ctl.GetEndWork)
	endworks.PUT(":id", ctl.UpdateEndWork)
	endworks.DELETE(":id", ctl.DeleteEndWork)
 }
 