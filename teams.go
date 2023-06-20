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
	"net/http"
	"net/url"
	"sync"
)

type TeamsNotifier struct {
	sync.RWMutex
	genericHTTPNotifier
}

type TeamsConfig struct {
	WebhookURL string `json:"teamsWebhookURL" yaml:"teamsWebhookURL"`
	Color      string `json:"teamsColor" yaml:"teamsColor"`

	parsedUrl *url.URL
}

func NewTeamsNotifier() *TeamsNotifier {
	return &TeamsNotifier{
		genericHTTPNotifier: genericHTTPNotifier{
			config: NewDefaultTeamsConfig(),
			closed: false,
		},
	}
}

func NewTeamsNotifierConfig(config *TeamsConfig) (*TeamsNotifier, error) {
	n := NewTeamsNotifier()
	return n, n.Configure(config)
}

func NewTeamsNotifierConfigMust(config *TeamsConfig) *TeamsNotifier {
	n := NewTeamsNotifier()
	if e := n.Configure(config); e != nil {
		panic(e)
	}

	return n
}

func NewDefaultTeamsConfig() *TeamsConfig {
	return &TeamsConfig{
		WebhookURL: "www.example.com",
		Color:      "0076D7",
	}
}

func (t *TeamsNotifier) SendMessage(msg *Message) error {
	return t.sendMessageInternal(msg, t.generateRequest, t.parseResponse, t.validateMessage)
}

func (t *TeamsNotifier) validateMessage(msg *Message) error {
	return nil
}

// Internal method to generate the request for Teams incoming webhook from
// a message. Callers should have a read lock already on the TeamsNotifier struct.
func (t *TeamsNotifier) generateRequest(msg *Message) (*http.Request, error) {
	body := &map[string]interface{}{
		"@type":      "MessageCard",
		"@context":   "http://schema.org/extensions",
		"summary":    msg.GetSubtitle(),
		"title":      msg.GetTitle(),
		"themeColor": t.config.GetData()["color"],

		"sections": []map[string]interface{}{
			{
				"text": msg.GetMessage(),
			},
		},
	}

	bdata, e := json.Marshal(body)
	if e != nil {
		return nil, e
	}

	return http.NewRequest("POST", t.config.GetData()["url"].(string), bytes.NewReader(bdata))
}

func (t *TeamsNotifier) parseResponse(resp *http.Response) error {
	return nil
}

func (c *TeamsConfig) Validate() []error {
	if url, e := url.Parse(c.WebhookURL); e != nil {
		return []error{e}
	} else {
		c.parsedUrl = url
		return nil
	}
}

func (c *TeamsConfig) GetData() map[string]interface{} {
	return map[string]interface{}{
		"url":   c.getWebhookURL(),
		"color": c.getColor(),
	}
}

func (c *TeamsConfig) getColor() string      { return c.Color }
func (c *TeamsConfig) getWebhookURL() string { return c.WebhookURL }
