// code template version: v3.0.0 67ab087b6ba2926de886c8a05e3188b18cd6567d 1745553000-20250425115000
// TEMPLATE CODE DO NOT EDIT IT.

package tab

import (
	"github.com/cd365/hey-example/db/model/biz"
)

type Schema struct {
	Company  *Company  // company
	Employee *Employee // employee
}

func NewSchema(schema *biz.Schema) *Schema {
	return &Schema{
		Company:  NewCompany(schema.Company),   // company
		Employee: NewEmployee(schema.Employee), // employee
	}
}
