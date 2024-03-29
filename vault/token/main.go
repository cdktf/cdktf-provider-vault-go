// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package token

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-vault.token.Token",
		reflect.TypeOf((*Token)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "clientToken", GoGetter: "ClientToken"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "displayName", GoGetter: "DisplayName"},
			_jsii_.MemberProperty{JsiiProperty: "displayNameInput", GoGetter: "DisplayNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "explicitMaxTtl", GoGetter: "ExplicitMaxTtl"},
			_jsii_.MemberProperty{JsiiProperty: "explicitMaxTtlInput", GoGetter: "ExplicitMaxTtlInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "leaseDuration", GoGetter: "LeaseDuration"},
			_jsii_.MemberProperty{JsiiProperty: "leaseStarted", GoGetter: "LeaseStarted"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "metadataInput", GoGetter: "MetadataInput"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "namespace", GoGetter: "Namespace"},
			_jsii_.MemberProperty{JsiiProperty: "namespaceInput", GoGetter: "NamespaceInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "noDefaultPolicy", GoGetter: "NoDefaultPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "noDefaultPolicyInput", GoGetter: "NoDefaultPolicyInput"},
			_jsii_.MemberProperty{JsiiProperty: "noParent", GoGetter: "NoParent"},
			_jsii_.MemberProperty{JsiiProperty: "noParentInput", GoGetter: "NoParentInput"},
			_jsii_.MemberProperty{JsiiProperty: "numUses", GoGetter: "NumUses"},
			_jsii_.MemberProperty{JsiiProperty: "numUsesInput", GoGetter: "NumUsesInput"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "period", GoGetter: "Period"},
			_jsii_.MemberProperty{JsiiProperty: "periodInput", GoGetter: "PeriodInput"},
			_jsii_.MemberProperty{JsiiProperty: "policies", GoGetter: "Policies"},
			_jsii_.MemberProperty{JsiiProperty: "policiesInput", GoGetter: "PoliciesInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "renewable", GoGetter: "Renewable"},
			_jsii_.MemberProperty{JsiiProperty: "renewableInput", GoGetter: "RenewableInput"},
			_jsii_.MemberProperty{JsiiProperty: "renewIncrement", GoGetter: "RenewIncrement"},
			_jsii_.MemberProperty{JsiiProperty: "renewIncrementInput", GoGetter: "RenewIncrementInput"},
			_jsii_.MemberProperty{JsiiProperty: "renewMinLease", GoGetter: "RenewMinLease"},
			_jsii_.MemberProperty{JsiiProperty: "renewMinLeaseInput", GoGetter: "RenewMinLeaseInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetDisplayName", GoMethod: "ResetDisplayName"},
			_jsii_.MemberMethod{JsiiMethod: "resetExplicitMaxTtl", GoMethod: "ResetExplicitMaxTtl"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetMetadata", GoMethod: "ResetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "resetNamespace", GoMethod: "ResetNamespace"},
			_jsii_.MemberMethod{JsiiMethod: "resetNoDefaultPolicy", GoMethod: "ResetNoDefaultPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "resetNoParent", GoMethod: "ResetNoParent"},
			_jsii_.MemberMethod{JsiiMethod: "resetNumUses", GoMethod: "ResetNumUses"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPeriod", GoMethod: "ResetPeriod"},
			_jsii_.MemberMethod{JsiiMethod: "resetPolicies", GoMethod: "ResetPolicies"},
			_jsii_.MemberMethod{JsiiMethod: "resetRenewable", GoMethod: "ResetRenewable"},
			_jsii_.MemberMethod{JsiiMethod: "resetRenewIncrement", GoMethod: "ResetRenewIncrement"},
			_jsii_.MemberMethod{JsiiMethod: "resetRenewMinLease", GoMethod: "ResetRenewMinLease"},
			_jsii_.MemberMethod{JsiiMethod: "resetRoleName", GoMethod: "ResetRoleName"},
			_jsii_.MemberMethod{JsiiMethod: "resetTtl", GoMethod: "ResetTtl"},
			_jsii_.MemberMethod{JsiiMethod: "resetWrappingTtl", GoMethod: "ResetWrappingTtl"},
			_jsii_.MemberProperty{JsiiProperty: "roleName", GoGetter: "RoleName"},
			_jsii_.MemberProperty{JsiiProperty: "roleNameInput", GoGetter: "RoleNameInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "ttl", GoGetter: "Ttl"},
			_jsii_.MemberProperty{JsiiProperty: "ttlInput", GoGetter: "TtlInput"},
			_jsii_.MemberProperty{JsiiProperty: "wrappedToken", GoGetter: "WrappedToken"},
			_jsii_.MemberProperty{JsiiProperty: "wrappingAccessor", GoGetter: "WrappingAccessor"},
			_jsii_.MemberProperty{JsiiProperty: "wrappingTtl", GoGetter: "WrappingTtl"},
			_jsii_.MemberProperty{JsiiProperty: "wrappingTtlInput", GoGetter: "WrappingTtlInput"},
		},
		func() interface{} {
			j := jsiiProxy_Token{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-vault.token.TokenConfig",
		reflect.TypeOf((*TokenConfig)(nil)).Elem(),
	)
}
