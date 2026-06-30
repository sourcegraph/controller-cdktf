package observegcp

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/observegcp/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/observegcp/internal"
)

// Defines an Observegcp based on a Terraform module.
//
// Docs at Terraform Registry: {@link https://registry.terraform.io/modules/observeinc/collection/google/1.0.2 observeinc/collection/google}
type Observegcp interface {
	cdktf.TerraformModule
	BucketLifecycleAbortUploadDays() *float64
	SetBucketLifecycleAbortUploadDays(val *float64)
	BucketLifecycleDeleteDays() *float64
	SetBucketLifecycleDeleteDays(val *float64)
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CloudFunctionDebugLevel() *string
	SetCloudFunctionDebugLevel(val *string)
	// Experimental.
	ConstructNodeMetadata() *map[string]any
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EnableAssetTracking() *bool
	SetEnableAssetTracking(val *bool)
	EnableFunction() *bool
	SetEnableFunction(val *bool)
	FolderIncludeChildren() *bool
	SetFolderIncludeChildren(val *bool)
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	FunctionAvailableMemoryMb() *float64
	SetFunctionAvailableMemoryMb(val *float64)
	FunctionBucket() *string
	SetFunctionBucket(val *string)
	FunctionDisableLogging() *bool
	SetFunctionDisableLogging(val *bool)
	FunctionMaxInstances() *float64
	SetFunctionMaxInstances(val *float64)
	FunctionObject() *string
	SetFunctionObject(val *string)
	FunctionRoles() *[]*string
	SetFunctionRoles(val *[]*string)
	FunctionScheduleFrequency() *string
	SetFunctionScheduleFrequency(val *string)
	FunctionScheduleFrequencyRestOfAssets() *string
	SetFunctionScheduleFrequencyRestOfAssets(val *string)
	FunctionTimeout() *float64
	SetFunctionTimeout(val *float64)
	GcpRegion() *string
	SetGcpRegion(val *string)
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LoggingExclusions() any
	SetLoggingExclusions(val any)
	LoggingFilter() *string
	SetLoggingFilter(val *string)
	MaxAttempts() *float64
	SetMaxAttempts(val *float64)
	MaxConcurrentDispatches() *float64
	SetMaxConcurrentDispatches(val *float64)
	MaxDispatchesPerSecond() *float64
	SetMaxDispatchesPerSecond(val *float64)
	MaxRetryDuration() *string
	SetMaxRetryDuration(val *string)
	MinBackoff() *string
	SetMinBackoff(val *string)
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	PollerRoles() *[]*string
	SetPollerRoles(val *[]*string)
	ProjectId() *string
	SetProjectId(val *string)
	ProjectOutput() *string
	// Experimental.
	Providers() *[]any
	PubsubAckDeadlineSeconds() *float64
	SetPubsubAckDeadlineSeconds(val *float64)
	PubsubMaximumBackoff() *string
	SetPubsubMaximumBackoff(val *string)
	PubsubMessageRetentionDuration() *string
	SetPubsubMessageRetentionDuration(val *string)
	PubsubMinimumBackoff() *string
	SetPubsubMinimumBackoff(val *string)
	// Experimental.
	RawOverrides() any
	Resource() *string
	SetResource(val *string)
	ServiceAccountKeyOutput() *string
	// Experimental.
	SkipAssetCreationFromLocalModules() *bool
	// Experimental.
	Source() *string
	SubscriptionOutput() *string
	TopicOutput() *string
	// Experimental.
	Version() *string
	// Experimental.
	AddOverride(path *string, value any)
	// Experimental.
	AddProvider(provider any)
	// Experimental.
	GetString(output *string) *string
	// Experimental.
	InterpolationForOutput(moduleOutput *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]any
	SynthesizeHclAttributes() *map[string]any
	// Experimental.
	ToHclTerraform() any
	// Experimental.
	ToMetadata() any
	// Returns a string representation of this construct.
	ToString() *string
	// Experimental.
	ToTerraform() any
}

// The jsii proxy struct for Observegcp
type jsiiProxy_Observegcp struct {
	internal.Type__cdktfTerraformModule
}

func (j *jsiiProxy_Observegcp) BucketLifecycleAbortUploadDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bucketLifecycleAbortUploadDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Observegcp) BucketLifecycleDeleteDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bucketLifecycleDeleteDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Observegcp) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Observegcp) CloudFunctionDebugLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudFunctionDebugLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Observegcp) ConstructNodeMetadata() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Observegcp) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Observegcp) EnableAssetTracking() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enableAssetTracking",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Observegcp) EnableFunction() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enableFunction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Observegcp) FolderIncludeChildren() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"folderIncludeChildren",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Observegcp) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Observegcp) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Observegcp) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Observegcp) FunctionAvailableMemoryMb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"functionAvailableMemoryMb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Observegcp) FunctionBucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionBucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Observegcp) FunctionDisableLogging() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"functionDisableLogging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Observegcp) FunctionMaxInstances() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"functionMaxInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Observegcp) FunctionObject() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionObject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Observegcp) FunctionRoles() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"functionRoles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Observegcp) FunctionScheduleFrequency() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionScheduleFrequency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Observegcp) FunctionScheduleFrequencyRestOfAssets() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionScheduleFrequencyRestOfAssets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Observegcp) FunctionTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"functionTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Observegcp) GcpRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gcpRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Observegcp) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Observegcp) LoggingExclusions() any {
	var returns any
	_jsii_.Get(
		j,
		"loggingExclusions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Observegcp) LoggingFilter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loggingFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Observegcp) MaxAttempts() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAttempts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Observegcp) MaxConcurrentDispatches() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrentDispatches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Observegcp) MaxDispatchesPerSecond() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxDispatchesPerSecond",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Observegcp) MaxRetryDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxRetryDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Observegcp) MinBackoff() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minBackoff",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Observegcp) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Observegcp) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Observegcp) PollerRoles() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"pollerRoles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Observegcp) ProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Observegcp) ProjectOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Observegcp) Providers() *[]any {
	var returns *[]any
	_jsii_.Get(
		j,
		"providers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Observegcp) PubsubAckDeadlineSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"pubsubAckDeadlineSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Observegcp) PubsubMaximumBackoff() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pubsubMaximumBackoff",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Observegcp) PubsubMessageRetentionDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pubsubMessageRetentionDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Observegcp) PubsubMinimumBackoff() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pubsubMinimumBackoff",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Observegcp) RawOverrides() any {
	var returns any
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Observegcp) Resource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Observegcp) ServiceAccountKeyOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountKeyOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Observegcp) SkipAssetCreationFromLocalModules() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"skipAssetCreationFromLocalModules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Observegcp) Source() *string {
	var returns *string
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Observegcp) SubscriptionOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subscriptionOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Observegcp) TopicOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"topicOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Observegcp) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func NewObservegcp(scope constructs.Construct, id *string, config *ObservegcpConfig) Observegcp {
	_init_.Initialize()

	if err := validateNewObservegcpParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Observegcp{}

	_jsii_.Create(
		"@cdktf/provider-observegcp.Observegcp",
		[]any{scope, id, config},
		&j,
	)

	return &j
}

func NewObservegcp_Override(o Observegcp, scope constructs.Construct, id *string, config *ObservegcpConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-observegcp.Observegcp",
		[]any{scope, id, config},
		o,
	)
}

func (j *jsiiProxy_Observegcp) SetBucketLifecycleAbortUploadDays(val *float64) {
	_jsii_.Set(
		j,
		"bucketLifecycleAbortUploadDays",
		val,
	)
}

func (j *jsiiProxy_Observegcp) SetBucketLifecycleDeleteDays(val *float64) {
	_jsii_.Set(
		j,
		"bucketLifecycleDeleteDays",
		val,
	)
}

func (j *jsiiProxy_Observegcp) SetCloudFunctionDebugLevel(val *string) {
	_jsii_.Set(
		j,
		"cloudFunctionDebugLevel",
		val,
	)
}

func (j *jsiiProxy_Observegcp) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Observegcp) SetEnableAssetTracking(val *bool) {
	_jsii_.Set(
		j,
		"enableAssetTracking",
		val,
	)
}

func (j *jsiiProxy_Observegcp) SetEnableFunction(val *bool) {
	_jsii_.Set(
		j,
		"enableFunction",
		val,
	)
}

func (j *jsiiProxy_Observegcp) SetFolderIncludeChildren(val *bool) {
	_jsii_.Set(
		j,
		"folderIncludeChildren",
		val,
	)
}

func (j *jsiiProxy_Observegcp) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Observegcp) SetFunctionAvailableMemoryMb(val *float64) {
	_jsii_.Set(
		j,
		"functionAvailableMemoryMb",
		val,
	)
}

func (j *jsiiProxy_Observegcp) SetFunctionBucket(val *string) {
	_jsii_.Set(
		j,
		"functionBucket",
		val,
	)
}

func (j *jsiiProxy_Observegcp) SetFunctionDisableLogging(val *bool) {
	_jsii_.Set(
		j,
		"functionDisableLogging",
		val,
	)
}

func (j *jsiiProxy_Observegcp) SetFunctionMaxInstances(val *float64) {
	_jsii_.Set(
		j,
		"functionMaxInstances",
		val,
	)
}

func (j *jsiiProxy_Observegcp) SetFunctionObject(val *string) {
	_jsii_.Set(
		j,
		"functionObject",
		val,
	)
}

func (j *jsiiProxy_Observegcp) SetFunctionRoles(val *[]*string) {
	_jsii_.Set(
		j,
		"functionRoles",
		val,
	)
}

func (j *jsiiProxy_Observegcp) SetFunctionScheduleFrequency(val *string) {
	_jsii_.Set(
		j,
		"functionScheduleFrequency",
		val,
	)
}

func (j *jsiiProxy_Observegcp) SetFunctionScheduleFrequencyRestOfAssets(val *string) {
	_jsii_.Set(
		j,
		"functionScheduleFrequencyRestOfAssets",
		val,
	)
}

func (j *jsiiProxy_Observegcp) SetFunctionTimeout(val *float64) {
	_jsii_.Set(
		j,
		"functionTimeout",
		val,
	)
}

func (j *jsiiProxy_Observegcp) SetGcpRegion(val *string) {
	_jsii_.Set(
		j,
		"gcpRegion",
		val,
	)
}

func (j *jsiiProxy_Observegcp) SetLabels(val *map[string]*string) {
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_Observegcp) SetLoggingExclusions(val any) {
	if err := j.validateSetLoggingExclusionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loggingExclusions",
		val,
	)
}

func (j *jsiiProxy_Observegcp) SetLoggingFilter(val *string) {
	_jsii_.Set(
		j,
		"loggingFilter",
		val,
	)
}

func (j *jsiiProxy_Observegcp) SetMaxAttempts(val *float64) {
	_jsii_.Set(
		j,
		"maxAttempts",
		val,
	)
}

func (j *jsiiProxy_Observegcp) SetMaxConcurrentDispatches(val *float64) {
	_jsii_.Set(
		j,
		"maxConcurrentDispatches",
		val,
	)
}

func (j *jsiiProxy_Observegcp) SetMaxDispatchesPerSecond(val *float64) {
	_jsii_.Set(
		j,
		"maxDispatchesPerSecond",
		val,
	)
}

func (j *jsiiProxy_Observegcp) SetMaxRetryDuration(val *string) {
	_jsii_.Set(
		j,
		"maxRetryDuration",
		val,
	)
}

func (j *jsiiProxy_Observegcp) SetMinBackoff(val *string) {
	_jsii_.Set(
		j,
		"minBackoff",
		val,
	)
}

func (j *jsiiProxy_Observegcp) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_Observegcp) SetPollerRoles(val *[]*string) {
	_jsii_.Set(
		j,
		"pollerRoles",
		val,
	)
}

func (j *jsiiProxy_Observegcp) SetProjectId(val *string) {
	_jsii_.Set(
		j,
		"projectId",
		val,
	)
}

func (j *jsiiProxy_Observegcp) SetPubsubAckDeadlineSeconds(val *float64) {
	_jsii_.Set(
		j,
		"pubsubAckDeadlineSeconds",
		val,
	)
}

func (j *jsiiProxy_Observegcp) SetPubsubMaximumBackoff(val *string) {
	_jsii_.Set(
		j,
		"pubsubMaximumBackoff",
		val,
	)
}

func (j *jsiiProxy_Observegcp) SetPubsubMessageRetentionDuration(val *string) {
	_jsii_.Set(
		j,
		"pubsubMessageRetentionDuration",
		val,
	)
}

func (j *jsiiProxy_Observegcp) SetPubsubMinimumBackoff(val *string) {
	_jsii_.Set(
		j,
		"pubsubMinimumBackoff",
		val,
	)
}

func (j *jsiiProxy_Observegcp) SetResource(val *string) {
	if err := j.validateSetResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resource",
		val,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func Observegcp_IsConstruct(x any) *bool {
	_init_.Initialize()

	if err := validateObservegcp_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-observegcp.Observegcp",
		"isConstruct",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func Observegcp_IsTerraformElement(x any) *bool {
	_init_.Initialize()

	if err := validateObservegcp_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-observegcp.Observegcp",
		"isTerraformElement",
		[]any{x},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_Observegcp) AddOverride(path *string, value any) {
	if err := o.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"addOverride",
		[]any{path, value},
	)
}

func (o *jsiiProxy_Observegcp) AddProvider(provider any) {
	if err := o.validateAddProviderParameters(provider); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"addProvider",
		[]any{provider},
	)
}

func (o *jsiiProxy_Observegcp) GetString(output *string) *string {
	if err := o.validateGetStringParameters(output); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		o,
		"getString",
		[]any{output},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_Observegcp) InterpolationForOutput(moduleOutput *string) cdktf.IResolvable {
	if err := o.validateInterpolationForOutputParameters(moduleOutput); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForOutput",
		[]any{moduleOutput},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_Observegcp) OverrideLogicalId(newLogicalId *string) {
	if err := o.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"overrideLogicalId",
		[]any{newLogicalId},
	)
}

func (o *jsiiProxy_Observegcp) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		o,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_Observegcp) SynthesizeAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		o,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_Observegcp) SynthesizeHclAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		o,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_Observegcp) ToHclTerraform() any {
	var returns any

	_jsii_.Invoke(
		o,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_Observegcp) ToMetadata() any {
	var returns any

	_jsii_.Invoke(
		o,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_Observegcp) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_Observegcp) ToTerraform() any {
	var returns any

	_jsii_.Invoke(
		o,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}
