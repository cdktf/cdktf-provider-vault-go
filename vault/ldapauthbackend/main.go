// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ldapauthbackend

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-vault.ldapAuthBackend.LdapAuthBackend",
		reflect.TypeOf((*LdapAuthBackend)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accessor", GoGetter: "Accessor"},
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
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
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberProperty{JsiiProperty: "insecureTls", GoGetter: "InsecureTls"},
			_jsii_.MemberProperty{JsiiProperty: "insecureTlsInput", GoGetter: "InsecureTlsInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "local", GoGetter: "Local"},
			_jsii_.MemberProperty{JsiiProperty: "localInput", GoGetter: "LocalInput"},
			_jsii_.MemberProperty{JsiiProperty: "maxPageSize", GoGetter: "MaxPageSize"},
			_jsii_.MemberProperty{JsiiProperty: "maxPageSizeInput", GoGetter: "MaxPageSizeInput"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberProperty{JsiiProperty: "namespace", GoGetter: "Namespace"},
			_jsii_.MemberProperty{JsiiProperty: "namespaceInput", GoGetter: "NamespaceInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "path", GoGetter: "Path"},
			_jsii_.MemberProperty{JsiiProperty: "pathInput", GoGetter: "PathInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetBinddn", GoMethod: "ResetBinddn"},
			_jsii_.MemberMethod{JsiiMethod: "resetBindpass", GoMethod: "ResetBindpass"},
			_jsii_.MemberMethod{JsiiMethod: "resetCaseSensitiveNames", GoMethod: "ResetCaseSensitiveNames"},
			_jsii_.MemberMethod{JsiiMethod: "resetCertificate", GoMethod: "ResetCertificate"},
			_jsii_.MemberMethod{JsiiMethod: "resetClientTlsCert", GoMethod: "ResetClientTlsCert"},
			_jsii_.MemberMethod{JsiiMethod: "resetClientTlsKey", GoMethod: "ResetClientTlsKey"},
			_jsii_.MemberMethod{JsiiMethod: "resetDenyNullBind", GoMethod: "ResetDenyNullBind"},
			_jsii_.MemberMethod{JsiiMethod: "resetDescription", GoMethod: "ResetDescription"},
			_jsii_.MemberMethod{JsiiMethod: "resetDisableRemount", GoMethod: "ResetDisableRemount"},
			_jsii_.MemberMethod{JsiiMethod: "resetDiscoverdn", GoMethod: "ResetDiscoverdn"},
			_jsii_.MemberMethod{JsiiMethod: "resetGroupattr", GoMethod: "ResetGroupattr"},
			_jsii_.MemberMethod{JsiiMethod: "resetGroupdn", GoMethod: "ResetGroupdn"},
			_jsii_.MemberMethod{JsiiMethod: "resetGroupfilter", GoMethod: "ResetGroupfilter"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetInsecureTls", GoMethod: "ResetInsecureTls"},
			_jsii_.MemberMethod{JsiiMethod: "resetLocal", GoMethod: "ResetLocal"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxPageSize", GoMethod: "ResetMaxPageSize"},
			_jsii_.MemberMethod{JsiiMethod: "resetNamespace", GoMethod: "ResetNamespace"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPath", GoMethod: "ResetPath"},
			_jsii_.MemberMethod{JsiiMethod: "resetStarttls", GoMethod: "ResetStarttls"},
			_jsii_.MemberMethod{JsiiMethod: "resetTlsMaxVersion", GoMethod: "ResetTlsMaxVersion"},
			_jsii_.MemberMethod{JsiiMethod: "resetTlsMinVersion", GoMethod: "ResetTlsMinVersion"},
			_jsii_.MemberMethod{JsiiMethod: "resetTokenBoundCidrs", GoMethod: "ResetTokenBoundCidrs"},
			_jsii_.MemberMethod{JsiiMethod: "resetTokenExplicitMaxTtl", GoMethod: "ResetTokenExplicitMaxTtl"},
			_jsii_.MemberMethod{JsiiMethod: "resetTokenMaxTtl", GoMethod: "ResetTokenMaxTtl"},
			_jsii_.MemberMethod{JsiiMethod: "resetTokenNoDefaultPolicy", GoMethod: "ResetTokenNoDefaultPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "resetTokenNumUses", GoMethod: "ResetTokenNumUses"},
			_jsii_.MemberMethod{JsiiMethod: "resetTokenPeriod", GoMethod: "ResetTokenPeriod"},
			_jsii_.MemberMethod{JsiiMethod: "resetTokenPolicies", GoMethod: "ResetTokenPolicies"},
			_jsii_.MemberMethod{JsiiMethod: "resetTokenTtl", GoMethod: "ResetTokenTtl"},
			_jsii_.MemberMethod{JsiiMethod: "resetTokenType", GoMethod: "ResetTokenType"},
			_jsii_.MemberMethod{JsiiMethod: "resetUpndomain", GoMethod: "ResetUpndomain"},
			_jsii_.MemberMethod{JsiiMethod: "resetUserattr", GoMethod: "ResetUserattr"},
			_jsii_.MemberMethod{JsiiMethod: "resetUserdn", GoMethod: "ResetUserdn"},
			_jsii_.MemberMethod{JsiiMethod: "resetUserfilter", GoMethod: "ResetUserfilter"},
			_jsii_.MemberMethod{JsiiMethod: "resetUsernameAsAlias", GoMethod: "ResetUsernameAsAlias"},
			_jsii_.MemberMethod{JsiiMethod: "resetUseTokenGroups", GoMethod: "ResetUseTokenGroups"},
			_jsii_.MemberProperty{JsiiProperty: "starttls", GoGetter: "Starttls"},
			_jsii_.MemberProperty{JsiiProperty: "starttlsInput", GoGetter: "StarttlsInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "tlsMaxVersion", GoGetter: "TlsMaxVersion"},
			_jsii_.MemberProperty{JsiiProperty: "tlsMaxVersionInput", GoGetter: "TlsMaxVersionInput"},
			_jsii_.MemberProperty{JsiiProperty: "tlsMinVersion", GoGetter: "TlsMinVersion"},
			_jsii_.MemberProperty{JsiiProperty: "tlsMinVersionInput", GoGetter: "TlsMinVersionInput"},
			_jsii_.MemberProperty{JsiiProperty: "tokenBoundCidrs", GoGetter: "TokenBoundCidrs"},
			_jsii_.MemberProperty{JsiiProperty: "tokenBoundCidrsInput", GoGetter: "TokenBoundCidrsInput"},
			_jsii_.MemberProperty{JsiiProperty: "tokenExplicitMaxTtl", GoGetter: "TokenExplicitMaxTtl"},
			_jsii_.MemberProperty{JsiiProperty: "tokenExplicitMaxTtlInput", GoGetter: "TokenExplicitMaxTtlInput"},
			_jsii_.MemberProperty{JsiiProperty: "tokenMaxTtl", GoGetter: "TokenMaxTtl"},
			_jsii_.MemberProperty{JsiiProperty: "tokenMaxTtlInput", GoGetter: "TokenMaxTtlInput"},
			_jsii_.MemberProperty{JsiiProperty: "tokenNoDefaultPolicy", GoGetter: "TokenNoDefaultPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "tokenNoDefaultPolicyInput", GoGetter: "TokenNoDefaultPolicyInput"},
			_jsii_.MemberProperty{JsiiProperty: "tokenNumUses", GoGetter: "TokenNumUses"},
			_jsii_.MemberProperty{JsiiProperty: "tokenNumUsesInput", GoGetter: "TokenNumUsesInput"},
			_jsii_.MemberProperty{JsiiProperty: "tokenPeriod", GoGetter: "TokenPeriod"},
			_jsii_.MemberProperty{JsiiProperty: "tokenPeriodInput", GoGetter: "TokenPeriodInput"},
			_jsii_.MemberProperty{JsiiProperty: "tokenPolicies", GoGetter: "TokenPolicies"},
			_jsii_.MemberProperty{JsiiProperty: "tokenPoliciesInput", GoGetter: "TokenPoliciesInput"},
			_jsii_.MemberProperty{JsiiProperty: "tokenTtl", GoGetter: "TokenTtl"},
			_jsii_.MemberProperty{JsiiProperty: "tokenTtlInput", GoGetter: "TokenTtlInput"},
			_jsii_.MemberProperty{JsiiProperty: "tokenType", GoGetter: "TokenType"},
			_jsii_.MemberProperty{JsiiProperty: "tokenTypeInput", GoGetter: "TokenTypeInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "userfilter", GoGetter: "Userfilter"},
			_jsii_.MemberProperty{JsiiProperty: "userfilterInput", GoGetter: "UserfilterInput"},
			_jsii_.MemberProperty{JsiiProperty: "usernameAsAlias", GoGetter: "UsernameAsAlias"},
			_jsii_.MemberProperty{JsiiProperty: "usernameAsAliasInput", GoGetter: "UsernameAsAliasInput"},
			_jsii_.MemberProperty{JsiiProperty: "useTokenGroups", GoGetter: "UseTokenGroups"},
			_jsii_.MemberProperty{JsiiProperty: "useTokenGroupsInput", GoGetter: "UseTokenGroupsInput"},
		},
		func() interface{} {
			j := jsiiProxy_LdapAuthBackend{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-vault.ldapAuthBackend.LdapAuthBackendConfig",
		reflect.TypeOf((*LdapAuthBackendConfig)(nil)).Elem(),
	)
}
