package http

type Request struct {
	payload     interface{}
	queryString map[string]string
	headers     map[string]string
}

func (r Request) GetParam(key string) string {
	return ""
}
func (r Request) GetURI() string {
	return ""
}

const MethodOptions = "OPTIONS"
const MethodGet = "GET"
const MethodHead = "HEAD"
const MethodPost = "POST"
const MethodPatch = "PATCH"
const MethodPut = "PUT"
const MethodDelete = "DELETE"
const MethodTrace = "TRACE"
const MethodConnect = "CONNECT"
