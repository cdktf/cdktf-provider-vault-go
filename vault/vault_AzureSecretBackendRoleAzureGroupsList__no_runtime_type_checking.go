//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt vault Provider for Terraform CDK (cdktf)
package vault

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AzureSecretBackendRoleAzureGroupsList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AzureSecretBackendRoleAzureGroupsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AzureSecretBackendRoleAzureGroupsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AzureSecretBackendRoleAzureGroupsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AzureSecretBackendRoleAzureGroupsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AzureSecretBackendRoleAzureGroupsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAzureSecretBackendRoleAzureGroupsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}
