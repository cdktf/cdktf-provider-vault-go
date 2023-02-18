package gcpsecretstaticaccount


type GcpSecretStaticAccountBinding struct {
	// Resource name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/gcp_secret_static_account#resource GcpSecretStaticAccount#resource}
	Resource *string `field:"required" json:"resource" yaml:"resource"`
	// List of roles to apply to the resource.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/gcp_secret_static_account#roles GcpSecretStaticAccount#roles}
	Roles *[]*string `field:"required" json:"roles" yaml:"roles"`
}

