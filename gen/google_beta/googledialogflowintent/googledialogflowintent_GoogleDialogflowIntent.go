package googledialogflowintent

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledialogflowintent/internal"
)

// Represents a {@link https://www.terraform.io/docs/providers/google-beta/r/google_dialogflow_intent google_dialogflow_intent}.
type GoogleDialogflowIntent interface {
	cdktf.TerraformResource
	Action() *string
	SetAction(val *string)
	ActionInput() *string
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
	DefaultResponsePlatforms() *[]*string
	SetDefaultResponsePlatforms(val *[]*string)
	DefaultResponsePlatformsInput() *[]*string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	Events() *[]*string
	SetEvents(val *[]*string)
	EventsInput() *[]*string
	FollowupIntentInfo() GoogleDialogflowIntentFollowupIntentInfoList
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	InputContextNames() *[]*string
	SetInputContextNames(val *[]*string)
	InputContextNamesInput() *[]*string
	IsFallback() interface{}
	SetIsFallback(val interface{})
	IsFallbackInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MlDisabled() interface{}
	SetMlDisabled(val interface{})
	MlDisabledInput() interface{}
	Name() *string
	// The tree node.
	Node() constructs.Node
	ParentFollowupIntentName() *string
	SetParentFollowupIntentName(val *string)
	ParentFollowupIntentNameInput() *string
	Priority() *float64
	SetPriority(val *float64)
	PriorityInput() *float64
	Project() *string
	SetProject(val *string)
	ProjectInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	ResetContexts() interface{}
	SetResetContexts(val interface{})
	ResetContextsInput() interface{}
	RootFollowupIntentName() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleDialogflowIntentTimeoutsOutputReference
	TimeoutsInput() interface{}
	WebhookState() *string
	SetWebhookState(val *string)
	WebhookStateInput() *string
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
	PutTimeouts(value *GoogleDialogflowIntentTimeouts)
	ResetAction()
	ResetDefaultResponsePlatforms()
	ResetEvents()
	ResetId()
	ResetInputContextNames()
	ResetIsFallback()
	ResetMlDisabled()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetParentFollowupIntentName()
	ResetPriority()
	ResetProject()
	ResetResetContexts()
	ResetTimeouts()
	ResetWebhookState()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for GoogleDialogflowIntent
type jsiiProxy_GoogleDialogflowIntent struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleDialogflowIntent) Action() *string {
	var returns *string
	_jsii_.Get(
		j,
		"action",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowIntent) ActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowIntent) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowIntent) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowIntent) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowIntent) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowIntent) DefaultResponsePlatforms() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"defaultResponsePlatforms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowIntent) DefaultResponsePlatformsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"defaultResponsePlatformsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowIntent) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowIntent) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowIntent) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowIntent) Events() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"events",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowIntent) EventsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"eventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowIntent) FollowupIntentInfo() GoogleDialogflowIntentFollowupIntentInfoList {
	var returns GoogleDialogflowIntentFollowupIntentInfoList
	_jsii_.Get(
		j,
		"followupIntentInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowIntent) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowIntent) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowIntent) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowIntent) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowIntent) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowIntent) InputContextNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"inputContextNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowIntent) InputContextNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"inputContextNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowIntent) IsFallback() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isFallback",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowIntent) IsFallbackInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isFallbackInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowIntent) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowIntent) MlDisabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mlDisabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowIntent) MlDisabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mlDisabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowIntent) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowIntent) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowIntent) ParentFollowupIntentName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentFollowupIntentName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowIntent) ParentFollowupIntentNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentFollowupIntentNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowIntent) Priority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowIntent) PriorityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowIntent) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowIntent) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowIntent) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowIntent) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowIntent) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowIntent) ResetContexts() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resetContexts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowIntent) ResetContextsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resetContextsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowIntent) RootFollowupIntentName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rootFollowupIntentName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowIntent) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowIntent) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowIntent) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowIntent) Timeouts() GoogleDialogflowIntentTimeoutsOutputReference {
	var returns GoogleDialogflowIntentTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowIntent) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowIntent) WebhookState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webhookState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowIntent) WebhookStateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webhookStateInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/google-beta/r/google_dialogflow_intent google_dialogflow_intent} Resource.
func NewGoogleDialogflowIntent(scope constructs.Construct, id *string, config *GoogleDialogflowIntentConfig) GoogleDialogflowIntent {
	_init_.Initialize()

	if err := validateNewGoogleDialogflowIntentParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDialogflowIntent{}

	_jsii_.Create(
		"google-beta.googleDialogflowIntent.GoogleDialogflowIntent",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/google-beta/r/google_dialogflow_intent google_dialogflow_intent} Resource.
func NewGoogleDialogflowIntent_Override(g GoogleDialogflowIntent, scope constructs.Construct, id *string, config *GoogleDialogflowIntentConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleDialogflowIntent.GoogleDialogflowIntent",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleDialogflowIntent)SetAction(val *string) {
	if err := j.validateSetActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"action",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowIntent)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowIntent)SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowIntent)SetDefaultResponsePlatforms(val *[]*string) {
	if err := j.validateSetDefaultResponsePlatformsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultResponsePlatforms",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowIntent)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowIntent)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowIntent)SetEvents(val *[]*string) {
	if err := j.validateSetEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"events",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowIntent)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowIntent)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowIntent)SetInputContextNames(val *[]*string) {
	if err := j.validateSetInputContextNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputContextNames",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowIntent)SetIsFallback(val interface{}) {
	if err := j.validateSetIsFallbackParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isFallback",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowIntent)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowIntent)SetMlDisabled(val interface{}) {
	if err := j.validateSetMlDisabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mlDisabled",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowIntent)SetParentFollowupIntentName(val *string) {
	if err := j.validateSetParentFollowupIntentNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parentFollowupIntentName",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowIntent)SetPriority(val *float64) {
	if err := j.validateSetPriorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"priority",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowIntent)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowIntent)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowIntent)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowIntent)SetResetContexts(val interface{}) {
	if err := j.validateSetResetContextsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resetContexts",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowIntent)SetWebhookState(val *string) {
	if err := j.validateSetWebhookStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"webhookState",
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
func GoogleDialogflowIntent_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleDialogflowIntent_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"google-beta.googleDialogflowIntent.GoogleDialogflowIntent",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleDialogflowIntent_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"google-beta.googleDialogflowIntent.GoogleDialogflowIntent",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleDialogflowIntent) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleDialogflowIntent) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDialogflowIntent) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowIntent) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDialogflowIntent) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDialogflowIntent) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDialogflowIntent) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDialogflowIntent) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDialogflowIntent) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDialogflowIntent) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDialogflowIntent) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowIntent) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleDialogflowIntent) PutTimeouts(value *GoogleDialogflowIntentTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDialogflowIntent) ResetAction() {
	_jsii_.InvokeVoid(
		g,
		"resetAction",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowIntent) ResetDefaultResponsePlatforms() {
	_jsii_.InvokeVoid(
		g,
		"resetDefaultResponsePlatforms",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowIntent) ResetEvents() {
	_jsii_.InvokeVoid(
		g,
		"resetEvents",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowIntent) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowIntent) ResetInputContextNames() {
	_jsii_.InvokeVoid(
		g,
		"resetInputContextNames",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowIntent) ResetIsFallback() {
	_jsii_.InvokeVoid(
		g,
		"resetIsFallback",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowIntent) ResetMlDisabled() {
	_jsii_.InvokeVoid(
		g,
		"resetMlDisabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowIntent) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowIntent) ResetParentFollowupIntentName() {
	_jsii_.InvokeVoid(
		g,
		"resetParentFollowupIntentName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowIntent) ResetPriority() {
	_jsii_.InvokeVoid(
		g,
		"resetPriority",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowIntent) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowIntent) ResetResetContexts() {
	_jsii_.InvokeVoid(
		g,
		"resetResetContexts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowIntent) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowIntent) ResetWebhookState() {
	_jsii_.InvokeVoid(
		g,
		"resetWebhookState",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowIntent) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowIntent) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowIntent) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowIntent) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

