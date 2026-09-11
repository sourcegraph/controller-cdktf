package appsignonpolicyrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/okta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/okta/appsignonpolicyrule/internal"
)

// Represents a {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_signon_policy_rule okta_app_signon_policy_rule}.
type AppSignonPolicyRule interface {
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
	Constraints() *[]*string
	SetConstraints(val *[]*string)
	ConstraintsInput() *[]*string
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CustomExpression() *string
	SetCustomExpression(val *string)
	CustomExpressionInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DeviceAssurancesIncluded() *[]*string
	SetDeviceAssurancesIncluded(val *[]*string)
	DeviceAssurancesIncludedInput() *[]*string
	DeviceIsManaged() interface{}
	SetDeviceIsManaged(val interface{})
	DeviceIsManagedInput() interface{}
	DeviceIsRegistered() interface{}
	SetDeviceIsRegistered(val interface{})
	DeviceIsRegisteredInput() interface{}
	FactorMode() *string
	SetFactorMode(val *string)
	FactorModeInput() *string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GroupsExcluded() *[]*string
	SetGroupsExcluded(val *[]*string)
	GroupsExcludedInput() *[]*string
	GroupsIncluded() *[]*string
	SetGroupsIncluded(val *[]*string)
	GroupsIncludedInput() *[]*string
	Id() *string
	SetId(val *string)
	IdInput() *string
	InactivityPeriod() *string
	SetInactivityPeriod(val *string)
	InactivityPeriodInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	NetworkConnection() *string
	SetNetworkConnection(val *string)
	NetworkConnectionInput() *string
	NetworkExcludes() *[]*string
	SetNetworkExcludes(val *[]*string)
	NetworkExcludesInput() *[]*string
	NetworkIncludes() *[]*string
	SetNetworkIncludes(val *[]*string)
	NetworkIncludesInput() *[]*string
	// The tree node.
	Node() constructs.Node
	PlatformInclude() AppSignonPolicyRulePlatformIncludeList
	PlatformIncludeInput() interface{}
	PolicyId() *string
	SetPolicyId(val *string)
	PolicyIdInput() *string
	Priority() *float64
	SetPriority(val *float64)
	PriorityInput() *float64
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
	ReAuthenticationFrequency() *string
	SetReAuthenticationFrequency(val *string)
	ReAuthenticationFrequencyInput() *string
	RiskScore() *string
	SetRiskScore(val *string)
	RiskScoreInput() *string
	Status() *string
	SetStatus(val *string)
	StatusInput() *string
	SystemAttribute() cdktn.IResolvable
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Type() *string
	SetType(val *string)
	TypeInput() *string
	UsersExcluded() *[]*string
	SetUsersExcluded(val *[]*string)
	UsersExcludedInput() *[]*string
	UsersIncluded() *[]*string
	SetUsersIncluded(val *[]*string)
	UsersIncludedInput() *[]*string
	UserTypesExcluded() *[]*string
	SetUserTypesExcluded(val *[]*string)
	UserTypesExcludedInput() *[]*string
	UserTypesIncluded() *[]*string
	SetUserTypesIncluded(val *[]*string)
	UserTypesIncludedInput() *[]*string
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
	PutPlatformInclude(value interface{})
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
	ResetConstraints()
	ResetCustomExpression()
	ResetDeviceAssurancesIncluded()
	ResetDeviceIsManaged()
	ResetDeviceIsRegistered()
	ResetFactorMode()
	ResetGroupsExcluded()
	ResetGroupsIncluded()
	ResetId()
	ResetInactivityPeriod()
	ResetNetworkConnection()
	ResetNetworkExcludes()
	ResetNetworkIncludes()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPlatformInclude()
	ResetPriority()
	ResetReAuthenticationFrequency()
	ResetRiskScore()
	ResetStatus()
	ResetType()
	ResetUsersExcluded()
	ResetUsersIncluded()
	ResetUserTypesExcluded()
	ResetUserTypesIncluded()
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

// The jsii proxy struct for AppSignonPolicyRule
type jsiiProxy_AppSignonPolicyRule struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AppSignonPolicyRule) Access() *string {
	var returns *string
	_jsii_.Get(
		j,
		"access",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSignonPolicyRule) AccessInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSignonPolicyRule) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSignonPolicyRule) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSignonPolicyRule) Constraints() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"constraints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSignonPolicyRule) ConstraintsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"constraintsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSignonPolicyRule) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSignonPolicyRule) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSignonPolicyRule) CustomExpression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSignonPolicyRule) CustomExpressionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customExpressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSignonPolicyRule) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSignonPolicyRule) DeviceAssurancesIncluded() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"deviceAssurancesIncluded",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSignonPolicyRule) DeviceAssurancesIncludedInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"deviceAssurancesIncludedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSignonPolicyRule) DeviceIsManaged() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deviceIsManaged",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSignonPolicyRule) DeviceIsManagedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deviceIsManagedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSignonPolicyRule) DeviceIsRegistered() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deviceIsRegistered",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSignonPolicyRule) DeviceIsRegisteredInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deviceIsRegisteredInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSignonPolicyRule) FactorMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"factorMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSignonPolicyRule) FactorModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"factorModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSignonPolicyRule) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSignonPolicyRule) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSignonPolicyRule) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSignonPolicyRule) GroupsExcluded() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupsExcluded",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSignonPolicyRule) GroupsExcludedInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupsExcludedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSignonPolicyRule) GroupsIncluded() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupsIncluded",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSignonPolicyRule) GroupsIncludedInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupsIncludedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSignonPolicyRule) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSignonPolicyRule) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSignonPolicyRule) InactivityPeriod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inactivityPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSignonPolicyRule) InactivityPeriodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inactivityPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSignonPolicyRule) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSignonPolicyRule) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSignonPolicyRule) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSignonPolicyRule) NetworkConnection() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkConnection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSignonPolicyRule) NetworkConnectionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkConnectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSignonPolicyRule) NetworkExcludes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"networkExcludes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSignonPolicyRule) NetworkExcludesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"networkExcludesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSignonPolicyRule) NetworkIncludes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"networkIncludes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSignonPolicyRule) NetworkIncludesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"networkIncludesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSignonPolicyRule) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSignonPolicyRule) PlatformInclude() AppSignonPolicyRulePlatformIncludeList {
	var returns AppSignonPolicyRulePlatformIncludeList
	_jsii_.Get(
		j,
		"platformInclude",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSignonPolicyRule) PlatformIncludeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"platformIncludeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSignonPolicyRule) PolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSignonPolicyRule) PolicyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSignonPolicyRule) Priority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSignonPolicyRule) PriorityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSignonPolicyRule) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSignonPolicyRule) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSignonPolicyRule) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSignonPolicyRule) ReAuthenticationFrequency() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reAuthenticationFrequency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSignonPolicyRule) ReAuthenticationFrequencyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reAuthenticationFrequencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSignonPolicyRule) RiskScore() *string {
	var returns *string
	_jsii_.Get(
		j,
		"riskScore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSignonPolicyRule) RiskScoreInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"riskScoreInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSignonPolicyRule) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSignonPolicyRule) StatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSignonPolicyRule) SystemAttribute() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"systemAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSignonPolicyRule) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSignonPolicyRule) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSignonPolicyRule) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSignonPolicyRule) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSignonPolicyRule) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSignonPolicyRule) UsersExcluded() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"usersExcluded",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSignonPolicyRule) UsersExcludedInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"usersExcludedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSignonPolicyRule) UsersIncluded() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"usersIncluded",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSignonPolicyRule) UsersIncludedInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"usersIncludedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSignonPolicyRule) UserTypesExcluded() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"userTypesExcluded",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSignonPolicyRule) UserTypesExcludedInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"userTypesExcludedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSignonPolicyRule) UserTypesIncluded() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"userTypesIncluded",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSignonPolicyRule) UserTypesIncludedInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"userTypesIncludedInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_signon_policy_rule okta_app_signon_policy_rule} Resource.
func NewAppSignonPolicyRule(scope constructs.Construct, id *string, config *AppSignonPolicyRuleConfig) AppSignonPolicyRule {
	_init_.Initialize()

	if err := validateNewAppSignonPolicyRuleParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppSignonPolicyRule{}

	_jsii_.Create(
		"@cdktn/provider-okta.appSignonPolicyRule.AppSignonPolicyRule",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_signon_policy_rule okta_app_signon_policy_rule} Resource.
func NewAppSignonPolicyRule_Override(a AppSignonPolicyRule, scope constructs.Construct, id *string, config *AppSignonPolicyRuleConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-okta.appSignonPolicyRule.AppSignonPolicyRule",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AppSignonPolicyRule)SetAccess(val *string) {
	if err := j.validateSetAccessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"access",
		val,
	)
}

func (j *jsiiProxy_AppSignonPolicyRule)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AppSignonPolicyRule)SetConstraints(val *[]*string) {
	if err := j.validateSetConstraintsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"constraints",
		val,
	)
}

func (j *jsiiProxy_AppSignonPolicyRule)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AppSignonPolicyRule)SetCustomExpression(val *string) {
	if err := j.validateSetCustomExpressionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customExpression",
		val,
	)
}

func (j *jsiiProxy_AppSignonPolicyRule)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AppSignonPolicyRule)SetDeviceAssurancesIncluded(val *[]*string) {
	if err := j.validateSetDeviceAssurancesIncludedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deviceAssurancesIncluded",
		val,
	)
}

func (j *jsiiProxy_AppSignonPolicyRule)SetDeviceIsManaged(val interface{}) {
	if err := j.validateSetDeviceIsManagedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deviceIsManaged",
		val,
	)
}

func (j *jsiiProxy_AppSignonPolicyRule)SetDeviceIsRegistered(val interface{}) {
	if err := j.validateSetDeviceIsRegisteredParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deviceIsRegistered",
		val,
	)
}

func (j *jsiiProxy_AppSignonPolicyRule)SetFactorMode(val *string) {
	if err := j.validateSetFactorModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"factorMode",
		val,
	)
}

func (j *jsiiProxy_AppSignonPolicyRule)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AppSignonPolicyRule)SetGroupsExcluded(val *[]*string) {
	if err := j.validateSetGroupsExcludedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupsExcluded",
		val,
	)
}

func (j *jsiiProxy_AppSignonPolicyRule)SetGroupsIncluded(val *[]*string) {
	if err := j.validateSetGroupsIncludedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupsIncluded",
		val,
	)
}

func (j *jsiiProxy_AppSignonPolicyRule)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AppSignonPolicyRule)SetInactivityPeriod(val *string) {
	if err := j.validateSetInactivityPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inactivityPeriod",
		val,
	)
}

func (j *jsiiProxy_AppSignonPolicyRule)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AppSignonPolicyRule)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AppSignonPolicyRule)SetNetworkConnection(val *string) {
	if err := j.validateSetNetworkConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkConnection",
		val,
	)
}

func (j *jsiiProxy_AppSignonPolicyRule)SetNetworkExcludes(val *[]*string) {
	if err := j.validateSetNetworkExcludesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkExcludes",
		val,
	)
}

func (j *jsiiProxy_AppSignonPolicyRule)SetNetworkIncludes(val *[]*string) {
	if err := j.validateSetNetworkIncludesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkIncludes",
		val,
	)
}

func (j *jsiiProxy_AppSignonPolicyRule)SetPolicyId(val *string) {
	if err := j.validateSetPolicyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policyId",
		val,
	)
}

func (j *jsiiProxy_AppSignonPolicyRule)SetPriority(val *float64) {
	if err := j.validateSetPriorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"priority",
		val,
	)
}

func (j *jsiiProxy_AppSignonPolicyRule)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AppSignonPolicyRule)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AppSignonPolicyRule)SetReAuthenticationFrequency(val *string) {
	if err := j.validateSetReAuthenticationFrequencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reAuthenticationFrequency",
		val,
	)
}

func (j *jsiiProxy_AppSignonPolicyRule)SetRiskScore(val *string) {
	if err := j.validateSetRiskScoreParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"riskScore",
		val,
	)
}

func (j *jsiiProxy_AppSignonPolicyRule)SetStatus(val *string) {
	if err := j.validateSetStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_AppSignonPolicyRule)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_AppSignonPolicyRule)SetUsersExcluded(val *[]*string) {
	if err := j.validateSetUsersExcludedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"usersExcluded",
		val,
	)
}

func (j *jsiiProxy_AppSignonPolicyRule)SetUsersIncluded(val *[]*string) {
	if err := j.validateSetUsersIncludedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"usersIncluded",
		val,
	)
}

func (j *jsiiProxy_AppSignonPolicyRule)SetUserTypesExcluded(val *[]*string) {
	if err := j.validateSetUserTypesExcludedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userTypesExcluded",
		val,
	)
}

func (j *jsiiProxy_AppSignonPolicyRule)SetUserTypesIncluded(val *[]*string) {
	if err := j.validateSetUserTypesIncludedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userTypesIncluded",
		val,
	)
}

// Generates CDKTN code for importing a AppSignonPolicyRule resource upon running "cdktn plan <stack-name>".
func AppSignonPolicyRule_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAppSignonPolicyRule_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-okta.appSignonPolicyRule.AppSignonPolicyRule",
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
func AppSignonPolicyRule_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAppSignonPolicyRule_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-okta.appSignonPolicyRule.AppSignonPolicyRule",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AppSignonPolicyRule_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAppSignonPolicyRule_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-okta.appSignonPolicyRule.AppSignonPolicyRule",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AppSignonPolicyRule_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAppSignonPolicyRule_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-okta.appSignonPolicyRule.AppSignonPolicyRule",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AppSignonPolicyRule_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-okta.appSignonPolicyRule.AppSignonPolicyRule",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AppSignonPolicyRule) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AppSignonPolicyRule) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AppSignonPolicyRule) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSignonPolicyRule) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSignonPolicyRule) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSignonPolicyRule) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSignonPolicyRule) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSignonPolicyRule) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSignonPolicyRule) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSignonPolicyRule) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSignonPolicyRule) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSignonPolicyRule) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSignonPolicyRule) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AppSignonPolicyRule) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSignonPolicyRule) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := a.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSignonPolicyRule) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AppSignonPolicyRule) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AppSignonPolicyRule) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AppSignonPolicyRule) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AppSignonPolicyRule) PutPlatformInclude(value interface{}) {
	if err := a.validatePutPlatformIncludeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPlatformInclude",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppSignonPolicyRule) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AppSignonPolicyRule) ResetAccess() {
	_jsii_.InvokeVoid(
		a,
		"resetAccess",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSignonPolicyRule) ResetConstraints() {
	_jsii_.InvokeVoid(
		a,
		"resetConstraints",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSignonPolicyRule) ResetCustomExpression() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomExpression",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSignonPolicyRule) ResetDeviceAssurancesIncluded() {
	_jsii_.InvokeVoid(
		a,
		"resetDeviceAssurancesIncluded",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSignonPolicyRule) ResetDeviceIsManaged() {
	_jsii_.InvokeVoid(
		a,
		"resetDeviceIsManaged",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSignonPolicyRule) ResetDeviceIsRegistered() {
	_jsii_.InvokeVoid(
		a,
		"resetDeviceIsRegistered",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSignonPolicyRule) ResetFactorMode() {
	_jsii_.InvokeVoid(
		a,
		"resetFactorMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSignonPolicyRule) ResetGroupsExcluded() {
	_jsii_.InvokeVoid(
		a,
		"resetGroupsExcluded",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSignonPolicyRule) ResetGroupsIncluded() {
	_jsii_.InvokeVoid(
		a,
		"resetGroupsIncluded",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSignonPolicyRule) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSignonPolicyRule) ResetInactivityPeriod() {
	_jsii_.InvokeVoid(
		a,
		"resetInactivityPeriod",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSignonPolicyRule) ResetNetworkConnection() {
	_jsii_.InvokeVoid(
		a,
		"resetNetworkConnection",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSignonPolicyRule) ResetNetworkExcludes() {
	_jsii_.InvokeVoid(
		a,
		"resetNetworkExcludes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSignonPolicyRule) ResetNetworkIncludes() {
	_jsii_.InvokeVoid(
		a,
		"resetNetworkIncludes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSignonPolicyRule) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSignonPolicyRule) ResetPlatformInclude() {
	_jsii_.InvokeVoid(
		a,
		"resetPlatformInclude",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSignonPolicyRule) ResetPriority() {
	_jsii_.InvokeVoid(
		a,
		"resetPriority",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSignonPolicyRule) ResetReAuthenticationFrequency() {
	_jsii_.InvokeVoid(
		a,
		"resetReAuthenticationFrequency",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSignonPolicyRule) ResetRiskScore() {
	_jsii_.InvokeVoid(
		a,
		"resetRiskScore",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSignonPolicyRule) ResetStatus() {
	_jsii_.InvokeVoid(
		a,
		"resetStatus",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSignonPolicyRule) ResetType() {
	_jsii_.InvokeVoid(
		a,
		"resetType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSignonPolicyRule) ResetUsersExcluded() {
	_jsii_.InvokeVoid(
		a,
		"resetUsersExcluded",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSignonPolicyRule) ResetUsersIncluded() {
	_jsii_.InvokeVoid(
		a,
		"resetUsersIncluded",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSignonPolicyRule) ResetUserTypesExcluded() {
	_jsii_.InvokeVoid(
		a,
		"resetUserTypesExcluded",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSignonPolicyRule) ResetUserTypesIncluded() {
	_jsii_.InvokeVoid(
		a,
		"resetUserTypesIncluded",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppSignonPolicyRule) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSignonPolicyRule) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSignonPolicyRule) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSignonPolicyRule) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSignonPolicyRule) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSignonPolicyRule) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSignonPolicyRule) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		a,
		"with",
		args,
		&returns,
	)

	return returns
}

