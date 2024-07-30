package model

var (
	LogLevel        = "log-level"
	LogLevelInfo    = "info"
	LogLevelError   = "error"
	LegLevelDebug   = "debug"
	LogLevelWarning = "warn"
)

var (
	ApiPackage        = "api"
	StorePackage      = "store"
	ControllerPackage = "controller"
	ModelPackage      = "model"
	UtilPackage       = "util"
	MainPackage       = "main"
)

var (
	Controller = "controller"
	Store      = "store"
	Api        = "api"
	Main       = "main"
)

var (
	NewServer = "new-server"
	NewStore  = "new-store"

	CreateUser      = "create-user"
	GetUser         = "get-user"
	GetUsers        = "get-users"
	GetUserByFilter = "get-user-by-filter"
	UpdateUser      = "update-user"
	DeleteUser      = "delete-user"
)

// General
var (
	Value    = "value"
	Email    = "email"
	Password = "password"
	UserID   = "userID"
	Expire   = "exp"

	Authorization = "X-Token"

	DSN = "host=localhost user=iot password=iot dbname=homeautomation port=5432 sslmode=disable"

	DataPerPage = "limit"
	PageNumber  = "page"
	StartDate   = "start_date"
	EndDate     = "end_date"
	TimeLayout  = "2006-01-02 15:04:05.000 -0700"
)
