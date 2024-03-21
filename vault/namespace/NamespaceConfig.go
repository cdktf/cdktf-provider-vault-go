// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package namespace

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NamespaceConfig struct {
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
	// Namespace path.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.1.0/docs/resources/namespace#path Namespace#path}
	Path *string `field:"required" json:"path" yaml:"path"`
	// Custom metadata describing this namespace. Value type is map[string]string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.1.0/docs/resources/namespace#custom_metadata Namespace#custom_metadata}
	CustomMetadata *map[string]*string `field:"optional" json:"customMetadata" yaml:"customMetadata"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.1.0/docs/resources/namespace#id Namespace#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.1.0/docs/resources/namespace#namespace Namespace#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// The fully qualified namespace path.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.1.0/docs/resources/namespace#path_fq Namespace#path_fq}
	PathFq *string `field:"optional" json:"pathFq" yaml:"pathFq"`
}

