package policyrulesignon

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/okta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/okta/policyrulesignon/internal"
)

// Represents a {@link https://www.terraform.io/docs/providers/okta/r/policy_rule_signon okta_policy_rule_signon}.
type PolicyRuleSignon interface {
	cdktf.TerraformResource
	Access() *string
	SetAccess(val *string)
	AccessInput() *string
	Authtype() *string
	SetAuthtype(val *string)
	AuthtypeInput() *string
	Behaviors() *[]*string
	SetBehaviors(val *[]*string)
	BehaviorsInput() *[]*string
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
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	FactorSequence() PolicyRuleSignonFactorSequenceList
	FactorSequenceInput() interface{}
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
	IdentityProvider() *string
	SetIdentityProvider(val *string)
	IdentityProviderIds() *[]*string
	SetIdentityProviderIds(val *[]*string)
	IdentityProviderIdsInput() *[]*string
	IdentityProviderInput() *string
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MfaLifetime() *float64
	SetMfaLifetime(val *float64)
	MfaLifetimeInput() *float64
	MfaPrompt() *string
	SetMfaPrompt(val *string)
	MfaPromptInput() *string
	MfaRememberDevice() interface{}
	SetMfaRememberDevice(val interface{})
	MfaRememberDeviceInput() interface{}
	MfaRequired() interface{}
	SetMfaRequired(val interface{})
	MfaRequiredInput() interface{}
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
	Policyid() *string
	SetPolicyid(val *string)
	PolicyId() *string
	SetPolicyId(val *string)
	PolicyidInput() *string
	PolicyIdInput() *string
	PrimaryFactor() *string
	SetPrimaryFactor(val *string)
	PrimaryFactorInput() *string
	Priority() *float64
	SetPriority(val *float64)
	PriorityInput() *float64
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
	RiscLevel() *string
	SetRiscLevel(val *string)
	RiscLevelInput() *string
	SessionIdle() *float64
	SetSessionIdle(val *float64)
	SessionIdleInput() *float64
	SessionLifetime() *float64
	SetSessionLifetime(val *float64)
	SessionLifetimeInput() *float64
	SessionPersistent() interface{}
	SetSessionPersistent(val interface{})
	SessionPersistentInput() interface{}
	Status() *string
	SetStatus(val *string)
	StatusInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	UsersExcluded() *[]*string
	SetUsersExcluded(val *[]*string)
	UsersExcludedInput() *[]*string
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
	PutFactorSequence(value interface{})
	ResetAccess()
	ResetAuthtype()
	ResetBehaviors()
	ResetFactorSequence()
	ResetId()
	ResetIdentityProvider()
	ResetIdentityProviderIds()
	ResetMfaLifetime()
	ResetMfaPrompt()
	ResetMfaRememberDevice()
	ResetMfaRequired()
	ResetNetworkConnection()
	ResetNetworkExcludes()
	ResetNetworkIncludes()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPolicyid()
	ResetPolicyId()
	ResetPrimaryFactor()
	ResetPriority()
	ResetRiscLevel()
	ResetSessionIdle()
	ResetSessionLifetime()
	ResetSessionPersistent()
	ResetStatus()
	ResetUsersExcluded()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for PolicyRuleSignon
type jsiiProxy_PolicyRuleSignon struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_PolicyRuleSignon) Access() *string {
	var returns *string
	_jsii_.Get(
		j,
		"access",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleSignon) AccessInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleSignon) Authtype() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authtype",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleSignon) AuthtypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authtypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleSignon) Behaviors() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"behaviors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleSignon) BehaviorsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"behaviorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleSignon) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleSignon) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleSignon) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleSignon) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleSignon) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleSignon) FactorSequence() PolicyRuleSignonFactorSequenceList {
	var returns PolicyRuleSignonFactorSequenceList
	_jsii_.Get(
		j,
		"factorSequence",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleSignon) FactorSequenceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"factorSequenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleSignon) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleSignon) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleSignon) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleSignon) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleSignon) IdentityProvider() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleSignon) IdentityProviderIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"identityProviderIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleSignon) IdentityProviderIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"identityProviderIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleSignon) IdentityProviderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityProviderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleSignon) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleSignon) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleSignon) MfaLifetime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"mfaLifetime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleSignon) MfaLifetimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"mfaLifetimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleSignon) MfaPrompt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mfaPrompt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleSignon) MfaPromptInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mfaPromptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleSignon) MfaRememberDevice() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mfaRememberDevice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleSignon) MfaRememberDeviceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mfaRememberDeviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleSignon) MfaRequired() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mfaRequired",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleSignon) MfaRequiredInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mfaRequiredInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleSignon) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleSignon) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleSignon) NetworkConnection() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkConnection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleSignon) NetworkConnectionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkConnectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleSignon) NetworkExcludes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"networkExcludes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleSignon) NetworkExcludesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"networkExcludesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleSignon) NetworkIncludes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"networkIncludes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleSignon) NetworkIncludesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"networkIncludesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleSignon) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleSignon) Policyid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleSignon) PolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleSignon) PolicyidInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleSignon) PolicyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleSignon) PrimaryFactor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryFactor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleSignon) PrimaryFactorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryFactorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleSignon) Priority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleSignon) PriorityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleSignon) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleSignon) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleSignon) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleSignon) RiscLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"riscLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleSignon) RiscLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"riscLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleSignon) SessionIdle() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sessionIdle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleSignon) SessionIdleInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sessionIdleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleSignon) SessionLifetime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sessionLifetime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleSignon) SessionLifetimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sessionLifetimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleSignon) SessionPersistent() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sessionPersistent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleSignon) SessionPersistentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sessionPersistentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleSignon) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleSignon) StatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleSignon) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleSignon) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleSignon) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleSignon) UsersExcluded() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"usersExcluded",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleSignon) UsersExcludedInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"usersExcludedInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/okta/r/policy_rule_signon okta_policy_rule_signon} Resource.
func NewPolicyRuleSignon(scope constructs.Construct, id *string, config *PolicyRuleSignonConfig) PolicyRuleSignon {
	_init_.Initialize()

	if err := validateNewPolicyRuleSignonParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_PolicyRuleSignon{}

	_jsii_.Create(
		"@cdktf/provider-okta.policyRuleSignon.PolicyRuleSignon",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/okta/r/policy_rule_signon okta_policy_rule_signon} Resource.
func NewPolicyRuleSignon_Override(p PolicyRuleSignon, scope constructs.Construct, id *string, config *PolicyRuleSignonConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-okta.policyRuleSignon.PolicyRuleSignon",
		[]interface{}{scope, id, config},
		p,
	)
}

func (j *jsiiProxy_PolicyRuleSignon)SetAccess(val *string) {
	if err := j.validateSetAccessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"access",
		val,
	)
}

func (j *jsiiProxy_PolicyRuleSignon)SetAuthtype(val *string) {
	if err := j.validateSetAuthtypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authtype",
		val,
	)
}

func (j *jsiiProxy_PolicyRuleSignon)SetBehaviors(val *[]*string) {
	if err := j.validateSetBehaviorsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"behaviors",
		val,
	)
}

func (j *jsiiProxy_PolicyRuleSignon)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_PolicyRuleSignon)SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_PolicyRuleSignon)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_PolicyRuleSignon)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_PolicyRuleSignon)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_PolicyRuleSignon)SetIdentityProvider(val *string) {
	if err := j.validateSetIdentityProviderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"identityProvider",
		val,
	)
}

func (j *jsiiProxy_PolicyRuleSignon)SetIdentityProviderIds(val *[]*string) {
	if err := j.validateSetIdentityProviderIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"identityProviderIds",
		val,
	)
}

func (j *jsiiProxy_PolicyRuleSignon)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_PolicyRuleSignon)SetMfaLifetime(val *float64) {
	if err := j.validateSetMfaLifetimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mfaLifetime",
		val,
	)
}

func (j *jsiiProxy_PolicyRuleSignon)SetMfaPrompt(val *string) {
	if err := j.validateSetMfaPromptParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mfaPrompt",
		val,
	)
}

func (j *jsiiProxy_PolicyRuleSignon)SetMfaRememberDevice(val interface{}) {
	if err := j.validateSetMfaRememberDeviceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mfaRememberDevice",
		val,
	)
}

func (j *jsiiProxy_PolicyRuleSignon)SetMfaRequired(val interface{}) {
	if err := j.validateSetMfaRequiredParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mfaRequired",
		val,
	)
}

func (j *jsiiProxy_PolicyRuleSignon)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_PolicyRuleSignon)SetNetworkConnection(val *string) {
	if err := j.validateSetNetworkConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkConnection",
		val,
	)
}

func (j *jsiiProxy_PolicyRuleSignon)SetNetworkExcludes(val *[]*string) {
	if err := j.validateSetNetworkExcludesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkExcludes",
		val,
	)
}

func (j *jsiiProxy_PolicyRuleSignon)SetNetworkIncludes(val *[]*string) {
	if err := j.validateSetNetworkIncludesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkIncludes",
		val,
	)
}

func (j *jsiiProxy_PolicyRuleSignon)SetPolicyid(val *string) {
	if err := j.validateSetPolicyidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policyid",
		val,
	)
}

func (j *jsiiProxy_PolicyRuleSignon)SetPolicyId(val *string) {
	if err := j.validateSetPolicyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policyId",
		val,
	)
}

func (j *jsiiProxy_PolicyRuleSignon)SetPrimaryFactor(val *string) {
	if err := j.validateSetPrimaryFactorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"primaryFactor",
		val,
	)
}

func (j *jsiiProxy_PolicyRuleSignon)SetPriority(val *float64) {
	if err := j.validateSetPriorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"priority",
		val,
	)
}

func (j *jsiiProxy_PolicyRuleSignon)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_PolicyRuleSignon)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_PolicyRuleSignon)SetRiscLevel(val *string) {
	if err := j.validateSetRiscLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"riscLevel",
		val,
	)
}

func (j *jsiiProxy_PolicyRuleSignon)SetSessionIdle(val *float64) {
	if err := j.validateSetSessionIdleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sessionIdle",
		val,
	)
}

func (j *jsiiProxy_PolicyRuleSignon)SetSessionLifetime(val *float64) {
	if err := j.validateSetSessionLifetimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sessionLifetime",
		val,
	)
}

func (j *jsiiProxy_PolicyRuleSignon)SetSessionPersistent(val interface{}) {
	if err := j.validateSetSessionPersistentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sessionPersistent",
		val,
	)
}

func (j *jsiiProxy_PolicyRuleSignon)SetStatus(val *string) {
	if err := j.validateSetStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_PolicyRuleSignon)SetUsersExcluded(val *[]*string) {
	if err := j.validateSetUsersExcludedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"usersExcluded",
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
func PolicyRuleSignon_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePolicyRuleSignon_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.policyRuleSignon.PolicyRuleSignon",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func PolicyRuleSignon_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePolicyRuleSignon_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.policyRuleSignon.PolicyRuleSignon",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func PolicyRuleSignon_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePolicyRuleSignon_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.policyRuleSignon.PolicyRuleSignon",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func PolicyRuleSignon_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-okta.policyRuleSignon.PolicyRuleSignon",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (p *jsiiProxy_PolicyRuleSignon) AddOverride(path *string, value interface{}) {
	if err := p.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (p *jsiiProxy_PolicyRuleSignon) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PolicyRuleSignon) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyRuleSignon) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PolicyRuleSignon) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PolicyRuleSignon) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PolicyRuleSignon) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PolicyRuleSignon) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PolicyRuleSignon) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PolicyRuleSignon) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PolicyRuleSignon) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyRuleSignon) OverrideLogicalId(newLogicalId *string) {
	if err := p.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (p *jsiiProxy_PolicyRuleSignon) PutFactorSequence(value interface{}) {
	if err := p.validatePutFactorSequenceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putFactorSequence",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PolicyRuleSignon) ResetAccess() {
	_jsii_.InvokeVoid(
		p,
		"resetAccess",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyRuleSignon) ResetAuthtype() {
	_jsii_.InvokeVoid(
		p,
		"resetAuthtype",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyRuleSignon) ResetBehaviors() {
	_jsii_.InvokeVoid(
		p,
		"resetBehaviors",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyRuleSignon) ResetFactorSequence() {
	_jsii_.InvokeVoid(
		p,
		"resetFactorSequence",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyRuleSignon) ResetId() {
	_jsii_.InvokeVoid(
		p,
		"resetId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyRuleSignon) ResetIdentityProvider() {
	_jsii_.InvokeVoid(
		p,
		"resetIdentityProvider",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyRuleSignon) ResetIdentityProviderIds() {
	_jsii_.InvokeVoid(
		p,
		"resetIdentityProviderIds",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyRuleSignon) ResetMfaLifetime() {
	_jsii_.InvokeVoid(
		p,
		"resetMfaLifetime",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyRuleSignon) ResetMfaPrompt() {
	_jsii_.InvokeVoid(
		p,
		"resetMfaPrompt",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyRuleSignon) ResetMfaRememberDevice() {
	_jsii_.InvokeVoid(
		p,
		"resetMfaRememberDevice",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyRuleSignon) ResetMfaRequired() {
	_jsii_.InvokeVoid(
		p,
		"resetMfaRequired",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyRuleSignon) ResetNetworkConnection() {
	_jsii_.InvokeVoid(
		p,
		"resetNetworkConnection",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyRuleSignon) ResetNetworkExcludes() {
	_jsii_.InvokeVoid(
		p,
		"resetNetworkExcludes",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyRuleSignon) ResetNetworkIncludes() {
	_jsii_.InvokeVoid(
		p,
		"resetNetworkIncludes",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyRuleSignon) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		p,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyRuleSignon) ResetPolicyid() {
	_jsii_.InvokeVoid(
		p,
		"resetPolicyid",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyRuleSignon) ResetPolicyId() {
	_jsii_.InvokeVoid(
		p,
		"resetPolicyId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyRuleSignon) ResetPrimaryFactor() {
	_jsii_.InvokeVoid(
		p,
		"resetPrimaryFactor",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyRuleSignon) ResetPriority() {
	_jsii_.InvokeVoid(
		p,
		"resetPriority",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyRuleSignon) ResetRiscLevel() {
	_jsii_.InvokeVoid(
		p,
		"resetRiscLevel",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyRuleSignon) ResetSessionIdle() {
	_jsii_.InvokeVoid(
		p,
		"resetSessionIdle",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyRuleSignon) ResetSessionLifetime() {
	_jsii_.InvokeVoid(
		p,
		"resetSessionLifetime",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyRuleSignon) ResetSessionPersistent() {
	_jsii_.InvokeVoid(
		p,
		"resetSessionPersistent",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyRuleSignon) ResetStatus() {
	_jsii_.InvokeVoid(
		p,
		"resetStatus",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyRuleSignon) ResetUsersExcluded() {
	_jsii_.InvokeVoid(
		p,
		"resetUsersExcluded",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyRuleSignon) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyRuleSignon) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyRuleSignon) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyRuleSignon) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

