package googlecomputefirewall

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecomputefirewall/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_compute_firewall google_compute_firewall}.
type GoogleComputeFirewall interface {
	cdktn.TerraformResource
	Allow() GoogleComputeFirewallAllowList
	AllowInput() interface{}
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
	CreationTimestamp() *string
	Deny() GoogleComputeFirewallDenyList
	DenyInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DestinationRanges() *[]*string
	SetDestinationRanges(val *[]*string)
	DestinationRangesInput() *[]*string
	Direction() *string
	SetDirection(val *string)
	DirectionInput() *string
	Disabled() interface{}
	SetDisabled(val interface{})
	DisabledInput() interface{}
	EnableLogging() interface{}
	SetEnableLogging(val interface{})
	EnableLoggingInput() interface{}
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
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	LogConfig() GoogleComputeFirewallLogConfigOutputReference
	LogConfigInput() *GoogleComputeFirewallLogConfig
	Name() *string
	SetName(val *string)
	NameInput() *string
	Network() *string
	SetNetwork(val *string)
	NetworkInput() *string
	// The tree node.
	Node() constructs.Node
	Params() GoogleComputeFirewallParamsOutputReference
	ParamsInput() *GoogleComputeFirewallParams
	Priority() *float64
	SetPriority(val *float64)
	PriorityInput() *float64
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
	SelfLink() *string
	SourceRanges() *[]*string
	SetSourceRanges(val *[]*string)
	SourceRangesInput() *[]*string
	SourceServiceAccounts() *[]*string
	SetSourceServiceAccounts(val *[]*string)
	SourceServiceAccountsInput() *[]*string
	SourceTags() *[]*string
	SetSourceTags(val *[]*string)
	SourceTagsInput() *[]*string
	TargetServiceAccounts() *[]*string
	SetTargetServiceAccounts(val *[]*string)
	TargetServiceAccountsInput() *[]*string
	TargetTags() *[]*string
	SetTargetTags(val *[]*string)
	TargetTagsInput() *[]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleComputeFirewallTimeoutsOutputReference
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
	PutAllow(value interface{})
	PutDeny(value interface{})
	PutLogConfig(value *GoogleComputeFirewallLogConfig)
	PutParams(value *GoogleComputeFirewallParams)
	PutTimeouts(value *GoogleComputeFirewallTimeouts)
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
	ResetAllow()
	ResetDeny()
	ResetDescription()
	ResetDestinationRanges()
	ResetDirection()
	ResetDisabled()
	ResetEnableLogging()
	ResetId()
	ResetLogConfig()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetParams()
	ResetPriority()
	ResetProject()
	ResetSourceRanges()
	ResetSourceServiceAccounts()
	ResetSourceTags()
	ResetTargetServiceAccounts()
	ResetTargetTags()
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

// The jsii proxy struct for GoogleComputeFirewall
type jsiiProxy_GoogleComputeFirewall struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_GoogleComputeFirewall) Allow() GoogleComputeFirewallAllowList {
	var returns GoogleComputeFirewallAllowList
	_jsii_.Get(
		j,
		"allow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFirewall) AllowInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFirewall) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFirewall) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFirewall) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFirewall) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFirewall) CreationTimestamp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creationTimestamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFirewall) Deny() GoogleComputeFirewallDenyList {
	var returns GoogleComputeFirewallDenyList
	_jsii_.Get(
		j,
		"deny",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFirewall) DenyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"denyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFirewall) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFirewall) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFirewall) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFirewall) DestinationRanges() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFirewall) DestinationRangesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFirewall) Direction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"direction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFirewall) DirectionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFirewall) Disabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFirewall) DisabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFirewall) EnableLogging() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableLogging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFirewall) EnableLoggingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableLoggingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFirewall) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFirewall) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFirewall) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFirewall) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFirewall) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFirewall) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFirewall) LogConfig() GoogleComputeFirewallLogConfigOutputReference {
	var returns GoogleComputeFirewallLogConfigOutputReference
	_jsii_.Get(
		j,
		"logConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFirewall) LogConfigInput() *GoogleComputeFirewallLogConfig {
	var returns *GoogleComputeFirewallLogConfig
	_jsii_.Get(
		j,
		"logConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFirewall) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFirewall) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFirewall) Network() *string {
	var returns *string
	_jsii_.Get(
		j,
		"network",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFirewall) NetworkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFirewall) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFirewall) Params() GoogleComputeFirewallParamsOutputReference {
	var returns GoogleComputeFirewallParamsOutputReference
	_jsii_.Get(
		j,
		"params",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFirewall) ParamsInput() *GoogleComputeFirewallParams {
	var returns *GoogleComputeFirewallParams
	_jsii_.Get(
		j,
		"paramsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFirewall) Priority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFirewall) PriorityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFirewall) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFirewall) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFirewall) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFirewall) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFirewall) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFirewall) SelfLink() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selfLink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFirewall) SourceRanges() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourceRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFirewall) SourceRangesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourceRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFirewall) SourceServiceAccounts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourceServiceAccounts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFirewall) SourceServiceAccountsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourceServiceAccountsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFirewall) SourceTags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourceTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFirewall) SourceTagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourceTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFirewall) TargetServiceAccounts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetServiceAccounts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFirewall) TargetServiceAccountsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetServiceAccountsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFirewall) TargetTags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFirewall) TargetTagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFirewall) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFirewall) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFirewall) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFirewall) Timeouts() GoogleComputeFirewallTimeoutsOutputReference {
	var returns GoogleComputeFirewallTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFirewall) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_compute_firewall google_compute_firewall} Resource.
func NewGoogleComputeFirewall(scope constructs.Construct, id *string, config *GoogleComputeFirewallConfig) GoogleComputeFirewall {
	_init_.Initialize()

	if err := validateNewGoogleComputeFirewallParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeFirewall{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleComputeFirewall.GoogleComputeFirewall",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_compute_firewall google_compute_firewall} Resource.
func NewGoogleComputeFirewall_Override(g GoogleComputeFirewall, scope constructs.Construct, id *string, config *GoogleComputeFirewallConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleComputeFirewall.GoogleComputeFirewall",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeFirewall)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeFirewall)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeFirewall)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeFirewall)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeFirewall)SetDestinationRanges(val *[]*string) {
	if err := j.validateSetDestinationRangesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationRanges",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeFirewall)SetDirection(val *string) {
	if err := j.validateSetDirectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"direction",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeFirewall)SetDisabled(val interface{}) {
	if err := j.validateSetDisabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disabled",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeFirewall)SetEnableLogging(val interface{}) {
	if err := j.validateSetEnableLoggingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableLogging",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeFirewall)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeFirewall)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeFirewall)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeFirewall)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeFirewall)SetNetwork(val *string) {
	if err := j.validateSetNetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"network",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeFirewall)SetPriority(val *float64) {
	if err := j.validateSetPriorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"priority",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeFirewall)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeFirewall)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeFirewall)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeFirewall)SetSourceRanges(val *[]*string) {
	if err := j.validateSetSourceRangesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceRanges",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeFirewall)SetSourceServiceAccounts(val *[]*string) {
	if err := j.validateSetSourceServiceAccountsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceServiceAccounts",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeFirewall)SetSourceTags(val *[]*string) {
	if err := j.validateSetSourceTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceTags",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeFirewall)SetTargetServiceAccounts(val *[]*string) {
	if err := j.validateSetTargetServiceAccountsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetServiceAccounts",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeFirewall)SetTargetTags(val *[]*string) {
	if err := j.validateSetTargetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetTags",
		val,
	)
}

// Generates CDKTN code for importing a GoogleComputeFirewall resource upon running "cdktn plan <stack-name>".
func GoogleComputeFirewall_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateGoogleComputeFirewall_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-google-beta.googleComputeFirewall.GoogleComputeFirewall",
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
func GoogleComputeFirewall_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleComputeFirewall_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google-beta.googleComputeFirewall.GoogleComputeFirewall",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleComputeFirewall_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleComputeFirewall_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google-beta.googleComputeFirewall.GoogleComputeFirewall",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleComputeFirewall_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleComputeFirewall_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google-beta.googleComputeFirewall.GoogleComputeFirewall",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleComputeFirewall_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-google-beta.googleComputeFirewall.GoogleComputeFirewall",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleComputeFirewall) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GoogleComputeFirewall) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleComputeFirewall) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleComputeFirewall) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleComputeFirewall) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComputeFirewall) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComputeFirewall) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComputeFirewall) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComputeFirewall) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComputeFirewall) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComputeFirewall) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComputeFirewall) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeFirewall) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GoogleComputeFirewall) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleComputeFirewall) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (g *jsiiProxy_GoogleComputeFirewall) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleComputeFirewall) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GoogleComputeFirewall) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleComputeFirewall) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleComputeFirewall) PutAllow(value interface{}) {
	if err := g.validatePutAllowParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAllow",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeFirewall) PutDeny(value interface{}) {
	if err := g.validatePutDenyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDeny",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeFirewall) PutLogConfig(value *GoogleComputeFirewallLogConfig) {
	if err := g.validatePutLogConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putLogConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeFirewall) PutParams(value *GoogleComputeFirewallParams) {
	if err := g.validatePutParamsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putParams",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeFirewall) PutTimeouts(value *GoogleComputeFirewallTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeFirewall) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := g.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (g *jsiiProxy_GoogleComputeFirewall) ResetAllow() {
	_jsii_.InvokeVoid(
		g,
		"resetAllow",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeFirewall) ResetDeny() {
	_jsii_.InvokeVoid(
		g,
		"resetDeny",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeFirewall) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeFirewall) ResetDestinationRanges() {
	_jsii_.InvokeVoid(
		g,
		"resetDestinationRanges",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeFirewall) ResetDirection() {
	_jsii_.InvokeVoid(
		g,
		"resetDirection",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeFirewall) ResetDisabled() {
	_jsii_.InvokeVoid(
		g,
		"resetDisabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeFirewall) ResetEnableLogging() {
	_jsii_.InvokeVoid(
		g,
		"resetEnableLogging",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeFirewall) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeFirewall) ResetLogConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetLogConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeFirewall) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeFirewall) ResetParams() {
	_jsii_.InvokeVoid(
		g,
		"resetParams",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeFirewall) ResetPriority() {
	_jsii_.InvokeVoid(
		g,
		"resetPriority",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeFirewall) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeFirewall) ResetSourceRanges() {
	_jsii_.InvokeVoid(
		g,
		"resetSourceRanges",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeFirewall) ResetSourceServiceAccounts() {
	_jsii_.InvokeVoid(
		g,
		"resetSourceServiceAccounts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeFirewall) ResetSourceTags() {
	_jsii_.InvokeVoid(
		g,
		"resetSourceTags",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeFirewall) ResetTargetServiceAccounts() {
	_jsii_.InvokeVoid(
		g,
		"resetTargetServiceAccounts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeFirewall) ResetTargetTags() {
	_jsii_.InvokeVoid(
		g,
		"resetTargetTags",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeFirewall) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeFirewall) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeFirewall) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeFirewall) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeFirewall) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeFirewall) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeFirewall) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeFirewall) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

