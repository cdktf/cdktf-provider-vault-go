// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mongodbatlassecretbackend

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MongodbatlasSecretBackendConfig struct {
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
	// Path where MongoDB Atlas secret backend is mounted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.4.0/docs/resources/mongodbatlas_secret_backend#mount MongodbatlasSecretBackend#mount}
	Mount *string `field:"required" json:"mount" yaml:"mount"`
	// The Private Programmatic API Key used to connect with MongoDB Atlas API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.4.0/docs/resources/mongodbatlas_secret_backend#private_key MongodbatlasSecretBackend#private_key}
	PrivateKey *string `field:"required" json:"privateKey" yaml:"privateKey"`
	// The Public Programmatic API Key used to authenticate with the MongoDB Atlas API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.4.0/docs/resources/mongodbatlas_secret_backend#public_key MongodbatlasSecretBackend#public_key}
	PublicKey *string `field:"required" json:"publicKey" yaml:"publicKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.4.0/docs/resources/mongodbatlas_secret_backend#id MongodbatlasSecretBackend#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.4.0/docs/resources/mongodbatlas_secret_backend#namespace MongodbatlasSecretBackend#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

