package slack

import (
	"fmt"
	"github.com/browserutils/kooky"
	_ "github.com/browserutils/kooky/browser/all" // register cookie store finders!
	"io"
	"net/http"
	"regexp"
	"sort"
)

// Slack _needs_ user agent pretending to be Chrome, otherwise it does not give back a token
const fakeUserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36"

func GetDCookie() (*http.Cookie, error) {
	cookies := kooky.ReadCookies(kooky.Valid, kooky.DomainHasSuffix(`slack.com`), kooky.Name("d"))
	// Sort found cookies by expiry date, in reverse, to get the most "fresh" cookie
	sort.Slice(cookies, func(i, j int) bool {
		return cookies[i].Expires.After(cookies[j].Expires)
	})
	if len(cookies) > 0 {
		return &cookies[0].Cookie, nil
	} else {
		return nil, fmt.Errorf("slack's 'd' cookie wasn't found")
	}
}

func ScrapeToken(team string, dCookie *http.Cookie) (string, error) {
	// Just a page where token is exposed as part of JSON in the page body itself
	teamUrl := fmt.Sprintf("https://%s.slack.com/customize/emoji", team)

	req, err := http.NewRequest("GET", teamUrl, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("User-Agent", fakeUserAgent)
	req.AddCookie(dCookie)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	if (resp.StatusCode < 200) || (resp.StatusCode > 299) {
		return "", fmt.Errorf("unexpected HTTP response status code: %d", resp.StatusCode)
	}

	// Make sure the body is closed
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	html, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	// Find the API token in the page body
	re := regexp.MustCompile(`"api_token":"(xo[^"]+)"`)
	match := re.FindSubmatch(html)
	if match == nil {
		return "", fmt.Errorf("token is not found in the HTML body")
	}
	return string(match[1]), nil
}
