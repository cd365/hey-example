// code template version: v3.0.0 8232b1ce979cdaf7365eb708a30dddfa0cbaa290 1742361115-20250319131155
// TEMPLATE CODE DO NOT EDIT IT.

package tab

import (
	"github.com/cd365/hey-example/db/model/biz"
)

type Schema struct {
	Company  *Company  // company
	Employee *Employee // employee
}

func NewSchema(schema *biz.Schema, custom *Custom) *Schema {
	return &Schema{
		Company:  NewCompany(schema.Company, custom),   // company
		Employee: NewEmployee(schema.Employee, custom), // employee
	}
}
