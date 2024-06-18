// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider


type VaultProviderClientAuth struct {
	// Path to a file containing the client certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.3.0/docs#cert_file VaultProvider#cert_file}
	CertFile *string `field:"optional" json:"certFile" yaml:"certFile"`
	// Path to a file containing the private key that the certificate was issued for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.3.0/docs#key_file VaultProvider#key_file}
	KeyFile *string `field:"optional" json:"keyFile" yaml:"keyFile"`
}

