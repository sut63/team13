package controllers
 
import (
   "context"
   "fmt"
   "strconv"
   "github.com/team13/app/ent"
   "github.com/team13/app/ent/getoffwork"
   "github.com/gin-gonic/gin"
)
 
// GetOffWorkController defines the struct for the getoffwork controller
type GetOffWorkController struct {
   client *ent.Client
   router gin.IRouter
}

// CreateGetOffWork handles POST requests for adding getoffwork entities
// @Summary Create getoffwork
// @Description Create getoffwork
// @ID create-getoffwork
// @Accept   json
// @Produce  json
// @Param getoffwork body ent.GetOffWork true "GetOffWork entity"
// @Success 200 {object} ent.GetOffWork
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /getoffworks [post]
func (ctl *GetOffWorkController) CreateGetOffWork(c *gin.Context) {
	obj := ent.GetOffWork{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "getoffwork binding failed",
		})
		return
	}
  
	u, err := ctl.client.GetOffWork.
		Create().
		SetGetOffWork(obj.GetOffWork).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}
  
	c.JSON(200, u)
 }
 
 // GetGetOffWork handles GET requests to retrieve a getoffwork entity
// @Summary Get a getoffwork entity by ID
// @Description get getoffwork by ID
// @ID get-getoffwork
// @Produce  json
// @Param id path int true "GetOffWork ID"
// @Success 200 {object} ent.GetOffWork
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /getoffworks/{id} [get]
func (ctl *GetOffWorkController) GetGetOffWork(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	u, err := ctl.client.GetOffWork.
		Query().
		Where(getoffwork.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	c.JSON(200, u)
 }

 // ListGetOffWork handles request to get a list of getoffwork entities
// @Summary List getoffwork entities
// @Description list getoffwork entities
// @ID list-getoffwork
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.GetOffWork
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /getoffworks [get]
func (ctl *GetOffWorkController) ListGetOffWork(c *gin.Context) {
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
  
	getoffworks, err := ctl.client.GetOffWork.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
		if err != nil {
		c.JSON(400, gin.H{"error": err.Error(),})
		return
	}
  
	c.JSON(200, getoffworks)
 }

 // DeleteGetOffWork handles DELETE requests to delete a getoffwork entity
// @Summary Delete a getoffwork entity by ID
// @Description get getoffwork by ID
// @ID delete-getoffwork
// @Produce  json
// @Param id path int true "GetOffWork ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /getoffworks/{id} [delete]
func (ctl *GetOffWorkController) DeleteGetOffWork(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	err = ctl.client.GetOffWork.
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

 // UpdateGetOffWork handles PUT requests to update a getoffwork entity
// @Summary Update a getoffwork entity by ID
// @Description update getoffwork by ID
// @ID update-getoffwork
// @Accept   json
// @Produce  json
// @Param id path int true "GetOffWork ID"
// @Param getoffwork body ent.GetOffWork true "GetOffWork entity"
// @Success 200 {object} ent.GetOffWork
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /getoffworks/{id} [put]
func (ctl *GetOffWorkController) UpdateGetOffWork(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	obj := ent.GetOffWork{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "getoffwork binding failed",
		})
		return
	}
	obj.ID = int(id)
	u, err := ctl.client.GetOffWork.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed",})
		return
	}
  
	c.JSON(200, u)
 }

 // NewGetOffWorkController creates and registers handles for the getoffwork controller
func NewGetOffWorkController(router gin.IRouter, client *ent.Client) *GetOffWorkController {
	uc := &GetOffWorkController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
 }
  
 // InitGetOffWorkController registers routes to the main engine
 func (ctl *GetOffWorkController) register() {
	getoffworks := ctl.router.Group("/getoffworks")
  
	getoffworks.GET("", ctl.ListGetOffWork)
  
	// CRUD
	getoffworks.POST("", ctl.CreateGetOffWork)
	getoffworks.GET(":id", ctl.GetGetOffWork)
	getoffworks.PUT(":id", ctl.UpdateGetOffWork)
	getoffworks.DELETE(":id", ctl.DeleteGetOffWork)
 }
 