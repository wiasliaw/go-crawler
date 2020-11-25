package worker

import (
	"crypto/md5"
	"fmt"
	"net/url"
	"regexp"
)

var (
	contentTypeRegex = regexp.MustCompile(`[a-z]+[\/]{1}[a-z+-]+`)
)

func queryFilter(absURL string) (string, error) {
	u, err := url.Parse(absURL)
	if err != nil {
		return "", err
	}
	u.RawQuery = ""
	return u.String(), nil
}

func contentTypeFilter(str string) string {
	match := contentTypeRegex.Find([]byte(str))
	return string(match)
}

func hashMD5(src string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(src)))
}
