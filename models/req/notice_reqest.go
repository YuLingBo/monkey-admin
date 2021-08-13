package req

import "monkey-admin/pkg/base"

type NoticeQuery struct {
	base.GlobalQuery
	NoticeTitle string `form:"noticeTitle"`
	NoticeType  string `form:"noticeType"`
	CreateBy    string `form:"createBy"`
}
