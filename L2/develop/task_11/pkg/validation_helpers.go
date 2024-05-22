package pkg

import (
    "net/url"
    "strconv"
    "time"
	"errors"
)

func ParseInt(values url.Values, key string) (int, error) {
    valStr := values.Get(key)
    if valStr == "" {
        return 0, errors.New("missing " + key)
    }
    return strconv.Atoi(valStr)
}

func ParseDate(values url.Values, key string) (time.Time, error) {
    valStr := values.Get(key)
    if valStr == "" {
        return time.Time{}, errors.New("missing " + key)
    }
    return time.Parse("2006-01-02", valStr)
}
