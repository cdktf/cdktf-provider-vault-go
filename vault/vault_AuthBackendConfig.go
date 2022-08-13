// Prebuilt vault Provider for Terraform CDK (cdktf)
package vault

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AuthBackendConfig struct {
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
	// Name of the auth backend.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/auth_backend#type AuthBackend#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// The description of the auth backend.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/auth_backend#description AuthBackend#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/auth_backend#id AuthBackend#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Specifies if the auth method is local only.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/auth_backend#local AuthBackend#local}
	Local interface{} `field:"optional" json:"local" yaml:"local"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/auth_backend#namespace AuthBackend#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// path to mount the backend. This defaults to the type.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/auth_backend#path AuthBackend#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/auth_backend#tune AuthBackend#tune}.
	Tune interface{} `field:"optional" json:"tune" yaml:"tune"`
}

