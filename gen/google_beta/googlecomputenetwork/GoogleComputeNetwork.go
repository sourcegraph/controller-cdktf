package googlecomputenetwork

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecomputenetwork/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_compute_network google_compute_network}.
type GoogleComputeNetwork interface {
	cdktn.TerraformResource
	AutoCreateSubnetworks() interface{}
	SetAutoCreateSubnetworks(val interface{})
	AutoCreateSubnetworksInput() interface{}
	BgpAlwaysCompareMed() interface{}
	SetBgpAlwaysCompareMed(val interface{})
	BgpAlwaysCompareMedInput() interface{}
	BgpBestPathSelectionMode() *string
	SetBgpBestPathSelectionMode(val *string)
	BgpBestPathSelectionModeInput() *string
	BgpInterRegionCost() *string
	SetBgpInterRegionCost(val *string)
	BgpInterRegionCostInput() *string
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
	DeleteBgpAlwaysCompareMed() interface{}
	SetDeleteBgpAlwaysCompareMed(val interface{})
	DeleteBgpAlwaysCompareMedInput() interface{}
	DeleteDefaultRoutesOnCreate() interface{}
	SetDeleteDefaultRoutesOnCreate(val interface{})
	DeleteDefaultRoutesOnCreateInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	EnableUlaInternalIpv6() interface{}
	SetEnableUlaInternalIpv6(val interface{})
	EnableUlaInternalIpv6Input() interface{}
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GatewayIpv4() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	InternalIpv6Range() *string
	SetInternalIpv6Range(val *string)
	InternalIpv6RangeInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	Mtu() *float64
	SetMtu(val *float64)
	MtuInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	NetworkFirewallPolicyEnforcementOrder() *string
	SetNetworkFirewallPolicyEnforcementOrder(val *string)
	NetworkFirewallPolicyEnforcementOrderInput() *string
	NetworkId() *string
	NetworkProfile() *string
	SetNetworkProfile(val *string)
	NetworkProfileInput() *string
	// The tree node.
	Node() constructs.Node
	NumericId() *string
	Params() GoogleComputeNetworkParamsOutputReference
	ParamsInput() *GoogleComputeNetworkParams
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
	RoutingMode() *string
	SetRoutingMode(val *string)
	RoutingModeInput() *string
	SelfLink() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleComputeNetworkTimeoutsOutputReference
	TimeoutsInput() interface{}
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
	PutParams(value *GoogleComputeNetworkParams)
	PutTimeouts(value *GoogleComputeNetworkTimeouts)
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
	ResetAutoCreateSubnetworks()
	ResetBgpAlwaysCompareMed()
	ResetBgpBestPathSelectionMode()
	ResetBgpInterRegionCost()
	ResetDeleteBgpAlwaysCompareMed()
	ResetDeleteDefaultRoutesOnCreate()
	ResetDescription()
	ResetEnableUlaInternalIpv6()
	ResetId()
	ResetInternalIpv6Range()
	ResetMtu()
	ResetNetworkFirewallPolicyEnforcementOrder()
	ResetNetworkProfile()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetParams()
	ResetProject()
	ResetRoutingMode()
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

// The jsii proxy struct for GoogleComputeNetwork
type jsiiProxy_GoogleComputeNetwork struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_GoogleComputeNetwork) AutoCreateSubnetworks() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoCreateSubnetworks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetwork) AutoCreateSubnetworksInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoCreateSubnetworksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetwork) BgpAlwaysCompareMed() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bgpAlwaysCompareMed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetwork) BgpAlwaysCompareMedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bgpAlwaysCompareMedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetwork) BgpBestPathSelectionMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bgpBestPathSelectionMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetwork) BgpBestPathSelectionModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bgpBestPathSelectionModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetwork) BgpInterRegionCost() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bgpInterRegionCost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetwork) BgpInterRegionCostInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bgpInterRegionCostInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetwork) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetwork) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetwork) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetwork) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetwork) DeleteBgpAlwaysCompareMed() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteBgpAlwaysCompareMed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetwork) DeleteBgpAlwaysCompareMedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteBgpAlwaysCompareMedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetwork) DeleteDefaultRoutesOnCreate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteDefaultRoutesOnCreate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetwork) DeleteDefaultRoutesOnCreateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteDefaultRoutesOnCreateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetwork) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetwork) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetwork) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetwork) EnableUlaInternalIpv6() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableUlaInternalIpv6",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetwork) EnableUlaInternalIpv6Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableUlaInternalIpv6Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetwork) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetwork) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetwork) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetwork) GatewayIpv4() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayIpv4",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetwork) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetwork) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetwork) InternalIpv6Range() *string {
	var returns *string
	_jsii_.Get(
		j,
		"internalIpv6Range",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetwork) InternalIpv6RangeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"internalIpv6RangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetwork) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetwork) Mtu() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"mtu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetwork) MtuInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"mtuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetwork) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetwork) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetwork) NetworkFirewallPolicyEnforcementOrder() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkFirewallPolicyEnforcementOrder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetwork) NetworkFirewallPolicyEnforcementOrderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkFirewallPolicyEnforcementOrderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetwork) NetworkId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetwork) NetworkProfile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetwork) NetworkProfileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetwork) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetwork) NumericId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"numericId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetwork) Params() GoogleComputeNetworkParamsOutputReference {
	var returns GoogleComputeNetworkParamsOutputReference
	_jsii_.Get(
		j,
		"params",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetwork) ParamsInput() *GoogleComputeNetworkParams {
	var returns *GoogleComputeNetworkParams
	_jsii_.Get(
		j,
		"paramsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetwork) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetwork) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetwork) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetwork) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetwork) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetwork) RoutingMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetwork) RoutingModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetwork) SelfLink() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selfLink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetwork) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetwork) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetwork) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetwork) Timeouts() GoogleComputeNetworkTimeoutsOutputReference {
	var returns GoogleComputeNetworkTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetwork) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_compute_network google_compute_network} Resource.
func NewGoogleComputeNetwork(scope constructs.Construct, id *string, config *GoogleComputeNetworkConfig) GoogleComputeNetwork {
	_init_.Initialize()

	if err := validateNewGoogleComputeNetworkParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeNetwork{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleComputeNetwork.GoogleComputeNetwork",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_compute_network google_compute_network} Resource.
func NewGoogleComputeNetwork_Override(g GoogleComputeNetwork, scope constructs.Construct, id *string, config *GoogleComputeNetworkConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleComputeNetwork.GoogleComputeNetwork",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeNetwork)SetAutoCreateSubnetworks(val interface{}) {
	if err := j.validateSetAutoCreateSubnetworksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoCreateSubnetworks",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeNetwork)SetBgpAlwaysCompareMed(val interface{}) {
	if err := j.validateSetBgpAlwaysCompareMedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bgpAlwaysCompareMed",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeNetwork)SetBgpBestPathSelectionMode(val *string) {
	if err := j.validateSetBgpBestPathSelectionModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bgpBestPathSelectionMode",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeNetwork)SetBgpInterRegionCost(val *string) {
	if err := j.validateSetBgpInterRegionCostParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bgpInterRegionCost",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeNetwork)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeNetwork)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeNetwork)SetDeleteBgpAlwaysCompareMed(val interface{}) {
	if err := j.validateSetDeleteBgpAlwaysCompareMedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteBgpAlwaysCompareMed",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeNetwork)SetDeleteDefaultRoutesOnCreate(val interface{}) {
	if err := j.validateSetDeleteDefaultRoutesOnCreateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteDefaultRoutesOnCreate",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeNetwork)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeNetwork)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeNetwork)SetEnableUlaInternalIpv6(val interface{}) {
	if err := j.validateSetEnableUlaInternalIpv6Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableUlaInternalIpv6",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeNetwork)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeNetwork)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeNetwork)SetInternalIpv6Range(val *string) {
	if err := j.validateSetInternalIpv6RangeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalIpv6Range",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeNetwork)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeNetwork)SetMtu(val *float64) {
	if err := j.validateSetMtuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mtu",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeNetwork)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeNetwork)SetNetworkFirewallPolicyEnforcementOrder(val *string) {
	if err := j.validateSetNetworkFirewallPolicyEnforcementOrderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkFirewallPolicyEnforcementOrder",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeNetwork)SetNetworkProfile(val *string) {
	if err := j.validateSetNetworkProfileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkProfile",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeNetwork)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeNetwork)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeNetwork)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeNetwork)SetRoutingMode(val *string) {
	if err := j.validateSetRoutingModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routingMode",
		val,
	)
}

// Generates CDKTN code for importing a GoogleComputeNetwork resource upon running "cdktn plan <stack-name>".
func GoogleComputeNetwork_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateGoogleComputeNetwork_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-google-beta.googleComputeNetwork.GoogleComputeNetwork",
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
func GoogleComputeNetwork_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleComputeNetwork_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google-beta.googleComputeNetwork.GoogleComputeNetwork",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleComputeNetwork_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleComputeNetwork_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google-beta.googleComputeNetwork.GoogleComputeNetwork",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleComputeNetwork_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleComputeNetwork_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google-beta.googleComputeNetwork.GoogleComputeNetwork",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleComputeNetwork_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-google-beta.googleComputeNetwork.GoogleComputeNetwork",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleComputeNetwork) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GoogleComputeNetwork) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleComputeNetwork) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleComputeNetwork) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleComputeNetwork) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComputeNetwork) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComputeNetwork) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComputeNetwork) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComputeNetwork) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComputeNetwork) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComputeNetwork) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComputeNetwork) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeNetwork) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GoogleComputeNetwork) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleComputeNetwork) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (g *jsiiProxy_GoogleComputeNetwork) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleComputeNetwork) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GoogleComputeNetwork) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleComputeNetwork) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleComputeNetwork) PutParams(value *GoogleComputeNetworkParams) {
	if err := g.validatePutParamsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putParams",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeNetwork) PutTimeouts(value *GoogleComputeNetworkTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeNetwork) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := g.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (g *jsiiProxy_GoogleComputeNetwork) ResetAutoCreateSubnetworks() {
	_jsii_.InvokeVoid(
		g,
		"resetAutoCreateSubnetworks",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeNetwork) ResetBgpAlwaysCompareMed() {
	_jsii_.InvokeVoid(
		g,
		"resetBgpAlwaysCompareMed",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeNetwork) ResetBgpBestPathSelectionMode() {
	_jsii_.InvokeVoid(
		g,
		"resetBgpBestPathSelectionMode",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeNetwork) ResetBgpInterRegionCost() {
	_jsii_.InvokeVoid(
		g,
		"resetBgpInterRegionCost",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeNetwork) ResetDeleteBgpAlwaysCompareMed() {
	_jsii_.InvokeVoid(
		g,
		"resetDeleteBgpAlwaysCompareMed",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeNetwork) ResetDeleteDefaultRoutesOnCreate() {
	_jsii_.InvokeVoid(
		g,
		"resetDeleteDefaultRoutesOnCreate",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeNetwork) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeNetwork) ResetEnableUlaInternalIpv6() {
	_jsii_.InvokeVoid(
		g,
		"resetEnableUlaInternalIpv6",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeNetwork) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeNetwork) ResetInternalIpv6Range() {
	_jsii_.InvokeVoid(
		g,
		"resetInternalIpv6Range",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeNetwork) ResetMtu() {
	_jsii_.InvokeVoid(
		g,
		"resetMtu",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeNetwork) ResetNetworkFirewallPolicyEnforcementOrder() {
	_jsii_.InvokeVoid(
		g,
		"resetNetworkFirewallPolicyEnforcementOrder",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeNetwork) ResetNetworkProfile() {
	_jsii_.InvokeVoid(
		g,
		"resetNetworkProfile",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeNetwork) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeNetwork) ResetParams() {
	_jsii_.InvokeVoid(
		g,
		"resetParams",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeNetwork) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeNetwork) ResetRoutingMode() {
	_jsii_.InvokeVoid(
		g,
		"resetRoutingMode",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeNetwork) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeNetwork) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeNetwork) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeNetwork) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeNetwork) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeNetwork) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeNetwork) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeNetwork) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

