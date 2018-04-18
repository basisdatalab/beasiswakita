package errors

var (
	BadRequest          = New(400, "Bad Request", "")
	Unauthorized        = New(401, "Unauthorized", "")
	Forbidden           = New(403, "Forbidden", "")
	NotFound            = New(404, "Not Found", "")
	UnprocessableEntity = New(422, "Unprocessable Entity", "")
	InternalServerError = New(500, "Internal Server Error", "")
)
