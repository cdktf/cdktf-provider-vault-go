// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gcpauthbackend

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GcpAuthBackendConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.0.0/docs/resources/gcp_auth_backend#client_email GcpAuthBackend#client_email}.
	ClientEmail *string `field:"optional" json:"clientEmail" yaml:"clientEmail"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.0.0/docs/resources/gcp_auth_backend#client_id GcpAuthBackend#client_id}.
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.0.0/docs/resources/gcp_auth_backend#credentials GcpAuthBackend#credentials}.
	Credentials *string `field:"optional" json:"credentials" yaml:"credentials"`
	// custom_endpoint block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.0.0/docs/resources/gcp_auth_backend#custom_endpoint GcpAuthBackend#custom_endpoint}
	CustomEndpoint *GcpAuthBackendCustomEndpoint `field:"optional" json:"customEndpoint" yaml:"customEndpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.0.0/docs/resources/gcp_auth_backend#description GcpAuthBackend#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Stops rotation of the root credential until set to false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.0.0/docs/resources/gcp_auth_backend#disable_automated_rotation GcpAuthBackend#disable_automated_rotation}
	DisableAutomatedRotation interface{} `field:"optional" json:"disableAutomatedRotation" yaml:"disableAutomatedRotation"`
	// If set, opts out of mount migration on path updates.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.0.0/docs/resources/gcp_auth_backend#disable_remount GcpAuthBackend#disable_remount}
	DisableRemount interface{} `field:"optional" json:"disableRemount" yaml:"disableRemount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.0.0/docs/resources/gcp_auth_backend#id GcpAuthBackend#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The audience claim value for plugin identity tokens.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.0.0/docs/resources/gcp_auth_backend#identity_token_audience GcpAuthBackend#identity_token_audience}
	IdentityTokenAudience *string `field:"optional" json:"identityTokenAudience" yaml:"identityTokenAudience"`
	// The key to use for signing identity tokens.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.0.0/docs/resources/gcp_auth_backend#identity_token_key GcpAuthBackend#identity_token_key}
	IdentityTokenKey *string `field:"optional" json:"identityTokenKey" yaml:"identityTokenKey"`
	// The TTL of generated tokens.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.0.0/docs/resources/gcp_auth_backend#identity_token_ttl GcpAuthBackend#identity_token_ttl}
	IdentityTokenTtl *float64 `field:"optional" json:"identityTokenTtl" yaml:"identityTokenTtl"`
	// Specifies if the auth method is local only.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.0.0/docs/resources/gcp_auth_backend#local GcpAuthBackend#local}
	Local interface{} `field:"optional" json:"local" yaml:"local"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.0.0/docs/resources/gcp_auth_backend#namespace GcpAuthBackend#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.0.0/docs/resources/gcp_auth_backend#path GcpAuthBackend#path}.
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.0.0/docs/resources/gcp_auth_backend#private_key_id GcpAuthBackend#private_key_id}.
	PrivateKeyId *string `field:"optional" json:"privateKeyId" yaml:"privateKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.0.0/docs/resources/gcp_auth_backend#project_id GcpAuthBackend#project_id}.
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
	// The period of time in seconds between each rotation of the root credential. Cannot be used with rotation_schedule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.0.0/docs/resources/gcp_auth_backend#rotation_period GcpAuthBackend#rotation_period}
	RotationPeriod *float64 `field:"optional" json:"rotationPeriod" yaml:"rotationPeriod"`
	// The cron-style schedule for the root credential to be rotated on. Cannot be used with rotation_period.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.0.0/docs/resources/gcp_auth_backend#rotation_schedule GcpAuthBackend#rotation_schedule}
	RotationSchedule *string `field:"optional" json:"rotationSchedule" yaml:"rotationSchedule"`
	// The maximum amount of time in seconds Vault is allowed to complete a rotation once a scheduled rotation is triggered.
	//
	// Can only be used with rotation_schedule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.0.0/docs/resources/gcp_auth_backend#rotation_window GcpAuthBackend#rotation_window}
	RotationWindow *float64 `field:"optional" json:"rotationWindow" yaml:"rotationWindow"`
	// Service Account to impersonate for plugin workload identity federation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.0.0/docs/resources/gcp_auth_backend#service_account_email GcpAuthBackend#service_account_email}
	ServiceAccountEmail *string `field:"optional" json:"serviceAccountEmail" yaml:"serviceAccountEmail"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.0.0/docs/resources/gcp_auth_backend#tune GcpAuthBackend#tune}.
	Tune interface{} `field:"optional" json:"tune" yaml:"tune"`
}

