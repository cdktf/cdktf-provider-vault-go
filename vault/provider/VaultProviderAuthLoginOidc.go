package provider


type VaultProviderAuthLoginOidc struct {
	// Name of the login role.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault#role VaultProvider#role}
	Role *string `field:"required" json:"role" yaml:"role"`
	// The callback address. Must be a valid URI without the path.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault#callback_address VaultProvider#callback_address}
	CallbackAddress *string `field:"optional" json:"callbackAddress" yaml:"callbackAddress"`
	// The callback listener's address. Must be a valid URI without the path.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault#callback_listener_address VaultProvider#callback_listener_address}
	CallbackListenerAddress *string `field:"optional" json:"callbackListenerAddress" yaml:"callbackListenerAddress"`
	// The path where the authentication engine is mounted.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault#mount VaultProvider#mount}
	Mount *string `field:"optional" json:"mount" yaml:"mount"`
	// The authentication engine's namespace.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault#namespace VaultProvider#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

