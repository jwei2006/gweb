package sys

import "time"

type Context struct {
	PrincipalCode string
	TenantCode string
	CompanyId string
	UserId string
	UserName string
	RequestTime time.Time
	RequestIp string
	RequestId string
	IsManager int
	Auth int
	Business int
	CompanyType int
}