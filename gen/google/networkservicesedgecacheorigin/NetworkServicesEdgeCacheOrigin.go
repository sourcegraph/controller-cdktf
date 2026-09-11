package networkservicesedgecacheorigin

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google/networkservicesedgecacheorigin/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/network_services_edge_cache_origin google_network_services_edge_cache_origin}.
type NetworkServicesEdgeCacheOrigin interface {
	cdktn.TerraformResource
	AwsV4Authentication() NetworkServicesEdgeCacheOriginAwsV4AuthenticationOutputReference
	AwsV4AuthenticationInput() *NetworkServicesEdgeCacheOriginAwsV4Authentication
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
	EffectiveLabels() cdktn.StringMap
	FailoverOrigin() *string
	SetFailoverOrigin(val *string)
	FailoverOriginInput() *string
	FlexShielding() NetworkServicesEdgeCacheOriginFlexShieldingOutputReference
	FlexShieldingInput() *NetworkServicesEdgeCacheOriginFlexShielding
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
	MaxAttempts() *float64
	SetMaxAttempts(val *float64)
	MaxAttemptsInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	OriginAddress() *string
	SetOriginAddress(val *string)
	OriginAddressInput() *string
	OriginOverrideAction() NetworkServicesEdgeCacheOriginOriginOverrideActionOutputReference
	OriginOverrideActionInput() *NetworkServicesEdgeCacheOriginOriginOverrideAction
	OriginRedirect() NetworkServicesEdgeCacheOriginOriginRedirectOutputReference
	OriginRedirectInput() *NetworkServicesEdgeCacheOriginOriginRedirect
	Port() *float64
	SetPort(val *float64)
	PortInput() *float64
	Project() *string
	SetProject(val *string)
	ProjectInput() *string
	Protocol() *string
	SetProtocol(val *string)
	ProtocolInput() *string
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
	RetryConditions() *[]*string
	SetRetryConditions(val *[]*string)
	RetryConditionsInput() *[]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	TerraformLabels() cdktn.StringMap
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeout() NetworkServicesEdgeCacheOriginTimeoutOutputReference
	TimeoutInput() *NetworkServicesEdgeCacheOriginTimeout
	Timeouts() NetworkServicesEdgeCacheOriginTimeoutsOutputReference
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
	PutAwsV4Authentication(value *NetworkServicesEdgeCacheOriginAwsV4Authentication)
	PutFlexShielding(value *NetworkServicesEdgeCacheOriginFlexShielding)
	PutOriginOverrideAction(value *NetworkServicesEdgeCacheOriginOriginOverrideAction)
	PutOriginRedirect(value *NetworkServicesEdgeCacheOriginOriginRedirect)
	PutTimeout(value *NetworkServicesEdgeCacheOriginTimeout)
	PutTimeouts(value *NetworkServicesEdgeCacheOriginTimeouts)
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
	ResetAwsV4Authentication()
	ResetDescription()
	ResetFailoverOrigin()
	ResetFlexShielding()
	ResetId()
	ResetLabels()
	ResetMaxAttempts()
	ResetOriginOverrideAction()
	ResetOriginRedirect()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPort()
	ResetProject()
	ResetProtocol()
	ResetRetryConditions()
	ResetTimeout()
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

// The jsii proxy struct for NetworkServicesEdgeCacheOrigin
type jsiiProxy_NetworkServicesEdgeCacheOrigin struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) AwsV4Authentication() NetworkServicesEdgeCacheOriginAwsV4AuthenticationOutputReference {
	var returns NetworkServicesEdgeCacheOriginAwsV4AuthenticationOutputReference
	_jsii_.Get(
		j,
		"awsV4Authentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) AwsV4AuthenticationInput() *NetworkServicesEdgeCacheOriginAwsV4Authentication {
	var returns *NetworkServicesEdgeCacheOriginAwsV4Authentication
	_jsii_.Get(
		j,
		"awsV4AuthenticationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) EffectiveLabels() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"effectiveLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) FailoverOrigin() *string {
	var returns *string
	_jsii_.Get(
		j,
		"failoverOrigin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) FailoverOriginInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"failoverOriginInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) FlexShielding() NetworkServicesEdgeCacheOriginFlexShieldingOutputReference {
	var returns NetworkServicesEdgeCacheOriginFlexShieldingOutputReference
	_jsii_.Get(
		j,
		"flexShielding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) FlexShieldingInput() *NetworkServicesEdgeCacheOriginFlexShielding {
	var returns *NetworkServicesEdgeCacheOriginFlexShielding
	_jsii_.Get(
		j,
		"flexShieldingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) MaxAttempts() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAttempts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) MaxAttemptsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAttemptsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) OriginAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) OriginAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) OriginOverrideAction() NetworkServicesEdgeCacheOriginOriginOverrideActionOutputReference {
	var returns NetworkServicesEdgeCacheOriginOriginOverrideActionOutputReference
	_jsii_.Get(
		j,
		"originOverrideAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) OriginOverrideActionInput() *NetworkServicesEdgeCacheOriginOriginOverrideAction {
	var returns *NetworkServicesEdgeCacheOriginOriginOverrideAction
	_jsii_.Get(
		j,
		"originOverrideActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) OriginRedirect() NetworkServicesEdgeCacheOriginOriginRedirectOutputReference {
	var returns NetworkServicesEdgeCacheOriginOriginRedirectOutputReference
	_jsii_.Get(
		j,
		"originRedirect",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) OriginRedirectInput() *NetworkServicesEdgeCacheOriginOriginRedirect {
	var returns *NetworkServicesEdgeCacheOriginOriginRedirect
	_jsii_.Get(
		j,
		"originRedirectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) ProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) RetryConditions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"retryConditions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) RetryConditionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"retryConditionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) TerraformLabels() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"terraformLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) Timeout() NetworkServicesEdgeCacheOriginTimeoutOutputReference {
	var returns NetworkServicesEdgeCacheOriginTimeoutOutputReference
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) TimeoutInput() *NetworkServicesEdgeCacheOriginTimeout {
	var returns *NetworkServicesEdgeCacheOriginTimeout
	_jsii_.Get(
		j,
		"timeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) Timeouts() NetworkServicesEdgeCacheOriginTimeoutsOutputReference {
	var returns NetworkServicesEdgeCacheOriginTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/network_services_edge_cache_origin google_network_services_edge_cache_origin} Resource.
func NewNetworkServicesEdgeCacheOrigin(scope constructs.Construct, id *string, config *NetworkServicesEdgeCacheOriginConfig) NetworkServicesEdgeCacheOrigin {
	_init_.Initialize()

	if err := validateNewNetworkServicesEdgeCacheOriginParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetworkServicesEdgeCacheOrigin{}

	_jsii_.Create(
		"@cdktn/provider-google.networkServicesEdgeCacheOrigin.NetworkServicesEdgeCacheOrigin",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/network_services_edge_cache_origin google_network_services_edge_cache_origin} Resource.
func NewNetworkServicesEdgeCacheOrigin_Override(n NetworkServicesEdgeCacheOrigin, scope constructs.Construct, id *string, config *NetworkServicesEdgeCacheOriginConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.networkServicesEdgeCacheOrigin.NetworkServicesEdgeCacheOrigin",
		[]interface{}{scope, id, config},
		n,
	)
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin)SetFailoverOrigin(val *string) {
	if err := j.validateSetFailoverOriginParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"failoverOrigin",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin)SetMaxAttempts(val *float64) {
	if err := j.validateSetMaxAttemptsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxAttempts",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin)SetOriginAddress(val *string) {
	if err := j.validateSetOriginAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"originAddress",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin)SetPort(val *float64) {
	if err := j.validateSetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin)SetProtocol(val *string) {
	if err := j.validateSetProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesEdgeCacheOrigin)SetRetryConditions(val *[]*string) {
	if err := j.validateSetRetryConditionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retryConditions",
		val,
	)
}

// Generates CDKTN code for importing a NetworkServicesEdgeCacheOrigin resource upon running "cdktn plan <stack-name>".
func NetworkServicesEdgeCacheOrigin_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateNetworkServicesEdgeCacheOrigin_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.networkServicesEdgeCacheOrigin.NetworkServicesEdgeCacheOrigin",
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
func NetworkServicesEdgeCacheOrigin_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNetworkServicesEdgeCacheOrigin_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.networkServicesEdgeCacheOrigin.NetworkServicesEdgeCacheOrigin",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func NetworkServicesEdgeCacheOrigin_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNetworkServicesEdgeCacheOrigin_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.networkServicesEdgeCacheOrigin.NetworkServicesEdgeCacheOrigin",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func NetworkServicesEdgeCacheOrigin_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNetworkServicesEdgeCacheOrigin_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.networkServicesEdgeCacheOrigin.NetworkServicesEdgeCacheOrigin",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func NetworkServicesEdgeCacheOrigin_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-google.networkServicesEdgeCacheOrigin.NetworkServicesEdgeCacheOrigin",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (n *jsiiProxy_NetworkServicesEdgeCacheOrigin) AddMoveTarget(moveTarget *string) {
	if err := n.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (n *jsiiProxy_NetworkServicesEdgeCacheOrigin) AddOverride(path *string, value interface{}) {
	if err := n.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (n *jsiiProxy_NetworkServicesEdgeCacheOrigin) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (n *jsiiProxy_NetworkServicesEdgeCacheOrigin) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (n *jsiiProxy_NetworkServicesEdgeCacheOrigin) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (n *jsiiProxy_NetworkServicesEdgeCacheOrigin) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (n *jsiiProxy_NetworkServicesEdgeCacheOrigin) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (n *jsiiProxy_NetworkServicesEdgeCacheOrigin) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (n *jsiiProxy_NetworkServicesEdgeCacheOrigin) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (n *jsiiProxy_NetworkServicesEdgeCacheOrigin) GetStringAttribute(terraformAttribute *string) *string {
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

func (n *jsiiProxy_NetworkServicesEdgeCacheOrigin) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (n *jsiiProxy_NetworkServicesEdgeCacheOrigin) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesEdgeCacheOrigin) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := n.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (n *jsiiProxy_NetworkServicesEdgeCacheOrigin) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (n *jsiiProxy_NetworkServicesEdgeCacheOrigin) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (n *jsiiProxy_NetworkServicesEdgeCacheOrigin) MoveFromId(id *string) {
	if err := n.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"moveFromId",
		[]interface{}{id},
	)
}

func (n *jsiiProxy_NetworkServicesEdgeCacheOrigin) MoveTo(moveTarget *string, index interface{}) {
	if err := n.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (n *jsiiProxy_NetworkServicesEdgeCacheOrigin) MoveToId(id *string) {
	if err := n.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"moveToId",
		[]interface{}{id},
	)
}

func (n *jsiiProxy_NetworkServicesEdgeCacheOrigin) OverrideLogicalId(newLogicalId *string) {
	if err := n.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (n *jsiiProxy_NetworkServicesEdgeCacheOrigin) PutAwsV4Authentication(value *NetworkServicesEdgeCacheOriginAwsV4Authentication) {
	if err := n.validatePutAwsV4AuthenticationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putAwsV4Authentication",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkServicesEdgeCacheOrigin) PutFlexShielding(value *NetworkServicesEdgeCacheOriginFlexShielding) {
	if err := n.validatePutFlexShieldingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putFlexShielding",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkServicesEdgeCacheOrigin) PutOriginOverrideAction(value *NetworkServicesEdgeCacheOriginOriginOverrideAction) {
	if err := n.validatePutOriginOverrideActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putOriginOverrideAction",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkServicesEdgeCacheOrigin) PutOriginRedirect(value *NetworkServicesEdgeCacheOriginOriginRedirect) {
	if err := n.validatePutOriginRedirectParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putOriginRedirect",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkServicesEdgeCacheOrigin) PutTimeout(value *NetworkServicesEdgeCacheOriginTimeout) {
	if err := n.validatePutTimeoutParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putTimeout",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkServicesEdgeCacheOrigin) PutTimeouts(value *NetworkServicesEdgeCacheOriginTimeouts) {
	if err := n.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkServicesEdgeCacheOrigin) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := n.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (n *jsiiProxy_NetworkServicesEdgeCacheOrigin) ResetAwsV4Authentication() {
	_jsii_.InvokeVoid(
		n,
		"resetAwsV4Authentication",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkServicesEdgeCacheOrigin) ResetDescription() {
	_jsii_.InvokeVoid(
		n,
		"resetDescription",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkServicesEdgeCacheOrigin) ResetFailoverOrigin() {
	_jsii_.InvokeVoid(
		n,
		"resetFailoverOrigin",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkServicesEdgeCacheOrigin) ResetFlexShielding() {
	_jsii_.InvokeVoid(
		n,
		"resetFlexShielding",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkServicesEdgeCacheOrigin) ResetId() {
	_jsii_.InvokeVoid(
		n,
		"resetId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkServicesEdgeCacheOrigin) ResetLabels() {
	_jsii_.InvokeVoid(
		n,
		"resetLabels",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkServicesEdgeCacheOrigin) ResetMaxAttempts() {
	_jsii_.InvokeVoid(
		n,
		"resetMaxAttempts",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkServicesEdgeCacheOrigin) ResetOriginOverrideAction() {
	_jsii_.InvokeVoid(
		n,
		"resetOriginOverrideAction",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkServicesEdgeCacheOrigin) ResetOriginRedirect() {
	_jsii_.InvokeVoid(
		n,
		"resetOriginRedirect",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkServicesEdgeCacheOrigin) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		n,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkServicesEdgeCacheOrigin) ResetPort() {
	_jsii_.InvokeVoid(
		n,
		"resetPort",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkServicesEdgeCacheOrigin) ResetProject() {
	_jsii_.InvokeVoid(
		n,
		"resetProject",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkServicesEdgeCacheOrigin) ResetProtocol() {
	_jsii_.InvokeVoid(
		n,
		"resetProtocol",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkServicesEdgeCacheOrigin) ResetRetryConditions() {
	_jsii_.InvokeVoid(
		n,
		"resetRetryConditions",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkServicesEdgeCacheOrigin) ResetTimeout() {
	_jsii_.InvokeVoid(
		n,
		"resetTimeout",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkServicesEdgeCacheOrigin) ResetTimeouts() {
	_jsii_.InvokeVoid(
		n,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkServicesEdgeCacheOrigin) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesEdgeCacheOrigin) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesEdgeCacheOrigin) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesEdgeCacheOrigin) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesEdgeCacheOrigin) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesEdgeCacheOrigin) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesEdgeCacheOrigin) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

