package pkisecretbackendcrlconfig

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PkiSecretBackendCrlConfigConfig struct {
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
	// The path of the PKI secret backend the resource belongs to.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/pki_secret_backend_crl_config#backend PkiSecretBackendCrlConfig#backend}
	Backend *string `field:"required" json:"backend" yaml:"backend"`
	// Disables or enables CRL building.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/pki_secret_backend_crl_config#disable PkiSecretBackendCrlConfig#disable}
	Disable interface{} `field:"optional" json:"disable" yaml:"disable"`
	// Specifies the time until expiration.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/pki_secret_backend_crl_config#expiry PkiSecretBackendCrlConfig#expiry}
	Expiry *string `field:"optional" json:"expiry" yaml:"expiry"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/pki_secret_backend_crl_config#id PkiSecretBackendCrlConfig#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/pki_secret_backend_crl_config#namespace PkiSecretBackendCrlConfig#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

