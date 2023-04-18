package rabbitmqsecretbackendrole


type RabbitmqSecretBackendRoleVhostTopicVhost struct {
	// The read permissions for this vhost.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.15.0/docs/resources/rabbitmq_secret_backend_role#read RabbitmqSecretBackendRole#read}
	Read *string `field:"required" json:"read" yaml:"read"`
	// The vhost to set permissions for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.15.0/docs/resources/rabbitmq_secret_backend_role#topic RabbitmqSecretBackendRole#topic}
	Topic *string `field:"required" json:"topic" yaml:"topic"`
	// The write permissions for this vhost.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/3.15.0/docs/resources/rabbitmq_secret_backend_role#write RabbitmqSecretBackendRole#write}
	Write *string `field:"required" json:"write" yaml:"write"`
}
