//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt vault Provider for Terraform CDK (cdktf)
package vault

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataVaultIdentityEntityAliasesList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataVaultIdentityEntityAliasesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataVaultIdentityEntityAliasesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataVaultIdentityEntityAliasesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataVaultIdentityEntityAliasesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataVaultIdentityEntityAliasesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

