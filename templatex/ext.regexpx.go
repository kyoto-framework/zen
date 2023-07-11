package templatex

import (
	"html/template"

	"github.com/kyoto-framework/zen/v3/regexpx"
)

/*
RegexpxExtension is a template compatibility wrapper
around existing regexpx zen package.
*/
type RegexpxExtension struct{}

func (e *RegexpxExtension) Replace(str, old, new string) string {
	return regexpx.Replace(str, old, new)
}

func (e *RegexpxExtension) FuncMap() template.FuncMap {
	return template.FuncMap{
		"replacergx": e.Replace,
	}
}