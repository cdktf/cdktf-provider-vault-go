// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package ociauthbackend

// Building without runtime type checking enabled, so all the below just return nil

func (o *jsiiProxy_OciAuthBackendTuneList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (o *jsiiProxy_OciAuthBackendTuneList) validateGetParameters(index *float64) error {
	return nil
}

func (o *jsiiProxy_OciAuthBackendTuneList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_OciAuthBackendTuneList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_OciAuthBackendTuneList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_OciAuthBackendTuneList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_OciAuthBackendTuneList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewOciAuthBackendTuneListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

