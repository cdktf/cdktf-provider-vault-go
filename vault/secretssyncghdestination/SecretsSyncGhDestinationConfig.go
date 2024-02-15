// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package secretssyncghdestination

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SecretsSyncGhDestinationConfig struct {
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
	// Unique name of the github destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.25.0/docs/resources/secrets_sync_gh_destination#name SecretsSyncGhDestination#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Fine-grained or personal access token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.25.0/docs/resources/secrets_sync_gh_destination#access_token SecretsSyncGhDestination#access_token}
	AccessToken *string `field:"optional" json:"accessToken" yaml:"accessToken"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.25.0/docs/resources/secrets_sync_gh_destination#id SecretsSyncGhDestination#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.25.0/docs/resources/secrets_sync_gh_destination#namespace SecretsSyncGhDestination#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Name of the repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.25.0/docs/resources/secrets_sync_gh_destination#repository_name SecretsSyncGhDestination#repository_name}
	RepositoryName *string `field:"optional" json:"repositoryName" yaml:"repositoryName"`
	// GitHub organization or username that owns the repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.25.0/docs/resources/secrets_sync_gh_destination#repository_owner SecretsSyncGhDestination#repository_owner}
	RepositoryOwner *string `field:"optional" json:"repositoryOwner" yaml:"repositoryOwner"`
	// Template describing how to generate external secret names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.25.0/docs/resources/secrets_sync_gh_destination#secret_name_template SecretsSyncGhDestination#secret_name_template}
	SecretNameTemplate *string `field:"optional" json:"secretNameTemplate" yaml:"secretNameTemplate"`
}

