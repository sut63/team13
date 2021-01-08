package controllers
 
import (
   "context"
   "fmt"
   "strconv"
   "github.com/team13/app/ent"
   "github.com/team13/app/ent/shift"
   "github.com/gin-gonic/gin"
)
 
// ShiftController defines the struct for the shift controller
type ShiftController struct {
   client *ent.Client
   router gin.IRouter
}

// CreateShift handles POST requests for adding shift entities
// @Summary Create shift
// @Description Create shift
// @ID create-shift
// @Accept   json
// @Produce  json
// @Param shift body ent.Shift true "Shift entity"
// @Success 200 {object} ent.Shift
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /shifts [post]
func (ctl *ShiftController) CreateShift(c *gin.Context) {
	obj := ent.Shift{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "shift binding failed",
		})
		return
	}
  
	u, err := ctl.client.Shift.
		Create().
		SetTimeStart(obj.TimeStart).
		SetTimeEnd(obj.TimeEnd).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}
  
	c.JSON(200, u)
 }
 
 // GetShift handles GET requests to retrieve a shift entity
// @Summary Get a shift entity by ID
// @Description get shift by ID
// @ID get-shift
// @Produce  json
// @Param id path int true "Shift ID"
// @Success 200 {object} ent.Shift
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /shifts/{id} [get]
func (ctl *ShiftController) GetShift(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	u, err := ctl.client.Shift.
		Query().
		Where(shift.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	c.JSON(200, u)
 }

 // ListShift handles request to get a list of shift entities
// @Summary List shift entities
// @Description list shift entities
// @ID list-shift
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Shift
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /shifts [get]
func (ctl *ShiftController) ListShift(c *gin.Context) {
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
  
	shifts, err := ctl.client.Shift.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
		if err != nil {
		c.JSON(400, gin.H{"error": err.Error(),})
		return
	}
  
	c.JSON(200, shifts)
 }

 // DeleteShift handles DELETE requests to delete a shift entity
// @Summary Delete a shift entity by ID
// @Description get shift by ID
// @ID delete-shift
// @Produce  json
// @Param id path int true "Shift ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /shifts/{id} [delete]
func (ctl *ShiftController) DeleteShift(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	err = ctl.client.Shift.
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

 // UpdateShift handles PUT requests to update a shift entity
// @Summary Update a shift entity by ID
// @Description update shift by ID
// @ID update-shift
// @Accept   json
// @Produce  json
// @Param id path int true "Shift ID"
// @Param shift body ent.Shift true "Shift entity"
// @Success 200 {object} ent.Shift
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /shifts/{id} [put]
func (ctl *ShiftController) UpdateShift(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	obj := ent.Shift{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "shift binding failed",
		})
		return
	}
	obj.ID = int(id)
	u, err := ctl.client.Shift.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed",})
		return
	}
  
	c.JSON(200, u)
 }

 // NewShiftController creates and registers handles for the shift controller
func NewShiftController(router gin.IRouter, client *ent.Client) *ShiftController {
	uc := &ShiftController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
 }
  
 // InitShiftController registers routes to the main engine
 func (ctl *ShiftController) register() {
	shifts := ctl.router.Group("/shifts")
  
	shifts.GET("", ctl.ListShift)
  
	// CRUD
	shifts.POST("", ctl.CreateShift)
	shifts.GET(":id", ctl.GetShift)
	shifts.PUT(":id", ctl.UpdateShift)
	shifts.DELETE(":id", ctl.DeleteShift)
 }
 