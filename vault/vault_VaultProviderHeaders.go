// Prebuilt vault Provider for Terraform CDK (cdktf)
package vault


type VaultProviderHeaders struct {
	// The header name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault#name VaultProvider#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The header value.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault#value VaultProvider#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

