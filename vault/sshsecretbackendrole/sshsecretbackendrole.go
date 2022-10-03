package sshsecretbackendrole

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-vault.sshSecretBackendRole.SshSecretBackendRole",
		reflect.TypeOf((*SshSecretBackendRole)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "algorithmSigner", GoGetter: "AlgorithmSigner"},
			_jsii_.MemberProperty{JsiiProperty: "algorithmSignerInput", GoGetter: "AlgorithmSignerInput"},
			_jsii_.MemberProperty{JsiiProperty: "allowBareDomains", GoGetter: "AllowBareDomains"},
			_jsii_.MemberProperty{JsiiProperty: "allowBareDomainsInput", GoGetter: "AllowBareDomainsInput"},
			_jsii_.MemberProperty{JsiiProperty: "allowedCriticalOptions", GoGetter: "AllowedCriticalOptions"},
			_jsii_.MemberProperty{JsiiProperty: "allowedCriticalOptionsInput", GoGetter: "AllowedCriticalOptionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "allowedDomains", GoGetter: "AllowedDomains"},
			_jsii_.MemberProperty{JsiiProperty: "allowedDomainsInput", GoGetter: "AllowedDomainsInput"},
			_jsii_.MemberProperty{JsiiProperty: "allowedExtensions", GoGetter: "AllowedExtensions"},
			_jsii_.MemberProperty{JsiiProperty: "allowedExtensionsInput", GoGetter: "AllowedExtensionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "allowedUserKeyConfig", GoGetter: "AllowedUserKeyConfig"},
			_jsii_.MemberProperty{JsiiProperty: "allowedUserKeyConfigInput", GoGetter: "AllowedUserKeyConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "allowedUserKeyLengths", GoGetter: "AllowedUserKeyLengths"},
			_jsii_.MemberProperty{JsiiProperty: "allowedUserKeyLengthsInput", GoGetter: "AllowedUserKeyLengthsInput"},
			_jsii_.MemberProperty{JsiiProperty: "allowedUsers", GoGetter: "AllowedUsers"},
			_jsii_.MemberProperty{JsiiProperty: "allowedUsersInput", GoGetter: "AllowedUsersInput"},
			_jsii_.MemberProperty{JsiiProperty: "allowedUsersTemplate", GoGetter: "AllowedUsersTemplate"},
			_jsii_.MemberProperty{JsiiProperty: "allowedUsersTemplateInput", GoGetter: "AllowedUsersTemplateInput"},
			_jsii_.MemberProperty{JsiiProperty: "allowHostCertificates", GoGetter: "AllowHostCertificates"},
			_jsii_.MemberProperty{JsiiProperty: "allowHostCertificatesInput", GoGetter: "AllowHostCertificatesInput"},
			_jsii_.MemberProperty{JsiiProperty: "allowSubdomains", GoGetter: "AllowSubdomains"},
			_jsii_.MemberProperty{JsiiProperty: "allowSubdomainsInput", GoGetter: "AllowSubdomainsInput"},
			_jsii_.MemberProperty{JsiiProperty: "allowUserCertificates", GoGetter: "AllowUserCertificates"},
			_jsii_.MemberProperty{JsiiProperty: "allowUserCertificatesInput", GoGetter: "AllowUserCertificatesInput"},
			_jsii_.MemberProperty{JsiiProperty: "allowUserKeyIds", GoGetter: "AllowUserKeyIds"},
			_jsii_.MemberProperty{JsiiProperty: "allowUserKeyIdsInput", GoGetter: "AllowUserKeyIdsInput"},
			_jsii_.MemberProperty{JsiiProperty: "backend", GoGetter: "Backend"},
			_jsii_.MemberProperty{JsiiProperty: "backendInput", GoGetter: "BackendInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "cidrList", GoGetter: "CidrList"},
			_jsii_.MemberProperty{JsiiProperty: "cidrListInput", GoGetter: "CidrListInput"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "defaultCriticalOptions", GoGetter: "DefaultCriticalOptions"},
			_jsii_.MemberProperty{JsiiProperty: "defaultCriticalOptionsInput", GoGetter: "DefaultCriticalOptionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "defaultExtensions", GoGetter: "DefaultExtensions"},
			_jsii_.MemberProperty{JsiiProperty: "defaultExtensionsInput", GoGetter: "DefaultExtensionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "defaultUser", GoGetter: "DefaultUser"},
			_jsii_.MemberProperty{JsiiProperty: "defaultUserInput", GoGetter: "DefaultUserInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "keyIdFormat", GoGetter: "KeyIdFormat"},
			_jsii_.MemberProperty{JsiiProperty: "keyIdFormatInput", GoGetter: "KeyIdFormatInput"},
			_jsii_.MemberProperty{JsiiProperty: "keyType", GoGetter: "KeyType"},
			_jsii_.MemberProperty{JsiiProperty: "keyTypeInput", GoGetter: "KeyTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "maxTtl", GoGetter: "MaxTtl"},
			_jsii_.MemberProperty{JsiiProperty: "maxTtlInput", GoGetter: "MaxTtlInput"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "namespace", GoGetter: "Namespace"},
			_jsii_.MemberProperty{JsiiProperty: "namespaceInput", GoGetter: "NamespaceInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putAllowedUserKeyConfig", GoMethod: "PutAllowedUserKeyConfig"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetAlgorithmSigner", GoMethod: "ResetAlgorithmSigner"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowBareDomains", GoMethod: "ResetAllowBareDomains"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowedCriticalOptions", GoMethod: "ResetAllowedCriticalOptions"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowedDomains", GoMethod: "ResetAllowedDomains"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowedExtensions", GoMethod: "ResetAllowedExtensions"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowedUserKeyConfig", GoMethod: "ResetAllowedUserKeyConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowedUserKeyLengths", GoMethod: "ResetAllowedUserKeyLengths"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowedUsers", GoMethod: "ResetAllowedUsers"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowedUsersTemplate", GoMethod: "ResetAllowedUsersTemplate"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowHostCertificates", GoMethod: "ResetAllowHostCertificates"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowSubdomains", GoMethod: "ResetAllowSubdomains"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowUserCertificates", GoMethod: "ResetAllowUserCertificates"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowUserKeyIds", GoMethod: "ResetAllowUserKeyIds"},
			_jsii_.MemberMethod{JsiiMethod: "resetCidrList", GoMethod: "ResetCidrList"},
			_jsii_.MemberMethod{JsiiMethod: "resetDefaultCriticalOptions", GoMethod: "ResetDefaultCriticalOptions"},
			_jsii_.MemberMethod{JsiiMethod: "resetDefaultExtensions", GoMethod: "ResetDefaultExtensions"},
			_jsii_.MemberMethod{JsiiMethod: "resetDefaultUser", GoMethod: "ResetDefaultUser"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetKeyIdFormat", GoMethod: "ResetKeyIdFormat"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxTtl", GoMethod: "ResetMaxTtl"},
			_jsii_.MemberMethod{JsiiMethod: "resetNamespace", GoMethod: "ResetNamespace"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetTtl", GoMethod: "ResetTtl"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "ttl", GoGetter: "Ttl"},
			_jsii_.MemberProperty{JsiiProperty: "ttlInput", GoGetter: "TtlInput"},
		},
		func() interface{} {
			j := jsiiProxy_SshSecretBackendRole{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-vault.sshSecretBackendRole.SshSecretBackendRoleAllowedUserKeyConfig",
		reflect.TypeOf((*SshSecretBackendRoleAllowedUserKeyConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-vault.sshSecretBackendRole.SshSecretBackendRoleAllowedUserKeyConfigList",
		reflect.TypeOf((*SshSecretBackendRoleAllowedUserKeyConfigList)(nil)).Elem(),
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
			j := jsiiProxy_SshSecretBackendRoleAllowedUserKeyConfigList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-vault.sshSecretBackendRole.SshSecretBackendRoleAllowedUserKeyConfigOutputReference",
		reflect.TypeOf((*SshSecretBackendRoleAllowedUserKeyConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "lengths", GoGetter: "Lengths"},
			_jsii_.MemberProperty{JsiiProperty: "lengthsInput", GoGetter: "LengthsInput"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "typeInput", GoGetter: "TypeInput"},
		},
		func() interface{} {
			j := jsiiProxy_SshSecretBackendRoleAllowedUserKeyConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-vault.sshSecretBackendRole.SshSecretBackendRoleConfig",
		reflect.TypeOf((*SshSecretBackendRoleConfig)(nil)).Elem(),
	)
}
