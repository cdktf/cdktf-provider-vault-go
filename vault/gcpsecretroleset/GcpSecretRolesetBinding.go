// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gcpsecretroleset


type GcpSecretRolesetBinding struct {
	// Resource name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.2.0/docs/resources/gcp_secret_roleset#resource GcpSecretRoleset#resource}
	Resource *string `field:"required" json:"resource" yaml:"resource"`
	// List of roles to apply to the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.2.0/docs/resources/gcp_secret_roleset#roles GcpSecretRoleset#roles}
	Roles *[]*string `field:"required" json:"roles" yaml:"roles"`
}

