package policyrulepassword

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/okta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/okta/policyrulepassword/internal"
)

// Represents a {@link https://registry.terraform.io/providers/okta/okta/3.41.0/docs/resources/policy_rule_password okta_policy_rule_password}.
type PolicyRulePassword interface {
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
	Id() *string
	SetId(val *string)
	IdInput() *string
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
	PasswordChange() *string
	SetPasswordChange(val *string)
	PasswordChangeInput() *string
	PasswordReset() *string
	SetPasswordReset(val *string)
	PasswordResetInput() *string
	PasswordUnlock() *string
	SetPasswordUnlock(val *string)
	PasswordUnlockInput() *string
	Policyid() *string
	SetPolicyid(val *string)
	PolicyId() *string
	SetPolicyId(val *string)
	PolicyidInput() *string
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
	ResetId()
	ResetNetworkConnection()
	ResetNetworkExcludes()
	ResetNetworkIncludes()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPasswordChange()
	ResetPasswordReset()
	ResetPasswordUnlock()
	ResetPolicyid()
	ResetPolicyId()
	ResetPriority()
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

// The jsii proxy struct for PolicyRulePassword
type jsiiProxy_PolicyRulePassword struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_PolicyRulePassword) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRulePassword) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRulePassword) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRulePassword) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRulePassword) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRulePassword) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRulePassword) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRulePassword) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRulePassword) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRulePassword) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRulePassword) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRulePassword) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRulePassword) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRulePassword) NetworkConnection() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkConnection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRulePassword) NetworkConnectionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkConnectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRulePassword) NetworkExcludes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"networkExcludes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRulePassword) NetworkExcludesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"networkExcludesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRulePassword) NetworkIncludes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"networkIncludes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRulePassword) NetworkIncludesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"networkIncludesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRulePassword) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRulePassword) PasswordChange() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordChange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRulePassword) PasswordChangeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordChangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRulePassword) PasswordReset() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordReset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRulePassword) PasswordResetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordResetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRulePassword) PasswordUnlock() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordUnlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRulePassword) PasswordUnlockInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordUnlockInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRulePassword) Policyid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRulePassword) PolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRulePassword) PolicyidInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRulePassword) PolicyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRulePassword) Priority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRulePassword) PriorityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRulePassword) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRulePassword) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRulePassword) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRulePassword) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRulePassword) StatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRulePassword) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRulePassword) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRulePassword) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRulePassword) UsersExcluded() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"usersExcluded",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRulePassword) UsersExcludedInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"usersExcludedInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/okta/okta/3.41.0/docs/resources/policy_rule_password okta_policy_rule_password} Resource.
func NewPolicyRulePassword(scope constructs.Construct, id *string, config *PolicyRulePasswordConfig) PolicyRulePassword {
	_init_.Initialize()

	if err := validateNewPolicyRulePasswordParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_PolicyRulePassword{}

	_jsii_.Create(
		"@cdktf/provider-okta.policyRulePassword.PolicyRulePassword",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/okta/okta/3.41.0/docs/resources/policy_rule_password okta_policy_rule_password} Resource.
func NewPolicyRulePassword_Override(p PolicyRulePassword, scope constructs.Construct, id *string, config *PolicyRulePasswordConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-okta.policyRulePassword.PolicyRulePassword",
		[]interface{}{scope, id, config},
		p,
	)
}

func (j *jsiiProxy_PolicyRulePassword)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_PolicyRulePassword)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_PolicyRulePassword)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_PolicyRulePassword)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_PolicyRulePassword)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_PolicyRulePassword)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_PolicyRulePassword)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_PolicyRulePassword)SetNetworkConnection(val *string) {
	if err := j.validateSetNetworkConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkConnection",
		val,
	)
}

func (j *jsiiProxy_PolicyRulePassword)SetNetworkExcludes(val *[]*string) {
	if err := j.validateSetNetworkExcludesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkExcludes",
		val,
	)
}

func (j *jsiiProxy_PolicyRulePassword)SetNetworkIncludes(val *[]*string) {
	if err := j.validateSetNetworkIncludesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkIncludes",
		val,
	)
}

func (j *jsiiProxy_PolicyRulePassword)SetPasswordChange(val *string) {
	if err := j.validateSetPasswordChangeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passwordChange",
		val,
	)
}

func (j *jsiiProxy_PolicyRulePassword)SetPasswordReset(val *string) {
	if err := j.validateSetPasswordResetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passwordReset",
		val,
	)
}

func (j *jsiiProxy_PolicyRulePassword)SetPasswordUnlock(val *string) {
	if err := j.validateSetPasswordUnlockParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passwordUnlock",
		val,
	)
}

func (j *jsiiProxy_PolicyRulePassword)SetPolicyid(val *string) {
	if err := j.validateSetPolicyidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policyid",
		val,
	)
}

func (j *jsiiProxy_PolicyRulePassword)SetPolicyId(val *string) {
	if err := j.validateSetPolicyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policyId",
		val,
	)
}

func (j *jsiiProxy_PolicyRulePassword)SetPriority(val *float64) {
	if err := j.validateSetPriorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"priority",
		val,
	)
}

func (j *jsiiProxy_PolicyRulePassword)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_PolicyRulePassword)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_PolicyRulePassword)SetStatus(val *string) {
	if err := j.validateSetStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_PolicyRulePassword)SetUsersExcluded(val *[]*string) {
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
func PolicyRulePassword_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePolicyRulePassword_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.policyRulePassword.PolicyRulePassword",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func PolicyRulePassword_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePolicyRulePassword_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.policyRulePassword.PolicyRulePassword",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func PolicyRulePassword_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePolicyRulePassword_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.policyRulePassword.PolicyRulePassword",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func PolicyRulePassword_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-okta.policyRulePassword.PolicyRulePassword",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (p *jsiiProxy_PolicyRulePassword) AddOverride(path *string, value interface{}) {
	if err := p.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (p *jsiiProxy_PolicyRulePassword) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PolicyRulePassword) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PolicyRulePassword) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PolicyRulePassword) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PolicyRulePassword) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PolicyRulePassword) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PolicyRulePassword) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PolicyRulePassword) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PolicyRulePassword) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PolicyRulePassword) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PolicyRulePassword) OverrideLogicalId(newLogicalId *string) {
	if err := p.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (p *jsiiProxy_PolicyRulePassword) ResetId() {
	_jsii_.InvokeVoid(
		p,
		"resetId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyRulePassword) ResetNetworkConnection() {
	_jsii_.InvokeVoid(
		p,
		"resetNetworkConnection",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyRulePassword) ResetNetworkExcludes() {
	_jsii_.InvokeVoid(
		p,
		"resetNetworkExcludes",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyRulePassword) ResetNetworkIncludes() {
	_jsii_.InvokeVoid(
		p,
		"resetNetworkIncludes",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyRulePassword) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		p,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyRulePassword) ResetPasswordChange() {
	_jsii_.InvokeVoid(
		p,
		"resetPasswordChange",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyRulePassword) ResetPasswordReset() {
	_jsii_.InvokeVoid(
		p,
		"resetPasswordReset",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyRulePassword) ResetPasswordUnlock() {
	_jsii_.InvokeVoid(
		p,
		"resetPasswordUnlock",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyRulePassword) ResetPolicyid() {
	_jsii_.InvokeVoid(
		p,
		"resetPolicyid",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyRulePassword) ResetPolicyId() {
	_jsii_.InvokeVoid(
		p,
		"resetPolicyId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyRulePassword) ResetPriority() {
	_jsii_.InvokeVoid(
		p,
		"resetPriority",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyRulePassword) ResetStatus() {
	_jsii_.InvokeVoid(
		p,
		"resetStatus",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyRulePassword) ResetUsersExcluded() {
	_jsii_.InvokeVoid(
		p,
		"resetUsersExcluded",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyRulePassword) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyRulePassword) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyRulePassword) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyRulePassword) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

