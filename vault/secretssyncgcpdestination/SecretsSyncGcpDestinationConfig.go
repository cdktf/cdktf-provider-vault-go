// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package secretssyncgcpdestination

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SecretsSyncGcpDestinationConfig struct {
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
	// Unique name of the GCP destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.25.0/docs/resources/secrets_sync_gcp_destination#name SecretsSyncGcpDestination#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// JSON-encoded credentials to use to connect to GCP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.25.0/docs/resources/secrets_sync_gcp_destination#credentials SecretsSyncGcpDestination#credentials}
	Credentials *string `field:"optional" json:"credentials" yaml:"credentials"`
	// Custom tags to set on the secret managed at the destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.25.0/docs/resources/secrets_sync_gcp_destination#custom_tags SecretsSyncGcpDestination#custom_tags}
	CustomTags *map[string]*string `field:"optional" json:"customTags" yaml:"customTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.25.0/docs/resources/secrets_sync_gcp_destination#id SecretsSyncGcpDestination#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.25.0/docs/resources/secrets_sync_gcp_destination#namespace SecretsSyncGcpDestination#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Template describing how to generate external secret names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.25.0/docs/resources/secrets_sync_gcp_destination#secret_name_template SecretsSyncGcpDestination#secret_name_template}
	SecretNameTemplate *string `field:"optional" json:"secretNameTemplate" yaml:"secretNameTemplate"`
}

