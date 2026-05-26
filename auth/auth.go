// Copyright 2015 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package auth

import (
	"github.com/pingcap/parser/format"
)

// UserIdentity represents username and hostname.
type UserIdentity struct {
	Username     string
	Hostname     string
	CurrentUser  bool
	AuthUsername string // Username matched in privileges system
	AuthHostname string // Match in privs system (i.e. could be a wildcard)
}

// Restore implements Node interface.
func (user *UserIdentity) Restore(ctx *format.RestoreCtx) error {
	_ = "STUB: not implemented"
	return nil
}

// String converts UserIdentity to the format user@host.
func (user *UserIdentity) String() string {
	_ = "STUB: not implemented"
	// TODO: Escape username and hostname.
	return ""
}

// AuthIdentityString returns matched identity in user@host format
func (user *UserIdentity) AuthIdentityString() string {
	_ = "STUB: not implemented"
	// TODO: Escape username and hostname.
	return ""
}

type RoleIdentity struct {
	Username string
	Hostname string
}

func (role *RoleIdentity) Restore(ctx *format.RestoreCtx) error {
	_ = "STUB: not implemented"
	return nil
}

// String converts UserIdentity to the format user@host.
func (role *RoleIdentity) String() string {
	_ = "STUB: not implemented"
	// TODO: Escape username and hostname.
	return ""
}
