// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datavaulttransitsign

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-vault.dataVaultTransitSign.DataVaultTransitSign",
		reflect.TypeOf((*DataVaultTransitSign)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "batchInput", GoGetter: "BatchInput"},
			_jsii_.MemberProperty{JsiiProperty: "batchInputInput", GoGetter: "BatchInputInput"},
			_jsii_.MemberProperty{JsiiProperty: "batchResults", GoGetter: "BatchResults"},
			_jsii_.MemberProperty{JsiiProperty: "batchResultsInput", GoGetter: "BatchResultsInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "context", GoGetter: "Context"},
			_jsii_.MemberProperty{JsiiProperty: "contextInput", GoGetter: "ContextInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "hashAlgorithm", GoGetter: "HashAlgorithm"},
			_jsii_.MemberProperty{JsiiProperty: "hashAlgorithmInput", GoGetter: "HashAlgorithmInput"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberProperty{JsiiProperty: "input", GoGetter: "Input"},
			_jsii_.MemberProperty{JsiiProperty: "inputInput", GoGetter: "InputInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "keyVersion", GoGetter: "KeyVersion"},
			_jsii_.MemberProperty{JsiiProperty: "keyVersionInput", GoGetter: "KeyVersionInput"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "marshalingAlgorithm", GoGetter: "MarshalingAlgorithm"},
			_jsii_.MemberProperty{JsiiProperty: "marshalingAlgorithmInput", GoGetter: "MarshalingAlgorithmInput"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "namespace", GoGetter: "Namespace"},
			_jsii_.MemberProperty{JsiiProperty: "namespaceInput", GoGetter: "NamespaceInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "path", GoGetter: "Path"},
			_jsii_.MemberProperty{JsiiProperty: "pathInput", GoGetter: "PathInput"},
			_jsii_.MemberProperty{JsiiProperty: "prehashed", GoGetter: "Prehashed"},
			_jsii_.MemberProperty{JsiiProperty: "prehashedInput", GoGetter: "PrehashedInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "reference", GoGetter: "Reference"},
			_jsii_.MemberProperty{JsiiProperty: "referenceInput", GoGetter: "ReferenceInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetBatchInput", GoMethod: "ResetBatchInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetBatchResults", GoMethod: "ResetBatchResults"},
			_jsii_.MemberMethod{JsiiMethod: "resetContext", GoMethod: "ResetContext"},
			_jsii_.MemberMethod{JsiiMethod: "resetHashAlgorithm", GoMethod: "ResetHashAlgorithm"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetInput", GoMethod: "ResetInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetKeyVersion", GoMethod: "ResetKeyVersion"},
			_jsii_.MemberMethod{JsiiMethod: "resetMarshalingAlgorithm", GoMethod: "ResetMarshalingAlgorithm"},
			_jsii_.MemberMethod{JsiiMethod: "resetNamespace", GoMethod: "ResetNamespace"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrehashed", GoMethod: "ResetPrehashed"},
			_jsii_.MemberMethod{JsiiMethod: "resetReference", GoMethod: "ResetReference"},
			_jsii_.MemberMethod{JsiiMethod: "resetSaltLength", GoMethod: "ResetSaltLength"},
			_jsii_.MemberMethod{JsiiMethod: "resetSignature", GoMethod: "ResetSignature"},
			_jsii_.MemberMethod{JsiiMethod: "resetSignatureAlgorithm", GoMethod: "ResetSignatureAlgorithm"},
			_jsii_.MemberMethod{JsiiMethod: "resetSignatureContext", GoMethod: "ResetSignatureContext"},
			_jsii_.MemberProperty{JsiiProperty: "saltLength", GoGetter: "SaltLength"},
			_jsii_.MemberProperty{JsiiProperty: "saltLengthInput", GoGetter: "SaltLengthInput"},
			_jsii_.MemberProperty{JsiiProperty: "signature", GoGetter: "Signature"},
			_jsii_.MemberProperty{JsiiProperty: "signatureAlgorithm", GoGetter: "SignatureAlgorithm"},
			_jsii_.MemberProperty{JsiiProperty: "signatureAlgorithmInput", GoGetter: "SignatureAlgorithmInput"},
			_jsii_.MemberProperty{JsiiProperty: "signatureContext", GoGetter: "SignatureContext"},
			_jsii_.MemberProperty{JsiiProperty: "signatureContextInput", GoGetter: "SignatureContextInput"},
			_jsii_.MemberProperty{JsiiProperty: "signatureInput", GoGetter: "SignatureInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_DataVaultTransitSign{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformDataSource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-vault.dataVaultTransitSign.DataVaultTransitSignConfig",
		reflect.TypeOf((*DataVaultTransitSignConfig)(nil)).Elem(),
	)
}
