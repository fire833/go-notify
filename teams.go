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
	"net/url"
	"sync"

	"github.com/fire833/go-notify/pkg/common"
)

var ()

type TeamsNotifier struct {
	sync.RWMutex
	config *TeamsConfig

	closed bool
}

type TeamsConfig struct {
	WebhookURL *url.URL

	Color string
}

func NewTeamsNotifier() *TeamsNotifier {
	return &TeamsNotifier{
		closed: false,
		config: nil,
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

func (t *TeamsNotifier) SendMessage(msg *Message) error {
	t.RLock()
	defer t.RUnlock()

	if t.isClosed() {
		return common.ErrorNotifierClosed
	}

	if t.isReady() {
		req, e := t.generateRequest(msg)
		if e != nil {
			return common.ErrorNotifierSerializationError
		}

		resp, e1 := common.NotifyHTTPTransporter.RoundTrip(req)
		e2 := t.parseResponse(resp)

		if e1 != nil || e2 != nil {
			return common.ErrorNotificationSendError
		}

	} else {
		return common.ErrorNotifierNotReady
	}

	return nil
}

// Configure configures the notifier with proper configuration for its operation.
func (t *TeamsNotifier) Configure(config *TeamsConfig) error {
	if e := config.Validate(); e != nil {
		return common.ErrorInvalidConfiguration
	}

	t.Lock()
	t.config = config
	t.Unlock()
	return nil
}

// Close closes out the notifier. Returns an error if unable to or if the Notifier
// has already been closed.
func (t *TeamsNotifier) Close() error {

	if t.isClosed() {
		return common.ErrorNotifierClosed
	}

	t.Lock()
	t.closed = true
	t.Unlock()
	return nil
}

func (t *TeamsNotifier) isReady() bool {
	t.RLock()
	defer t.RUnlock()
	return t.config != nil
}

func (t *TeamsNotifier) isClosed() bool {
	t.RLock()
	defer t.RUnlock()
	return t.closed
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
		"themeColor": t.config.Color,

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

	return http.NewRequest("POST", t.config.WebhookURL.String(), bytes.NewReader(bdata))

}

func (t *TeamsNotifier) parseResponse(*http.Response) error {
	return nil
}

func (c *TeamsConfig) Validate() []error {
	return nil
}
