package provider


type VaultProviderAuthLoginRadius struct {
	// The Radius password for username.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault#password VaultProvider#password}
	Password *string `field:"required" json:"password" yaml:"password"`
	// The Radius username.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault#username VaultProvider#username}
	Username *string `field:"required" json:"username" yaml:"username"`
	// The path where the authentication engine is mounted.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault#mount VaultProvider#mount}
	Mount *string `field:"optional" json:"mount" yaml:"mount"`
	// The authentication engine's namespace.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault#namespace VaultProvider#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

