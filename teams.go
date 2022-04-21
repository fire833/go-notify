/*
*	Copyright (C) 2022  Kendall Tauser
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
	"sync"
)

var ()

type TeamsNotifier struct {
	sync.RWMutex
	genericHTTPNotifier
}

type TeamsConfig struct {
	WebhookURL string

	Color string
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
	return t.sendMessageInternal(msg, t.generateRequest, t.parseResponse)
}

// Internal method to generate the request for Teams incoming webhook from
// a message. Callers should have a read lock already on the TeamsNotifier struct.
func (t *TeamsNotifier) generateRequest(msg *Message) (*http.Request, error) {

	// generate the kv object.
	fields := map[string]interface{}{}
	for key, value := range msg.metadata {
		fields[key] = value
	}

	body := &map[string]interface{}{
		"@type":      "MessageCard",
		"@context":   "http://schema.org/extensions",
		"summary":    msg.subtitle,
		"title":      msg.title,
		"themeColor": "",

		"sections": []map[string]interface{}{
			{
				"text": msg.msg,
			},
		},
	}

	bdata, e := json.Marshal(body)
	if e != nil {
		return nil, e
	}

	return http.NewRequest("POST", "", bytes.NewReader(bdata))

}

func (t *TeamsNotifier) parseResponse(*http.Response) error {
	return nil
}

func (c *TeamsConfig) Validate() []error {
	return nil
}

func (c *TeamsConfig) GetData() map[string]interface{} {
	return map[string]interface{}{
		"url":   c.WebhookURL,
		"color": c.Color,
	}
}
