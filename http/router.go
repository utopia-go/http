package http

import (
	"fmt"
	"slices"
	"strings"
)

const PlaceholderToken = ":::"
const WildcardToken = "*"

var allowOverride = false

var routes = map[string]map[string]Route{
	MethodGet:    make(map[string]Route),
	MethodPost:   make(map[string]Route),
	MethodPut:    make(map[string]Route),
	MethodPatch:  make(map[string]Route),
	MethodDelete: make(map[string]Route),
}

var params = []int{}

func GetRoutes() map[string]map[string]Route {
	return routes
}

func GetAllowOverride() bool {
	return allowOverride
}

func SetAllowOverride(value bool) {
	allowOverride = value
}

func AddRoute(route Route) {
	path, localParams := preparePath(route.Path)
	if _, ok := routes[route.Method]; !ok {
		panic(fmt.Sprintf("Method %s not supported.", route.Method))
	}

	if _, ok := routes[route.Method][path]; ok {
		panic(fmt.Sprintf("Route for %s:%s already registered.", route.Method, path))
	}

	for key, index := range localParams {
		route.SetPathParam(key, index)
	}

	routes[route.Method][path] = route
}

func addRouteAlias(path string, route Route) {
	localPath, _ := preparePath(path)

	if _, ok := routes[route.Method][localPath]; ok {
		panic(fmt.Sprintf("Route for %s:%s already registered.", route.Method, path))
	}

	routes[route.Method][path] = route
}

func preparePath(path string) (string, map[string]int) {
	parts := GetCleanParts(path)

	prepare := ""
	localParams := map[string]int{}

	for i, part := range parts {
		if i != 0 {
			prepare += "/"
		}

		if strings.Index(part, ":") == 0 {
			prepare += PlaceholderToken
			localParams[strings.Trim(part, ":")] = i
			if !slices.Contains(params, i) {
				params = append(params, i)
			}
		} else {
			prepare += part
		}
	}

	return prepare, localParams
}

func GetCleanParts(path string) []string {
	parts := []string{}
	for _, str := range strings.Split(path, `/`) {
		if str == "" {
			continue
		}
		parts = append(parts, str)
	}
	return parts
}

func Reset() {
	params = []int{}
	routes = map[string]map[string]Route{
		MethodGet:    make(map[string]Route),
		MethodPost:   make(map[string]Route),
		MethodPut:    make(map[string]Route),
		MethodPatch:  make(map[string]Route),
		MethodDelete: make(map[string]Route),
	}
}
