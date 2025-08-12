package router

const (
	GET    Method = "GET"
	POST   Method = "POST"
	PUT    Method = "PUT"
	PATCH  Method = "PATCH"
	DELETE Method = "DELETE"
)

const (
	FormatMethodNotAllowed = "could not find any route matching %s %s"
)

// Methods vector of available HTTP MEthods
var MethodsList = [5]Method{GET, PUT, POST, PATCH, DELETE}
