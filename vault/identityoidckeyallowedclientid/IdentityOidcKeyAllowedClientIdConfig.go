package identityoidckeyallowedclientid

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IdentityOidcKeyAllowedClientIdConfig struct {
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
	// Role Client ID allowed to use the key for signing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.18.0/docs/resources/identity_oidc_key_allowed_client_id#allowed_client_id IdentityOidcKeyAllowedClientId#allowed_client_id}
	AllowedClientId *string `field:"required" json:"allowedClientId" yaml:"allowedClientId"`
	// Name of the key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.18.0/docs/resources/identity_oidc_key_allowed_client_id#key_name IdentityOidcKeyAllowedClientId#key_name}
	KeyName *string `field:"required" json:"keyName" yaml:"keyName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.18.0/docs/resources/identity_oidc_key_allowed_client_id#id IdentityOidcKeyAllowedClientId#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.18.0/docs/resources/identity_oidc_key_allowed_client_id#namespace IdentityOidcKeyAllowedClientId#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

