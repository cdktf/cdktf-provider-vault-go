package databasesecretbackendconnection


type DatabaseSecretBackendConnectionHana struct {
	// Connection string to use to connect to the database.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/database_secret_backend_connection#connection_url DatabaseSecretBackendConnection#connection_url}
	ConnectionUrl *string `field:"optional" json:"connectionUrl" yaml:"connectionUrl"`
	// Disable special character escaping in username and password.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/database_secret_backend_connection#disable_escaping DatabaseSecretBackendConnection#disable_escaping}
	DisableEscaping interface{} `field:"optional" json:"disableEscaping" yaml:"disableEscaping"`
	// Maximum number of seconds a connection may be reused.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/database_secret_backend_connection#max_connection_lifetime DatabaseSecretBackendConnection#max_connection_lifetime}
	MaxConnectionLifetime *float64 `field:"optional" json:"maxConnectionLifetime" yaml:"maxConnectionLifetime"`
	// Maximum number of idle connections to the database.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/database_secret_backend_connection#max_idle_connections DatabaseSecretBackendConnection#max_idle_connections}
	MaxIdleConnections *float64 `field:"optional" json:"maxIdleConnections" yaml:"maxIdleConnections"`
	// Maximum number of open connections to the database.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/database_secret_backend_connection#max_open_connections DatabaseSecretBackendConnection#max_open_connections}
	MaxOpenConnections *float64 `field:"optional" json:"maxOpenConnections" yaml:"maxOpenConnections"`
	// The root credential password used in the connection URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/database_secret_backend_connection#password DatabaseSecretBackendConnection#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// The root credential username used in the connection URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/database_secret_backend_connection#username DatabaseSecretBackendConnection#username}
	Username *string `field:"optional" json:"username" yaml:"username"`
}
