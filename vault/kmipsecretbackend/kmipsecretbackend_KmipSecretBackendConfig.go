package kmipsecretbackend

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KmipSecretBackendConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
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
	// Path where KMIP secret backend will be mounted.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/kmip_secret_backend#path KmipSecretBackend#path}
	Path *string `field:"required" json:"path" yaml:"path"`
	// Client certificate key bits, valid values depend on key type.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/kmip_secret_backend#default_tls_client_key_bits KmipSecretBackend#default_tls_client_key_bits}
	DefaultTlsClientKeyBits *float64 `field:"optional" json:"defaultTlsClientKeyBits" yaml:"defaultTlsClientKeyBits"`
	// Client certificate key type, rsa or ec.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/kmip_secret_backend#default_tls_client_key_type KmipSecretBackend#default_tls_client_key_type}
	DefaultTlsClientKeyType *string `field:"optional" json:"defaultTlsClientKeyType" yaml:"defaultTlsClientKeyType"`
	// Client certificate TTL in seconds.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/kmip_secret_backend#default_tls_client_ttl KmipSecretBackend#default_tls_client_ttl}
	DefaultTlsClientTtl *float64 `field:"optional" json:"defaultTlsClientTtl" yaml:"defaultTlsClientTtl"`
	// Human-friendly description of the mount for the backend.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/kmip_secret_backend#description KmipSecretBackend#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/kmip_secret_backend#id KmipSecretBackend#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Addresses the KMIP server should listen on (host:port).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/kmip_secret_backend#listen_addrs KmipSecretBackend#listen_addrs}
	ListenAddrs *[]*string `field:"optional" json:"listenAddrs" yaml:"listenAddrs"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/kmip_secret_backend#namespace KmipSecretBackend#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Hostnames to include in the server's TLS certificate as SAN DNS names.
	//
	// The first will be used as the common name (CN)
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/kmip_secret_backend#server_hostnames KmipSecretBackend#server_hostnames}
	ServerHostnames *[]*string `field:"optional" json:"serverHostnames" yaml:"serverHostnames"`
	// IPs to include in the server's TLS certificate as SAN IP addresses.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/kmip_secret_backend#server_ips KmipSecretBackend#server_ips}
	ServerIps *[]*string `field:"optional" json:"serverIps" yaml:"serverIps"`
	// CA key bits, valid values depend on key type.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/kmip_secret_backend#tls_ca_key_bits KmipSecretBackend#tls_ca_key_bits}
	TlsCaKeyBits *float64 `field:"optional" json:"tlsCaKeyBits" yaml:"tlsCaKeyBits"`
	// CA key type, rsa or ec.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/kmip_secret_backend#tls_ca_key_type KmipSecretBackend#tls_ca_key_type}
	TlsCaKeyType *string `field:"optional" json:"tlsCaKeyType" yaml:"tlsCaKeyType"`
	// Minimum TLS version to accept.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/kmip_secret_backend#tls_min_version KmipSecretBackend#tls_min_version}
	TlsMinVersion *string `field:"optional" json:"tlsMinVersion" yaml:"tlsMinVersion"`
}

