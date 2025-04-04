// code template version: v3.0.0 c3e763620528071cd91f9f9535dd9700e721d7a5 1743124166-20250328090926
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
