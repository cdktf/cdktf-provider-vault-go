// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pkisecretbackendissuer

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-vault.pkiSecretBackendIssuer.PkiSecretBackendIssuer",
		reflect.TypeOf((*PkiSecretBackendIssuer)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "backend", GoGetter: "Backend"},
			_jsii_.MemberProperty{JsiiProperty: "backendInput", GoGetter: "BackendInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "crlDistributionPoints", GoGetter: "CrlDistributionPoints"},
			_jsii_.MemberProperty{JsiiProperty: "crlDistributionPointsInput", GoGetter: "CrlDistributionPointsInput"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "enableAiaUrlTemplating", GoGetter: "EnableAiaUrlTemplating"},
			_jsii_.MemberProperty{JsiiProperty: "enableAiaUrlTemplatingInput", GoGetter: "EnableAiaUrlTemplatingInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "hasResourceMove", GoMethod: "HasResourceMove"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "issuerId", GoGetter: "IssuerId"},
			_jsii_.MemberProperty{JsiiProperty: "issuerName", GoGetter: "IssuerName"},
			_jsii_.MemberProperty{JsiiProperty: "issuerNameInput", GoGetter: "IssuerNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "issuerRef", GoGetter: "IssuerRef"},
			_jsii_.MemberProperty{JsiiProperty: "issuerRefInput", GoGetter: "IssuerRefInput"},
			_jsii_.MemberProperty{JsiiProperty: "issuingCertificates", GoGetter: "IssuingCertificates"},
			_jsii_.MemberProperty{JsiiProperty: "issuingCertificatesInput", GoGetter: "IssuingCertificatesInput"},
			_jsii_.MemberProperty{JsiiProperty: "leafNotAfterBehavior", GoGetter: "LeafNotAfterBehavior"},
			_jsii_.MemberProperty{JsiiProperty: "leafNotAfterBehaviorInput", GoGetter: "LeafNotAfterBehaviorInput"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "manualChain", GoGetter: "ManualChain"},
			_jsii_.MemberProperty{JsiiProperty: "manualChainInput", GoGetter: "ManualChainInput"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "namespace", GoGetter: "Namespace"},
			_jsii_.MemberProperty{JsiiProperty: "namespaceInput", GoGetter: "NamespaceInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "ocspServers", GoGetter: "OcspServers"},
			_jsii_.MemberProperty{JsiiProperty: "ocspServersInput", GoGetter: "OcspServersInput"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetCrlDistributionPoints", GoMethod: "ResetCrlDistributionPoints"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnableAiaUrlTemplating", GoMethod: "ResetEnableAiaUrlTemplating"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetIssuerName", GoMethod: "ResetIssuerName"},
			_jsii_.MemberMethod{JsiiMethod: "resetIssuingCertificates", GoMethod: "ResetIssuingCertificates"},
			_jsii_.MemberMethod{JsiiMethod: "resetLeafNotAfterBehavior", GoMethod: "ResetLeafNotAfterBehavior"},
			_jsii_.MemberMethod{JsiiMethod: "resetManualChain", GoMethod: "ResetManualChain"},
			_jsii_.MemberMethod{JsiiMethod: "resetNamespace", GoMethod: "ResetNamespace"},
			_jsii_.MemberMethod{JsiiMethod: "resetOcspServers", GoMethod: "ResetOcspServers"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetRevocationSignatureAlgorithm", GoMethod: "ResetRevocationSignatureAlgorithm"},
			_jsii_.MemberMethod{JsiiMethod: "resetUsage", GoMethod: "ResetUsage"},
			_jsii_.MemberProperty{JsiiProperty: "revocationSignatureAlgorithm", GoGetter: "RevocationSignatureAlgorithm"},
			_jsii_.MemberProperty{JsiiProperty: "revocationSignatureAlgorithmInput", GoGetter: "RevocationSignatureAlgorithmInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "usage", GoGetter: "Usage"},
			_jsii_.MemberProperty{JsiiProperty: "usageInput", GoGetter: "UsageInput"},
		},
		func() interface{} {
			j := jsiiProxy_PkiSecretBackendIssuer{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-vault.pkiSecretBackendIssuer.PkiSecretBackendIssuerConfig",
		reflect.TypeOf((*PkiSecretBackendIssuerConfig)(nil)).Elem(),
	)
}
