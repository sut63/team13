package controllers
 
import (
   "context"
   "fmt"
   "strconv"
   "github.com/team13/app/ent"
   "github.com/team13/app/ent/startwork"
   "github.com/gin-gonic/gin"
)
 
// StartWorkController defines the struct for the startwork controller
type StartWorkController struct {
   client *ent.Client
   router gin.IRouter
}

// CreateStartWork handles POST requests for adding startwork entities
// @Summary Create startwork
// @Description Create startwork
// @ID create-startwork
// @Accept   json
// @Produce  json
// @Param startwork body ent.StartWork true "StartWork entity"
// @Success 200 {object} ent.StartWork
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /startworks [post]
func (ctl *StartWorkController) CreateStartWork(c *gin.Context) {
	obj := ent.StartWork{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "startwork binding failed",
		})
		return
	}
  
	u, err := ctl.client.StartWork.
		Create().
		SetStartWork(obj.StartWork).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}
  
	c.JSON(200, u)
 }
 
 // GetStartWork handles GET requests to retrieve a startwork entity
// @Summary Get a startwork entity by ID
// @Description get startwork by ID
// @ID get-startwork
// @Produce  json
// @Param id path int true "StartWork ID"
// @Success 200 {object} ent.StartWork
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /startworks/{id} [get]
func (ctl *StartWorkController) GetStartWork(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	u, err := ctl.client.StartWork.
		Query().
		Where(startwork.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	c.JSON(200, u)
 }

 // ListStartWork handles request to get a list of startwork entities
// @Summary List startwork entities
// @Description list startwork entities
// @ID list-startwork
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.StartWork
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /startworks [get]
func (ctl *StartWorkController) ListStartWork(c *gin.Context) {
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
  
	startworks, err := ctl.client.StartWork.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
		if err != nil {
		c.JSON(400, gin.H{"error": err.Error(),})
		return
	}
  
	c.JSON(200, startworks)
 }

 // DeleteStartWork handles DELETE requests to delete a startwork entity
// @Summary Delete a startwork entity by ID
// @Description get startwork by ID
// @ID delete-startwork
// @Produce  json
// @Param id path int true "StartWork ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /startworks/{id} [delete]
func (ctl *StartWorkController) DeleteStartWork(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	err = ctl.client.StartWork.
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

 // UpdateStartWork handles PUT requests to update a startwork entity
// @Summary Update a startwork entity by ID
// @Description update startwork by ID
// @ID update-startwork
// @Accept   json
// @Produce  json
// @Param id path int true "StartWork ID"
// @Param startwork body ent.StartWork true "StartWork entity"
// @Success 200 {object} ent.StartWork
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /startworks/{id} [put]
func (ctl *StartWorkController) UpdateStartWork(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	obj := ent.StartWork{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "startwork binding failed",
		})
		return
	}
	obj.ID = int(id)
	u, err := ctl.client.StartWork.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed",})
		return
	}
  
	c.JSON(200, u)
 }

 // NewStartWorkController creates and registers handles for the startwork controller
func NewStartWorkController(router gin.IRouter, client *ent.Client) *StartWorkController {
	uc := &StartWorkController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
 }
  
 // InitStartWorkController registers routes to the main engine
 func (ctl *StartWorkController) register() {
	startworks := ctl.router.Group("/startworks")
  
	startworks.GET("", ctl.ListStartWork)
  
	// CRUD
	startworks.POST("", ctl.CreateStartWork)
	startworks.GET(":id", ctl.GetStartWork)
	startworks.PUT(":id", ctl.UpdateStartWork)
	startworks.DELETE(":id", ctl.DeleteStartWork)
 }
 