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

package utils

import (
	"reflect"
	"testing"
)

func TestAppendJSONKV(t *testing.T) {
	type args struct {
		buf   []byte
		key   string
		value interface{}
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "1",
			args: args{
				buf:   []byte(""),
				key:   "value1",
				value: 12345,
			},
			want: []byte(`"value1":12345`),
		},
		{
			name: "2",
			args: args{
				buf:   []byte(""),
				key:   "value2",
				value: "string here",
			},
			want: []byte(`"value2":"string here"`),
		},
		{
			name: "3",
			args: args{
				buf:   []byte(""),
				key:   "value3",
				value: -34789573489,
			},
			want: []byte(`"value3":-34789573489`),
		},
		{
			name: "4",
			args: args{
				buf:   []byte(""),
				key:   "value4",
				value: 'a',
			},
			want: []byte(`"value4":97`),
		},
		{
			name: "5",
			args: args{
				buf:   []byte(""),
				key:   "value5",
				value: "a string here",
			},
			want: []byte(`"value5":"a string here"`),
		},
		{
			name: "6",
			args: args{
				buf:   []byte(""),
				key:   "value6",
				value: 20.3456,
			},
			want: []byte(`"value6":20.3456`),
		},
		{
			name: "7",
			args: args{
				buf:   []byte(""),
				key:   "value7",
				value: true,
			},
			want: []byte(`"value7":true`),
		},
		{
			name: "8",
			args: args{
				buf:   []byte(""),
				key:   "value8",
				value: false,
			},
			want: []byte(`"value8":false`),
		},
		{
			name: "9",
			args: args{
				buf:   []byte(""),
				key:   "value9",
				value: '1',
			},
			want: []byte(`"value9":49`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AppendJSONKV(tt.args.buf, tt.args.key, tt.args.value)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AppendJSONKV() = %s, want %s\n", string(got), string(tt.want))
			}
		})
	}
}
