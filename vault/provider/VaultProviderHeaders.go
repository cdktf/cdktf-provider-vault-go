// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider


type VaultProviderHeaders struct {
	// The header name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.0.0/docs#name VaultProvider#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The header value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.0.0/docs#value VaultProvider#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

