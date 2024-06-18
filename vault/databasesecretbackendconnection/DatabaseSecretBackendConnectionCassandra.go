// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package databasesecretbackendconnection


type DatabaseSecretBackendConnectionCassandra struct {
	// The number of seconds to use as a connection timeout.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.3.0/docs/resources/database_secret_backend_connection#connect_timeout DatabaseSecretBackendConnection#connect_timeout}
	ConnectTimeout *float64 `field:"optional" json:"connectTimeout" yaml:"connectTimeout"`
	// Cassandra hosts to connect to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.3.0/docs/resources/database_secret_backend_connection#hosts DatabaseSecretBackendConnection#hosts}
	Hosts *[]*string `field:"optional" json:"hosts" yaml:"hosts"`
	// Whether to skip verification of the server certificate when using TLS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.3.0/docs/resources/database_secret_backend_connection#insecure_tls DatabaseSecretBackendConnection#insecure_tls}
	InsecureTls interface{} `field:"optional" json:"insecureTls" yaml:"insecureTls"`
	// The password to use when authenticating with Cassandra.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.3.0/docs/resources/database_secret_backend_connection#password DatabaseSecretBackendConnection#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Concatenated PEM blocks containing a certificate and private key;
	//
	// a certificate, private key, and issuing CA certificate; or just a CA certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.3.0/docs/resources/database_secret_backend_connection#pem_bundle DatabaseSecretBackendConnection#pem_bundle}
	PemBundle *string `field:"optional" json:"pemBundle" yaml:"pemBundle"`
	// Specifies JSON containing a certificate and private key;
	//
	// a certificate, private key, and issuing CA certificate; or just a CA certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.3.0/docs/resources/database_secret_backend_connection#pem_json DatabaseSecretBackendConnection#pem_json}
	PemJson *string `field:"optional" json:"pemJson" yaml:"pemJson"`
	// The transport port to use to connect to Cassandra.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.3.0/docs/resources/database_secret_backend_connection#port DatabaseSecretBackendConnection#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// The CQL protocol version to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.3.0/docs/resources/database_secret_backend_connection#protocol_version DatabaseSecretBackendConnection#protocol_version}
	ProtocolVersion *float64 `field:"optional" json:"protocolVersion" yaml:"protocolVersion"`
	// Whether to use TLS when connecting to Cassandra.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.3.0/docs/resources/database_secret_backend_connection#tls DatabaseSecretBackendConnection#tls}
	Tls interface{} `field:"optional" json:"tls" yaml:"tls"`
	// The username to use when authenticating with Cassandra.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.3.0/docs/resources/database_secret_backend_connection#username DatabaseSecretBackendConnection#username}
	Username *string `field:"optional" json:"username" yaml:"username"`
}

