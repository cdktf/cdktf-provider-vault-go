// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package terraformcloudsecretbackend

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TerraformCloudSecretBackendConfig struct {
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
	// Specifies the address of the Terraform Cloud instance, provided as "host:port" like "127.0.0.1:8500".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.6.0/docs/resources/terraform_cloud_secret_backend#address TerraformCloudSecretBackend#address}
	Address *string `field:"optional" json:"address" yaml:"address"`
	// Unique name of the Vault Terraform Cloud mount to configure.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.6.0/docs/resources/terraform_cloud_secret_backend#backend TerraformCloudSecretBackend#backend}
	Backend *string `field:"optional" json:"backend" yaml:"backend"`
	// Specifies the base path for the Terraform Cloud or Enterprise API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.6.0/docs/resources/terraform_cloud_secret_backend#base_path TerraformCloudSecretBackend#base_path}
	BasePath *string `field:"optional" json:"basePath" yaml:"basePath"`
	// Default lease duration for secrets in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.6.0/docs/resources/terraform_cloud_secret_backend#default_lease_ttl_seconds TerraformCloudSecretBackend#default_lease_ttl_seconds}
	DefaultLeaseTtlSeconds *float64 `field:"optional" json:"defaultLeaseTtlSeconds" yaml:"defaultLeaseTtlSeconds"`
	// Human-friendly description of the mount for the backend.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.6.0/docs/resources/terraform_cloud_secret_backend#description TerraformCloudSecretBackend#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// If set, opts out of mount migration on path updates.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.6.0/docs/resources/terraform_cloud_secret_backend#disable_remount TerraformCloudSecretBackend#disable_remount}
	DisableRemount interface{} `field:"optional" json:"disableRemount" yaml:"disableRemount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.6.0/docs/resources/terraform_cloud_secret_backend#id TerraformCloudSecretBackend#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Maximum possible lease duration for secrets in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.6.0/docs/resources/terraform_cloud_secret_backend#max_lease_ttl_seconds TerraformCloudSecretBackend#max_lease_ttl_seconds}
	MaxLeaseTtlSeconds *float64 `field:"optional" json:"maxLeaseTtlSeconds" yaml:"maxLeaseTtlSeconds"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.6.0/docs/resources/terraform_cloud_secret_backend#namespace TerraformCloudSecretBackend#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Specifies the Terraform Cloud access token to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.6.0/docs/resources/terraform_cloud_secret_backend#token TerraformCloudSecretBackend#token}
	Token *string `field:"optional" json:"token" yaml:"token"`
}

