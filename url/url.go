package url

import "strings"

func StripPort(host string) string {
	// strip the port from the host if it is present
	return strings.TrimRight(host, ":0123456789")
}
