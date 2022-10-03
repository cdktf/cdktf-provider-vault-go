package transformtransformation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TransformTransformationConfig struct {
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
	// The name of the transformation.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/transform_transformation#name TransformTransformation#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The mount path for a back-end, for example, the path given in "$ vault auth enable -path=my-aws aws".
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/transform_transformation#path TransformTransformation#path}
	Path *string `field:"required" json:"path" yaml:"path"`
	// The set of roles allowed to perform this transformation.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/transform_transformation#allowed_roles TransformTransformation#allowed_roles}
	AllowedRoles *[]*string `field:"optional" json:"allowedRoles" yaml:"allowedRoles"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/transform_transformation#id TransformTransformation#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The character used to replace data when in masking mode.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/transform_transformation#masking_character TransformTransformation#masking_character}
	MaskingCharacter *string `field:"optional" json:"maskingCharacter" yaml:"maskingCharacter"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/transform_transformation#namespace TransformTransformation#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// The name of the template to use.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/transform_transformation#template TransformTransformation#template}
	Template *string `field:"optional" json:"template" yaml:"template"`
	// Templates configured for transformation.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/transform_transformation#templates TransformTransformation#templates}
	Templates *[]*string `field:"optional" json:"templates" yaml:"templates"`
	// The source of where the tweak value comes from. Only valid when in FPE mode.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/transform_transformation#tweak_source TransformTransformation#tweak_source}
	TweakSource *string `field:"optional" json:"tweakSource" yaml:"tweakSource"`
	// The type of transformation to perform.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/transform_transformation#type TransformTransformation#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

