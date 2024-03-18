// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package azuresecretbackendrole


type AzureSecretBackendRoleAzureGroups struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.0.0/docs/resources/azure_secret_backend_role#group_name AzureSecretBackendRole#group_name}.
	GroupName *string `field:"required" json:"groupName" yaml:"groupName"`
}

