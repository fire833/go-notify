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
)

type GotifyNotifier struct {
	genericHTTPNotifier
}

type GotifyConfig struct {
	Host   string `json:"gotifyHost" yaml:"gotifyHost"`
	APIKey string `json:"gotifyAPIKey" yaml:"gotifyAPIKey"`
}

func NewGotifyNotifier() *GotifyNotifier {
	return &GotifyNotifier{}
}

func NewGotifyNotifierConfig(config *GotifyConfig) (*GotifyNotifier, error) {
	n := NewGotifyNotifier()
	return n, n.Configure(config)
}

func NewGotifyNotifierConfigMust(config *GotifyConfig) *GotifyNotifier {
	n := NewGotifyNotifier()
	if e := n.Configure(config); e != nil {
		panic(e)
	}

	return n
}

func NewDefaultGotifyConfig() *GotifyConfig {
	return &GotifyConfig{}
}

func (g *GotifyNotifier) SendMessage(msg *Message) error {
	return g.sendMessageInternal(msg, g.generateRequest, g.parseResponse, g.validateMessage)
}

func (g *GotifyNotifier) validateMessage(msg *Message) error {
	if msg.GetMessage() == "" {
		return errors.New("gotify: message body is required for notification")
	}

	if msg.GetTitle() == "" {
		return errors.New("gotify: message title is required for notification")
	}

	if msg.GetPriority() < 0 {
		return errors.New("gotify: message priority is required for notification")
	}

	// if msg.GetMetadata() == nil {
	// 	return errors.New("gotify: message metadata must not be empty or must be an empty interface")
	// }

	return nil
}

func (g *GotifyNotifier) generateRequest(msg *Message) (*http.Request, error) {
	body := &map[string]interface{}{
		"message":  msg.GetMessage(),
		"priority": msg.GetPriority(),
		"title":    msg.GetTitle(),
		"extras":   msg.GetMetadata(),
	}

	bdata, e := json.Marshal(body)
	if e != nil {
		return nil, e
	}

	url := g.config.GetData()["host"].(string) + "/message"

	req, e1 := http.NewRequest("POST", url, bytes.NewReader(bdata))
	if e1 != nil {
		return nil, e1
	}

	req.Header["X-Gotify-Key"] = []string{g.config.GetData()["apiKey"].(string)}
	return req, nil
}

func (g *GotifyNotifier) parseResponse(resp *http.Response) error {
	if resp.StatusCode != 200 {
		return errors.New("gotify: unable to send message request successfully")
	}

	return nil
}

func (g *GotifyConfig) Validate() []error {
	return nil
}

func (g *GotifyConfig) GetData() map[string]interface{} {
	return map[string]interface{}{
		"host":   g.Host,
		"apiKey": g.APIKey,
	}
}
