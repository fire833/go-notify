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
	"errors"
	"net/url"
	"sync"
	"time"

	"github.com/fire833/go-notify/pkg/utils"
)

var errorNoURL error = errors.New("url object is not loaded within message")

type Message struct {
	m sync.RWMutex

	msg string // `json:"message" yaml:"message"`

	title    string   // `json:"title,omitempty" yaml:"title,omitempty"`
	subtitle string   // `json:"subtitle,omitempty" yaml:"subtitle,omitempty"`
	priority int      // `json:"priority,omitempty" yaml:"priority,omitempty"`
	privacy  int      // `json:"privacy,omitempty" yaml:"privacy,omitempty"`
	url      *url.URL // `json:"url,omitempty" yaml:"url,omitempty"`

	metadata map[string]interface{} // `json:"metadata,omitempty" yaml:"metadata,omitempty"`

	timestamp int64 // `json:"timestamp" yaml:"timestamp"`
}

func NewMessage(msg string) *Message {
	m := &Message{
		msg:       msg,
		timestamp: time.Now().Unix(),
		title:     "",
		subtitle:  "",
		metadata:  map[string]interface{}{},
		url:       nil,
		priority:  -1,
		privacy:   -1,
	}

	return m
}

func (msg *Message) SetMessage(new string) *Message {
	msg.m.Lock()
	msg.msg = new
	msg.m.Unlock()
	return msg
}

func (msg *Message) SetTitle(title string) *Message {
	msg.m.Lock()
	msg.title = title
	msg.m.Unlock()
	return msg
}

func (msg *Message) SetSubtitle(subtitle string) *Message {
	msg.m.Lock()
	msg.subtitle = subtitle
	msg.m.Unlock()
	return msg
}

func (msg *Message) SetPriority(prio uint) *Message {
	msg.m.Lock()
	msg.priority = int(prio)
	msg.m.Unlock()
	return msg
}

func (msg *Message) SetPrivacy(priv uint) *Message {
	msg.m.Lock()
	msg.privacy = int(priv)
	msg.m.Unlock()
	return msg
}

func (msg *Message) SetURL(url *url.URL) *Message {
	msg.m.Lock()
	msg.url = url
	msg.m.Unlock()
	return msg
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

func (msg *Message) GetPrivacy() int {
	msg.m.RLock()
	defer msg.m.RUnlock()
	return msg.privacy
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

func (msg *Message) AddKVMetadata(key string, value interface{}) {
	msg.m.Lock()
	msg.metadata[key] = value
	msg.m.Unlock()
}

func (msg *Message) GetMetadata() map[string]interface{} {
	msg.m.RLock()
	defer msg.m.RUnlock()
	return msg.metadata
}

// Returns whether the defined key has a value inside of it.
func (msg *Message) KeyExists(key string) bool {
	msg.m.RLock()
	defer msg.m.RUnlock()
	return msg.metadata[key] != nil
}

func (msg *Message) GetValue(key string) interface{} {
	msg.m.RLock()
	defer msg.m.RUnlock()
	return msg.metadata[key]
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

func (msg *Message) MarshalJSON() ([]byte, error) {
	var buf []byte

	buf = append(buf, '{')

	msg.m.RLock()

	found := false

	if msg.title != "" {
		if !found {
			found = true
		}
		buf = utils.AppendJSONKV(buf, "title", msg.title)
	}

	if msg.subtitle != "" {
		if !found {
			found = true
		} else {
			buf = utils.AppendComma(buf)
		}
		buf = utils.AppendJSONKV(buf, "subtitle", msg.subtitle)
	}

	if msg.msg != "" {
		if !found {
			found = true
		} else {
			buf = utils.AppendComma(buf)
		}
		buf = utils.AppendJSONKV(buf, "message", msg.msg)
	}

	if msg.url != nil {
		if !found {
			found = true
		} else {
			buf = utils.AppendComma(buf)
		}
		buf = utils.AppendJSONKV(buf, "url", msg.url.String())
	}

	if msg.priority != -1 {
		if !found {
			found = true
		} else {
			buf = utils.AppendComma(buf)
		}
		buf = utils.AppendJSONKV(buf, "priority", msg.priority)
	}

	if msg.privacy != -1 {
		if !found {
			found = true
		} else {
			buf = utils.AppendComma(buf)
		}
		buf = utils.AppendJSONKV(buf, "privacy", msg.privacy)
	}

	buf = append(buf, '}')

	msg.m.RUnlock()
	return buf, nil
}

// Implement the io.Reader interface for the message, and thus
// the message can be read into a byte buffer if needed.
//
// Utilizes the String() method underneath for the initial serialization
// of the object.
// func (msg *Message) Read(p []byte) (n int, err error) {

// }
