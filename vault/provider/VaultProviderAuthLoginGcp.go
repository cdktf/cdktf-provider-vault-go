package provider


type VaultProviderAuthLoginGcp struct {
	// Name of the login role.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault#role VaultProvider#role}
	Role *string `field:"required" json:"role" yaml:"role"`
	// Path to the Google Cloud credentials file.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault#credentials VaultProvider#credentials}
	Credentials *string `field:"optional" json:"credentials" yaml:"credentials"`
	// A signed JSON Web Token.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault#jwt VaultProvider#jwt}
	Jwt *string `field:"optional" json:"jwt" yaml:"jwt"`
	// The path where the authentication engine is mounted.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault#mount VaultProvider#mount}
	Mount *string `field:"optional" json:"mount" yaml:"mount"`
	// The authentication engine's namespace.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault#namespace VaultProvider#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// IAM service account.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault#service_account VaultProvider#service_account}
	ServiceAccount *string `field:"optional" json:"serviceAccount" yaml:"serviceAccount"`
}

