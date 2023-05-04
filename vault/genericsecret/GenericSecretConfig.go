package genericsecret

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GenericSecretConfig struct {
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
	// JSON-encoded secret data to write.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.15.1/docs/resources/generic_secret#data_json GenericSecret#data_json}
	DataJson *string `field:"required" json:"dataJson" yaml:"dataJson"`
	// Full path where the generic secret will be written.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.15.1/docs/resources/generic_secret#path GenericSecret#path}
	Path *string `field:"required" json:"path" yaml:"path"`
	// Only applicable for kv-v2 stores. If set, permanently deletes all versions for the specified key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.15.1/docs/resources/generic_secret#delete_all_versions GenericSecret#delete_all_versions}
	DeleteAllVersions interface{} `field:"optional" json:"deleteAllVersions" yaml:"deleteAllVersions"`
	// Don't attempt to read the token from Vault if true; drift won't be detected.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.15.1/docs/resources/generic_secret#disable_read GenericSecret#disable_read}
	DisableRead interface{} `field:"optional" json:"disableRead" yaml:"disableRead"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.15.1/docs/resources/generic_secret#id GenericSecret#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.15.1/docs/resources/generic_secret#namespace GenericSecret#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

