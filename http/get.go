package get

import (
	"net/http"
)

func Do(url string, headers map[string]string) (*http.Response, error) {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	for key, val := range headers {
		request.Header.Add(key, val)
	}

	client := http.Client{
		//CheckRedirect: func(req *http.Request, via []*http.Request) error {
		//	log.Println("Redirect:", req)
		//	return nil
		//},
	}
	return client.Do(request)
}
