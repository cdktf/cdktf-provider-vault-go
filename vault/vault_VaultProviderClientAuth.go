// Prebuilt vault Provider for Terraform CDK (cdktf)
package vault


type VaultProviderClientAuth struct {
	// Path to a file containing the client certificate.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault#cert_file VaultProvider#cert_file}
	CertFile *string `field:"optional" json:"certFile" yaml:"certFile"`
	// Path to a file containing the private key that the certificate was issued for.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault#key_file VaultProvider#key_file}
	KeyFile *string `field:"optional" json:"keyFile" yaml:"keyFile"`
}

