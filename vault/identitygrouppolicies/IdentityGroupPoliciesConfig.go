package identitygrouppolicies

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IdentityGroupPoliciesConfig struct {
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
	// ID of the group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.15.2/docs/resources/identity_group_policies#group_id IdentityGroupPolicies#group_id}
	GroupId *string `field:"required" json:"groupId" yaml:"groupId"`
	// Policies to be tied to the group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.15.2/docs/resources/identity_group_policies#policies IdentityGroupPolicies#policies}
	Policies *[]*string `field:"required" json:"policies" yaml:"policies"`
	// Should the resource manage policies exclusively? Beware of race conditions when disabling exclusive management.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.15.2/docs/resources/identity_group_policies#exclusive IdentityGroupPolicies#exclusive}
	Exclusive interface{} `field:"optional" json:"exclusive" yaml:"exclusive"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.15.2/docs/resources/identity_group_policies#id IdentityGroupPolicies#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.15.2/docs/resources/identity_group_policies#namespace IdentityGroupPolicies#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

