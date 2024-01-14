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

package utils

import "fmt"

func AppendNSpaces(buf []byte, num int) []byte {
	for i := 0; i < num; i++ {
		buf = append(buf, ' ')
	}
	return buf
}

func AppendNL(buf []byte) []byte {
	buf = append(buf, '\n')
	return buf
}

func AppendCommaNL(buf []byte) []byte {
	buf = append(buf, ',')
	AppendNL(buf)
	return buf
}

func AppendComma(buf []byte) []byte {
	buf = append(buf, ',')
	return buf
}

// Adds a new line of JSON w/ a k/v pair.
// Like this:
// "key":value
//
// Done with no preceding spaces and no succeeding comma.
func AppendJSONKV(buf []byte, key string, value interface{}) []byte {
	kbytes := []byte(key)

	buf = append(buf, '"')
	buf = append(buf, kbytes...)

	buf = append(buf, `":`...)

	switch value.(type) {
	case string:
		{
			b := []byte(fmt.Sprintf("%s", value))

			buf = append(buf, `"`...)
			buf = append(buf, b...)
			buf = append(buf, `"`...)
		}
	default:
		{
			b := []byte(fmt.Sprintf("%v", value))

			buf = append(buf, b...)
		}
	}

	return buf
}
