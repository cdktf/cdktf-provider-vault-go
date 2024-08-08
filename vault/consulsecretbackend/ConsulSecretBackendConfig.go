// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package consulsecretbackend

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ConsulSecretBackendConfig struct {
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
	// Specifies the address of the Consul instance, provided as "host:port" like "127.0.0.1:8500".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.4.0/docs/resources/consul_secret_backend#address ConsulSecretBackend#address}
	Address *string `field:"required" json:"address" yaml:"address"`
	// Denotes a backend resource that is used to bootstrap the Consul ACL system.
	//
	// Only one resource may be used to bootstrap.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.4.0/docs/resources/consul_secret_backend#bootstrap ConsulSecretBackend#bootstrap}
	Bootstrap interface{} `field:"optional" json:"bootstrap" yaml:"bootstrap"`
	// CA certificate to use when verifying Consul server certificate, must be x509 PEM encoded.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.4.0/docs/resources/consul_secret_backend#ca_cert ConsulSecretBackend#ca_cert}
	CaCert *string `field:"optional" json:"caCert" yaml:"caCert"`
	// Client certificate used for Consul's TLS communication, must be x509 PEM encoded and if this is set you need to also set client_key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.4.0/docs/resources/consul_secret_backend#client_cert ConsulSecretBackend#client_cert}
	ClientCert *string `field:"optional" json:"clientCert" yaml:"clientCert"`
	// Client key used for Consul's TLS communication, must be x509 PEM encoded and if this is set you need to also set client_cert.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.4.0/docs/resources/consul_secret_backend#client_key ConsulSecretBackend#client_key}
	ClientKey *string `field:"optional" json:"clientKey" yaml:"clientKey"`
	// Default lease duration for secrets in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.4.0/docs/resources/consul_secret_backend#default_lease_ttl_seconds ConsulSecretBackend#default_lease_ttl_seconds}
	DefaultLeaseTtlSeconds *float64 `field:"optional" json:"defaultLeaseTtlSeconds" yaml:"defaultLeaseTtlSeconds"`
	// Human-friendly description of the mount for the backend.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.4.0/docs/resources/consul_secret_backend#description ConsulSecretBackend#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// If set, opts out of mount migration on path updates.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.4.0/docs/resources/consul_secret_backend#disable_remount ConsulSecretBackend#disable_remount}
	DisableRemount interface{} `field:"optional" json:"disableRemount" yaml:"disableRemount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.4.0/docs/resources/consul_secret_backend#id ConsulSecretBackend#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Specifies if the secret backend is local only.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.4.0/docs/resources/consul_secret_backend#local ConsulSecretBackend#local}
	Local interface{} `field:"optional" json:"local" yaml:"local"`
	// Maximum possible lease duration for secrets in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.4.0/docs/resources/consul_secret_backend#max_lease_ttl_seconds ConsulSecretBackend#max_lease_ttl_seconds}
	MaxLeaseTtlSeconds *float64 `field:"optional" json:"maxLeaseTtlSeconds" yaml:"maxLeaseTtlSeconds"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.4.0/docs/resources/consul_secret_backend#namespace ConsulSecretBackend#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Unique name of the Vault Consul mount to configure.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.4.0/docs/resources/consul_secret_backend#path ConsulSecretBackend#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Specifies the URL scheme to use. Defaults to "http".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.4.0/docs/resources/consul_secret_backend#scheme ConsulSecretBackend#scheme}
	Scheme *string `field:"optional" json:"scheme" yaml:"scheme"`
	// Specifies the Consul token to use when managing or issuing new tokens.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.4.0/docs/resources/consul_secret_backend#token ConsulSecretBackend#token}
	Token *string `field:"optional" json:"token" yaml:"token"`
}

