// Prebuilt vault Provider for Terraform CDK (cdktf)
package vault


type SshSecretBackendRoleAllowedUserKeyConfig struct {
	// List of allowed key lengths, vault-1.10 and above.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/ssh_secret_backend_role#lengths SshSecretBackendRole#lengths}
	Lengths *[]*float64 `field:"required" json:"lengths" yaml:"lengths"`
	// Key type, choices: rsa, ecdsa, ec, dsa, ed25519, ssh-rsa, ssh-dss, ssh-ed25519, ecdsa-sha2-nistp256, ecdsa-sha2-nistp384, ecdsa-sha2-nistp521.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/ssh_secret_backend_role#type SshSecretBackendRole#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

