// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package passwordpolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PasswordPolicyConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Name of the password policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.0.0/docs/resources/password_policy#name PasswordPolicy#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The password policy document.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.0.0/docs/resources/password_policy#policy PasswordPolicy#policy}
	Policy *string `field:"required" json:"policy" yaml:"policy"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.0.0/docs/resources/password_policy#namespace PasswordPolicy#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

