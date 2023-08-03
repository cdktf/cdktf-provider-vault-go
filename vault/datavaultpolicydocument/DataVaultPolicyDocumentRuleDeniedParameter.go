package datavaultpolicydocument


type DataVaultPolicyDocumentRuleDeniedParameter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.19.0/docs/data-sources/policy_document#key DataVaultPolicyDocument#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.19.0/docs/data-sources/policy_document#value DataVaultPolicyDocument#value}.
	Value *[]*string `field:"required" json:"value" yaml:"value"`
}

