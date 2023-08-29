// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package awssecretbackendrole

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AwsSecretBackendRoleConfig struct {
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
	// The path of the AWS Secret Backend the role belongs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.19.0/docs/resources/aws_secret_backend_role#backend AwsSecretBackendRole#backend}
	Backend *string `field:"required" json:"backend" yaml:"backend"`
	// Role credential type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.19.0/docs/resources/aws_secret_backend_role#credential_type AwsSecretBackendRole#credential_type}
	CredentialType *string `field:"required" json:"credentialType" yaml:"credentialType"`
	// Unique name for the role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.19.0/docs/resources/aws_secret_backend_role#name AwsSecretBackendRole#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The default TTL in seconds for STS credentials.
	//
	// When a TTL is not specified when STS credentials are requested, and a default TTL is specified on the role, then this default TTL will be used. Valid only when credential_type is one of assumed_role or federation_token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.19.0/docs/resources/aws_secret_backend_role#default_sts_ttl AwsSecretBackendRole#default_sts_ttl}
	DefaultStsTtl *float64 `field:"optional" json:"defaultStsTtl" yaml:"defaultStsTtl"`
	// A list of IAM group names.
	//
	// IAM users generated against this vault role will be added to these IAM Groups. For a credential type of assumed_role or federation_token, the policies sent to the corresponding AWS call (sts:AssumeRole or sts:GetFederation) will be the policies from each group in iam_groups combined with the policy_document and policy_arns parameters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.19.0/docs/resources/aws_secret_backend_role#iam_groups AwsSecretBackendRole#iam_groups}
	IamGroups *[]*string `field:"optional" json:"iamGroups" yaml:"iamGroups"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.19.0/docs/resources/aws_secret_backend_role#id AwsSecretBackendRole#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The max allowed TTL in seconds for STS credentials (credentials TTL are capped to max_sts_ttl).
	//
	// Valid only when credential_type is one of assumed_role or federation_token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.19.0/docs/resources/aws_secret_backend_role#max_sts_ttl AwsSecretBackendRole#max_sts_ttl}
	MaxStsTtl *float64 `field:"optional" json:"maxStsTtl" yaml:"maxStsTtl"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.19.0/docs/resources/aws_secret_backend_role#namespace AwsSecretBackendRole#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// The ARN of the AWS Permissions Boundary to attach to IAM users created in the role.
	//
	// Valid only when credential_type is iam_user. If not specified, then no permissions boundary policy will be attached.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.19.0/docs/resources/aws_secret_backend_role#permissions_boundary_arn AwsSecretBackendRole#permissions_boundary_arn}
	PermissionsBoundaryArn *string `field:"optional" json:"permissionsBoundaryArn" yaml:"permissionsBoundaryArn"`
	// ARN for an existing IAM policy the role should use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.19.0/docs/resources/aws_secret_backend_role#policy_arns AwsSecretBackendRole#policy_arns}
	PolicyArns *[]*string `field:"optional" json:"policyArns" yaml:"policyArns"`
	// IAM policy the role should use in JSON format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.19.0/docs/resources/aws_secret_backend_role#policy_document AwsSecretBackendRole#policy_document}
	PolicyDocument *string `field:"optional" json:"policyDocument" yaml:"policyDocument"`
	// ARNs of AWS roles allowed to be assumed. Only valid when credential_type is 'assumed_role'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.19.0/docs/resources/aws_secret_backend_role#role_arns AwsSecretBackendRole#role_arns}
	RoleArns *[]*string `field:"optional" json:"roleArns" yaml:"roleArns"`
	// The path for the user name. Valid only when credential_type is iam_user. Default is /.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.19.0/docs/resources/aws_secret_backend_role#user_path AwsSecretBackendRole#user_path}
	UserPath *string `field:"optional" json:"userPath" yaml:"userPath"`
}

