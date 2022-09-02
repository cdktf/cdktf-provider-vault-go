//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt vault Provider for Terraform CDK (cdktf)
package vault

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DatabaseSecretsMountCouchbaseList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DatabaseSecretsMountCouchbaseList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DatabaseSecretsMountCouchbaseList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DatabaseSecretsMountCouchbaseList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DatabaseSecretsMountCouchbaseList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DatabaseSecretsMountCouchbaseList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDatabaseSecretsMountCouchbaseListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}
