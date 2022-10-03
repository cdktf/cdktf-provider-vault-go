package databasesecretsmount


type DatabaseSecretsMountInfluxdb struct {
	// Influxdb host to connect to.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/database_secrets_mount#host DatabaseSecretsMount#host}
	Host *string `field:"required" json:"host" yaml:"host"`
	// Name of the database connection.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/database_secrets_mount#name DatabaseSecretsMount#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Specifies the password corresponding to the given username.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/database_secrets_mount#password DatabaseSecretsMount#password}
	Password *string `field:"required" json:"password" yaml:"password"`
	// Specifies the username to use for superuser access.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/database_secrets_mount#username DatabaseSecretsMount#username}
	Username *string `field:"required" json:"username" yaml:"username"`
	// A list of roles that are allowed to use this connection.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/database_secrets_mount#allowed_roles DatabaseSecretsMount#allowed_roles}
	AllowedRoles *[]*string `field:"optional" json:"allowedRoles" yaml:"allowedRoles"`
	// The number of seconds to use as a connection timeout.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/database_secrets_mount#connect_timeout DatabaseSecretsMount#connect_timeout}
	ConnectTimeout *float64 `field:"optional" json:"connectTimeout" yaml:"connectTimeout"`
	// A map of sensitive data to pass to the endpoint. Useful for templated connection strings.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/database_secrets_mount#data DatabaseSecretsMount#data}
	Data *map[string]*string `field:"optional" json:"data" yaml:"data"`
	// Whether to skip verification of the server certificate when using TLS.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/database_secrets_mount#insecure_tls DatabaseSecretsMount#insecure_tls}
	InsecureTls interface{} `field:"optional" json:"insecureTls" yaml:"insecureTls"`
	// Concatenated PEM blocks containing a certificate and private key;
	//
	// a certificate, private key, and issuing CA certificate; or just a CA certificate.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/database_secrets_mount#pem_bundle DatabaseSecretsMount#pem_bundle}
	PemBundle *string `field:"optional" json:"pemBundle" yaml:"pemBundle"`
	// Specifies JSON containing a certificate and private key;
	//
	// a certificate, private key, and issuing CA certificate; or just a CA certificate.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/database_secrets_mount#pem_json DatabaseSecretsMount#pem_json}
	PemJson *string `field:"optional" json:"pemJson" yaml:"pemJson"`
	// Specifies the name of the plugin to use for this connection.
	//
	// Must be prefixed with the name of one of the supported database engine types.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/database_secrets_mount#plugin_name DatabaseSecretsMount#plugin_name}
	PluginName *string `field:"optional" json:"pluginName" yaml:"pluginName"`
	// The transport port to use to connect to Influxdb.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/database_secrets_mount#port DatabaseSecretsMount#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// A list of database statements to be executed to rotate the root user's credentials.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/database_secrets_mount#root_rotation_statements DatabaseSecretsMount#root_rotation_statements}
	RootRotationStatements *[]*string `field:"optional" json:"rootRotationStatements" yaml:"rootRotationStatements"`
	// Whether to use TLS when connecting to Influxdb.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/database_secrets_mount#tls DatabaseSecretsMount#tls}
	Tls interface{} `field:"optional" json:"tls" yaml:"tls"`
	// Template describing how dynamic usernames are generated.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/database_secrets_mount#username_template DatabaseSecretsMount#username_template}
	UsernameTemplate *string `field:"optional" json:"usernameTemplate" yaml:"usernameTemplate"`
	// Specifies if the connection is verified during initial configuration.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault/r/database_secrets_mount#verify_connection DatabaseSecretsMount#verify_connection}
	VerifyConnection interface{} `field:"optional" json:"verifyConnection" yaml:"verifyConnection"`
}
