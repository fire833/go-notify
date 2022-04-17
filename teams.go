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
	"net/url"
	"sync"

	"github.com/fire833/go-notify/pkg/common"
)

type TeamsNotifier struct {
	sync.RWMutex

	closed bool

	config *TeamsConfig
}

type TeamsConfig struct {
	webhooks map[string]*url.URL
}

func NewTeamsNotifier() *TeamsNotifier {
	return &TeamsNotifier{
		closed: false,
		config: nil,
	}
}

func (t *TeamsNotifier) SendMessage(msg *Message) error {
	t.RLock()
	defer t.RUnlock()

	if t.isReady() {

	}

	return nil
}

func (t *TeamsNotifier) Configure(config *TeamsConfig) error {

	return nil
}

func (t *TeamsNotifier) Close() error {
	t.Lock()
	defer t.Unlock()

	if t.isClosed() {
		return common.ErrorNotifierClosed
	}

	t.closed = true

	return nil
}

func (t *TeamsNotifier) isReady() bool {
	t.RLock()
	defer t.RUnlock()

	// If the notifier is closed, auto failout
	if t.isClosed() {
		return false
	}

	// If the notifier has not been configured, auto failout
	if t.config == nil {
		return false
	}

	return true
}

func (t *TeamsNotifier) isClosed() bool {
	t.RLock()
	defer t.RUnlock()

	if t.closed {
		return true
	} else {
		return false
	}
}

func (c *TeamsConfig) Validate() []error {
	return nil
}
