package provider


type VaultProviderAuthLoginJwt struct {
	// A signed JSON Web Token.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault#jwt VaultProvider#jwt}
	Jwt *string `field:"required" json:"jwt" yaml:"jwt"`
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

