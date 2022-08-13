// Prebuilt vault Provider for Terraform CDK (cdktf)
package vault

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GithubUserConfig struct {
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
	// GitHub user name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/github_user#user GithubUser#user}
	User *string `field:"required" json:"user" yaml:"user"`
	// Auth backend to which user mapping will be congigured.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/github_user#backend GithubUser#backend}
	Backend *string `field:"optional" json:"backend" yaml:"backend"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/github_user#id GithubUser#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/github_user#namespace GithubUser#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Policies to be assigned to this user.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/github_user#policies GithubUser#policies}
	Policies *[]*string `field:"optional" json:"policies" yaml:"policies"`
}

