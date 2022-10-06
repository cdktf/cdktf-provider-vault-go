package managedkeys

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-vault.managedKeys.ManagedKeys",
		reflect.TypeOf((*ManagedKeys)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "aws", GoGetter: "Aws"},
			_jsii_.MemberProperty{JsiiProperty: "awsInput", GoGetter: "AwsInput"},
			_jsii_.MemberProperty{JsiiProperty: "azure", GoGetter: "Azure"},
			_jsii_.MemberProperty{JsiiProperty: "azureInput", GoGetter: "AzureInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "namespace", GoGetter: "Namespace"},
			_jsii_.MemberProperty{JsiiProperty: "namespaceInput", GoGetter: "NamespaceInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "pkcs", GoGetter: "Pkcs"},
			_jsii_.MemberProperty{JsiiProperty: "pkcsInput", GoGetter: "PkcsInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putAws", GoMethod: "PutAws"},
			_jsii_.MemberMethod{JsiiMethod: "putAzure", GoMethod: "PutAzure"},
			_jsii_.MemberMethod{JsiiMethod: "putPkcs", GoMethod: "PutPkcs"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetAws", GoMethod: "ResetAws"},
			_jsii_.MemberMethod{JsiiMethod: "resetAzure", GoMethod: "ResetAzure"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetNamespace", GoMethod: "ResetNamespace"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPkcs", GoMethod: "ResetPkcs"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_ManagedKeys{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-vault.managedKeys.ManagedKeysAws",
		reflect.TypeOf((*ManagedKeysAws)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-vault.managedKeys.ManagedKeysAwsList",
		reflect.TypeOf((*ManagedKeysAwsList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_ManagedKeysAwsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-vault.managedKeys.ManagedKeysAwsOutputReference",
		reflect.TypeOf((*ManagedKeysAwsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accessKey", GoGetter: "AccessKey"},
			_jsii_.MemberProperty{JsiiProperty: "accessKeyInput", GoGetter: "AccessKeyInput"},
			_jsii_.MemberProperty{JsiiProperty: "allowGenerateKey", GoGetter: "AllowGenerateKey"},
			_jsii_.MemberProperty{JsiiProperty: "allowGenerateKeyInput", GoGetter: "AllowGenerateKeyInput"},
			_jsii_.MemberProperty{JsiiProperty: "allowReplaceKey", GoGetter: "AllowReplaceKey"},
			_jsii_.MemberProperty{JsiiProperty: "allowReplaceKeyInput", GoGetter: "AllowReplaceKeyInput"},
			_jsii_.MemberProperty{JsiiProperty: "allowStoreKey", GoGetter: "AllowStoreKey"},
			_jsii_.MemberProperty{JsiiProperty: "allowStoreKeyInput", GoGetter: "AllowStoreKeyInput"},
			_jsii_.MemberProperty{JsiiProperty: "anyMount", GoGetter: "AnyMount"},
			_jsii_.MemberProperty{JsiiProperty: "anyMountInput", GoGetter: "AnyMountInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "curve", GoGetter: "Curve"},
			_jsii_.MemberProperty{JsiiProperty: "curveInput", GoGetter: "CurveInput"},
			_jsii_.MemberProperty{JsiiProperty: "endpoint", GoGetter: "Endpoint"},
			_jsii_.MemberProperty{JsiiProperty: "endpointInput", GoGetter: "EndpointInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "keyBits", GoGetter: "KeyBits"},
			_jsii_.MemberProperty{JsiiProperty: "keyBitsInput", GoGetter: "KeyBitsInput"},
			_jsii_.MemberProperty{JsiiProperty: "keyType", GoGetter: "KeyType"},
			_jsii_.MemberProperty{JsiiProperty: "keyTypeInput", GoGetter: "KeyTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKey", GoGetter: "KmsKey"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKeyInput", GoGetter: "KmsKeyInput"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "region", GoGetter: "Region"},
			_jsii_.MemberProperty{JsiiProperty: "regionInput", GoGetter: "RegionInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowGenerateKey", GoMethod: "ResetAllowGenerateKey"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowReplaceKey", GoMethod: "ResetAllowReplaceKey"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowStoreKey", GoMethod: "ResetAllowStoreKey"},
			_jsii_.MemberMethod{JsiiMethod: "resetAnyMount", GoMethod: "ResetAnyMount"},
			_jsii_.MemberMethod{JsiiMethod: "resetCurve", GoMethod: "ResetCurve"},
			_jsii_.MemberMethod{JsiiMethod: "resetEndpoint", GoMethod: "ResetEndpoint"},
			_jsii_.MemberMethod{JsiiMethod: "resetRegion", GoMethod: "ResetRegion"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "secretKey", GoGetter: "SecretKey"},
			_jsii_.MemberProperty{JsiiProperty: "secretKeyInput", GoGetter: "SecretKeyInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "uuid", GoGetter: "Uuid"},
		},
		func() interface{} {
			j := jsiiProxy_ManagedKeysAwsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-vault.managedKeys.ManagedKeysAzure",
		reflect.TypeOf((*ManagedKeysAzure)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-vault.managedKeys.ManagedKeysAzureList",
		reflect.TypeOf((*ManagedKeysAzureList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_ManagedKeysAzureList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-vault.managedKeys.ManagedKeysAzureOutputReference",
		reflect.TypeOf((*ManagedKeysAzureOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allowGenerateKey", GoGetter: "AllowGenerateKey"},
			_jsii_.MemberProperty{JsiiProperty: "allowGenerateKeyInput", GoGetter: "AllowGenerateKeyInput"},
			_jsii_.MemberProperty{JsiiProperty: "allowReplaceKey", GoGetter: "AllowReplaceKey"},
			_jsii_.MemberProperty{JsiiProperty: "allowReplaceKeyInput", GoGetter: "AllowReplaceKeyInput"},
			_jsii_.MemberProperty{JsiiProperty: "allowStoreKey", GoGetter: "AllowStoreKey"},
			_jsii_.MemberProperty{JsiiProperty: "allowStoreKeyInput", GoGetter: "AllowStoreKeyInput"},
			_jsii_.MemberProperty{JsiiProperty: "anyMount", GoGetter: "AnyMount"},
			_jsii_.MemberProperty{JsiiProperty: "anyMountInput", GoGetter: "AnyMountInput"},
			_jsii_.MemberProperty{JsiiProperty: "clientId", GoGetter: "ClientId"},
			_jsii_.MemberProperty{JsiiProperty: "clientIdInput", GoGetter: "ClientIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "clientSecret", GoGetter: "ClientSecret"},
			_jsii_.MemberProperty{JsiiProperty: "clientSecretInput", GoGetter: "ClientSecretInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "environment", GoGetter: "Environment"},
			_jsii_.MemberProperty{JsiiProperty: "environmentInput", GoGetter: "EnvironmentInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "keyBits", GoGetter: "KeyBits"},
			_jsii_.MemberProperty{JsiiProperty: "keyBitsInput", GoGetter: "KeyBitsInput"},
			_jsii_.MemberProperty{JsiiProperty: "keyName", GoGetter: "KeyName"},
			_jsii_.MemberProperty{JsiiProperty: "keyNameInput", GoGetter: "KeyNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "keyType", GoGetter: "KeyType"},
			_jsii_.MemberProperty{JsiiProperty: "keyTypeInput", GoGetter: "KeyTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowGenerateKey", GoMethod: "ResetAllowGenerateKey"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowReplaceKey", GoMethod: "ResetAllowReplaceKey"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowStoreKey", GoMethod: "ResetAllowStoreKey"},
			_jsii_.MemberMethod{JsiiMethod: "resetAnyMount", GoMethod: "ResetAnyMount"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnvironment", GoMethod: "ResetEnvironment"},
			_jsii_.MemberMethod{JsiiMethod: "resetKeyBits", GoMethod: "ResetKeyBits"},
			_jsii_.MemberMethod{JsiiMethod: "resetResource", GoMethod: "ResetResource"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "resource", GoGetter: "Resource"},
			_jsii_.MemberProperty{JsiiProperty: "resourceInput", GoGetter: "ResourceInput"},
			_jsii_.MemberProperty{JsiiProperty: "tenantId", GoGetter: "TenantId"},
			_jsii_.MemberProperty{JsiiProperty: "tenantIdInput", GoGetter: "TenantIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "uuid", GoGetter: "Uuid"},
			_jsii_.MemberProperty{JsiiProperty: "vaultName", GoGetter: "VaultName"},
			_jsii_.MemberProperty{JsiiProperty: "vaultNameInput", GoGetter: "VaultNameInput"},
		},
		func() interface{} {
			j := jsiiProxy_ManagedKeysAzureOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-vault.managedKeys.ManagedKeysConfig",
		reflect.TypeOf((*ManagedKeysConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-vault.managedKeys.ManagedKeysPkcs",
		reflect.TypeOf((*ManagedKeysPkcs)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-vault.managedKeys.ManagedKeysPkcsList",
		reflect.TypeOf((*ManagedKeysPkcsList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_ManagedKeysPkcsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-vault.managedKeys.ManagedKeysPkcsOutputReference",
		reflect.TypeOf((*ManagedKeysPkcsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allowGenerateKey", GoGetter: "AllowGenerateKey"},
			_jsii_.MemberProperty{JsiiProperty: "allowGenerateKeyInput", GoGetter: "AllowGenerateKeyInput"},
			_jsii_.MemberProperty{JsiiProperty: "allowReplaceKey", GoGetter: "AllowReplaceKey"},
			_jsii_.MemberProperty{JsiiProperty: "allowReplaceKeyInput", GoGetter: "AllowReplaceKeyInput"},
			_jsii_.MemberProperty{JsiiProperty: "allowStoreKey", GoGetter: "AllowStoreKey"},
			_jsii_.MemberProperty{JsiiProperty: "allowStoreKeyInput", GoGetter: "AllowStoreKeyInput"},
			_jsii_.MemberProperty{JsiiProperty: "anyMount", GoGetter: "AnyMount"},
			_jsii_.MemberProperty{JsiiProperty: "anyMountInput", GoGetter: "AnyMountInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "curve", GoGetter: "Curve"},
			_jsii_.MemberProperty{JsiiProperty: "curveInput", GoGetter: "CurveInput"},
			_jsii_.MemberProperty{JsiiProperty: "forceRwSession", GoGetter: "ForceRwSession"},
			_jsii_.MemberProperty{JsiiProperty: "forceRwSessionInput", GoGetter: "ForceRwSessionInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "keyBits", GoGetter: "KeyBits"},
			_jsii_.MemberProperty{JsiiProperty: "keyBitsInput", GoGetter: "KeyBitsInput"},
			_jsii_.MemberProperty{JsiiProperty: "keyId", GoGetter: "KeyId"},
			_jsii_.MemberProperty{JsiiProperty: "keyIdInput", GoGetter: "KeyIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "keyLabel", GoGetter: "KeyLabel"},
			_jsii_.MemberProperty{JsiiProperty: "keyLabelInput", GoGetter: "KeyLabelInput"},
			_jsii_.MemberProperty{JsiiProperty: "library", GoGetter: "Library"},
			_jsii_.MemberProperty{JsiiProperty: "libraryInput", GoGetter: "LibraryInput"},
			_jsii_.MemberProperty{JsiiProperty: "mechanism", GoGetter: "Mechanism"},
			_jsii_.MemberProperty{JsiiProperty: "mechanismInput", GoGetter: "MechanismInput"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "pin", GoGetter: "Pin"},
			_jsii_.MemberProperty{JsiiProperty: "pinInput", GoGetter: "PinInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowGenerateKey", GoMethod: "ResetAllowGenerateKey"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowReplaceKey", GoMethod: "ResetAllowReplaceKey"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowStoreKey", GoMethod: "ResetAllowStoreKey"},
			_jsii_.MemberMethod{JsiiMethod: "resetAnyMount", GoMethod: "ResetAnyMount"},
			_jsii_.MemberMethod{JsiiMethod: "resetCurve", GoMethod: "ResetCurve"},
			_jsii_.MemberMethod{JsiiMethod: "resetForceRwSession", GoMethod: "ResetForceRwSession"},
			_jsii_.MemberMethod{JsiiMethod: "resetKeyBits", GoMethod: "ResetKeyBits"},
			_jsii_.MemberMethod{JsiiMethod: "resetSlot", GoMethod: "ResetSlot"},
			_jsii_.MemberMethod{JsiiMethod: "resetTokenLabel", GoMethod: "ResetTokenLabel"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "slot", GoGetter: "Slot"},
			_jsii_.MemberProperty{JsiiProperty: "slotInput", GoGetter: "SlotInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "tokenLabel", GoGetter: "TokenLabel"},
			_jsii_.MemberProperty{JsiiProperty: "tokenLabelInput", GoGetter: "TokenLabelInput"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "uuid", GoGetter: "Uuid"},
		},
		func() interface{} {
			j := jsiiProxy_ManagedKeysPkcsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}