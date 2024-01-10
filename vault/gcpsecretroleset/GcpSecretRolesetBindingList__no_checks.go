// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package gcpsecretroleset

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GcpSecretRolesetBindingList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (g *jsiiProxy_GcpSecretRolesetBindingList) validateGetParameters(index *float64) error {
	return nil
}

func (g *jsiiProxy_GcpSecretRolesetBindingList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_GcpSecretRolesetBindingList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_GcpSecretRolesetBindingList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_GcpSecretRolesetBindingList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_GcpSecretRolesetBindingList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewGcpSecretRolesetBindingListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

