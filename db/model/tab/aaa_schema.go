// code template version: v3.0.0 876382ccafbc7ec905331e01d9c66afa58a11d6b 1744869629-20250417140029
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
