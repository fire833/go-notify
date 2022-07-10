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

func TestPushoverNotifier_generateRequest(t *testing.T) {
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
			p := &PushoverNotifier{}
			got, err := p.generateRequest(tt.msg)
			if (err != nil) != tt.wantErr {
				t.Errorf("PushoverNotifier.generateRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PushoverNotifier.generateRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPushoverNotifier_parseResponse(t *testing.T) {
	tests := []struct {
		name    string
		resp    *http.Response
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PushoverNotifier{}
			if err := p.parseResponse(tt.resp); (err != nil) != tt.wantErr {
				t.Errorf("PushoverNotifier.parseResponse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPushoverConfig_Validate(t *testing.T) {
	type fields struct {
		APIKey  string
		Userkey string
	}
	tests := []struct {
		name   string
		fields fields
		want   []error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PushoverConfig{
				APIKey:  tt.fields.APIKey,
				Userkey: tt.fields.Userkey,
			}
			if got := p.Validate(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PushoverConfig.Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPushoverConfig_GetData(t *testing.T) {
	type fields struct {
		APIKey  string
		Userkey string
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PushoverConfig{
				APIKey:  tt.fields.APIKey,
				Userkey: tt.fields.Userkey,
			}
			if got := p.GetData(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PushoverConfig.GetData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPushoverNotifier_validateMessage(t *testing.T) {
	tests := []struct {
		name    string
		msg     *Message
		wantErr bool
	}{
		{
			name:    "1",
			msg:     NewMessage(""),
			wantErr: true,
		},
		{
			name:    "2",
			msg:     NewMessage("3489rsdfhsdjkvhxcjkvhxcv"),
			wantErr: false,
		},
		{
			name:    "3",
			msg:     NewMessage("a random notification message here"),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PushoverNotifier{}
			if err := p.validateMessage(tt.msg); (err != nil) != tt.wantErr {
				t.Errorf("PushoverNotifier.validateMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
