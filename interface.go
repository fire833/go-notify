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

// Notifier is a simple interface for objects that implements a
// notification backend. This is not included in the interface
// spec, but all Notifiers should also have a New() function,
// so that the end user can easily instantiate the container.
type Notifier interface {

	// The notifier should get set with it's new configuration.
	// This should be thread safe, and can be called at both
	// initialization of the notifier or at runtime whenever
	// the configuration legitimately changes for whatever reason.
	//
	// Should be thread-safe to call at any point at runtime.
	Configure(config NotifierConfig) error

	// SendMessage basically does what you think it would do, it sends
	// a notification message over the desired Notifier and reports back
	// the results. Returns an error if the operation failed. The possible
	// errors can be found in pkg/common/errors.go.
	SendMessage(msg *Message) error

	// Close closes out the transport and deinitialize any state with the Notifier.
	// After Close is called, the Notifier should automatically fail out any operations
	// that could return an error, returning ErrorNotifierClosed.
	Close() error

	isReady() bool
	isClosed() bool
}

// NotifierConfig implements a generic interface for specifying the
// configurations for Notifiers.
type NotifierConfig interface {

	// Validates the configuration struct and returns the errors
	// that are associated with said configuration. If len(errors)
	// is 0, then the configuration should be considered valid.
	Validate() []error

	// Returns the key value data from the configuration.
	GetData() map[string]interface{}
}
