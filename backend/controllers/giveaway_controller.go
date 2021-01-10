package controllers
 
import (
   "context"
   "fmt"
   "strconv"
   "github.com/team13/app/ent"
   "github.com/team13/app/ent/giveaway"
   "github.com/gin-gonic/gin"
)
 
// GiveawayController defines the struct for the giveaway controller
type GiveawayController struct {
   client *ent.Client
   router gin.IRouter
}

// CreateGiveaway handles POST requests for adding giveaway entities
// @Summary Create giveaway
// @Description Create giveaway
// @ID create-giveaway
// @Accept   json
// @Produce  json
// @Param giveaway body ent.Giveaway true "Giveaway entity"
// @Success 200 {object} ent.Giveaway
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /giveaways [post]
func (ctl *GiveawayController) CreateGiveaway(c *gin.Context) {
	obj := ent.Giveaway{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "giveaway binding failed",
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

 // GetGiveaway handles GET requests to retrieve a giveaway entity
// @Summary Get a giveaway entity by ID
// @Description get giveaway by ID
// @ID get-giveaway
// @Produce  json
// @Param id path int true "Giveaway ID"
// @Success 200 {object} ent.Giveaway
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /giveaways/{id} [get]
func (ctl *GiveawayController) GetGiveaway(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	u, err := ctl.client.Giveaway.
		Query().
		Where(giveaway.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	c.JSON(200, u)
 }

 // ListGiveaway handles request to get a list of giveaway entities
// @Summary List giveaway entities
// @Description list giveaway entities
// @ID list-giveaway
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Giveaway
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /giveaways [get]
func (ctl *GiveawayController) ListGiveaway(c *gin.Context) {
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
  
	giveaways, err := ctl.client.Giveaway.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
		if err != nil {
		c.JSON(400, gin.H{"error": err.Error(),})
		return
	}
  
	c.JSON(200, giveaways)
 }

 // DeleteGiveaway handles DELETE requests to delete a giveaway entity
// @Summary Delete a giveaway entity by ID
// @Description get giveaway by ID
// @ID delete-giveaway
// @Produce  json
// @Param id path int true "Giveaway ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /giveaways/{id} [delete]
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

 // UpdateGiveaway handles PUT requests to update a giveaway entity
// @Summary Update a giveaway entity by ID
// @Description update giveaway by ID
// @ID update-giveaway
// @Accept   json
// @Produce  json
// @Param id path int true "Giveaway ID"
// @Param giveaway body ent.Giveaway true "Giveaway entity"
// @Success 200 {object} ent.Giveaway
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /giveaways/{id} [put]
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
			"error": "giveaway binding failed",
		})
		return
	}
	obj.ID = int(id)
	u, err := ctl.client.Giveaway.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed",})
		return
	}
  
	c.JSON(200, u)
 }

 // NewGiveawayController creates and registers handles for the giveaway controller
func NewGiveawayController(router gin.IRouter, client *ent.Client) *GiveawayController {
	uc := &GiveawayController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
 }
  
 // InitGiveawayController registers routes to the main engine
 func (ctl *GiveawayController) register() {
	giveaways := ctl.router.Group("/giveaways")
  
	giveaways.GET("", ctl.ListGiveaway)
  
	// CRUD
	giveaways.POST("", ctl.CreateGiveaway)
	giveaways.GET(":id", ctl.GetGiveaway)
	giveaways.PUT(":id", ctl.UpdateGiveaway)
	giveaways.DELETE(":id", ctl.DeleteGiveaway)
 }
 
