package provider


type VaultProviderConfig struct {
	// URL of the root of the target Vault server.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault#address VaultProvider#address}
	Address *string `field:"required" json:"address" yaml:"address"`
	// If true, adds the value of the `address` argument to the Terraform process environment.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault#add_address_to_env VaultProvider#add_address_to_env}
	AddAddressToEnv *string `field:"optional" json:"addAddressToEnv" yaml:"addAddressToEnv"`
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault#alias VaultProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// auth_login block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault#auth_login VaultProvider#auth_login}
	AuthLogin *VaultProviderAuthLogin `field:"optional" json:"authLogin" yaml:"authLogin"`
	// auth_login_aws block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault#auth_login_aws VaultProvider#auth_login_aws}
	AuthLoginAws *VaultProviderAuthLoginAws `field:"optional" json:"authLoginAws" yaml:"authLoginAws"`
	// auth_login_azure block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault#auth_login_azure VaultProvider#auth_login_azure}
	AuthLoginAzure *VaultProviderAuthLoginAzure `field:"optional" json:"authLoginAzure" yaml:"authLoginAzure"`
	// auth_login_cert block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault#auth_login_cert VaultProvider#auth_login_cert}
	AuthLoginCert *VaultProviderAuthLoginCert `field:"optional" json:"authLoginCert" yaml:"authLoginCert"`
	// auth_login_gcp block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault#auth_login_gcp VaultProvider#auth_login_gcp}
	AuthLoginGcp *VaultProviderAuthLoginGcp `field:"optional" json:"authLoginGcp" yaml:"authLoginGcp"`
	// auth_login_jwt block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault#auth_login_jwt VaultProvider#auth_login_jwt}
	AuthLoginJwt *VaultProviderAuthLoginJwt `field:"optional" json:"authLoginJwt" yaml:"authLoginJwt"`
	// auth_login_kerberos block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault#auth_login_kerberos VaultProvider#auth_login_kerberos}
	AuthLoginKerberos *VaultProviderAuthLoginKerberos `field:"optional" json:"authLoginKerberos" yaml:"authLoginKerberos"`
	// auth_login_oci block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault#auth_login_oci VaultProvider#auth_login_oci}
	AuthLoginOci *VaultProviderAuthLoginOci `field:"optional" json:"authLoginOci" yaml:"authLoginOci"`
	// auth_login_oidc block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault#auth_login_oidc VaultProvider#auth_login_oidc}
	AuthLoginOidc *VaultProviderAuthLoginOidc `field:"optional" json:"authLoginOidc" yaml:"authLoginOidc"`
	// auth_login_radius block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault#auth_login_radius VaultProvider#auth_login_radius}
	AuthLoginRadius *VaultProviderAuthLoginRadius `field:"optional" json:"authLoginRadius" yaml:"authLoginRadius"`
	// auth_login_userpass block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault#auth_login_userpass VaultProvider#auth_login_userpass}
	AuthLoginUserpass *VaultProviderAuthLoginUserpass `field:"optional" json:"authLoginUserpass" yaml:"authLoginUserpass"`
	// Path to directory containing CA certificate files to validate the server's certificate.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault#ca_cert_dir VaultProvider#ca_cert_dir}
	CaCertDir *string `field:"optional" json:"caCertDir" yaml:"caCertDir"`
	// Path to a CA certificate file to validate the server's certificate.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault#ca_cert_file VaultProvider#ca_cert_file}
	CaCertFile *string `field:"optional" json:"caCertFile" yaml:"caCertFile"`
	// client_auth block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault#client_auth VaultProvider#client_auth}
	ClientAuth *VaultProviderClientAuth `field:"optional" json:"clientAuth" yaml:"clientAuth"`
	// headers block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault#headers VaultProvider#headers}
	Headers interface{} `field:"optional" json:"headers" yaml:"headers"`
	// Maximum TTL for secret leases requested by this provider.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault#max_lease_ttl_seconds VaultProvider#max_lease_ttl_seconds}
	MaxLeaseTtlSeconds *float64 `field:"optional" json:"maxLeaseTtlSeconds" yaml:"maxLeaseTtlSeconds"`
	// Maximum number of retries when a 5xx error code is encountered.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault#max_retries VaultProvider#max_retries}
	MaxRetries *float64 `field:"optional" json:"maxRetries" yaml:"maxRetries"`
	// Maximum number of retries for Client Controlled Consistency related operations.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault#max_retries_ccc VaultProvider#max_retries_ccc}
	MaxRetriesCcc *float64 `field:"optional" json:"maxRetriesCcc" yaml:"maxRetriesCcc"`
	// The namespace to use. Available only for Vault Enterprise.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault#namespace VaultProvider#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Set this to true to prevent the creation of ephemeral child token used by this provider.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault#skip_child_token VaultProvider#skip_child_token}
	SkipChildToken interface{} `field:"optional" json:"skipChildToken" yaml:"skipChildToken"`
	// Set this to true only if the target Vault server is an insecure development instance.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault#skip_tls_verify VaultProvider#skip_tls_verify}
	SkipTlsVerify interface{} `field:"optional" json:"skipTlsVerify" yaml:"skipTlsVerify"`
	// Name to use as the SNI host when connecting via TLS.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault#tls_server_name VaultProvider#tls_server_name}
	TlsServerName *string `field:"optional" json:"tlsServerName" yaml:"tlsServerName"`
	// Token to use to authenticate to Vault.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault#token VaultProvider#token}
	Token *string `field:"optional" json:"token" yaml:"token"`
	// Token name to use for creating the Vault child token.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/vault#token_name VaultProvider#token_name}
	TokenName *string `field:"optional" json:"tokenName" yaml:"tokenName"`
}

