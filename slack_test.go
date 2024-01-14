/*
*	Copyright (C) 2024 Kendall Tauser
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

func TestSlackNotifier_generateRequest(t *testing.T) {
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
			s := &SlackNotifier{}
			got, err := s.generateRequest(tt.msg)
			if (err != nil) != tt.wantErr {
				t.Errorf("SlackNotifier.generateRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SlackNotifier.generateRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSlackNotifier_parseResponse(t *testing.T) {
	tests := []struct {
		name    string
		resp    *http.Response
		wantErr bool
	}{
		{
			name: "1",
			resp: &http.Response{
				StatusCode: http.StatusNotAcceptable,
			},
			wantErr: true,
		},
		{
			name: "2",
			resp: &http.Response{
				StatusCode: http.StatusOK,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SlackNotifier{}
			if err := s.parseResponse(tt.resp); (err != nil) != tt.wantErr {
				t.Errorf("SlackNotifier.parseResponse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSlackNotifier_validateMessage(t *testing.T) {
	tests := []struct {
		name    string
		msg     *Message
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SlackNotifier{}
			if err := s.validateMessage(tt.msg); (err != nil) != tt.wantErr {
				t.Errorf("SlackNotifier.validateMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSlackConfig_GetData(t *testing.T) {
	type fields struct {
		WebhookURL string
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
			c := &SlackConfig{
				WebhookURL: tt.fields.WebhookURL,
			}
			if got := c.GetData(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SlackConfig.GetData() = %v, want %v", got, tt.want)
			}
		})
	}
}
