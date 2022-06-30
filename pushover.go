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
	"errors"
	"net/http"
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

	// http.NewRequest("POST", urlgen.createURL())

	return nil, nil
}

func (p *PushoverNotifier) parseResponse(resp *http.Response) error {
	return nil
}

func (p *PushoverConfig) Validate() []error {
	return nil
}

func (p *PushoverConfig) GetData() map[string]interface{} {
	return map[string]interface{}{}
}
