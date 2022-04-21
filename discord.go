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

type DiscordNotifier struct {
	sync.RWMutex
	config *DiscordConfig

	closed bool
}

type DiscordConfig struct {
	WebhookURL *url.URL
}

func NewDiscordNotifier() *DiscordNotifier {
	return &DiscordNotifier{
		closed: false,
		config: nil,
	}
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

func (d *DiscordNotifier) SendMessage(msg *Message) error {
	d.RLock()
	defer d.RUnlock()

	if d.isClosed() {
		return common.ErrorNotifierClosed
	}

	if d.isReady() {
		req, e := d.generateRequest(msg)
		if e != nil {
			return common.ErrorNotifierSerializationError
		}

		resp, e1 := common.NotifyHTTPTransporter.RoundTrip(req)
		e2 := d.parseResponse(resp)

		if e1 != nil || e2 != nil {
			return common.ErrorNotificationSendError
		}

	} else {
		return common.ErrorNotifierNotReady
	}

	return nil
}

// Configure configures the notifier with proper configuration for its operation.
func (d *DiscordNotifier) Configure(config *DiscordConfig) error {
	if e := config.Validate(); e != nil {
		return common.ErrorInvalidConfiguration
	}

	d.Lock()
	d.config = config
	d.Unlock()
	return nil
}

// Close closes out the notifier. Returns an error if unable to or if the Notifier
// has already been closed.
func (d *DiscordNotifier) Close() error {

	if d.isClosed() {
		return common.ErrorNotifierClosed
	}

	d.Lock()
	d.closed = true
	d.Unlock()
	return nil
}

func (d *DiscordNotifier) isReady() bool {
	d.RLock()
	defer d.RUnlock()
	return d.config != nil
}

func (d *DiscordNotifier) isClosed() bool {
	d.RLock()
	defer d.RUnlock()
	return d.closed
}

func (d *DiscordNotifier) generateRequest(msg *Message) (*http.Request, error) {
	return nil, nil
}

func (d *DiscordNotifier) parseResponse(*http.Response) error {
	return nil
}

func (c *DiscordConfig) Validate() []error {
	return nil
}
