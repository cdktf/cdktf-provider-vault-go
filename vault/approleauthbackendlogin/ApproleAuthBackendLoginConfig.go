// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package approleauthbackendlogin

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApproleAuthBackendLoginConfig struct {
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
	// The RoleID to log in with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.20.0/docs/resources/approle_auth_backend_login#role_id ApproleAuthBackendLogin#role_id}
	RoleId *string `field:"required" json:"roleId" yaml:"roleId"`
	// Unique name of the auth backend to configure.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.20.0/docs/resources/approle_auth_backend_login#backend ApproleAuthBackendLogin#backend}
	Backend *string `field:"optional" json:"backend" yaml:"backend"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.20.0/docs/resources/approle_auth_backend_login#id ApproleAuthBackendLogin#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.20.0/docs/resources/approle_auth_backend_login#namespace ApproleAuthBackendLogin#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// The SecretID to log in with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.20.0/docs/resources/approle_auth_backend_login#secret_id ApproleAuthBackendLogin#secret_id}
	SecretId *string `field:"optional" json:"secretId" yaml:"secretId"`
}

