// @cdktf/provider-observegcp
package observegcp

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-observegcp.Observegcp",
		reflect.TypeOf((*Observegcp)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addProvider", GoMethod: "AddProvider"},
			_jsii_.MemberProperty{JsiiProperty: "bucketLifecycleAbortUploadDays", GoGetter: "BucketLifecycleAbortUploadDays"},
			_jsii_.MemberProperty{JsiiProperty: "bucketLifecycleDeleteDays", GoGetter: "BucketLifecycleDeleteDays"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "cloudFunctionDebugLevel", GoGetter: "CloudFunctionDebugLevel"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "enableAssetTracking", GoGetter: "EnableAssetTracking"},
			_jsii_.MemberProperty{JsiiProperty: "enableFunction", GoGetter: "EnableFunction"},
			_jsii_.MemberProperty{JsiiProperty: "folderIncludeChildren", GoGetter: "FolderIncludeChildren"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberProperty{JsiiProperty: "functionAvailableMemoryMb", GoGetter: "FunctionAvailableMemoryMb"},
			_jsii_.MemberProperty{JsiiProperty: "functionBucket", GoGetter: "FunctionBucket"},
			_jsii_.MemberProperty{JsiiProperty: "functionDisableLogging", GoGetter: "FunctionDisableLogging"},
			_jsii_.MemberProperty{JsiiProperty: "functionMaxInstances", GoGetter: "FunctionMaxInstances"},
			_jsii_.MemberProperty{JsiiProperty: "functionObject", GoGetter: "FunctionObject"},
			_jsii_.MemberProperty{JsiiProperty: "functionRoles", GoGetter: "FunctionRoles"},
			_jsii_.MemberProperty{JsiiProperty: "functionScheduleFrequency", GoGetter: "FunctionScheduleFrequency"},
			_jsii_.MemberProperty{JsiiProperty: "functionScheduleFrequencyRestOfAssets", GoGetter: "FunctionScheduleFrequencyRestOfAssets"},
			_jsii_.MemberProperty{JsiiProperty: "functionTimeout", GoGetter: "FunctionTimeout"},
			_jsii_.MemberProperty{JsiiProperty: "gcpRegion", GoGetter: "GcpRegion"},
			_jsii_.MemberMethod{JsiiMethod: "getString", GoMethod: "GetString"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForOutput", GoMethod: "InterpolationForOutput"},
			_jsii_.MemberProperty{JsiiProperty: "labels", GoGetter: "Labels"},
			_jsii_.MemberProperty{JsiiProperty: "loggingExclusions", GoGetter: "LoggingExclusions"},
			_jsii_.MemberProperty{JsiiProperty: "loggingFilter", GoGetter: "LoggingFilter"},
			_jsii_.MemberProperty{JsiiProperty: "maxAttempts", GoGetter: "MaxAttempts"},
			_jsii_.MemberProperty{JsiiProperty: "maxConcurrentDispatches", GoGetter: "MaxConcurrentDispatches"},
			_jsii_.MemberProperty{JsiiProperty: "maxDispatchesPerSecond", GoGetter: "MaxDispatchesPerSecond"},
			_jsii_.MemberProperty{JsiiProperty: "maxRetryDuration", GoGetter: "MaxRetryDuration"},
			_jsii_.MemberProperty{JsiiProperty: "minBackoff", GoGetter: "MinBackoff"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "pollerRoles", GoGetter: "PollerRoles"},
			_jsii_.MemberProperty{JsiiProperty: "projectId", GoGetter: "ProjectId"},
			_jsii_.MemberProperty{JsiiProperty: "projectOutput", GoGetter: "ProjectOutput"},
			_jsii_.MemberProperty{JsiiProperty: "providers", GoGetter: "Providers"},
			_jsii_.MemberProperty{JsiiProperty: "pubsubAckDeadlineSeconds", GoGetter: "PubsubAckDeadlineSeconds"},
			_jsii_.MemberProperty{JsiiProperty: "pubsubMaximumBackoff", GoGetter: "PubsubMaximumBackoff"},
			_jsii_.MemberProperty{JsiiProperty: "pubsubMessageRetentionDuration", GoGetter: "PubsubMessageRetentionDuration"},
			_jsii_.MemberProperty{JsiiProperty: "pubsubMinimumBackoff", GoGetter: "PubsubMinimumBackoff"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "resource", GoGetter: "Resource"},
			_jsii_.MemberProperty{JsiiProperty: "serviceAccountKeyOutput", GoGetter: "ServiceAccountKeyOutput"},
			_jsii_.MemberProperty{JsiiProperty: "skipAssetCreationFromLocalModules", GoGetter: "SkipAssetCreationFromLocalModules"},
			_jsii_.MemberProperty{JsiiProperty: "source", GoGetter: "Source"},
			_jsii_.MemberProperty{JsiiProperty: "subscriptionOutput", GoGetter: "SubscriptionOutput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "topicOutput", GoGetter: "TopicOutput"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "version", GoGetter: "Version"},
		},
		func() interface{} {
			j := jsiiProxy_Observegcp{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformModule)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-observegcp.ObservegcpConfig",
		reflect.TypeOf((*ObservegcpConfig)(nil)).Elem(),
	)
}
