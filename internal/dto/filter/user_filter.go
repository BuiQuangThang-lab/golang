package filter

type UserFilter struct {
	Name     string `form:"name"`
	Page     int    `form:"page"`
	PageSize int    `form:"pageSize"`
}
