// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pkisecretbackendconfigscep


type PkiSecretBackendConfigScepAuthenticators struct {
	// The accessor and cert_role properties for cert auth backends.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.2.1/docs/resources/pki_secret_backend_config_scep#cert PkiSecretBackendConfigScep#cert}
	Cert *map[string]*string `field:"optional" json:"cert" yaml:"cert"`
	// The accessor property for SCEP auth backends.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.2.1/docs/resources/pki_secret_backend_config_scep#scep PkiSecretBackendConfigScep#scep}
	Scep *map[string]*string `field:"optional" json:"scep" yaml:"scep"`
}

