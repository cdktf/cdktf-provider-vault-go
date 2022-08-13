// Prebuilt vault Provider for Terraform CDK (cdktf)
package vault


type PkiSecretBackendRolePolicyIdentifier struct {
	// OID.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/pki_secret_backend_role#oid PkiSecretBackendRole#oid}
	Oid *string `field:"required" json:"oid" yaml:"oid"`
	// Optional CPS URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/pki_secret_backend_role#cps PkiSecretBackendRole#cps}
	Cps *string `field:"optional" json:"cps" yaml:"cps"`
	// Optional notice.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/pki_secret_backend_role#notice PkiSecretBackendRole#notice}
	Notice *string `field:"optional" json:"notice" yaml:"notice"`
}

