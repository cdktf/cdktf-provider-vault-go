// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package terraformcloudsecretrole

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TerraformCloudSecretRoleConfig struct {
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
	// The name of an existing role against which to create this Terraform Cloud credential.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.4.0/docs/resources/terraform_cloud_secret_role#name TerraformCloudSecretRole#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The path of the Terraform Cloud Secret Backend the role belongs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.4.0/docs/resources/terraform_cloud_secret_role#backend TerraformCloudSecretRole#backend}
	Backend *string `field:"optional" json:"backend" yaml:"backend"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.4.0/docs/resources/terraform_cloud_secret_role#id TerraformCloudSecretRole#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Maximum allowed lease for generated credentials. If not set or set to 0, will use system default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.4.0/docs/resources/terraform_cloud_secret_role#max_ttl TerraformCloudSecretRole#max_ttl}
	MaxTtl *float64 `field:"optional" json:"maxTtl" yaml:"maxTtl"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.4.0/docs/resources/terraform_cloud_secret_role#namespace TerraformCloudSecretRole#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Name of the Terraform Cloud or Enterprise organization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.4.0/docs/resources/terraform_cloud_secret_role#organization TerraformCloudSecretRole#organization}
	Organization *string `field:"optional" json:"organization" yaml:"organization"`
	// ID of the Terraform Cloud or Enterprise team under organization (e.g., settings/teams/team-xxxxxxxxxxxxx).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.4.0/docs/resources/terraform_cloud_secret_role#team_id TerraformCloudSecretRole#team_id}
	TeamId *string `field:"optional" json:"teamId" yaml:"teamId"`
	// Default lease for generated credentials. If not set or set to 0, will use system default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.4.0/docs/resources/terraform_cloud_secret_role#ttl TerraformCloudSecretRole#ttl}
	Ttl *float64 `field:"optional" json:"ttl" yaml:"ttl"`
	// ID of the Terraform Cloud or Enterprise user (e.g., user-xxxxxxxxxxxxxxxx).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.4.0/docs/resources/terraform_cloud_secret_role#user_id TerraformCloudSecretRole#user_id}
	UserId *string `field:"optional" json:"userId" yaml:"userId"`
}

