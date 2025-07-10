// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datavaultidentitygroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataVaultIdentityGroupConfig struct {
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
	// ID of the alias.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.1.0/docs/data-sources/identity_group#alias_id DataVaultIdentityGroup#alias_id}
	AliasId *string `field:"optional" json:"aliasId" yaml:"aliasId"`
	// Accessor of the mount to which the alias belongs to. This should be supplied in conjunction with `alias_name`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.1.0/docs/data-sources/identity_group#alias_mount_accessor DataVaultIdentityGroup#alias_mount_accessor}
	AliasMountAccessor *string `field:"optional" json:"aliasMountAccessor" yaml:"aliasMountAccessor"`
	// Name of the alias. This should be supplied in conjunction with `alias_mount_accessor`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.1.0/docs/data-sources/identity_group#alias_name DataVaultIdentityGroup#alias_name}
	AliasName *string `field:"optional" json:"aliasName" yaml:"aliasName"`
	// ID of the group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.1.0/docs/data-sources/identity_group#group_id DataVaultIdentityGroup#group_id}
	GroupId *string `field:"optional" json:"groupId" yaml:"groupId"`
	// Name of the group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.1.0/docs/data-sources/identity_group#group_name DataVaultIdentityGroup#group_name}
	GroupName *string `field:"optional" json:"groupName" yaml:"groupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.1.0/docs/data-sources/identity_group#id DataVaultIdentityGroup#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.1.0/docs/data-sources/identity_group#namespace DataVaultIdentityGroup#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

