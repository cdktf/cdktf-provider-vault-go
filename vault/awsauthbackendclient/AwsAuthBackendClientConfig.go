// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package awsauthbackendclient

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AwsAuthBackendClientConfig struct {
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
	// AWS Access key with permissions to query AWS APIs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.4.0/docs/resources/aws_auth_backend_client#access_key AwsAuthBackendClient#access_key}
	AccessKey *string `field:"optional" json:"accessKey" yaml:"accessKey"`
	// Unique name of the auth backend to configure.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.4.0/docs/resources/aws_auth_backend_client#backend AwsAuthBackendClient#backend}
	Backend *string `field:"optional" json:"backend" yaml:"backend"`
	// URL to override the default generated endpoint for making AWS EC2 API calls.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.4.0/docs/resources/aws_auth_backend_client#ec2_endpoint AwsAuthBackendClient#ec2_endpoint}
	Ec2Endpoint *string `field:"optional" json:"ec2Endpoint" yaml:"ec2Endpoint"`
	// URL to override the default generated endpoint for making AWS IAM API calls.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.4.0/docs/resources/aws_auth_backend_client#iam_endpoint AwsAuthBackendClient#iam_endpoint}
	IamEndpoint *string `field:"optional" json:"iamEndpoint" yaml:"iamEndpoint"`
	// The value to require in the X-Vault-AWS-IAM-Server-ID header as part of GetCallerIdentity requests that are used in the iam auth method.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.4.0/docs/resources/aws_auth_backend_client#iam_server_id_header_value AwsAuthBackendClient#iam_server_id_header_value}
	IamServerIdHeaderValue *string `field:"optional" json:"iamServerIdHeaderValue" yaml:"iamServerIdHeaderValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.4.0/docs/resources/aws_auth_backend_client#id AwsAuthBackendClient#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The audience claim value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.4.0/docs/resources/aws_auth_backend_client#identity_token_audience AwsAuthBackendClient#identity_token_audience}
	IdentityTokenAudience *string `field:"optional" json:"identityTokenAudience" yaml:"identityTokenAudience"`
	// The TTL of generated identity tokens in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.4.0/docs/resources/aws_auth_backend_client#identity_token_ttl AwsAuthBackendClient#identity_token_ttl}
	IdentityTokenTtl *float64 `field:"optional" json:"identityTokenTtl" yaml:"identityTokenTtl"`
	// Number of max retries the client should use for recoverable errors.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.4.0/docs/resources/aws_auth_backend_client#max_retries AwsAuthBackendClient#max_retries}
	MaxRetries *float64 `field:"optional" json:"maxRetries" yaml:"maxRetries"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.4.0/docs/resources/aws_auth_backend_client#namespace AwsAuthBackendClient#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Role ARN to assume for plugin identity token federation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.4.0/docs/resources/aws_auth_backend_client#role_arn AwsAuthBackendClient#role_arn}
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// AWS Secret key with permissions to query AWS APIs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.4.0/docs/resources/aws_auth_backend_client#secret_key AwsAuthBackendClient#secret_key}
	SecretKey *string `field:"optional" json:"secretKey" yaml:"secretKey"`
	// URL to override the default generated endpoint for making AWS STS API calls.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.4.0/docs/resources/aws_auth_backend_client#sts_endpoint AwsAuthBackendClient#sts_endpoint}
	StsEndpoint *string `field:"optional" json:"stsEndpoint" yaml:"stsEndpoint"`
	// Region to override the default region for making AWS STS API calls.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.4.0/docs/resources/aws_auth_backend_client#sts_region AwsAuthBackendClient#sts_region}
	StsRegion *string `field:"optional" json:"stsRegion" yaml:"stsRegion"`
	// If set, will override sts_region and use the region from the client request's header.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.4.0/docs/resources/aws_auth_backend_client#use_sts_region_from_client AwsAuthBackendClient#use_sts_region_from_client}
	UseStsRegionFromClient interface{} `field:"optional" json:"useStsRegionFromClient" yaml:"useStsRegionFromClient"`
}

