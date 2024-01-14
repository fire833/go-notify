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

import "errors"

var (
	ErrorNoBackend         error = errors.New("backend not found")
	ErrorNoBackendInstance error = errors.New("backend instance not found")

	ErrorNotifierClosed             error = errors.New("the desired notifier has been closed and decomissioned")
	ErrorNotifierNotReady           error = errors.New("the desired notifier is not ready yet")
	ErrorNotifierSerializationError error = errors.New("unable to parse message to request for backend to process")

	ErrorInvalidConfiguration error = errors.New("the configuration provided is invalid")

	ErrorNotificationSendError error = errors.New("unable to send notification via the desired backend")
)
