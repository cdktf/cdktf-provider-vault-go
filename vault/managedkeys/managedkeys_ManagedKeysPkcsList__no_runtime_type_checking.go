//go:build no_runtime_type_checking

package managedkeys

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_ManagedKeysPkcsList) validateGetParameters(index *float64) error {
	return nil
}

func (m *jsiiProxy_ManagedKeysPkcsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ManagedKeysPkcsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ManagedKeysPkcsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ManagedKeysPkcsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ManagedKeysPkcsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewManagedKeysPkcsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

