package signonpolicyrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/okta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/okta/signonpolicyrule/internal"
)

// Represents a {@link https://www.terraform.io/docs/providers/okta/r/signon_policy_rule okta_signon_policy_rule}.
type SignonPolicyRule interface {
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
	FactorSequence() SignonPolicyRuleFactorSequenceList
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

// The jsii proxy struct for SignonPolicyRule
type jsiiProxy_SignonPolicyRule struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SignonPolicyRule) Access() *string {
	var returns *string
	_jsii_.Get(
		j,
		"access",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignonPolicyRule) AccessInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignonPolicyRule) Authtype() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authtype",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignonPolicyRule) AuthtypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authtypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignonPolicyRule) Behaviors() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"behaviors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignonPolicyRule) BehaviorsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"behaviorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignonPolicyRule) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignonPolicyRule) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignonPolicyRule) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignonPolicyRule) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignonPolicyRule) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignonPolicyRule) FactorSequence() SignonPolicyRuleFactorSequenceList {
	var returns SignonPolicyRuleFactorSequenceList
	_jsii_.Get(
		j,
		"factorSequence",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignonPolicyRule) FactorSequenceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"factorSequenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignonPolicyRule) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignonPolicyRule) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignonPolicyRule) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignonPolicyRule) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignonPolicyRule) IdentityProvider() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignonPolicyRule) IdentityProviderIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"identityProviderIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignonPolicyRule) IdentityProviderIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"identityProviderIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignonPolicyRule) IdentityProviderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityProviderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignonPolicyRule) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignonPolicyRule) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignonPolicyRule) MfaLifetime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"mfaLifetime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignonPolicyRule) MfaLifetimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"mfaLifetimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignonPolicyRule) MfaPrompt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mfaPrompt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignonPolicyRule) MfaPromptInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mfaPromptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignonPolicyRule) MfaRememberDevice() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mfaRememberDevice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignonPolicyRule) MfaRememberDeviceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mfaRememberDeviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignonPolicyRule) MfaRequired() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mfaRequired",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignonPolicyRule) MfaRequiredInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mfaRequiredInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignonPolicyRule) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignonPolicyRule) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignonPolicyRule) NetworkConnection() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkConnection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignonPolicyRule) NetworkConnectionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkConnectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignonPolicyRule) NetworkExcludes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"networkExcludes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignonPolicyRule) NetworkExcludesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"networkExcludesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignonPolicyRule) NetworkIncludes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"networkIncludes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignonPolicyRule) NetworkIncludesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"networkIncludesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignonPolicyRule) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignonPolicyRule) Policyid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignonPolicyRule) PolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignonPolicyRule) PolicyidInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignonPolicyRule) PolicyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignonPolicyRule) PrimaryFactor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryFactor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignonPolicyRule) PrimaryFactorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryFactorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignonPolicyRule) Priority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignonPolicyRule) PriorityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignonPolicyRule) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignonPolicyRule) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignonPolicyRule) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignonPolicyRule) RiscLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"riscLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignonPolicyRule) RiscLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"riscLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignonPolicyRule) SessionIdle() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sessionIdle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignonPolicyRule) SessionIdleInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sessionIdleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignonPolicyRule) SessionLifetime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sessionLifetime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignonPolicyRule) SessionLifetimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sessionLifetimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignonPolicyRule) SessionPersistent() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sessionPersistent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignonPolicyRule) SessionPersistentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sessionPersistentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignonPolicyRule) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignonPolicyRule) StatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignonPolicyRule) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignonPolicyRule) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignonPolicyRule) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignonPolicyRule) UsersExcluded() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"usersExcluded",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignonPolicyRule) UsersExcludedInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"usersExcludedInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/okta/r/signon_policy_rule okta_signon_policy_rule} Resource.
func NewSignonPolicyRule(scope constructs.Construct, id *string, config *SignonPolicyRuleConfig) SignonPolicyRule {
	_init_.Initialize()

	if err := validateNewSignonPolicyRuleParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_SignonPolicyRule{}

	_jsii_.Create(
		"okta.signonPolicyRule.SignonPolicyRule",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/okta/r/signon_policy_rule okta_signon_policy_rule} Resource.
func NewSignonPolicyRule_Override(s SignonPolicyRule, scope constructs.Construct, id *string, config *SignonPolicyRuleConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"okta.signonPolicyRule.SignonPolicyRule",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SignonPolicyRule)SetAccess(val *string) {
	if err := j.validateSetAccessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"access",
		val,
	)
}

func (j *jsiiProxy_SignonPolicyRule)SetAuthtype(val *string) {
	if err := j.validateSetAuthtypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authtype",
		val,
	)
}

func (j *jsiiProxy_SignonPolicyRule)SetBehaviors(val *[]*string) {
	if err := j.validateSetBehaviorsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"behaviors",
		val,
	)
}

func (j *jsiiProxy_SignonPolicyRule)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_SignonPolicyRule)SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SignonPolicyRule)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SignonPolicyRule)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_SignonPolicyRule)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_SignonPolicyRule)SetIdentityProvider(val *string) {
	if err := j.validateSetIdentityProviderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"identityProvider",
		val,
	)
}

func (j *jsiiProxy_SignonPolicyRule)SetIdentityProviderIds(val *[]*string) {
	if err := j.validateSetIdentityProviderIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"identityProviderIds",
		val,
	)
}

func (j *jsiiProxy_SignonPolicyRule)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SignonPolicyRule)SetMfaLifetime(val *float64) {
	if err := j.validateSetMfaLifetimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mfaLifetime",
		val,
	)
}

func (j *jsiiProxy_SignonPolicyRule)SetMfaPrompt(val *string) {
	if err := j.validateSetMfaPromptParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mfaPrompt",
		val,
	)
}

func (j *jsiiProxy_SignonPolicyRule)SetMfaRememberDevice(val interface{}) {
	if err := j.validateSetMfaRememberDeviceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mfaRememberDevice",
		val,
	)
}

func (j *jsiiProxy_SignonPolicyRule)SetMfaRequired(val interface{}) {
	if err := j.validateSetMfaRequiredParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mfaRequired",
		val,
	)
}

func (j *jsiiProxy_SignonPolicyRule)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_SignonPolicyRule)SetNetworkConnection(val *string) {
	if err := j.validateSetNetworkConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkConnection",
		val,
	)
}

func (j *jsiiProxy_SignonPolicyRule)SetNetworkExcludes(val *[]*string) {
	if err := j.validateSetNetworkExcludesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkExcludes",
		val,
	)
}

func (j *jsiiProxy_SignonPolicyRule)SetNetworkIncludes(val *[]*string) {
	if err := j.validateSetNetworkIncludesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkIncludes",
		val,
	)
}

func (j *jsiiProxy_SignonPolicyRule)SetPolicyid(val *string) {
	if err := j.validateSetPolicyidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policyid",
		val,
	)
}

func (j *jsiiProxy_SignonPolicyRule)SetPolicyId(val *string) {
	if err := j.validateSetPolicyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policyId",
		val,
	)
}

func (j *jsiiProxy_SignonPolicyRule)SetPrimaryFactor(val *string) {
	if err := j.validateSetPrimaryFactorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"primaryFactor",
		val,
	)
}

func (j *jsiiProxy_SignonPolicyRule)SetPriority(val *float64) {
	if err := j.validateSetPriorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"priority",
		val,
	)
}

func (j *jsiiProxy_SignonPolicyRule)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SignonPolicyRule)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_SignonPolicyRule)SetRiscLevel(val *string) {
	if err := j.validateSetRiscLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"riscLevel",
		val,
	)
}

func (j *jsiiProxy_SignonPolicyRule)SetSessionIdle(val *float64) {
	if err := j.validateSetSessionIdleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sessionIdle",
		val,
	)
}

func (j *jsiiProxy_SignonPolicyRule)SetSessionLifetime(val *float64) {
	if err := j.validateSetSessionLifetimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sessionLifetime",
		val,
	)
}

func (j *jsiiProxy_SignonPolicyRule)SetSessionPersistent(val interface{}) {
	if err := j.validateSetSessionPersistentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sessionPersistent",
		val,
	)
}

func (j *jsiiProxy_SignonPolicyRule)SetStatus(val *string) {
	if err := j.validateSetStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_SignonPolicyRule)SetUsersExcluded(val *[]*string) {
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
func SignonPolicyRule_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSignonPolicyRule_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"okta.signonPolicyRule.SignonPolicyRule",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SignonPolicyRule_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"okta.signonPolicyRule.SignonPolicyRule",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_SignonPolicyRule) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_SignonPolicyRule) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SignonPolicyRule) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SignonPolicyRule) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SignonPolicyRule) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SignonPolicyRule) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SignonPolicyRule) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SignonPolicyRule) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SignonPolicyRule) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SignonPolicyRule) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SignonPolicyRule) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SignonPolicyRule) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SignonPolicyRule) PutFactorSequence(value interface{}) {
	if err := s.validatePutFactorSequenceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putFactorSequence",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SignonPolicyRule) ResetAccess() {
	_jsii_.InvokeVoid(
		s,
		"resetAccess",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SignonPolicyRule) ResetAuthtype() {
	_jsii_.InvokeVoid(
		s,
		"resetAuthtype",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SignonPolicyRule) ResetBehaviors() {
	_jsii_.InvokeVoid(
		s,
		"resetBehaviors",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SignonPolicyRule) ResetFactorSequence() {
	_jsii_.InvokeVoid(
		s,
		"resetFactorSequence",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SignonPolicyRule) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SignonPolicyRule) ResetIdentityProvider() {
	_jsii_.InvokeVoid(
		s,
		"resetIdentityProvider",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SignonPolicyRule) ResetIdentityProviderIds() {
	_jsii_.InvokeVoid(
		s,
		"resetIdentityProviderIds",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SignonPolicyRule) ResetMfaLifetime() {
	_jsii_.InvokeVoid(
		s,
		"resetMfaLifetime",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SignonPolicyRule) ResetMfaPrompt() {
	_jsii_.InvokeVoid(
		s,
		"resetMfaPrompt",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SignonPolicyRule) ResetMfaRememberDevice() {
	_jsii_.InvokeVoid(
		s,
		"resetMfaRememberDevice",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SignonPolicyRule) ResetMfaRequired() {
	_jsii_.InvokeVoid(
		s,
		"resetMfaRequired",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SignonPolicyRule) ResetNetworkConnection() {
	_jsii_.InvokeVoid(
		s,
		"resetNetworkConnection",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SignonPolicyRule) ResetNetworkExcludes() {
	_jsii_.InvokeVoid(
		s,
		"resetNetworkExcludes",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SignonPolicyRule) ResetNetworkIncludes() {
	_jsii_.InvokeVoid(
		s,
		"resetNetworkIncludes",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SignonPolicyRule) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SignonPolicyRule) ResetPolicyid() {
	_jsii_.InvokeVoid(
		s,
		"resetPolicyid",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SignonPolicyRule) ResetPolicyId() {
	_jsii_.InvokeVoid(
		s,
		"resetPolicyId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SignonPolicyRule) ResetPrimaryFactor() {
	_jsii_.InvokeVoid(
		s,
		"resetPrimaryFactor",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SignonPolicyRule) ResetPriority() {
	_jsii_.InvokeVoid(
		s,
		"resetPriority",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SignonPolicyRule) ResetRiscLevel() {
	_jsii_.InvokeVoid(
		s,
		"resetRiscLevel",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SignonPolicyRule) ResetSessionIdle() {
	_jsii_.InvokeVoid(
		s,
		"resetSessionIdle",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SignonPolicyRule) ResetSessionLifetime() {
	_jsii_.InvokeVoid(
		s,
		"resetSessionLifetime",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SignonPolicyRule) ResetSessionPersistent() {
	_jsii_.InvokeVoid(
		s,
		"resetSessionPersistent",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SignonPolicyRule) ResetStatus() {
	_jsii_.InvokeVoid(
		s,
		"resetStatus",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SignonPolicyRule) ResetUsersExcluded() {
	_jsii_.InvokeVoid(
		s,
		"resetUsersExcluded",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SignonPolicyRule) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SignonPolicyRule) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SignonPolicyRule) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SignonPolicyRule) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

