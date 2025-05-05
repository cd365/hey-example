// code template version: v3.0.0 6e51d011dc279801cc620f872d835f27cb05e3af 1743124166-20250328090926
// You can add your own business logic code in this file.

package tab

import (
	"github.com/cd365/hey-example/db/model/biz"
)

type Employee struct {
	biz.Employee
}

func NewEmployee(employee biz.Employee) *Employee {
	return &Employee{
		Employee: employee,
	}
}

/* Your custom method. */
