package ldapsecretbackend

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-vault.ldapSecretBackend.LdapSecretBackend",
		reflect.TypeOf((*LdapSecretBackend)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accessor", GoGetter: "Accessor"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "allowedManagedKeys", GoGetter: "AllowedManagedKeys"},
			_jsii_.MemberProperty{JsiiProperty: "allowedManagedKeysInput", GoGetter: "AllowedManagedKeysInput"},
			_jsii_.MemberProperty{JsiiProperty: "auditNonHmacRequestKeys", GoGetter: "AuditNonHmacRequestKeys"},
			_jsii_.MemberProperty{JsiiProperty: "auditNonHmacRequestKeysInput", GoGetter: "AuditNonHmacRequestKeysInput"},
			_jsii_.MemberProperty{JsiiProperty: "auditNonHmacResponseKeys", GoGetter: "AuditNonHmacResponseKeys"},
			_jsii_.MemberProperty{JsiiProperty: "auditNonHmacResponseKeysInput", GoGetter: "AuditNonHmacResponseKeysInput"},
			_jsii_.MemberProperty{JsiiProperty: "binddn", GoGetter: "Binddn"},
			_jsii_.MemberProperty{JsiiProperty: "binddnInput", GoGetter: "BinddnInput"},
			_jsii_.MemberProperty{JsiiProperty: "bindpass", GoGetter: "Bindpass"},
			_jsii_.MemberProperty{JsiiProperty: "bindpassInput", GoGetter: "BindpassInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "certificate", GoGetter: "Certificate"},
			_jsii_.MemberProperty{JsiiProperty: "certificateInput", GoGetter: "CertificateInput"},
			_jsii_.MemberProperty{JsiiProperty: "clientTlsCert", GoGetter: "ClientTlsCert"},
			_jsii_.MemberProperty{JsiiProperty: "clientTlsCertInput", GoGetter: "ClientTlsCertInput"},
			_jsii_.MemberProperty{JsiiProperty: "clientTlsKey", GoGetter: "ClientTlsKey"},
			_jsii_.MemberProperty{JsiiProperty: "clientTlsKeyInput", GoGetter: "ClientTlsKeyInput"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "connectionTimeout", GoGetter: "ConnectionTimeout"},
			_jsii_.MemberProperty{JsiiProperty: "connectionTimeoutInput", GoGetter: "ConnectionTimeoutInput"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "defaultLeaseTtlSeconds", GoGetter: "DefaultLeaseTtlSeconds"},
			_jsii_.MemberProperty{JsiiProperty: "defaultLeaseTtlSecondsInput", GoGetter: "DefaultLeaseTtlSecondsInput"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "descriptionInput", GoGetter: "DescriptionInput"},
			_jsii_.MemberProperty{JsiiProperty: "disableRemount", GoGetter: "DisableRemount"},
			_jsii_.MemberProperty{JsiiProperty: "disableRemountInput", GoGetter: "DisableRemountInput"},
			_jsii_.MemberProperty{JsiiProperty: "externalEntropyAccess", GoGetter: "ExternalEntropyAccess"},
			_jsii_.MemberProperty{JsiiProperty: "externalEntropyAccessInput", GoGetter: "ExternalEntropyAccessInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "insecureTls", GoGetter: "InsecureTls"},
			_jsii_.MemberProperty{JsiiProperty: "insecureTlsInput", GoGetter: "InsecureTlsInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "length", GoGetter: "Length"},
			_jsii_.MemberProperty{JsiiProperty: "lengthInput", GoGetter: "LengthInput"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "local", GoGetter: "Local"},
			_jsii_.MemberProperty{JsiiProperty: "localInput", GoGetter: "LocalInput"},
			_jsii_.MemberProperty{JsiiProperty: "maxLeaseTtlSeconds", GoGetter: "MaxLeaseTtlSeconds"},
			_jsii_.MemberProperty{JsiiProperty: "maxLeaseTtlSecondsInput", GoGetter: "MaxLeaseTtlSecondsInput"},
			_jsii_.MemberProperty{JsiiProperty: "namespace", GoGetter: "Namespace"},
			_jsii_.MemberProperty{JsiiProperty: "namespaceInput", GoGetter: "NamespaceInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "options", GoGetter: "Options"},
			_jsii_.MemberProperty{JsiiProperty: "optionsInput", GoGetter: "OptionsInput"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "passwordPolicy", GoGetter: "PasswordPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "passwordPolicyInput", GoGetter: "PasswordPolicyInput"},
			_jsii_.MemberProperty{JsiiProperty: "path", GoGetter: "Path"},
			_jsii_.MemberProperty{JsiiProperty: "pathInput", GoGetter: "PathInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "requestTimeout", GoGetter: "RequestTimeout"},
			_jsii_.MemberProperty{JsiiProperty: "requestTimeoutInput", GoGetter: "RequestTimeoutInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowedManagedKeys", GoMethod: "ResetAllowedManagedKeys"},
			_jsii_.MemberMethod{JsiiMethod: "resetAuditNonHmacRequestKeys", GoMethod: "ResetAuditNonHmacRequestKeys"},
			_jsii_.MemberMethod{JsiiMethod: "resetAuditNonHmacResponseKeys", GoMethod: "ResetAuditNonHmacResponseKeys"},
			_jsii_.MemberMethod{JsiiMethod: "resetCertificate", GoMethod: "ResetCertificate"},
			_jsii_.MemberMethod{JsiiMethod: "resetClientTlsCert", GoMethod: "ResetClientTlsCert"},
			_jsii_.MemberMethod{JsiiMethod: "resetClientTlsKey", GoMethod: "ResetClientTlsKey"},
			_jsii_.MemberMethod{JsiiMethod: "resetConnectionTimeout", GoMethod: "ResetConnectionTimeout"},
			_jsii_.MemberMethod{JsiiMethod: "resetDefaultLeaseTtlSeconds", GoMethod: "ResetDefaultLeaseTtlSeconds"},
			_jsii_.MemberMethod{JsiiMethod: "resetDescription", GoMethod: "ResetDescription"},
			_jsii_.MemberMethod{JsiiMethod: "resetDisableRemount", GoMethod: "ResetDisableRemount"},
			_jsii_.MemberMethod{JsiiMethod: "resetExternalEntropyAccess", GoMethod: "ResetExternalEntropyAccess"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetInsecureTls", GoMethod: "ResetInsecureTls"},
			_jsii_.MemberMethod{JsiiMethod: "resetLength", GoMethod: "ResetLength"},
			_jsii_.MemberMethod{JsiiMethod: "resetLocal", GoMethod: "ResetLocal"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxLeaseTtlSeconds", GoMethod: "ResetMaxLeaseTtlSeconds"},
			_jsii_.MemberMethod{JsiiMethod: "resetNamespace", GoMethod: "ResetNamespace"},
			_jsii_.MemberMethod{JsiiMethod: "resetOptions", GoMethod: "ResetOptions"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPasswordPolicy", GoMethod: "ResetPasswordPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "resetPath", GoMethod: "ResetPath"},
			_jsii_.MemberMethod{JsiiMethod: "resetRequestTimeout", GoMethod: "ResetRequestTimeout"},
			_jsii_.MemberMethod{JsiiMethod: "resetSchema", GoMethod: "ResetSchema"},
			_jsii_.MemberMethod{JsiiMethod: "resetSealWrap", GoMethod: "ResetSealWrap"},
			_jsii_.MemberMethod{JsiiMethod: "resetStarttls", GoMethod: "ResetStarttls"},
			_jsii_.MemberMethod{JsiiMethod: "resetUpndomain", GoMethod: "ResetUpndomain"},
			_jsii_.MemberMethod{JsiiMethod: "resetUrl", GoMethod: "ResetUrl"},
			_jsii_.MemberMethod{JsiiMethod: "resetUserattr", GoMethod: "ResetUserattr"},
			_jsii_.MemberMethod{JsiiMethod: "resetUserdn", GoMethod: "ResetUserdn"},
			_jsii_.MemberProperty{JsiiProperty: "schema", GoGetter: "Schema"},
			_jsii_.MemberProperty{JsiiProperty: "schemaInput", GoGetter: "SchemaInput"},
			_jsii_.MemberProperty{JsiiProperty: "sealWrap", GoGetter: "SealWrap"},
			_jsii_.MemberProperty{JsiiProperty: "sealWrapInput", GoGetter: "SealWrapInput"},
			_jsii_.MemberProperty{JsiiProperty: "starttls", GoGetter: "Starttls"},
			_jsii_.MemberProperty{JsiiProperty: "starttlsInput", GoGetter: "StarttlsInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "upndomain", GoGetter: "Upndomain"},
			_jsii_.MemberProperty{JsiiProperty: "upndomainInput", GoGetter: "UpndomainInput"},
			_jsii_.MemberProperty{JsiiProperty: "url", GoGetter: "Url"},
			_jsii_.MemberProperty{JsiiProperty: "urlInput", GoGetter: "UrlInput"},
			_jsii_.MemberProperty{JsiiProperty: "userattr", GoGetter: "Userattr"},
			_jsii_.MemberProperty{JsiiProperty: "userattrInput", GoGetter: "UserattrInput"},
			_jsii_.MemberProperty{JsiiProperty: "userdn", GoGetter: "Userdn"},
			_jsii_.MemberProperty{JsiiProperty: "userdnInput", GoGetter: "UserdnInput"},
		},
		func() interface{} {
			j := jsiiProxy_LdapSecretBackend{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-vault.ldapSecretBackend.LdapSecretBackendConfig",
		reflect.TypeOf((*LdapSecretBackendConfig)(nil)).Elem(),
	)
}
