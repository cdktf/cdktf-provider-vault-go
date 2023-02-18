package provider


type VaultProviderAuthLoginCert struct {
	// Path to a file containing the client certificate.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault#cert_file VaultProvider#cert_file}
	CertFile *string `field:"required" json:"certFile" yaml:"certFile"`
	// Path to a file containing the private key that the certificate was issued for.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault#key_file VaultProvider#key_file}
	KeyFile *string `field:"required" json:"keyFile" yaml:"keyFile"`
	// The path where the authentication engine is mounted.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault#mount VaultProvider#mount}
	Mount *string `field:"optional" json:"mount" yaml:"mount"`
	// Name of the certificate's role.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault#name VaultProvider#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The authentication engine's namespace.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault#namespace VaultProvider#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

