package computeserviceattachment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google/computeserviceattachment/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/compute_service_attachment google_compute_service_attachment}.
type ComputeServiceAttachment interface {
	cdktn.TerraformResource
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	ConnectedEndpoints() ComputeServiceAttachmentConnectedEndpointsList
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	ConnectionPreference() *string
	SetConnectionPreference(val *string)
	ConnectionPreferenceInput() *string
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ConsumerAcceptLists() ComputeServiceAttachmentConsumerAcceptListsList
	ConsumerAcceptListsInput() interface{}
	ConsumerRejectLists() *[]*string
	SetConsumerRejectLists(val *[]*string)
	ConsumerRejectListsInput() *[]*string
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
	DomainNames() *[]*string
	SetDomainNames(val *[]*string)
	DomainNamesInput() *[]*string
	EnableProxyProtocol() interface{}
	SetEnableProxyProtocol(val interface{})
	EnableProxyProtocolInput() interface{}
	Fingerprint() *string
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
	Name() *string
	SetName(val *string)
	NameInput() *string
	NatSubnets() *[]*string
	SetNatSubnets(val *[]*string)
	NatSubnetsInput() *[]*string
	// The tree node.
	Node() constructs.Node
	Project() *string
	SetProject(val *string)
	ProjectInput() *string
	PropagatedConnectionLimit() *float64
	SetPropagatedConnectionLimit(val *float64)
	PropagatedConnectionLimitInput() *float64
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	PscServiceAttachmentId() ComputeServiceAttachmentPscServiceAttachmentIdList
	// Experimental.
	RawOverrides() interface{}
	ReconcileConnections() interface{}
	SetReconcileConnections(val interface{})
	ReconcileConnectionsInput() interface{}
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	SelfLink() *string
	SendPropagatedConnectionLimitIfZero() interface{}
	SetSendPropagatedConnectionLimitIfZero(val interface{})
	SendPropagatedConnectionLimitIfZeroInput() interface{}
	ShowNatIps() interface{}
	SetShowNatIps(val interface{})
	ShowNatIpsInput() interface{}
	TargetService() *string
	SetTargetService(val *string)
	TargetServiceInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() ComputeServiceAttachmentTimeoutsOutputReference
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
	PutConsumerAcceptLists(value interface{})
	PutTimeouts(value *ComputeServiceAttachmentTimeouts)
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
	ResetConsumerAcceptLists()
	ResetConsumerRejectLists()
	ResetDescription()
	ResetDomainNames()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetPropagatedConnectionLimit()
	ResetReconcileConnections()
	ResetRegion()
	ResetSendPropagatedConnectionLimitIfZero()
	ResetShowNatIps()
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

// The jsii proxy struct for ComputeServiceAttachment
type jsiiProxy_ComputeServiceAttachment struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_ComputeServiceAttachment) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeServiceAttachment) ConnectedEndpoints() ComputeServiceAttachmentConnectedEndpointsList {
	var returns ComputeServiceAttachmentConnectedEndpointsList
	_jsii_.Get(
		j,
		"connectedEndpoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeServiceAttachment) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeServiceAttachment) ConnectionPreference() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionPreference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeServiceAttachment) ConnectionPreferenceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionPreferenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeServiceAttachment) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeServiceAttachment) ConsumerAcceptLists() ComputeServiceAttachmentConsumerAcceptListsList {
	var returns ComputeServiceAttachmentConsumerAcceptListsList
	_jsii_.Get(
		j,
		"consumerAcceptLists",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeServiceAttachment) ConsumerAcceptListsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"consumerAcceptListsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeServiceAttachment) ConsumerRejectLists() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"consumerRejectLists",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeServiceAttachment) ConsumerRejectListsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"consumerRejectListsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeServiceAttachment) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeServiceAttachment) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeServiceAttachment) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeServiceAttachment) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeServiceAttachment) DomainNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"domainNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeServiceAttachment) DomainNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"domainNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeServiceAttachment) EnableProxyProtocol() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableProxyProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeServiceAttachment) EnableProxyProtocolInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableProxyProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeServiceAttachment) Fingerprint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fingerprint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeServiceAttachment) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeServiceAttachment) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeServiceAttachment) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeServiceAttachment) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeServiceAttachment) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeServiceAttachment) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeServiceAttachment) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeServiceAttachment) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeServiceAttachment) NatSubnets() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"natSubnets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeServiceAttachment) NatSubnetsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"natSubnetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeServiceAttachment) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeServiceAttachment) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeServiceAttachment) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeServiceAttachment) PropagatedConnectionLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"propagatedConnectionLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeServiceAttachment) PropagatedConnectionLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"propagatedConnectionLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeServiceAttachment) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeServiceAttachment) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeServiceAttachment) PscServiceAttachmentId() ComputeServiceAttachmentPscServiceAttachmentIdList {
	var returns ComputeServiceAttachmentPscServiceAttachmentIdList
	_jsii_.Get(
		j,
		"pscServiceAttachmentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeServiceAttachment) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeServiceAttachment) ReconcileConnections() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"reconcileConnections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeServiceAttachment) ReconcileConnectionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"reconcileConnectionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeServiceAttachment) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeServiceAttachment) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeServiceAttachment) SelfLink() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selfLink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeServiceAttachment) SendPropagatedConnectionLimitIfZero() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sendPropagatedConnectionLimitIfZero",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeServiceAttachment) SendPropagatedConnectionLimitIfZeroInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sendPropagatedConnectionLimitIfZeroInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeServiceAttachment) ShowNatIps() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showNatIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeServiceAttachment) ShowNatIpsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showNatIpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeServiceAttachment) TargetService() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeServiceAttachment) TargetServiceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetServiceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeServiceAttachment) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeServiceAttachment) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeServiceAttachment) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeServiceAttachment) Timeouts() ComputeServiceAttachmentTimeoutsOutputReference {
	var returns ComputeServiceAttachmentTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeServiceAttachment) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/compute_service_attachment google_compute_service_attachment} Resource.
func NewComputeServiceAttachment(scope constructs.Construct, id *string, config *ComputeServiceAttachmentConfig) ComputeServiceAttachment {
	_init_.Initialize()

	if err := validateNewComputeServiceAttachmentParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeServiceAttachment{}

	_jsii_.Create(
		"@cdktn/provider-google.computeServiceAttachment.ComputeServiceAttachment",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/compute_service_attachment google_compute_service_attachment} Resource.
func NewComputeServiceAttachment_Override(c ComputeServiceAttachment, scope constructs.Construct, id *string, config *ComputeServiceAttachmentConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.computeServiceAttachment.ComputeServiceAttachment",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_ComputeServiceAttachment)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ComputeServiceAttachment)SetConnectionPreference(val *string) {
	if err := j.validateSetConnectionPreferenceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionPreference",
		val,
	)
}

func (j *jsiiProxy_ComputeServiceAttachment)SetConsumerRejectLists(val *[]*string) {
	if err := j.validateSetConsumerRejectListsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"consumerRejectLists",
		val,
	)
}

func (j *jsiiProxy_ComputeServiceAttachment)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ComputeServiceAttachment)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ComputeServiceAttachment)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ComputeServiceAttachment)SetDomainNames(val *[]*string) {
	if err := j.validateSetDomainNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainNames",
		val,
	)
}

func (j *jsiiProxy_ComputeServiceAttachment)SetEnableProxyProtocol(val interface{}) {
	if err := j.validateSetEnableProxyProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableProxyProtocol",
		val,
	)
}

func (j *jsiiProxy_ComputeServiceAttachment)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ComputeServiceAttachment)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ComputeServiceAttachment)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ComputeServiceAttachment)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ComputeServiceAttachment)SetNatSubnets(val *[]*string) {
	if err := j.validateSetNatSubnetsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"natSubnets",
		val,
	)
}

func (j *jsiiProxy_ComputeServiceAttachment)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_ComputeServiceAttachment)SetPropagatedConnectionLimit(val *float64) {
	if err := j.validateSetPropagatedConnectionLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"propagatedConnectionLimit",
		val,
	)
}

func (j *jsiiProxy_ComputeServiceAttachment)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ComputeServiceAttachment)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ComputeServiceAttachment)SetReconcileConnections(val interface{}) {
	if err := j.validateSetReconcileConnectionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reconcileConnections",
		val,
	)
}

func (j *jsiiProxy_ComputeServiceAttachment)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_ComputeServiceAttachment)SetSendPropagatedConnectionLimitIfZero(val interface{}) {
	if err := j.validateSetSendPropagatedConnectionLimitIfZeroParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sendPropagatedConnectionLimitIfZero",
		val,
	)
}

func (j *jsiiProxy_ComputeServiceAttachment)SetShowNatIps(val interface{}) {
	if err := j.validateSetShowNatIpsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"showNatIps",
		val,
	)
}

func (j *jsiiProxy_ComputeServiceAttachment)SetTargetService(val *string) {
	if err := j.validateSetTargetServiceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetService",
		val,
	)
}

// Generates CDKTN code for importing a ComputeServiceAttachment resource upon running "cdktn plan <stack-name>".
func ComputeServiceAttachment_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateComputeServiceAttachment_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.computeServiceAttachment.ComputeServiceAttachment",
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
func ComputeServiceAttachment_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateComputeServiceAttachment_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.computeServiceAttachment.ComputeServiceAttachment",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ComputeServiceAttachment_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateComputeServiceAttachment_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.computeServiceAttachment.ComputeServiceAttachment",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ComputeServiceAttachment_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateComputeServiceAttachment_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.computeServiceAttachment.ComputeServiceAttachment",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ComputeServiceAttachment_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-google.computeServiceAttachment.ComputeServiceAttachment",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_ComputeServiceAttachment) AddMoveTarget(moveTarget *string) {
	if err := c.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (c *jsiiProxy_ComputeServiceAttachment) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_ComputeServiceAttachment) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ComputeServiceAttachment) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ComputeServiceAttachment) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ComputeServiceAttachment) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ComputeServiceAttachment) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ComputeServiceAttachment) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ComputeServiceAttachment) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ComputeServiceAttachment) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ComputeServiceAttachment) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ComputeServiceAttachment) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeServiceAttachment) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := c.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (c *jsiiProxy_ComputeServiceAttachment) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ComputeServiceAttachment) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (c *jsiiProxy_ComputeServiceAttachment) MoveFromId(id *string) {
	if err := c.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveFromId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_ComputeServiceAttachment) MoveTo(moveTarget *string, index interface{}) {
	if err := c.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (c *jsiiProxy_ComputeServiceAttachment) MoveToId(id *string) {
	if err := c.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveToId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_ComputeServiceAttachment) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_ComputeServiceAttachment) PutConsumerAcceptLists(value interface{}) {
	if err := c.validatePutConsumerAcceptListsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putConsumerAcceptLists",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeServiceAttachment) PutTimeouts(value *ComputeServiceAttachmentTimeouts) {
	if err := c.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeServiceAttachment) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := c.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (c *jsiiProxy_ComputeServiceAttachment) ResetConsumerAcceptLists() {
	_jsii_.InvokeVoid(
		c,
		"resetConsumerAcceptLists",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeServiceAttachment) ResetConsumerRejectLists() {
	_jsii_.InvokeVoid(
		c,
		"resetConsumerRejectLists",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeServiceAttachment) ResetDescription() {
	_jsii_.InvokeVoid(
		c,
		"resetDescription",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeServiceAttachment) ResetDomainNames() {
	_jsii_.InvokeVoid(
		c,
		"resetDomainNames",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeServiceAttachment) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeServiceAttachment) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeServiceAttachment) ResetProject() {
	_jsii_.InvokeVoid(
		c,
		"resetProject",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeServiceAttachment) ResetPropagatedConnectionLimit() {
	_jsii_.InvokeVoid(
		c,
		"resetPropagatedConnectionLimit",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeServiceAttachment) ResetReconcileConnections() {
	_jsii_.InvokeVoid(
		c,
		"resetReconcileConnections",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeServiceAttachment) ResetRegion() {
	_jsii_.InvokeVoid(
		c,
		"resetRegion",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeServiceAttachment) ResetSendPropagatedConnectionLimitIfZero() {
	_jsii_.InvokeVoid(
		c,
		"resetSendPropagatedConnectionLimitIfZero",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeServiceAttachment) ResetShowNatIps() {
	_jsii_.InvokeVoid(
		c,
		"resetShowNatIps",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeServiceAttachment) ResetTimeouts() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeServiceAttachment) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeServiceAttachment) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeServiceAttachment) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeServiceAttachment) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeServiceAttachment) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeServiceAttachment) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeServiceAttachment) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

