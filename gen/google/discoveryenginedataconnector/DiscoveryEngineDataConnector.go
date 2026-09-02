package discoveryenginedataconnector

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google/discoveryenginedataconnector/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/discovery_engine_data_connector google_discovery_engine_data_connector}.
type DiscoveryEngineDataConnector interface {
	cdktn.TerraformResource
	ActionConfig() DiscoveryEngineDataConnectorActionConfigOutputReference
	ActionConfigInput() *DiscoveryEngineDataConnectorActionConfig
	ActionState() *string
	AutoRunDisabled() interface{}
	SetAutoRunDisabled(val interface{})
	AutoRunDisabledInput() interface{}
	BapConfig() DiscoveryEngineDataConnectorBapConfigOutputReference
	BapConfigInput() *DiscoveryEngineDataConnectorBapConfig
	BlockingReasons() *[]*string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	CollectionDisplayName() *string
	SetCollectionDisplayName(val *string)
	CollectionDisplayNameInput() *string
	CollectionId() *string
	SetCollectionId(val *string)
	CollectionIdInput() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	ConnectorModes() *[]*string
	SetConnectorModes(val *[]*string)
	ConnectorModesInput() *[]*string
	ConnectorType() *string
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreateTime() *string
	DataSource() *string
	SetDataSource(val *string)
	DataSourceInput() *string
	DataSourceVersion() *float64
	SetDataSourceVersion(val *float64)
	DataSourceVersionInput() *float64
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DestinationConfigs() DiscoveryEngineDataConnectorDestinationConfigsList
	DestinationConfigsInput() interface{}
	Entities() DiscoveryEngineDataConnectorEntitiesList
	EntitiesInput() interface{}
	Errors() DiscoveryEngineDataConnectorErrorsList
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	IncrementalRefreshInterval() *string
	SetIncrementalRefreshInterval(val *string)
	IncrementalRefreshIntervalInput() *string
	IncrementalSyncDisabled() interface{}
	SetIncrementalSyncDisabled(val interface{})
	IncrementalSyncDisabledInput() interface{}
	JsonParams() *string
	SetJsonParams(val *string)
	JsonParamsInput() *string
	KmsKeyName() *string
	SetKmsKeyName(val *string)
	KmsKeyNameInput() *string
	LastSyncTime() *string
	LatestPauseTime() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	Name() *string
	// The tree node.
	Node() constructs.Node
	Params() *map[string]*string
	SetParams(val *map[string]*string)
	ParamsInput() *map[string]*string
	PrivateConnectivityProjectId() *string
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
	RealtimeState() *string
	RefreshInterval() *string
	SetRefreshInterval(val *string)
	RefreshIntervalInput() *string
	State() *string
	StaticIpAddresses() *[]*string
	StaticIpEnabled() interface{}
	SetStaticIpEnabled(val interface{})
	StaticIpEnabledInput() interface{}
	SyncMode() *string
	SetSyncMode(val *string)
	SyncModeInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() DiscoveryEngineDataConnectorTimeoutsOutputReference
	TimeoutsInput() interface{}
	UpdateTime() *string
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
	PutActionConfig(value *DiscoveryEngineDataConnectorActionConfig)
	PutBapConfig(value *DiscoveryEngineDataConnectorBapConfig)
	PutDestinationConfigs(value interface{})
	PutEntities(value interface{})
	PutTimeouts(value *DiscoveryEngineDataConnectorTimeouts)
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
	ResetActionConfig()
	ResetAutoRunDisabled()
	ResetBapConfig()
	ResetConnectorModes()
	ResetDataSourceVersion()
	ResetDestinationConfigs()
	ResetEntities()
	ResetId()
	ResetIncrementalRefreshInterval()
	ResetIncrementalSyncDisabled()
	ResetJsonParams()
	ResetKmsKeyName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetParams()
	ResetProject()
	ResetStaticIpEnabled()
	ResetSyncMode()
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

// The jsii proxy struct for DiscoveryEngineDataConnector
type jsiiProxy_DiscoveryEngineDataConnector struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_DiscoveryEngineDataConnector) ActionConfig() DiscoveryEngineDataConnectorActionConfigOutputReference {
	var returns DiscoveryEngineDataConnectorActionConfigOutputReference
	_jsii_.Get(
		j,
		"actionConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataConnector) ActionConfigInput() *DiscoveryEngineDataConnectorActionConfig {
	var returns *DiscoveryEngineDataConnectorActionConfig
	_jsii_.Get(
		j,
		"actionConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataConnector) ActionState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataConnector) AutoRunDisabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoRunDisabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataConnector) AutoRunDisabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoRunDisabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataConnector) BapConfig() DiscoveryEngineDataConnectorBapConfigOutputReference {
	var returns DiscoveryEngineDataConnectorBapConfigOutputReference
	_jsii_.Get(
		j,
		"bapConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataConnector) BapConfigInput() *DiscoveryEngineDataConnectorBapConfig {
	var returns *DiscoveryEngineDataConnectorBapConfig
	_jsii_.Get(
		j,
		"bapConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataConnector) BlockingReasons() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"blockingReasons",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataConnector) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataConnector) CollectionDisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"collectionDisplayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataConnector) CollectionDisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"collectionDisplayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataConnector) CollectionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"collectionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataConnector) CollectionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"collectionIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataConnector) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataConnector) ConnectorModes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"connectorModes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataConnector) ConnectorModesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"connectorModesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataConnector) ConnectorType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectorType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataConnector) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataConnector) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataConnector) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataConnector) DataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataConnector) DataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataConnector) DataSourceVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataSourceVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataConnector) DataSourceVersionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataSourceVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataConnector) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataConnector) DestinationConfigs() DiscoveryEngineDataConnectorDestinationConfigsList {
	var returns DiscoveryEngineDataConnectorDestinationConfigsList
	_jsii_.Get(
		j,
		"destinationConfigs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataConnector) DestinationConfigsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"destinationConfigsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataConnector) Entities() DiscoveryEngineDataConnectorEntitiesList {
	var returns DiscoveryEngineDataConnectorEntitiesList
	_jsii_.Get(
		j,
		"entities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataConnector) EntitiesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"entitiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataConnector) Errors() DiscoveryEngineDataConnectorErrorsList {
	var returns DiscoveryEngineDataConnectorErrorsList
	_jsii_.Get(
		j,
		"errors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataConnector) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataConnector) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataConnector) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataConnector) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataConnector) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataConnector) IncrementalRefreshInterval() *string {
	var returns *string
	_jsii_.Get(
		j,
		"incrementalRefreshInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataConnector) IncrementalRefreshIntervalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"incrementalRefreshIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataConnector) IncrementalSyncDisabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"incrementalSyncDisabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataConnector) IncrementalSyncDisabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"incrementalSyncDisabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataConnector) JsonParams() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jsonParams",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataConnector) JsonParamsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jsonParamsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataConnector) KmsKeyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataConnector) KmsKeyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataConnector) LastSyncTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastSyncTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataConnector) LatestPauseTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"latestPauseTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataConnector) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataConnector) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataConnector) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataConnector) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataConnector) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataConnector) Params() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"params",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataConnector) ParamsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"paramsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataConnector) PrivateConnectivityProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateConnectivityProjectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataConnector) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataConnector) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataConnector) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataConnector) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataConnector) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataConnector) RealtimeState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"realtimeState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataConnector) RefreshInterval() *string {
	var returns *string
	_jsii_.Get(
		j,
		"refreshInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataConnector) RefreshIntervalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"refreshIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataConnector) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataConnector) StaticIpAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"staticIpAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataConnector) StaticIpEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"staticIpEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataConnector) StaticIpEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"staticIpEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataConnector) SyncMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"syncMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataConnector) SyncModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"syncModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataConnector) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataConnector) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataConnector) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataConnector) Timeouts() DiscoveryEngineDataConnectorTimeoutsOutputReference {
	var returns DiscoveryEngineDataConnectorTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataConnector) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataConnector) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/discovery_engine_data_connector google_discovery_engine_data_connector} Resource.
func NewDiscoveryEngineDataConnector(scope constructs.Construct, id *string, config *DiscoveryEngineDataConnectorConfig) DiscoveryEngineDataConnector {
	_init_.Initialize()

	if err := validateNewDiscoveryEngineDataConnectorParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DiscoveryEngineDataConnector{}

	_jsii_.Create(
		"@cdktn/provider-google.discoveryEngineDataConnector.DiscoveryEngineDataConnector",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/discovery_engine_data_connector google_discovery_engine_data_connector} Resource.
func NewDiscoveryEngineDataConnector_Override(d DiscoveryEngineDataConnector, scope constructs.Construct, id *string, config *DiscoveryEngineDataConnectorConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.discoveryEngineDataConnector.DiscoveryEngineDataConnector",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DiscoveryEngineDataConnector)SetAutoRunDisabled(val interface{}) {
	if err := j.validateSetAutoRunDisabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoRunDisabled",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineDataConnector)SetCollectionDisplayName(val *string) {
	if err := j.validateSetCollectionDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"collectionDisplayName",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineDataConnector)SetCollectionId(val *string) {
	if err := j.validateSetCollectionIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"collectionId",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineDataConnector)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineDataConnector)SetConnectorModes(val *[]*string) {
	if err := j.validateSetConnectorModesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectorModes",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineDataConnector)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineDataConnector)SetDataSource(val *string) {
	if err := j.validateSetDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataSource",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineDataConnector)SetDataSourceVersion(val *float64) {
	if err := j.validateSetDataSourceVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataSourceVersion",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineDataConnector)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineDataConnector)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineDataConnector)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineDataConnector)SetIncrementalRefreshInterval(val *string) {
	if err := j.validateSetIncrementalRefreshIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"incrementalRefreshInterval",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineDataConnector)SetIncrementalSyncDisabled(val interface{}) {
	if err := j.validateSetIncrementalSyncDisabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"incrementalSyncDisabled",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineDataConnector)SetJsonParams(val *string) {
	if err := j.validateSetJsonParamsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jsonParams",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineDataConnector)SetKmsKeyName(val *string) {
	if err := j.validateSetKmsKeyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyName",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineDataConnector)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineDataConnector)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineDataConnector)SetParams(val *map[string]*string) {
	if err := j.validateSetParamsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"params",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineDataConnector)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineDataConnector)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineDataConnector)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineDataConnector)SetRefreshInterval(val *string) {
	if err := j.validateSetRefreshIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"refreshInterval",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineDataConnector)SetStaticIpEnabled(val interface{}) {
	if err := j.validateSetStaticIpEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"staticIpEnabled",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineDataConnector)SetSyncMode(val *string) {
	if err := j.validateSetSyncModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"syncMode",
		val,
	)
}

// Generates CDKTN code for importing a DiscoveryEngineDataConnector resource upon running "cdktn plan <stack-name>".
func DiscoveryEngineDataConnector_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateDiscoveryEngineDataConnector_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.discoveryEngineDataConnector.DiscoveryEngineDataConnector",
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
func DiscoveryEngineDataConnector_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDiscoveryEngineDataConnector_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.discoveryEngineDataConnector.DiscoveryEngineDataConnector",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DiscoveryEngineDataConnector_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDiscoveryEngineDataConnector_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.discoveryEngineDataConnector.DiscoveryEngineDataConnector",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DiscoveryEngineDataConnector_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDiscoveryEngineDataConnector_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.discoveryEngineDataConnector.DiscoveryEngineDataConnector",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DiscoveryEngineDataConnector_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-google.discoveryEngineDataConnector.DiscoveryEngineDataConnector",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DiscoveryEngineDataConnector) AddMoveTarget(moveTarget *string) {
	if err := d.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (d *jsiiProxy_DiscoveryEngineDataConnector) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DiscoveryEngineDataConnector) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineDataConnector) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineDataConnector) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineDataConnector) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineDataConnector) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineDataConnector) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineDataConnector) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineDataConnector) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineDataConnector) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineDataConnector) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineDataConnector) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := d.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (d *jsiiProxy_DiscoveryEngineDataConnector) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineDataConnector) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := d.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineDataConnector) MoveFromId(id *string) {
	if err := d.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveFromId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DiscoveryEngineDataConnector) MoveTo(moveTarget *string, index interface{}) {
	if err := d.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (d *jsiiProxy_DiscoveryEngineDataConnector) MoveToId(id *string) {
	if err := d.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveToId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DiscoveryEngineDataConnector) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DiscoveryEngineDataConnector) PutActionConfig(value *DiscoveryEngineDataConnectorActionConfig) {
	if err := d.validatePutActionConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putActionConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DiscoveryEngineDataConnector) PutBapConfig(value *DiscoveryEngineDataConnectorBapConfig) {
	if err := d.validatePutBapConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putBapConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DiscoveryEngineDataConnector) PutDestinationConfigs(value interface{}) {
	if err := d.validatePutDestinationConfigsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putDestinationConfigs",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DiscoveryEngineDataConnector) PutEntities(value interface{}) {
	if err := d.validatePutEntitiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putEntities",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DiscoveryEngineDataConnector) PutTimeouts(value *DiscoveryEngineDataConnectorTimeouts) {
	if err := d.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DiscoveryEngineDataConnector) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := d.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (d *jsiiProxy_DiscoveryEngineDataConnector) ResetActionConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetActionConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineDataConnector) ResetAutoRunDisabled() {
	_jsii_.InvokeVoid(
		d,
		"resetAutoRunDisabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineDataConnector) ResetBapConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetBapConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineDataConnector) ResetConnectorModes() {
	_jsii_.InvokeVoid(
		d,
		"resetConnectorModes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineDataConnector) ResetDataSourceVersion() {
	_jsii_.InvokeVoid(
		d,
		"resetDataSourceVersion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineDataConnector) ResetDestinationConfigs() {
	_jsii_.InvokeVoid(
		d,
		"resetDestinationConfigs",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineDataConnector) ResetEntities() {
	_jsii_.InvokeVoid(
		d,
		"resetEntities",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineDataConnector) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineDataConnector) ResetIncrementalRefreshInterval() {
	_jsii_.InvokeVoid(
		d,
		"resetIncrementalRefreshInterval",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineDataConnector) ResetIncrementalSyncDisabled() {
	_jsii_.InvokeVoid(
		d,
		"resetIncrementalSyncDisabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineDataConnector) ResetJsonParams() {
	_jsii_.InvokeVoid(
		d,
		"resetJsonParams",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineDataConnector) ResetKmsKeyName() {
	_jsii_.InvokeVoid(
		d,
		"resetKmsKeyName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineDataConnector) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineDataConnector) ResetParams() {
	_jsii_.InvokeVoid(
		d,
		"resetParams",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineDataConnector) ResetProject() {
	_jsii_.InvokeVoid(
		d,
		"resetProject",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineDataConnector) ResetStaticIpEnabled() {
	_jsii_.InvokeVoid(
		d,
		"resetStaticIpEnabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineDataConnector) ResetSyncMode() {
	_jsii_.InvokeVoid(
		d,
		"resetSyncMode",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineDataConnector) ResetTimeouts() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineDataConnector) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineDataConnector) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineDataConnector) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineDataConnector) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineDataConnector) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineDataConnector) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineDataConnector) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		d,
		"with",
		args,
		&returns,
	)

	return returns
}

