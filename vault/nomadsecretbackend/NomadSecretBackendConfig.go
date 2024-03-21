// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package nomadsecretbackend

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NomadSecretBackendConfig struct {
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
	// Specifies the address of the Nomad instance, provided as "protocol://host:port" like "http://127.0.0.1:4646".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.1.0/docs/resources/nomad_secret_backend#address NomadSecretBackend#address}
	Address *string `field:"optional" json:"address" yaml:"address"`
	// The mount path for the Nomad backend.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.1.0/docs/resources/nomad_secret_backend#backend NomadSecretBackend#backend}
	Backend *string `field:"optional" json:"backend" yaml:"backend"`
	// CA certificate to use when verifying Nomad server certificate, must be x509 PEM encoded.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.1.0/docs/resources/nomad_secret_backend#ca_cert NomadSecretBackend#ca_cert}
	CaCert *string `field:"optional" json:"caCert" yaml:"caCert"`
	// Client certificate used for Nomad's TLS communication, must be x509 PEM encoded and if this is set you need to also set client_key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.1.0/docs/resources/nomad_secret_backend#client_cert NomadSecretBackend#client_cert}
	ClientCert *string `field:"optional" json:"clientCert" yaml:"clientCert"`
	// Client key used for Nomad's TLS communication, must be x509 PEM encoded and if this is set you need to also set client_cert.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.1.0/docs/resources/nomad_secret_backend#client_key NomadSecretBackend#client_key}
	ClientKey *string `field:"optional" json:"clientKey" yaml:"clientKey"`
	// Default lease duration for secrets in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.1.0/docs/resources/nomad_secret_backend#default_lease_ttl_seconds NomadSecretBackend#default_lease_ttl_seconds}
	DefaultLeaseTtlSeconds *float64 `field:"optional" json:"defaultLeaseTtlSeconds" yaml:"defaultLeaseTtlSeconds"`
	// Human-friendly description of the mount for the backend.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.1.0/docs/resources/nomad_secret_backend#description NomadSecretBackend#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// If set, opts out of mount migration on path updates.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.1.0/docs/resources/nomad_secret_backend#disable_remount NomadSecretBackend#disable_remount}
	DisableRemount interface{} `field:"optional" json:"disableRemount" yaml:"disableRemount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.1.0/docs/resources/nomad_secret_backend#id NomadSecretBackend#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Mark the secrets engine as local-only.
	//
	// Local engines are not replicated or removed by replication. Tolerance duration to use when checking the last rotation time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.1.0/docs/resources/nomad_secret_backend#local NomadSecretBackend#local}
	Local interface{} `field:"optional" json:"local" yaml:"local"`
	// Maximum possible lease duration for secrets in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.1.0/docs/resources/nomad_secret_backend#max_lease_ttl_seconds NomadSecretBackend#max_lease_ttl_seconds}
	MaxLeaseTtlSeconds *float64 `field:"optional" json:"maxLeaseTtlSeconds" yaml:"maxLeaseTtlSeconds"`
	// Specifies the maximum length to use for the name of the Nomad token generated with Generate Credential.
	//
	// If omitted, 0 is used and ignored, defaulting to the max value allowed by the Nomad version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.1.0/docs/resources/nomad_secret_backend#max_token_name_length NomadSecretBackend#max_token_name_length}
	MaxTokenNameLength *float64 `field:"optional" json:"maxTokenNameLength" yaml:"maxTokenNameLength"`
	// Maximum possible lease duration for secrets in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.1.0/docs/resources/nomad_secret_backend#max_ttl NomadSecretBackend#max_ttl}
	MaxTtl *float64 `field:"optional" json:"maxTtl" yaml:"maxTtl"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.1.0/docs/resources/nomad_secret_backend#namespace NomadSecretBackend#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Specifies the Nomad Management token to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.1.0/docs/resources/nomad_secret_backend#token NomadSecretBackend#token}
	Token *string `field:"optional" json:"token" yaml:"token"`
	// Maximum possible lease duration for secrets in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.1.0/docs/resources/nomad_secret_backend#ttl NomadSecretBackend#ttl}
	Ttl *float64 `field:"optional" json:"ttl" yaml:"ttl"`
}

