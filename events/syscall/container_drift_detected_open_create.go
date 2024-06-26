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
	"os"

	"github.com/falcosecurity/event-generator/events"
)

var _ = events.Register(
	ContainerDriftDetectedOpenCreate,
	events.WithDisabled(), // this rules is not included in falco_rules.yaml (stable rules), so disable the action
)

func ContainerDriftDetectedOpenCreate(h events.Helper) error {
	if h.InContainer() {
		// Create a unique file under tmp dir
		file, err := os.CreateTemp(os.TempDir(), "created-by-falco-event-generator-")
		if err != nil {
			h.Log().WithError(err).Error("Error Creating an empty file")
			return err
		}
		defer os.Remove(file.Name()) // Remove the file after function return
		h.Log().Infof("writing to %s", file.Name())
		return os.WriteFile(file.Name(), nil, os.FileMode(0755)) // Also set execute permission
	}
	return &events.ErrSkipped{
		Reason: "'Container Drift Detected (open+create)' is applicable only to containers.",
	}
}
