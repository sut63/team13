// Code generated by entc, DO NOT EDIT.

package assessment

const (
	// Label holds the string label denoting the assessment type in the database.
	Label = "assessment"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAssessment holds the string denoting the assessment field in the database.
	FieldAssessment = "assessment"

	// EdgeFormassessment holds the string denoting the formassessment edge name in mutations.
	EdgeFormassessment = "formassessment"

	// Table holds the table name of the assessment in the database.
	Table = "assessments"
	// FormassessmentTable is the table the holds the formassessment relation/edge.
	FormassessmentTable = "salaries"
	// FormassessmentInverseTable is the table name for the Salary entity.
	// It exists in this package in order to avoid circular dependency with the "salary" package.
	FormassessmentInverseTable = "salaries"
	// FormassessmentColumn is the table column denoting the formassessment relation/edge.
	FormassessmentColumn = "assessment_formassessment"
)

// Columns holds all SQL columns for assessment fields.
var Columns = []string{
	FieldID,
	FieldAssessment,
}

var (
	// AssessmentValidator is a validator for the "assessment" field. It is called by the builders before save.
	AssessmentValidator func(string) error
)
