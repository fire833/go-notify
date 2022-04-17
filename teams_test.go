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
	"reflect"
	"sync"
	"testing"
)

func TestTeamsNotifier_generateRequest(t *testing.T) {
	type fields struct {
		RWMutex sync.RWMutex
		config  *TeamsConfig
		closed  bool
	}
	type args struct {
		msg *Message
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *http.Request
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &TeamsNotifier{
				RWMutex: tt.fields.RWMutex,
				config:  tt.fields.config,
				closed:  tt.fields.closed,
			}
			if got := tr.generateRequest(tt.args.msg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TeamsNotifier.generateRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTeamsNotifier_parseResponse(t *testing.T) {
	type fields struct {
		RWMutex sync.RWMutex
		config  *TeamsConfig
		closed  bool
	}
	type args struct {
		in0 *http.Response
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &TeamsNotifier{
				RWMutex: tt.fields.RWMutex,
				config:  tt.fields.config,
				closed:  tt.fields.closed,
			}
			if err := tr.parseResponse(tt.args.in0); (err != nil) != tt.wantErr {
				t.Errorf("TeamsNotifier.parseResponse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTeamsNotifier_SendMessage(t *testing.T) {
	type fields struct {
		RWMutex sync.RWMutex
		config  *TeamsConfig
		closed  bool
	}
	type args struct {
		msg *Message
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &TeamsNotifier{
				RWMutex: tt.fields.RWMutex,
				config:  tt.fields.config,
				closed:  tt.fields.closed,
			}
			if err := tr.SendMessage(tt.args.msg); (err != nil) != tt.wantErr {
				t.Errorf("TeamsNotifier.SendMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTeamsConfig_Validate(t *testing.T) {
	type fields struct {
		WebhookURL *url.URL
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
			c := &TeamsConfig{
				WebhookURL: tt.fields.WebhookURL,
			}
			if got := c.Validate(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TeamsConfig.Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}
