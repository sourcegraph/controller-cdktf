package waitingroom

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/cloudflare/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/cloudflare/waitingroom/internal"
)

// Represents a {@link https://www.terraform.io/docs/providers/cloudflare/r/waiting_room cloudflare_waiting_room}.
type WaitingRoom interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() *float64
	// Experimental.
	SetCount(val *float64)
	CustomPageHtml() *string
	SetCustomPageHtml(val *string)
	CustomPageHtmlInput() *string
	DefaultTemplateLanguage() *string
	SetDefaultTemplateLanguage(val *string)
	DefaultTemplateLanguageInput() *string
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
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Host() *string
	SetHost(val *string)
	HostInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	JsonResponseEnabled() interface{}
	SetJsonResponseEnabled(val interface{})
	JsonResponseEnabledInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	NewUsersPerMinute() *float64
	SetNewUsersPerMinute(val *float64)
	NewUsersPerMinuteInput() *float64
	// The tree node.
	Node() constructs.Node
	Path() *string
	SetPath(val *string)
	PathInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	QueueAll() interface{}
	SetQueueAll(val interface{})
	QueueAllInput() interface{}
	QueueingMethod() *string
	SetQueueingMethod(val *string)
	QueueingMethodInput() *string
	// Experimental.
	RawOverrides() interface{}
	SessionDuration() *float64
	SetSessionDuration(val *float64)
	SessionDurationInput() *float64
	Suspended() interface{}
	SetSuspended(val interface{})
	SuspendedInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() WaitingRoomTimeoutsOutputReference
	TimeoutsInput() interface{}
	TotalActiveUsers() *float64
	SetTotalActiveUsers(val *float64)
	TotalActiveUsersInput() *float64
	ZoneId() *string
	SetZoneId(val *string)
	ZoneIdInput() *string
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutTimeouts(value *WaitingRoomTimeouts)
	ResetCustomPageHtml()
	ResetDefaultTemplateLanguage()
	ResetDescription()
	ResetDisableSessionRenewal()
	ResetId()
	ResetJsonResponseEnabled()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPath()
	ResetQueueAll()
	ResetQueueingMethod()
	ResetSessionDuration()
	ResetSuspended()
	ResetTimeouts()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for WaitingRoom
type jsiiProxy_WaitingRoom struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_WaitingRoom) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaitingRoom) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaitingRoom) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaitingRoom) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaitingRoom) CustomPageHtml() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customPageHtml",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaitingRoom) CustomPageHtmlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customPageHtmlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaitingRoom) DefaultTemplateLanguage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultTemplateLanguage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaitingRoom) DefaultTemplateLanguageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultTemplateLanguageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaitingRoom) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaitingRoom) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaitingRoom) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaitingRoom) DisableSessionRenewal() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableSessionRenewal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaitingRoom) DisableSessionRenewalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableSessionRenewalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaitingRoom) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaitingRoom) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaitingRoom) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaitingRoom) Host() *string {
	var returns *string
	_jsii_.Get(
		j,
		"host",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaitingRoom) HostInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaitingRoom) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaitingRoom) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaitingRoom) JsonResponseEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jsonResponseEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaitingRoom) JsonResponseEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jsonResponseEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaitingRoom) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaitingRoom) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaitingRoom) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaitingRoom) NewUsersPerMinute() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"newUsersPerMinute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaitingRoom) NewUsersPerMinuteInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"newUsersPerMinuteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaitingRoom) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaitingRoom) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaitingRoom) PathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaitingRoom) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaitingRoom) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaitingRoom) QueueAll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"queueAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaitingRoom) QueueAllInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"queueAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaitingRoom) QueueingMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queueingMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaitingRoom) QueueingMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queueingMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaitingRoom) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaitingRoom) SessionDuration() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sessionDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaitingRoom) SessionDurationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sessionDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaitingRoom) Suspended() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"suspended",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaitingRoom) SuspendedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"suspendedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaitingRoom) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaitingRoom) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaitingRoom) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaitingRoom) Timeouts() WaitingRoomTimeoutsOutputReference {
	var returns WaitingRoomTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaitingRoom) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaitingRoom) TotalActiveUsers() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"totalActiveUsers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaitingRoom) TotalActiveUsersInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"totalActiveUsersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaitingRoom) ZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaitingRoom) ZoneIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/cloudflare/r/waiting_room cloudflare_waiting_room} Resource.
func NewWaitingRoom(scope constructs.Construct, id *string, config *WaitingRoomConfig) WaitingRoom {
	_init_.Initialize()

	if err := validateNewWaitingRoomParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_WaitingRoom{}

	_jsii_.Create(
		"cloudflare.waitingRoom.WaitingRoom",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/cloudflare/r/waiting_room cloudflare_waiting_room} Resource.
func NewWaitingRoom_Override(w WaitingRoom, scope constructs.Construct, id *string, config *WaitingRoomConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"cloudflare.waitingRoom.WaitingRoom",
		[]interface{}{scope, id, config},
		w,
	)
}

func (j *jsiiProxy_WaitingRoom)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_WaitingRoom)SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_WaitingRoom)SetCustomPageHtml(val *string) {
	if err := j.validateSetCustomPageHtmlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customPageHtml",
		val,
	)
}

func (j *jsiiProxy_WaitingRoom)SetDefaultTemplateLanguage(val *string) {
	if err := j.validateSetDefaultTemplateLanguageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultTemplateLanguage",
		val,
	)
}

func (j *jsiiProxy_WaitingRoom)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_WaitingRoom)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_WaitingRoom)SetDisableSessionRenewal(val interface{}) {
	if err := j.validateSetDisableSessionRenewalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableSessionRenewal",
		val,
	)
}

func (j *jsiiProxy_WaitingRoom)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_WaitingRoom)SetHost(val *string) {
	if err := j.validateSetHostParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"host",
		val,
	)
}

func (j *jsiiProxy_WaitingRoom)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_WaitingRoom)SetJsonResponseEnabled(val interface{}) {
	if err := j.validateSetJsonResponseEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jsonResponseEnabled",
		val,
	)
}

func (j *jsiiProxy_WaitingRoom)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_WaitingRoom)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_WaitingRoom)SetNewUsersPerMinute(val *float64) {
	if err := j.validateSetNewUsersPerMinuteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"newUsersPerMinute",
		val,
	)
}

func (j *jsiiProxy_WaitingRoom)SetPath(val *string) {
	if err := j.validateSetPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"path",
		val,
	)
}

func (j *jsiiProxy_WaitingRoom)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_WaitingRoom)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_WaitingRoom)SetQueueAll(val interface{}) {
	if err := j.validateSetQueueAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queueAll",
		val,
	)
}

func (j *jsiiProxy_WaitingRoom)SetQueueingMethod(val *string) {
	if err := j.validateSetQueueingMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queueingMethod",
		val,
	)
}

func (j *jsiiProxy_WaitingRoom)SetSessionDuration(val *float64) {
	if err := j.validateSetSessionDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sessionDuration",
		val,
	)
}

func (j *jsiiProxy_WaitingRoom)SetSuspended(val interface{}) {
	if err := j.validateSetSuspendedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"suspended",
		val,
	)
}

func (j *jsiiProxy_WaitingRoom)SetTotalActiveUsers(val *float64) {
	if err := j.validateSetTotalActiveUsersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"totalActiveUsers",
		val,
	)
}

func (j *jsiiProxy_WaitingRoom)SetZoneId(val *string) {
	if err := j.validateSetZoneIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zoneId",
		val,
	)
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
func WaitingRoom_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateWaitingRoom_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cloudflare.waitingRoom.WaitingRoom",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func WaitingRoom_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"cloudflare.waitingRoom.WaitingRoom",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (w *jsiiProxy_WaitingRoom) AddOverride(path *string, value interface{}) {
	if err := w.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (w *jsiiProxy_WaitingRoom) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (w *jsiiProxy_WaitingRoom) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := w.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WaitingRoom) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (w *jsiiProxy_WaitingRoom) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (w *jsiiProxy_WaitingRoom) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (w *jsiiProxy_WaitingRoom) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (w *jsiiProxy_WaitingRoom) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (w *jsiiProxy_WaitingRoom) GetStringAttribute(terraformAttribute *string) *string {
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

func (w *jsiiProxy_WaitingRoom) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (w *jsiiProxy_WaitingRoom) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := w.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WaitingRoom) OverrideLogicalId(newLogicalId *string) {
	if err := w.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (w *jsiiProxy_WaitingRoom) PutTimeouts(value *WaitingRoomTimeouts) {
	if err := w.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WaitingRoom) ResetCustomPageHtml() {
	_jsii_.InvokeVoid(
		w,
		"resetCustomPageHtml",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WaitingRoom) ResetDefaultTemplateLanguage() {
	_jsii_.InvokeVoid(
		w,
		"resetDefaultTemplateLanguage",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WaitingRoom) ResetDescription() {
	_jsii_.InvokeVoid(
		w,
		"resetDescription",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WaitingRoom) ResetDisableSessionRenewal() {
	_jsii_.InvokeVoid(
		w,
		"resetDisableSessionRenewal",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WaitingRoom) ResetId() {
	_jsii_.InvokeVoid(
		w,
		"resetId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WaitingRoom) ResetJsonResponseEnabled() {
	_jsii_.InvokeVoid(
		w,
		"resetJsonResponseEnabled",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WaitingRoom) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		w,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WaitingRoom) ResetPath() {
	_jsii_.InvokeVoid(
		w,
		"resetPath",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WaitingRoom) ResetQueueAll() {
	_jsii_.InvokeVoid(
		w,
		"resetQueueAll",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WaitingRoom) ResetQueueingMethod() {
	_jsii_.InvokeVoid(
		w,
		"resetQueueingMethod",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WaitingRoom) ResetSessionDuration() {
	_jsii_.InvokeVoid(
		w,
		"resetSessionDuration",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WaitingRoom) ResetSuspended() {
	_jsii_.InvokeVoid(
		w,
		"resetSuspended",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WaitingRoom) ResetTimeouts() {
	_jsii_.InvokeVoid(
		w,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WaitingRoom) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WaitingRoom) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WaitingRoom) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WaitingRoom) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

