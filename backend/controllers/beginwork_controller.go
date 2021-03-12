package controllers
 
import (
   "context"
   "fmt"
   "strconv"
   "github.com/team13/app/ent"
   "github.com/team13/app/ent/beginwork"
   "github.com/gin-gonic/gin"
)
 
// BeginWorkController defines the struct for the beginwork controller
type BeginWorkController struct {
   client *ent.Client
   router gin.IRouter
}

// CreateBeginWork handles POST requests for adding beginwork entities
// @Summary Create beginwork
// @Description Create beginwork
// @ID create-beginwork
// @Accept   json
// @Produce  json
// @Param beginwork body ent.BeginWork true "BeginWork entity"
// @Success 200 {object} ent.BeginWork
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /beginworks [post]
func (ctl *BeginWorkController) CreateBeginWork(c *gin.Context) {
	obj := ent.BeginWork{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "beginwork binding failed",
		})
		return
	}
  
	u, err := ctl.client.BeginWork.
		Create().
		SetBeginWork(obj.BeginWork).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}
  
	c.JSON(200, u)
 }
 
 // GetBeginWork handles GET requests to retrieve a beginwork entity
// @Summary Get a beginwork entity by ID
// @Description get beginwork by ID
// @ID get-beginwork
// @Produce  json
// @Param id path int true "BeginWork ID"
// @Success 200 {object} ent.BeginWork
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /beginworks/{id} [get]
func (ctl *BeginWorkController) GetBeginWork(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	u, err := ctl.client.BeginWork.
		Query().
		Where(beginwork.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	c.JSON(200, u)
 }

 // ListBeginWork handles request to get a list of beginwork entities
// @Summary List beginwork entities
// @Description list beginwork entities
// @ID list-beginwork
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.BeginWork
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /beginworks [get]
func (ctl *BeginWorkController) ListBeginWork(c *gin.Context) {
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
  
	beginworks, err := ctl.client.BeginWork.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
		if err != nil {
		c.JSON(400, gin.H{"error": err.Error(),})
		return
	}
  
	c.JSON(200, beginworks)
 }

 // DeleteBeginWork handles DELETE requests to delete a beginwork entity
// @Summary Delete a beginwork entity by ID
// @Description get beginwork by ID
// @ID delete-beginwork
// @Produce  json
// @Param id path int true "BeginWork ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /beginworks/{id} [delete]
func (ctl *BeginWorkController) DeleteBeginWork(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	err = ctl.client.BeginWork.
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

 // UpdateBeginWork handles PUT requests to update a beginwork entity
// @Summary Update a beginwork entity by ID
// @Description update beginwork by ID
// @ID update-beginwork
// @Accept   json
// @Produce  json
// @Param id path int true "BeginWork ID"
// @Param beginwork body ent.BeginWork true "BeginWork entity"
// @Success 200 {object} ent.BeginWork
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /beginworks/{id} [put]
func (ctl *BeginWorkController) UpdateBeginWork(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	obj := ent.BeginWork{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "beginwork binding failed",
		})
		return
	}
	obj.ID = int(id)
	u, err := ctl.client.BeginWork.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed",})
		return
	}
  
	c.JSON(200, u)
 }

 // NewBeginWorkController creates and registers handles for the beginwork controller
func NewBeginWorkController(router gin.IRouter, client *ent.Client) *BeginWorkController {
	uc := &BeginWorkController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
 }
  
 // InitBeginWorkController registers routes to the main engine
 func (ctl *BeginWorkController) register() {
	beginworks := ctl.router.Group("/beginworks")
  
	beginworks.GET("", ctl.ListBeginWork)
  
	// CRUD
	beginworks.POST("", ctl.CreateBeginWork)
	beginworks.GET(":id", ctl.GetBeginWork)
	beginworks.PUT(":id", ctl.UpdateBeginWork)
	beginworks.DELETE(":id", ctl.DeleteBeginWork)
 }
 