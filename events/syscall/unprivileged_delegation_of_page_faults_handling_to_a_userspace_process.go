//go:build linux
// +build linux

// SPDX-License-Identifier: Apache-2.0
/*
Copyright (C) 2024 The Falco Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package syscall

import (
	"github.com/falcosecurity/event-generator/events"
	"golang.org/x/sys/unix"
)

var _ = events.Register(
	UnprivilegedDelegationofPageFaultsHandlingtoaUserspaceProcess,
	events.WithDisabled(), // this rules is not included in falco_rules.yaml (stable rules), so disable the action
)

func UnprivilegedDelegationofPageFaultsHandlingtoaUserspaceProcess(h events.Helper) error {
	// To make user.uid != 0
	if err := becameUser(h, "daemon"); err != nil {
		return err
	}
	// Attempt to create userfaultfd syscall is enough
	unix.Syscall(unix.SYS_USERFAULTFD, 0, 0, 0)
	return nil
}
