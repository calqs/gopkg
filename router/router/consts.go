package router

const (
	GET    Method = "GET"
	POST   Method = "POST"
	PUT    Method = "PUT"
	PATCH  Method = "PATCH"
	DELETE Method = "DELETE"
)

// Methods vector of available HTTP MEthods
var MethodsList = [5]Method{GET, PUT, POST, PATCH, DELETE}
