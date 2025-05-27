// code template version: v3.0.0 a1e877e692cab7668466ba74010a8e88e78e039e 1748326418-20250527141338
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
