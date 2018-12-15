package config

type SysConfig struct {
	Port int
	IsDebug bool
	IsNeedOpen bool

	LogGinPath string
	LogInfoPath string
	LogErrorPath string
	LogLevel string

	ShowSql bool
	//LogSql bool
	Db bool
	DbType string
	DbUrl string
	MaxIdle int
	MaxActive int

	CookieId string
	SessionId string

	LoginPathWx string
	LoginPath string
}
const RequestTime string = "requestTime"
const Key_Ctx = "k_ctx"
const Key_RequestId = "X-REQUEST-ID"
const Key_Open = "k_open"
const Openid = "openid"
const Key_UserId = "k_userId"
const Key_UserName = "k_userName"

const Key_principal string = "k_principal"
const Value_principal string = "PINGX"
const Key_Tenant string = "k_tenant"
const Value_tenant string = "PX"
const Key_CompanyId string = "k_companyId"
const Key_Company_Type string = "k_companyType"
const Key_Business string = "k_business"

var Conf SysConfig = SysConfig{}