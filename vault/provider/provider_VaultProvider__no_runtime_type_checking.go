//go:build no_runtime_type_checking
// +build no_runtime_type_checking

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

func (j *jsiiProxy_VaultProvider) validateSetAuthLoginParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_VaultProvider) validateSetClientAuthParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_VaultProvider) validateSetHeadersParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_VaultProvider) validateSetSkipChildTokenParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_VaultProvider) validateSetSkipTlsVerifyParameters(val interface{}) error {
	return nil
}

func validateNewVaultProviderParameters(scope constructs.Construct, id *string, config *VaultProviderConfig) error {
	return nil
}

