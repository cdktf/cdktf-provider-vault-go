package provider


type VaultProviderAuthLoginOci struct {
	// Authentication type to use when getting OCI credentials.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault#auth_type VaultProvider#auth_type}
	AuthType *string `field:"required" json:"authType" yaml:"authType"`
	// Name of the login role.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault#role VaultProvider#role}
	Role *string `field:"required" json:"role" yaml:"role"`
	// The path where the authentication engine is mounted.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault#mount VaultProvider#mount}
	Mount *string `field:"optional" json:"mount" yaml:"mount"`
	// The authentication engine's namespace.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault#namespace VaultProvider#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

