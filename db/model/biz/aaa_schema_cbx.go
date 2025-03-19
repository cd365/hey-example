// code template version: v3.0.0 8232b1ce979cdaf7365eb708a30dddfa0cbaa290 1742361115-20250319131155
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
