package http

import "slices"

var counter = 0

type Route struct {
	Method     string
	Hook       bool
	Path       string
	PathParams map[string]int
	Order      int
}

func NewRoute(method string, path string) *Route {
	counter++
	return &Route{Method: method, Hook: true, Path: path, Order: 100}
}

func GetCounter() int {
	return counter
}

func (r *Route) Alias(path string) *Route {
	addRouteAlias(path, *r)
	return r
}

func (r *Route) SetPathParam(key string, value int) {
	r.PathParams[key] = value
}

func (r *Route) GetValuePath(request Request) *Route {
	pathValues := GetCleanParts(request.GetURI())

	for key, index := range r.PathParams {
		if slices.Contains(pathValues, key) {
			pathValues
		}
	}

}
