// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package certauthbackendrole

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-vault.certAuthBackendRole.CertAuthBackendRole",
		reflect.TypeOf((*CertAuthBackendRole)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "allowedCommonNames", GoGetter: "AllowedCommonNames"},
			_jsii_.MemberProperty{JsiiProperty: "allowedCommonNamesInput", GoGetter: "AllowedCommonNamesInput"},
			_jsii_.MemberProperty{JsiiProperty: "allowedDnsSans", GoGetter: "AllowedDnsSans"},
			_jsii_.MemberProperty{JsiiProperty: "allowedDnsSansInput", GoGetter: "AllowedDnsSansInput"},
			_jsii_.MemberProperty{JsiiProperty: "allowedEmailSans", GoGetter: "AllowedEmailSans"},
			_jsii_.MemberProperty{JsiiProperty: "allowedEmailSansInput", GoGetter: "AllowedEmailSansInput"},
			_jsii_.MemberProperty{JsiiProperty: "allowedNames", GoGetter: "AllowedNames"},
			_jsii_.MemberProperty{JsiiProperty: "allowedNamesInput", GoGetter: "AllowedNamesInput"},
			_jsii_.MemberProperty{JsiiProperty: "allowedOrganizationalUnits", GoGetter: "AllowedOrganizationalUnits"},
			_jsii_.MemberProperty{JsiiProperty: "allowedOrganizationalUnitsInput", GoGetter: "AllowedOrganizationalUnitsInput"},
			_jsii_.MemberProperty{JsiiProperty: "allowedOrganizationUnits", GoGetter: "AllowedOrganizationUnits"},
			_jsii_.MemberProperty{JsiiProperty: "allowedOrganizationUnitsInput", GoGetter: "AllowedOrganizationUnitsInput"},
			_jsii_.MemberProperty{JsiiProperty: "allowedUriSans", GoGetter: "AllowedUriSans"},
			_jsii_.MemberProperty{JsiiProperty: "allowedUriSansInput", GoGetter: "AllowedUriSansInput"},
			_jsii_.MemberProperty{JsiiProperty: "backend", GoGetter: "Backend"},
			_jsii_.MemberProperty{JsiiProperty: "backendInput", GoGetter: "BackendInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "certificate", GoGetter: "Certificate"},
			_jsii_.MemberProperty{JsiiProperty: "certificateInput", GoGetter: "CertificateInput"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "displayName", GoGetter: "DisplayName"},
			_jsii_.MemberProperty{JsiiProperty: "displayNameInput", GoGetter: "DisplayNameInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "requiredExtensions", GoGetter: "RequiredExtensions"},
			_jsii_.MemberProperty{JsiiProperty: "requiredExtensionsInput", GoGetter: "RequiredExtensionsInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowedCommonNames", GoMethod: "ResetAllowedCommonNames"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowedDnsSans", GoMethod: "ResetAllowedDnsSans"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowedEmailSans", GoMethod: "ResetAllowedEmailSans"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowedNames", GoMethod: "ResetAllowedNames"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowedOrganizationalUnits", GoMethod: "ResetAllowedOrganizationalUnits"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowedOrganizationUnits", GoMethod: "ResetAllowedOrganizationUnits"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowedUriSans", GoMethod: "ResetAllowedUriSans"},
			_jsii_.MemberMethod{JsiiMethod: "resetBackend", GoMethod: "ResetBackend"},
			_jsii_.MemberMethod{JsiiMethod: "resetDisplayName", GoMethod: "ResetDisplayName"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetNamespace", GoMethod: "ResetNamespace"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetRequiredExtensions", GoMethod: "ResetRequiredExtensions"},
			_jsii_.MemberMethod{JsiiMethod: "resetTokenBoundCidrs", GoMethod: "ResetTokenBoundCidrs"},
			_jsii_.MemberMethod{JsiiMethod: "resetTokenExplicitMaxTtl", GoMethod: "ResetTokenExplicitMaxTtl"},
			_jsii_.MemberMethod{JsiiMethod: "resetTokenMaxTtl", GoMethod: "ResetTokenMaxTtl"},
			_jsii_.MemberMethod{JsiiMethod: "resetTokenNoDefaultPolicy", GoMethod: "ResetTokenNoDefaultPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "resetTokenNumUses", GoMethod: "ResetTokenNumUses"},
			_jsii_.MemberMethod{JsiiMethod: "resetTokenPeriod", GoMethod: "ResetTokenPeriod"},
			_jsii_.MemberMethod{JsiiMethod: "resetTokenPolicies", GoMethod: "ResetTokenPolicies"},
			_jsii_.MemberMethod{JsiiMethod: "resetTokenTtl", GoMethod: "ResetTokenTtl"},
			_jsii_.MemberMethod{JsiiMethod: "resetTokenType", GoMethod: "ResetTokenType"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
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
		},
		func() interface{} {
			j := jsiiProxy_CertAuthBackendRole{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-vault.certAuthBackendRole.CertAuthBackendRoleConfig",
		reflect.TypeOf((*CertAuthBackendRoleConfig)(nil)).Elem(),
	)
}
