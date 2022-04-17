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
	"time"
)

type Message struct {
	msg string

	title    string
	priority int
	url      *url.URL

	timestamp int64
}

func NewMessage(msg string) *Message {

	m := &Message{
		msg:       msg,
		timestamp: time.Now().UnixNano(),
		title:     "Generic Notification",
	}

	return m
}

func (msg *Message) SetMessage(new string) {
	msg.msg = new
}

func (msg *Message) SetTitle(title string) {
	msg.title = title
}

func (msg *Message) SetPriority(prio int) {
	msg.priority = prio
}

func (msg *Message) SetURL(url *url.URL) {
	msg.url = url
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
	return "Title: " + msg.title + "\n" + "Message: " + msg.msg + "\n"
}

// Implement the io.Reader interface for the message, and thus
// the message can be read into a byte buffer if needed.
//
// Utilizes the String() method underneath for the initial serialization
// of the object.
// func (msg *Message) Read(p []byte) (n int, err error) {

// }
