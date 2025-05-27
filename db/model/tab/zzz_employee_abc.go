// code template version: v3.0.0 a1e877e692cab7668466ba74010a8e88e78e039e 1748326418-20250527141338
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
