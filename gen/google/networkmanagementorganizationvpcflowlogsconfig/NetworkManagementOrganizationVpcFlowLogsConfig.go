package networkmanagementorganizationvpcflowlogsconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google/networkmanagementorganizationvpcflowlogsconfig/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/network_management_organization_vpc_flow_logs_config google_network_management_organization_vpc_flow_logs_config}.
type NetworkManagementOrganizationVpcFlowLogsConfig interface {
	cdktn.TerraformResource
	AggregationInterval() *string
	SetAggregationInterval(val *string)
	AggregationIntervalInput() *string
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
	CreateTime() *string
	CrossProjectMetadata() *string
	SetCrossProjectMetadata(val *string)
	CrossProjectMetadataInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	EffectiveLabels() cdktn.StringMap
	FilterExpr() *string
	SetFilterExpr(val *string)
	FilterExprInput() *string
	FlowSampling() *float64
	SetFlowSampling(val *float64)
	FlowSamplingInput() *float64
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
	Metadata() *string
	SetMetadata(val *string)
	MetadataFields() *[]*string
	SetMetadataFields(val *[]*string)
	MetadataFieldsInput() *[]*string
	MetadataInput() *string
	Name() *string
	// The tree node.
	Node() constructs.Node
	Organization() *string
	SetOrganization(val *string)
	OrganizationInput() *string
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
	State() *string
	SetState(val *string)
	StateInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	TerraformLabels() cdktn.StringMap
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() NetworkManagementOrganizationVpcFlowLogsConfigTimeoutsOutputReference
	TimeoutsInput() interface{}
	UpdateTime() *string
	VpcFlowLogsConfigId() *string
	SetVpcFlowLogsConfigId(val *string)
	VpcFlowLogsConfigIdInput() *string
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
	PutTimeouts(value *NetworkManagementOrganizationVpcFlowLogsConfigTimeouts)
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
	ResetAggregationInterval()
	ResetCrossProjectMetadata()
	ResetDescription()
	ResetFilterExpr()
	ResetFlowSampling()
	ResetId()
	ResetLabels()
	ResetMetadata()
	ResetMetadataFields()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetState()
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

// The jsii proxy struct for NetworkManagementOrganizationVpcFlowLogsConfig
type jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig) AggregationInterval() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aggregationInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig) AggregationIntervalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aggregationIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig) CrossProjectMetadata() *string {
	var returns *string
	_jsii_.Get(
		j,
		"crossProjectMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig) CrossProjectMetadataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"crossProjectMetadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig) EffectiveLabels() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"effectiveLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig) FilterExpr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filterExpr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig) FilterExprInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filterExprInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig) FlowSampling() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"flowSampling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig) FlowSamplingInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"flowSamplingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig) Metadata() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig) MetadataFields() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"metadataFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig) MetadataFieldsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"metadataFieldsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig) MetadataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig) Organization() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig) OrganizationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig) StateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig) TerraformLabels() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"terraformLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig) Timeouts() NetworkManagementOrganizationVpcFlowLogsConfigTimeoutsOutputReference {
	var returns NetworkManagementOrganizationVpcFlowLogsConfigTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig) VpcFlowLogsConfigId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcFlowLogsConfigId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig) VpcFlowLogsConfigIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcFlowLogsConfigIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/network_management_organization_vpc_flow_logs_config google_network_management_organization_vpc_flow_logs_config} Resource.
func NewNetworkManagementOrganizationVpcFlowLogsConfig(scope constructs.Construct, id *string, config *NetworkManagementOrganizationVpcFlowLogsConfigConfig) NetworkManagementOrganizationVpcFlowLogsConfig {
	_init_.Initialize()

	if err := validateNewNetworkManagementOrganizationVpcFlowLogsConfigParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig{}

	_jsii_.Create(
		"@cdktn/provider-google.networkManagementOrganizationVpcFlowLogsConfig.NetworkManagementOrganizationVpcFlowLogsConfig",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/network_management_organization_vpc_flow_logs_config google_network_management_organization_vpc_flow_logs_config} Resource.
func NewNetworkManagementOrganizationVpcFlowLogsConfig_Override(n NetworkManagementOrganizationVpcFlowLogsConfig, scope constructs.Construct, id *string, config *NetworkManagementOrganizationVpcFlowLogsConfigConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.networkManagementOrganizationVpcFlowLogsConfig.NetworkManagementOrganizationVpcFlowLogsConfig",
		[]interface{}{scope, id, config},
		n,
	)
}

func (j *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig)SetAggregationInterval(val *string) {
	if err := j.validateSetAggregationIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aggregationInterval",
		val,
	)
}

func (j *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig)SetCrossProjectMetadata(val *string) {
	if err := j.validateSetCrossProjectMetadataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"crossProjectMetadata",
		val,
	)
}

func (j *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig)SetFilterExpr(val *string) {
	if err := j.validateSetFilterExprParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"filterExpr",
		val,
	)
}

func (j *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig)SetFlowSampling(val *float64) {
	if err := j.validateSetFlowSamplingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"flowSampling",
		val,
	)
}

func (j *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig)SetMetadata(val *string) {
	if err := j.validateSetMetadataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metadata",
		val,
	)
}

func (j *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig)SetMetadataFields(val *[]*string) {
	if err := j.validateSetMetadataFieldsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metadataFields",
		val,
	)
}

func (j *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig)SetOrganization(val *string) {
	if err := j.validateSetOrganizationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"organization",
		val,
	)
}

func (j *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig)SetState(val *string) {
	if err := j.validateSetStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"state",
		val,
	)
}

func (j *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig)SetVpcFlowLogsConfigId(val *string) {
	if err := j.validateSetVpcFlowLogsConfigIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcFlowLogsConfigId",
		val,
	)
}

// Generates CDKTN code for importing a NetworkManagementOrganizationVpcFlowLogsConfig resource upon running "cdktn plan <stack-name>".
func NetworkManagementOrganizationVpcFlowLogsConfig_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateNetworkManagementOrganizationVpcFlowLogsConfig_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.networkManagementOrganizationVpcFlowLogsConfig.NetworkManagementOrganizationVpcFlowLogsConfig",
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
func NetworkManagementOrganizationVpcFlowLogsConfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNetworkManagementOrganizationVpcFlowLogsConfig_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.networkManagementOrganizationVpcFlowLogsConfig.NetworkManagementOrganizationVpcFlowLogsConfig",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func NetworkManagementOrganizationVpcFlowLogsConfig_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNetworkManagementOrganizationVpcFlowLogsConfig_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.networkManagementOrganizationVpcFlowLogsConfig.NetworkManagementOrganizationVpcFlowLogsConfig",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func NetworkManagementOrganizationVpcFlowLogsConfig_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNetworkManagementOrganizationVpcFlowLogsConfig_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.networkManagementOrganizationVpcFlowLogsConfig.NetworkManagementOrganizationVpcFlowLogsConfig",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func NetworkManagementOrganizationVpcFlowLogsConfig_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-google.networkManagementOrganizationVpcFlowLogsConfig.NetworkManagementOrganizationVpcFlowLogsConfig",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (n *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig) AddMoveTarget(moveTarget *string) {
	if err := n.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (n *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig) AddOverride(path *string, value interface{}) {
	if err := n.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (n *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := n.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := n.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		n,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := n.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		n,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := n.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := n.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		n,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := n.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		n,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := n.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		n,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig) GetStringAttribute(terraformAttribute *string) *string {
	if err := n.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		n,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := n.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		n,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := n.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (n *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := n.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := n.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		n,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig) MoveFromId(id *string) {
	if err := n.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"moveFromId",
		[]interface{}{id},
	)
}

func (n *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig) MoveTo(moveTarget *string, index interface{}) {
	if err := n.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (n *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig) MoveToId(id *string) {
	if err := n.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"moveToId",
		[]interface{}{id},
	)
}

func (n *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig) OverrideLogicalId(newLogicalId *string) {
	if err := n.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (n *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig) PutTimeouts(value *NetworkManagementOrganizationVpcFlowLogsConfigTimeouts) {
	if err := n.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := n.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (n *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig) ResetAggregationInterval() {
	_jsii_.InvokeVoid(
		n,
		"resetAggregationInterval",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig) ResetCrossProjectMetadata() {
	_jsii_.InvokeVoid(
		n,
		"resetCrossProjectMetadata",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig) ResetDescription() {
	_jsii_.InvokeVoid(
		n,
		"resetDescription",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig) ResetFilterExpr() {
	_jsii_.InvokeVoid(
		n,
		"resetFilterExpr",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig) ResetFlowSampling() {
	_jsii_.InvokeVoid(
		n,
		"resetFlowSampling",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig) ResetId() {
	_jsii_.InvokeVoid(
		n,
		"resetId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig) ResetLabels() {
	_jsii_.InvokeVoid(
		n,
		"resetLabels",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig) ResetMetadata() {
	_jsii_.InvokeVoid(
		n,
		"resetMetadata",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig) ResetMetadataFields() {
	_jsii_.InvokeVoid(
		n,
		"resetMetadataFields",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		n,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig) ResetState() {
	_jsii_.InvokeVoid(
		n,
		"resetState",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig) ResetTimeouts() {
	_jsii_.InvokeVoid(
		n,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkManagementOrganizationVpcFlowLogsConfig) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		n,
		"with",
		args,
		&returns,
	)

	return returns
}

