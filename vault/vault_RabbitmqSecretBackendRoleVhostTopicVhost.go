// Prebuilt vault Provider for Terraform CDK (cdktf)
package vault


type RabbitmqSecretBackendRoleVhostTopicVhost struct {
	// The read permissions for this vhost.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/rabbitmq_secret_backend_role#read RabbitmqSecretBackendRole#read}
	Read *string `field:"required" json:"read" yaml:"read"`
	// The vhost to set permissions for.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/rabbitmq_secret_backend_role#topic RabbitmqSecretBackendRole#topic}
	Topic *string `field:"required" json:"topic" yaml:"topic"`
	// The write permissions for this vhost.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/rabbitmq_secret_backend_role#write RabbitmqSecretBackendRole#write}
	Write *string `field:"required" json:"write" yaml:"write"`
}

