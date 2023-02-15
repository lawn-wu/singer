package domain

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

func Request(
	urlPrefix string,
	urlHasParamsNoSig string,
	sig string,
	systemParams SortedMaps,
	applicationParams SortedMaps,
	requestMethodType RequestMethodType,
) (body []byte, err error) {
	var u string
	if requestMethodType == REQUEST_METHOD_TYPE_POST {
		u = UrlApp.GenOnlyHasSysParamsAndSigUrl(urlPrefix, systemParams, sig)
		body, err = Post(u, applicationParams.UrlValues())
	} else {
		u = UrlApp.GenUrlForGetRequest(urlHasParamsNoSig, sig)
		body, err = Get(u)
	}

	return
}

func Get(url string) (body []byte, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	return
}

func Post(url string, data url.Values) (body []byte, err error) {
	resp, err := http.PostForm(url, data)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	return
}
