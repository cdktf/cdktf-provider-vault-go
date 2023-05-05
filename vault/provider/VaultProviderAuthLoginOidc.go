package provider


type VaultProviderAuthLoginOidc struct {
	// Name of the login role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.15.2/docs#role VaultProvider#role}
	Role *string `field:"required" json:"role" yaml:"role"`
	// The callback address. Must be a valid URI without the path.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.15.2/docs#callback_address VaultProvider#callback_address}
	CallbackAddress *string `field:"optional" json:"callbackAddress" yaml:"callbackAddress"`
	// The callback listener's address. Must be a valid URI without the path.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.15.2/docs#callback_listener_address VaultProvider#callback_listener_address}
	CallbackListenerAddress *string `field:"optional" json:"callbackListenerAddress" yaml:"callbackListenerAddress"`
	// The path where the authentication engine is mounted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.15.2/docs#mount VaultProvider#mount}
	Mount *string `field:"optional" json:"mount" yaml:"mount"`
	// The authentication engine's namespace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.15.2/docs#namespace VaultProvider#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

