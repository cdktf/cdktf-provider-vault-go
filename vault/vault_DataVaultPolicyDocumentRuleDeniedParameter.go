// Prebuilt vault Provider for Terraform CDK (cdktf)
package vault


type DataVaultPolicyDocumentRuleDeniedParameter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/d/policy_document#key DataVaultPolicyDocument#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/d/policy_document#value DataVaultPolicyDocument#value}.
	Value *[]*string `field:"required" json:"value" yaml:"value"`
}

