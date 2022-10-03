package transformalphabet

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TransformAlphabetConfig struct {
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
	// The name of the alphabet.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/transform_alphabet#name TransformAlphabet#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The mount path for a back-end, for example, the path given in "$ vault auth enable -path=my-aws aws".
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/transform_alphabet#path TransformAlphabet#path}
	Path *string `field:"required" json:"path" yaml:"path"`
	// A string of characters that contains the alphabet set.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/transform_alphabet#alphabet TransformAlphabet#alphabet}
	Alphabet *string `field:"optional" json:"alphabet" yaml:"alphabet"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/transform_alphabet#id TransformAlphabet#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/transform_alphabet#namespace TransformAlphabet#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}
