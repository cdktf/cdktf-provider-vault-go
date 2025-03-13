// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kmipsecretrole

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KmipSecretRoleConfig struct {
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
	// Path where KMIP backend is mounted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.7.0/docs/resources/kmip_secret_role#path KmipSecretRole#path}
	Path *string `field:"required" json:"path" yaml:"path"`
	// Name of the role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.7.0/docs/resources/kmip_secret_role#role KmipSecretRole#role}
	Role *string `field:"required" json:"role" yaml:"role"`
	// Name of the scope.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.7.0/docs/resources/kmip_secret_role#scope KmipSecretRole#scope}
	Scope *string `field:"required" json:"scope" yaml:"scope"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.7.0/docs/resources/kmip_secret_role#id KmipSecretRole#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.7.0/docs/resources/kmip_secret_role#namespace KmipSecretRole#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Grant permission to use the KMIP Activate operation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.7.0/docs/resources/kmip_secret_role#operation_activate KmipSecretRole#operation_activate}
	OperationActivate interface{} `field:"optional" json:"operationActivate" yaml:"operationActivate"`
	// Grant permission to use the KMIP Add Attribute operation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.7.0/docs/resources/kmip_secret_role#operation_add_attribute KmipSecretRole#operation_add_attribute}
	OperationAddAttribute interface{} `field:"optional" json:"operationAddAttribute" yaml:"operationAddAttribute"`
	// Grant all permissions to this role. May not be specified with any other operation_* params.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.7.0/docs/resources/kmip_secret_role#operation_all KmipSecretRole#operation_all}
	OperationAll interface{} `field:"optional" json:"operationAll" yaml:"operationAll"`
	// Grant permission to use the KMIP Create operation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.7.0/docs/resources/kmip_secret_role#operation_create KmipSecretRole#operation_create}
	OperationCreate interface{} `field:"optional" json:"operationCreate" yaml:"operationCreate"`
	// Grant permission to use the KMIP Destroy operation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.7.0/docs/resources/kmip_secret_role#operation_destroy KmipSecretRole#operation_destroy}
	OperationDestroy interface{} `field:"optional" json:"operationDestroy" yaml:"operationDestroy"`
	// Grant permission to use the KMIP Discover Version operation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.7.0/docs/resources/kmip_secret_role#operation_discover_versions KmipSecretRole#operation_discover_versions}
	OperationDiscoverVersions interface{} `field:"optional" json:"operationDiscoverVersions" yaml:"operationDiscoverVersions"`
	// Grant permission to use the KMIP Get operation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.7.0/docs/resources/kmip_secret_role#operation_get KmipSecretRole#operation_get}
	OperationGet interface{} `field:"optional" json:"operationGet" yaml:"operationGet"`
	// Grant permission to use the KMIP Get Attribute List operation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.7.0/docs/resources/kmip_secret_role#operation_get_attribute_list KmipSecretRole#operation_get_attribute_list}
	OperationGetAttributeList interface{} `field:"optional" json:"operationGetAttributeList" yaml:"operationGetAttributeList"`
	// Grant permission to use the KMIP Get Attributes operation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.7.0/docs/resources/kmip_secret_role#operation_get_attributes KmipSecretRole#operation_get_attributes}
	OperationGetAttributes interface{} `field:"optional" json:"operationGetAttributes" yaml:"operationGetAttributes"`
	// Grant permission to use the KMIP Locate operation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.7.0/docs/resources/kmip_secret_role#operation_locate KmipSecretRole#operation_locate}
	OperationLocate interface{} `field:"optional" json:"operationLocate" yaml:"operationLocate"`
	// Remove all permissions from this role. May not be specified with any other operation_* params.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.7.0/docs/resources/kmip_secret_role#operation_none KmipSecretRole#operation_none}
	OperationNone interface{} `field:"optional" json:"operationNone" yaml:"operationNone"`
	// Grant permission to use the KMIP Register operation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.7.0/docs/resources/kmip_secret_role#operation_register KmipSecretRole#operation_register}
	OperationRegister interface{} `field:"optional" json:"operationRegister" yaml:"operationRegister"`
	// Grant permission to use the KMIP Rekey operation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.7.0/docs/resources/kmip_secret_role#operation_rekey KmipSecretRole#operation_rekey}
	OperationRekey interface{} `field:"optional" json:"operationRekey" yaml:"operationRekey"`
	// Grant permission to use the KMIP Revoke operation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.7.0/docs/resources/kmip_secret_role#operation_revoke KmipSecretRole#operation_revoke}
	OperationRevoke interface{} `field:"optional" json:"operationRevoke" yaml:"operationRevoke"`
	// Client certificate key bits, valid values depend on key type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.7.0/docs/resources/kmip_secret_role#tls_client_key_bits KmipSecretRole#tls_client_key_bits}
	TlsClientKeyBits *float64 `field:"optional" json:"tlsClientKeyBits" yaml:"tlsClientKeyBits"`
	// Client certificate key type, rsa or ec.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.7.0/docs/resources/kmip_secret_role#tls_client_key_type KmipSecretRole#tls_client_key_type}
	TlsClientKeyType *string `field:"optional" json:"tlsClientKeyType" yaml:"tlsClientKeyType"`
	// Client certificate TTL in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.7.0/docs/resources/kmip_secret_role#tls_client_ttl KmipSecretRole#tls_client_ttl}
	TlsClientTtl *float64 `field:"optional" json:"tlsClientTtl" yaml:"tlsClientTtl"`
}

