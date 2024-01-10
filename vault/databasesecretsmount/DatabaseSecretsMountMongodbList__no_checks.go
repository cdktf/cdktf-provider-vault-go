// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package databasesecretsmount

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DatabaseSecretsMountMongodbList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DatabaseSecretsMountMongodbList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DatabaseSecretsMountMongodbList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DatabaseSecretsMountMongodbList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DatabaseSecretsMountMongodbList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DatabaseSecretsMountMongodbList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DatabaseSecretsMountMongodbList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDatabaseSecretsMountMongodbListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

