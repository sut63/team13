package controllers

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/team13/app/ent"
	"github.com/team13/app/ent/assessment"
	"github.com/team13/app/ent/employee"
	"github.com/team13/app/ent/position"
)

// SalaryController defines the struct for the salary controller
type SalaryController struct {
	client *ent.Client
	router gin.IRouter
}

// Salary defines the struct for the salary
type Salary struct {
	EmployeeID     int
	PositionID     int
	AssessmentID   int
	salarys		   float32
	SalaryDate     string

}

// CreateSalary handles POST requests for adding salary entities
// @Summary Create salary
// @Description Create salary
// @ID create-salary
// @Accept   json
// @Produce  json
// @Param salary body Salary true "Salary entity"
// @Success 200 {object} ent.Salary
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /salarys [post]
func (ctl *SalaryController) CreateSalary(c *gin.Context) {
	obj := Salary{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "salary binding failed",
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

	po, err := ctl.client.Position.
		Query().
		Where(position.IDEQ(int(obj.PositionID))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "position diagnostic  not found",
		})
		return
	}

	ass, err := ctl.client.Assessment.
		Query().
		Where(assessment.IDEQ(int(obj.AssessmentID))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "assessment not found",
		})
		return
	}

	salaryDatetime, err := time.Parse(time.RFC3339, obj.SalaryDate)
	

	se, err := ctl.client.Salary.
		Create().
		SetEmployee(em).
		SetPosition(po).
		SetAssessment(ass).
		SetSalaryDatetime(salaryDatetime).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, se)
}

// ListSalary handles request to get a list of salary entities
// @Summary List salary entities
// @Description list salary entities
// @ID list-salary
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Salary
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /salarys [get]
func (ctl *SalaryController) ListSalary(c *gin.Context) {
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

	salarys, err := ctl.client.Salary.
		Query().
		WithEmployee().
		WithAssessment().
		WithPosition().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, salarys)
}

// DeleteSalary handles DELETE requests to delete a salary entity
// @Summary Delete a salary entity by ID
// @Description get salary by ID
// @ID delete-salary
// @Produce  json
// @Param id path int true "Salary ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /salarys/{id} [delete]
func (ctl *SalaryController) DeleteSalary(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Salary.
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

// NewSalaryController creates and registers handles for the salary controller
func NewSalaryController(router gin.IRouter, client *ent.Client) *SalaryController {
	drc := &SalaryController{
		client: client,
		router: router,
	}
	drc.register()
	return drc
}

// InitUserController registers routes to the main engine
func (ctl *SalaryController) register() {
	salarys := ctl.router.Group("/salarys")

	salarys.GET("", ctl.ListSalary)

	// CRUD
	salarys.POST("", ctl.CreateSalary)
	salarys.DELETE(":id", ctl.DeleteSalary)
}