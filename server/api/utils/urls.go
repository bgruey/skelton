package api

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func IdFromUrl(r *http.Request) (int, error) {
	parts := strings.Split(r.URL.String(), "/")
	if len(parts) != 3 {
		return 0, errors.New("not found")
	}
	id, err := strconv.Atoi(parts[len(parts)-1])
	if err != nil {
		return 0, errors.New("not found")
	}
	return id, nil
}

func UuidFromUrl(prefix string, r *http.Request) (string, error) {
	parts := strings.Split(r.URL.String(), prefix)
	if len(parts) < 2 {
		return "", fmt.Errorf("unparsable url for %s", prefix)
	}
	// Remove query params
	target := strings.Split(parts[1], "?")[0]
	ret := strings.Replace(target, "/", "", -1)

	return ret, nil
}
