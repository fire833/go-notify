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
)

type DiscordNotifier struct {
	genericHTTPNotifier
}

type DiscordConfig struct {
	WebhookURL *url.URL
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
	return &DiscordConfig{}
}

func (d *DiscordNotifier) SendMessage(msg *Message) error {
	return d.sendMessageInternal(msg, d.generateRequest, d.parseResponse, d.validateMessage)
}

func (d *DiscordNotifier) validateMessage(msg *Message) error {
	return nil
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

func (c *DiscordConfig) GetData() map[string]interface{} {
	return map[string]interface{}{}
}
