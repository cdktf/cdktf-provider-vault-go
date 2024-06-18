// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package awssecretbackendrole

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-vault.awsSecretBackendRole.AwsSecretBackendRole",
		reflect.TypeOf((*AwsSecretBackendRole)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "backend", GoGetter: "Backend"},
			_jsii_.MemberProperty{JsiiProperty: "backendInput", GoGetter: "BackendInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "credentialType", GoGetter: "CredentialType"},
			_jsii_.MemberProperty{JsiiProperty: "credentialTypeInput", GoGetter: "CredentialTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "defaultStsTtl", GoGetter: "DefaultStsTtl"},
			_jsii_.MemberProperty{JsiiProperty: "defaultStsTtlInput", GoGetter: "DefaultStsTtlInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "hasResourceMove", GoMethod: "HasResourceMove"},
			_jsii_.MemberProperty{JsiiProperty: "iamGroups", GoGetter: "IamGroups"},
			_jsii_.MemberProperty{JsiiProperty: "iamGroupsInput", GoGetter: "IamGroupsInput"},
			_jsii_.MemberProperty{JsiiProperty: "iamTags", GoGetter: "IamTags"},
			_jsii_.MemberProperty{JsiiProperty: "iamTagsInput", GoGetter: "IamTagsInput"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "maxStsTtl", GoGetter: "MaxStsTtl"},
			_jsii_.MemberProperty{JsiiProperty: "maxStsTtlInput", GoGetter: "MaxStsTtlInput"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "namespace", GoGetter: "Namespace"},
			_jsii_.MemberProperty{JsiiProperty: "namespaceInput", GoGetter: "NamespaceInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "permissionsBoundaryArn", GoGetter: "PermissionsBoundaryArn"},
			_jsii_.MemberProperty{JsiiProperty: "permissionsBoundaryArnInput", GoGetter: "PermissionsBoundaryArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "policyArns", GoGetter: "PolicyArns"},
			_jsii_.MemberProperty{JsiiProperty: "policyArnsInput", GoGetter: "PolicyArnsInput"},
			_jsii_.MemberProperty{JsiiProperty: "policyDocument", GoGetter: "PolicyDocument"},
			_jsii_.MemberProperty{JsiiProperty: "policyDocumentInput", GoGetter: "PolicyDocumentInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetDefaultStsTtl", GoMethod: "ResetDefaultStsTtl"},
			_jsii_.MemberMethod{JsiiMethod: "resetIamGroups", GoMethod: "ResetIamGroups"},
			_jsii_.MemberMethod{JsiiMethod: "resetIamTags", GoMethod: "ResetIamTags"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxStsTtl", GoMethod: "ResetMaxStsTtl"},
			_jsii_.MemberMethod{JsiiMethod: "resetNamespace", GoMethod: "ResetNamespace"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPermissionsBoundaryArn", GoMethod: "ResetPermissionsBoundaryArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetPolicyArns", GoMethod: "ResetPolicyArns"},
			_jsii_.MemberMethod{JsiiMethod: "resetPolicyDocument", GoMethod: "ResetPolicyDocument"},
			_jsii_.MemberMethod{JsiiMethod: "resetRoleArns", GoMethod: "ResetRoleArns"},
			_jsii_.MemberMethod{JsiiMethod: "resetUserPath", GoMethod: "ResetUserPath"},
			_jsii_.MemberProperty{JsiiProperty: "roleArns", GoGetter: "RoleArns"},
			_jsii_.MemberProperty{JsiiProperty: "roleArnsInput", GoGetter: "RoleArnsInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "userPath", GoGetter: "UserPath"},
			_jsii_.MemberProperty{JsiiProperty: "userPathInput", GoGetter: "UserPathInput"},
		},
		func() interface{} {
			j := jsiiProxy_AwsSecretBackendRole{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-vault.awsSecretBackendRole.AwsSecretBackendRoleConfig",
		reflect.TypeOf((*AwsSecretBackendRoleConfig)(nil)).Elem(),
	)
}
