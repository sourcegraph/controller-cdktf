package googlenetworkmanagementorganizationvpcflowlogsconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlenetworkmanagementorganizationvpcflowlogsconfig/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_network_management_organization_vpc_flow_logs_config google_network_management_organization_vpc_flow_logs_config}.
type GoogleNetworkManagementOrganizationVpcFlowLogsConfig interface {
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
	Timeouts() GoogleNetworkManagementOrganizationVpcFlowLogsConfigTimeoutsOutputReference
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
	PutTimeouts(value *GoogleNetworkManagementOrganizationVpcFlowLogsConfigTimeouts)
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

// The jsii proxy struct for GoogleNetworkManagementOrganizationVpcFlowLogsConfig
type jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig) AggregationInterval() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aggregationInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig) AggregationIntervalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aggregationIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig) CrossProjectMetadata() *string {
	var returns *string
	_jsii_.Get(
		j,
		"crossProjectMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig) CrossProjectMetadataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"crossProjectMetadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig) EffectiveLabels() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"effectiveLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig) FilterExpr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filterExpr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig) FilterExprInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filterExprInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig) FlowSampling() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"flowSampling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig) FlowSamplingInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"flowSamplingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig) Metadata() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig) MetadataFields() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"metadataFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig) MetadataFieldsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"metadataFieldsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig) MetadataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig) Organization() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig) OrganizationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig) StateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig) TerraformLabels() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"terraformLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig) Timeouts() GoogleNetworkManagementOrganizationVpcFlowLogsConfigTimeoutsOutputReference {
	var returns GoogleNetworkManagementOrganizationVpcFlowLogsConfigTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig) VpcFlowLogsConfigId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcFlowLogsConfigId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig) VpcFlowLogsConfigIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcFlowLogsConfigIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_network_management_organization_vpc_flow_logs_config google_network_management_organization_vpc_flow_logs_config} Resource.
func NewGoogleNetworkManagementOrganizationVpcFlowLogsConfig(scope constructs.Construct, id *string, config *GoogleNetworkManagementOrganizationVpcFlowLogsConfigConfig) GoogleNetworkManagementOrganizationVpcFlowLogsConfig {
	_init_.Initialize()

	if err := validateNewGoogleNetworkManagementOrganizationVpcFlowLogsConfigParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleNetworkManagementOrganizationVpcFlowLogsConfig.GoogleNetworkManagementOrganizationVpcFlowLogsConfig",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_network_management_organization_vpc_flow_logs_config google_network_management_organization_vpc_flow_logs_config} Resource.
func NewGoogleNetworkManagementOrganizationVpcFlowLogsConfig_Override(g GoogleNetworkManagementOrganizationVpcFlowLogsConfig, scope constructs.Construct, id *string, config *GoogleNetworkManagementOrganizationVpcFlowLogsConfigConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleNetworkManagementOrganizationVpcFlowLogsConfig.GoogleNetworkManagementOrganizationVpcFlowLogsConfig",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig)SetAggregationInterval(val *string) {
	if err := j.validateSetAggregationIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aggregationInterval",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig)SetCrossProjectMetadata(val *string) {
	if err := j.validateSetCrossProjectMetadataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"crossProjectMetadata",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig)SetFilterExpr(val *string) {
	if err := j.validateSetFilterExprParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"filterExpr",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig)SetFlowSampling(val *float64) {
	if err := j.validateSetFlowSamplingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"flowSampling",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig)SetMetadata(val *string) {
	if err := j.validateSetMetadataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metadata",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig)SetMetadataFields(val *[]*string) {
	if err := j.validateSetMetadataFieldsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metadataFields",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig)SetOrganization(val *string) {
	if err := j.validateSetOrganizationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"organization",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig)SetState(val *string) {
	if err := j.validateSetStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"state",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig)SetVpcFlowLogsConfigId(val *string) {
	if err := j.validateSetVpcFlowLogsConfigIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcFlowLogsConfigId",
		val,
	)
}

// Generates CDKTN code for importing a GoogleNetworkManagementOrganizationVpcFlowLogsConfig resource upon running "cdktn plan <stack-name>".
func GoogleNetworkManagementOrganizationVpcFlowLogsConfig_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateGoogleNetworkManagementOrganizationVpcFlowLogsConfig_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-google-beta.googleNetworkManagementOrganizationVpcFlowLogsConfig.GoogleNetworkManagementOrganizationVpcFlowLogsConfig",
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
func GoogleNetworkManagementOrganizationVpcFlowLogsConfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleNetworkManagementOrganizationVpcFlowLogsConfig_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google-beta.googleNetworkManagementOrganizationVpcFlowLogsConfig.GoogleNetworkManagementOrganizationVpcFlowLogsConfig",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleNetworkManagementOrganizationVpcFlowLogsConfig_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleNetworkManagementOrganizationVpcFlowLogsConfig_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google-beta.googleNetworkManagementOrganizationVpcFlowLogsConfig.GoogleNetworkManagementOrganizationVpcFlowLogsConfig",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleNetworkManagementOrganizationVpcFlowLogsConfig_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleNetworkManagementOrganizationVpcFlowLogsConfig_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google-beta.googleNetworkManagementOrganizationVpcFlowLogsConfig.GoogleNetworkManagementOrganizationVpcFlowLogsConfig",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleNetworkManagementOrganizationVpcFlowLogsConfig_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-google-beta.googleNetworkManagementOrganizationVpcFlowLogsConfig.GoogleNetworkManagementOrganizationVpcFlowLogsConfig",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (g *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig) PutTimeouts(value *GoogleNetworkManagementOrganizationVpcFlowLogsConfigTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := g.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (g *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig) ResetAggregationInterval() {
	_jsii_.InvokeVoid(
		g,
		"resetAggregationInterval",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig) ResetCrossProjectMetadata() {
	_jsii_.InvokeVoid(
		g,
		"resetCrossProjectMetadata",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig) ResetFilterExpr() {
	_jsii_.InvokeVoid(
		g,
		"resetFilterExpr",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig) ResetFlowSampling() {
	_jsii_.InvokeVoid(
		g,
		"resetFlowSampling",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig) ResetLabels() {
	_jsii_.InvokeVoid(
		g,
		"resetLabels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig) ResetMetadata() {
	_jsii_.InvokeVoid(
		g,
		"resetMetadata",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig) ResetMetadataFields() {
	_jsii_.InvokeVoid(
		g,
		"resetMetadataFields",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig) ResetState() {
	_jsii_.InvokeVoid(
		g,
		"resetState",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetworkManagementOrganizationVpcFlowLogsConfig) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

