// Copyright 2015 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package user

// Interface provides access to the "appengine/users" API methods.
type Interface interface {
	Current() *User
	CurrentOAuth(scopes ...string) (*User, error)

	IsAdmin() bool

	LoginURL(dest string) (string, error)
	LoginURLFederated(dest, identity string) (string, error)
	LogoutURL(dest string) (string, error)

	OAuthConsumerKey() (string, error)
}
