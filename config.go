package hive

import (
	"net/url"
	"strconv"
	"time"
)

type ConnectOptions struct {
	Host      string
	Timeout   time.Duration
	AuthMode  string
	Username  string
	Password  string
	BatchSize int64
}

func connectOptionsFromURL(u *url.URL) (ConnectOptions, error) {
	var opts ConnectOptions

	opts.Host = u.Host
	queryParams := u.Query()

	opts.AuthMode = queryParams.Get("auth")
	if batchSize, err := strconv.ParseInt(queryParams.Get("batch"), 10, 64); err != nil {
		return ConnectOptions{}, err
	} else {
		opts.BatchSize = batchSize
	}
	if timeoutSeconds, err := strconv.ParseInt(queryParams.Get("connect_timeout"), 10, 0); err != nil {
		return ConnectOptions{}, err
	} else {
		opts.Timeout = time.Duration(timeoutSeconds) * time.Second
	}

	switch opts.AuthMode {
	case "sasl":
		if name := u.User.Username(); name != "" {
			opts.Username = name
		}
		if password, ok := u.User.Password(); ok {
			opts.Password = password
		} else {
			return ConnectOptions{}, ErrNoPassword
		}
	}

	return opts, nil
}
