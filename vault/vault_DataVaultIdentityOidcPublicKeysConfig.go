// Prebuilt vault Provider for Terraform CDK (cdktf)
package vault

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataVaultIdentityOidcPublicKeysConfig struct {
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
	// The name of the provider.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/d/identity_oidc_public_keys#name DataVaultIdentityOidcPublicKeys#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/d/identity_oidc_public_keys#id DataVaultIdentityOidcPublicKeys#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/d/identity_oidc_public_keys#namespace DataVaultIdentityOidcPublicKeys#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

