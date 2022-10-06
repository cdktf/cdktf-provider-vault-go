//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package managedkeys

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_ManagedKeysAwsList) validateGetParameters(index *float64) error {
	return nil
}

func (m *jsiiProxy_ManagedKeysAwsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ManagedKeysAwsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ManagedKeysAwsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ManagedKeysAwsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ManagedKeysAwsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewManagedKeysAwsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}
