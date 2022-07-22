// SPDX-License-Identifier: BSL-1.0

package innosetup

import "github.com/moderncircuits/paket"

type InstallerConfig struct {
	UUID       string           `hcl:"uuid,optional" json:"uuid,omitempty"`
	Welcome    string           `hcl:"welcome,optional" json:"welcome,omitempty"`
	Conclusion string           `hcl:"conclusion,optional" json:"conclusion,omitempty"`
	Artifacts  []paket.Artifact `hcl:"artifact,block" json:"artifacts"`
}
