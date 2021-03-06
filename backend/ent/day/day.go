// Code generated by entc, DO NOT EDIT.

package day

const (
	// Label holds the string label denoting the day type in the database.
	Label = "day"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldDay holds the string denoting the day field in the database.
	FieldDay = "day"

	// EdgeWhatday holds the string denoting the whatday edge name in mutations.
	EdgeWhatday = "whatday"

	// Table holds the table name of the day in the database.
	Table = "days"
	// WhatdayTable is the table the holds the whatday relation/edge.
	WhatdayTable = "employee_working_hours"
	// WhatdayInverseTable is the table name for the EmployeeWorkingHours entity.
	// It exists in this package in order to avoid circular dependency with the "employeeworkinghours" package.
	WhatdayInverseTable = "employee_working_hours"
	// WhatdayColumn is the table column denoting the whatday relation/edge.
	WhatdayColumn = "day_whatday"
)

// Columns holds all SQL columns for day fields.
var Columns = []string{
	FieldID,
	FieldDay,
}
