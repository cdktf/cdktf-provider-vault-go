// Prebuilt vault Provider for Terraform CDK (cdktf)
package vault


type GcpSecretRolesetBinding struct {
	// Resource name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/gcp_secret_roleset#resource GcpSecretRoleset#resource}
	Resource *string `field:"required" json:"resource" yaml:"resource"`
	// List of roles to apply to the resource.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/gcp_secret_roleset#roles GcpSecretRoleset#roles}
	Roles *[]*string `field:"required" json:"roles" yaml:"roles"`
}

