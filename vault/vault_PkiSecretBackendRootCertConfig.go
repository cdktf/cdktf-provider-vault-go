// Prebuilt vault Provider for Terraform CDK (cdktf)
package vault

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PkiSecretBackendRootCertConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/pki_secret_backend_root_cert#backend PkiSecretBackendRootCert#backend}
	Backend *string `field:"required" json:"backend" yaml:"backend"`
	// CN of root to create.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/pki_secret_backend_root_cert#common_name PkiSecretBackendRootCert#common_name}
	CommonName *string `field:"required" json:"commonName" yaml:"commonName"`
	// Type of root to create. Must be either "exported" or "internal".
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/pki_secret_backend_root_cert#type PkiSecretBackendRootCert#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// List of alternative names.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/pki_secret_backend_root_cert#alt_names PkiSecretBackendRootCert#alt_names}
	AltNames *[]*string `field:"optional" json:"altNames" yaml:"altNames"`
	// The country.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/pki_secret_backend_root_cert#country PkiSecretBackendRootCert#country}
	Country *string `field:"optional" json:"country" yaml:"country"`
	// Flag to exclude CN from SANs.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/pki_secret_backend_root_cert#exclude_cn_from_sans PkiSecretBackendRootCert#exclude_cn_from_sans}
	ExcludeCnFromSans interface{} `field:"optional" json:"excludeCnFromSans" yaml:"excludeCnFromSans"`
	// The format of data.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/pki_secret_backend_root_cert#format PkiSecretBackendRootCert#format}
	Format *string `field:"optional" json:"format" yaml:"format"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/pki_secret_backend_root_cert#id PkiSecretBackendRootCert#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// List of alternative IPs.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/pki_secret_backend_root_cert#ip_sans PkiSecretBackendRootCert#ip_sans}
	IpSans *[]*string `field:"optional" json:"ipSans" yaml:"ipSans"`
	// The number of bits to use.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/pki_secret_backend_root_cert#key_bits PkiSecretBackendRootCert#key_bits}
	KeyBits *float64 `field:"optional" json:"keyBits" yaml:"keyBits"`
	// The desired key type.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/pki_secret_backend_root_cert#key_type PkiSecretBackendRootCert#key_type}
	KeyType *string `field:"optional" json:"keyType" yaml:"keyType"`
	// The locality.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/pki_secret_backend_root_cert#locality PkiSecretBackendRootCert#locality}
	Locality *string `field:"optional" json:"locality" yaml:"locality"`
	// The maximum path length to encode in the generated certificate.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/pki_secret_backend_root_cert#max_path_length PkiSecretBackendRootCert#max_path_length}
	MaxPathLength *float64 `field:"optional" json:"maxPathLength" yaml:"maxPathLength"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/pki_secret_backend_root_cert#namespace PkiSecretBackendRootCert#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// The organization.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/pki_secret_backend_root_cert#organization PkiSecretBackendRootCert#organization}
	Organization *string `field:"optional" json:"organization" yaml:"organization"`
	// List of other SANs.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/pki_secret_backend_root_cert#other_sans PkiSecretBackendRootCert#other_sans}
	OtherSans *[]*string `field:"optional" json:"otherSans" yaml:"otherSans"`
	// The organization unit.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/pki_secret_backend_root_cert#ou PkiSecretBackendRootCert#ou}
	Ou *string `field:"optional" json:"ou" yaml:"ou"`
	// List of domains for which certificates are allowed to be issued.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/pki_secret_backend_root_cert#permitted_dns_domains PkiSecretBackendRootCert#permitted_dns_domains}
	PermittedDnsDomains *[]*string `field:"optional" json:"permittedDnsDomains" yaml:"permittedDnsDomains"`
	// The postal code.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/pki_secret_backend_root_cert#postal_code PkiSecretBackendRootCert#postal_code}
	PostalCode *string `field:"optional" json:"postalCode" yaml:"postalCode"`
	// The private key format.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/pki_secret_backend_root_cert#private_key_format PkiSecretBackendRootCert#private_key_format}
	PrivateKeyFormat *string `field:"optional" json:"privateKeyFormat" yaml:"privateKeyFormat"`
	// The province.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/pki_secret_backend_root_cert#province PkiSecretBackendRootCert#province}
	Province *string `field:"optional" json:"province" yaml:"province"`
	// The street address.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/pki_secret_backend_root_cert#street_address PkiSecretBackendRootCert#street_address}
	StreetAddress *string `field:"optional" json:"streetAddress" yaml:"streetAddress"`
	// Time to live.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/pki_secret_backend_root_cert#ttl PkiSecretBackendRootCert#ttl}
	Ttl *string `field:"optional" json:"ttl" yaml:"ttl"`
	// List of alternative URIs.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/pki_secret_backend_root_cert#uri_sans PkiSecretBackendRootCert#uri_sans}
	UriSans *[]*string `field:"optional" json:"uriSans" yaml:"uriSans"`
}

