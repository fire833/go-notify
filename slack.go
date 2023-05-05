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
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
)

type SlackNotifier struct {
	genericHTTPNotifier
}

type SlackConfig struct {
	WebhookURL string

	parsedUrl *url.URL
}

func NewSlackNotifier() *SlackNotifier {
	return &SlackNotifier{
		genericHTTPNotifier: genericHTTPNotifier{
			config: NewDefaultSlackConfig(),
			closed: false,
		},
	}
}

func NewSlackNotifierConfig(config *SlackConfig) (*SlackNotifier, error) {
	n := NewSlackNotifier()
	return n, n.Configure(config)
}

func NewSlackNotifierConfigMust(config *SlackConfig) *SlackNotifier {
	n := NewSlackNotifier()
	if e := n.Configure(config); e != nil {
		panic(e)
	}

	return n
}

func NewDefaultSlackConfig() *SlackConfig {
	return &SlackConfig{
		WebhookURL: "www.example.com",
	}
}

func (s *SlackNotifier) SendMessage(msg *Message) error {
	return s.sendMessageInternal(msg, s.generateRequest, s.parseResponse, s.validateMessage)
}

func (s *SlackNotifier) validateMessage(msg *Message) error {
	return nil
}

func (s *SlackNotifier) generateRequest(msg *Message) (*http.Request, error) {

	body := &map[string]interface{}{
		"text": msg.GetMessage(),
	}

	bdata, e := json.Marshal(body)
	if e != nil {
		return nil, e
	}

	return http.NewRequest("POST", s.config.GetData()["url"].(string), bytes.NewReader(bdata))
}

func (s *SlackNotifier) parseResponse(resp *http.Response) error {
	if resp.StatusCode != 200 {
		return errors.New("slack: unable to send message request successfully")
	}

	return nil
}

func (c *SlackConfig) Validate() []error {
	if url, e := url.Parse(c.WebhookURL); e != nil {
		return []error{e}
	} else {
		c.parsedUrl = url
		return nil
	}
}

func (c *SlackConfig) GetData() map[string]interface{} {
	return map[string]interface{}{
		"url":    c.WebhookURL,
		"parsed": c.parsedUrl,
	}
}
