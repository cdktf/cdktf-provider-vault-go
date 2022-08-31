//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt vault Provider for Terraform CDK (cdktf)
package vault

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SshSecretBackendRoleAllowedUserKeyConfigList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SshSecretBackendRoleAllowedUserKeyConfigList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SshSecretBackendRoleAllowedUserKeyConfigList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SshSecretBackendRoleAllowedUserKeyConfigList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SshSecretBackendRoleAllowedUserKeyConfigList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SshSecretBackendRoleAllowedUserKeyConfigList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSshSecretBackendRoleAllowedUserKeyConfigListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

