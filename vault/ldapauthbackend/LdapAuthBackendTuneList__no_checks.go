// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package ldapauthbackend

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LdapAuthBackendTuneList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (l *jsiiProxy_LdapAuthBackendTuneList) validateGetParameters(index *float64) error {
	return nil
}

func (l *jsiiProxy_LdapAuthBackendTuneList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_LdapAuthBackendTuneList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_LdapAuthBackendTuneList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LdapAuthBackendTuneList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_LdapAuthBackendTuneList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewLdapAuthBackendTuneListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

