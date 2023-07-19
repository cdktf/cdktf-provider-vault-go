package pkisecretbackendrole


type PkiSecretBackendRolePolicyIdentifier struct {
	// OID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.18.0/docs/resources/pki_secret_backend_role#oid PkiSecretBackendRole#oid}
	Oid *string `field:"required" json:"oid" yaml:"oid"`
	// Optional CPS URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.18.0/docs/resources/pki_secret_backend_role#cps PkiSecretBackendRole#cps}
	Cps *string `field:"optional" json:"cps" yaml:"cps"`
	// Optional notice.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.18.0/docs/resources/pki_secret_backend_role#notice PkiSecretBackendRole#notice}
	Notice *string `field:"optional" json:"notice" yaml:"notice"`
}

