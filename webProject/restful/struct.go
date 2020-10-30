package restful

var (
	UnknownErr = -1
	NotFound   = NewApiCode(404, "HandleNotFound")
	ServerErr  = 500
	DefaultOptions = &Options{
		MaxAge: 		  1800,
		AllowCredentials: true,
	}
)

const SUCCESS = 0

type ApiCode struct {
	Code int
	Msg string
}

func (e *ApiCode) Error() string {
	return e.Msg
}

type Result struct {
	Errno int			`json:"errno"`
	Data  interface{}	`json:"data,omitempty"`
	Error string		`json:"error,omitempty"`
}

type Options struct {
	MaxAge 				int
	AllowCredentials	bool
}