// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rabbitmqsecretbackendrole


type RabbitmqSecretBackendRoleVhostTopic struct {
	// The vhost to set permissions for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.2.0/docs/resources/rabbitmq_secret_backend_role#host RabbitmqSecretBackendRole#host}
	Host *string `field:"required" json:"host" yaml:"host"`
	// vhost block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.2.0/docs/resources/rabbitmq_secret_backend_role#vhost RabbitmqSecretBackendRole#vhost}
	Vhost interface{} `field:"optional" json:"vhost" yaml:"vhost"`
}

