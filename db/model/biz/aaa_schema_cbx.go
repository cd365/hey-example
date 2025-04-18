// code template version: v3.0.0 876382ccafbc7ec905331e01d9c66afa58a11d6b 1744869629-20250417140029
// TEMPLATE CODE DO NOT EDIT IT.

package biz

import (
	"github.com/cd365/hey-example/db/model"
)

type Schema struct {
	Company  Company  // company
	Employee Employee // employee
}

func NewSchema(db *model.Database, initialize func(db *model.Database, schema *Schema) error) (*Schema, error) {
	tmp := &Schema{
		Company:  NewCompany(db.COMPANY),   // company
		Employee: NewEmployee(db.EMPLOYEE), // employee
	}
	if initialize != nil {
		if err := initialize(db, tmp); err != nil {
			return nil, err
		}
	}
	return tmp, nil
}
