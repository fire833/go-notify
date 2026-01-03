/*
*	Copyright (C) 2026 Kendall Tauser
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
	"reflect"
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

func TestMessage_String(t *testing.T) {
	t1 := NewMessage("A new message.")
	t2 := NewMessage("")
	t2.SetMessage("3485yusdfhdfjkshfjkshf8934hjsdhfg")
	t3 := NewMessage("jksdhjklgio54u89ghdfjkghfm,ghiory")
	t3.SetMessage("12345")
	t3.SetMessage("jksdhjklgio54u89ghdfjkghfm,ghiory")
	t4 := NewMessage("12345")
	t4.SetTitle("A Title")

	tests := []struct {
		name string
		msg  *Message
		want string
	}{
		{
			name: "1",
			msg:  t1,
			want: "Title: \nMessage: A new message.\n",
		},
		{
			name: "2",
			msg:  t2,
			want: "Title: \nMessage: 3485yusdfhdfjkshfjkshf8934hjsdhfg\n",
		},
		{
			name: "3",
			msg:  t3,
			want: "Title: \nMessage: jksdhjklgio54u89ghdfjkghfm,ghiory\n",
		},
		{
			name: "4",
			msg:  t4,
			want: "Title: A Title\nMessage: 12345\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.msg.String(); got != tt.want {
				t.Errorf("Message.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMessage_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		msg     *Message
		want    []byte
		wantErr bool
	}{
		{
			name:    "1",
			msg:     NewMessage("foo bar").SetTitle("Hello World").SetPrivacy(2).SetPriority(5),
			want:    []byte(`{"title":"Hello World","message":"foo bar","priority":5,"privacy":2}`),
			wantErr: false,
		},
		{
			name:    "2",
			msg:     NewMessage("1234").SetTitle("foo").SetSubtitle("bar"),
			want:    []byte(`{"title":"foo","subtitle":"bar","message":"1234"}`),
			wantErr: false,
		},
		{
			name:    "3",
			msg:     NewMessage("").SetPrivacy(3),
			want:    []byte(`{"privacy":3}`),
			wantErr: false,
		},
		{
			name:    "4",
			msg:     NewMessage(""),
			want:    []byte(`{}`),
			wantErr: false,
		},
		{
			name:    "5",
			msg:     NewMessage("hello there world, this is a n important message"),
			want:    []byte(`{"message":"hello there world, this is a n important message"}`),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.msg.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("Message.MarshalJSON() error = %s, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Message.MarshalJSON() = %s, want %s", got, tt.want)
			}
		})
	}
}

func TestMessage_SetTitle(t *testing.T) {
	t1 := NewMessage("1234")
	t1.title = "foo1"

	tests := []struct {
		name  string
		msg   *Message
		title string
		want  *Message
	}{
		{
			name:  "1",
			msg:   NewMessage("1234"),
			want:  t1,
			title: "foo1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.msg.SetTitle(tt.title); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Message.SetTitle() = %v, want %v", got, tt.want)
			}
		})
	}
}
