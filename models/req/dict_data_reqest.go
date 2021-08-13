package req

import "monkey-admin/pkg/base"

type DiceDataQuery struct {
	base.GlobalQuery
	DictType  string `form:"dictType"`
	DictLabel string `form:"dictLabel"`
	Status    string `form:"status"`
}
