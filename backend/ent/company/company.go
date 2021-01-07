// Code generated by entc, DO NOT EDIT.

package company

const (
	// Label holds the string label denoting the company type in the database.
	Label = "company"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"

	// EdgeCompanys holds the string denoting the companys edge name in mutations.
	EdgeCompanys = "companys"

	// Table holds the table name of the company in the database.
	Table = "companies"
	// CompanysTable is the table the holds the companys relation/edge.
	CompanysTable = "orderproducts"
	// CompanysInverseTable is the table name for the Orderproduct entity.
	// It exists in this package in order to avoid circular dependency with the "orderproduct" package.
	CompanysInverseTable = "orderproducts"
	// CompanysColumn is the table column denoting the companys relation/edge.
	CompanysColumn = "company_companys"
)

// Columns holds all SQL columns for company fields.
var Columns = []string{
	FieldID,
	FieldName,
}