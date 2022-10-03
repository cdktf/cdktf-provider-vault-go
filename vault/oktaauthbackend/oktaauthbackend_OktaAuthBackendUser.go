package oktaauthbackend


type OktaAuthBackendUser struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/okta_auth_backend#groups OktaAuthBackend#groups}.
	Groups *[]*string `field:"optional" json:"groups" yaml:"groups"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/okta_auth_backend#policies OktaAuthBackend#policies}.
	Policies *[]*string `field:"optional" json:"policies" yaml:"policies"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/okta_auth_backend#username OktaAuthBackend#username}.
	Username *string `field:"optional" json:"username" yaml:"username"`
}

