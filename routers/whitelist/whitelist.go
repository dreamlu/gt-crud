package whitelist

import "strings"

// IP 和 PATH 白名单
var (
	WLIp   Whitelist = &IPWhitelist{}
	WLPath Whitelist = &PathWhitelist{}
)

func init() {
	WLIp.Add(
		"127.0.0.1",
		"localhost",
	)
	WLPath.Add(
		"/login",
		"/static/file",
		"/wx/notify",
	)
}

type Whitelist interface {
	Contains(string) bool
	Add(...string)
}

// PathWhitelist IP白名单
type PathWhitelist []string

func (w *PathWhitelist) Add(paths ...string) {
	*w = append(*w, paths...)
}

func (w PathWhitelist) Contains(path string) bool {
	for _, s := range w {
		if strings.Contains(s, path) || strings.Contains(path, s) {
			return true
		}
	}
	return false
}

// IPWhitelist IP白名单
type IPWhitelist map[string]string

func (w IPWhitelist) Add(ips ...string) {
	for _, ip := range ips {
		w[ip] = ip
	}
}

func (w IPWhitelist) Contains(ip string) bool {
	if strings.Contains(ip, ":") {
		ip = strings.Split(ip, ":")[0]
	}
	return w[ip] == ip
}
