// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package secretssyncverceldestination

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SecretsSyncVercelDestinationConfig struct {
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
	// Vercel API access token with the permissions to manage environment variables.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.3.0/docs/resources/secrets_sync_vercel_destination#access_token SecretsSyncVercelDestination#access_token}
	AccessToken *string `field:"required" json:"accessToken" yaml:"accessToken"`
	// Deployment environments where the environment variables are available. Accepts 'development', 'preview' & 'production'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.3.0/docs/resources/secrets_sync_vercel_destination#deployment_environments SecretsSyncVercelDestination#deployment_environments}
	DeploymentEnvironments *[]*string `field:"required" json:"deploymentEnvironments" yaml:"deploymentEnvironments"`
	// Unique name of the Vercel destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.3.0/docs/resources/secrets_sync_vercel_destination#name SecretsSyncVercelDestination#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Project ID where to manage environment variables.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.3.0/docs/resources/secrets_sync_vercel_destination#project_id SecretsSyncVercelDestination#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Determines what level of information is synced as a distinct resource at the destination. Can be 'secret-path' or 'secret-key'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.3.0/docs/resources/secrets_sync_vercel_destination#granularity SecretsSyncVercelDestination#granularity}
	Granularity *string `field:"optional" json:"granularity" yaml:"granularity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.3.0/docs/resources/secrets_sync_vercel_destination#id SecretsSyncVercelDestination#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.3.0/docs/resources/secrets_sync_vercel_destination#namespace SecretsSyncVercelDestination#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Template describing how to generate external secret names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.3.0/docs/resources/secrets_sync_vercel_destination#secret_name_template SecretsSyncVercelDestination#secret_name_template}
	SecretNameTemplate *string `field:"optional" json:"secretNameTemplate" yaml:"secretNameTemplate"`
	// Team ID the project belongs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.3.0/docs/resources/secrets_sync_vercel_destination#team_id SecretsSyncVercelDestination#team_id}
	TeamId *string `field:"optional" json:"teamId" yaml:"teamId"`
}

