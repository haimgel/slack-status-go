package cmd

import (
	"fmt"
	slack2 "github.com/haimgel/slack-status/internal/slack"
	"net/http"
	"time"
)

func Execute(version string, exit func(int), args []string) {
	//cookie := slack.GetDCookie()
	cookie := &http.Cookie{
		Name:     "d",
		Value:    "xoxd-9%2FFyocxzZ2E17ACCYqcRfCu0x2Er48eXT%2FBtF7IeSNRERAxJLkZ5fFkHl5EZZCsGYvkUd09wdYRwJiwGjYvlPa8QD0mEeBRg3JtRLEcujPfN2ifMB0TGJoJSpTNUxeM1wcO3V5duoC0RltMmVTEhci9WzPH8aChg8dis4EZ5A97QXysEPWNsdhrQ",
		Path:     "/",
		Domain:   ".slack.com",
		Expires:  time.Now().AddDate(1, 0, 0),
		Secure:   true,
		HttpOnly: true,
		SameSite: 0,
	}
	fmt.Printf("d-cookie=%s", cookie.Value)
	result, err := slack2.ScrapeToken("hageltech", cookie)
	fmt.Printf("%s, %s", result, err)
	err = slack2.SetStatus(result, cookie.Value, "Lunch", ":hamburger:", 30*time.Minute)
	// Not authorized is `slack.SlackErrorResponse` with Err = "not_authed"
	fmt.Printf("%s", err)
	// slack.New("xoxc", slack.OptionCookie("d", ""))
}
