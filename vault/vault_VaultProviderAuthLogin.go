// Prebuilt vault Provider for Terraform CDK (cdktf)
package vault


type VaultProviderAuthLogin struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault#path VaultProvider#path}.
	Path *string `field:"required" json:"path" yaml:"path"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault#method VaultProvider#method}.
	Method *string `field:"optional" json:"method" yaml:"method"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault#namespace VaultProvider#namespace}.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault#parameters VaultProvider#parameters}.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
}

