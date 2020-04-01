package wxapi

import (
	"crypto/sha1"
	"fmt"
	"sort"
	"strings"
)

func ValidateSignature(signature string, timestamp string, nonce string) bool {
	tempArr := []string{Token, timestamp, nonce}
	sort.Strings(tempArr)
	temp := strings.Join(tempArr, "")
	result := fmt.Sprintf("%x", sha1.Sum([]byte(temp)))
	if result == signature {
		return true
	} else {
		return false
	}
}

func CalculateSignature(noncestr string, jsapiTicket string, timestamp string, url string) string {
	temp := "jsapi_ticket=" + jsapiTicket + "&noncestr=" + noncestr + "&timestamp=" + timestamp + "&url=" + url
	result := fmt.Sprintf("%x", sha1.Sum([]byte(temp)))
	return result
}
