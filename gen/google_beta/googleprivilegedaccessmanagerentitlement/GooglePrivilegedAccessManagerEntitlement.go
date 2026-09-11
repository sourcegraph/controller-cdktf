package googleprivilegedaccessmanagerentitlement

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleprivilegedaccessmanagerentitlement/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_privileged_access_manager_entitlement google_privileged_access_manager_entitlement}.
type GooglePrivilegedAccessManagerEntitlement interface {
	cdktn.TerraformResource
	AdditionalNotificationTargets() GooglePrivilegedAccessManagerEntitlementAdditionalNotificationTargetsOutputReference
	AdditionalNotificationTargetsInput() *GooglePrivilegedAccessManagerEntitlementAdditionalNotificationTargets
	ApprovalWorkflow() GooglePrivilegedAccessManagerEntitlementApprovalWorkflowOutputReference
	ApprovalWorkflowInput() *GooglePrivilegedAccessManagerEntitlementApprovalWorkflow
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
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EligibleUsers() GooglePrivilegedAccessManagerEntitlementEligibleUsersList
	EligibleUsersInput() interface{}
	EntitlementId() *string
	SetEntitlementId(val *string)
	EntitlementIdInput() *string
	Etag() *string
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
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	MaxRequestDuration() *string
	SetMaxRequestDuration(val *string)
	MaxRequestDurationInput() *string
	Name() *string
	// The tree node.
	Node() constructs.Node
	Parent() *string
	SetParent(val *string)
	ParentInput() *string
	PrivilegedAccess() GooglePrivilegedAccessManagerEntitlementPrivilegedAccessOutputReference
	PrivilegedAccessInput() *GooglePrivilegedAccessManagerEntitlementPrivilegedAccess
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
	RequesterJustificationConfig() GooglePrivilegedAccessManagerEntitlementRequesterJustificationConfigOutputReference
	RequesterJustificationConfigInput() *GooglePrivilegedAccessManagerEntitlementRequesterJustificationConfig
	State() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GooglePrivilegedAccessManagerEntitlementTimeoutsOutputReference
	TimeoutsInput() interface{}
	UpdateTime() *string
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
	PutAdditionalNotificationTargets(value *GooglePrivilegedAccessManagerEntitlementAdditionalNotificationTargets)
	PutApprovalWorkflow(value *GooglePrivilegedAccessManagerEntitlementApprovalWorkflow)
	PutEligibleUsers(value interface{})
	PutPrivilegedAccess(value *GooglePrivilegedAccessManagerEntitlementPrivilegedAccess)
	PutRequesterJustificationConfig(value *GooglePrivilegedAccessManagerEntitlementRequesterJustificationConfig)
	PutTimeouts(value *GooglePrivilegedAccessManagerEntitlementTimeouts)
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
	ResetAdditionalNotificationTargets()
	ResetApprovalWorkflow()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
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

// The jsii proxy struct for GooglePrivilegedAccessManagerEntitlement
type jsiiProxy_GooglePrivilegedAccessManagerEntitlement struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerEntitlement) AdditionalNotificationTargets() GooglePrivilegedAccessManagerEntitlementAdditionalNotificationTargetsOutputReference {
	var returns GooglePrivilegedAccessManagerEntitlementAdditionalNotificationTargetsOutputReference
	_jsii_.Get(
		j,
		"additionalNotificationTargets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerEntitlement) AdditionalNotificationTargetsInput() *GooglePrivilegedAccessManagerEntitlementAdditionalNotificationTargets {
	var returns *GooglePrivilegedAccessManagerEntitlementAdditionalNotificationTargets
	_jsii_.Get(
		j,
		"additionalNotificationTargetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerEntitlement) ApprovalWorkflow() GooglePrivilegedAccessManagerEntitlementApprovalWorkflowOutputReference {
	var returns GooglePrivilegedAccessManagerEntitlementApprovalWorkflowOutputReference
	_jsii_.Get(
		j,
		"approvalWorkflow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerEntitlement) ApprovalWorkflowInput() *GooglePrivilegedAccessManagerEntitlementApprovalWorkflow {
	var returns *GooglePrivilegedAccessManagerEntitlementApprovalWorkflow
	_jsii_.Get(
		j,
		"approvalWorkflowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerEntitlement) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerEntitlement) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerEntitlement) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerEntitlement) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerEntitlement) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerEntitlement) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerEntitlement) EligibleUsers() GooglePrivilegedAccessManagerEntitlementEligibleUsersList {
	var returns GooglePrivilegedAccessManagerEntitlementEligibleUsersList
	_jsii_.Get(
		j,
		"eligibleUsers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerEntitlement) EligibleUsersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"eligibleUsersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerEntitlement) EntitlementId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entitlementId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerEntitlement) EntitlementIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entitlementIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerEntitlement) Etag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerEntitlement) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerEntitlement) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerEntitlement) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerEntitlement) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerEntitlement) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerEntitlement) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerEntitlement) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerEntitlement) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerEntitlement) MaxRequestDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxRequestDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerEntitlement) MaxRequestDurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxRequestDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerEntitlement) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerEntitlement) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerEntitlement) Parent() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerEntitlement) ParentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerEntitlement) PrivilegedAccess() GooglePrivilegedAccessManagerEntitlementPrivilegedAccessOutputReference {
	var returns GooglePrivilegedAccessManagerEntitlementPrivilegedAccessOutputReference
	_jsii_.Get(
		j,
		"privilegedAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerEntitlement) PrivilegedAccessInput() *GooglePrivilegedAccessManagerEntitlementPrivilegedAccess {
	var returns *GooglePrivilegedAccessManagerEntitlementPrivilegedAccess
	_jsii_.Get(
		j,
		"privilegedAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerEntitlement) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerEntitlement) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerEntitlement) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerEntitlement) RequesterJustificationConfig() GooglePrivilegedAccessManagerEntitlementRequesterJustificationConfigOutputReference {
	var returns GooglePrivilegedAccessManagerEntitlementRequesterJustificationConfigOutputReference
	_jsii_.Get(
		j,
		"requesterJustificationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerEntitlement) RequesterJustificationConfigInput() *GooglePrivilegedAccessManagerEntitlementRequesterJustificationConfig {
	var returns *GooglePrivilegedAccessManagerEntitlementRequesterJustificationConfig
	_jsii_.Get(
		j,
		"requesterJustificationConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerEntitlement) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerEntitlement) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerEntitlement) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerEntitlement) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerEntitlement) Timeouts() GooglePrivilegedAccessManagerEntitlementTimeoutsOutputReference {
	var returns GooglePrivilegedAccessManagerEntitlementTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerEntitlement) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerEntitlement) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_privileged_access_manager_entitlement google_privileged_access_manager_entitlement} Resource.
func NewGooglePrivilegedAccessManagerEntitlement(scope constructs.Construct, id *string, config *GooglePrivilegedAccessManagerEntitlementConfig) GooglePrivilegedAccessManagerEntitlement {
	_init_.Initialize()

	if err := validateNewGooglePrivilegedAccessManagerEntitlementParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GooglePrivilegedAccessManagerEntitlement{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googlePrivilegedAccessManagerEntitlement.GooglePrivilegedAccessManagerEntitlement",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_privileged_access_manager_entitlement google_privileged_access_manager_entitlement} Resource.
func NewGooglePrivilegedAccessManagerEntitlement_Override(g GooglePrivilegedAccessManagerEntitlement, scope constructs.Construct, id *string, config *GooglePrivilegedAccessManagerEntitlementConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googlePrivilegedAccessManagerEntitlement.GooglePrivilegedAccessManagerEntitlement",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerEntitlement)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerEntitlement)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerEntitlement)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerEntitlement)SetEntitlementId(val *string) {
	if err := j.validateSetEntitlementIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"entitlementId",
		val,
	)
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerEntitlement)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerEntitlement)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerEntitlement)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerEntitlement)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerEntitlement)SetMaxRequestDuration(val *string) {
	if err := j.validateSetMaxRequestDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxRequestDuration",
		val,
	)
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerEntitlement)SetParent(val *string) {
	if err := j.validateSetParentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parent",
		val,
	)
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerEntitlement)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerEntitlement)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTN code for importing a GooglePrivilegedAccessManagerEntitlement resource upon running "cdktn plan <stack-name>".
func GooglePrivilegedAccessManagerEntitlement_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateGooglePrivilegedAccessManagerEntitlement_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-google-beta.googlePrivilegedAccessManagerEntitlement.GooglePrivilegedAccessManagerEntitlement",
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
func GooglePrivilegedAccessManagerEntitlement_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGooglePrivilegedAccessManagerEntitlement_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google-beta.googlePrivilegedAccessManagerEntitlement.GooglePrivilegedAccessManagerEntitlement",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GooglePrivilegedAccessManagerEntitlement_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGooglePrivilegedAccessManagerEntitlement_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google-beta.googlePrivilegedAccessManagerEntitlement.GooglePrivilegedAccessManagerEntitlement",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GooglePrivilegedAccessManagerEntitlement_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGooglePrivilegedAccessManagerEntitlement_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google-beta.googlePrivilegedAccessManagerEntitlement.GooglePrivilegedAccessManagerEntitlement",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GooglePrivilegedAccessManagerEntitlement_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-google-beta.googlePrivilegedAccessManagerEntitlement.GooglePrivilegedAccessManagerEntitlement",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GooglePrivilegedAccessManagerEntitlement) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GooglePrivilegedAccessManagerEntitlement) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GooglePrivilegedAccessManagerEntitlement) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GooglePrivilegedAccessManagerEntitlement) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GooglePrivilegedAccessManagerEntitlement) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GooglePrivilegedAccessManagerEntitlement) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GooglePrivilegedAccessManagerEntitlement) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GooglePrivilegedAccessManagerEntitlement) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GooglePrivilegedAccessManagerEntitlement) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GooglePrivilegedAccessManagerEntitlement) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GooglePrivilegedAccessManagerEntitlement) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GooglePrivilegedAccessManagerEntitlement) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GooglePrivilegedAccessManagerEntitlement) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GooglePrivilegedAccessManagerEntitlement) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GooglePrivilegedAccessManagerEntitlement) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (g *jsiiProxy_GooglePrivilegedAccessManagerEntitlement) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GooglePrivilegedAccessManagerEntitlement) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GooglePrivilegedAccessManagerEntitlement) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GooglePrivilegedAccessManagerEntitlement) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GooglePrivilegedAccessManagerEntitlement) PutAdditionalNotificationTargets(value *GooglePrivilegedAccessManagerEntitlementAdditionalNotificationTargets) {
	if err := g.validatePutAdditionalNotificationTargetsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAdditionalNotificationTargets",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GooglePrivilegedAccessManagerEntitlement) PutApprovalWorkflow(value *GooglePrivilegedAccessManagerEntitlementApprovalWorkflow) {
	if err := g.validatePutApprovalWorkflowParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putApprovalWorkflow",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GooglePrivilegedAccessManagerEntitlement) PutEligibleUsers(value interface{}) {
	if err := g.validatePutEligibleUsersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putEligibleUsers",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GooglePrivilegedAccessManagerEntitlement) PutPrivilegedAccess(value *GooglePrivilegedAccessManagerEntitlementPrivilegedAccess) {
	if err := g.validatePutPrivilegedAccessParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPrivilegedAccess",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GooglePrivilegedAccessManagerEntitlement) PutRequesterJustificationConfig(value *GooglePrivilegedAccessManagerEntitlementRequesterJustificationConfig) {
	if err := g.validatePutRequesterJustificationConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRequesterJustificationConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GooglePrivilegedAccessManagerEntitlement) PutTimeouts(value *GooglePrivilegedAccessManagerEntitlementTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GooglePrivilegedAccessManagerEntitlement) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := g.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (g *jsiiProxy_GooglePrivilegedAccessManagerEntitlement) ResetAdditionalNotificationTargets() {
	_jsii_.InvokeVoid(
		g,
		"resetAdditionalNotificationTargets",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePrivilegedAccessManagerEntitlement) ResetApprovalWorkflow() {
	_jsii_.InvokeVoid(
		g,
		"resetApprovalWorkflow",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePrivilegedAccessManagerEntitlement) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePrivilegedAccessManagerEntitlement) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePrivilegedAccessManagerEntitlement) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePrivilegedAccessManagerEntitlement) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GooglePrivilegedAccessManagerEntitlement) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GooglePrivilegedAccessManagerEntitlement) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GooglePrivilegedAccessManagerEntitlement) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GooglePrivilegedAccessManagerEntitlement) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GooglePrivilegedAccessManagerEntitlement) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GooglePrivilegedAccessManagerEntitlement) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

