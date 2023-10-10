// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider


type VaultProviderAuthLoginTokenFile struct {
	// The name of a file containing a single line that is a valid Vault token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.21.0/docs#filename VaultProvider#filename}
	Filename *string `field:"required" json:"filename" yaml:"filename"`
	// The authentication engine's namespace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.21.0/docs#namespace VaultProvider#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

