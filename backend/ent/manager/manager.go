// Code generated by entc, DO NOT EDIT.

package manager

const (
	// Label holds the string label denoting the manager type in the database.
	Label = "manager"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"

	// EdgeManagers holds the string denoting the managers edge name in mutations.
	EdgeManagers = "managers"

	// Table holds the table name of the manager in the database.
	Table = "managers"
	// ManagersTable is the table the holds the managers relation/edge.
	ManagersTable = "orderproducts"
	// ManagersInverseTable is the table name for the Orderproduct entity.
	// It exists in this package in order to avoid circular dependency with the "orderproduct" package.
	ManagersInverseTable = "orderproducts"
	// ManagersColumn is the table column denoting the managers relation/edge.
	ManagersColumn = "manager_managers"
)

// Columns holds all SQL columns for manager fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldEmail,
	FieldPassword,
}

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// EmailValidator is a validator for the "email" field. It is called by the builders before save.
	EmailValidator func(string) error
	// PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	PasswordValidator func(string) error
)
