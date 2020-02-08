package get

import (
	"fmt"
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
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println("Redirect:", req)
			return nil
		},
	}
	return client.Do(request)
}
