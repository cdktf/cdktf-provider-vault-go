//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt vault Provider for Terraform CDK (cdktf)
package vault

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DatabaseSecretsMountMysqlRdsList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DatabaseSecretsMountMysqlRdsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DatabaseSecretsMountMysqlRdsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DatabaseSecretsMountMysqlRdsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DatabaseSecretsMountMysqlRdsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DatabaseSecretsMountMysqlRdsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDatabaseSecretsMountMysqlRdsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}
