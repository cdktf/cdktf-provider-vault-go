// Prebuilt vault Provider for Terraform CDK (cdktf)
package vault

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TransitSecretCacheConfigConfig struct {
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
	// The Transit secret backend the resource belongs to.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/transit_secret_cache_config#backend TransitSecretCacheConfig#backend}
	Backend *string `field:"required" json:"backend" yaml:"backend"`
	// Number of cache entries. A size of 0 mean unlimited.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/transit_secret_cache_config#size TransitSecretCacheConfig#size}
	Size *float64 `field:"required" json:"size" yaml:"size"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/transit_secret_cache_config#id TransitSecretCacheConfig#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/transit_secret_cache_config#namespace TransitSecretCacheConfig#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

