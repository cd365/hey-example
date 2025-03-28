// code template version: v3.0.0 c3e763620528071cd91f9f9535dd9700e721d7a5 1743124166-20250328090926
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
