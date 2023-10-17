// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kmipsecretbackend

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-vault.kmipSecretBackend.KmipSecretBackend",
		reflect.TypeOf((*KmipSecretBackend)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "defaultTlsClientKeyBits", GoGetter: "DefaultTlsClientKeyBits"},
			_jsii_.MemberProperty{JsiiProperty: "defaultTlsClientKeyBitsInput", GoGetter: "DefaultTlsClientKeyBitsInput"},
			_jsii_.MemberProperty{JsiiProperty: "defaultTlsClientKeyType", GoGetter: "DefaultTlsClientKeyType"},
			_jsii_.MemberProperty{JsiiProperty: "defaultTlsClientKeyTypeInput", GoGetter: "DefaultTlsClientKeyTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "defaultTlsClientTtl", GoGetter: "DefaultTlsClientTtl"},
			_jsii_.MemberProperty{JsiiProperty: "defaultTlsClientTtlInput", GoGetter: "DefaultTlsClientTtlInput"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "descriptionInput", GoGetter: "DescriptionInput"},
			_jsii_.MemberProperty{JsiiProperty: "disableRemount", GoGetter: "DisableRemount"},
			_jsii_.MemberProperty{JsiiProperty: "disableRemountInput", GoGetter: "DisableRemountInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "listenAddrs", GoGetter: "ListenAddrs"},
			_jsii_.MemberProperty{JsiiProperty: "listenAddrsInput", GoGetter: "ListenAddrsInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetDefaultTlsClientKeyBits", GoMethod: "ResetDefaultTlsClientKeyBits"},
			_jsii_.MemberMethod{JsiiMethod: "resetDefaultTlsClientKeyType", GoMethod: "ResetDefaultTlsClientKeyType"},
			_jsii_.MemberMethod{JsiiMethod: "resetDefaultTlsClientTtl", GoMethod: "ResetDefaultTlsClientTtl"},
			_jsii_.MemberMethod{JsiiMethod: "resetDescription", GoMethod: "ResetDescription"},
			_jsii_.MemberMethod{JsiiMethod: "resetDisableRemount", GoMethod: "ResetDisableRemount"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetListenAddrs", GoMethod: "ResetListenAddrs"},
			_jsii_.MemberMethod{JsiiMethod: "resetNamespace", GoMethod: "ResetNamespace"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetServerHostnames", GoMethod: "ResetServerHostnames"},
			_jsii_.MemberMethod{JsiiMethod: "resetServerIps", GoMethod: "ResetServerIps"},
			_jsii_.MemberMethod{JsiiMethod: "resetTlsCaKeyBits", GoMethod: "ResetTlsCaKeyBits"},
			_jsii_.MemberMethod{JsiiMethod: "resetTlsCaKeyType", GoMethod: "ResetTlsCaKeyType"},
			_jsii_.MemberMethod{JsiiMethod: "resetTlsMinVersion", GoMethod: "ResetTlsMinVersion"},
			_jsii_.MemberProperty{JsiiProperty: "serverHostnames", GoGetter: "ServerHostnames"},
			_jsii_.MemberProperty{JsiiProperty: "serverHostnamesInput", GoGetter: "ServerHostnamesInput"},
			_jsii_.MemberProperty{JsiiProperty: "serverIps", GoGetter: "ServerIps"},
			_jsii_.MemberProperty{JsiiProperty: "serverIpsInput", GoGetter: "ServerIpsInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "tlsCaKeyBits", GoGetter: "TlsCaKeyBits"},
			_jsii_.MemberProperty{JsiiProperty: "tlsCaKeyBitsInput", GoGetter: "TlsCaKeyBitsInput"},
			_jsii_.MemberProperty{JsiiProperty: "tlsCaKeyType", GoGetter: "TlsCaKeyType"},
			_jsii_.MemberProperty{JsiiProperty: "tlsCaKeyTypeInput", GoGetter: "TlsCaKeyTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "tlsMinVersion", GoGetter: "TlsMinVersion"},
			_jsii_.MemberProperty{JsiiProperty: "tlsMinVersionInput", GoGetter: "TlsMinVersionInput"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_KmipSecretBackend{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-vault.kmipSecretBackend.KmipSecretBackendConfig",
		reflect.TypeOf((*KmipSecretBackendConfig)(nil)).Elem(),
	)
}
