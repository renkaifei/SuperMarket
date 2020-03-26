package wxapi

import (
	"crypto/sha1"
	"fmt"
	"sort"
	"strings"
)

func ValidateSignature(signature string, timestamp string, nonce string) bool {
	tempArr := []string{token, timestamp, nonce}
	sort.Strings(tempArr)
	temp := strings.Join(tempArr, "")
	result := fmt.Sprintf("%x", sha1.Sum([]byte(temp)))
	if result == signature {
		return true
	} else {
		return false
	}
}
