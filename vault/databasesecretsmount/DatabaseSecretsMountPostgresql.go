// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package databasesecretsmount


type DatabaseSecretsMountPostgresql struct {
	// Name of the database connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.5.0/docs/resources/database_secrets_mount#name DatabaseSecretsMount#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A list of roles that are allowed to use this connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.5.0/docs/resources/database_secrets_mount#allowed_roles DatabaseSecretsMount#allowed_roles}
	AllowedRoles *[]*string `field:"optional" json:"allowedRoles" yaml:"allowedRoles"`
	// Specify alternative authorization type. (Only 'gcp_iam' is valid currently).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.5.0/docs/resources/database_secrets_mount#auth_type DatabaseSecretsMount#auth_type}
	AuthType *string `field:"optional" json:"authType" yaml:"authType"`
	// Connection string to use to connect to the database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.5.0/docs/resources/database_secrets_mount#connection_url DatabaseSecretsMount#connection_url}
	ConnectionUrl *string `field:"optional" json:"connectionUrl" yaml:"connectionUrl"`
	// A map of sensitive data to pass to the endpoint. Useful for templated connection strings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.5.0/docs/resources/database_secrets_mount#data DatabaseSecretsMount#data}
	Data *map[string]*string `field:"optional" json:"data" yaml:"data"`
	// Disable special character escaping in username and password.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.5.0/docs/resources/database_secrets_mount#disable_escaping DatabaseSecretsMount#disable_escaping}
	DisableEscaping interface{} `field:"optional" json:"disableEscaping" yaml:"disableEscaping"`
	// Maximum number of seconds a connection may be reused.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.5.0/docs/resources/database_secrets_mount#max_connection_lifetime DatabaseSecretsMount#max_connection_lifetime}
	MaxConnectionLifetime *float64 `field:"optional" json:"maxConnectionLifetime" yaml:"maxConnectionLifetime"`
	// Maximum number of idle connections to the database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.5.0/docs/resources/database_secrets_mount#max_idle_connections DatabaseSecretsMount#max_idle_connections}
	MaxIdleConnections *float64 `field:"optional" json:"maxIdleConnections" yaml:"maxIdleConnections"`
	// Maximum number of open connections to the database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.5.0/docs/resources/database_secrets_mount#max_open_connections DatabaseSecretsMount#max_open_connections}
	MaxOpenConnections *float64 `field:"optional" json:"maxOpenConnections" yaml:"maxOpenConnections"`
	// The root credential password used in the connection URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.5.0/docs/resources/database_secrets_mount#password DatabaseSecretsMount#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Specifies the name of the plugin to use for this connection.
	//
	// Must be prefixed with the name of one of the supported database engine types.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.5.0/docs/resources/database_secrets_mount#plugin_name DatabaseSecretsMount#plugin_name}
	PluginName *string `field:"optional" json:"pluginName" yaml:"pluginName"`
	// The secret key used for the x509 client certificate. Must be PEM encoded.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.5.0/docs/resources/database_secrets_mount#private_key DatabaseSecretsMount#private_key}
	PrivateKey *string `field:"optional" json:"privateKey" yaml:"privateKey"`
	// A list of database statements to be executed to rotate the root user's credentials.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.5.0/docs/resources/database_secrets_mount#root_rotation_statements DatabaseSecretsMount#root_rotation_statements}
	RootRotationStatements *[]*string `field:"optional" json:"rootRotationStatements" yaml:"rootRotationStatements"`
	// If set, allows onboarding static roles with a rootless connection configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.5.0/docs/resources/database_secrets_mount#self_managed DatabaseSecretsMount#self_managed}
	SelfManaged interface{} `field:"optional" json:"selfManaged" yaml:"selfManaged"`
	// A JSON encoded credential for use with IAM authorization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.5.0/docs/resources/database_secrets_mount#service_account_json DatabaseSecretsMount#service_account_json}
	ServiceAccountJson *string `field:"optional" json:"serviceAccountJson" yaml:"serviceAccountJson"`
	// The x509 CA file for validating the certificate presented by the PostgreSQL server. Must be PEM encoded.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.5.0/docs/resources/database_secrets_mount#tls_ca DatabaseSecretsMount#tls_ca}
	TlsCa *string `field:"optional" json:"tlsCa" yaml:"tlsCa"`
	// The x509 client certificate for connecting to the database. Must be PEM encoded.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.5.0/docs/resources/database_secrets_mount#tls_certificate DatabaseSecretsMount#tls_certificate}
	TlsCertificate *string `field:"optional" json:"tlsCertificate" yaml:"tlsCertificate"`
	// The root credential username used in the connection URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.5.0/docs/resources/database_secrets_mount#username DatabaseSecretsMount#username}
	Username *string `field:"optional" json:"username" yaml:"username"`
	// Username generation template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.5.0/docs/resources/database_secrets_mount#username_template DatabaseSecretsMount#username_template}
	UsernameTemplate *string `field:"optional" json:"usernameTemplate" yaml:"usernameTemplate"`
	// Specifies if the connection is verified during initial configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/4.5.0/docs/resources/database_secrets_mount#verify_connection DatabaseSecretsMount#verify_connection}
	VerifyConnection interface{} `field:"optional" json:"verifyConnection" yaml:"verifyConnection"`
}

