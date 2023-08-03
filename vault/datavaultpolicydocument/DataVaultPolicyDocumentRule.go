package datavaultpolicydocument


type DataVaultPolicyDocumentRule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.19.0/docs/data-sources/policy_document#capabilities DataVaultPolicyDocument#capabilities}.
	Capabilities *[]*string `field:"required" json:"capabilities" yaml:"capabilities"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.19.0/docs/data-sources/policy_document#path DataVaultPolicyDocument#path}.
	Path *string `field:"required" json:"path" yaml:"path"`
	// allowed_parameter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.19.0/docs/data-sources/policy_document#allowed_parameter DataVaultPolicyDocument#allowed_parameter}
	AllowedParameter interface{} `field:"optional" json:"allowedParameter" yaml:"allowedParameter"`
	// denied_parameter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.19.0/docs/data-sources/policy_document#denied_parameter DataVaultPolicyDocument#denied_parameter}
	DeniedParameter interface{} `field:"optional" json:"deniedParameter" yaml:"deniedParameter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.19.0/docs/data-sources/policy_document#description DataVaultPolicyDocument#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.19.0/docs/data-sources/policy_document#max_wrapping_ttl DataVaultPolicyDocument#max_wrapping_ttl}.
	MaxWrappingTtl *string `field:"optional" json:"maxWrappingTtl" yaml:"maxWrappingTtl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.19.0/docs/data-sources/policy_document#min_wrapping_ttl DataVaultPolicyDocument#min_wrapping_ttl}.
	MinWrappingTtl *string `field:"optional" json:"minWrappingTtl" yaml:"minWrappingTtl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.19.0/docs/data-sources/policy_document#required_parameters DataVaultPolicyDocument#required_parameters}.
	RequiredParameters *[]*string `field:"optional" json:"requiredParameters" yaml:"requiredParameters"`
}

