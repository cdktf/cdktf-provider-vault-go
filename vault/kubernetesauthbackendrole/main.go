// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kubernetesauthbackendrole

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-vault.kubernetesAuthBackendRole.KubernetesAuthBackendRole",
		reflect.TypeOf((*KubernetesAuthBackendRole)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "aliasNameSource", GoGetter: "AliasNameSource"},
			_jsii_.MemberProperty{JsiiProperty: "aliasNameSourceInput", GoGetter: "AliasNameSourceInput"},
			_jsii_.MemberProperty{JsiiProperty: "audience", GoGetter: "Audience"},
			_jsii_.MemberProperty{JsiiProperty: "audienceInput", GoGetter: "AudienceInput"},
			_jsii_.MemberProperty{JsiiProperty: "backend", GoGetter: "Backend"},
			_jsii_.MemberProperty{JsiiProperty: "backendInput", GoGetter: "BackendInput"},
			_jsii_.MemberProperty{JsiiProperty: "boundServiceAccountNames", GoGetter: "BoundServiceAccountNames"},
			_jsii_.MemberProperty{JsiiProperty: "boundServiceAccountNamesInput", GoGetter: "BoundServiceAccountNamesInput"},
			_jsii_.MemberProperty{JsiiProperty: "boundServiceAccountNamespaces", GoGetter: "BoundServiceAccountNamespaces"},
			_jsii_.MemberProperty{JsiiProperty: "boundServiceAccountNamespacesInput", GoGetter: "BoundServiceAccountNamespacesInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
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
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "namespace", GoGetter: "Namespace"},
			_jsii_.MemberProperty{JsiiProperty: "namespaceInput", GoGetter: "NamespaceInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetAliasNameSource", GoMethod: "ResetAliasNameSource"},
			_jsii_.MemberMethod{JsiiMethod: "resetAudience", GoMethod: "ResetAudience"},
			_jsii_.MemberMethod{JsiiMethod: "resetBackend", GoMethod: "ResetBackend"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetNamespace", GoMethod: "ResetNamespace"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetTokenBoundCidrs", GoMethod: "ResetTokenBoundCidrs"},
			_jsii_.MemberMethod{JsiiMethod: "resetTokenExplicitMaxTtl", GoMethod: "ResetTokenExplicitMaxTtl"},
			_jsii_.MemberMethod{JsiiMethod: "resetTokenMaxTtl", GoMethod: "ResetTokenMaxTtl"},
			_jsii_.MemberMethod{JsiiMethod: "resetTokenNoDefaultPolicy", GoMethod: "ResetTokenNoDefaultPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "resetTokenNumUses", GoMethod: "ResetTokenNumUses"},
			_jsii_.MemberMethod{JsiiMethod: "resetTokenPeriod", GoMethod: "ResetTokenPeriod"},
			_jsii_.MemberMethod{JsiiMethod: "resetTokenPolicies", GoMethod: "ResetTokenPolicies"},
			_jsii_.MemberMethod{JsiiMethod: "resetTokenTtl", GoMethod: "ResetTokenTtl"},
			_jsii_.MemberMethod{JsiiMethod: "resetTokenType", GoMethod: "ResetTokenType"},
			_jsii_.MemberProperty{JsiiProperty: "roleName", GoGetter: "RoleName"},
			_jsii_.MemberProperty{JsiiProperty: "roleNameInput", GoGetter: "RoleNameInput"},
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
			j := jsiiProxy_KubernetesAuthBackendRole{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-vault.kubernetesAuthBackendRole.KubernetesAuthBackendRoleConfig",
		reflect.TypeOf((*KubernetesAuthBackendRoleConfig)(nil)).Elem(),
	)
}
