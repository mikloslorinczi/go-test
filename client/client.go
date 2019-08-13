package client

import (
	"github.com/ddliu/go-httpclient"
	"github.com/pkg/errors"
)

func Call(url string) (string, error) {
	httpclient.Defaults(httpclient.Map{
		httpclient.OPT_USERAGENT: "my awsome httpclient",
		"Accept-Language":        "en-us",
	})

	res, err := httpclient.Get(url)
	if err != nil {
		return "", errors.Wrapf(err, "Failed to get %s", url)
	}

	defer res.Body.Close()

	if resStr, err := res.ToString(); err == nil {
		return resStr, nil
	}

	return "", errors.New("An error has occured...")
}
