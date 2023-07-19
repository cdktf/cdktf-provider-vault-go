package transitsecretbackendkey

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TransitSecretBackendKeyConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The Transit secret backend the resource belongs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.18.0/docs/resources/transit_secret_backend_key#backend TransitSecretBackendKey#backend}
	Backend *string `field:"required" json:"backend" yaml:"backend"`
	// Name of the encryption key to create.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.18.0/docs/resources/transit_secret_backend_key#name TransitSecretBackendKey#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// If set, enables taking backup of named key in the plaintext format. Once set, this cannot be disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.18.0/docs/resources/transit_secret_backend_key#allow_plaintext_backup TransitSecretBackendKey#allow_plaintext_backup}
	AllowPlaintextBackup interface{} `field:"optional" json:"allowPlaintextBackup" yaml:"allowPlaintextBackup"`
	// Amount of time the key should live before being automatically rotated.
	//
	// A value of 0 disables automatic rotation for the key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.18.0/docs/resources/transit_secret_backend_key#auto_rotate_interval TransitSecretBackendKey#auto_rotate_interval}
	AutoRotateInterval *float64 `field:"optional" json:"autoRotateInterval" yaml:"autoRotateInterval"`
	// Amount of time the key should live before being automatically rotated.
	//
	// A value of 0 disables automatic rotation for the key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.18.0/docs/resources/transit_secret_backend_key#auto_rotate_period TransitSecretBackendKey#auto_rotate_period}
	AutoRotatePeriod *float64 `field:"optional" json:"autoRotatePeriod" yaml:"autoRotatePeriod"`
	// Whether or not to support convergent encryption, where the same plaintext creates the same ciphertext.
	//
	// This requires derived to be set to true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.18.0/docs/resources/transit_secret_backend_key#convergent_encryption TransitSecretBackendKey#convergent_encryption}
	ConvergentEncryption interface{} `field:"optional" json:"convergentEncryption" yaml:"convergentEncryption"`
	// Specifies if the key is allowed to be deleted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.18.0/docs/resources/transit_secret_backend_key#deletion_allowed TransitSecretBackendKey#deletion_allowed}
	DeletionAllowed interface{} `field:"optional" json:"deletionAllowed" yaml:"deletionAllowed"`
	// Specifies if key derivation is to be used.
	//
	// If enabled, all encrypt/decrypt requests to this key must provide a context which is used for key derivation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.18.0/docs/resources/transit_secret_backend_key#derived TransitSecretBackendKey#derived}
	Derived interface{} `field:"optional" json:"derived" yaml:"derived"`
	// Enables keys to be exportable.
	//
	// This allows for all the valid keys in the key ring to be exported. Once set, this cannot be disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.18.0/docs/resources/transit_secret_backend_key#exportable TransitSecretBackendKey#exportable}
	Exportable interface{} `field:"optional" json:"exportable" yaml:"exportable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.18.0/docs/resources/transit_secret_backend_key#id TransitSecretBackendKey#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Minimum key version to use for decryption.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.18.0/docs/resources/transit_secret_backend_key#min_decryption_version TransitSecretBackendKey#min_decryption_version}
	MinDecryptionVersion *float64 `field:"optional" json:"minDecryptionVersion" yaml:"minDecryptionVersion"`
	// Minimum key version to use for encryption.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.18.0/docs/resources/transit_secret_backend_key#min_encryption_version TransitSecretBackendKey#min_encryption_version}
	MinEncryptionVersion *float64 `field:"optional" json:"minEncryptionVersion" yaml:"minEncryptionVersion"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.18.0/docs/resources/transit_secret_backend_key#namespace TransitSecretBackendKey#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Specifies the type of key to create.
	//
	// The currently-supported types are: aes128-gcm96, aes256-gcm96, chacha20-poly1305, ed25519, ecdsa-p256, ecdsa-p384, ecdsa-p521, rsa-2048, rsa-3072, rsa-4096
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.18.0/docs/resources/transit_secret_backend_key#type TransitSecretBackendKey#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

