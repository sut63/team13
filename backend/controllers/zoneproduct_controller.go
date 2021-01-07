package controllers
 
import (
   "context"
   "fmt"
   "strconv"
   "github.com/team13/app/ent"
   "github.com/team13/app/ent/zoneproduct"
   "github.com/gin-gonic/gin"
)
 
// ZoneproductController defines the struct for the zoneproduct controller
type ZoneproductController struct {
   client *ent.Client
   router gin.IRouter
}


// CreateZoneproduct handles POST requests for adding zoneproduct entities
// @Summary Create zoneproduct
// @Description Create zoneproduct
// @ID create-zoneproduct
// @Accept   json
// @Produce  json
// @Param zoneproduct body ent.Zoneproduct true "Zoneproduct entity"
// @Success 200 {object} ent.Zoneproduct
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /zoneproducts [post]
func (ctl *ZoneproductController) CreateZoneproduct(c *gin.Context) {
    obj := ent.Zoneproduct{}
    if err := c.ShouldBind(&obj); err != nil {
        c.JSON(400, gin.H{
            "error": "zoneproduct binding failed",
        })
        return
    }
  
    z, err := ctl.client.Zoneproduct.
    
			 Create().
			 SetZone(obj.Zone).
			 Save(context.Background())

    if err != nil {
        c.JSON(400, gin.H{
            "error": "saving failed",
        })
        return
    }
  
    c.JSON(200, z)
 }
// GetZoneproduct handles GET requests to retrieve a zoneproduct entity
// @Summary Get a zoneproduct entity by ID
// @Description get zoneproduct by ID
// @ID get-zoneproduct
// @Produce  json
// @Param id path int true "Zoneproduct ID"
// @Success 200 {object} ent.Zoneproduct
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /zoneproducts/{id} [get]
func (ctl *ZoneproductController) GetZoneproduct(c *gin.Context) {
    id, err := strconv.ParseInt(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(400, gin.H{
            "error": err.Error(),
        })
        return
    }
  
    z, err := ctl.client.Zoneproduct.
        Query().
        Where(zoneproduct.IDEQ(int(id))).
        Only(context.Background())
    if err != nil {
        c.JSON(404, gin.H{
            "error": err.Error(),
        })
        return
    }
  
    c.JSON(200, z)
 }
// ListZoneproduct handles request to get a list of zoneproduct entities
// @Summary List zoneproduct entities
// @Description list zoneproduct entities
// @ID list-zoneproduct
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Zoneproduct
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /zoneproducts [get]
func (ctl *ZoneproductController) ListZoneproduct(c *gin.Context) {
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
  
    zoneproducts, err := ctl.client.Zoneproduct.
        Query().
        Limit(limit).
        Offset(offset).
        All(context.Background())
        if err != nil {
        c.JSON(400, gin.H{"error": err.Error(),})
        return
    }
  
    c.JSON(200, zoneproducts)
 }
// DeleteZoneproduct handles DELETE requests to delete a zoneproduct entity
// @Summary Delete a zoneproduct entity by ID
// @Description get zoneproduct by ID
// @ID delete-zoneproduct
// @Produce  json
// @Param id path int true "Zoneproduct ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /zoneproducts/{id} [delete]
func (ctl *ZoneproductController) DeleteZoneproduct(c *gin.Context) {
    id, err := strconv.ParseInt(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(400, gin.H{
            "error": err.Error(),
        })
        return
    }
  
    err = ctl.client.Zoneproduct.
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
// UpdateZoneproduct handles PUT requests to update a zoneproduct entity
// @Summary Update a zoneproduct entity by ID
// @Description update zoneproduct by ID
// @ID update-zoneproduct
// @Accept   json
// @Produce  json
// @Param id path int true "Zoneproduct ID"
// @Param zoneproduct body ent.Zoneproduct true "Zoneproduct entity"
// @Success 200 {object} ent.Zoneproduct
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /zoneproducts/{id} [put]
func (ctl *ZoneproductController) UpdateZoneproduct(c *gin.Context) {
    id, err := strconv.ParseInt(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(400, gin.H{
            "error": err.Error(),
        })
        return
    }
  
    obj := ent.Zoneproduct{}
    if err := c.ShouldBind(&obj); err != nil {
        c.JSON(400, gin.H{
            "error": "zoneproduct binding failed",
        })
        return
    }
    obj.ID = int(id)
    z, err := ctl.client.Zoneproduct.
        UpdateOne(&obj).
        Save(context.Background())
    if err != nil {
        c.JSON(400, gin.H{"error": "update failed",})
        return
    }
  
    c.JSON(200, z)
 }
// NewZoneproductController creates and registers handles for the zoneproduct controller
func NewZoneproductController(router gin.IRouter, client *ent.Client) *ZoneproductController {
    pc := &ZoneproductController{
        client: client,
        router: router,
    }
    pc.register()
    return pc
 }
  
 // InitZoneproductController registers routes to the main engine
 func (ctl *ZoneproductController) register() {
    zoneproducts := ctl.router.Group("/zoneproducts")
  
    zoneproducts.GET("", ctl.ListZoneproduct)
  
    // CRUD
    zoneproducts.POST("", ctl.CreateZoneproduct)
    zoneproducts.GET(":id", ctl.GetZoneproduct)
    zoneproducts.PUT(":id", ctl.UpdateZoneproduct)
    zoneproducts.DELETE(":id", ctl.DeleteZoneproduct)
 }
      