package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"github.com/project-reclass/gopr/config"
)

type SlackResponse struct {
	OK bool `json:"ok"`
}

func main() {
	slackConfig := config.New("config.yml")

	if err := sendSlackMessage(os.Args[1], slackConfig.Slack); err != nil {
		panic(err)
	}
}

func sendSlackMessage(text string, config config.SlackConfig) error {
	slackURL := url.URL{
		RawQuery: fmt.Sprintf("token=%s&channel=%s&text=%s&pretty=%s&icon_emoji=%s", config.Token,
			config.Channel, url.QueryEscape(text), "1", config.IconEmoji),
	}

	resp, err := http.Get(config.URI + slackURL.String())
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	slackResponse := &SlackResponse{}
	if err = json.Unmarshal(body, slackResponse); err != nil {
		return err
	}

	if !slackResponse.OK {
		return errors.New("Response was not ok")
	}
	return nil
}
