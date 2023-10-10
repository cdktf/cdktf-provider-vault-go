// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sshsecretbackendca

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SshSecretBackendCaConfig struct {
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
	// The path of the SSH Secret Backend where the CA should be configured.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.21.0/docs/resources/ssh_secret_backend_ca#backend SshSecretBackendCa#backend}
	Backend *string `field:"optional" json:"backend" yaml:"backend"`
	// Whether Vault should generate the signing key pair internally.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.21.0/docs/resources/ssh_secret_backend_ca#generate_signing_key SshSecretBackendCa#generate_signing_key}
	GenerateSigningKey interface{} `field:"optional" json:"generateSigningKey" yaml:"generateSigningKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.21.0/docs/resources/ssh_secret_backend_ca#id SshSecretBackendCa#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.21.0/docs/resources/ssh_secret_backend_ca#namespace SshSecretBackendCa#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Private key part the SSH CA key pair; required if generate_signing_key is false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.21.0/docs/resources/ssh_secret_backend_ca#private_key SshSecretBackendCa#private_key}
	PrivateKey *string `field:"optional" json:"privateKey" yaml:"privateKey"`
	// Public key part the SSH CA key pair; required if generate_signing_key is false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.21.0/docs/resources/ssh_secret_backend_ca#public_key SshSecretBackendCa#public_key}
	PublicKey *string `field:"optional" json:"publicKey" yaml:"publicKey"`
}

