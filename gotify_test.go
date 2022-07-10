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
	"reflect"
	"testing"
)

func TestGotifyConfig_Validate(t *testing.T) {
	tests := []struct {
		name string
		g    *GotifyConfig
		want []error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GotifyConfig{}
			if got := g.Validate(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GotifyConfig.Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGotifyConfig_GetData(t *testing.T) {
	tests := []struct {
		name string
		g    *GotifyConfig
		want map[string]interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GotifyConfig{}
			if got := g.GetData(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GotifyConfig.GetData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGotifyNotifier_validateMessage(t *testing.T) {
	msg1 := NewMessage("a message goes here")
	msg1.SetTitle("")
	msg1.SetPriority(0)
	msg2 := NewMessage("a message that occurred")
	msg2.SetPriority(2)
	msg2.SetTitle("title")
	msg3 := NewMessage("")
	msg4 := NewMessage("a message")
	msg4.SetPriority(6)
	msg4.SetTitle("")
	msg5 := NewMessage("237483749")
	msg5.SetTitle("The title of this message")
	msg5.SetPriority(-3)

	tests := []struct {
		name    string
		msg     *Message
		wantErr bool
	}{
		{
			name:    "1",
			msg:     msg1,
			wantErr: true,
		},
		{
			name:    "2",
			msg:     msg2,
			wantErr: false,
		},
		{
			name:    "3",
			msg:     msg3,
			wantErr: true,
		},
		{
			name:    "4",
			msg:     msg4,
			wantErr: true,
		},
		{
			name:    "5",
			msg:     msg5,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GotifyNotifier{}
			if err := g.validateMessage(tt.msg); (err != nil) != tt.wantErr {
				t.Errorf("GotifyNotifier.validateMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGotifyNotifier_generateRequest(t *testing.T) {
	tests := []struct {
		name    string
		msg     *Message
		want    *http.Request
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GotifyNotifier{}
			got, err := g.generateRequest(tt.msg)
			if (err != nil) != tt.wantErr {
				t.Errorf("GotifyNotifier.generateRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GotifyNotifier.generateRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGotifyNotifier_parseResponse(t *testing.T) {
	tests := []struct {
		name    string
		resp    *http.Response
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GotifyNotifier{}
			if err := g.parseResponse(tt.resp); (err != nil) != tt.wantErr {
				t.Errorf("GotifyNotifier.parseResponse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
