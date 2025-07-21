/*
*	Copyright (C) 2025 Kendall Tauser
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

package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	gonotify "github.com/fire833/go-notify"
)

func makeRequest(method, path, bodyString string, headers map[string]string) (string, error) {
	req, e := http.NewRequest(method, "http://localhost:8080"+path, bytes.NewReader([]byte(bodyString)))

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	res, e := http.DefaultClient.Do(req)
	if e != nil {
		return "", e
	}

	if res.StatusCode != 200 {
		return "", fmt.Errorf("returned %d not 200 to login", res.StatusCode)
	}

	body, e := io.ReadAll(res.Body)
	if e != nil {
		return "", e
	}

	data := map[string]interface{}{}
	if e := json.Unmarshal(body, &data); e != nil {
		return "", e
	}

	tok, ok := data["token"]
	if !ok {
		return "", errors.New("unable to retrieve app token")
	}

	return tok.(string), nil
}

func main() {
	notifier, e := gonotify.NewGotifyNotifierConfig(&gonotify.GotifyConfig{
		BaseURL: "http://localhost:8080",
		APIKey:  os.Getenv("GOTIFY_KEY"),
	})
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}

	if e := notifier.SendMessage(gonotify.NewMessage("Something happened").SetTitle("Random title").SetPriority(2).AddKVMetadata("Hello", "Idiot")); e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
}
