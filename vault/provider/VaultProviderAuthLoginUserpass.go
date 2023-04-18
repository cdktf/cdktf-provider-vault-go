package provider


type VaultProviderAuthLoginUserpass struct {
	// Login with username.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.15.0/docs#username VaultProvider#username}
	Username *string `field:"required" json:"username" yaml:"username"`
	// The path where the authentication engine is mounted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.15.0/docs#mount VaultProvider#mount}
	Mount *string `field:"optional" json:"mount" yaml:"mount"`
	// The authentication engine's namespace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.15.0/docs#namespace VaultProvider#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Login with password.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.15.0/docs#password VaultProvider#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Login with password from a file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.15.0/docs#password_file VaultProvider#password_file}
	PasswordFile *string `field:"optional" json:"passwordFile" yaml:"passwordFile"`
}

