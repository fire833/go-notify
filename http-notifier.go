/*
*	Copyright (C) 2025 Kendall Tauser
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
	"sync"

	"github.com/fire833/go-notify/pkg/common"
)

// A generic implementation of an HTTP-based notifier,
// which sends requests through HTTP and returns errors
// on bad responses.
type genericHTTPNotifier struct {
	sync.RWMutex
	config NotifierConfig

	closed bool
}

func (n *genericHTTPNotifier) sendMessageInternal(msg *Message,
	genReq func(msg *Message) (*http.Request, error),
	parseResp func(*http.Response) error,
	validateFunc func(msg *Message) error,
) error {
	n.RLock()
	defer n.RUnlock()

	if n.isClosed() {
		return common.ErrorNotifierClosed
	}

	if e := validateFunc(msg); e != nil {
		return e
	}

	if n.isReady() {
		req, e := genReq(msg)
		if e != nil {
			return common.ErrorNotifierSerializationError
		}

		resp, e1 := common.NotifyHTTPTransporter.RoundTrip(req)
		e2 := parseResp(resp)

		if e1 != nil || e2 != nil {
			return common.ErrorNotificationSendError
		}

	} else {
		return common.ErrorNotifierNotReady
	}

	return nil
}

// Configure configures the notifier with proper configuration for its operation.
func (n *genericHTTPNotifier) Configure(config NotifierConfig) error {
	if e := config.Validate(); len(e) != 0 {
		return common.ErrorInvalidConfiguration
	}

	n.Lock()
	n.config = config
	n.Unlock()
	return nil
}

// Close closes out the notifier. Returns an error if unable to or if the Notifier
// has already been closed.
func (n *genericHTTPNotifier) Close() error {
	if n.isClosed() {
		return common.ErrorNotifierClosed
	}

	n.Lock()
	n.closed = true
	n.Unlock()
	return nil
}

func (n *genericHTTPNotifier) isReady() bool {
	n.RLock()
	defer n.RUnlock()
	return n.config != nil
}

func (n *genericHTTPNotifier) isClosed() bool {
	n.RLock()
	defer n.RUnlock()
	return n.closed
}
