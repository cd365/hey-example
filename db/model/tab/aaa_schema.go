// code template version: v3.0.0 c3e763620528071cd91f9f9535dd9700e721d7a5 1743124166-20250328090926
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
