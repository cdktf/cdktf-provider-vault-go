package pkisecretbackendrole

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-vault.pkiSecretBackendRole.PkiSecretBackendRole",
		reflect.TypeOf((*PkiSecretBackendRole)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "allowAnyName", GoGetter: "AllowAnyName"},
			_jsii_.MemberProperty{JsiiProperty: "allowAnyNameInput", GoGetter: "AllowAnyNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "allowBareDomains", GoGetter: "AllowBareDomains"},
			_jsii_.MemberProperty{JsiiProperty: "allowBareDomainsInput", GoGetter: "AllowBareDomainsInput"},
			_jsii_.MemberProperty{JsiiProperty: "allowedDomains", GoGetter: "AllowedDomains"},
			_jsii_.MemberProperty{JsiiProperty: "allowedDomainsInput", GoGetter: "AllowedDomainsInput"},
			_jsii_.MemberProperty{JsiiProperty: "allowedDomainsTemplate", GoGetter: "AllowedDomainsTemplate"},
			_jsii_.MemberProperty{JsiiProperty: "allowedDomainsTemplateInput", GoGetter: "AllowedDomainsTemplateInput"},
			_jsii_.MemberProperty{JsiiProperty: "allowedOtherSans", GoGetter: "AllowedOtherSans"},
			_jsii_.MemberProperty{JsiiProperty: "allowedOtherSansInput", GoGetter: "AllowedOtherSansInput"},
			_jsii_.MemberProperty{JsiiProperty: "allowedSerialNumbers", GoGetter: "AllowedSerialNumbers"},
			_jsii_.MemberProperty{JsiiProperty: "allowedSerialNumbersInput", GoGetter: "AllowedSerialNumbersInput"},
			_jsii_.MemberProperty{JsiiProperty: "allowedUriSans", GoGetter: "AllowedUriSans"},
			_jsii_.MemberProperty{JsiiProperty: "allowedUriSansInput", GoGetter: "AllowedUriSansInput"},
			_jsii_.MemberProperty{JsiiProperty: "allowGlobDomains", GoGetter: "AllowGlobDomains"},
			_jsii_.MemberProperty{JsiiProperty: "allowGlobDomainsInput", GoGetter: "AllowGlobDomainsInput"},
			_jsii_.MemberProperty{JsiiProperty: "allowIpSans", GoGetter: "AllowIpSans"},
			_jsii_.MemberProperty{JsiiProperty: "allowIpSansInput", GoGetter: "AllowIpSansInput"},
			_jsii_.MemberProperty{JsiiProperty: "allowLocalhost", GoGetter: "AllowLocalhost"},
			_jsii_.MemberProperty{JsiiProperty: "allowLocalhostInput", GoGetter: "AllowLocalhostInput"},
			_jsii_.MemberProperty{JsiiProperty: "allowSubdomains", GoGetter: "AllowSubdomains"},
			_jsii_.MemberProperty{JsiiProperty: "allowSubdomainsInput", GoGetter: "AllowSubdomainsInput"},
			_jsii_.MemberProperty{JsiiProperty: "backend", GoGetter: "Backend"},
			_jsii_.MemberProperty{JsiiProperty: "backendInput", GoGetter: "BackendInput"},
			_jsii_.MemberProperty{JsiiProperty: "basicConstraintsValidForNonCa", GoGetter: "BasicConstraintsValidForNonCa"},
			_jsii_.MemberProperty{JsiiProperty: "basicConstraintsValidForNonCaInput", GoGetter: "BasicConstraintsValidForNonCaInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "clientFlag", GoGetter: "ClientFlag"},
			_jsii_.MemberProperty{JsiiProperty: "clientFlagInput", GoGetter: "ClientFlagInput"},
			_jsii_.MemberProperty{JsiiProperty: "codeSigningFlag", GoGetter: "CodeSigningFlag"},
			_jsii_.MemberProperty{JsiiProperty: "codeSigningFlagInput", GoGetter: "CodeSigningFlagInput"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "country", GoGetter: "Country"},
			_jsii_.MemberProperty{JsiiProperty: "countryInput", GoGetter: "CountryInput"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "emailProtectionFlag", GoGetter: "EmailProtectionFlag"},
			_jsii_.MemberProperty{JsiiProperty: "emailProtectionFlagInput", GoGetter: "EmailProtectionFlagInput"},
			_jsii_.MemberProperty{JsiiProperty: "enforceHostnames", GoGetter: "EnforceHostnames"},
			_jsii_.MemberProperty{JsiiProperty: "enforceHostnamesInput", GoGetter: "EnforceHostnamesInput"},
			_jsii_.MemberProperty{JsiiProperty: "extKeyUsage", GoGetter: "ExtKeyUsage"},
			_jsii_.MemberProperty{JsiiProperty: "extKeyUsageInput", GoGetter: "ExtKeyUsageInput"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberProperty{JsiiProperty: "generateLease", GoGetter: "GenerateLease"},
			_jsii_.MemberProperty{JsiiProperty: "generateLeaseInput", GoGetter: "GenerateLeaseInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "keyBits", GoGetter: "KeyBits"},
			_jsii_.MemberProperty{JsiiProperty: "keyBitsInput", GoGetter: "KeyBitsInput"},
			_jsii_.MemberProperty{JsiiProperty: "keyType", GoGetter: "KeyType"},
			_jsii_.MemberProperty{JsiiProperty: "keyTypeInput", GoGetter: "KeyTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "keyUsage", GoGetter: "KeyUsage"},
			_jsii_.MemberProperty{JsiiProperty: "keyUsageInput", GoGetter: "KeyUsageInput"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "locality", GoGetter: "Locality"},
			_jsii_.MemberProperty{JsiiProperty: "localityInput", GoGetter: "LocalityInput"},
			_jsii_.MemberProperty{JsiiProperty: "maxTtl", GoGetter: "MaxTtl"},
			_jsii_.MemberProperty{JsiiProperty: "maxTtlInput", GoGetter: "MaxTtlInput"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "namespace", GoGetter: "Namespace"},
			_jsii_.MemberProperty{JsiiProperty: "namespaceInput", GoGetter: "NamespaceInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "noStore", GoGetter: "NoStore"},
			_jsii_.MemberProperty{JsiiProperty: "noStoreInput", GoGetter: "NoStoreInput"},
			_jsii_.MemberProperty{JsiiProperty: "notBeforeDuration", GoGetter: "NotBeforeDuration"},
			_jsii_.MemberProperty{JsiiProperty: "notBeforeDurationInput", GoGetter: "NotBeforeDurationInput"},
			_jsii_.MemberProperty{JsiiProperty: "organization", GoGetter: "Organization"},
			_jsii_.MemberProperty{JsiiProperty: "organizationInput", GoGetter: "OrganizationInput"},
			_jsii_.MemberProperty{JsiiProperty: "ou", GoGetter: "Ou"},
			_jsii_.MemberProperty{JsiiProperty: "ouInput", GoGetter: "OuInput"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "policyIdentifier", GoGetter: "PolicyIdentifier"},
			_jsii_.MemberProperty{JsiiProperty: "policyIdentifierInput", GoGetter: "PolicyIdentifierInput"},
			_jsii_.MemberProperty{JsiiProperty: "policyIdentifiers", GoGetter: "PolicyIdentifiers"},
			_jsii_.MemberProperty{JsiiProperty: "policyIdentifiersInput", GoGetter: "PolicyIdentifiersInput"},
			_jsii_.MemberProperty{JsiiProperty: "postalCode", GoGetter: "PostalCode"},
			_jsii_.MemberProperty{JsiiProperty: "postalCodeInput", GoGetter: "PostalCodeInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "province", GoGetter: "Province"},
			_jsii_.MemberProperty{JsiiProperty: "provinceInput", GoGetter: "ProvinceInput"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putPolicyIdentifier", GoMethod: "PutPolicyIdentifier"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "requireCn", GoGetter: "RequireCn"},
			_jsii_.MemberProperty{JsiiProperty: "requireCnInput", GoGetter: "RequireCnInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowAnyName", GoMethod: "ResetAllowAnyName"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowBareDomains", GoMethod: "ResetAllowBareDomains"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowedDomains", GoMethod: "ResetAllowedDomains"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowedDomainsTemplate", GoMethod: "ResetAllowedDomainsTemplate"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowedOtherSans", GoMethod: "ResetAllowedOtherSans"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowedSerialNumbers", GoMethod: "ResetAllowedSerialNumbers"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowedUriSans", GoMethod: "ResetAllowedUriSans"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowGlobDomains", GoMethod: "ResetAllowGlobDomains"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowIpSans", GoMethod: "ResetAllowIpSans"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowLocalhost", GoMethod: "ResetAllowLocalhost"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowSubdomains", GoMethod: "ResetAllowSubdomains"},
			_jsii_.MemberMethod{JsiiMethod: "resetBasicConstraintsValidForNonCa", GoMethod: "ResetBasicConstraintsValidForNonCa"},
			_jsii_.MemberMethod{JsiiMethod: "resetClientFlag", GoMethod: "ResetClientFlag"},
			_jsii_.MemberMethod{JsiiMethod: "resetCodeSigningFlag", GoMethod: "ResetCodeSigningFlag"},
			_jsii_.MemberMethod{JsiiMethod: "resetCountry", GoMethod: "ResetCountry"},
			_jsii_.MemberMethod{JsiiMethod: "resetEmailProtectionFlag", GoMethod: "ResetEmailProtectionFlag"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnforceHostnames", GoMethod: "ResetEnforceHostnames"},
			_jsii_.MemberMethod{JsiiMethod: "resetExtKeyUsage", GoMethod: "ResetExtKeyUsage"},
			_jsii_.MemberMethod{JsiiMethod: "resetGenerateLease", GoMethod: "ResetGenerateLease"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetKeyBits", GoMethod: "ResetKeyBits"},
			_jsii_.MemberMethod{JsiiMethod: "resetKeyType", GoMethod: "ResetKeyType"},
			_jsii_.MemberMethod{JsiiMethod: "resetKeyUsage", GoMethod: "ResetKeyUsage"},
			_jsii_.MemberMethod{JsiiMethod: "resetLocality", GoMethod: "ResetLocality"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxTtl", GoMethod: "ResetMaxTtl"},
			_jsii_.MemberMethod{JsiiMethod: "resetNamespace", GoMethod: "ResetNamespace"},
			_jsii_.MemberMethod{JsiiMethod: "resetNoStore", GoMethod: "ResetNoStore"},
			_jsii_.MemberMethod{JsiiMethod: "resetNotBeforeDuration", GoMethod: "ResetNotBeforeDuration"},
			_jsii_.MemberMethod{JsiiMethod: "resetOrganization", GoMethod: "ResetOrganization"},
			_jsii_.MemberMethod{JsiiMethod: "resetOu", GoMethod: "ResetOu"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPolicyIdentifier", GoMethod: "ResetPolicyIdentifier"},
			_jsii_.MemberMethod{JsiiMethod: "resetPolicyIdentifiers", GoMethod: "ResetPolicyIdentifiers"},
			_jsii_.MemberMethod{JsiiMethod: "resetPostalCode", GoMethod: "ResetPostalCode"},
			_jsii_.MemberMethod{JsiiMethod: "resetProvince", GoMethod: "ResetProvince"},
			_jsii_.MemberMethod{JsiiMethod: "resetRequireCn", GoMethod: "ResetRequireCn"},
			_jsii_.MemberMethod{JsiiMethod: "resetServerFlag", GoMethod: "ResetServerFlag"},
			_jsii_.MemberMethod{JsiiMethod: "resetStreetAddress", GoMethod: "ResetStreetAddress"},
			_jsii_.MemberMethod{JsiiMethod: "resetTtl", GoMethod: "ResetTtl"},
			_jsii_.MemberMethod{JsiiMethod: "resetUseCsrCommonName", GoMethod: "ResetUseCsrCommonName"},
			_jsii_.MemberMethod{JsiiMethod: "resetUseCsrSans", GoMethod: "ResetUseCsrSans"},
			_jsii_.MemberProperty{JsiiProperty: "serverFlag", GoGetter: "ServerFlag"},
			_jsii_.MemberProperty{JsiiProperty: "serverFlagInput", GoGetter: "ServerFlagInput"},
			_jsii_.MemberProperty{JsiiProperty: "streetAddress", GoGetter: "StreetAddress"},
			_jsii_.MemberProperty{JsiiProperty: "streetAddressInput", GoGetter: "StreetAddressInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "ttl", GoGetter: "Ttl"},
			_jsii_.MemberProperty{JsiiProperty: "ttlInput", GoGetter: "TtlInput"},
			_jsii_.MemberProperty{JsiiProperty: "useCsrCommonName", GoGetter: "UseCsrCommonName"},
			_jsii_.MemberProperty{JsiiProperty: "useCsrCommonNameInput", GoGetter: "UseCsrCommonNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "useCsrSans", GoGetter: "UseCsrSans"},
			_jsii_.MemberProperty{JsiiProperty: "useCsrSansInput", GoGetter: "UseCsrSansInput"},
		},
		func() interface{} {
			j := jsiiProxy_PkiSecretBackendRole{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-vault.pkiSecretBackendRole.PkiSecretBackendRoleConfig",
		reflect.TypeOf((*PkiSecretBackendRoleConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-vault.pkiSecretBackendRole.PkiSecretBackendRolePolicyIdentifier",
		reflect.TypeOf((*PkiSecretBackendRolePolicyIdentifier)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-vault.pkiSecretBackendRole.PkiSecretBackendRolePolicyIdentifierList",
		reflect.TypeOf((*PkiSecretBackendRolePolicyIdentifierList)(nil)).Elem(),
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
			j := jsiiProxy_PkiSecretBackendRolePolicyIdentifierList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-vault.pkiSecretBackendRole.PkiSecretBackendRolePolicyIdentifierOutputReference",
		reflect.TypeOf((*PkiSecretBackendRolePolicyIdentifierOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "cps", GoGetter: "Cps"},
			_jsii_.MemberProperty{JsiiProperty: "cpsInput", GoGetter: "CpsInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "notice", GoGetter: "Notice"},
			_jsii_.MemberProperty{JsiiProperty: "noticeInput", GoGetter: "NoticeInput"},
			_jsii_.MemberProperty{JsiiProperty: "oid", GoGetter: "Oid"},
			_jsii_.MemberProperty{JsiiProperty: "oidInput", GoGetter: "OidInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetCps", GoMethod: "ResetCps"},
			_jsii_.MemberMethod{JsiiMethod: "resetNotice", GoMethod: "ResetNotice"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_PkiSecretBackendRolePolicyIdentifierOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
