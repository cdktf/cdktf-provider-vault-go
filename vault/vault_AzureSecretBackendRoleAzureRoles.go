// Prebuilt vault Provider for Terraform CDK (cdktf)
package vault


type AzureSecretBackendRoleAzureRoles struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/azure_secret_backend_role#role_name AzureSecretBackendRole#role_name}.
	RoleName *string `field:"required" json:"roleName" yaml:"roleName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/azure_secret_backend_role#scope AzureSecretBackendRole#scope}.
	Scope *string `field:"required" json:"scope" yaml:"scope"`
}

