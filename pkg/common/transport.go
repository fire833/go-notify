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

package common

import "net/http"

// NotifyHTTPTransporter is the defacto HTTP client that should
// be utilized for all Notifier notification transportation that must
// take place over HTTP. This allow for custom transports to be
// implemented for testing purposes of this library.
//
// Should be by default set to the http.DefaultTransport, but a custom
// transport **MAY** be implemented in the future tooptimize for the
// nature of Notifications, but that is still up in the air.
var NotifyHTTPTransporter http.RoundTripper = http.DefaultTransport
