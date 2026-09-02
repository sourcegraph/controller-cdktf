package workstationsworkstationconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google/workstationsworkstationconfig/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/workstations_workstation_config google_workstations_workstation_config}.
type WorkstationsWorkstationConfigA interface {
	cdktn.TerraformResource
	AllowedPorts() WorkstationsWorkstationConfigAllowedPortsList
	AllowedPortsInput() interface{}
	Annotations() *map[string]*string
	SetAnnotations(val *map[string]*string)
	AnnotationsInput() *map[string]*string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	Conditions() WorkstationsWorkstationConfigConditionsList
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	Container() WorkstationsWorkstationConfigContainerOutputReference
	ContainerInput() *WorkstationsWorkstationConfigContainer
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreateTime() *string
	Degraded() cdktn.IResolvable
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DisableTcpConnections() interface{}
	SetDisableTcpConnections(val interface{})
	DisableTcpConnectionsInput() interface{}
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	EffectiveAnnotations() cdktn.StringMap
	EffectiveLabels() cdktn.StringMap
	EnableAuditAgent() interface{}
	SetEnableAuditAgent(val interface{})
	EnableAuditAgentInput() interface{}
	EncryptionKey() WorkstationsWorkstationConfigEncryptionKeyOutputReference
	EncryptionKeyInput() *WorkstationsWorkstationConfigEncryptionKey
	EphemeralDirectories() WorkstationsWorkstationConfigEphemeralDirectoriesList
	EphemeralDirectoriesInput() interface{}
	Etag() *string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Host() WorkstationsWorkstationConfigHostOutputReference
	HostInput() *WorkstationsWorkstationConfigHost
	Id() *string
	SetId(val *string)
	IdInput() *string
	IdleTimeout() *string
	SetIdleTimeout(val *string)
	IdleTimeoutInput() *string
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	MaxUsableWorkstations() *float64
	SetMaxUsableWorkstations(val *float64)
	MaxUsableWorkstationsInput() *float64
	Name() *string
	// The tree node.
	Node() constructs.Node
	PersistentDirectories() WorkstationsWorkstationConfigPersistentDirectoriesList
	PersistentDirectoriesInput() interface{}
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
	ReadinessChecks() WorkstationsWorkstationConfigReadinessChecksList
	ReadinessChecksInput() interface{}
	ReplicaZones() *[]*string
	SetReplicaZones(val *[]*string)
	ReplicaZonesInput() *[]*string
	RunningTimeout() *string
	SetRunningTimeout(val *string)
	RunningTimeoutInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	TerraformLabels() cdktn.StringMap
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() WorkstationsWorkstationConfigTimeoutsOutputReference
	TimeoutsInput() interface{}
	Uid() *string
	WorkstationClusterId() *string
	SetWorkstationClusterId(val *string)
	WorkstationClusterIdInput() *string
	WorkstationConfigId() *string
	SetWorkstationConfigId(val *string)
	WorkstationConfigIdInput() *string
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
	PutAllowedPorts(value interface{})
	PutContainer(value *WorkstationsWorkstationConfigContainer)
	PutEncryptionKey(value *WorkstationsWorkstationConfigEncryptionKey)
	PutEphemeralDirectories(value interface{})
	PutHost(value *WorkstationsWorkstationConfigHost)
	PutPersistentDirectories(value interface{})
	PutReadinessChecks(value interface{})
	PutTimeouts(value *WorkstationsWorkstationConfigTimeouts)
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
	ResetAllowedPorts()
	ResetAnnotations()
	ResetContainer()
	ResetDisableTcpConnections()
	ResetDisplayName()
	ResetEnableAuditAgent()
	ResetEncryptionKey()
	ResetEphemeralDirectories()
	ResetHost()
	ResetId()
	ResetIdleTimeout()
	ResetLabels()
	ResetMaxUsableWorkstations()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPersistentDirectories()
	ResetProject()
	ResetReadinessChecks()
	ResetReplicaZones()
	ResetRunningTimeout()
	ResetTimeouts()
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

// The jsii proxy struct for WorkstationsWorkstationConfigA
type jsiiProxy_WorkstationsWorkstationConfigA struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_WorkstationsWorkstationConfigA) AllowedPorts() WorkstationsWorkstationConfigAllowedPortsList {
	var returns WorkstationsWorkstationConfigAllowedPortsList
	_jsii_.Get(
		j,
		"allowedPorts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigA) AllowedPortsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowedPortsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigA) Annotations() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"annotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigA) AnnotationsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"annotationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigA) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigA) Conditions() WorkstationsWorkstationConfigConditionsList {
	var returns WorkstationsWorkstationConfigConditionsList
	_jsii_.Get(
		j,
		"conditions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigA) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigA) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigA) Container() WorkstationsWorkstationConfigContainerOutputReference {
	var returns WorkstationsWorkstationConfigContainerOutputReference
	_jsii_.Get(
		j,
		"container",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigA) ContainerInput() *WorkstationsWorkstationConfigContainer {
	var returns *WorkstationsWorkstationConfigContainer
	_jsii_.Get(
		j,
		"containerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigA) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigA) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigA) Degraded() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"degraded",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigA) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigA) DisableTcpConnections() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableTcpConnections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigA) DisableTcpConnectionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableTcpConnectionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigA) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigA) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigA) EffectiveAnnotations() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"effectiveAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigA) EffectiveLabels() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"effectiveLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigA) EnableAuditAgent() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAuditAgent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigA) EnableAuditAgentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAuditAgentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigA) EncryptionKey() WorkstationsWorkstationConfigEncryptionKeyOutputReference {
	var returns WorkstationsWorkstationConfigEncryptionKeyOutputReference
	_jsii_.Get(
		j,
		"encryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigA) EncryptionKeyInput() *WorkstationsWorkstationConfigEncryptionKey {
	var returns *WorkstationsWorkstationConfigEncryptionKey
	_jsii_.Get(
		j,
		"encryptionKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigA) EphemeralDirectories() WorkstationsWorkstationConfigEphemeralDirectoriesList {
	var returns WorkstationsWorkstationConfigEphemeralDirectoriesList
	_jsii_.Get(
		j,
		"ephemeralDirectories",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigA) EphemeralDirectoriesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ephemeralDirectoriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigA) Etag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigA) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigA) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigA) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigA) Host() WorkstationsWorkstationConfigHostOutputReference {
	var returns WorkstationsWorkstationConfigHostOutputReference
	_jsii_.Get(
		j,
		"host",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigA) HostInput() *WorkstationsWorkstationConfigHost {
	var returns *WorkstationsWorkstationConfigHost
	_jsii_.Get(
		j,
		"hostInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigA) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigA) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigA) IdleTimeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idleTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigA) IdleTimeoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idleTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigA) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigA) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigA) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigA) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigA) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigA) MaxUsableWorkstations() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxUsableWorkstations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigA) MaxUsableWorkstationsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxUsableWorkstationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigA) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigA) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigA) PersistentDirectories() WorkstationsWorkstationConfigPersistentDirectoriesList {
	var returns WorkstationsWorkstationConfigPersistentDirectoriesList
	_jsii_.Get(
		j,
		"persistentDirectories",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigA) PersistentDirectoriesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"persistentDirectoriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigA) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigA) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigA) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigA) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigA) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigA) ReadinessChecks() WorkstationsWorkstationConfigReadinessChecksList {
	var returns WorkstationsWorkstationConfigReadinessChecksList
	_jsii_.Get(
		j,
		"readinessChecks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigA) ReadinessChecksInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readinessChecksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigA) ReplicaZones() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"replicaZones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigA) ReplicaZonesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"replicaZonesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigA) RunningTimeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runningTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigA) RunningTimeoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runningTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigA) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigA) TerraformLabels() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"terraformLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigA) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigA) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigA) Timeouts() WorkstationsWorkstationConfigTimeoutsOutputReference {
	var returns WorkstationsWorkstationConfigTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigA) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigA) Uid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigA) WorkstationClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workstationClusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigA) WorkstationClusterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workstationClusterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigA) WorkstationConfigId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workstationConfigId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigA) WorkstationConfigIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workstationConfigIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/workstations_workstation_config google_workstations_workstation_config} Resource.
func NewWorkstationsWorkstationConfigA(scope constructs.Construct, id *string, config *WorkstationsWorkstationConfigAConfig) WorkstationsWorkstationConfigA {
	_init_.Initialize()

	if err := validateNewWorkstationsWorkstationConfigAParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_WorkstationsWorkstationConfigA{}

	_jsii_.Create(
		"@cdktn/provider-google.workstationsWorkstationConfig.WorkstationsWorkstationConfigA",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/workstations_workstation_config google_workstations_workstation_config} Resource.
func NewWorkstationsWorkstationConfigA_Override(w WorkstationsWorkstationConfigA, scope constructs.Construct, id *string, config *WorkstationsWorkstationConfigAConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.workstationsWorkstationConfig.WorkstationsWorkstationConfigA",
		[]interface{}{scope, id, config},
		w,
	)
}

func (j *jsiiProxy_WorkstationsWorkstationConfigA)SetAnnotations(val *map[string]*string) {
	if err := j.validateSetAnnotationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"annotations",
		val,
	)
}

func (j *jsiiProxy_WorkstationsWorkstationConfigA)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_WorkstationsWorkstationConfigA)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_WorkstationsWorkstationConfigA)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_WorkstationsWorkstationConfigA)SetDisableTcpConnections(val interface{}) {
	if err := j.validateSetDisableTcpConnectionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableTcpConnections",
		val,
	)
}

func (j *jsiiProxy_WorkstationsWorkstationConfigA)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_WorkstationsWorkstationConfigA)SetEnableAuditAgent(val interface{}) {
	if err := j.validateSetEnableAuditAgentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableAuditAgent",
		val,
	)
}

func (j *jsiiProxy_WorkstationsWorkstationConfigA)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_WorkstationsWorkstationConfigA)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_WorkstationsWorkstationConfigA)SetIdleTimeout(val *string) {
	if err := j.validateSetIdleTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"idleTimeout",
		val,
	)
}

func (j *jsiiProxy_WorkstationsWorkstationConfigA)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_WorkstationsWorkstationConfigA)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_WorkstationsWorkstationConfigA)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_WorkstationsWorkstationConfigA)SetMaxUsableWorkstations(val *float64) {
	if err := j.validateSetMaxUsableWorkstationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxUsableWorkstations",
		val,
	)
}

func (j *jsiiProxy_WorkstationsWorkstationConfigA)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_WorkstationsWorkstationConfigA)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_WorkstationsWorkstationConfigA)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_WorkstationsWorkstationConfigA)SetReplicaZones(val *[]*string) {
	if err := j.validateSetReplicaZonesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replicaZones",
		val,
	)
}

func (j *jsiiProxy_WorkstationsWorkstationConfigA)SetRunningTimeout(val *string) {
	if err := j.validateSetRunningTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runningTimeout",
		val,
	)
}

func (j *jsiiProxy_WorkstationsWorkstationConfigA)SetWorkstationClusterId(val *string) {
	if err := j.validateSetWorkstationClusterIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workstationClusterId",
		val,
	)
}

func (j *jsiiProxy_WorkstationsWorkstationConfigA)SetWorkstationConfigId(val *string) {
	if err := j.validateSetWorkstationConfigIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workstationConfigId",
		val,
	)
}

// Generates CDKTN code for importing a WorkstationsWorkstationConfigA resource upon running "cdktn plan <stack-name>".
func WorkstationsWorkstationConfigA_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateWorkstationsWorkstationConfigA_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.workstationsWorkstationConfig.WorkstationsWorkstationConfigA",
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
func WorkstationsWorkstationConfigA_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateWorkstationsWorkstationConfigA_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.workstationsWorkstationConfig.WorkstationsWorkstationConfigA",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func WorkstationsWorkstationConfigA_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateWorkstationsWorkstationConfigA_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.workstationsWorkstationConfig.WorkstationsWorkstationConfigA",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func WorkstationsWorkstationConfigA_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateWorkstationsWorkstationConfigA_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.workstationsWorkstationConfig.WorkstationsWorkstationConfigA",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func WorkstationsWorkstationConfigA_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-google.workstationsWorkstationConfig.WorkstationsWorkstationConfigA",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (w *jsiiProxy_WorkstationsWorkstationConfigA) AddMoveTarget(moveTarget *string) {
	if err := w.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (w *jsiiProxy_WorkstationsWorkstationConfigA) AddOverride(path *string, value interface{}) {
	if err := w.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (w *jsiiProxy_WorkstationsWorkstationConfigA) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := w.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkstationsWorkstationConfigA) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := w.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkstationsWorkstationConfigA) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := w.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkstationsWorkstationConfigA) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := w.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkstationsWorkstationConfigA) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := w.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkstationsWorkstationConfigA) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := w.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkstationsWorkstationConfigA) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := w.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkstationsWorkstationConfigA) GetStringAttribute(terraformAttribute *string) *string {
	if err := w.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkstationsWorkstationConfigA) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := w.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkstationsWorkstationConfigA) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkstationsWorkstationConfigA) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := w.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (w *jsiiProxy_WorkstationsWorkstationConfigA) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := w.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkstationsWorkstationConfigA) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := w.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		w,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkstationsWorkstationConfigA) MoveFromId(id *string) {
	if err := w.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"moveFromId",
		[]interface{}{id},
	)
}

func (w *jsiiProxy_WorkstationsWorkstationConfigA) MoveTo(moveTarget *string, index interface{}) {
	if err := w.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (w *jsiiProxy_WorkstationsWorkstationConfigA) MoveToId(id *string) {
	if err := w.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"moveToId",
		[]interface{}{id},
	)
}

func (w *jsiiProxy_WorkstationsWorkstationConfigA) OverrideLogicalId(newLogicalId *string) {
	if err := w.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (w *jsiiProxy_WorkstationsWorkstationConfigA) PutAllowedPorts(value interface{}) {
	if err := w.validatePutAllowedPortsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putAllowedPorts",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkstationsWorkstationConfigA) PutContainer(value *WorkstationsWorkstationConfigContainer) {
	if err := w.validatePutContainerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putContainer",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkstationsWorkstationConfigA) PutEncryptionKey(value *WorkstationsWorkstationConfigEncryptionKey) {
	if err := w.validatePutEncryptionKeyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putEncryptionKey",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkstationsWorkstationConfigA) PutEphemeralDirectories(value interface{}) {
	if err := w.validatePutEphemeralDirectoriesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putEphemeralDirectories",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkstationsWorkstationConfigA) PutHost(value *WorkstationsWorkstationConfigHost) {
	if err := w.validatePutHostParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putHost",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkstationsWorkstationConfigA) PutPersistentDirectories(value interface{}) {
	if err := w.validatePutPersistentDirectoriesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putPersistentDirectories",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkstationsWorkstationConfigA) PutReadinessChecks(value interface{}) {
	if err := w.validatePutReadinessChecksParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putReadinessChecks",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkstationsWorkstationConfigA) PutTimeouts(value *WorkstationsWorkstationConfigTimeouts) {
	if err := w.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkstationsWorkstationConfigA) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := w.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (w *jsiiProxy_WorkstationsWorkstationConfigA) ResetAllowedPorts() {
	_jsii_.InvokeVoid(
		w,
		"resetAllowedPorts",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkstationsWorkstationConfigA) ResetAnnotations() {
	_jsii_.InvokeVoid(
		w,
		"resetAnnotations",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkstationsWorkstationConfigA) ResetContainer() {
	_jsii_.InvokeVoid(
		w,
		"resetContainer",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkstationsWorkstationConfigA) ResetDisableTcpConnections() {
	_jsii_.InvokeVoid(
		w,
		"resetDisableTcpConnections",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkstationsWorkstationConfigA) ResetDisplayName() {
	_jsii_.InvokeVoid(
		w,
		"resetDisplayName",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkstationsWorkstationConfigA) ResetEnableAuditAgent() {
	_jsii_.InvokeVoid(
		w,
		"resetEnableAuditAgent",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkstationsWorkstationConfigA) ResetEncryptionKey() {
	_jsii_.InvokeVoid(
		w,
		"resetEncryptionKey",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkstationsWorkstationConfigA) ResetEphemeralDirectories() {
	_jsii_.InvokeVoid(
		w,
		"resetEphemeralDirectories",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkstationsWorkstationConfigA) ResetHost() {
	_jsii_.InvokeVoid(
		w,
		"resetHost",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkstationsWorkstationConfigA) ResetId() {
	_jsii_.InvokeVoid(
		w,
		"resetId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkstationsWorkstationConfigA) ResetIdleTimeout() {
	_jsii_.InvokeVoid(
		w,
		"resetIdleTimeout",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkstationsWorkstationConfigA) ResetLabels() {
	_jsii_.InvokeVoid(
		w,
		"resetLabels",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkstationsWorkstationConfigA) ResetMaxUsableWorkstations() {
	_jsii_.InvokeVoid(
		w,
		"resetMaxUsableWorkstations",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkstationsWorkstationConfigA) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		w,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkstationsWorkstationConfigA) ResetPersistentDirectories() {
	_jsii_.InvokeVoid(
		w,
		"resetPersistentDirectories",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkstationsWorkstationConfigA) ResetProject() {
	_jsii_.InvokeVoid(
		w,
		"resetProject",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkstationsWorkstationConfigA) ResetReadinessChecks() {
	_jsii_.InvokeVoid(
		w,
		"resetReadinessChecks",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkstationsWorkstationConfigA) ResetReplicaZones() {
	_jsii_.InvokeVoid(
		w,
		"resetReplicaZones",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkstationsWorkstationConfigA) ResetRunningTimeout() {
	_jsii_.InvokeVoid(
		w,
		"resetRunningTimeout",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkstationsWorkstationConfigA) ResetTimeouts() {
	_jsii_.InvokeVoid(
		w,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkstationsWorkstationConfigA) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkstationsWorkstationConfigA) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkstationsWorkstationConfigA) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkstationsWorkstationConfigA) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkstationsWorkstationConfigA) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkstationsWorkstationConfigA) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkstationsWorkstationConfigA) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		w,
		"with",
		args,
		&returns,
	)

	return returns
}

