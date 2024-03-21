package authserverpolicyrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/okta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/okta/authserverpolicyrule/internal"
)

// Represents a {@link https://registry.terraform.io/providers/okta/okta/3.41.0/docs/resources/auth_server_policy_rule okta_auth_server_policy_rule}.
type AuthServerPolicyRule interface {
	cdktf.TerraformResource
	AccessTokenLifetimeMinutes() *float64
	SetAccessTokenLifetimeMinutes(val *float64)
	AccessTokenLifetimeMinutesInput() *float64
	AuthServerId() *string
	SetAuthServerId(val *string)
	AuthServerIdInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
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
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GrantTypeWhitelist() *[]*string
	SetGrantTypeWhitelist(val *[]*string)
	GrantTypeWhitelistInput() *[]*string
	GroupBlacklist() *[]*string
	SetGroupBlacklist(val *[]*string)
	GroupBlacklistInput() *[]*string
	GroupWhitelist() *[]*string
	SetGroupWhitelist(val *[]*string)
	GroupWhitelistInput() *[]*string
	Id() *string
	SetId(val *string)
	IdInput() *string
	InlineHookId() *string
	SetInlineHookId(val *string)
	InlineHookIdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
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
	RefreshTokenLifetimeMinutes() *float64
	SetRefreshTokenLifetimeMinutes(val *float64)
	RefreshTokenLifetimeMinutesInput() *float64
	RefreshTokenWindowMinutes() *float64
	SetRefreshTokenWindowMinutes(val *float64)
	RefreshTokenWindowMinutesInput() *float64
	ScopeWhitelist() *[]*string
	SetScopeWhitelist(val *[]*string)
	ScopeWhitelistInput() *[]*string
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
	UserBlacklist() *[]*string
	SetUserBlacklist(val *[]*string)
	UserBlacklistInput() *[]*string
	UserWhitelist() *[]*string
	SetUserWhitelist(val *[]*string)
	UserWhitelistInput() *[]*string
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
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
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
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
	ResetAccessTokenLifetimeMinutes()
	ResetGroupBlacklist()
	ResetGroupWhitelist()
	ResetId()
	ResetInlineHookId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRefreshTokenLifetimeMinutes()
	ResetRefreshTokenWindowMinutes()
	ResetScopeWhitelist()
	ResetStatus()
	ResetType()
	ResetUserBlacklist()
	ResetUserWhitelist()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for AuthServerPolicyRule
type jsiiProxy_AuthServerPolicyRule struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_AuthServerPolicyRule) AccessTokenLifetimeMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"accessTokenLifetimeMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthServerPolicyRule) AccessTokenLifetimeMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"accessTokenLifetimeMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthServerPolicyRule) AuthServerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authServerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthServerPolicyRule) AuthServerIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authServerIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthServerPolicyRule) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthServerPolicyRule) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthServerPolicyRule) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthServerPolicyRule) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthServerPolicyRule) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthServerPolicyRule) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthServerPolicyRule) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthServerPolicyRule) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthServerPolicyRule) GrantTypeWhitelist() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"grantTypeWhitelist",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthServerPolicyRule) GrantTypeWhitelistInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"grantTypeWhitelistInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthServerPolicyRule) GroupBlacklist() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupBlacklist",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthServerPolicyRule) GroupBlacklistInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupBlacklistInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthServerPolicyRule) GroupWhitelist() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupWhitelist",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthServerPolicyRule) GroupWhitelistInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupWhitelistInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthServerPolicyRule) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthServerPolicyRule) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthServerPolicyRule) InlineHookId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inlineHookId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthServerPolicyRule) InlineHookIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inlineHookIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthServerPolicyRule) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthServerPolicyRule) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthServerPolicyRule) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthServerPolicyRule) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthServerPolicyRule) PolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthServerPolicyRule) PolicyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthServerPolicyRule) Priority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthServerPolicyRule) PriorityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthServerPolicyRule) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthServerPolicyRule) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthServerPolicyRule) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthServerPolicyRule) RefreshTokenLifetimeMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"refreshTokenLifetimeMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthServerPolicyRule) RefreshTokenLifetimeMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"refreshTokenLifetimeMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthServerPolicyRule) RefreshTokenWindowMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"refreshTokenWindowMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthServerPolicyRule) RefreshTokenWindowMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"refreshTokenWindowMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthServerPolicyRule) ScopeWhitelist() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"scopeWhitelist",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthServerPolicyRule) ScopeWhitelistInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"scopeWhitelistInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthServerPolicyRule) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthServerPolicyRule) StatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthServerPolicyRule) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthServerPolicyRule) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthServerPolicyRule) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthServerPolicyRule) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthServerPolicyRule) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthServerPolicyRule) UserBlacklist() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"userBlacklist",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthServerPolicyRule) UserBlacklistInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"userBlacklistInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthServerPolicyRule) UserWhitelist() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"userWhitelist",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthServerPolicyRule) UserWhitelistInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"userWhitelistInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/okta/okta/3.41.0/docs/resources/auth_server_policy_rule okta_auth_server_policy_rule} Resource.
func NewAuthServerPolicyRule(scope constructs.Construct, id *string, config *AuthServerPolicyRuleConfig) AuthServerPolicyRule {
	_init_.Initialize()

	if err := validateNewAuthServerPolicyRuleParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AuthServerPolicyRule{}

	_jsii_.Create(
		"@cdktf/provider-okta.authServerPolicyRule.AuthServerPolicyRule",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/okta/okta/3.41.0/docs/resources/auth_server_policy_rule okta_auth_server_policy_rule} Resource.
func NewAuthServerPolicyRule_Override(a AuthServerPolicyRule, scope constructs.Construct, id *string, config *AuthServerPolicyRuleConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-okta.authServerPolicyRule.AuthServerPolicyRule",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AuthServerPolicyRule)SetAccessTokenLifetimeMinutes(val *float64) {
	if err := j.validateSetAccessTokenLifetimeMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessTokenLifetimeMinutes",
		val,
	)
}

func (j *jsiiProxy_AuthServerPolicyRule)SetAuthServerId(val *string) {
	if err := j.validateSetAuthServerIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authServerId",
		val,
	)
}

func (j *jsiiProxy_AuthServerPolicyRule)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AuthServerPolicyRule)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AuthServerPolicyRule)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AuthServerPolicyRule)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AuthServerPolicyRule)SetGrantTypeWhitelist(val *[]*string) {
	if err := j.validateSetGrantTypeWhitelistParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"grantTypeWhitelist",
		val,
	)
}

func (j *jsiiProxy_AuthServerPolicyRule)SetGroupBlacklist(val *[]*string) {
	if err := j.validateSetGroupBlacklistParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupBlacklist",
		val,
	)
}

func (j *jsiiProxy_AuthServerPolicyRule)SetGroupWhitelist(val *[]*string) {
	if err := j.validateSetGroupWhitelistParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupWhitelist",
		val,
	)
}

func (j *jsiiProxy_AuthServerPolicyRule)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AuthServerPolicyRule)SetInlineHookId(val *string) {
	if err := j.validateSetInlineHookIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inlineHookId",
		val,
	)
}

func (j *jsiiProxy_AuthServerPolicyRule)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AuthServerPolicyRule)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AuthServerPolicyRule)SetPolicyId(val *string) {
	if err := j.validateSetPolicyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policyId",
		val,
	)
}

func (j *jsiiProxy_AuthServerPolicyRule)SetPriority(val *float64) {
	if err := j.validateSetPriorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"priority",
		val,
	)
}

func (j *jsiiProxy_AuthServerPolicyRule)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AuthServerPolicyRule)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AuthServerPolicyRule)SetRefreshTokenLifetimeMinutes(val *float64) {
	if err := j.validateSetRefreshTokenLifetimeMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"refreshTokenLifetimeMinutes",
		val,
	)
}

func (j *jsiiProxy_AuthServerPolicyRule)SetRefreshTokenWindowMinutes(val *float64) {
	if err := j.validateSetRefreshTokenWindowMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"refreshTokenWindowMinutes",
		val,
	)
}

func (j *jsiiProxy_AuthServerPolicyRule)SetScopeWhitelist(val *[]*string) {
	if err := j.validateSetScopeWhitelistParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scopeWhitelist",
		val,
	)
}

func (j *jsiiProxy_AuthServerPolicyRule)SetStatus(val *string) {
	if err := j.validateSetStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_AuthServerPolicyRule)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_AuthServerPolicyRule)SetUserBlacklist(val *[]*string) {
	if err := j.validateSetUserBlacklistParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userBlacklist",
		val,
	)
}

func (j *jsiiProxy_AuthServerPolicyRule)SetUserWhitelist(val *[]*string) {
	if err := j.validateSetUserWhitelistParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userWhitelist",
		val,
	)
}

// Generates CDKTF code for importing a AuthServerPolicyRule resource upon running "cdktf plan <stack-name>".
func AuthServerPolicyRule_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateAuthServerPolicyRule_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.authServerPolicyRule.AuthServerPolicyRule",
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
func AuthServerPolicyRule_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAuthServerPolicyRule_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.authServerPolicyRule.AuthServerPolicyRule",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AuthServerPolicyRule_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAuthServerPolicyRule_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.authServerPolicyRule.AuthServerPolicyRule",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AuthServerPolicyRule_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAuthServerPolicyRule_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.authServerPolicyRule.AuthServerPolicyRule",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AuthServerPolicyRule_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-okta.authServerPolicyRule.AuthServerPolicyRule",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AuthServerPolicyRule) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AuthServerPolicyRule) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AuthServerPolicyRule) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AuthServerPolicyRule) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AuthServerPolicyRule) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AuthServerPolicyRule) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AuthServerPolicyRule) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AuthServerPolicyRule) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AuthServerPolicyRule) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AuthServerPolicyRule) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AuthServerPolicyRule) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AuthServerPolicyRule) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AuthServerPolicyRule) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AuthServerPolicyRule) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AuthServerPolicyRule) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AuthServerPolicyRule) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AuthServerPolicyRule) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AuthServerPolicyRule) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AuthServerPolicyRule) ResetAccessTokenLifetimeMinutes() {
	_jsii_.InvokeVoid(
		a,
		"resetAccessTokenLifetimeMinutes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AuthServerPolicyRule) ResetGroupBlacklist() {
	_jsii_.InvokeVoid(
		a,
		"resetGroupBlacklist",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AuthServerPolicyRule) ResetGroupWhitelist() {
	_jsii_.InvokeVoid(
		a,
		"resetGroupWhitelist",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AuthServerPolicyRule) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AuthServerPolicyRule) ResetInlineHookId() {
	_jsii_.InvokeVoid(
		a,
		"resetInlineHookId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AuthServerPolicyRule) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AuthServerPolicyRule) ResetRefreshTokenLifetimeMinutes() {
	_jsii_.InvokeVoid(
		a,
		"resetRefreshTokenLifetimeMinutes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AuthServerPolicyRule) ResetRefreshTokenWindowMinutes() {
	_jsii_.InvokeVoid(
		a,
		"resetRefreshTokenWindowMinutes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AuthServerPolicyRule) ResetScopeWhitelist() {
	_jsii_.InvokeVoid(
		a,
		"resetScopeWhitelist",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AuthServerPolicyRule) ResetStatus() {
	_jsii_.InvokeVoid(
		a,
		"resetStatus",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AuthServerPolicyRule) ResetType() {
	_jsii_.InvokeVoid(
		a,
		"resetType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AuthServerPolicyRule) ResetUserBlacklist() {
	_jsii_.InvokeVoid(
		a,
		"resetUserBlacklist",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AuthServerPolicyRule) ResetUserWhitelist() {
	_jsii_.InvokeVoid(
		a,
		"resetUserWhitelist",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AuthServerPolicyRule) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AuthServerPolicyRule) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AuthServerPolicyRule) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AuthServerPolicyRule) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

