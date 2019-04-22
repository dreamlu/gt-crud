// @author  dreamlu
package xss

import "html"

//type Xss struct {
//
//}
func XssMap(args map[string][]string) {
	for _, v := range args {
		v[0] = html.EscapeString(v[0])
	}
	return
}
