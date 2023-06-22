package pkisecretbackendsign

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-vault.pkiSecretBackendSign.PkiSecretBackendSign",
		reflect.TypeOf((*PkiSecretBackendSign)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "altNames", GoGetter: "AltNames"},
			_jsii_.MemberProperty{JsiiProperty: "altNamesInput", GoGetter: "AltNamesInput"},
			_jsii_.MemberProperty{JsiiProperty: "autoRenew", GoGetter: "AutoRenew"},
			_jsii_.MemberProperty{JsiiProperty: "autoRenewInput", GoGetter: "AutoRenewInput"},
			_jsii_.MemberProperty{JsiiProperty: "backend", GoGetter: "Backend"},
			_jsii_.MemberProperty{JsiiProperty: "backendInput", GoGetter: "BackendInput"},
			_jsii_.MemberProperty{JsiiProperty: "caChain", GoGetter: "CaChain"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "certificate", GoGetter: "Certificate"},
			_jsii_.MemberProperty{JsiiProperty: "commonName", GoGetter: "CommonName"},
			_jsii_.MemberProperty{JsiiProperty: "commonNameInput", GoGetter: "CommonNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "csr", GoGetter: "Csr"},
			_jsii_.MemberProperty{JsiiProperty: "csrInput", GoGetter: "CsrInput"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "excludeCnFromSans", GoGetter: "ExcludeCnFromSans"},
			_jsii_.MemberProperty{JsiiProperty: "excludeCnFromSansInput", GoGetter: "ExcludeCnFromSansInput"},
			_jsii_.MemberProperty{JsiiProperty: "expiration", GoGetter: "Expiration"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "format", GoGetter: "Format"},
			_jsii_.MemberProperty{JsiiProperty: "formatInput", GoGetter: "FormatInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "ipSans", GoGetter: "IpSans"},
			_jsii_.MemberProperty{JsiiProperty: "ipSansInput", GoGetter: "IpSansInput"},
			_jsii_.MemberProperty{JsiiProperty: "issuerRef", GoGetter: "IssuerRef"},
			_jsii_.MemberProperty{JsiiProperty: "issuerRefInput", GoGetter: "IssuerRefInput"},
			_jsii_.MemberProperty{JsiiProperty: "issuingCa", GoGetter: "IssuingCa"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "minSecondsRemaining", GoGetter: "MinSecondsRemaining"},
			_jsii_.MemberProperty{JsiiProperty: "minSecondsRemainingInput", GoGetter: "MinSecondsRemainingInput"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "namespace", GoGetter: "Namespace"},
			_jsii_.MemberProperty{JsiiProperty: "namespaceInput", GoGetter: "NamespaceInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "otherSans", GoGetter: "OtherSans"},
			_jsii_.MemberProperty{JsiiProperty: "otherSansInput", GoGetter: "OtherSansInput"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "renewPending", GoGetter: "RenewPending"},
			_jsii_.MemberMethod{JsiiMethod: "resetAltNames", GoMethod: "ResetAltNames"},
			_jsii_.MemberMethod{JsiiMethod: "resetAutoRenew", GoMethod: "ResetAutoRenew"},
			_jsii_.MemberMethod{JsiiMethod: "resetExcludeCnFromSans", GoMethod: "ResetExcludeCnFromSans"},
			_jsii_.MemberMethod{JsiiMethod: "resetFormat", GoMethod: "ResetFormat"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetIpSans", GoMethod: "ResetIpSans"},
			_jsii_.MemberMethod{JsiiMethod: "resetIssuerRef", GoMethod: "ResetIssuerRef"},
			_jsii_.MemberMethod{JsiiMethod: "resetMinSecondsRemaining", GoMethod: "ResetMinSecondsRemaining"},
			_jsii_.MemberMethod{JsiiMethod: "resetNamespace", GoMethod: "ResetNamespace"},
			_jsii_.MemberMethod{JsiiMethod: "resetOtherSans", GoMethod: "ResetOtherSans"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetTtl", GoMethod: "ResetTtl"},
			_jsii_.MemberMethod{JsiiMethod: "resetUriSans", GoMethod: "ResetUriSans"},
			_jsii_.MemberProperty{JsiiProperty: "serial", GoGetter: "Serial"},
			_jsii_.MemberProperty{JsiiProperty: "serialNumber", GoGetter: "SerialNumber"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "ttl", GoGetter: "Ttl"},
			_jsii_.MemberProperty{JsiiProperty: "ttlInput", GoGetter: "TtlInput"},
			_jsii_.MemberProperty{JsiiProperty: "uriSans", GoGetter: "UriSans"},
			_jsii_.MemberProperty{JsiiProperty: "uriSansInput", GoGetter: "UriSansInput"},
		},
		func() interface{} {
			j := jsiiProxy_PkiSecretBackendSign{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-vault.pkiSecretBackendSign.PkiSecretBackendSignConfig",
		reflect.TypeOf((*PkiSecretBackendSignConfig)(nil)).Elem(),
	)
}
