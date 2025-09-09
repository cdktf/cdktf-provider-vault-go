// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pkisecretbackendconfigscep


type PkiSecretBackendConfigScepExternalValidation struct {
	// The credentials to enable Microsoft Intune validation of SCEP requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.3.0/docs/resources/pki_secret_backend_config_scep#intune PkiSecretBackendConfigScep#intune}
	Intune *map[string]*string `field:"optional" json:"intune" yaml:"intune"`
}

