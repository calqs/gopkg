package router

import "strings"

func CleanPath(path string) string {
	path = strings.ReplaceAll(path, "//", "/")
	path = strings.TrimRight(path, "/")
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	return path
}
