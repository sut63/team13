package controllers

import (
	"context"
	"fmt"
	"strconv"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/team13/app/ent"
	"github.com/team13/app/ent/employee"
	"github.com/team13/app/ent/product"
	"github.com/team13/app/ent/typeproduct"
	"github.com/team13/app/ent/zoneproduct"
	"github.com/team13/app/ent/stock"
	
)

// StockController defines the struct for the stock controller
type StockController struct {
	client *ent.Client
	router gin.IRouter
}

// Stock defines the struct for the stock
type Stock struct {
	ProductID	  int
	ZoneID        int
	EmployeeID    int
	TypeproductID int
	Priceproduct  float64
	Amount        int
	Time          string

}

// CreateStock handles POST requests for adding stock entities
// @Summary Create stock
// @Description Create stock
// @ID create-stock
// @Accept   json
// @Produce  json
// @Param stock body Stock true "Stock entity"
// @Success 200 {object} ent.Stock
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /stocks [post]
func (ctl *StockController) CreateStock(c *gin.Context) {
	obj := Stock{}
		if err := c.ShouldBind(&obj); err != nil {
			c.JSON(400, gin.H{
				"error": "stock binding failed",
			})
			return
		}
	
		em, err := ctl.client.Employee.
			Query().
			Where(employee.IDEQ(int(obj.EmployeeID))).
			Only(context.Background())
	
		if err != nil {
			c.JSON(400, gin.H{
				"error": "employee not found",
			})
			return
		}
	
		pr, err := ctl.client.Product.
			Query().
			Where(product.IDEQ(int(obj.ProductID))).
			Only(context.Background())
	
		if err != nil {
			c.JSON(400, gin.H{
				"error": "product diagnostic  not found",
			})
			return
		}
	
		ty, err := ctl.client.Typeproduct.
			Query().
			Where(typeproduct.IDEQ(int(obj.TypeproductID))).
			Only(context.Background())
	
		if err != nil {
			c.JSON(400, gin.H{
				"error": "typeproduct not found",
			})
			return
		}

		zo, err := ctl.client.Zoneproduct.
			Query().
			Where(zoneproduct.IDEQ(int(obj.ZoneID))).
			Only(context.Background())
	
		if err != nil {
			c.JSON(400, gin.H{
				"error": "zoneproduct not found",
			})
			return
		}
	
		times, err := time.Parse(time.RFC3339, obj.Time)
		
	

	s, err := ctl.client.Stock.
		Create().
		SetProduct(pr).
		SetZoneproduct(zo).
		SetTypeproduct(ty).
		SetEmployee(em).
		SetAmount(obj.Amount).
		SetPriceproduct(obj.Priceproduct).
		SetTime(times).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, s)
}

// GetStock handles GET requests to retrieve a stock entity
// @Summary Get a stock entity by ID
// @Description get stock by ID
// @ID get-stock
// @Produce  json
// @Param id path int true "Stock ID"
// @Success 200 {object} ent.Stock
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /stocks/{id} [get]
func (ctl *StockController) GetStock(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	pa, err := ctl.client.Stock.
		Query().
		Where(stock.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, pa)
 }
// ListStock handles request to get a list of stock entities
// @Summary List stock entities
// @Description list stock entities
// @ID list-stock
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Stock
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /stocks [get]
func (ctl *StockController) ListStock(c *gin.Context) {
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

	stocks, err := ctl.client.Stock.
		Query().
		WithProduct().
		WithZoneproduct().
		WithTypeproduct().
		WithEmployee().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, stocks)
}

// DeleteStock handles DELETE requests to delete a stock entity
// @Summary Delete a stock entity by ID
// @Description get stock by ID
// @ID delete-stock
// @Produce  json
// @Param id path int true "Stock ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /stocks/{id} [delete]
func (ctl *StockController) DeleteStock(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Stock.
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

// NewStockController creates and registers handles for the stock controller
func NewStockController(router gin.IRouter, client *ent.Client) *StockController {
	drc := &StockController{
		client: client,
		router: router,
	}
	drc.register()
	return drc
}

// InitUserController registers routes to the main engine
func (ctl *StockController) register() {
	stocks := ctl.router.Group("/stocks")

	stocks.GET("", ctl.ListStock)

	// CRUD
	stocks.POST("", ctl.CreateStock)
	stocks.DELETE(":id", ctl.DeleteStock)
}