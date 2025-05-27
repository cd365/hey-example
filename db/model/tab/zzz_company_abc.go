// code template version: v3.0.0 a1e877e692cab7668466ba74010a8e88e78e039e 1748326418-20250527141338
// You can add your own business logic code in this file.

package tab

import (
	"github.com/cd365/hey-example/db/model/biz"
)

type Company struct {
	biz.Company
}

func NewCompany(company biz.Company) *Company {
	return &Company{
		Company: company,
	}
}

/* Your custom method. */
