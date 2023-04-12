package policyruleprofileenrollment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/okta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/okta/policyruleprofileenrollment/internal"
)

// Represents a {@link https://www.terraform.io/docs/providers/okta/r/policy_rule_profile_enrollment okta_policy_rule_profile_enrollment}.
type PolicyRuleProfileEnrollment interface {
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
	EmailVerification() interface{}
	SetEmailVerification(val interface{})
	EmailVerificationInput() interface{}
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
	InlineHookId() *string
	SetInlineHookId(val *string)
	InlineHookIdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	// The tree node.
	Node() constructs.Node
	PolicyId() *string
	SetPolicyId(val *string)
	PolicyIdInput() *string
	ProfileAttributes() PolicyRuleProfileEnrollmentProfileAttributesList
	ProfileAttributesInput() interface{}
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
	TargetGroupId() *string
	SetTargetGroupId(val *string)
	TargetGroupIdInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
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
	PutProfileAttributes(value interface{})
	ResetAccess()
	ResetEmailVerification()
	ResetId()
	ResetInlineHookId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProfileAttributes()
	ResetTargetGroupId()
	ResetUiSchemaId()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for PolicyRuleProfileEnrollment
type jsiiProxy_PolicyRuleProfileEnrollment struct {
	internal.Type__cdktfTerraformResource
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

func (j *jsiiProxy_PolicyRuleProfileEnrollment) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
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

func (j *jsiiProxy_PolicyRuleProfileEnrollment) Count() *float64 {
	var returns *float64
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

func (j *jsiiProxy_PolicyRuleProfileEnrollment) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
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

func (j *jsiiProxy_PolicyRuleProfileEnrollment) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
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

func (j *jsiiProxy_PolicyRuleProfileEnrollment) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
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

func (j *jsiiProxy_PolicyRuleProfileEnrollment) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
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


// Create a new {@link https://www.terraform.io/docs/providers/okta/r/policy_rule_profile_enrollment okta_policy_rule_profile_enrollment} Resource.
func NewPolicyRuleProfileEnrollment(scope constructs.Construct, id *string, config *PolicyRuleProfileEnrollmentConfig) PolicyRuleProfileEnrollment {
	_init_.Initialize()

	if err := validateNewPolicyRuleProfileEnrollmentParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_PolicyRuleProfileEnrollment{}

	_jsii_.Create(
		"@cdktf/provider-okta.policyRuleProfileEnrollment.PolicyRuleProfileEnrollment",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/okta/r/policy_rule_profile_enrollment okta_policy_rule_profile_enrollment} Resource.
func NewPolicyRuleProfileEnrollment_Override(p PolicyRuleProfileEnrollment, scope constructs.Construct, id *string, config *PolicyRuleProfileEnrollmentConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-okta.policyRuleProfileEnrollment.PolicyRuleProfileEnrollment",
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

func (j *jsiiProxy_PolicyRuleProfileEnrollment)SetCount(val *float64) {
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

func (j *jsiiProxy_PolicyRuleProfileEnrollment)SetForEach(val cdktf.ITerraformIterator) {
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

func (j *jsiiProxy_PolicyRuleProfileEnrollment)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
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

func (j *jsiiProxy_PolicyRuleProfileEnrollment)SetProvider(val cdktf.TerraformProvider) {
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
		"@cdktf/provider-okta.policyRuleProfileEnrollment.PolicyRuleProfileEnrollment",
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
		"@cdktf/provider-okta.policyRuleProfileEnrollment.PolicyRuleProfileEnrollment",
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
		"@cdktf/provider-okta.policyRuleProfileEnrollment.PolicyRuleProfileEnrollment",
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
		"@cdktf/provider-okta.policyRuleProfileEnrollment.PolicyRuleProfileEnrollment",
		"tfResourceType",
		&returns,
	)
	return returns
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

func (p *jsiiProxy_PolicyRuleProfileEnrollment) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PolicyRuleProfileEnrollment) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

