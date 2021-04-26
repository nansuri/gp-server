package util

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	model "github.com/nansuri/gp-server/model"
)

func SendNotification(token string, ticketDetail model.JiraRequest, key string) {
	url := "https://oapi.dingtalk.com/robot/send?access_token=" + token
	method := "POST"

	rawPayload := `{
		"msgtype": "text",
		"text": {
			"content": "Hi member on duty, you have a new issue posted by ` + ticketDetail.Reporter + `\n- Ticket ID : ` + key + `\n- Title : ` + ticketDetail.Summary + `\n- Link : https://danaindonesia.atlassian.net/browse/` + key + ` "
		},
		"at": {
			"atMobiles": [
				"6285224056939"
			],
			"isAtAll": ` + ticketDetail.IsUrgent + `
		}
	}`

	payload := strings.NewReader(rawPayload)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("\n" + string(body))
}
