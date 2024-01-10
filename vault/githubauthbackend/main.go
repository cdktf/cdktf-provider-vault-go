// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package githubauthbackend

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-vault.githubAuthBackend.GithubAuthBackend",
		reflect.TypeOf((*GithubAuthBackend)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accessor", GoGetter: "Accessor"},
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "baseUrl", GoGetter: "BaseUrl"},
			_jsii_.MemberProperty{JsiiProperty: "baseUrlInput", GoGetter: "BaseUrlInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
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
			_jsii_.MemberMethod{JsiiMethod: "hasResourceMove", GoMethod: "HasResourceMove"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "namespace", GoGetter: "Namespace"},
			_jsii_.MemberProperty{JsiiProperty: "namespaceInput", GoGetter: "NamespaceInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "organization", GoGetter: "Organization"},
			_jsii_.MemberProperty{JsiiProperty: "organizationId", GoGetter: "OrganizationId"},
			_jsii_.MemberProperty{JsiiProperty: "organizationIdInput", GoGetter: "OrganizationIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "organizationInput", GoGetter: "OrganizationInput"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "path", GoGetter: "Path"},
			_jsii_.MemberProperty{JsiiProperty: "pathInput", GoGetter: "PathInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putTune", GoMethod: "PutTune"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetBaseUrl", GoMethod: "ResetBaseUrl"},
			_jsii_.MemberMethod{JsiiMethod: "resetDescription", GoMethod: "ResetDescription"},
			_jsii_.MemberMethod{JsiiMethod: "resetDisableRemount", GoMethod: "ResetDisableRemount"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetNamespace", GoMethod: "ResetNamespace"},
			_jsii_.MemberMethod{JsiiMethod: "resetOrganizationId", GoMethod: "ResetOrganizationId"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPath", GoMethod: "ResetPath"},
			_jsii_.MemberMethod{JsiiMethod: "resetTokenBoundCidrs", GoMethod: "ResetTokenBoundCidrs"},
			_jsii_.MemberMethod{JsiiMethod: "resetTokenExplicitMaxTtl", GoMethod: "ResetTokenExplicitMaxTtl"},
			_jsii_.MemberMethod{JsiiMethod: "resetTokenMaxTtl", GoMethod: "ResetTokenMaxTtl"},
			_jsii_.MemberMethod{JsiiMethod: "resetTokenNoDefaultPolicy", GoMethod: "ResetTokenNoDefaultPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "resetTokenNumUses", GoMethod: "ResetTokenNumUses"},
			_jsii_.MemberMethod{JsiiMethod: "resetTokenPeriod", GoMethod: "ResetTokenPeriod"},
			_jsii_.MemberMethod{JsiiMethod: "resetTokenPolicies", GoMethod: "ResetTokenPolicies"},
			_jsii_.MemberMethod{JsiiMethod: "resetTokenTtl", GoMethod: "ResetTokenTtl"},
			_jsii_.MemberMethod{JsiiMethod: "resetTokenType", GoMethod: "ResetTokenType"},
			_jsii_.MemberMethod{JsiiMethod: "resetTune", GoMethod: "ResetTune"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
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
			_jsii_.MemberProperty{JsiiProperty: "tune", GoGetter: "Tune"},
			_jsii_.MemberProperty{JsiiProperty: "tuneInput", GoGetter: "TuneInput"},
		},
		func() interface{} {
			j := jsiiProxy_GithubAuthBackend{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-vault.githubAuthBackend.GithubAuthBackendConfig",
		reflect.TypeOf((*GithubAuthBackendConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-vault.githubAuthBackend.GithubAuthBackendTune",
		reflect.TypeOf((*GithubAuthBackendTune)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-vault.githubAuthBackend.GithubAuthBackendTuneList",
		reflect.TypeOf((*GithubAuthBackendTuneList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_GithubAuthBackendTuneList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-vault.githubAuthBackend.GithubAuthBackendTuneOutputReference",
		reflect.TypeOf((*GithubAuthBackendTuneOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allowedResponseHeaders", GoGetter: "AllowedResponseHeaders"},
			_jsii_.MemberProperty{JsiiProperty: "allowedResponseHeadersInput", GoGetter: "AllowedResponseHeadersInput"},
			_jsii_.MemberProperty{JsiiProperty: "auditNonHmacRequestKeys", GoGetter: "AuditNonHmacRequestKeys"},
			_jsii_.MemberProperty{JsiiProperty: "auditNonHmacRequestKeysInput", GoGetter: "AuditNonHmacRequestKeysInput"},
			_jsii_.MemberProperty{JsiiProperty: "auditNonHmacResponseKeys", GoGetter: "AuditNonHmacResponseKeys"},
			_jsii_.MemberProperty{JsiiProperty: "auditNonHmacResponseKeysInput", GoGetter: "AuditNonHmacResponseKeysInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "defaultLeaseTtl", GoGetter: "DefaultLeaseTtl"},
			_jsii_.MemberProperty{JsiiProperty: "defaultLeaseTtlInput", GoGetter: "DefaultLeaseTtlInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "listingVisibility", GoGetter: "ListingVisibility"},
			_jsii_.MemberProperty{JsiiProperty: "listingVisibilityInput", GoGetter: "ListingVisibilityInput"},
			_jsii_.MemberProperty{JsiiProperty: "maxLeaseTtl", GoGetter: "MaxLeaseTtl"},
			_jsii_.MemberProperty{JsiiProperty: "maxLeaseTtlInput", GoGetter: "MaxLeaseTtlInput"},
			_jsii_.MemberProperty{JsiiProperty: "passthroughRequestHeaders", GoGetter: "PassthroughRequestHeaders"},
			_jsii_.MemberProperty{JsiiProperty: "passthroughRequestHeadersInput", GoGetter: "PassthroughRequestHeadersInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowedResponseHeaders", GoMethod: "ResetAllowedResponseHeaders"},
			_jsii_.MemberMethod{JsiiMethod: "resetAuditNonHmacRequestKeys", GoMethod: "ResetAuditNonHmacRequestKeys"},
			_jsii_.MemberMethod{JsiiMethod: "resetAuditNonHmacResponseKeys", GoMethod: "ResetAuditNonHmacResponseKeys"},
			_jsii_.MemberMethod{JsiiMethod: "resetDefaultLeaseTtl", GoMethod: "ResetDefaultLeaseTtl"},
			_jsii_.MemberMethod{JsiiMethod: "resetListingVisibility", GoMethod: "ResetListingVisibility"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxLeaseTtl", GoMethod: "ResetMaxLeaseTtl"},
			_jsii_.MemberMethod{JsiiMethod: "resetPassthroughRequestHeaders", GoMethod: "ResetPassthroughRequestHeaders"},
			_jsii_.MemberMethod{JsiiMethod: "resetTokenType", GoMethod: "ResetTokenType"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "tokenType", GoGetter: "TokenType"},
			_jsii_.MemberProperty{JsiiProperty: "tokenTypeInput", GoGetter: "TokenTypeInput"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GithubAuthBackendTuneOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
