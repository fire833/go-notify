/*
*	Copyright (C) 2023 Kendall Tauser
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
	"testing"
)

func Test_genericHTTPNotifier_sendMessageInternal(t *testing.T) {
	type fields struct {
		config NotifierConfig
		closed bool
	}
	type args struct {
		msg          *Message
		genReq       func(msg *Message) (*http.Request, error)
		parseResp    func(*http.Response) error
		validateFunc func(msg *Message) error
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
			n := &genericHTTPNotifier{
				RWMutex: sync.RWMutex{},
				config:  tt.fields.config,
				closed:  tt.fields.closed,
			}
			if err := n.sendMessageInternal(tt.args.msg, tt.args.genReq, tt.args.parseResp, tt.args.validateFunc); (err != nil) != tt.wantErr {
				t.Errorf("genericHTTPNotifier.sendMessageInternal() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_genericHTTPNotifier_Configure(t *testing.T) {
	type fields struct {
		config NotifierConfig
		closed bool
	}
	type args struct {
		config NotifierConfig
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
			n := &genericHTTPNotifier{
				RWMutex: sync.RWMutex{},
				config:  tt.fields.config,
				closed:  tt.fields.closed,
			}
			if err := n.Configure(tt.args.config); (err != nil) != tt.wantErr {
				t.Errorf("genericHTTPNotifier.Configure() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_genericHTTPNotifier_Close(t *testing.T) {
	type fields struct {
		config NotifierConfig
		closed bool
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &genericHTTPNotifier{
				RWMutex: sync.RWMutex{},
				config:  tt.fields.config,
				closed:  tt.fields.closed,
			}
			if err := n.Close(); (err != nil) != tt.wantErr {
				t.Errorf("genericHTTPNotifier.Close() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_genericHTTPNotifier_isReady(t *testing.T) {
	type fields struct {
		config NotifierConfig
		closed bool
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &genericHTTPNotifier{
				RWMutex: sync.RWMutex{},
				config:  tt.fields.config,
				closed:  tt.fields.closed,
			}
			if got := n.isReady(); got != tt.want {
				t.Errorf("genericHTTPNotifier.isReady() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_genericHTTPNotifier_isClosed(t *testing.T) {
	type fields struct {
		config NotifierConfig
		closed bool
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &genericHTTPNotifier{
				RWMutex: sync.RWMutex{},
				config:  tt.fields.config,
				closed:  tt.fields.closed,
			}
			if got := n.isClosed(); got != tt.want {
				t.Errorf("genericHTTPNotifier.isClosed() = %v, want %v", got, tt.want)
			}
		})
	}
}
