// Ref: https://golangcode.com/validate-an-email-address/

package mailmgr

import (
	"net"
	"regexp"
	"strings"
)

// isEmailValid checks if the email provided passes the required structure
// and length test. It also checks the domain has a valid MX record.

func IsEmailValid(exp string) bool {

	var isValid bool = true
	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	if len(exp) < 3 || len(exp) > 254 {
		isValid = false
	} else if !emailRegex.MatchString(exp) {
		isValid = false
	} else {
		parts := strings.Split(exp, "@")
		mx, err := net.LookupMX(parts[1])
		if err != nil || len(mx) == 0 {
			isValid = false
		}
	}
	return isValid
}
