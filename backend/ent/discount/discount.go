// Code generated by entc, DO NOT EDIT.

package discount

const (
	// Label holds the string label denoting the discount type in the database.
	Label = "discount"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldSale holds the string denoting the sale field in the database.
	FieldSale = "sale"

	// EdgeFordiscount holds the string denoting the fordiscount edge name in mutations.
	EdgeFordiscount = "fordiscount"

	// Table holds the table name of the discount in the database.
	Table = "discounts"
	// FordiscountTable is the table the holds the fordiscount relation/edge.
	FordiscountTable = "promotions"
	// FordiscountInverseTable is the table name for the Promotion entity.
	// It exists in this package in order to avoid circular dependency with the "promotion" package.
	FordiscountInverseTable = "promotions"
	// FordiscountColumn is the table column denoting the fordiscount relation/edge.
	FordiscountColumn = "discount_fordiscount"
)

// Columns holds all SQL columns for discount fields.
var Columns = []string{
	FieldID,
	FieldSale,
}