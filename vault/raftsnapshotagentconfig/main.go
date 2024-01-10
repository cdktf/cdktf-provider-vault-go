// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package raftsnapshotagentconfig

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-vault.raftSnapshotAgentConfig.RaftSnapshotAgentConfig",
		reflect.TypeOf((*RaftSnapshotAgentConfig)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "awsAccessKeyId", GoGetter: "AwsAccessKeyId"},
			_jsii_.MemberProperty{JsiiProperty: "awsAccessKeyIdInput", GoGetter: "AwsAccessKeyIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "awsS3Bucket", GoGetter: "AwsS3Bucket"},
			_jsii_.MemberProperty{JsiiProperty: "awsS3BucketInput", GoGetter: "AwsS3BucketInput"},
			_jsii_.MemberProperty{JsiiProperty: "awsS3DisableTls", GoGetter: "AwsS3DisableTls"},
			_jsii_.MemberProperty{JsiiProperty: "awsS3DisableTlsInput", GoGetter: "AwsS3DisableTlsInput"},
			_jsii_.MemberProperty{JsiiProperty: "awsS3EnableKms", GoGetter: "AwsS3EnableKms"},
			_jsii_.MemberProperty{JsiiProperty: "awsS3EnableKmsInput", GoGetter: "AwsS3EnableKmsInput"},
			_jsii_.MemberProperty{JsiiProperty: "awsS3Endpoint", GoGetter: "AwsS3Endpoint"},
			_jsii_.MemberProperty{JsiiProperty: "awsS3EndpointInput", GoGetter: "AwsS3EndpointInput"},
			_jsii_.MemberProperty{JsiiProperty: "awsS3ForcePathStyle", GoGetter: "AwsS3ForcePathStyle"},
			_jsii_.MemberProperty{JsiiProperty: "awsS3ForcePathStyleInput", GoGetter: "AwsS3ForcePathStyleInput"},
			_jsii_.MemberProperty{JsiiProperty: "awsS3KmsKey", GoGetter: "AwsS3KmsKey"},
			_jsii_.MemberProperty{JsiiProperty: "awsS3KmsKeyInput", GoGetter: "AwsS3KmsKeyInput"},
			_jsii_.MemberProperty{JsiiProperty: "awsS3Region", GoGetter: "AwsS3Region"},
			_jsii_.MemberProperty{JsiiProperty: "awsS3RegionInput", GoGetter: "AwsS3RegionInput"},
			_jsii_.MemberProperty{JsiiProperty: "awsS3ServerSideEncryption", GoGetter: "AwsS3ServerSideEncryption"},
			_jsii_.MemberProperty{JsiiProperty: "awsS3ServerSideEncryptionInput", GoGetter: "AwsS3ServerSideEncryptionInput"},
			_jsii_.MemberProperty{JsiiProperty: "awsSecretAccessKey", GoGetter: "AwsSecretAccessKey"},
			_jsii_.MemberProperty{JsiiProperty: "awsSecretAccessKeyInput", GoGetter: "AwsSecretAccessKeyInput"},
			_jsii_.MemberProperty{JsiiProperty: "awsSessionToken", GoGetter: "AwsSessionToken"},
			_jsii_.MemberProperty{JsiiProperty: "awsSessionTokenInput", GoGetter: "AwsSessionTokenInput"},
			_jsii_.MemberProperty{JsiiProperty: "azureAccountKey", GoGetter: "AzureAccountKey"},
			_jsii_.MemberProperty{JsiiProperty: "azureAccountKeyInput", GoGetter: "AzureAccountKeyInput"},
			_jsii_.MemberProperty{JsiiProperty: "azureAccountName", GoGetter: "AzureAccountName"},
			_jsii_.MemberProperty{JsiiProperty: "azureAccountNameInput", GoGetter: "AzureAccountNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "azureBlobEnvironment", GoGetter: "AzureBlobEnvironment"},
			_jsii_.MemberProperty{JsiiProperty: "azureBlobEnvironmentInput", GoGetter: "AzureBlobEnvironmentInput"},
			_jsii_.MemberProperty{JsiiProperty: "azureContainerName", GoGetter: "AzureContainerName"},
			_jsii_.MemberProperty{JsiiProperty: "azureContainerNameInput", GoGetter: "AzureContainerNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "azureEndpoint", GoGetter: "AzureEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "azureEndpointInput", GoGetter: "AzureEndpointInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "filePrefix", GoGetter: "FilePrefix"},
			_jsii_.MemberProperty{JsiiProperty: "filePrefixInput", GoGetter: "FilePrefixInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "googleDisableTls", GoGetter: "GoogleDisableTls"},
			_jsii_.MemberProperty{JsiiProperty: "googleDisableTlsInput", GoGetter: "GoogleDisableTlsInput"},
			_jsii_.MemberProperty{JsiiProperty: "googleEndpoint", GoGetter: "GoogleEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "googleEndpointInput", GoGetter: "GoogleEndpointInput"},
			_jsii_.MemberProperty{JsiiProperty: "googleGcsBucket", GoGetter: "GoogleGcsBucket"},
			_jsii_.MemberProperty{JsiiProperty: "googleGcsBucketInput", GoGetter: "GoogleGcsBucketInput"},
			_jsii_.MemberProperty{JsiiProperty: "googleServiceAccountKey", GoGetter: "GoogleServiceAccountKey"},
			_jsii_.MemberProperty{JsiiProperty: "googleServiceAccountKeyInput", GoGetter: "GoogleServiceAccountKeyInput"},
			_jsii_.MemberMethod{JsiiMethod: "hasResourceMove", GoMethod: "HasResourceMove"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "intervalSeconds", GoGetter: "IntervalSeconds"},
			_jsii_.MemberProperty{JsiiProperty: "intervalSecondsInput", GoGetter: "IntervalSecondsInput"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "localMaxSpace", GoGetter: "LocalMaxSpace"},
			_jsii_.MemberProperty{JsiiProperty: "localMaxSpaceInput", GoGetter: "LocalMaxSpaceInput"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "namespace", GoGetter: "Namespace"},
			_jsii_.MemberProperty{JsiiProperty: "namespaceInput", GoGetter: "NamespaceInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "pathPrefix", GoGetter: "PathPrefix"},
			_jsii_.MemberProperty{JsiiProperty: "pathPrefixInput", GoGetter: "PathPrefixInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetAwsAccessKeyId", GoMethod: "ResetAwsAccessKeyId"},
			_jsii_.MemberMethod{JsiiMethod: "resetAwsS3Bucket", GoMethod: "ResetAwsS3Bucket"},
			_jsii_.MemberMethod{JsiiMethod: "resetAwsS3DisableTls", GoMethod: "ResetAwsS3DisableTls"},
			_jsii_.MemberMethod{JsiiMethod: "resetAwsS3EnableKms", GoMethod: "ResetAwsS3EnableKms"},
			_jsii_.MemberMethod{JsiiMethod: "resetAwsS3Endpoint", GoMethod: "ResetAwsS3Endpoint"},
			_jsii_.MemberMethod{JsiiMethod: "resetAwsS3ForcePathStyle", GoMethod: "ResetAwsS3ForcePathStyle"},
			_jsii_.MemberMethod{JsiiMethod: "resetAwsS3KmsKey", GoMethod: "ResetAwsS3KmsKey"},
			_jsii_.MemberMethod{JsiiMethod: "resetAwsS3Region", GoMethod: "ResetAwsS3Region"},
			_jsii_.MemberMethod{JsiiMethod: "resetAwsS3ServerSideEncryption", GoMethod: "ResetAwsS3ServerSideEncryption"},
			_jsii_.MemberMethod{JsiiMethod: "resetAwsSecretAccessKey", GoMethod: "ResetAwsSecretAccessKey"},
			_jsii_.MemberMethod{JsiiMethod: "resetAwsSessionToken", GoMethod: "ResetAwsSessionToken"},
			_jsii_.MemberMethod{JsiiMethod: "resetAzureAccountKey", GoMethod: "ResetAzureAccountKey"},
			_jsii_.MemberMethod{JsiiMethod: "resetAzureAccountName", GoMethod: "ResetAzureAccountName"},
			_jsii_.MemberMethod{JsiiMethod: "resetAzureBlobEnvironment", GoMethod: "ResetAzureBlobEnvironment"},
			_jsii_.MemberMethod{JsiiMethod: "resetAzureContainerName", GoMethod: "ResetAzureContainerName"},
			_jsii_.MemberMethod{JsiiMethod: "resetAzureEndpoint", GoMethod: "ResetAzureEndpoint"},
			_jsii_.MemberMethod{JsiiMethod: "resetFilePrefix", GoMethod: "ResetFilePrefix"},
			_jsii_.MemberMethod{JsiiMethod: "resetGoogleDisableTls", GoMethod: "ResetGoogleDisableTls"},
			_jsii_.MemberMethod{JsiiMethod: "resetGoogleEndpoint", GoMethod: "ResetGoogleEndpoint"},
			_jsii_.MemberMethod{JsiiMethod: "resetGoogleGcsBucket", GoMethod: "ResetGoogleGcsBucket"},
			_jsii_.MemberMethod{JsiiMethod: "resetGoogleServiceAccountKey", GoMethod: "ResetGoogleServiceAccountKey"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetLocalMaxSpace", GoMethod: "ResetLocalMaxSpace"},
			_jsii_.MemberMethod{JsiiMethod: "resetNamespace", GoMethod: "ResetNamespace"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetRetain", GoMethod: "ResetRetain"},
			_jsii_.MemberProperty{JsiiProperty: "retain", GoGetter: "Retain"},
			_jsii_.MemberProperty{JsiiProperty: "retainInput", GoGetter: "RetainInput"},
			_jsii_.MemberProperty{JsiiProperty: "storageType", GoGetter: "StorageType"},
			_jsii_.MemberProperty{JsiiProperty: "storageTypeInput", GoGetter: "StorageTypeInput"},
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
			j := jsiiProxy_RaftSnapshotAgentConfig{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-vault.raftSnapshotAgentConfig.RaftSnapshotAgentConfigConfig",
		reflect.TypeOf((*RaftSnapshotAgentConfigConfig)(nil)).Elem(),
	)
}
