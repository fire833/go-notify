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
	"net/http"
	"reflect"
	"testing"
)

func TestDiscordConfig_Validate(t *testing.T) {
	var nilErr []error

	type fields struct {
		Id    string
		Token string
	}
	tests := []struct {
		name   string
		fields fields
		want   []error
	}{
		{
			name: "1",
			fields: fields{
				Id:    "34789573489",
				Token: "347589hsdjkfhsd",
			},
			want: nilErr,
		},
		{
			name: "2",
			fields: fields{
				Id:    "",
				Token: "2374897348956uasga",
			},
			want: []error{errors.New("discord: webhook id must not be nil")},
		},
		{
			name: "3",
			fields: fields{
				Id:    "374897234895",
				Token: "",
			},
			want: []error{errors.New("discord: webhook token must not be nil")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &DiscordConfig{
				Id:    tt.fields.Id,
				Token: tt.fields.Token,
			}
			if got := c.Validate(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DiscordConfig.Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDiscordConfig_GetData(t *testing.T) {
	type fields struct {
		Id    string
		Token string
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]interface{}
	}{
		{
			name: "1",
			fields: fields{
				Id:    "83457",
				Token: "234895730sdmkfbhsdjkfh",
			},
			want: map[string]interface{}{
				"id":    "83457",
				"token": "234895730sdmkfbhsdjkfh",
			},
		},
		{
			name: "2",
			fields: fields{
				Id:    "something",
				Token: "a token",
			},
			want: map[string]interface{}{
				"id":    "something",
				"token": "a token",
			},
		},
		{
			name: "3",
			fields: fields{
				Id:    "8349347t3489908-0fguxcklgh",
				Token: "347895yfuisghfmndged478tye8fgve4fvyegcfdfjydjytdci",
			},
			want: map[string]interface{}{
				"id":    "8349347t3489908-0fguxcklgh",
				"token": "347895yfuisghfmndged478tye8fgve4fvyegcfdfjydjytdci",
			},
		},
		{
			name: "4",
			fields: fields{
				Id:    "8349347t3489908-3489573894573489",
				Token: "8346589hdfjksgfsdjkfgkuhvndjvidvhdcruivh",
			},
			want: map[string]interface{}{
				"id":    "8349347t3489908-3489573894573489",
				"token": "8346589hdfjksgfsdjkfgkuhvndjvidvhdcruivh",
			},
		},
		{
			name: "5",
			fields: fields{
				Id:    "348994705832901731890237190",
				Token: "sdhfjksfhsjkfgweruighxcjkvbjnxcbjhafduyhasgdjkbnacmnxcvmxchvxc",
			},
			want: map[string]interface{}{
				"id":    "348994705832901731890237190",
				"token": "sdhfjksfhsjkfgweruighxcjkvbjnxcbjhafduyhasgdjkbnacmnxcvmxchvxc",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &DiscordConfig{
				Id:    tt.fields.Id,
				Token: tt.fields.Token,
			}
			if got := c.GetData(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DiscordConfig.GetData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDiscordNotifier_validateMessage(t *testing.T) {
	tests := []struct {
		name    string
		msg     *Message
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DiscordNotifier{}
			if err := d.validateMessage(tt.msg); (err != nil) != tt.wantErr {
				t.Errorf("DiscordNotifier.validateMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDiscordNotifier_generateRequest(t *testing.T) {
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
			d := &DiscordNotifier{}
			got, err := d.generateRequest(tt.msg)
			if (err != nil) != tt.wantErr {
				t.Errorf("DiscordNotifier.generateRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DiscordNotifier.generateRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDiscordNotifier_parseResponse(t *testing.T) {
	tests := []struct {
		name    string
		resp    *http.Response
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DiscordNotifier{}
			if err := d.parseResponse(tt.resp); (err != nil) != tt.wantErr {
				t.Errorf("DiscordNotifier.parseResponse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_discordURLgen(t *testing.T) {
	type args struct {
		id    string
		token string
	}
	tests := []struct {
		name    string
		args    args
		wantUrl string
	}{
		{
			name: "1",
			args: args{
				id:    "73489573890723612",
				token: "8234rhuisdbfjksdfbi9023uycb589e4y5v78ctusyfuisy",
			},
			wantUrl: "https://discord.com/api/webhooks/73489573890723612/8234rhuisdbfjksdfbi9023uycb589e4y5v78ctusyfuisy",
		},
		{
			name: "2",
			args: args{
				id:    "12345",
				token: "892347589ydusgfsh12735x247as4gs3ycrtuykuzw;3r58",
			},
			wantUrl: "https://discord.com/api/webhooks/12345/892347589ydusgfsh12735x247as4gs3ycrtuykuzw;3r58",
		},
		{
			name: "3",

			args: args{
				id:    "734895734",
				token: "c3giugyxdvjkf89wey54ve896rv96s7v5mkwectgsr7kys89",
			},
			wantUrl: "https://discord.com/api/webhooks/734895734/c3giugyxdvjkf89wey54ve896rv96s7v5mkwectgsr7kys89",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotUrl := discordURLgen(tt.args.id, tt.args.token); gotUrl != tt.wantUrl {
				t.Errorf("discordURLgen() = %v, want %v", gotUrl, tt.wantUrl)
			}
		})
	}
}
