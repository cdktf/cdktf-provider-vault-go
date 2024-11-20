// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package egppolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EgpPolicyConfig struct {
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
	// Enforcement level of Sentinel policy. Can be one of: 'advisory', 'soft-mandatory' or 'hard-mandatory'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.5.0/docs/resources/egp_policy#enforcement_level EgpPolicy#enforcement_level}
	EnforcementLevel *string `field:"required" json:"enforcementLevel" yaml:"enforcementLevel"`
	// Name of the policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.5.0/docs/resources/egp_policy#name EgpPolicy#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// List of paths to which the policy will be applied.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.5.0/docs/resources/egp_policy#paths EgpPolicy#paths}
	Paths *[]*string `field:"required" json:"paths" yaml:"paths"`
	// The policy document.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.5.0/docs/resources/egp_policy#policy EgpPolicy#policy}
	Policy *string `field:"required" json:"policy" yaml:"policy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.5.0/docs/resources/egp_policy#id EgpPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.5.0/docs/resources/egp_policy#namespace EgpPolicy#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

