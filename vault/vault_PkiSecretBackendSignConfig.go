// Prebuilt vault Provider for Terraform CDK (cdktf)
package vault

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PkiSecretBackendSignConfig struct {
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
	// The PKI secret backend the resource belongs to.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/pki_secret_backend_sign#backend PkiSecretBackendSign#backend}
	Backend *string `field:"required" json:"backend" yaml:"backend"`
	// CN of intermediate to create.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/pki_secret_backend_sign#common_name PkiSecretBackendSign#common_name}
	CommonName *string `field:"required" json:"commonName" yaml:"commonName"`
	// The CSR.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/pki_secret_backend_sign#csr PkiSecretBackendSign#csr}
	Csr *string `field:"required" json:"csr" yaml:"csr"`
	// Name of the role to create the certificate against.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/pki_secret_backend_sign#name PkiSecretBackendSign#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// List of alternative names.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/pki_secret_backend_sign#alt_names PkiSecretBackendSign#alt_names}
	AltNames *[]*string `field:"optional" json:"altNames" yaml:"altNames"`
	// If enabled, a new certificate will be generated if the expiration is within min_seconds_remaining.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/pki_secret_backend_sign#auto_renew PkiSecretBackendSign#auto_renew}
	AutoRenew interface{} `field:"optional" json:"autoRenew" yaml:"autoRenew"`
	// Flag to exclude CN from SANs.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/pki_secret_backend_sign#exclude_cn_from_sans PkiSecretBackendSign#exclude_cn_from_sans}
	ExcludeCnFromSans interface{} `field:"optional" json:"excludeCnFromSans" yaml:"excludeCnFromSans"`
	// The format of data.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/pki_secret_backend_sign#format PkiSecretBackendSign#format}
	Format *string `field:"optional" json:"format" yaml:"format"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/pki_secret_backend_sign#id PkiSecretBackendSign#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// List of alternative IPs.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/pki_secret_backend_sign#ip_sans PkiSecretBackendSign#ip_sans}
	IpSans *[]*string `field:"optional" json:"ipSans" yaml:"ipSans"`
	// Generate a new certificate when the expiration is within this number of seconds.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/pki_secret_backend_sign#min_seconds_remaining PkiSecretBackendSign#min_seconds_remaining}
	MinSecondsRemaining *float64 `field:"optional" json:"minSecondsRemaining" yaml:"minSecondsRemaining"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/pki_secret_backend_sign#namespace PkiSecretBackendSign#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// List of other SANs.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/pki_secret_backend_sign#other_sans PkiSecretBackendSign#other_sans}
	OtherSans *[]*string `field:"optional" json:"otherSans" yaml:"otherSans"`
	// Time to live.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/pki_secret_backend_sign#ttl PkiSecretBackendSign#ttl}
	Ttl *string `field:"optional" json:"ttl" yaml:"ttl"`
	// List of alternative URIs.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/pki_secret_backend_sign#uri_sans PkiSecretBackendSign#uri_sans}
	UriSans *[]*string `field:"optional" json:"uriSans" yaml:"uriSans"`
}

