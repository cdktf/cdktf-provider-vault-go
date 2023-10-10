// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pkisecretbackendissuer

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PkiSecretBackendIssuerConfig struct {
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
	// Full path where PKI backend is mounted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.21.0/docs/resources/pki_secret_backend_issuer#backend PkiSecretBackendIssuer#backend}
	Backend *string `field:"required" json:"backend" yaml:"backend"`
	// Reference to an existing issuer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.21.0/docs/resources/pki_secret_backend_issuer#issuer_ref PkiSecretBackendIssuer#issuer_ref}
	IssuerRef *string `field:"required" json:"issuerRef" yaml:"issuerRef"`
	// Specifies the URL values for the CRL Distribution Points field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.21.0/docs/resources/pki_secret_backend_issuer#crl_distribution_points PkiSecretBackendIssuer#crl_distribution_points}
	CrlDistributionPoints *[]*string `field:"optional" json:"crlDistributionPoints" yaml:"crlDistributionPoints"`
	// Specifies that the AIA URL values should be templated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.21.0/docs/resources/pki_secret_backend_issuer#enable_aia_url_templating PkiSecretBackendIssuer#enable_aia_url_templating}
	EnableAiaUrlTemplating interface{} `field:"optional" json:"enableAiaUrlTemplating" yaml:"enableAiaUrlTemplating"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.21.0/docs/resources/pki_secret_backend_issuer#id PkiSecretBackendIssuer#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Reference to an existing issuer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.21.0/docs/resources/pki_secret_backend_issuer#issuer_name PkiSecretBackendIssuer#issuer_name}
	IssuerName *string `field:"optional" json:"issuerName" yaml:"issuerName"`
	// Specifies the URL values for the Issuing Certificate field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.21.0/docs/resources/pki_secret_backend_issuer#issuing_certificates PkiSecretBackendIssuer#issuing_certificates}
	IssuingCertificates *[]*string `field:"optional" json:"issuingCertificates" yaml:"issuingCertificates"`
	// Behavior of a leaf's 'NotAfter' field during issuance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.21.0/docs/resources/pki_secret_backend_issuer#leaf_not_after_behavior PkiSecretBackendIssuer#leaf_not_after_behavior}
	LeafNotAfterBehavior *string `field:"optional" json:"leafNotAfterBehavior" yaml:"leafNotAfterBehavior"`
	// Chain of issuer references to build this issuer's computed CAChain field from, when non-empty.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.21.0/docs/resources/pki_secret_backend_issuer#manual_chain PkiSecretBackendIssuer#manual_chain}
	ManualChain *[]*string `field:"optional" json:"manualChain" yaml:"manualChain"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.21.0/docs/resources/pki_secret_backend_issuer#namespace PkiSecretBackendIssuer#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Specifies the URL values for the OCSP Servers field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.21.0/docs/resources/pki_secret_backend_issuer#ocsp_servers PkiSecretBackendIssuer#ocsp_servers}
	OcspServers *[]*string `field:"optional" json:"ocspServers" yaml:"ocspServers"`
	// Which signature algorithm to use when building CRLs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.21.0/docs/resources/pki_secret_backend_issuer#revocation_signature_algorithm PkiSecretBackendIssuer#revocation_signature_algorithm}
	RevocationSignatureAlgorithm *string `field:"optional" json:"revocationSignatureAlgorithm" yaml:"revocationSignatureAlgorithm"`
	// Comma-separated list of allowed usages for this issuer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.21.0/docs/resources/pki_secret_backend_issuer#usage PkiSecretBackendIssuer#usage}
	Usage *string `field:"optional" json:"usage" yaml:"usage"`
}

