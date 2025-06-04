// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datavaulttransitverify

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataVaultTransitVerifyConfig struct {
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
	// Specifies the name of the encryption key that was used to generate the signature or HMAC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.0.0/docs/data-sources/transit_verify#name DataVaultTransitVerify#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The Transit secret backend the key belongs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.0.0/docs/data-sources/transit_verify#path DataVaultTransitVerify#path}
	Path *string `field:"required" json:"path" yaml:"path"`
	// Specifies a list of items for processing.
	//
	// When this parameter is set, any supplied 'input' or 'context' parameters will be ignored. Responses are returned in the 'batch_results' array component of the 'data' element of the response. Any batch output will preserve the order of the batch input. If the input data value of an item is invalid, the corresponding item in the 'batch_results' will have the key 'error' with a value describing the error.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.0.0/docs/data-sources/transit_verify#batch_input DataVaultTransitVerify#batch_input}
	BatchInput interface{} `field:"optional" json:"batchInput" yaml:"batchInput"`
	// The results returned from Vault if using batch_input.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.0.0/docs/data-sources/transit_verify#batch_results DataVaultTransitVerify#batch_results}
	BatchResults interface{} `field:"optional" json:"batchResults" yaml:"batchResults"`
	// (Enterprise only) Specifies the signature output from the /transit/cmac function.
	//
	// One of the following arguments must be supplied signature, hmac or cmac.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.0.0/docs/data-sources/transit_verify#cmac DataVaultTransitVerify#cmac}
	Cmac *string `field:"optional" json:"cmac" yaml:"cmac"`
	// Base64 encoded context for key derivation. Required if key derivation is enabled; currently only available with ed25519 keys.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.0.0/docs/data-sources/transit_verify#context DataVaultTransitVerify#context}
	Context *string `field:"optional" json:"context" yaml:"context"`
	// Specifies the hash algorithm to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.0.0/docs/data-sources/transit_verify#hash_algorithm DataVaultTransitVerify#hash_algorithm}
	HashAlgorithm *string `field:"optional" json:"hashAlgorithm" yaml:"hashAlgorithm"`
	// Specifies the signature output from the /transit/hmac function.
	//
	// One of the following arguments must be supplied signature, hmac or cmac.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.0.0/docs/data-sources/transit_verify#hmac DataVaultTransitVerify#hmac}
	Hmac *string `field:"optional" json:"hmac" yaml:"hmac"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.0.0/docs/data-sources/transit_verify#id DataVaultTransitVerify#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Specifies the base64 encoded input data. One of input or batch_input must be supplied.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.0.0/docs/data-sources/transit_verify#input DataVaultTransitVerify#input}
	Input *string `field:"optional" json:"input" yaml:"input"`
	// Specifies the way in which the signature was originally marshaled. This currently only applies to ECDSA keys.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.0.0/docs/data-sources/transit_verify#marshaling_algorithm DataVaultTransitVerify#marshaling_algorithm}
	MarshalingAlgorithm *string `field:"optional" json:"marshalingAlgorithm" yaml:"marshalingAlgorithm"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.0.0/docs/data-sources/transit_verify#namespace DataVaultTransitVerify#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Set to true when the input is already hashed.
	//
	// If the key type is rsa-2048, rsa-3072 or rsa-4096, then the algorithm used to hash the input should be indicated by the hash_algorithm parameter. Just as the value to sign should be the base64-encoded representation of the exact binary data you want signed, when set, input is expected to be base64-encoded binary hashed data, not hex-formatted. (As an example, on the command line, you could generate a suitable input via openssl dgst -sha256 -binary | base64.)
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.0.0/docs/data-sources/transit_verify#prehashed DataVaultTransitVerify#prehashed}
	Prehashed interface{} `field:"optional" json:"prehashed" yaml:"prehashed"`
	// A user-supplied string that will be present in the reference field on the corresponding batch_results item in the response, to assist in understanding which result corresponds to a particular input.
	//
	// Only valid on batch requests when using ‘batch_input’ below.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.0.0/docs/data-sources/transit_verify#reference DataVaultTransitVerify#reference}
	Reference *string `field:"optional" json:"reference" yaml:"reference"`
	// The salt length used to sign. This currently only applies to the RSA PSS signature scheme.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.0.0/docs/data-sources/transit_verify#salt_length DataVaultTransitVerify#salt_length}
	SaltLength *string `field:"optional" json:"saltLength" yaml:"saltLength"`
	// Specifies the signature output from the /transit/sign function.
	//
	// One of the following arguments must be supplied signature, hmac or cmac.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.0.0/docs/data-sources/transit_verify#signature DataVaultTransitVerify#signature}
	Signature *string `field:"optional" json:"signature" yaml:"signature"`
	// When using a RSA key, specifies the RSA signature algorithm to use for signature verification.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.0.0/docs/data-sources/transit_verify#signature_algorithm DataVaultTransitVerify#signature_algorithm}
	SignatureAlgorithm *string `field:"optional" json:"signatureAlgorithm" yaml:"signatureAlgorithm"`
	// Base64 encoded context for Ed25519ctx and Ed25519ph signatures.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.0.0/docs/data-sources/transit_verify#signature_context DataVaultTransitVerify#signature_context}
	SignatureContext *string `field:"optional" json:"signatureContext" yaml:"signatureContext"`
	// Indicates whether verification succeeded.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.0.0/docs/data-sources/transit_verify#valid DataVaultTransitVerify#valid}
	Valid interface{} `field:"optional" json:"valid" yaml:"valid"`
}

