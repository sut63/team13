package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/team13/app/ent"
	"github.com/team13/app/ent/day"
	"github.com/team13/app/ent/employee"
	"github.com/team13/app/ent/employeeworkinghours"
	"github.com/team13/app/ent/role"
	"github.com/team13/app/ent/startwork"
	"github.com/team13/app/ent/endwork"
	"github.com/gin-gonic/gin"
)

// EmployeeWorkingHoursController defines the struct for the employeeWorkingHours
type EmployeeWorkingHours struct {
	CodeWork	string
	IDEmployee	string
	IDNumber	string
	Employee 	int
	Day     	int
	Role	  	int
	StartWork	int
	EndWork		int
	Wages		float64
}

// EmployeeWorkingHoursController defines the struct for the employeeworkinghours controller
type EmployeeWorkingHoursController struct {
   client *ent.Client
   router gin.IRouter
}

// CreateEmployeeWorkingHours handles POST requests for adding employeeworkinghours entities
// @Summary Create employeeworkinghours
// @Description Create employeeworkinghours
// @ID create-employeeworkinghours
// @Accept   json
// @Produce  json
// @Param employeeworkinghours body EmployeeWorkingHours true "EmployeeWorkingHours entity"
// @Success 200 {object} ent.EmployeeWorkingHours
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /employeeworkinghourss [post]
func (ctl *EmployeeWorkingHoursController) CreateEmployeeWorkingHours(c *gin.Context) {
	obj := EmployeeWorkingHours{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "employeeworkinghours binding failed",
		})
		return
	}

	em, err := ctl.client.Employee.
		Query().
		Where(employee.IDEQ(int(obj.Employee))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Employee not found",
		})
		return
	}

	d, err := ctl.client.Day.
		Query().
		Where(day.IDEQ(int(obj.Day))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Day not found",
		})
		return
	}

	r, err := ctl.client.Role.
		Query().
		Where(role.IDEQ(int(obj.Role))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Role not found",
		})
		return
	}

	sw, err := ctl.client.StartWork.
		Query().
		Where(startwork.IDEQ(int(obj.StartWork))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "StartWork not found",
		})
		return
	}

	ew, err := ctl.client.EndWork.
		Query().
		Where(endwork.IDEQ(int(obj.EndWork))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "EndWork not found",
		})
		return
	}

	e, err := ctl.client.EmployeeWorkingHours.
		Create().
		SetCodeWork(obj.CodeWork).
		SetIDNumber(obj.IDNumber).
		SetWages(obj.Wages).
		SetEmployee(em).
		SetDay(d).
		SetRole(r).
		SetStartwork(sw).
		SetEndwork(ew).
		Save(context.Background())

	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"status" : false,
			"error"  : err,
		})
		return
	}

	c.JSON(200, gin.H{
		"status": true,
		"data"  :  e,
	})
 }
 
 // GetEmployeeWorkingHours handles GET requests to retrieve a employeeworkinghours entity
// @Summary Get a employeeworkinghours entity by ID
// @Description get employeeworkinghours by ID
// @ID get-employeeworkinghours
// @Produce  json
// @Param id path int true "EmployeeWorkingHours ID"
// @Success 200 {array} ent.EmployeeWorkingHours
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /employeeworkinghourss/{id} [get]
func (ctl *EmployeeWorkingHoursController) GetEmployeeWorkingHours(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	u, err := ctl.client.EmployeeWorkingHours.
		Query().
		WithDay().
		WithEmployee().
		WithRole().
		WithStartwork().
		WithEndwork().
		Where(employeeworkinghours.HasEmployeeWith(employee.IDEQ(int(id)))).
		All(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	c.JSON(200, u)
 }

 // ListEmployeeWorkingHours handles request to get a list of employeeworkinghours entities
// @Summary List employeeworkinghours entities
// @Description list employeeworkinghours entities
// @ID list-employeeworkinghours
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.EmployeeWorkingHours
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /employeeworkinghourss [get]
func (ctl *EmployeeWorkingHoursController) ListEmployeeWorkingHours(c *gin.Context) {
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
  
	employeeworkinghourss, err := ctl.client.EmployeeWorkingHours.
		Query().
		Limit(limit).
		Offset(offset).
		WithDay().
		WithEmployee().
		WithRole().
		WithStartwork().
		WithEndwork().
		All(context.Background())
		if err != nil {
		c.JSON(400, gin.H{"error": err.Error(),})
		return
	}
  
	c.JSON(200, employeeworkinghourss)
 }

 // DeleteEmployeeWorkingHours handles DELETE requests to delete a employeeworkinghours entity
// @Summary Delete a employeeworkinghours entity by ID
// @Description get employeeworkinghours by ID
// @ID delete-employeeworkinghours
// @Produce  json
// @Param id path int true "EmployeeWorkingHours ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /employeeworkinghourss/{id} [delete]
func (ctl *EmployeeWorkingHoursController) DeleteEmployeeWorkingHours(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	err = ctl.client.EmployeeWorkingHours.
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

 // UpdateEmployeeWorkingHours handles PUT requests to update a employeeworkinghours entity
// @Summary Update a employeeworkinghours entity by ID
// @Description update employeeworkinghours by ID
// @ID update-employeeworkinghours
// @Accept   json
// @Produce  json
// @Param id path int true "EmployeeWorkingHours ID"
// @Param employeeworkinghours body ent.EmployeeWorkingHours true "EmployeeWorkingHours entity"
// @Success 200 {object} ent.EmployeeWorkingHours
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /employeeworkinghourss/{id} [put]
func (ctl *EmployeeWorkingHoursController) UpdateEmployeeWorkingHours(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	obj := ent.EmployeeWorkingHours{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "employeeworkinghours binding failed",
		})
		return
	}
	obj.ID = int(id)
	u, err := ctl.client.EmployeeWorkingHours.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed",})
		return
	}
  
	c.JSON(200, u)
 }

 // NewEmployeeWorkingHoursController creates and registers handles for the employeeworkinghours controller
func NewEmployeeWorkingHoursController(router gin.IRouter, client *ent.Client) *EmployeeWorkingHoursController {
	uc := &EmployeeWorkingHoursController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
 }
  
 // InitEmployeeWorkingHoursController registers routes to the main engine
 func (ctl *EmployeeWorkingHoursController) register() {
	employeeworkinghourss := ctl.router.Group("/employeeworkinghourss")
  
	employeeworkinghourss.GET("", ctl.ListEmployeeWorkingHours)
  
	// CRUD
	employeeworkinghourss.POST("", ctl.CreateEmployeeWorkingHours)
	employeeworkinghourss.GET(":id", ctl.GetEmployeeWorkingHours)
	employeeworkinghourss.PUT(":id", ctl.UpdateEmployeeWorkingHours)
	employeeworkinghourss.DELETE(":id", ctl.DeleteEmployeeWorkingHours)
 }
 