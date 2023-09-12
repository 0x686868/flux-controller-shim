/*
Copyright 2023 Hidde Beydals <yelling@hhh.computer>

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

package shim

import (
	"crypto/tls"

	"helm.sh/helm/v3/pkg/registry"

	intregistry "github.com/fluxcd/source-controller/internal/helm/registry"
)

// RegistryClientGenerator generates a registry client and a temporary
// credential file.
func RegistryClientGenerator(tlsConfig *tls.Config, isLogin bool) (*registry.Client, string, error) {
	return intregistry.ClientGenerator(tlsConfig, isLogin)
}
