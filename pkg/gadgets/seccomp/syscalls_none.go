// +build docs

// Copyright 2019-2021 The Inspektor Gadget authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package seccomp

import (
	"github.com/opencontainers/runtime-spec/specs-go"
	seccompprofilev1alpha1 "sigs.k8s.io/security-profiles-operator/api/seccompprofile/v1alpha1"
)

func syscallArrToLinuxSeccomp(v []byte) *specs.LinuxSeccomp {
	panic("Not implemented")
	return nil
}
func syscallArrToSeccompPolicy(namespace, name, generateName string, v []byte) *seccompprofilev1alpha1.SeccompProfile {
	panic("Not implemented")
	return nil
}