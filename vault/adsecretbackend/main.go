// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package adsecretbackend

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-vault.adSecretBackend.AdSecretBackend",
		reflect.TypeOf((*AdSecretBackend)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "anonymousGroupSearch", GoGetter: "AnonymousGroupSearch"},
			_jsii_.MemberProperty{JsiiProperty: "anonymousGroupSearchInput", GoGetter: "AnonymousGroupSearchInput"},
			_jsii_.MemberProperty{JsiiProperty: "backend", GoGetter: "Backend"},
			_jsii_.MemberProperty{JsiiProperty: "backendInput", GoGetter: "BackendInput"},
			_jsii_.MemberProperty{JsiiProperty: "binddn", GoGetter: "Binddn"},
			_jsii_.MemberProperty{JsiiProperty: "binddnInput", GoGetter: "BinddnInput"},
			_jsii_.MemberProperty{JsiiProperty: "bindpass", GoGetter: "Bindpass"},
			_jsii_.MemberProperty{JsiiProperty: "bindpassInput", GoGetter: "BindpassInput"},
			_jsii_.MemberProperty{JsiiProperty: "caseSensitiveNames", GoGetter: "CaseSensitiveNames"},
			_jsii_.MemberProperty{JsiiProperty: "caseSensitiveNamesInput", GoGetter: "CaseSensitiveNamesInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "certificate", GoGetter: "Certificate"},
			_jsii_.MemberProperty{JsiiProperty: "certificateInput", GoGetter: "CertificateInput"},
			_jsii_.MemberProperty{JsiiProperty: "clientTlsCert", GoGetter: "ClientTlsCert"},
			_jsii_.MemberProperty{JsiiProperty: "clientTlsCertInput", GoGetter: "ClientTlsCertInput"},
			_jsii_.MemberProperty{JsiiProperty: "clientTlsKey", GoGetter: "ClientTlsKey"},
			_jsii_.MemberProperty{JsiiProperty: "clientTlsKeyInput", GoGetter: "ClientTlsKeyInput"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "defaultLeaseTtlSeconds", GoGetter: "DefaultLeaseTtlSeconds"},
			_jsii_.MemberProperty{JsiiProperty: "defaultLeaseTtlSecondsInput", GoGetter: "DefaultLeaseTtlSecondsInput"},
			_jsii_.MemberProperty{JsiiProperty: "denyNullBind", GoGetter: "DenyNullBind"},
			_jsii_.MemberProperty{JsiiProperty: "denyNullBindInput", GoGetter: "DenyNullBindInput"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "descriptionInput", GoGetter: "DescriptionInput"},
			_jsii_.MemberProperty{JsiiProperty: "disableRemount", GoGetter: "DisableRemount"},
			_jsii_.MemberProperty{JsiiProperty: "disableRemountInput", GoGetter: "DisableRemountInput"},
			_jsii_.MemberProperty{JsiiProperty: "discoverdn", GoGetter: "Discoverdn"},
			_jsii_.MemberProperty{JsiiProperty: "discoverdnInput", GoGetter: "DiscoverdnInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "groupattr", GoGetter: "Groupattr"},
			_jsii_.MemberProperty{JsiiProperty: "groupattrInput", GoGetter: "GroupattrInput"},
			_jsii_.MemberProperty{JsiiProperty: "groupdn", GoGetter: "Groupdn"},
			_jsii_.MemberProperty{JsiiProperty: "groupdnInput", GoGetter: "GroupdnInput"},
			_jsii_.MemberProperty{JsiiProperty: "groupfilter", GoGetter: "Groupfilter"},
			_jsii_.MemberProperty{JsiiProperty: "groupfilterInput", GoGetter: "GroupfilterInput"},
			_jsii_.MemberMethod{JsiiMethod: "hasResourceMove", GoMethod: "HasResourceMove"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberProperty{JsiiProperty: "insecureTls", GoGetter: "InsecureTls"},
			_jsii_.MemberProperty{JsiiProperty: "insecureTlsInput", GoGetter: "InsecureTlsInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lastRotationTolerance", GoGetter: "LastRotationTolerance"},
			_jsii_.MemberProperty{JsiiProperty: "lastRotationToleranceInput", GoGetter: "LastRotationToleranceInput"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "local", GoGetter: "Local"},
			_jsii_.MemberProperty{JsiiProperty: "localInput", GoGetter: "LocalInput"},
			_jsii_.MemberProperty{JsiiProperty: "maxLeaseTtlSeconds", GoGetter: "MaxLeaseTtlSeconds"},
			_jsii_.MemberProperty{JsiiProperty: "maxLeaseTtlSecondsInput", GoGetter: "MaxLeaseTtlSecondsInput"},
			_jsii_.MemberProperty{JsiiProperty: "maxTtl", GoGetter: "MaxTtl"},
			_jsii_.MemberProperty{JsiiProperty: "maxTtlInput", GoGetter: "MaxTtlInput"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "namespace", GoGetter: "Namespace"},
			_jsii_.MemberProperty{JsiiProperty: "namespaceInput", GoGetter: "NamespaceInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "passwordPolicy", GoGetter: "PasswordPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "passwordPolicyInput", GoGetter: "PasswordPolicyInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "requestTimeout", GoGetter: "RequestTimeout"},
			_jsii_.MemberProperty{JsiiProperty: "requestTimeoutInput", GoGetter: "RequestTimeoutInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAnonymousGroupSearch", GoMethod: "ResetAnonymousGroupSearch"},
			_jsii_.MemberMethod{JsiiMethod: "resetBackend", GoMethod: "ResetBackend"},
			_jsii_.MemberMethod{JsiiMethod: "resetCaseSensitiveNames", GoMethod: "ResetCaseSensitiveNames"},
			_jsii_.MemberMethod{JsiiMethod: "resetCertificate", GoMethod: "ResetCertificate"},
			_jsii_.MemberMethod{JsiiMethod: "resetClientTlsCert", GoMethod: "ResetClientTlsCert"},
			_jsii_.MemberMethod{JsiiMethod: "resetClientTlsKey", GoMethod: "ResetClientTlsKey"},
			_jsii_.MemberMethod{JsiiMethod: "resetDefaultLeaseTtlSeconds", GoMethod: "ResetDefaultLeaseTtlSeconds"},
			_jsii_.MemberMethod{JsiiMethod: "resetDenyNullBind", GoMethod: "ResetDenyNullBind"},
			_jsii_.MemberMethod{JsiiMethod: "resetDescription", GoMethod: "ResetDescription"},
			_jsii_.MemberMethod{JsiiMethod: "resetDisableRemount", GoMethod: "ResetDisableRemount"},
			_jsii_.MemberMethod{JsiiMethod: "resetDiscoverdn", GoMethod: "ResetDiscoverdn"},
			_jsii_.MemberMethod{JsiiMethod: "resetGroupattr", GoMethod: "ResetGroupattr"},
			_jsii_.MemberMethod{JsiiMethod: "resetGroupdn", GoMethod: "ResetGroupdn"},
			_jsii_.MemberMethod{JsiiMethod: "resetGroupfilter", GoMethod: "ResetGroupfilter"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetInsecureTls", GoMethod: "ResetInsecureTls"},
			_jsii_.MemberMethod{JsiiMethod: "resetLastRotationTolerance", GoMethod: "ResetLastRotationTolerance"},
			_jsii_.MemberMethod{JsiiMethod: "resetLocal", GoMethod: "ResetLocal"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxLeaseTtlSeconds", GoMethod: "ResetMaxLeaseTtlSeconds"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxTtl", GoMethod: "ResetMaxTtl"},
			_jsii_.MemberMethod{JsiiMethod: "resetNamespace", GoMethod: "ResetNamespace"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPasswordPolicy", GoMethod: "ResetPasswordPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "resetRequestTimeout", GoMethod: "ResetRequestTimeout"},
			_jsii_.MemberMethod{JsiiMethod: "resetStarttls", GoMethod: "ResetStarttls"},
			_jsii_.MemberMethod{JsiiMethod: "resetTlsMaxVersion", GoMethod: "ResetTlsMaxVersion"},
			_jsii_.MemberMethod{JsiiMethod: "resetTlsMinVersion", GoMethod: "ResetTlsMinVersion"},
			_jsii_.MemberMethod{JsiiMethod: "resetTtl", GoMethod: "ResetTtl"},
			_jsii_.MemberMethod{JsiiMethod: "resetUpndomain", GoMethod: "ResetUpndomain"},
			_jsii_.MemberMethod{JsiiMethod: "resetUrl", GoMethod: "ResetUrl"},
			_jsii_.MemberMethod{JsiiMethod: "resetUsePre111GroupCnBehavior", GoMethod: "ResetUsePre111GroupCnBehavior"},
			_jsii_.MemberMethod{JsiiMethod: "resetUserattr", GoMethod: "ResetUserattr"},
			_jsii_.MemberMethod{JsiiMethod: "resetUserdn", GoMethod: "ResetUserdn"},
			_jsii_.MemberMethod{JsiiMethod: "resetUseTokenGroups", GoMethod: "ResetUseTokenGroups"},
			_jsii_.MemberProperty{JsiiProperty: "starttls", GoGetter: "Starttls"},
			_jsii_.MemberProperty{JsiiProperty: "starttlsInput", GoGetter: "StarttlsInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "tlsMaxVersion", GoGetter: "TlsMaxVersion"},
			_jsii_.MemberProperty{JsiiProperty: "tlsMaxVersionInput", GoGetter: "TlsMaxVersionInput"},
			_jsii_.MemberProperty{JsiiProperty: "tlsMinVersion", GoGetter: "TlsMinVersion"},
			_jsii_.MemberProperty{JsiiProperty: "tlsMinVersionInput", GoGetter: "TlsMinVersionInput"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "ttl", GoGetter: "Ttl"},
			_jsii_.MemberProperty{JsiiProperty: "ttlInput", GoGetter: "TtlInput"},
			_jsii_.MemberProperty{JsiiProperty: "upndomain", GoGetter: "Upndomain"},
			_jsii_.MemberProperty{JsiiProperty: "upndomainInput", GoGetter: "UpndomainInput"},
			_jsii_.MemberProperty{JsiiProperty: "url", GoGetter: "Url"},
			_jsii_.MemberProperty{JsiiProperty: "urlInput", GoGetter: "UrlInput"},
			_jsii_.MemberProperty{JsiiProperty: "usePre111GroupCnBehavior", GoGetter: "UsePre111GroupCnBehavior"},
			_jsii_.MemberProperty{JsiiProperty: "usePre111GroupCnBehaviorInput", GoGetter: "UsePre111GroupCnBehaviorInput"},
			_jsii_.MemberProperty{JsiiProperty: "userattr", GoGetter: "Userattr"},
			_jsii_.MemberProperty{JsiiProperty: "userattrInput", GoGetter: "UserattrInput"},
			_jsii_.MemberProperty{JsiiProperty: "userdn", GoGetter: "Userdn"},
			_jsii_.MemberProperty{JsiiProperty: "userdnInput", GoGetter: "UserdnInput"},
			_jsii_.MemberProperty{JsiiProperty: "useTokenGroups", GoGetter: "UseTokenGroups"},
			_jsii_.MemberProperty{JsiiProperty: "useTokenGroupsInput", GoGetter: "UseTokenGroupsInput"},
		},
		func() interface{} {
			j := jsiiProxy_AdSecretBackend{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-vault.adSecretBackend.AdSecretBackendConfig",
		reflect.TypeOf((*AdSecretBackendConfig)(nil)).Elem(),
	)
}
