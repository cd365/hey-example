// code template version: v3.0.0 8232b1ce979cdaf7365eb708a30dddfa0cbaa290 1742361115-20250319131155
// You can add your own business logic code in this file.

package tab

import (
	"github.com/cd365/hey-example/db/model/biz"
)

type Custom struct {
	schema *biz.Schema
}

func NewCustom(
	schema *biz.Schema,
) *Custom {
	return &Custom{
		schema: schema,
	}
}
