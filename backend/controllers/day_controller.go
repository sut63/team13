package controllers
 
import (
   "context"
   "fmt"
   "strconv"
   "github.com/team13/app/ent"
   "github.com/team13/app/ent/day"
   "github.com/gin-gonic/gin"
)
 
// DayController defines the struct for the day controller
type DayController struct {
   client *ent.Client
   router gin.IRouter
}

// CreateDay handles POST requests for adding day entities
// @Summary Create day
// @Description Create day
// @ID create-day
// @Accept   json
// @Produce  json
// @Param day body ent.Day true "Day entity"
// @Success 200 {object} ent.Day
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /days [post]
func (ctl *DayController) CreateDay(c *gin.Context) {
	obj := ent.Day{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "day binding failed",
		})
		return
	}
  
	u, err := ctl.client.Day.
		Create().
		SetDay(obj.Day).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}
  
	c.JSON(200, u)
 }
 
 // GetDay handles GET requests to retrieve a day entity
// @Summary Get a day entity by ID
// @Description get day by ID
// @ID get-day
// @Produce  json
// @Param id path int true "Day ID"
// @Success 200 {object} ent.Day
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /days/{id} [get]
func (ctl *DayController) GetDay(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	u, err := ctl.client.Day.
		Query().
		Where(day.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	c.JSON(200, u)
 }

 // ListDay handles request to get a list of day entities
// @Summary List day entities
// @Description list day entities
// @ID list-day
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Day
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /days [get]
func (ctl *DayController) ListDay(c *gin.Context) {
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
  
	days, err := ctl.client.Day.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
		if err != nil {
		c.JSON(400, gin.H{"error": err.Error(),})
		return
	}
  
	c.JSON(200, days)
 }

 // DeleteDay handles DELETE requests to delete a day entity
// @Summary Delete a day entity by ID
// @Description get day by ID
// @ID delete-day
// @Produce  json
// @Param id path int true "Day ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /days/{id} [delete]
func (ctl *DayController) DeleteDay(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	err = ctl.client.Day.
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

 // UpdateDay handles PUT requests to update a day entity
// @Summary Update a day entity by ID
// @Description update day by ID
// @ID update-day
// @Accept   json
// @Produce  json
// @Param id path int true "Day ID"
// @Param day body ent.Day true "Day entity"
// @Success 200 {object} ent.Day
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /days/{id} [put]
func (ctl *DayController) UpdateDay(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	obj := ent.Day{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "day binding failed",
		})
		return
	}
	obj.ID = int(id)
	u, err := ctl.client.Day.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed",})
		return
	}
  
	c.JSON(200, u)
 }

 // NewDayController creates and registers handles for the day controller
func NewDayController(router gin.IRouter, client *ent.Client) *DayController {
	uc := &DayController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
 }
  
 // InitDayController registers routes to the main engine
 func (ctl *DayController) register() {
	days := ctl.router.Group("/days")
  
	days.GET("", ctl.ListDay)
  
	// CRUD
	days.POST("", ctl.CreateDay)
	days.GET(":id", ctl.GetDay)
	days.PUT(":id", ctl.UpdateDay)
	days.DELETE(":id", ctl.DeleteDay)
 }
 