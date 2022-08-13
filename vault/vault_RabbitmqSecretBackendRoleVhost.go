// Prebuilt vault Provider for Terraform CDK (cdktf)
package vault


type RabbitmqSecretBackendRoleVhost struct {
	// The configure permissions for this vhost.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/rabbitmq_secret_backend_role#configure RabbitmqSecretBackendRole#configure}
	Configure *string `field:"required" json:"configure" yaml:"configure"`
	// The vhost to set permissions for.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/rabbitmq_secret_backend_role#host RabbitmqSecretBackendRole#host}
	Host *string `field:"required" json:"host" yaml:"host"`
	// The read permissions for this vhost.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/rabbitmq_secret_backend_role#read RabbitmqSecretBackendRole#read}
	Read *string `field:"required" json:"read" yaml:"read"`
	// The write permissions for this vhost.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/rabbitmq_secret_backend_role#write RabbitmqSecretBackendRole#write}
	Write *string `field:"required" json:"write" yaml:"write"`
}
