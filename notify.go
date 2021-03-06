package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"
)

type notifier struct{}

func (n *notifier) NotifyConnectionStatus() error {

	values := map[string]string{"mobile": strings.Join(config.Notify.Mobiles, ","), "content": config.Notify.Content}
	jsonValue, _ := json.Marshal(values)
	_, err := http.Post(config.Notify.Url, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		return err
	}
	return nil
}
