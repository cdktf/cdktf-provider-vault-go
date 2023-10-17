// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package transitsecretbackendkey

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-vault.transitSecretBackendKey.TransitSecretBackendKey",
		reflect.TypeOf((*TransitSecretBackendKey)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "allowPlaintextBackup", GoGetter: "AllowPlaintextBackup"},
			_jsii_.MemberProperty{JsiiProperty: "allowPlaintextBackupInput", GoGetter: "AllowPlaintextBackupInput"},
			_jsii_.MemberProperty{JsiiProperty: "autoRotateInterval", GoGetter: "AutoRotateInterval"},
			_jsii_.MemberProperty{JsiiProperty: "autoRotateIntervalInput", GoGetter: "AutoRotateIntervalInput"},
			_jsii_.MemberProperty{JsiiProperty: "autoRotatePeriod", GoGetter: "AutoRotatePeriod"},
			_jsii_.MemberProperty{JsiiProperty: "autoRotatePeriodInput", GoGetter: "AutoRotatePeriodInput"},
			_jsii_.MemberProperty{JsiiProperty: "backend", GoGetter: "Backend"},
			_jsii_.MemberProperty{JsiiProperty: "backendInput", GoGetter: "BackendInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "convergentEncryption", GoGetter: "ConvergentEncryption"},
			_jsii_.MemberProperty{JsiiProperty: "convergentEncryptionInput", GoGetter: "ConvergentEncryptionInput"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "deletionAllowed", GoGetter: "DeletionAllowed"},
			_jsii_.MemberProperty{JsiiProperty: "deletionAllowedInput", GoGetter: "DeletionAllowedInput"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "derived", GoGetter: "Derived"},
			_jsii_.MemberProperty{JsiiProperty: "derivedInput", GoGetter: "DerivedInput"},
			_jsii_.MemberProperty{JsiiProperty: "exportable", GoGetter: "Exportable"},
			_jsii_.MemberProperty{JsiiProperty: "exportableInput", GoGetter: "ExportableInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "keys", GoGetter: "Keys"},
			_jsii_.MemberProperty{JsiiProperty: "keySize", GoGetter: "KeySize"},
			_jsii_.MemberProperty{JsiiProperty: "keySizeInput", GoGetter: "KeySizeInput"},
			_jsii_.MemberProperty{JsiiProperty: "latestVersion", GoGetter: "LatestVersion"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "minAvailableVersion", GoGetter: "MinAvailableVersion"},
			_jsii_.MemberProperty{JsiiProperty: "minDecryptionVersion", GoGetter: "MinDecryptionVersion"},
			_jsii_.MemberProperty{JsiiProperty: "minDecryptionVersionInput", GoGetter: "MinDecryptionVersionInput"},
			_jsii_.MemberProperty{JsiiProperty: "minEncryptionVersion", GoGetter: "MinEncryptionVersion"},
			_jsii_.MemberProperty{JsiiProperty: "minEncryptionVersionInput", GoGetter: "MinEncryptionVersionInput"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "namespace", GoGetter: "Namespace"},
			_jsii_.MemberProperty{JsiiProperty: "namespaceInput", GoGetter: "NamespaceInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowPlaintextBackup", GoMethod: "ResetAllowPlaintextBackup"},
			_jsii_.MemberMethod{JsiiMethod: "resetAutoRotateInterval", GoMethod: "ResetAutoRotateInterval"},
			_jsii_.MemberMethod{JsiiMethod: "resetAutoRotatePeriod", GoMethod: "ResetAutoRotatePeriod"},
			_jsii_.MemberMethod{JsiiMethod: "resetConvergentEncryption", GoMethod: "ResetConvergentEncryption"},
			_jsii_.MemberMethod{JsiiMethod: "resetDeletionAllowed", GoMethod: "ResetDeletionAllowed"},
			_jsii_.MemberMethod{JsiiMethod: "resetDerived", GoMethod: "ResetDerived"},
			_jsii_.MemberMethod{JsiiMethod: "resetExportable", GoMethod: "ResetExportable"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetKeySize", GoMethod: "ResetKeySize"},
			_jsii_.MemberMethod{JsiiMethod: "resetMinDecryptionVersion", GoMethod: "ResetMinDecryptionVersion"},
			_jsii_.MemberMethod{JsiiMethod: "resetMinEncryptionVersion", GoMethod: "ResetMinEncryptionVersion"},
			_jsii_.MemberMethod{JsiiMethod: "resetNamespace", GoMethod: "ResetNamespace"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetType", GoMethod: "ResetType"},
			_jsii_.MemberProperty{JsiiProperty: "supportsDecryption", GoGetter: "SupportsDecryption"},
			_jsii_.MemberProperty{JsiiProperty: "supportsDerivation", GoGetter: "SupportsDerivation"},
			_jsii_.MemberProperty{JsiiProperty: "supportsEncryption", GoGetter: "SupportsEncryption"},
			_jsii_.MemberProperty{JsiiProperty: "supportsSigning", GoGetter: "SupportsSigning"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "typeInput", GoGetter: "TypeInput"},
		},
		func() interface{} {
			j := jsiiProxy_TransitSecretBackendKey{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-vault.transitSecretBackendKey.TransitSecretBackendKeyConfig",
		reflect.TypeOf((*TransitSecretBackendKeyConfig)(nil)).Elem(),
	)
}
