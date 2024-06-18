// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pkisecretbackendintermediatesetsigned

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PkiSecretBackendIntermediateSetSignedConfig struct {
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
	// The PKI secret backend the resource belongs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.3.0/docs/resources/pki_secret_backend_intermediate_set_signed#backend PkiSecretBackendIntermediateSetSigned#backend}
	Backend *string `field:"required" json:"backend" yaml:"backend"`
	// The certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.3.0/docs/resources/pki_secret_backend_intermediate_set_signed#certificate PkiSecretBackendIntermediateSetSigned#certificate}
	Certificate *string `field:"required" json:"certificate" yaml:"certificate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.3.0/docs/resources/pki_secret_backend_intermediate_set_signed#id PkiSecretBackendIntermediateSetSigned#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.3.0/docs/resources/pki_secret_backend_intermediate_set_signed#namespace PkiSecretBackendIntermediateSetSigned#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

