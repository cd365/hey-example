// code template version: v3.0.0 6e51d011dc279801cc620f872d835f27cb05e3af 1746444860-20250505193420
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
