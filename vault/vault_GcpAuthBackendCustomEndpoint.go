// Prebuilt vault Provider for Terraform CDK (cdktf)
package vault


type GcpAuthBackendCustomEndpoint struct {
	// Replaces the service endpoint used in API requests to https://www.googleapis.com.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/gcp_auth_backend#api GcpAuthBackend#api}
	Api *string `field:"optional" json:"api" yaml:"api"`
	// Replaces the service endpoint used in API requests to `https://compute.googleapis.com`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/gcp_auth_backend#compute GcpAuthBackend#compute}
	Compute *string `field:"optional" json:"compute" yaml:"compute"`
	// Replaces the service endpoint used in API requests to `https://cloudresourcemanager.googleapis.com`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/gcp_auth_backend#crm GcpAuthBackend#crm}
	Crm *string `field:"optional" json:"crm" yaml:"crm"`
	// Replaces the service endpoint used in API requests to `https://iam.googleapis.com`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/gcp_auth_backend#iam GcpAuthBackend#iam}
	Iam *string `field:"optional" json:"iam" yaml:"iam"`
}

