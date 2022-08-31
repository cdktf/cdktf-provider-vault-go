//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt vault Provider for Terraform CDK (cdktf)
package vault

// Building without runtime type checking enabled, so all the below just return nil

func (k *jsiiProxy_KvSecret) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (k *jsiiProxy_KvSecret) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (k *jsiiProxy_KvSecret) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (k *jsiiProxy_KvSecret) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (k *jsiiProxy_KvSecret) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (k *jsiiProxy_KvSecret) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (k *jsiiProxy_KvSecret) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (k *jsiiProxy_KvSecret) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (k *jsiiProxy_KvSecret) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (k *jsiiProxy_KvSecret) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (k *jsiiProxy_KvSecret) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (k *jsiiProxy_KvSecret) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateKvSecret_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_KvSecret) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_KvSecret) validateSetDataJsonParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_KvSecret) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_KvSecret) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_KvSecret) validateSetNamespaceParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_KvSecret) validateSetPathParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_KvSecret) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func validateNewKvSecretParameters(scope constructs.Construct, id *string, config *KvSecretConfig) error {
	return nil
}

