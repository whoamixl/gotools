package HttpUtil

type HttpState int

const (
	OK                  HttpState = 200
	BadRequest          HttpState = 400
	Unauthorized        HttpState = 401
	Forbidden           HttpState = 403
	NotFound            HttpState = 404
	InternalServerError HttpState = 500
)
