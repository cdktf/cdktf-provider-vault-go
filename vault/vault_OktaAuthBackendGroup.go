// Prebuilt vault Provider for Terraform CDK (cdktf)
package vault


type OktaAuthBackendGroup struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/okta_auth_backend#group_name OktaAuthBackend#group_name}.
	GroupName *string `field:"optional" json:"groupName" yaml:"groupName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/okta_auth_backend#policies OktaAuthBackend#policies}.
	Policies *[]*string `field:"optional" json:"policies" yaml:"policies"`
}

