// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package quotaratelimit

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type QuotaRateLimitConfig struct {
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
	// The name of the quota.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.20.1/docs/resources/quota_rate_limit#name QuotaRateLimit#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The maximum number of requests at any given second to be allowed by the quota rule.
	//
	// The rate must be positive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.20.1/docs/resources/quota_rate_limit#rate QuotaRateLimit#rate}
	Rate *float64 `field:"required" json:"rate" yaml:"rate"`
	// If set, when a client reaches a rate limit threshold, the client will be prohibited from any further requests until after the 'block_interval' in seconds has elapsed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.20.1/docs/resources/quota_rate_limit#block_interval QuotaRateLimit#block_interval}
	BlockInterval *float64 `field:"optional" json:"blockInterval" yaml:"blockInterval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.20.1/docs/resources/quota_rate_limit#id QuotaRateLimit#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The duration in seconds to enforce rate limiting for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.20.1/docs/resources/quota_rate_limit#interval QuotaRateLimit#interval}
	Interval *float64 `field:"optional" json:"interval" yaml:"interval"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.20.1/docs/resources/quota_rate_limit#namespace QuotaRateLimit#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Path of the mount or namespace to apply the quota. A blank path configures a global rate limit quota.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.20.1/docs/resources/quota_rate_limit#path QuotaRateLimit#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
	// If set on a quota where path is set to an auth mount with a concept of roles (such as /auth/approle/), this will make the quota restrict login requests to that mount that are made with the specified role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.20.1/docs/resources/quota_rate_limit#role QuotaRateLimit#role}
	Role *string `field:"optional" json:"role" yaml:"role"`
}

