package databasesecretbackendstaticrole

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DatabaseSecretBackendStaticRoleConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
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
	// The path of the Database Secret Backend the role belongs to.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/database_secret_backend_static_role#backend DatabaseSecretBackendStaticRole#backend}
	Backend *string `field:"required" json:"backend" yaml:"backend"`
	// Database connection to use for this role.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/database_secret_backend_static_role#db_name DatabaseSecretBackendStaticRole#db_name}
	DbName *string `field:"required" json:"dbName" yaml:"dbName"`
	// Unique name for the static role.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/database_secret_backend_static_role#name DatabaseSecretBackendStaticRole#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The amount of time Vault should wait before rotating the password, in seconds.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/database_secret_backend_static_role#rotation_period DatabaseSecretBackendStaticRole#rotation_period}
	RotationPeriod *float64 `field:"required" json:"rotationPeriod" yaml:"rotationPeriod"`
	// The database username that this role corresponds to.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/database_secret_backend_static_role#username DatabaseSecretBackendStaticRole#username}
	Username *string `field:"required" json:"username" yaml:"username"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/database_secret_backend_static_role#id DatabaseSecretBackendStaticRole#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/database_secret_backend_static_role#namespace DatabaseSecretBackendStaticRole#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Database statements to execute to rotate the password for the configured database user.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/database_secret_backend_static_role#rotation_statements DatabaseSecretBackendStaticRole#rotation_statements}
	RotationStatements *[]*string `field:"optional" json:"rotationStatements" yaml:"rotationStatements"`
}

