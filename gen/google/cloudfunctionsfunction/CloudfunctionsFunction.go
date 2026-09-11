package cloudfunctionsfunction

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google/cloudfunctionsfunction/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/cloudfunctions_function google_cloudfunctions_function}.
type CloudfunctionsFunction interface {
	cdktn.TerraformResource
	AutomaticUpdatePolicy() CloudfunctionsFunctionAutomaticUpdatePolicyOutputReference
	AutomaticUpdatePolicyInput() *CloudfunctionsFunctionAutomaticUpdatePolicy
	AvailableMemoryMb() *float64
	SetAvailableMemoryMb(val *float64)
	AvailableMemoryMbInput() *float64
	BuildEnvironmentVariables() *map[string]*string
	SetBuildEnvironmentVariables(val *map[string]*string)
	BuildEnvironmentVariablesInput() *map[string]*string
	BuildServiceAccount() *string
	SetBuildServiceAccount(val *string)
	BuildServiceAccountInput() *string
	BuildWorkerPool() *string
	SetBuildWorkerPool(val *string)
	BuildWorkerPoolInput() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DockerRegistry() *string
	SetDockerRegistry(val *string)
	DockerRegistryInput() *string
	DockerRepository() *string
	SetDockerRepository(val *string)
	DockerRepositoryInput() *string
	EffectiveLabels() cdktn.StringMap
	EntryPoint() *string
	SetEntryPoint(val *string)
	EntryPointInput() *string
	EnvironmentVariables() *map[string]*string
	SetEnvironmentVariables(val *map[string]*string)
	EnvironmentVariablesInput() *map[string]*string
	EventTrigger() CloudfunctionsFunctionEventTriggerOutputReference
	EventTriggerInput() *CloudfunctionsFunctionEventTrigger
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HttpsTriggerSecurityLevel() *string
	SetHttpsTriggerSecurityLevel(val *string)
	HttpsTriggerSecurityLevelInput() *string
	HttpsTriggerUrl() *string
	SetHttpsTriggerUrl(val *string)
	HttpsTriggerUrlInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	IngressSettings() *string
	SetIngressSettings(val *string)
	IngressSettingsInput() *string
	KmsKeyName() *string
	SetKmsKeyName(val *string)
	KmsKeyNameInput() *string
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	MaxInstances() *float64
	SetMaxInstances(val *float64)
	MaxInstancesInput() *float64
	MinInstances() *float64
	SetMinInstances(val *float64)
	MinInstancesInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	OnDeployUpdatePolicy() CloudfunctionsFunctionOnDeployUpdatePolicyOutputReference
	OnDeployUpdatePolicyInput() *CloudfunctionsFunctionOnDeployUpdatePolicy
	Project() *string
	SetProject(val *string)
	ProjectInput() *string
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	Runtime() *string
	SetRuntime(val *string)
	RuntimeInput() *string
	SecretEnvironmentVariables() CloudfunctionsFunctionSecretEnvironmentVariablesList
	SecretEnvironmentVariablesInput() interface{}
	SecretVolumes() CloudfunctionsFunctionSecretVolumesList
	SecretVolumesInput() interface{}
	ServiceAccountEmail() *string
	SetServiceAccountEmail(val *string)
	ServiceAccountEmailInput() *string
	SourceArchiveBucket() *string
	SetSourceArchiveBucket(val *string)
	SourceArchiveBucketInput() *string
	SourceArchiveObject() *string
	SetSourceArchiveObject(val *string)
	SourceArchiveObjectInput() *string
	SourceRepository() CloudfunctionsFunctionSourceRepositoryOutputReference
	SourceRepositoryInput() *CloudfunctionsFunctionSourceRepository
	Status() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	TerraformLabels() cdktn.StringMap
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeout() *float64
	SetTimeout(val *float64)
	TimeoutInput() *float64
	Timeouts() CloudfunctionsFunctionTimeoutsOutputReference
	TimeoutsInput() interface{}
	TriggerHttp() interface{}
	SetTriggerHttp(val interface{})
	TriggerHttpInput() interface{}
	VersionId() *string
	VpcConnector() *string
	SetVpcConnector(val *string)
	VpcConnectorEgressSettings() *string
	SetVpcConnectorEgressSettings(val *string)
	VpcConnectorEgressSettingsInput() *string
	VpcConnectorInput() *string
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktn.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	// Wraps a write-only attribute's already-mapped value so that `ProviderFeature.WRITE_ONLY_ATTRIBUTES` usage is registered at *resolve* time instead of at mutation time (setter/constructor). Called by generated bindings from `synthesizeAttributes()` and `synthesizeHclAttributes()`, e.g. `secret_key_wo: this.markWriteOnlyAttribute(cdktn.stringToTerraform(this._secretKeyWo))`; not intended to be called directly.
	//
	// `undefined` passes through completely unchanged, so the existing
	// undefined-filtering that omits unset attributes from synthesized
	// output (see `resolve()` in `tokens/private/resolve.ts`, and the
	// `value.value !== undefined` filter in generated
	// `synthesizeHclAttributes()`) keeps working untouched. `null` is also
	// passed through unchanged: it already renders as an explicit
	// null-out and must not arm the validation either.
	//
	// Any other value - including one that will itself resolve to nothing
	// (e.g. a `Lazy`/`IResolvable` producer with no value to contribute) -
	// is wrapped in a token whose `resolve()` defers to the real resolver
	// first and registers usage only if what comes back is not
	// `null`/`undefined`; the resolved value is then returned unchanged,
	// so what actually renders is untouched by this wrapper. A producer
	// that resolves to `undefined` therefore neither registers usage nor
	// leaves anything behind in the synthesized attribute - the omission
	// behaves exactly as if the attribute had never been set.
	//
	// Registration goes through `_registerResolveDiscoveredProviderFeatureUsage`
	// rather than `registerProviderFeatureUsage`: usage here is only known at
	// resolve time, and a given element can be resolved across many
	// synthesis passes over its lifetime (repeated `app.synth()` calls,
	// tests reusing a construct tree), so it must represent only the CURRENT
	// pass rather than accumulate forever. Every validation-enabled entry
	// point (`App.synth`; `Testing.synth`/`synthHcl` with validations;
	// `StackSynthesizer.synthesize`) runs a prepare step that deactivates any
	// stale registration and then resolves every element's `toTerraform()`
	// before that same entry point's validations run - see
	// `TerraformStack._runPreparingResolve` - so whatever this closure
	// (re-)registers during that prepare step is always visible to the
	// validation that reads it afterwards, and nothing left over from an
	// earlier pass leaks into the current one.
	// Experimental.
	MarkWriteOnlyAttribute(value interface{}) interface{}
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using its instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutAutomaticUpdatePolicy(value *CloudfunctionsFunctionAutomaticUpdatePolicy)
	PutEventTrigger(value *CloudfunctionsFunctionEventTrigger)
	PutOnDeployUpdatePolicy(value *CloudfunctionsFunctionOnDeployUpdatePolicy)
	PutSecretEnvironmentVariables(value interface{})
	PutSecretVolumes(value interface{})
	PutSourceRepository(value *CloudfunctionsFunctionSourceRepository)
	PutTimeouts(value *CloudfunctionsFunctionTimeouts)
	// Registers a synth-time validation that the project's declared targetVersions admit the given provider-protocol feature family.
	//
	// Called by generated provider bindings when a versioned feature is
	// structurally in use - the element's existence in the construct tree
	// already implies the feature is used, e.g. constructing a
	// `TerraformEphemeralResource` at all - so, unlike
	// `_registerResolveDiscoveredProviderFeatureUsage`, this registration is
	// never deactivated by `_resetResolveDiscoveredProviderFeatureUsage`. Not
	// intended to be called directly by user code. Lives on `TerraformElement`
	// (rather than `TerraformResource`) so it covers any element subclass
	// that needs it.
	// Experimental.
	RegisterProviderFeatureUsage(feature cdktn.ProviderFeature)
	ResetAutomaticUpdatePolicy()
	ResetAvailableMemoryMb()
	ResetBuildEnvironmentVariables()
	ResetBuildServiceAccount()
	ResetBuildWorkerPool()
	ResetDescription()
	ResetDockerRegistry()
	ResetDockerRepository()
	ResetEntryPoint()
	ResetEnvironmentVariables()
	ResetEventTrigger()
	ResetHttpsTriggerSecurityLevel()
	ResetHttpsTriggerUrl()
	ResetId()
	ResetIngressSettings()
	ResetKmsKeyName()
	ResetLabels()
	ResetMaxInstances()
	ResetMinInstances()
	ResetOnDeployUpdatePolicy()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetRegion()
	ResetSecretEnvironmentVariables()
	ResetSecretVolumes()
	ResetServiceAccountEmail()
	ResetSourceArchiveBucket()
	ResetSourceArchiveObject()
	ResetSourceRepository()
	ResetTimeout()
	ResetTimeouts()
	ResetTriggerHttp()
	ResetVpcConnector()
	ResetVpcConnectorEgressSettings()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
	// Applies one or more mixins to this construct.
	//
	// Mixins are applied in order. The list of constructs is captured at the
	// start of the call, so constructs added by a mixin will not be visited.
	// Use multiple `with()` calls if subsequent mixins should apply to added
	// constructs.
	//
	// Returns: This construct for chaining.
	With(mixins ...constructs.IMixin) constructs.IConstruct
}

// The jsii proxy struct for CloudfunctionsFunction
type jsiiProxy_CloudfunctionsFunction struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_CloudfunctionsFunction) AutomaticUpdatePolicy() CloudfunctionsFunctionAutomaticUpdatePolicyOutputReference {
	var returns CloudfunctionsFunctionAutomaticUpdatePolicyOutputReference
	_jsii_.Get(
		j,
		"automaticUpdatePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) AutomaticUpdatePolicyInput() *CloudfunctionsFunctionAutomaticUpdatePolicy {
	var returns *CloudfunctionsFunctionAutomaticUpdatePolicy
	_jsii_.Get(
		j,
		"automaticUpdatePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) AvailableMemoryMb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"availableMemoryMb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) AvailableMemoryMbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"availableMemoryMbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) BuildEnvironmentVariables() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"buildEnvironmentVariables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) BuildEnvironmentVariablesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"buildEnvironmentVariablesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) BuildServiceAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildServiceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) BuildServiceAccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildServiceAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) BuildWorkerPool() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildWorkerPool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) BuildWorkerPoolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildWorkerPoolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) DockerRegistry() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dockerRegistry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) DockerRegistryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dockerRegistryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) DockerRepository() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dockerRepository",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) DockerRepositoryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dockerRepositoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) EffectiveLabels() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"effectiveLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) EntryPoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entryPoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) EntryPointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entryPointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) EnvironmentVariables() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"environmentVariables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) EnvironmentVariablesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"environmentVariablesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) EventTrigger() CloudfunctionsFunctionEventTriggerOutputReference {
	var returns CloudfunctionsFunctionEventTriggerOutputReference
	_jsii_.Get(
		j,
		"eventTrigger",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) EventTriggerInput() *CloudfunctionsFunctionEventTrigger {
	var returns *CloudfunctionsFunctionEventTrigger
	_jsii_.Get(
		j,
		"eventTriggerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) HttpsTriggerSecurityLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpsTriggerSecurityLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) HttpsTriggerSecurityLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpsTriggerSecurityLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) HttpsTriggerUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpsTriggerUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) HttpsTriggerUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpsTriggerUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) IngressSettings() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ingressSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) IngressSettingsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ingressSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) KmsKeyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) KmsKeyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) MaxInstances() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) MaxInstancesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxInstancesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) MinInstances() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) MinInstancesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minInstancesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) OnDeployUpdatePolicy() CloudfunctionsFunctionOnDeployUpdatePolicyOutputReference {
	var returns CloudfunctionsFunctionOnDeployUpdatePolicyOutputReference
	_jsii_.Get(
		j,
		"onDeployUpdatePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) OnDeployUpdatePolicyInput() *CloudfunctionsFunctionOnDeployUpdatePolicy {
	var returns *CloudfunctionsFunctionOnDeployUpdatePolicy
	_jsii_.Get(
		j,
		"onDeployUpdatePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) Runtime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) RuntimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) SecretEnvironmentVariables() CloudfunctionsFunctionSecretEnvironmentVariablesList {
	var returns CloudfunctionsFunctionSecretEnvironmentVariablesList
	_jsii_.Get(
		j,
		"secretEnvironmentVariables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) SecretEnvironmentVariablesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secretEnvironmentVariablesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) SecretVolumes() CloudfunctionsFunctionSecretVolumesList {
	var returns CloudfunctionsFunctionSecretVolumesList
	_jsii_.Get(
		j,
		"secretVolumes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) SecretVolumesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secretVolumesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) ServiceAccountEmail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) ServiceAccountEmailInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountEmailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) SourceArchiveBucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceArchiveBucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) SourceArchiveBucketInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceArchiveBucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) SourceArchiveObject() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceArchiveObject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) SourceArchiveObjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceArchiveObjectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) SourceRepository() CloudfunctionsFunctionSourceRepositoryOutputReference {
	var returns CloudfunctionsFunctionSourceRepositoryOutputReference
	_jsii_.Get(
		j,
		"sourceRepository",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) SourceRepositoryInput() *CloudfunctionsFunctionSourceRepository {
	var returns *CloudfunctionsFunctionSourceRepository
	_jsii_.Get(
		j,
		"sourceRepositoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) TerraformLabels() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"terraformLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) Timeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) TimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) Timeouts() CloudfunctionsFunctionTimeoutsOutputReference {
	var returns CloudfunctionsFunctionTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) TriggerHttp() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"triggerHttp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) TriggerHttpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"triggerHttpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) VersionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) VpcConnector() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcConnector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) VpcConnectorEgressSettings() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcConnectorEgressSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) VpcConnectorEgressSettingsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcConnectorEgressSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) VpcConnectorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcConnectorInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/cloudfunctions_function google_cloudfunctions_function} Resource.
func NewCloudfunctionsFunction(scope constructs.Construct, id *string, config *CloudfunctionsFunctionConfig) CloudfunctionsFunction {
	_init_.Initialize()

	if err := validateNewCloudfunctionsFunctionParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudfunctionsFunction{}

	_jsii_.Create(
		"@cdktn/provider-google.cloudfunctionsFunction.CloudfunctionsFunction",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/cloudfunctions_function google_cloudfunctions_function} Resource.
func NewCloudfunctionsFunction_Override(c CloudfunctionsFunction, scope constructs.Construct, id *string, config *CloudfunctionsFunctionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.cloudfunctionsFunction.CloudfunctionsFunction",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CloudfunctionsFunction)SetAvailableMemoryMb(val *float64) {
	if err := j.validateSetAvailableMemoryMbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availableMemoryMb",
		val,
	)
}

func (j *jsiiProxy_CloudfunctionsFunction)SetBuildEnvironmentVariables(val *map[string]*string) {
	if err := j.validateSetBuildEnvironmentVariablesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"buildEnvironmentVariables",
		val,
	)
}

func (j *jsiiProxy_CloudfunctionsFunction)SetBuildServiceAccount(val *string) {
	if err := j.validateSetBuildServiceAccountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"buildServiceAccount",
		val,
	)
}

func (j *jsiiProxy_CloudfunctionsFunction)SetBuildWorkerPool(val *string) {
	if err := j.validateSetBuildWorkerPoolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"buildWorkerPool",
		val,
	)
}

func (j *jsiiProxy_CloudfunctionsFunction)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_CloudfunctionsFunction)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CloudfunctionsFunction)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CloudfunctionsFunction)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CloudfunctionsFunction)SetDockerRegistry(val *string) {
	if err := j.validateSetDockerRegistryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dockerRegistry",
		val,
	)
}

func (j *jsiiProxy_CloudfunctionsFunction)SetDockerRepository(val *string) {
	if err := j.validateSetDockerRepositoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dockerRepository",
		val,
	)
}

func (j *jsiiProxy_CloudfunctionsFunction)SetEntryPoint(val *string) {
	if err := j.validateSetEntryPointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"entryPoint",
		val,
	)
}

func (j *jsiiProxy_CloudfunctionsFunction)SetEnvironmentVariables(val *map[string]*string) {
	if err := j.validateSetEnvironmentVariablesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"environmentVariables",
		val,
	)
}

func (j *jsiiProxy_CloudfunctionsFunction)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_CloudfunctionsFunction)SetHttpsTriggerSecurityLevel(val *string) {
	if err := j.validateSetHttpsTriggerSecurityLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpsTriggerSecurityLevel",
		val,
	)
}

func (j *jsiiProxy_CloudfunctionsFunction)SetHttpsTriggerUrl(val *string) {
	if err := j.validateSetHttpsTriggerUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpsTriggerUrl",
		val,
	)
}

func (j *jsiiProxy_CloudfunctionsFunction)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_CloudfunctionsFunction)SetIngressSettings(val *string) {
	if err := j.validateSetIngressSettingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ingressSettings",
		val,
	)
}

func (j *jsiiProxy_CloudfunctionsFunction)SetKmsKeyName(val *string) {
	if err := j.validateSetKmsKeyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyName",
		val,
	)
}

func (j *jsiiProxy_CloudfunctionsFunction)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_CloudfunctionsFunction)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_CloudfunctionsFunction)SetMaxInstances(val *float64) {
	if err := j.validateSetMaxInstancesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxInstances",
		val,
	)
}

func (j *jsiiProxy_CloudfunctionsFunction)SetMinInstances(val *float64) {
	if err := j.validateSetMinInstancesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minInstances",
		val,
	)
}

func (j *jsiiProxy_CloudfunctionsFunction)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CloudfunctionsFunction)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_CloudfunctionsFunction)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_CloudfunctionsFunction)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_CloudfunctionsFunction)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_CloudfunctionsFunction)SetRuntime(val *string) {
	if err := j.validateSetRuntimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runtime",
		val,
	)
}

func (j *jsiiProxy_CloudfunctionsFunction)SetServiceAccountEmail(val *string) {
	if err := j.validateSetServiceAccountEmailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccountEmail",
		val,
	)
}

func (j *jsiiProxy_CloudfunctionsFunction)SetSourceArchiveBucket(val *string) {
	if err := j.validateSetSourceArchiveBucketParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceArchiveBucket",
		val,
	)
}

func (j *jsiiProxy_CloudfunctionsFunction)SetSourceArchiveObject(val *string) {
	if err := j.validateSetSourceArchiveObjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceArchiveObject",
		val,
	)
}

func (j *jsiiProxy_CloudfunctionsFunction)SetTimeout(val *float64) {
	if err := j.validateSetTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeout",
		val,
	)
}

func (j *jsiiProxy_CloudfunctionsFunction)SetTriggerHttp(val interface{}) {
	if err := j.validateSetTriggerHttpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"triggerHttp",
		val,
	)
}

func (j *jsiiProxy_CloudfunctionsFunction)SetVpcConnector(val *string) {
	if err := j.validateSetVpcConnectorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcConnector",
		val,
	)
}

func (j *jsiiProxy_CloudfunctionsFunction)SetVpcConnectorEgressSettings(val *string) {
	if err := j.validateSetVpcConnectorEgressSettingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcConnectorEgressSettings",
		val,
	)
}

// Generates CDKTN code for importing a CloudfunctionsFunction resource upon running "cdktn plan <stack-name>".
func CloudfunctionsFunction_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateCloudfunctionsFunction_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.cloudfunctionsFunction.CloudfunctionsFunction",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
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
func CloudfunctionsFunction_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCloudfunctionsFunction_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.cloudfunctionsFunction.CloudfunctionsFunction",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CloudfunctionsFunction_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCloudfunctionsFunction_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.cloudfunctionsFunction.CloudfunctionsFunction",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CloudfunctionsFunction_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCloudfunctionsFunction_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.cloudfunctionsFunction.CloudfunctionsFunction",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CloudfunctionsFunction_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-google.cloudfunctionsFunction.CloudfunctionsFunction",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CloudfunctionsFunction) AddMoveTarget(moveTarget *string) {
	if err := c.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (c *jsiiProxy_CloudfunctionsFunction) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CloudfunctionsFunction) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfunctionsFunction) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfunctionsFunction) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfunctionsFunction) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfunctionsFunction) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfunctionsFunction) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfunctionsFunction) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfunctionsFunction) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfunctionsFunction) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfunctionsFunction) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfunctionsFunction) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := c.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (c *jsiiProxy_CloudfunctionsFunction) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfunctionsFunction) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := c.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfunctionsFunction) MoveFromId(id *string) {
	if err := c.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveFromId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_CloudfunctionsFunction) MoveTo(moveTarget *string, index interface{}) {
	if err := c.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (c *jsiiProxy_CloudfunctionsFunction) MoveToId(id *string) {
	if err := c.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveToId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_CloudfunctionsFunction) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CloudfunctionsFunction) PutAutomaticUpdatePolicy(value *CloudfunctionsFunctionAutomaticUpdatePolicy) {
	if err := c.validatePutAutomaticUpdatePolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAutomaticUpdatePolicy",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfunctionsFunction) PutEventTrigger(value *CloudfunctionsFunctionEventTrigger) {
	if err := c.validatePutEventTriggerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putEventTrigger",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfunctionsFunction) PutOnDeployUpdatePolicy(value *CloudfunctionsFunctionOnDeployUpdatePolicy) {
	if err := c.validatePutOnDeployUpdatePolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putOnDeployUpdatePolicy",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfunctionsFunction) PutSecretEnvironmentVariables(value interface{}) {
	if err := c.validatePutSecretEnvironmentVariablesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSecretEnvironmentVariables",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfunctionsFunction) PutSecretVolumes(value interface{}) {
	if err := c.validatePutSecretVolumesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSecretVolumes",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfunctionsFunction) PutSourceRepository(value *CloudfunctionsFunctionSourceRepository) {
	if err := c.validatePutSourceRepositoryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSourceRepository",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfunctionsFunction) PutTimeouts(value *CloudfunctionsFunctionTimeouts) {
	if err := c.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfunctionsFunction) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := c.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (c *jsiiProxy_CloudfunctionsFunction) ResetAutomaticUpdatePolicy() {
	_jsii_.InvokeVoid(
		c,
		"resetAutomaticUpdatePolicy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfunctionsFunction) ResetAvailableMemoryMb() {
	_jsii_.InvokeVoid(
		c,
		"resetAvailableMemoryMb",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfunctionsFunction) ResetBuildEnvironmentVariables() {
	_jsii_.InvokeVoid(
		c,
		"resetBuildEnvironmentVariables",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfunctionsFunction) ResetBuildServiceAccount() {
	_jsii_.InvokeVoid(
		c,
		"resetBuildServiceAccount",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfunctionsFunction) ResetBuildWorkerPool() {
	_jsii_.InvokeVoid(
		c,
		"resetBuildWorkerPool",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfunctionsFunction) ResetDescription() {
	_jsii_.InvokeVoid(
		c,
		"resetDescription",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfunctionsFunction) ResetDockerRegistry() {
	_jsii_.InvokeVoid(
		c,
		"resetDockerRegistry",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfunctionsFunction) ResetDockerRepository() {
	_jsii_.InvokeVoid(
		c,
		"resetDockerRepository",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfunctionsFunction) ResetEntryPoint() {
	_jsii_.InvokeVoid(
		c,
		"resetEntryPoint",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfunctionsFunction) ResetEnvironmentVariables() {
	_jsii_.InvokeVoid(
		c,
		"resetEnvironmentVariables",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfunctionsFunction) ResetEventTrigger() {
	_jsii_.InvokeVoid(
		c,
		"resetEventTrigger",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfunctionsFunction) ResetHttpsTriggerSecurityLevel() {
	_jsii_.InvokeVoid(
		c,
		"resetHttpsTriggerSecurityLevel",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfunctionsFunction) ResetHttpsTriggerUrl() {
	_jsii_.InvokeVoid(
		c,
		"resetHttpsTriggerUrl",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfunctionsFunction) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfunctionsFunction) ResetIngressSettings() {
	_jsii_.InvokeVoid(
		c,
		"resetIngressSettings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfunctionsFunction) ResetKmsKeyName() {
	_jsii_.InvokeVoid(
		c,
		"resetKmsKeyName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfunctionsFunction) ResetLabels() {
	_jsii_.InvokeVoid(
		c,
		"resetLabels",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfunctionsFunction) ResetMaxInstances() {
	_jsii_.InvokeVoid(
		c,
		"resetMaxInstances",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfunctionsFunction) ResetMinInstances() {
	_jsii_.InvokeVoid(
		c,
		"resetMinInstances",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfunctionsFunction) ResetOnDeployUpdatePolicy() {
	_jsii_.InvokeVoid(
		c,
		"resetOnDeployUpdatePolicy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfunctionsFunction) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfunctionsFunction) ResetProject() {
	_jsii_.InvokeVoid(
		c,
		"resetProject",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfunctionsFunction) ResetRegion() {
	_jsii_.InvokeVoid(
		c,
		"resetRegion",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfunctionsFunction) ResetSecretEnvironmentVariables() {
	_jsii_.InvokeVoid(
		c,
		"resetSecretEnvironmentVariables",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfunctionsFunction) ResetSecretVolumes() {
	_jsii_.InvokeVoid(
		c,
		"resetSecretVolumes",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfunctionsFunction) ResetServiceAccountEmail() {
	_jsii_.InvokeVoid(
		c,
		"resetServiceAccountEmail",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfunctionsFunction) ResetSourceArchiveBucket() {
	_jsii_.InvokeVoid(
		c,
		"resetSourceArchiveBucket",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfunctionsFunction) ResetSourceArchiveObject() {
	_jsii_.InvokeVoid(
		c,
		"resetSourceArchiveObject",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfunctionsFunction) ResetSourceRepository() {
	_jsii_.InvokeVoid(
		c,
		"resetSourceRepository",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfunctionsFunction) ResetTimeout() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeout",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfunctionsFunction) ResetTimeouts() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfunctionsFunction) ResetTriggerHttp() {
	_jsii_.InvokeVoid(
		c,
		"resetTriggerHttp",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfunctionsFunction) ResetVpcConnector() {
	_jsii_.InvokeVoid(
		c,
		"resetVpcConnector",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfunctionsFunction) ResetVpcConnectorEgressSettings() {
	_jsii_.InvokeVoid(
		c,
		"resetVpcConnectorEgressSettings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfunctionsFunction) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfunctionsFunction) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfunctionsFunction) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfunctionsFunction) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfunctionsFunction) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfunctionsFunction) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfunctionsFunction) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		c,
		"with",
		args,
		&returns,
	)

	return returns
}

