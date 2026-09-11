package googlediscoveryenginedataconnector

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlediscoveryenginedataconnector/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_discovery_engine_data_connector google_discovery_engine_data_connector}.
type GoogleDiscoveryEngineDataConnector interface {
	cdktn.TerraformResource
	ActionConfig() GoogleDiscoveryEngineDataConnectorActionConfigOutputReference
	ActionConfigInput() *GoogleDiscoveryEngineDataConnectorActionConfig
	ActionState() *string
	AutoRunDisabled() interface{}
	SetAutoRunDisabled(val interface{})
	AutoRunDisabledInput() interface{}
	BapConfig() GoogleDiscoveryEngineDataConnectorBapConfigOutputReference
	BapConfigInput() *GoogleDiscoveryEngineDataConnectorBapConfig
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
	DestinationConfigs() GoogleDiscoveryEngineDataConnectorDestinationConfigsList
	DestinationConfigsInput() interface{}
	Entities() GoogleDiscoveryEngineDataConnectorEntitiesList
	EntitiesInput() interface{}
	Errors() GoogleDiscoveryEngineDataConnectorErrorsList
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
	Timeouts() GoogleDiscoveryEngineDataConnectorTimeoutsOutputReference
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
	PutActionConfig(value *GoogleDiscoveryEngineDataConnectorActionConfig)
	PutBapConfig(value *GoogleDiscoveryEngineDataConnectorBapConfig)
	PutDestinationConfigs(value interface{})
	PutEntities(value interface{})
	PutTimeouts(value *GoogleDiscoveryEngineDataConnectorTimeouts)
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

// The jsii proxy struct for GoogleDiscoveryEngineDataConnector
type jsiiProxy_GoogleDiscoveryEngineDataConnector struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnector) ActionConfig() GoogleDiscoveryEngineDataConnectorActionConfigOutputReference {
	var returns GoogleDiscoveryEngineDataConnectorActionConfigOutputReference
	_jsii_.Get(
		j,
		"actionConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnector) ActionConfigInput() *GoogleDiscoveryEngineDataConnectorActionConfig {
	var returns *GoogleDiscoveryEngineDataConnectorActionConfig
	_jsii_.Get(
		j,
		"actionConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnector) ActionState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnector) AutoRunDisabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoRunDisabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnector) AutoRunDisabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoRunDisabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnector) BapConfig() GoogleDiscoveryEngineDataConnectorBapConfigOutputReference {
	var returns GoogleDiscoveryEngineDataConnectorBapConfigOutputReference
	_jsii_.Get(
		j,
		"bapConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnector) BapConfigInput() *GoogleDiscoveryEngineDataConnectorBapConfig {
	var returns *GoogleDiscoveryEngineDataConnectorBapConfig
	_jsii_.Get(
		j,
		"bapConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnector) BlockingReasons() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"blockingReasons",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnector) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnector) CollectionDisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"collectionDisplayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnector) CollectionDisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"collectionDisplayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnector) CollectionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"collectionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnector) CollectionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"collectionIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnector) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnector) ConnectorModes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"connectorModes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnector) ConnectorModesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"connectorModesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnector) ConnectorType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectorType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnector) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnector) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnector) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnector) DataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnector) DataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnector) DataSourceVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataSourceVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnector) DataSourceVersionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataSourceVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnector) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnector) DestinationConfigs() GoogleDiscoveryEngineDataConnectorDestinationConfigsList {
	var returns GoogleDiscoveryEngineDataConnectorDestinationConfigsList
	_jsii_.Get(
		j,
		"destinationConfigs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnector) DestinationConfigsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"destinationConfigsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnector) Entities() GoogleDiscoveryEngineDataConnectorEntitiesList {
	var returns GoogleDiscoveryEngineDataConnectorEntitiesList
	_jsii_.Get(
		j,
		"entities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnector) EntitiesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"entitiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnector) Errors() GoogleDiscoveryEngineDataConnectorErrorsList {
	var returns GoogleDiscoveryEngineDataConnectorErrorsList
	_jsii_.Get(
		j,
		"errors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnector) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnector) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnector) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnector) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnector) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnector) IncrementalRefreshInterval() *string {
	var returns *string
	_jsii_.Get(
		j,
		"incrementalRefreshInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnector) IncrementalRefreshIntervalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"incrementalRefreshIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnector) IncrementalSyncDisabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"incrementalSyncDisabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnector) IncrementalSyncDisabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"incrementalSyncDisabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnector) JsonParams() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jsonParams",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnector) JsonParamsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jsonParamsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnector) KmsKeyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnector) KmsKeyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnector) LastSyncTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastSyncTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnector) LatestPauseTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"latestPauseTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnector) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnector) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnector) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnector) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnector) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnector) Params() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"params",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnector) ParamsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"paramsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnector) PrivateConnectivityProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateConnectivityProjectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnector) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnector) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnector) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnector) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnector) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnector) RealtimeState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"realtimeState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnector) RefreshInterval() *string {
	var returns *string
	_jsii_.Get(
		j,
		"refreshInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnector) RefreshIntervalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"refreshIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnector) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnector) StaticIpAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"staticIpAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnector) StaticIpEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"staticIpEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnector) StaticIpEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"staticIpEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnector) SyncMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"syncMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnector) SyncModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"syncModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnector) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnector) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnector) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnector) Timeouts() GoogleDiscoveryEngineDataConnectorTimeoutsOutputReference {
	var returns GoogleDiscoveryEngineDataConnectorTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnector) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnector) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_discovery_engine_data_connector google_discovery_engine_data_connector} Resource.
func NewGoogleDiscoveryEngineDataConnector(scope constructs.Construct, id *string, config *GoogleDiscoveryEngineDataConnectorConfig) GoogleDiscoveryEngineDataConnector {
	_init_.Initialize()

	if err := validateNewGoogleDiscoveryEngineDataConnectorParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDiscoveryEngineDataConnector{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleDiscoveryEngineDataConnector.GoogleDiscoveryEngineDataConnector",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_discovery_engine_data_connector google_discovery_engine_data_connector} Resource.
func NewGoogleDiscoveryEngineDataConnector_Override(g GoogleDiscoveryEngineDataConnector, scope constructs.Construct, id *string, config *GoogleDiscoveryEngineDataConnectorConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleDiscoveryEngineDataConnector.GoogleDiscoveryEngineDataConnector",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnector)SetAutoRunDisabled(val interface{}) {
	if err := j.validateSetAutoRunDisabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoRunDisabled",
		val,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnector)SetCollectionDisplayName(val *string) {
	if err := j.validateSetCollectionDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"collectionDisplayName",
		val,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnector)SetCollectionId(val *string) {
	if err := j.validateSetCollectionIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"collectionId",
		val,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnector)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnector)SetConnectorModes(val *[]*string) {
	if err := j.validateSetConnectorModesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectorModes",
		val,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnector)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnector)SetDataSource(val *string) {
	if err := j.validateSetDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataSource",
		val,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnector)SetDataSourceVersion(val *float64) {
	if err := j.validateSetDataSourceVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataSourceVersion",
		val,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnector)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnector)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnector)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnector)SetIncrementalRefreshInterval(val *string) {
	if err := j.validateSetIncrementalRefreshIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"incrementalRefreshInterval",
		val,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnector)SetIncrementalSyncDisabled(val interface{}) {
	if err := j.validateSetIncrementalSyncDisabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"incrementalSyncDisabled",
		val,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnector)SetJsonParams(val *string) {
	if err := j.validateSetJsonParamsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jsonParams",
		val,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnector)SetKmsKeyName(val *string) {
	if err := j.validateSetKmsKeyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyName",
		val,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnector)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnector)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnector)SetParams(val *map[string]*string) {
	if err := j.validateSetParamsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"params",
		val,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnector)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnector)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnector)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnector)SetRefreshInterval(val *string) {
	if err := j.validateSetRefreshIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"refreshInterval",
		val,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnector)SetStaticIpEnabled(val interface{}) {
	if err := j.validateSetStaticIpEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"staticIpEnabled",
		val,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnector)SetSyncMode(val *string) {
	if err := j.validateSetSyncModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"syncMode",
		val,
	)
}

// Generates CDKTN code for importing a GoogleDiscoveryEngineDataConnector resource upon running "cdktn plan <stack-name>".
func GoogleDiscoveryEngineDataConnector_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateGoogleDiscoveryEngineDataConnector_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-google-beta.googleDiscoveryEngineDataConnector.GoogleDiscoveryEngineDataConnector",
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
func GoogleDiscoveryEngineDataConnector_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleDiscoveryEngineDataConnector_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google-beta.googleDiscoveryEngineDataConnector.GoogleDiscoveryEngineDataConnector",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleDiscoveryEngineDataConnector_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleDiscoveryEngineDataConnector_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google-beta.googleDiscoveryEngineDataConnector.GoogleDiscoveryEngineDataConnector",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleDiscoveryEngineDataConnector_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleDiscoveryEngineDataConnector_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google-beta.googleDiscoveryEngineDataConnector.GoogleDiscoveryEngineDataConnector",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleDiscoveryEngineDataConnector_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-google-beta.googleDiscoveryEngineDataConnector.GoogleDiscoveryEngineDataConnector",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleDiscoveryEngineDataConnector) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineDataConnector) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineDataConnector) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDiscoveryEngineDataConnector) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDiscoveryEngineDataConnector) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDiscoveryEngineDataConnector) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDiscoveryEngineDataConnector) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDiscoveryEngineDataConnector) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDiscoveryEngineDataConnector) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDiscoveryEngineDataConnector) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDiscoveryEngineDataConnector) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDiscoveryEngineDataConnector) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDiscoveryEngineDataConnector) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineDataConnector) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDiscoveryEngineDataConnector) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := g.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDiscoveryEngineDataConnector) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineDataConnector) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineDataConnector) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineDataConnector) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineDataConnector) PutActionConfig(value *GoogleDiscoveryEngineDataConnectorActionConfig) {
	if err := g.validatePutActionConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putActionConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineDataConnector) PutBapConfig(value *GoogleDiscoveryEngineDataConnectorBapConfig) {
	if err := g.validatePutBapConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putBapConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineDataConnector) PutDestinationConfigs(value interface{}) {
	if err := g.validatePutDestinationConfigsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDestinationConfigs",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineDataConnector) PutEntities(value interface{}) {
	if err := g.validatePutEntitiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putEntities",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineDataConnector) PutTimeouts(value *GoogleDiscoveryEngineDataConnectorTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineDataConnector) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := g.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineDataConnector) ResetActionConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetActionConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineDataConnector) ResetAutoRunDisabled() {
	_jsii_.InvokeVoid(
		g,
		"resetAutoRunDisabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineDataConnector) ResetBapConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetBapConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineDataConnector) ResetConnectorModes() {
	_jsii_.InvokeVoid(
		g,
		"resetConnectorModes",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineDataConnector) ResetDataSourceVersion() {
	_jsii_.InvokeVoid(
		g,
		"resetDataSourceVersion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineDataConnector) ResetDestinationConfigs() {
	_jsii_.InvokeVoid(
		g,
		"resetDestinationConfigs",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineDataConnector) ResetEntities() {
	_jsii_.InvokeVoid(
		g,
		"resetEntities",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineDataConnector) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineDataConnector) ResetIncrementalRefreshInterval() {
	_jsii_.InvokeVoid(
		g,
		"resetIncrementalRefreshInterval",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineDataConnector) ResetIncrementalSyncDisabled() {
	_jsii_.InvokeVoid(
		g,
		"resetIncrementalSyncDisabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineDataConnector) ResetJsonParams() {
	_jsii_.InvokeVoid(
		g,
		"resetJsonParams",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineDataConnector) ResetKmsKeyName() {
	_jsii_.InvokeVoid(
		g,
		"resetKmsKeyName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineDataConnector) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineDataConnector) ResetParams() {
	_jsii_.InvokeVoid(
		g,
		"resetParams",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineDataConnector) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineDataConnector) ResetStaticIpEnabled() {
	_jsii_.InvokeVoid(
		g,
		"resetStaticIpEnabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineDataConnector) ResetSyncMode() {
	_jsii_.InvokeVoid(
		g,
		"resetSyncMode",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineDataConnector) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineDataConnector) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDiscoveryEngineDataConnector) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDiscoveryEngineDataConnector) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDiscoveryEngineDataConnector) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDiscoveryEngineDataConnector) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDiscoveryEngineDataConnector) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDiscoveryEngineDataConnector) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		g,
		"with",
		args,
		&returns,
	)

	return returns
}

