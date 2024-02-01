package intx

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"log"
)

func sign(method, path string, t int64, signingKey string, body []byte) string {
	key, err := base64.StdEncoding.DecodeString(signingKey)
	if err != nil {
		log.Fatalf("Error decoding signing key: %v", err)
	}

	message := fmt.Sprintf("%d%s%s", t, method, path)
	if len(body) > 0 {
		message += string(body)
	}

	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func appendQueryParam(queryParams, key, value string) string {
	if queryParams == "" {
		return "?" + key + "=" + value
	} else {
		return queryParams + "&" + key + "=" + value
	}
}

func appendPaginationParams(v string, p *PaginationParams) string {
	if p == nil {
		return v
	}

	if p.RefDatetime != "" {
		v = appendQueryParam(v, "ref_datetime", p.RefDatetime)
	}

	if p.ResultLimit > 0 {
		v = appendQueryParam(v, "result_limit", fmt.Sprint(p.ResultLimit))
	}

	if p.ResultOffset > 0 {
		v = appendQueryParam(v, "result_offset", fmt.Sprint(p.ResultOffset))
	}

	return v
}
