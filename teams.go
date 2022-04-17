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
	"net/http"
	"net/url"
	"sync"

	"github.com/fire833/go-notify/pkg/common"
)

var (
	localTransport http.RoundTripper = common.NotifyHTTPTransporter
)

type teamsNotificationReq struct {
}

type TeamsNotifier struct {
	sync.RWMutex
	config *TeamsConfig

	closed bool
}

type TeamsConfig struct {
	WebhookURL *url.URL
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
		r, e := localTransport.RoundTrip(t.generateRequest(msg))
		e1 := t.parseResponse(r)

		if e != nil || e1 != nil {
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

func (t *TeamsNotifier) generateRequest(msg *Message) *http.Request {
	return nil
}

func (t *TeamsNotifier) parseResponse(*http.Response) error {
	return nil
}

func (c *TeamsConfig) Validate() []error {
	return nil
}
