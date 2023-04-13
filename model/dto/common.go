package dto

type PageInfo struct {
	PageNum  int `uri:"pageNum" json:"pageNum" form:"pageNum" binding:"required,gt=0" errorMessage:"分页页码错误！" `
	PageSize int `uri:"pageSize" json:"pageSize" form:"pageSize" binding:"required,gt=0" errorMessage:"每页数量错误！" `
}

type PageResult[T any] struct {
	*PageInfo
	TotalCount int `json:"totalCount"`
	TotalPage  int `json:"totalPage"`
	Data       T   `json:"data"`
}

func (d *PageResult[T]) SetTheTotalNumberOfPages() {
	totalPage := d.TotalCount / d.PageSize
	remainder := d.TotalCount % d.PageSize
	if remainder > 0 {
		remainder = 1
	} else {
		remainder = 0
	}
	d.TotalPage = totalPage + remainder
}

type Id struct {
	Id string `form:"id" uri:"id" binding:"required" errorMessage:"请求错误参数错误！" `
}
