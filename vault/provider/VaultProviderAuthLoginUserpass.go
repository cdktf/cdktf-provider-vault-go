package provider


type VaultProviderAuthLoginUserpass struct {
	// Login with username.
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
	// Login with password.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault#password VaultProvider#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Login with password from a file.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault#password_file VaultProvider#password_file}
	PasswordFile *string `field:"optional" json:"passwordFile" yaml:"passwordFile"`
}

