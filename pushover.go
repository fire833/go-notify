/*
*	Copyright (C) 2023 Kendall Tauser
*
*	This program is free software; you can redistribute it and/or modify
*	it under the terms of the GNU General Public License as published by
*	the Free Software Foundation; either version 2 of the License, or
*	(at your option) any later version.
*
*	This program is distributed in the hope that it will be useful,
*	but WITHOUT ANY WARRANTY; without even the implied warranty of
*	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
*	GNU General Public License for more details.
*
*	You should have received a copy of the GNU General Public License along
*	with this program; if not, write to the Free Software Foundation, Inc.,
*	51 Franklin Street, Fifth Floor, Boston, MA 02110-1301 USA.
 */

package gonotify

import (
	"bytes"
	"errors"
	"net/http"
	"strconv"
)

const (
	pushoverEndpoint = "https://api.pushover.net/1/messages.json"
)

type PushoverNotifier struct {
	genericHTTPNotifier
}

type PushoverConfig struct {
	// The API key to validate the message.
	APIKey string
	// Can be either a group or user key, just identifies where pushover needs to send the message.
	Userkey string
}

func NewPushoverNotifier() *PushoverNotifier {
	return &PushoverNotifier{
		genericHTTPNotifier: genericHTTPNotifier{
			config: NewDefaultSlackConfig(),
			closed: false,
		},
	}
}

func NewPushoverNotifierConfig(config *PushoverConfig) (*PushoverNotifier, error) {
	n := NewPushoverNotifier()
	return n, n.Configure(config)
}

func NewPushoverNotifierConfigMust(config *PushoverConfig) *PushoverNotifier {
	n := NewPushoverNotifier()
	if e := n.Configure(config); e != nil {
		panic(e)
	}

	return n
}

func NewDefaultPushoverConfig() *PushoverConfig {
	return &PushoverConfig{}
}

func (p *PushoverNotifier) SendMessage(msg *Message) error {
	return p.sendMessageInternal(msg, p.generateRequest, p.parseResponse, p.validateMessage)
}

func (p *PushoverNotifier) validateMessage(msg *Message) error {
	if msg.GetMessage() == "" {
		return errors.New("pushover: message body is required for notification")
	}

	return nil
}

func (p *PushoverNotifier) generateRequest(msg *Message) (*http.Request, error) {
	req, e := http.NewRequest("POST", pushoverEndpoint, bytes.NewReader([]byte{}))
	if e != nil {
		return nil, e
	}
	params := req.URL.Query()
	params.Add("token", p.config.GetData()["token"].(string))
	params.Add("user", p.config.GetData()["user"].(string))
	params.Add("message", msg.GetMessage())
	params.Add("title", msg.GetTitle())
	if msg.url != nil {
		params.Add("url", msg.url.String())
	}
	params.Add("timestamp", strconv.Itoa(int(msg.timestamp)))

	parsedQuery := params.Encode()
	req.URL.RawQuery = parsedQuery

	return req, nil
}

func (p *PushoverNotifier) parseResponse(resp *http.Response) error {
	if resp.StatusCode != 200 {
		return errors.New("pushover: unable to send message request successfully")
	}

	return nil
}

func (p *PushoverConfig) Validate() []error {
	return nil
}

func (p *PushoverConfig) GetData() map[string]interface{} {
	return map[string]interface{}{
		"token": p.APIKey,
		"user":  p.Userkey,
	}
}
