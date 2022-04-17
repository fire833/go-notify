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
	"net/url"
	"sync"
	"time"
)

var (
	errorNoURL error = errors.New("url object is not loaded within message")
)

type Message struct {
	m sync.RWMutex

	msg string

	title    string
	subtitle string
	priority int
	url      *url.URL

	metadata map[string]string

	timestamp int64
}

func NewMessage(msg string) *Message {

	m := &Message{
		msg:       msg,
		timestamp: time.Now().UnixNano(),
		title:     "Generic Notification",
		subtitle:  "Generic Notification Subtitle",
	}

	return m
}

func (msg *Message) SetMessage(new string) {
	msg.m.Lock()
	msg.msg = new
	msg.m.Unlock()
}

func (msg *Message) SetTitle(title string) {
	msg.m.Lock()
	msg.title = title
	msg.m.Unlock()
}

func (msg *Message) SetSubtitle(subtitle string) {
	msg.m.Lock()
	msg.subtitle = subtitle
	msg.m.Unlock()
}

func (msg *Message) SetPriority(prio int) {
	msg.m.Lock()
	msg.priority = prio
	msg.m.Unlock()
}

func (msg *Message) SetURL(url *url.URL) {
	msg.m.Lock()
	msg.url = url
	msg.m.Unlock()
}

func (msg *Message) GetMessage() string {
	msg.m.RLock()
	defer msg.m.RUnlock()
	return msg.msg
}

func (msg *Message) GetTitle() string {
	msg.m.RLock()
	defer msg.m.RUnlock()
	return msg.title
}

func (msg *Message) GetSubtitle() string {
	msg.m.RLock()
	defer msg.m.RUnlock()
	return msg.subtitle
}

func (msg *Message) GetPriority() int {
	msg.m.RLock()
	defer msg.m.RUnlock()
	return msg.priority
}

func (msg *Message) GetURL() (*url.URL, error) {
	msg.m.RLock()
	defer msg.m.RUnlock()
	if msg.url != nil {
		return msg.url, nil
	} else {
		return nil, errorNoURL
	}
}

func (msg *Message) AddKVMetadata(key, value string) {
	msg.m.Lock()
	msg.metadata[key] = value
	msg.m.Unlock()
}

func (msg *Message) GetMetadata() map[string]string {
	msg.m.RLock()
	defer msg.m.RUnlock()
	return msg.metadata
}

func (msg *Message) KeyExists(key string) bool {
	msg.m.RLock()
	defer msg.m.RUnlock()
	return msg.metadata[key] != ""
}

func (msg *Message) SetURLString(rawurl string) error {
	if u, e := url.Parse(rawurl); e == nil {
		msg.SetURL(u)
		return nil
	} else {
		return e
	}
}

func (msg *Message) String() string {
	msg.m.RLock()
	defer msg.m.RUnlock()
	return "Title: " + msg.title + "\n" + "Message: " + msg.msg + "\n"
}

// Implement the io.Reader interface for the message, and thus
// the message can be read into a byte buffer if needed.
//
// Utilizes the String() method underneath for the initial serialization
// of the object.
// func (msg *Message) Read(p []byte) (n int, err error) {

// }
