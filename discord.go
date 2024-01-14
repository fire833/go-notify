/*
*	Copyright (C) 2024 Kendall Tauser
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

type DiscordNotifier struct {
	genericHTTPNotifier
}

type DiscordConfig struct {
	Id    string
	Token string
}

func NewDiscordNotifier() *DiscordNotifier {
	return &DiscordNotifier{}
}

func NewDiscordNotifierConfig(config *DiscordConfig) (*DiscordNotifier, error) {
	n := NewDiscordNotifier()
	return n, n.Configure(config)
}

func NewDiscordNotifierConfigMust(config *DiscordConfig) *DiscordNotifier {
	n := NewDiscordNotifier()
	if e := n.Configure(config); e != nil {
		panic(e)
	}

	return n
}

func NewDefaultDiscordConfig() *DiscordConfig {
	return &DiscordConfig{
		Id:    "",
		Token: "",
	}
}

func (d *DiscordNotifier) SendMessage(msg *Message) error {
	return d.sendMessageInternal(msg, d.generateRequest, d.parseResponse, d.validateMessage)
}

func (d *DiscordNotifier) validateMessage(msg *Message) error {
	if msg.GetMessage() == "" {
		return errors.New("discord: message body is required for notification")
	}

	return nil
}

func (d *DiscordNotifier) generateRequest(msg *Message) (*http.Request, error) {
	body := &map[string]interface{}{
		"content": msg.GetMessage(),
	}

	bdata, e := json.Marshal(body)
	if e != nil {
		return nil, e
	}

	req, e1 := http.NewRequest("POST", discordURLgen(d.config.GetData()["id"].(string), d.config.GetData()["token"].(string)), bytes.NewReader(bdata))
	if e1 != nil {
		return nil, e1
	}

	req.Header.Set("Content-Type", "application/json")
	return req, nil
}

func discordURLgen(id string, token string) (url string) {
	return "https://discord.com/api/webhooks/" + id + "/" + token
}

func (d *DiscordNotifier) parseResponse(resp *http.Response) error {
	if resp.StatusCode != 200 {
		return errors.New("discord: unable to send message request successfully")
	}

	return nil
}

func (c *DiscordConfig) Validate() []error {
	var retErr []error

	if c.Id == "" {
		retErr = append(retErr, errors.New("discord: webhook id must not be nil"))
	}

	if c.Token == "" {
		retErr = append(retErr, errors.New("discord: webhook token must not be nil"))
	}

	return retErr
}

func (c *DiscordConfig) GetData() map[string]interface{} {
	return map[string]interface{}{
		"id":    c.Id,
		"token": c.Token,
	}
}
