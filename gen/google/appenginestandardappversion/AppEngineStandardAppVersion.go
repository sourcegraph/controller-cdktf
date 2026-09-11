package appenginestandardappversion

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google/appenginestandardappversion/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/app_engine_standard_app_version google_app_engine_standard_app_version}.
type AppEngineStandardAppVersion interface {
	cdktn.TerraformResource
	AppEngineApis() interface{}
	SetAppEngineApis(val interface{})
	AppEngineApisInput() interface{}
	AutomaticScaling() AppEngineStandardAppVersionAutomaticScalingOutputReference
	AutomaticScalingInput() *AppEngineStandardAppVersionAutomaticScaling
	BasicScaling() AppEngineStandardAppVersionBasicScalingOutputReference
	BasicScalingInput() *AppEngineStandardAppVersionBasicScaling
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
	DeleteServiceOnDestroy() interface{}
	SetDeleteServiceOnDestroy(val interface{})
	DeleteServiceOnDestroyInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Deployment() AppEngineStandardAppVersionDeploymentOutputReference
	DeploymentInput() *AppEngineStandardAppVersionDeployment
	Entrypoint() AppEngineStandardAppVersionEntrypointOutputReference
	EntrypointInput() *AppEngineStandardAppVersionEntrypoint
	EnvVariables() *map[string]*string
	SetEnvVariables(val *map[string]*string)
	EnvVariablesInput() *map[string]*string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Handlers() AppEngineStandardAppVersionHandlersList
	HandlersInput() interface{}
	Id() *string
	SetId(val *string)
	IdInput() *string
	InboundServices() *[]*string
	SetInboundServices(val *[]*string)
	InboundServicesInput() *[]*string
	InstanceClass() *string
	SetInstanceClass(val *string)
	InstanceClassInput() *string
	Libraries() AppEngineStandardAppVersionLibrariesList
	LibrariesInput() interface{}
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	ManualScaling() AppEngineStandardAppVersionManualScalingOutputReference
	ManualScalingInput() *AppEngineStandardAppVersionManualScaling
	Name() *string
	// The tree node.
	Node() constructs.Node
	NoopOnDestroy() interface{}
	SetNoopOnDestroy(val interface{})
	NoopOnDestroyInput() interface{}
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
	Runtime() *string
	SetRuntime(val *string)
	RuntimeApiVersion() *string
	SetRuntimeApiVersion(val *string)
	RuntimeApiVersionInput() *string
	RuntimeInput() *string
	Service() *string
	SetService(val *string)
	ServiceAccount() *string
	SetServiceAccount(val *string)
	ServiceAccountInput() *string
	ServiceInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Threadsafe() interface{}
	SetThreadsafe(val interface{})
	ThreadsafeInput() interface{}
	Timeouts() AppEngineStandardAppVersionTimeoutsOutputReference
	TimeoutsInput() interface{}
	VersionId() *string
	SetVersionId(val *string)
	VersionIdInput() *string
	VpcAccessConnector() AppEngineStandardAppVersionVpcAccessConnectorOutputReference
	VpcAccessConnectorInput() *AppEngineStandardAppVersionVpcAccessConnector
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
	PutAutomaticScaling(value *AppEngineStandardAppVersionAutomaticScaling)
	PutBasicScaling(value *AppEngineStandardAppVersionBasicScaling)
	PutDeployment(value *AppEngineStandardAppVersionDeployment)
	PutEntrypoint(value *AppEngineStandardAppVersionEntrypoint)
	PutHandlers(value interface{})
	PutLibraries(value interface{})
	PutManualScaling(value *AppEngineStandardAppVersionManualScaling)
	PutTimeouts(value *AppEngineStandardAppVersionTimeouts)
	PutVpcAccessConnector(value *AppEngineStandardAppVersionVpcAccessConnector)
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
	ResetAppEngineApis()
	ResetAutomaticScaling()
	ResetBasicScaling()
	ResetDeleteServiceOnDestroy()
	ResetEnvVariables()
	ResetHandlers()
	ResetId()
	ResetInboundServices()
	ResetInstanceClass()
	ResetLibraries()
	ResetManualScaling()
	ResetNoopOnDestroy()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetRuntimeApiVersion()
	ResetServiceAccount()
	ResetThreadsafe()
	ResetTimeouts()
	ResetVersionId()
	ResetVpcAccessConnector()
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

// The jsii proxy struct for AppEngineStandardAppVersion
type jsiiProxy_AppEngineStandardAppVersion struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AppEngineStandardAppVersion) AppEngineApis() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"appEngineApis",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersion) AppEngineApisInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"appEngineApisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersion) AutomaticScaling() AppEngineStandardAppVersionAutomaticScalingOutputReference {
	var returns AppEngineStandardAppVersionAutomaticScalingOutputReference
	_jsii_.Get(
		j,
		"automaticScaling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersion) AutomaticScalingInput() *AppEngineStandardAppVersionAutomaticScaling {
	var returns *AppEngineStandardAppVersionAutomaticScaling
	_jsii_.Get(
		j,
		"automaticScalingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersion) BasicScaling() AppEngineStandardAppVersionBasicScalingOutputReference {
	var returns AppEngineStandardAppVersionBasicScalingOutputReference
	_jsii_.Get(
		j,
		"basicScaling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersion) BasicScalingInput() *AppEngineStandardAppVersionBasicScaling {
	var returns *AppEngineStandardAppVersionBasicScaling
	_jsii_.Get(
		j,
		"basicScalingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersion) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersion) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersion) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersion) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersion) DeleteServiceOnDestroy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteServiceOnDestroy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersion) DeleteServiceOnDestroyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteServiceOnDestroyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersion) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersion) Deployment() AppEngineStandardAppVersionDeploymentOutputReference {
	var returns AppEngineStandardAppVersionDeploymentOutputReference
	_jsii_.Get(
		j,
		"deployment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersion) DeploymentInput() *AppEngineStandardAppVersionDeployment {
	var returns *AppEngineStandardAppVersionDeployment
	_jsii_.Get(
		j,
		"deploymentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersion) Entrypoint() AppEngineStandardAppVersionEntrypointOutputReference {
	var returns AppEngineStandardAppVersionEntrypointOutputReference
	_jsii_.Get(
		j,
		"entrypoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersion) EntrypointInput() *AppEngineStandardAppVersionEntrypoint {
	var returns *AppEngineStandardAppVersionEntrypoint
	_jsii_.Get(
		j,
		"entrypointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersion) EnvVariables() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"envVariables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersion) EnvVariablesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"envVariablesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersion) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersion) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersion) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersion) Handlers() AppEngineStandardAppVersionHandlersList {
	var returns AppEngineStandardAppVersionHandlersList
	_jsii_.Get(
		j,
		"handlers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersion) HandlersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"handlersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersion) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersion) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersion) InboundServices() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"inboundServices",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersion) InboundServicesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"inboundServicesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersion) InstanceClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersion) InstanceClassInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceClassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersion) Libraries() AppEngineStandardAppVersionLibrariesList {
	var returns AppEngineStandardAppVersionLibrariesList
	_jsii_.Get(
		j,
		"libraries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersion) LibrariesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"librariesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersion) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersion) ManualScaling() AppEngineStandardAppVersionManualScalingOutputReference {
	var returns AppEngineStandardAppVersionManualScalingOutputReference
	_jsii_.Get(
		j,
		"manualScaling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersion) ManualScalingInput() *AppEngineStandardAppVersionManualScaling {
	var returns *AppEngineStandardAppVersionManualScaling
	_jsii_.Get(
		j,
		"manualScalingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersion) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersion) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersion) NoopOnDestroy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noopOnDestroy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersion) NoopOnDestroyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noopOnDestroyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersion) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersion) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersion) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersion) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersion) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersion) Runtime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersion) RuntimeApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeApiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersion) RuntimeApiVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeApiVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersion) RuntimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersion) Service() *string {
	var returns *string
	_jsii_.Get(
		j,
		"service",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersion) ServiceAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersion) ServiceAccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersion) ServiceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersion) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersion) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersion) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersion) Threadsafe() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"threadsafe",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersion) ThreadsafeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"threadsafeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersion) Timeouts() AppEngineStandardAppVersionTimeoutsOutputReference {
	var returns AppEngineStandardAppVersionTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersion) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersion) VersionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersion) VersionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersion) VpcAccessConnector() AppEngineStandardAppVersionVpcAccessConnectorOutputReference {
	var returns AppEngineStandardAppVersionVpcAccessConnectorOutputReference
	_jsii_.Get(
		j,
		"vpcAccessConnector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersion) VpcAccessConnectorInput() *AppEngineStandardAppVersionVpcAccessConnector {
	var returns *AppEngineStandardAppVersionVpcAccessConnector
	_jsii_.Get(
		j,
		"vpcAccessConnectorInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/app_engine_standard_app_version google_app_engine_standard_app_version} Resource.
func NewAppEngineStandardAppVersion(scope constructs.Construct, id *string, config *AppEngineStandardAppVersionConfig) AppEngineStandardAppVersion {
	_init_.Initialize()

	if err := validateNewAppEngineStandardAppVersionParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppEngineStandardAppVersion{}

	_jsii_.Create(
		"@cdktn/provider-google.appEngineStandardAppVersion.AppEngineStandardAppVersion",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/app_engine_standard_app_version google_app_engine_standard_app_version} Resource.
func NewAppEngineStandardAppVersion_Override(a AppEngineStandardAppVersion, scope constructs.Construct, id *string, config *AppEngineStandardAppVersionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.appEngineStandardAppVersion.AppEngineStandardAppVersion",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AppEngineStandardAppVersion)SetAppEngineApis(val interface{}) {
	if err := j.validateSetAppEngineApisParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appEngineApis",
		val,
	)
}

func (j *jsiiProxy_AppEngineStandardAppVersion)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AppEngineStandardAppVersion)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AppEngineStandardAppVersion)SetDeleteServiceOnDestroy(val interface{}) {
	if err := j.validateSetDeleteServiceOnDestroyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteServiceOnDestroy",
		val,
	)
}

func (j *jsiiProxy_AppEngineStandardAppVersion)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AppEngineStandardAppVersion)SetEnvVariables(val *map[string]*string) {
	if err := j.validateSetEnvVariablesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"envVariables",
		val,
	)
}

func (j *jsiiProxy_AppEngineStandardAppVersion)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AppEngineStandardAppVersion)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AppEngineStandardAppVersion)SetInboundServices(val *[]*string) {
	if err := j.validateSetInboundServicesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inboundServices",
		val,
	)
}

func (j *jsiiProxy_AppEngineStandardAppVersion)SetInstanceClass(val *string) {
	if err := j.validateSetInstanceClassParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceClass",
		val,
	)
}

func (j *jsiiProxy_AppEngineStandardAppVersion)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AppEngineStandardAppVersion)SetNoopOnDestroy(val interface{}) {
	if err := j.validateSetNoopOnDestroyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noopOnDestroy",
		val,
	)
}

func (j *jsiiProxy_AppEngineStandardAppVersion)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_AppEngineStandardAppVersion)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AppEngineStandardAppVersion)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AppEngineStandardAppVersion)SetRuntime(val *string) {
	if err := j.validateSetRuntimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runtime",
		val,
	)
}

func (j *jsiiProxy_AppEngineStandardAppVersion)SetRuntimeApiVersion(val *string) {
	if err := j.validateSetRuntimeApiVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runtimeApiVersion",
		val,
	)
}

func (j *jsiiProxy_AppEngineStandardAppVersion)SetService(val *string) {
	if err := j.validateSetServiceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"service",
		val,
	)
}

func (j *jsiiProxy_AppEngineStandardAppVersion)SetServiceAccount(val *string) {
	if err := j.validateSetServiceAccountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccount",
		val,
	)
}

func (j *jsiiProxy_AppEngineStandardAppVersion)SetThreadsafe(val interface{}) {
	if err := j.validateSetThreadsafeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"threadsafe",
		val,
	)
}

func (j *jsiiProxy_AppEngineStandardAppVersion)SetVersionId(val *string) {
	if err := j.validateSetVersionIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"versionId",
		val,
	)
}

// Generates CDKTN code for importing a AppEngineStandardAppVersion resource upon running "cdktn plan <stack-name>".
func AppEngineStandardAppVersion_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAppEngineStandardAppVersion_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.appEngineStandardAppVersion.AppEngineStandardAppVersion",
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
func AppEngineStandardAppVersion_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAppEngineStandardAppVersion_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.appEngineStandardAppVersion.AppEngineStandardAppVersion",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AppEngineStandardAppVersion_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAppEngineStandardAppVersion_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.appEngineStandardAppVersion.AppEngineStandardAppVersion",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AppEngineStandardAppVersion_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAppEngineStandardAppVersion_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.appEngineStandardAppVersion.AppEngineStandardAppVersion",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AppEngineStandardAppVersion_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-google.appEngineStandardAppVersion.AppEngineStandardAppVersion",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AppEngineStandardAppVersion) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AppEngineStandardAppVersion) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AppEngineStandardAppVersion) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppEngineStandardAppVersion) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppEngineStandardAppVersion) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppEngineStandardAppVersion) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppEngineStandardAppVersion) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppEngineStandardAppVersion) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppEngineStandardAppVersion) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppEngineStandardAppVersion) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppEngineStandardAppVersion) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppEngineStandardAppVersion) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppEngineStandardAppVersion) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AppEngineStandardAppVersion) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppEngineStandardAppVersion) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := a.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppEngineStandardAppVersion) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AppEngineStandardAppVersion) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AppEngineStandardAppVersion) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AppEngineStandardAppVersion) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AppEngineStandardAppVersion) PutAutomaticScaling(value *AppEngineStandardAppVersionAutomaticScaling) {
	if err := a.validatePutAutomaticScalingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAutomaticScaling",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppEngineStandardAppVersion) PutBasicScaling(value *AppEngineStandardAppVersionBasicScaling) {
	if err := a.validatePutBasicScalingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putBasicScaling",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppEngineStandardAppVersion) PutDeployment(value *AppEngineStandardAppVersionDeployment) {
	if err := a.validatePutDeploymentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDeployment",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppEngineStandardAppVersion) PutEntrypoint(value *AppEngineStandardAppVersionEntrypoint) {
	if err := a.validatePutEntrypointParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEntrypoint",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppEngineStandardAppVersion) PutHandlers(value interface{}) {
	if err := a.validatePutHandlersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putHandlers",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppEngineStandardAppVersion) PutLibraries(value interface{}) {
	if err := a.validatePutLibrariesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLibraries",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppEngineStandardAppVersion) PutManualScaling(value *AppEngineStandardAppVersionManualScaling) {
	if err := a.validatePutManualScalingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putManualScaling",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppEngineStandardAppVersion) PutTimeouts(value *AppEngineStandardAppVersionTimeouts) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppEngineStandardAppVersion) PutVpcAccessConnector(value *AppEngineStandardAppVersionVpcAccessConnector) {
	if err := a.validatePutVpcAccessConnectorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putVpcAccessConnector",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppEngineStandardAppVersion) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AppEngineStandardAppVersion) ResetAppEngineApis() {
	_jsii_.InvokeVoid(
		a,
		"resetAppEngineApis",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineStandardAppVersion) ResetAutomaticScaling() {
	_jsii_.InvokeVoid(
		a,
		"resetAutomaticScaling",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineStandardAppVersion) ResetBasicScaling() {
	_jsii_.InvokeVoid(
		a,
		"resetBasicScaling",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineStandardAppVersion) ResetDeleteServiceOnDestroy() {
	_jsii_.InvokeVoid(
		a,
		"resetDeleteServiceOnDestroy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineStandardAppVersion) ResetEnvVariables() {
	_jsii_.InvokeVoid(
		a,
		"resetEnvVariables",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineStandardAppVersion) ResetHandlers() {
	_jsii_.InvokeVoid(
		a,
		"resetHandlers",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineStandardAppVersion) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineStandardAppVersion) ResetInboundServices() {
	_jsii_.InvokeVoid(
		a,
		"resetInboundServices",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineStandardAppVersion) ResetInstanceClass() {
	_jsii_.InvokeVoid(
		a,
		"resetInstanceClass",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineStandardAppVersion) ResetLibraries() {
	_jsii_.InvokeVoid(
		a,
		"resetLibraries",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineStandardAppVersion) ResetManualScaling() {
	_jsii_.InvokeVoid(
		a,
		"resetManualScaling",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineStandardAppVersion) ResetNoopOnDestroy() {
	_jsii_.InvokeVoid(
		a,
		"resetNoopOnDestroy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineStandardAppVersion) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineStandardAppVersion) ResetProject() {
	_jsii_.InvokeVoid(
		a,
		"resetProject",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineStandardAppVersion) ResetRuntimeApiVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetRuntimeApiVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineStandardAppVersion) ResetServiceAccount() {
	_jsii_.InvokeVoid(
		a,
		"resetServiceAccount",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineStandardAppVersion) ResetThreadsafe() {
	_jsii_.InvokeVoid(
		a,
		"resetThreadsafe",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineStandardAppVersion) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineStandardAppVersion) ResetVersionId() {
	_jsii_.InvokeVoid(
		a,
		"resetVersionId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineStandardAppVersion) ResetVpcAccessConnector() {
	_jsii_.InvokeVoid(
		a,
		"resetVpcAccessConnector",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineStandardAppVersion) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppEngineStandardAppVersion) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppEngineStandardAppVersion) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppEngineStandardAppVersion) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppEngineStandardAppVersion) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppEngineStandardAppVersion) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppEngineStandardAppVersion) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		a,
		"with",
		args,
		&returns,
	)

	return returns
}

