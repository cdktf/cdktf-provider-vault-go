// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package databasesecretbackendconnection


type DatabaseSecretBackendConnectionOracle struct {
	// Connection string to use to connect to the database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.5.0/docs/resources/database_secret_backend_connection#connection_url DatabaseSecretBackendConnection#connection_url}
	ConnectionUrl *string `field:"optional" json:"connectionUrl" yaml:"connectionUrl"`
	// Set to true to disconnect any open sessions prior to running the revocation statements.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.5.0/docs/resources/database_secret_backend_connection#disconnect_sessions DatabaseSecretBackendConnection#disconnect_sessions}
	DisconnectSessions interface{} `field:"optional" json:"disconnectSessions" yaml:"disconnectSessions"`
	// Maximum number of seconds a connection may be reused.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.5.0/docs/resources/database_secret_backend_connection#max_connection_lifetime DatabaseSecretBackendConnection#max_connection_lifetime}
	MaxConnectionLifetime *float64 `field:"optional" json:"maxConnectionLifetime" yaml:"maxConnectionLifetime"`
	// Maximum number of idle connections to the database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.5.0/docs/resources/database_secret_backend_connection#max_idle_connections DatabaseSecretBackendConnection#max_idle_connections}
	MaxIdleConnections *float64 `field:"optional" json:"maxIdleConnections" yaml:"maxIdleConnections"`
	// Maximum number of open connections to the database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.5.0/docs/resources/database_secret_backend_connection#max_open_connections DatabaseSecretBackendConnection#max_open_connections}
	MaxOpenConnections *float64 `field:"optional" json:"maxOpenConnections" yaml:"maxOpenConnections"`
	// The root credential password used in the connection URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.5.0/docs/resources/database_secret_backend_connection#password DatabaseSecretBackendConnection#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Set to true in order to split statements after semi-colons.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.5.0/docs/resources/database_secret_backend_connection#split_statements DatabaseSecretBackendConnection#split_statements}
	SplitStatements interface{} `field:"optional" json:"splitStatements" yaml:"splitStatements"`
	// The root credential username used in the connection URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.5.0/docs/resources/database_secret_backend_connection#username DatabaseSecretBackendConnection#username}
	Username *string `field:"optional" json:"username" yaml:"username"`
	// Username generation template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.5.0/docs/resources/database_secret_backend_connection#username_template DatabaseSecretBackendConnection#username_template}
	UsernameTemplate *string `field:"optional" json:"usernameTemplate" yaml:"usernameTemplate"`
}

