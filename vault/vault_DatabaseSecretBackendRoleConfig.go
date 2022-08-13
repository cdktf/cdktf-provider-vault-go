// Prebuilt vault Provider for Terraform CDK (cdktf)
package vault

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DatabaseSecretBackendRoleConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/database_secret_backend_role#backend DatabaseSecretBackendRole#backend}
	Backend *string `field:"required" json:"backend" yaml:"backend"`
	// Database statements to execute to create and configure a user.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/database_secret_backend_role#creation_statements DatabaseSecretBackendRole#creation_statements}
	CreationStatements *[]*string `field:"required" json:"creationStatements" yaml:"creationStatements"`
	// Database connection to use for this role.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/database_secret_backend_role#db_name DatabaseSecretBackendRole#db_name}
	DbName *string `field:"required" json:"dbName" yaml:"dbName"`
	// Unique name for the role.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/database_secret_backend_role#name DatabaseSecretBackendRole#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Default TTL for leases associated with this role, in seconds.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/database_secret_backend_role#default_ttl DatabaseSecretBackendRole#default_ttl}
	DefaultTtl *float64 `field:"optional" json:"defaultTtl" yaml:"defaultTtl"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/database_secret_backend_role#id DatabaseSecretBackendRole#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Maximum TTL for leases associated with this role, in seconds.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/database_secret_backend_role#max_ttl DatabaseSecretBackendRole#max_ttl}
	MaxTtl *float64 `field:"optional" json:"maxTtl" yaml:"maxTtl"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/database_secret_backend_role#namespace DatabaseSecretBackendRole#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Database statements to execute to renew a user.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/database_secret_backend_role#renew_statements DatabaseSecretBackendRole#renew_statements}
	RenewStatements *[]*string `field:"optional" json:"renewStatements" yaml:"renewStatements"`
	// Database statements to execute to revoke a user.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/database_secret_backend_role#revocation_statements DatabaseSecretBackendRole#revocation_statements}
	RevocationStatements *[]*string `field:"optional" json:"revocationStatements" yaml:"revocationStatements"`
	// Database statements to execute to rollback a create operation in the event of an error.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/database_secret_backend_role#rollback_statements DatabaseSecretBackendRole#rollback_statements}
	RollbackStatements *[]*string `field:"optional" json:"rollbackStatements" yaml:"rollbackStatements"`
}

