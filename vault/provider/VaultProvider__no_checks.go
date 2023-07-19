//go:build no_runtime_type_checking

package provider

// Building without runtime type checking enabled, so all the below just return nil

func (v *jsiiProxy_VaultProvider) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (v *jsiiProxy_VaultProvider) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateVaultProvider_IsConstructParameters(x interface{}) error {
	return nil
}

func validateVaultProvider_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateVaultProvider_IsTerraformProviderParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_VaultProvider) validateSetAuthLoginParameters(val *VaultProviderAuthLogin) error {
	return nil
}

func (j *jsiiProxy_VaultProvider) validateSetAuthLoginAwsParameters(val *VaultProviderAuthLoginAws) error {
	return nil
}

func (j *jsiiProxy_VaultProvider) validateSetAuthLoginAzureParameters(val *VaultProviderAuthLoginAzure) error {
	return nil
}

func (j *jsiiProxy_VaultProvider) validateSetAuthLoginCertParameters(val *VaultProviderAuthLoginCert) error {
	return nil
}

func (j *jsiiProxy_VaultProvider) validateSetAuthLoginGcpParameters(val *VaultProviderAuthLoginGcp) error {
	return nil
}

func (j *jsiiProxy_VaultProvider) validateSetAuthLoginJwtParameters(val *VaultProviderAuthLoginJwt) error {
	return nil
}

func (j *jsiiProxy_VaultProvider) validateSetAuthLoginKerberosParameters(val *VaultProviderAuthLoginKerberos) error {
	return nil
}

func (j *jsiiProxy_VaultProvider) validateSetAuthLoginOciParameters(val *VaultProviderAuthLoginOci) error {
	return nil
}

func (j *jsiiProxy_VaultProvider) validateSetAuthLoginOidcParameters(val *VaultProviderAuthLoginOidc) error {
	return nil
}

func (j *jsiiProxy_VaultProvider) validateSetAuthLoginRadiusParameters(val *VaultProviderAuthLoginRadius) error {
	return nil
}

func (j *jsiiProxy_VaultProvider) validateSetAuthLoginTokenFileParameters(val *VaultProviderAuthLoginTokenFile) error {
	return nil
}

func (j *jsiiProxy_VaultProvider) validateSetAuthLoginUserpassParameters(val *VaultProviderAuthLoginUserpass) error {
	return nil
}

func (j *jsiiProxy_VaultProvider) validateSetClientAuthParameters(val *VaultProviderClientAuth) error {
	return nil
}

func (j *jsiiProxy_VaultProvider) validateSetHeadersParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_VaultProvider) validateSetSkipChildTokenParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_VaultProvider) validateSetSkipGetVaultVersionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_VaultProvider) validateSetSkipTlsVerifyParameters(val interface{}) error {
	return nil
}

func validateNewVaultProviderParameters(scope constructs.Construct, id *string, config *VaultProviderConfig) error {
	return nil
}

