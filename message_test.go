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
	crand "crypto/rand"
	"math/big"
	"math/rand"
	"testing"
)

func init() {
	i, _ := crand.Int(crand.Reader, big.NewInt(1234567808992831938))
	rand.Seed(i.Int64())
}

func TestMessage_KeyExists(t *testing.T) {
	msg1 := NewMessage("")
	msg1.metadata["lorem"] = 1234
	msg1.metadata["loremipsum"] = 12345
	msg1.metadata["loremipsumdolor"] = "string"
	msg1.metadata["loremipsumdolorsit"] = "string"
	msg2 := NewMessage("a test object")
	msg2.metadata["obj1"] = "a string"
	msg3 := NewMessage("12345")
	msg3.metadata["obj1"] = -1234567
	msg4 := NewMessage("12345")

	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want bool
		msg  *Message
	}{
		{
			name: "1",
			msg:  msg1,
			want: true,
			args: args{
				key: "lorem",
			},
		},
		{
			name: "2",
			msg:  msg1,
			want: true,
			args: args{
				key: "loremipsum",
			},
		},
		{
			name: "3",
			msg:  msg1,
			want: true,
			args: args{
				key: "loremipsumdolor",
			},
		},
		{
			name: "4",
			msg:  msg1,
			want: true,
			args: args{
				key: "loremipsumdolorsit",
			},
		},
		{
			name: "5",
			msg:  msg1,
			want: false,
			args: args{
				key: "loremipsumdolorsitamet",
			},
		},
		{
			name: "6",
			msg:  msg1,
			want: false,
			args: args{
				key: "12345",
			},
		},
		{
			name: "7",
			msg:  msg1,
			want: false,
			args: args{
				key: "3.14159",
			},
		},
		{
			name: "8",
			msg:  msg2,
			want: true,
			args: args{
				key: "obj1",
			},
		},
		{
			name: "9",
			msg:  msg2,
			want: false,
			args: args{
				key: "obj12345",
			},
		},
		{
			name: "10",
			msg:  msg2,
			want: false,
			args: args{
				key: "string",
			},
		},
		{
			name: "11",
			msg:  msg4,
			want: false,
			args: args{
				key: "string",
			},
		},
		{
			name: "12",
			msg:  msg3,
			want: false,
			args: args{
				key: "string",
			},
		},
		{
			name: "13",
			msg:  msg3,
			want: true,
			args: args{
				key: "obj1",
			},
		},
		{
			name: "14",
			msg:  msg4,
			want: false,
			args: args{
				key: "obj1",
			},
		},
		{
			name: "15",
			msg:  msg3,
			want: false,
			args: args{
				key: "obj12345",
			},
		},
		{
			name: "16",
			msg:  msg3,
			want: false,
			args: args{
				key: "a string",
			},
		},
		{
			name: "17",
			msg:  msg3,
			want: false,
			args: args{
				key: "a string 2",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.msg.KeyExists(tt.args.key); got != tt.want {
				t.Errorf("Message.KeyExists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMessage_SetMessage(t *testing.T) {
	// msg1 := NewMessage("12345")

	// wg := new(sync.WaitGroup)

	// count := 0

	// var arr []string
	// arr = make([]string, 100000)

	// for i := 0; i < 99999; i++ {
	// 	arr[i] = fmt.Sprintf("%v", rand.Int())
	// }

	// wg.Add(2)

	// go func() {
	// 	for i := 0; i < 99999; i++ {
	// 		msg1.SetMessage(arr[i])
	// 		count++
	// 	}

	// 	wg.Done()
	// }()

	// go func() {
	// 	for i := 0; i < 100; i++ {
	// 		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
	// 			m := msg1.GetMessage()

	// 			if m != arr[count] {
	// 				t.Errorf("Message.SetMessage() = %v, want %v", m, arr[count])
	// 			}
	// 		})
	// 	}
	// }()

	// wg.Wait()

}
