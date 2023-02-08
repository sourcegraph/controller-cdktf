package appsignonpolicyrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/okta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/okta/appsignonpolicyrule/internal"
)

// Represents a {@link https://www.terraform.io/docs/providers/okta/r/app_signon_policy_rule okta_app_signon_policy_rule}.
type AppSignonPolicyRule interface {
	cdktf.TerraformResource
	Access() *string
	SetAccess(val *string)
	AccessInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
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
	Count() *float64
	// Experimental.
	SetCount(val *float64)
	CustomExpression() *string
	SetCustomExpression(val *string)
	CustomExpressionInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
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
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
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
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
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
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	ReAuthenticationFrequency() *string
	SetReAuthenticationFrequency(val *string)
	ReAuthenticationFrequencyInput() *string
	Status() *string
	SetStatus(val *string)
	StatusInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
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
	PutPlatformInclude(value interface{})
	ResetAccess()
	ResetConstraints()
	ResetCustomExpression()
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
	ResetStatus()
	ResetType()
	ResetUsersExcluded()
	ResetUsersIncluded()
	ResetUserTypesExcluded()
	ResetUserTypesIncluded()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for AppSignonPolicyRule
type jsiiProxy_AppSignonPolicyRule struct {
	internal.Type__cdktfTerraformResource
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

func (j *jsiiProxy_AppSignonPolicyRule) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
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

func (j *jsiiProxy_AppSignonPolicyRule) Count() *float64 {
	var returns *float64
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

func (j *jsiiProxy_AppSignonPolicyRule) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
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

func (j *jsiiProxy_AppSignonPolicyRule) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
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

func (j *jsiiProxy_AppSignonPolicyRule) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
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

func (j *jsiiProxy_AppSignonPolicyRule) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
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


// Create a new {@link https://www.terraform.io/docs/providers/okta/r/app_signon_policy_rule okta_app_signon_policy_rule} Resource.
func NewAppSignonPolicyRule(scope constructs.Construct, id *string, config *AppSignonPolicyRuleConfig) AppSignonPolicyRule {
	_init_.Initialize()

	if err := validateNewAppSignonPolicyRuleParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppSignonPolicyRule{}

	_jsii_.Create(
		"okta.appSignonPolicyRule.AppSignonPolicyRule",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/okta/r/app_signon_policy_rule okta_app_signon_policy_rule} Resource.
func NewAppSignonPolicyRule_Override(a AppSignonPolicyRule, scope constructs.Construct, id *string, config *AppSignonPolicyRuleConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"okta.appSignonPolicyRule.AppSignonPolicyRule",
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

func (j *jsiiProxy_AppSignonPolicyRule)SetCount(val *float64) {
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

func (j *jsiiProxy_AppSignonPolicyRule)SetForEach(val cdktf.ITerraformIterator) {
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

func (j *jsiiProxy_AppSignonPolicyRule)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
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

func (j *jsiiProxy_AppSignonPolicyRule)SetProvider(val cdktf.TerraformProvider) {
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
		"okta.appSignonPolicyRule.AppSignonPolicyRule",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AppSignonPolicyRule_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"okta.appSignonPolicyRule.AppSignonPolicyRule",
		"tfResourceType",
		&returns,
	)
	return returns
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

func (a *jsiiProxy_AppSignonPolicyRule) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

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

func (a *jsiiProxy_AppSignonPolicyRule) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
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

