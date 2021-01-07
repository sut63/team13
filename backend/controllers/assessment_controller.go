package controllers
 
import (
   "context"
   "fmt"
   "strconv"
   "github.com/team13/app/ent"
   "github.com/team13/app/ent/assessment"
   "github.com/gin-gonic/gin"
)
 
// AssessmentController defines the struct for the assessment controller
type AssessmentController struct {
   client *ent.Client
   router gin.IRouter
}

// CreateAssessment handles POST requests for adding assessment entities
// @Summary Create assessment
// @Description Create assessment
// @ID create-assessment
// @Accept   json
// @Produce  json
// @Param assessment body ent.Assessment true "Assessment entity"
// @Success 200 {object} ent.Assessment
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /assessments [post]
func (ctl *AssessmentController) CreateAssessment(c *gin.Context) {
	obj := ent.Assessment{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "assessment binding failed",
		})
		return
	}
  
	u, err := ctl.client.Assessment.
		Create().
		SetAssessment(obj.Assessment).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}
  
	c.JSON(200, u)
 }

 // GetAssessment handles GET requests to retrieve a assessment entity
// @Summary Get a assessment entity by ID
// @Description get assessment by ID
// @ID get-assessment
// @Produce  json
// @Param id path int true "Assessment ID"
// @Success 200 {object} ent.Assessment
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /assessments/{id} [get]
func (ctl *AssessmentController) GetAssessment(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	u, err := ctl.client.Assessment.
		Query().
		Where(assessment.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	c.JSON(200, u)
 }

 // ListAssessment handles request to get a list of assessment entities
// @Summary List assessment entities
// @Description list assessment entities
// @ID list-assessment
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Assessment
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /assessments [get]
func (ctl *AssessmentController) ListAssessment(c *gin.Context) {
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
  
	assessments, err := ctl.client.Assessment.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
		if err != nil {
		c.JSON(400, gin.H{"error": err.Error(),})
		return
	}
  
	c.JSON(200, assessments)
 }

 // DeleteAssessment handles DELETE requests to delete a assessment entity
// @Summary Delete a assessment entity by ID
// @Description get assessment by ID
// @ID delete-assessment
// @Produce  json
// @Param id path int true "Assessment ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /assessments/{id} [delete]
func (ctl *AssessmentController) DeleteAssessment(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	err = ctl.client.Assessment.
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

 // UpdateAssessment handles PUT requests to update a assessment entity
// @Summary Update a assessment entity by ID
// @Description update assessment by ID
// @ID update-assessment
// @Accept   json
// @Produce  json
// @Param id path int true "Assessment ID"
// @Param assessment body ent.Assessment true "Assessment entity"
// @Success 200 {object} ent.Assessment
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /assessments/{id} [put]
func (ctl *AssessmentController) UpdateAssessment(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	obj := ent.Assessment{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "assessment binding failed",
		})
		return
	}
	obj.ID = int(id)
	u, err := ctl.client.Assessment.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed",})
		return
	}
  
	c.JSON(200, u)
 }

 // NewAssessmentController creates and registers handles for the assessment controller
func NewAssessmentController(router gin.IRouter, client *ent.Client) *AssessmentController {
	uc := &AssessmentController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
 }
  
 // InitAssessmentController registers routes to the main engine
 func (ctl *AssessmentController) register() {
	assessments := ctl.router.Group("/assessments")
  
	assessments.GET("", ctl.ListAssessment)
  
	// CRUD
	assessments.POST("", ctl.CreateAssessment)
	assessments.GET(":id", ctl.GetAssessment)
	assessments.PUT(":id", ctl.UpdateAssessment)
	assessments.DELETE(":id", ctl.DeleteAssessment)
 }
 
