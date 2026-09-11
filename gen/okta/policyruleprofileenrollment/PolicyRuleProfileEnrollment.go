package policyruleprofileenrollment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/okta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/okta/policyruleprofileenrollment/internal"
)

// Represents a {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_rule_profile_enrollment okta_policy_rule_profile_enrollment}.
type PolicyRuleProfileEnrollment interface {
	cdktn.TerraformResource
	Access() *string
	SetAccess(val *string)
	AccessInput() *string
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
	EmailVerification() interface{}
	SetEmailVerification(val interface{})
	EmailVerificationInput() interface{}
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
	InlineHookId() *string
	SetInlineHookId(val *string)
	InlineHookIdInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	Name() *string
	// The tree node.
	Node() constructs.Node
	PolicyId() *string
	SetPolicyId(val *string)
	PolicyIdInput() *string
	ProfileAttributes() PolicyRuleProfileEnrollmentProfileAttributesList
	ProfileAttributesInput() interface{}
	ProgressiveProfilingAction() *string
	SetProgressiveProfilingAction(val *string)
	ProgressiveProfilingActionInput() *string
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
	Status() *string
	TargetGroupId() *string
	SetTargetGroupId(val *string)
	TargetGroupIdInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	UiSchemaId() *string
	SetUiSchemaId(val *string)
	UiSchemaIdInput() *string
	UnknownUserAction() *string
	SetUnknownUserAction(val *string)
	UnknownUserActionInput() *string
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
	PutProfileAttributes(value interface{})
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
	ResetAccess()
	ResetEmailVerification()
	ResetId()
	ResetInlineHookId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProfileAttributes()
	ResetProgressiveProfilingAction()
	ResetTargetGroupId()
	ResetUiSchemaId()
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

// The jsii proxy struct for PolicyRuleProfileEnrollment
type jsiiProxy_PolicyRuleProfileEnrollment struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_PolicyRuleProfileEnrollment) Access() *string {
	var returns *string
	_jsii_.Get(
		j,
		"access",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleProfileEnrollment) AccessInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleProfileEnrollment) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleProfileEnrollment) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleProfileEnrollment) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleProfileEnrollment) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleProfileEnrollment) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleProfileEnrollment) EmailVerification() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"emailVerification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleProfileEnrollment) EmailVerificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"emailVerificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleProfileEnrollment) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleProfileEnrollment) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleProfileEnrollment) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleProfileEnrollment) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleProfileEnrollment) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleProfileEnrollment) InlineHookId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inlineHookId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleProfileEnrollment) InlineHookIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inlineHookIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleProfileEnrollment) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleProfileEnrollment) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleProfileEnrollment) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleProfileEnrollment) PolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleProfileEnrollment) PolicyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleProfileEnrollment) ProfileAttributes() PolicyRuleProfileEnrollmentProfileAttributesList {
	var returns PolicyRuleProfileEnrollmentProfileAttributesList
	_jsii_.Get(
		j,
		"profileAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleProfileEnrollment) ProfileAttributesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"profileAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleProfileEnrollment) ProgressiveProfilingAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"progressiveProfilingAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleProfileEnrollment) ProgressiveProfilingActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"progressiveProfilingActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleProfileEnrollment) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleProfileEnrollment) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleProfileEnrollment) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleProfileEnrollment) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleProfileEnrollment) TargetGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleProfileEnrollment) TargetGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleProfileEnrollment) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleProfileEnrollment) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleProfileEnrollment) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleProfileEnrollment) UiSchemaId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uiSchemaId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleProfileEnrollment) UiSchemaIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uiSchemaIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleProfileEnrollment) UnknownUserAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unknownUserAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleProfileEnrollment) UnknownUserActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unknownUserActionInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_rule_profile_enrollment okta_policy_rule_profile_enrollment} Resource.
func NewPolicyRuleProfileEnrollment(scope constructs.Construct, id *string, config *PolicyRuleProfileEnrollmentConfig) PolicyRuleProfileEnrollment {
	_init_.Initialize()

	if err := validateNewPolicyRuleProfileEnrollmentParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_PolicyRuleProfileEnrollment{}

	_jsii_.Create(
		"@cdktn/provider-okta.policyRuleProfileEnrollment.PolicyRuleProfileEnrollment",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/policy_rule_profile_enrollment okta_policy_rule_profile_enrollment} Resource.
func NewPolicyRuleProfileEnrollment_Override(p PolicyRuleProfileEnrollment, scope constructs.Construct, id *string, config *PolicyRuleProfileEnrollmentConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-okta.policyRuleProfileEnrollment.PolicyRuleProfileEnrollment",
		[]interface{}{scope, id, config},
		p,
	)
}

func (j *jsiiProxy_PolicyRuleProfileEnrollment)SetAccess(val *string) {
	if err := j.validateSetAccessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"access",
		val,
	)
}

func (j *jsiiProxy_PolicyRuleProfileEnrollment)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_PolicyRuleProfileEnrollment)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_PolicyRuleProfileEnrollment)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_PolicyRuleProfileEnrollment)SetEmailVerification(val interface{}) {
	if err := j.validateSetEmailVerificationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"emailVerification",
		val,
	)
}

func (j *jsiiProxy_PolicyRuleProfileEnrollment)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_PolicyRuleProfileEnrollment)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_PolicyRuleProfileEnrollment)SetInlineHookId(val *string) {
	if err := j.validateSetInlineHookIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inlineHookId",
		val,
	)
}

func (j *jsiiProxy_PolicyRuleProfileEnrollment)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_PolicyRuleProfileEnrollment)SetPolicyId(val *string) {
	if err := j.validateSetPolicyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policyId",
		val,
	)
}

func (j *jsiiProxy_PolicyRuleProfileEnrollment)SetProgressiveProfilingAction(val *string) {
	if err := j.validateSetProgressiveProfilingActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"progressiveProfilingAction",
		val,
	)
}

func (j *jsiiProxy_PolicyRuleProfileEnrollment)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_PolicyRuleProfileEnrollment)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_PolicyRuleProfileEnrollment)SetTargetGroupId(val *string) {
	if err := j.validateSetTargetGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetGroupId",
		val,
	)
}

func (j *jsiiProxy_PolicyRuleProfileEnrollment)SetUiSchemaId(val *string) {
	if err := j.validateSetUiSchemaIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"uiSchemaId",
		val,
	)
}

func (j *jsiiProxy_PolicyRuleProfileEnrollment)SetUnknownUserAction(val *string) {
	if err := j.validateSetUnknownUserActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"unknownUserAction",
		val,
	)
}

// Generates CDKTN code for importing a PolicyRuleProfileEnrollment resource upon running "cdktn plan <stack-name>".
func PolicyRuleProfileEnrollment_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validatePolicyRuleProfileEnrollment_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-okta.policyRuleProfileEnrollment.PolicyRuleProfileEnrollment",
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
func PolicyRuleProfileEnrollment_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePolicyRuleProfileEnrollment_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-okta.policyRuleProfileEnrollment.PolicyRuleProfileEnrollment",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func PolicyRuleProfileEnrollment_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePolicyRuleProfileEnrollment_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-okta.policyRuleProfileEnrollment.PolicyRuleProfileEnrollment",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func PolicyRuleProfileEnrollment_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePolicyRuleProfileEnrollment_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-okta.policyRuleProfileEnrollment.PolicyRuleProfileEnrollment",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func PolicyRuleProfileEnrollment_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-okta.policyRuleProfileEnrollment.PolicyRuleProfileEnrollment",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (p *jsiiProxy_PolicyRuleProfileEnrollment) AddMoveTarget(moveTarget *string) {
	if err := p.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (p *jsiiProxy_PolicyRuleProfileEnrollment) AddOverride(path *string, value interface{}) {
	if err := p.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (p *jsiiProxy_PolicyRuleProfileEnrollment) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyRuleProfileEnrollment) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyRuleProfileEnrollment) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyRuleProfileEnrollment) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyRuleProfileEnrollment) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyRuleProfileEnrollment) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyRuleProfileEnrollment) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyRuleProfileEnrollment) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyRuleProfileEnrollment) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyRuleProfileEnrollment) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyRuleProfileEnrollment) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := p.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (p *jsiiProxy_PolicyRuleProfileEnrollment) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyRuleProfileEnrollment) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := p.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		p,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyRuleProfileEnrollment) MoveFromId(id *string) {
	if err := p.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveFromId",
		[]interface{}{id},
	)
}

func (p *jsiiProxy_PolicyRuleProfileEnrollment) MoveTo(moveTarget *string, index interface{}) {
	if err := p.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (p *jsiiProxy_PolicyRuleProfileEnrollment) MoveToId(id *string) {
	if err := p.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveToId",
		[]interface{}{id},
	)
}

func (p *jsiiProxy_PolicyRuleProfileEnrollment) OverrideLogicalId(newLogicalId *string) {
	if err := p.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (p *jsiiProxy_PolicyRuleProfileEnrollment) PutProfileAttributes(value interface{}) {
	if err := p.validatePutProfileAttributesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putProfileAttributes",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PolicyRuleProfileEnrollment) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := p.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (p *jsiiProxy_PolicyRuleProfileEnrollment) ResetAccess() {
	_jsii_.InvokeVoid(
		p,
		"resetAccess",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyRuleProfileEnrollment) ResetEmailVerification() {
	_jsii_.InvokeVoid(
		p,
		"resetEmailVerification",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyRuleProfileEnrollment) ResetId() {
	_jsii_.InvokeVoid(
		p,
		"resetId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyRuleProfileEnrollment) ResetInlineHookId() {
	_jsii_.InvokeVoid(
		p,
		"resetInlineHookId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyRuleProfileEnrollment) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		p,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyRuleProfileEnrollment) ResetProfileAttributes() {
	_jsii_.InvokeVoid(
		p,
		"resetProfileAttributes",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyRuleProfileEnrollment) ResetProgressiveProfilingAction() {
	_jsii_.InvokeVoid(
		p,
		"resetProgressiveProfilingAction",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyRuleProfileEnrollment) ResetTargetGroupId() {
	_jsii_.InvokeVoid(
		p,
		"resetTargetGroupId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyRuleProfileEnrollment) ResetUiSchemaId() {
	_jsii_.InvokeVoid(
		p,
		"resetUiSchemaId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyRuleProfileEnrollment) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyRuleProfileEnrollment) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyRuleProfileEnrollment) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyRuleProfileEnrollment) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyRuleProfileEnrollment) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyRuleProfileEnrollment) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyRuleProfileEnrollment) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		p,
		"with",
		args,
		&returns,
	)

	return returns
}

