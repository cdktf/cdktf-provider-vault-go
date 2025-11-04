// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package databasesecretsmount

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DatabaseSecretsMountPostgresqlList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DatabaseSecretsMountPostgresqlList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DatabaseSecretsMountPostgresqlList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DatabaseSecretsMountPostgresqlList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DatabaseSecretsMountPostgresqlList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DatabaseSecretsMountPostgresqlList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DatabaseSecretsMountPostgresqlList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDatabaseSecretsMountPostgresqlListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

