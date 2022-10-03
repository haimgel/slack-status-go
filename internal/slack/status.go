package slack

import (
	slackApi "github.com/slack-go/slack"
	"time"
)

func GetStatus(token string, dCookie string) (string, error) {
	api := slackApi.New(token, slackApi.OptionCookie("d", dCookie))
	resp, err := api.GetUserProfile(&slackApi.GetUserProfileParameters{UserID: "", IncludeLabels: false})
	if err != nil {
		return "", err
	}
	return resp.StatusText, err
}

func SetStatus(token string, dCookie string, statusText string, statusEmoji string, statusExpiration time.Duration) error {
	api := slackApi.New(token, slackApi.OptionCookie("d", dCookie))
	return api.SetUserCustomStatus(statusText, statusEmoji, time.Now().Add(statusExpiration).Unix())
}
