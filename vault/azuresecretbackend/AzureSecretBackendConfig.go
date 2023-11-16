// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package azuresecretbackend

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AzureSecretBackendConfig struct {
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
	// The subscription id for the Azure Active Directory.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.23.0/docs/resources/azure_secret_backend#subscription_id AzureSecretBackend#subscription_id}
	SubscriptionId *string `field:"required" json:"subscriptionId" yaml:"subscriptionId"`
	// The tenant id for the Azure Active Directory organization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.23.0/docs/resources/azure_secret_backend#tenant_id AzureSecretBackend#tenant_id}
	TenantId *string `field:"required" json:"tenantId" yaml:"tenantId"`
	// The client id for credentials to query the Azure APIs. Currently read permissions to query compute resources are required.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.23.0/docs/resources/azure_secret_backend#client_id AzureSecretBackend#client_id}
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// The client secret for credentials to query the Azure APIs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.23.0/docs/resources/azure_secret_backend#client_secret AzureSecretBackend#client_secret}
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// Human-friendly description of the mount for the backend.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.23.0/docs/resources/azure_secret_backend#description AzureSecretBackend#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// If set, opts out of mount migration on path updates.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.23.0/docs/resources/azure_secret_backend#disable_remount AzureSecretBackend#disable_remount}
	DisableRemount interface{} `field:"optional" json:"disableRemount" yaml:"disableRemount"`
	// The Azure cloud environment. Valid values: AzurePublicCloud, AzureUSGovernmentCloud, AzureChinaCloud, AzureGermanCloud.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.23.0/docs/resources/azure_secret_backend#environment AzureSecretBackend#environment}
	Environment *string `field:"optional" json:"environment" yaml:"environment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.23.0/docs/resources/azure_secret_backend#id AzureSecretBackend#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.23.0/docs/resources/azure_secret_backend#namespace AzureSecretBackend#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Path to mount the backend at.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.23.0/docs/resources/azure_secret_backend#path AzureSecretBackend#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Use the Microsoft Graph API. Should be set to true on vault-1.10+.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.23.0/docs/resources/azure_secret_backend#use_microsoft_graph_api AzureSecretBackend#use_microsoft_graph_api}
	UseMicrosoftGraphApi interface{} `field:"optional" json:"useMicrosoftGraphApi" yaml:"useMicrosoftGraphApi"`
}

