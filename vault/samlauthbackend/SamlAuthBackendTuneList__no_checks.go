// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package samlauthbackend

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SamlAuthBackendTuneList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_SamlAuthBackendTuneList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SamlAuthBackendTuneList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SamlAuthBackendTuneList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SamlAuthBackendTuneList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SamlAuthBackendTuneList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SamlAuthBackendTuneList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSamlAuthBackendTuneListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

