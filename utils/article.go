package utils

import "github.com/microcosm-cc/bluemonday"


//AvoidXSS 避免XSS
func AvoidXSS(theHTML string) string{
	return bluemonday.UGCPolicy().Sanitize(theHTML)
}
