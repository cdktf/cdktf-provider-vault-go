// Prebuilt vault Provider for Terraform CDK (cdktf)
package vault


type AzureSecretBackendRoleAzureGroups struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/azure_secret_backend_role#group_name AzureSecretBackendRole#group_name}.
	GroupName *string `field:"required" json:"groupName" yaml:"groupName"`
}

