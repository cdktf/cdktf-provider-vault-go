// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datavaultpolicydocument


type DataVaultPolicyDocumentRuleAllowedParameter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.20.1/docs/data-sources/policy_document#key DataVaultPolicyDocument#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.20.1/docs/data-sources/policy_document#value DataVaultPolicyDocument#value}.
	Value *[]*string `field:"required" json:"value" yaml:"value"`
}

