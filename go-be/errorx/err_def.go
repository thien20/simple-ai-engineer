package errorx

var (
	BadRequest          = NewError(400, "Bad Request", nil)
	Unauthorized        = NewError(401, "Unauthorized", nil)
	Forbidden           = NewError(403, "Forbidden", nil)
	NotFound            = NewError(404, "Not Found", nil)
	Conflict            = NewError(409, "Conflict", nil)
	InternalServerError = NewError(500, "Internal Server Error", nil)
	BadGateway          = NewError(502, "Bad Gateway", nil)
)
