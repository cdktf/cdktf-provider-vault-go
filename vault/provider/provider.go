package provider

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-vault.provider.VaultProvider",
		reflect.TypeOf((*VaultProvider)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "addAddressToEnv", GoGetter: "AddAddressToEnv"},
			_jsii_.MemberProperty{JsiiProperty: "addAddressToEnvInput", GoGetter: "AddAddressToEnvInput"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "address", GoGetter: "Address"},
			_jsii_.MemberProperty{JsiiProperty: "addressInput", GoGetter: "AddressInput"},
			_jsii_.MemberProperty{JsiiProperty: "alias", GoGetter: "Alias"},
			_jsii_.MemberProperty{JsiiProperty: "aliasInput", GoGetter: "AliasInput"},
			_jsii_.MemberProperty{JsiiProperty: "authLogin", GoGetter: "AuthLogin"},
			_jsii_.MemberProperty{JsiiProperty: "authLoginInput", GoGetter: "AuthLoginInput"},
			_jsii_.MemberProperty{JsiiProperty: "caCertDir", GoGetter: "CaCertDir"},
			_jsii_.MemberProperty{JsiiProperty: "caCertDirInput", GoGetter: "CaCertDirInput"},
			_jsii_.MemberProperty{JsiiProperty: "caCertFile", GoGetter: "CaCertFile"},
			_jsii_.MemberProperty{JsiiProperty: "caCertFileInput", GoGetter: "CaCertFileInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "clientAuth", GoGetter: "ClientAuth"},
			_jsii_.MemberProperty{JsiiProperty: "clientAuthInput", GoGetter: "ClientAuthInput"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberProperty{JsiiProperty: "headers", GoGetter: "Headers"},
			_jsii_.MemberProperty{JsiiProperty: "headersInput", GoGetter: "HeadersInput"},
			_jsii_.MemberProperty{JsiiProperty: "maxLeaseTtlSeconds", GoGetter: "MaxLeaseTtlSeconds"},
			_jsii_.MemberProperty{JsiiProperty: "maxLeaseTtlSecondsInput", GoGetter: "MaxLeaseTtlSecondsInput"},
			_jsii_.MemberProperty{JsiiProperty: "maxRetries", GoGetter: "MaxRetries"},
			_jsii_.MemberProperty{JsiiProperty: "maxRetriesCcc", GoGetter: "MaxRetriesCcc"},
			_jsii_.MemberProperty{JsiiProperty: "maxRetriesCccInput", GoGetter: "MaxRetriesCccInput"},
			_jsii_.MemberProperty{JsiiProperty: "maxRetriesInput", GoGetter: "MaxRetriesInput"},
			_jsii_.MemberProperty{JsiiProperty: "metaAttributes", GoGetter: "MetaAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "namespace", GoGetter: "Namespace"},
			_jsii_.MemberProperty{JsiiProperty: "namespaceInput", GoGetter: "NamespaceInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetAddAddressToEnv", GoMethod: "ResetAddAddressToEnv"},
			_jsii_.MemberMethod{JsiiMethod: "resetAlias", GoMethod: "ResetAlias"},
			_jsii_.MemberMethod{JsiiMethod: "resetAuthLogin", GoMethod: "ResetAuthLogin"},
			_jsii_.MemberMethod{JsiiMethod: "resetCaCertDir", GoMethod: "ResetCaCertDir"},
			_jsii_.MemberMethod{JsiiMethod: "resetCaCertFile", GoMethod: "ResetCaCertFile"},
			_jsii_.MemberMethod{JsiiMethod: "resetClientAuth", GoMethod: "ResetClientAuth"},
			_jsii_.MemberMethod{JsiiMethod: "resetHeaders", GoMethod: "ResetHeaders"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxLeaseTtlSeconds", GoMethod: "ResetMaxLeaseTtlSeconds"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxRetries", GoMethod: "ResetMaxRetries"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxRetriesCcc", GoMethod: "ResetMaxRetriesCcc"},
			_jsii_.MemberMethod{JsiiMethod: "resetNamespace", GoMethod: "ResetNamespace"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetSkipChildToken", GoMethod: "ResetSkipChildToken"},
			_jsii_.MemberMethod{JsiiMethod: "resetSkipTlsVerify", GoMethod: "ResetSkipTlsVerify"},
			_jsii_.MemberMethod{JsiiMethod: "resetTlsServerName", GoMethod: "ResetTlsServerName"},
			_jsii_.MemberMethod{JsiiMethod: "resetToken", GoMethod: "ResetToken"},
			_jsii_.MemberMethod{JsiiMethod: "resetTokenName", GoMethod: "ResetTokenName"},
			_jsii_.MemberProperty{JsiiProperty: "skipChildToken", GoGetter: "SkipChildToken"},
			_jsii_.MemberProperty{JsiiProperty: "skipChildTokenInput", GoGetter: "SkipChildTokenInput"},
			_jsii_.MemberProperty{JsiiProperty: "skipTlsVerify", GoGetter: "SkipTlsVerify"},
			_jsii_.MemberProperty{JsiiProperty: "skipTlsVerifyInput", GoGetter: "SkipTlsVerifyInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformProviderSource", GoGetter: "TerraformProviderSource"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "tlsServerName", GoGetter: "TlsServerName"},
			_jsii_.MemberProperty{JsiiProperty: "tlsServerNameInput", GoGetter: "TlsServerNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "token", GoGetter: "Token"},
			_jsii_.MemberProperty{JsiiProperty: "tokenInput", GoGetter: "TokenInput"},
			_jsii_.MemberProperty{JsiiProperty: "tokenName", GoGetter: "TokenName"},
			_jsii_.MemberProperty{JsiiProperty: "tokenNameInput", GoGetter: "TokenNameInput"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_VaultProvider{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformProvider)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-vault.provider.VaultProviderAuthLogin",
		reflect.TypeOf((*VaultProviderAuthLogin)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-vault.provider.VaultProviderClientAuth",
		reflect.TypeOf((*VaultProviderClientAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-vault.provider.VaultProviderConfig",
		reflect.TypeOf((*VaultProviderConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-vault.provider.VaultProviderHeaders",
		reflect.TypeOf((*VaultProviderHeaders)(nil)).Elem(),
	)
}
