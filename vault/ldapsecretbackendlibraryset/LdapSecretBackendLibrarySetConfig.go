package ldapsecretbackendlibraryset

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LdapSecretBackendLibrarySetConfig struct {
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
	// The name of the set of service accounts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.18.0/docs/resources/ldap_secret_backend_library_set#name LdapSecretBackendLibrarySet#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The names of all the service accounts that can be checked out from this set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.18.0/docs/resources/ldap_secret_backend_library_set#service_account_names LdapSecretBackendLibrarySet#service_account_names}
	ServiceAccountNames *[]*string `field:"required" json:"serviceAccountNames" yaml:"serviceAccountNames"`
	// Disable enforcing that service accounts must be checked in by the entity or client token that checked them out.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.18.0/docs/resources/ldap_secret_backend_library_set#disable_check_in_enforcement LdapSecretBackendLibrarySet#disable_check_in_enforcement}
	DisableCheckInEnforcement interface{} `field:"optional" json:"disableCheckInEnforcement" yaml:"disableCheckInEnforcement"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.18.0/docs/resources/ldap_secret_backend_library_set#id LdapSecretBackendLibrarySet#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The maximum amount of time a check-out last with renewal before Vault automatically checks it back in.
	//
	// Defaults to 24 hours.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.18.0/docs/resources/ldap_secret_backend_library_set#max_ttl LdapSecretBackendLibrarySet#max_ttl}
	MaxTtl *float64 `field:"optional" json:"maxTtl" yaml:"maxTtl"`
	// The path where the LDAP secrets backend is mounted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.18.0/docs/resources/ldap_secret_backend_library_set#mount LdapSecretBackendLibrarySet#mount}
	Mount *string `field:"optional" json:"mount" yaml:"mount"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.18.0/docs/resources/ldap_secret_backend_library_set#namespace LdapSecretBackendLibrarySet#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// The maximum amount of time a single check-out lasts before Vault automatically checks it back in.
	//
	// Defaults to 24 hours.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.18.0/docs/resources/ldap_secret_backend_library_set#ttl LdapSecretBackendLibrarySet#ttl}
	Ttl *float64 `field:"optional" json:"ttl" yaml:"ttl"`
}

