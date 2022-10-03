package rabbitmqsecretbackendrole


type RabbitmqSecretBackendRoleVhostTopic struct {
	// The vhost to set permissions for.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/rabbitmq_secret_backend_role#host RabbitmqSecretBackendRole#host}
	Host *string `field:"required" json:"host" yaml:"host"`
	// vhost block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/rabbitmq_secret_backend_role#vhost RabbitmqSecretBackendRole#vhost}
	Vhost interface{} `field:"optional" json:"vhost" yaml:"vhost"`
}
