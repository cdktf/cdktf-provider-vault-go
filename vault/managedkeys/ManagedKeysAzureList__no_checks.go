// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package managedkeys

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_ManagedKeysAzureList) validateGetParameters(index *float64) error {
	return nil
}

func (m *jsiiProxy_ManagedKeysAzureList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ManagedKeysAzureList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ManagedKeysAzureList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ManagedKeysAzureList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ManagedKeysAzureList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewManagedKeysAzureListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

