package waitingroomevent

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/cloudflare/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/cloudflare/waitingroomevent/internal"
)

// Represents a {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/waiting_room_event cloudflare_waiting_room_event}.
type WaitingRoomEvent interface {
	cdktn.TerraformResource
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
	CreatedOn() *string
	CustomPageHtml() *string
	SetCustomPageHtml(val *string)
	CustomPageHtmlInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DisableSessionRenewal() interface{}
	SetDisableSessionRenewal(val interface{})
	DisableSessionRenewalInput() interface{}
	EventEndTime() *string
	SetEventEndTime(val *string)
	EventEndTimeInput() *string
	EventStartTime() *string
	SetEventStartTime(val *string)
	EventStartTimeInput() *string
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
	ModifiedOn() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	NewUsersPerMinute() *float64
	SetNewUsersPerMinute(val *float64)
	NewUsersPerMinuteInput() *float64
	// The tree node.
	Node() constructs.Node
	PrequeueStartTime() *string
	SetPrequeueStartTime(val *string)
	PrequeueStartTimeInput() *string
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	QueueingMethod() *string
	SetQueueingMethod(val *string)
	QueueingMethodInput() *string
	// Experimental.
	RawOverrides() interface{}
	SessionDuration() *float64
	SetSessionDuration(val *float64)
	SessionDurationInput() *float64
	ShuffleAtEventStart() interface{}
	SetShuffleAtEventStart(val interface{})
	ShuffleAtEventStartInput() interface{}
	Suspended() interface{}
	SetSuspended(val interface{})
	SuspendedInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TotalActiveUsers() *float64
	SetTotalActiveUsers(val *float64)
	TotalActiveUsersInput() *float64
	WaitingRoomId() *string
	SetWaitingRoomId(val *string)
	WaitingRoomIdInput() *string
	ZoneId() *string
	SetZoneId(val *string)
	ZoneIdInput() *string
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
	ResetCustomPageHtml()
	ResetDescription()
	ResetDisableSessionRenewal()
	ResetId()
	ResetNewUsersPerMinute()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPrequeueStartTime()
	ResetQueueingMethod()
	ResetSessionDuration()
	ResetShuffleAtEventStart()
	ResetSuspended()
	ResetTotalActiveUsers()
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

// The jsii proxy struct for WaitingRoomEvent
type jsiiProxy_WaitingRoomEvent struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_WaitingRoomEvent) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaitingRoomEvent) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaitingRoomEvent) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaitingRoomEvent) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaitingRoomEvent) CreatedOn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaitingRoomEvent) CustomPageHtml() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customPageHtml",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaitingRoomEvent) CustomPageHtmlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customPageHtmlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaitingRoomEvent) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaitingRoomEvent) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaitingRoomEvent) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaitingRoomEvent) DisableSessionRenewal() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableSessionRenewal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaitingRoomEvent) DisableSessionRenewalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableSessionRenewalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaitingRoomEvent) EventEndTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventEndTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaitingRoomEvent) EventEndTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventEndTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaitingRoomEvent) EventStartTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventStartTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaitingRoomEvent) EventStartTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventStartTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaitingRoomEvent) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaitingRoomEvent) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaitingRoomEvent) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaitingRoomEvent) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaitingRoomEvent) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaitingRoomEvent) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaitingRoomEvent) ModifiedOn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modifiedOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaitingRoomEvent) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaitingRoomEvent) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaitingRoomEvent) NewUsersPerMinute() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"newUsersPerMinute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaitingRoomEvent) NewUsersPerMinuteInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"newUsersPerMinuteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaitingRoomEvent) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaitingRoomEvent) PrequeueStartTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prequeueStartTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaitingRoomEvent) PrequeueStartTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prequeueStartTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaitingRoomEvent) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaitingRoomEvent) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaitingRoomEvent) QueueingMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queueingMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaitingRoomEvent) QueueingMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queueingMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaitingRoomEvent) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaitingRoomEvent) SessionDuration() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sessionDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaitingRoomEvent) SessionDurationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sessionDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaitingRoomEvent) ShuffleAtEventStart() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shuffleAtEventStart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaitingRoomEvent) ShuffleAtEventStartInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shuffleAtEventStartInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaitingRoomEvent) Suspended() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"suspended",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaitingRoomEvent) SuspendedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"suspendedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaitingRoomEvent) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaitingRoomEvent) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaitingRoomEvent) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaitingRoomEvent) TotalActiveUsers() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"totalActiveUsers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaitingRoomEvent) TotalActiveUsersInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"totalActiveUsersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaitingRoomEvent) WaitingRoomId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"waitingRoomId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaitingRoomEvent) WaitingRoomIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"waitingRoomIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaitingRoomEvent) ZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaitingRoomEvent) ZoneIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/waiting_room_event cloudflare_waiting_room_event} Resource.
func NewWaitingRoomEvent(scope constructs.Construct, id *string, config *WaitingRoomEventConfig) WaitingRoomEvent {
	_init_.Initialize()

	if err := validateNewWaitingRoomEventParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_WaitingRoomEvent{}

	_jsii_.Create(
		"@cdktn/provider-cloudflare.waitingRoomEvent.WaitingRoomEvent",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/waiting_room_event cloudflare_waiting_room_event} Resource.
func NewWaitingRoomEvent_Override(w WaitingRoomEvent, scope constructs.Construct, id *string, config *WaitingRoomEventConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-cloudflare.waitingRoomEvent.WaitingRoomEvent",
		[]interface{}{scope, id, config},
		w,
	)
}

func (j *jsiiProxy_WaitingRoomEvent)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_WaitingRoomEvent)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_WaitingRoomEvent)SetCustomPageHtml(val *string) {
	if err := j.validateSetCustomPageHtmlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customPageHtml",
		val,
	)
}

func (j *jsiiProxy_WaitingRoomEvent)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_WaitingRoomEvent)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_WaitingRoomEvent)SetDisableSessionRenewal(val interface{}) {
	if err := j.validateSetDisableSessionRenewalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableSessionRenewal",
		val,
	)
}

func (j *jsiiProxy_WaitingRoomEvent)SetEventEndTime(val *string) {
	if err := j.validateSetEventEndTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventEndTime",
		val,
	)
}

func (j *jsiiProxy_WaitingRoomEvent)SetEventStartTime(val *string) {
	if err := j.validateSetEventStartTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventStartTime",
		val,
	)
}

func (j *jsiiProxy_WaitingRoomEvent)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_WaitingRoomEvent)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_WaitingRoomEvent)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_WaitingRoomEvent)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_WaitingRoomEvent)SetNewUsersPerMinute(val *float64) {
	if err := j.validateSetNewUsersPerMinuteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"newUsersPerMinute",
		val,
	)
}

func (j *jsiiProxy_WaitingRoomEvent)SetPrequeueStartTime(val *string) {
	if err := j.validateSetPrequeueStartTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"prequeueStartTime",
		val,
	)
}

func (j *jsiiProxy_WaitingRoomEvent)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_WaitingRoomEvent)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_WaitingRoomEvent)SetQueueingMethod(val *string) {
	if err := j.validateSetQueueingMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queueingMethod",
		val,
	)
}

func (j *jsiiProxy_WaitingRoomEvent)SetSessionDuration(val *float64) {
	if err := j.validateSetSessionDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sessionDuration",
		val,
	)
}

func (j *jsiiProxy_WaitingRoomEvent)SetShuffleAtEventStart(val interface{}) {
	if err := j.validateSetShuffleAtEventStartParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shuffleAtEventStart",
		val,
	)
}

func (j *jsiiProxy_WaitingRoomEvent)SetSuspended(val interface{}) {
	if err := j.validateSetSuspendedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"suspended",
		val,
	)
}

func (j *jsiiProxy_WaitingRoomEvent)SetTotalActiveUsers(val *float64) {
	if err := j.validateSetTotalActiveUsersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"totalActiveUsers",
		val,
	)
}

func (j *jsiiProxy_WaitingRoomEvent)SetWaitingRoomId(val *string) {
	if err := j.validateSetWaitingRoomIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"waitingRoomId",
		val,
	)
}

func (j *jsiiProxy_WaitingRoomEvent)SetZoneId(val *string) {
	if err := j.validateSetZoneIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zoneId",
		val,
	)
}

// Generates CDKTN code for importing a WaitingRoomEvent resource upon running "cdktn plan <stack-name>".
func WaitingRoomEvent_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateWaitingRoomEvent_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-cloudflare.waitingRoomEvent.WaitingRoomEvent",
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
func WaitingRoomEvent_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateWaitingRoomEvent_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-cloudflare.waitingRoomEvent.WaitingRoomEvent",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func WaitingRoomEvent_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateWaitingRoomEvent_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-cloudflare.waitingRoomEvent.WaitingRoomEvent",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func WaitingRoomEvent_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateWaitingRoomEvent_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-cloudflare.waitingRoomEvent.WaitingRoomEvent",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func WaitingRoomEvent_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-cloudflare.waitingRoomEvent.WaitingRoomEvent",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (w *jsiiProxy_WaitingRoomEvent) AddMoveTarget(moveTarget *string) {
	if err := w.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (w *jsiiProxy_WaitingRoomEvent) AddOverride(path *string, value interface{}) {
	if err := w.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (w *jsiiProxy_WaitingRoomEvent) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (w *jsiiProxy_WaitingRoomEvent) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (w *jsiiProxy_WaitingRoomEvent) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (w *jsiiProxy_WaitingRoomEvent) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (w *jsiiProxy_WaitingRoomEvent) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (w *jsiiProxy_WaitingRoomEvent) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (w *jsiiProxy_WaitingRoomEvent) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (w *jsiiProxy_WaitingRoomEvent) GetStringAttribute(terraformAttribute *string) *string {
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

func (w *jsiiProxy_WaitingRoomEvent) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (w *jsiiProxy_WaitingRoomEvent) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WaitingRoomEvent) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := w.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (w *jsiiProxy_WaitingRoomEvent) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (w *jsiiProxy_WaitingRoomEvent) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (w *jsiiProxy_WaitingRoomEvent) MoveFromId(id *string) {
	if err := w.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"moveFromId",
		[]interface{}{id},
	)
}

func (w *jsiiProxy_WaitingRoomEvent) MoveTo(moveTarget *string, index interface{}) {
	if err := w.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (w *jsiiProxy_WaitingRoomEvent) MoveToId(id *string) {
	if err := w.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"moveToId",
		[]interface{}{id},
	)
}

func (w *jsiiProxy_WaitingRoomEvent) OverrideLogicalId(newLogicalId *string) {
	if err := w.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (w *jsiiProxy_WaitingRoomEvent) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := w.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (w *jsiiProxy_WaitingRoomEvent) ResetCustomPageHtml() {
	_jsii_.InvokeVoid(
		w,
		"resetCustomPageHtml",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WaitingRoomEvent) ResetDescription() {
	_jsii_.InvokeVoid(
		w,
		"resetDescription",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WaitingRoomEvent) ResetDisableSessionRenewal() {
	_jsii_.InvokeVoid(
		w,
		"resetDisableSessionRenewal",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WaitingRoomEvent) ResetId() {
	_jsii_.InvokeVoid(
		w,
		"resetId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WaitingRoomEvent) ResetNewUsersPerMinute() {
	_jsii_.InvokeVoid(
		w,
		"resetNewUsersPerMinute",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WaitingRoomEvent) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		w,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WaitingRoomEvent) ResetPrequeueStartTime() {
	_jsii_.InvokeVoid(
		w,
		"resetPrequeueStartTime",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WaitingRoomEvent) ResetQueueingMethod() {
	_jsii_.InvokeVoid(
		w,
		"resetQueueingMethod",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WaitingRoomEvent) ResetSessionDuration() {
	_jsii_.InvokeVoid(
		w,
		"resetSessionDuration",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WaitingRoomEvent) ResetShuffleAtEventStart() {
	_jsii_.InvokeVoid(
		w,
		"resetShuffleAtEventStart",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WaitingRoomEvent) ResetSuspended() {
	_jsii_.InvokeVoid(
		w,
		"resetSuspended",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WaitingRoomEvent) ResetTotalActiveUsers() {
	_jsii_.InvokeVoid(
		w,
		"resetTotalActiveUsers",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WaitingRoomEvent) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WaitingRoomEvent) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WaitingRoomEvent) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WaitingRoomEvent) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WaitingRoomEvent) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WaitingRoomEvent) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WaitingRoomEvent) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

