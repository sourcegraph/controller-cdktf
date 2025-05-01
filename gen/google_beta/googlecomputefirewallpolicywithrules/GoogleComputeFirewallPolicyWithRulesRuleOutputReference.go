package googlecomputefirewallpolicywithrules

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecomputefirewallpolicywithrules/internal"
)

type GoogleComputeFirewallPolicyWithRulesRuleOutputReference interface {
	cdktf.ComplexObject
	Action() *string
	SetAction(val *string)
	ActionInput() *string
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Direction() *string
	SetDirection(val *string)
	DirectionInput() *string
	Disabled() interface{}
	SetDisabled(val interface{})
	DisabledInput() interface{}
	EnableLogging() interface{}
	SetEnableLogging(val interface{})
	EnableLoggingInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Match() GoogleComputeFirewallPolicyWithRulesRuleMatchOutputReference
	MatchInput() *GoogleComputeFirewallPolicyWithRulesRuleMatch
	Priority() *float64
	SetPriority(val *float64)
	PriorityInput() *float64
	RuleName() *string
	SetRuleName(val *string)
	RuleNameInput() *string
	SecurityProfileGroup() *string
	SetSecurityProfileGroup(val *string)
	SecurityProfileGroupInput() *string
	TargetResources() *[]*string
	SetTargetResources(val *[]*string)
	TargetResourcesInput() *[]*string
	TargetServiceAccounts() *[]*string
	SetTargetServiceAccounts(val *[]*string)
	TargetServiceAccountsInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TlsInspect() interface{}
	SetTlsInspect(val interface{})
	TlsInspectInput() interface{}
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutMatch(value *GoogleComputeFirewallPolicyWithRulesRuleMatch)
	ResetDescription()
	ResetDirection()
	ResetDisabled()
	ResetEnableLogging()
	ResetRuleName()
	ResetSecurityProfileGroup()
	ResetTargetResources()
	ResetTargetServiceAccounts()
	ResetTlsInspect()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleComputeFirewallPolicyWithRulesRuleOutputReference
type jsiiProxy_GoogleComputeFirewallPolicyWithRulesRuleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleComputeFirewallPolicyWithRulesRuleOutputReference) Action() *string {
	var returns *string
	_jsii_.Get(
		j,
		"action",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFirewallPolicyWithRulesRuleOutputReference) ActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFirewallPolicyWithRulesRuleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFirewallPolicyWithRulesRuleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFirewallPolicyWithRulesRuleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFirewallPolicyWithRulesRuleOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFirewallPolicyWithRulesRuleOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFirewallPolicyWithRulesRuleOutputReference) Direction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"direction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFirewallPolicyWithRulesRuleOutputReference) DirectionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFirewallPolicyWithRulesRuleOutputReference) Disabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFirewallPolicyWithRulesRuleOutputReference) DisabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFirewallPolicyWithRulesRuleOutputReference) EnableLogging() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableLogging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFirewallPolicyWithRulesRuleOutputReference) EnableLoggingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableLoggingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFirewallPolicyWithRulesRuleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFirewallPolicyWithRulesRuleOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFirewallPolicyWithRulesRuleOutputReference) Match() GoogleComputeFirewallPolicyWithRulesRuleMatchOutputReference {
	var returns GoogleComputeFirewallPolicyWithRulesRuleMatchOutputReference
	_jsii_.Get(
		j,
		"match",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFirewallPolicyWithRulesRuleOutputReference) MatchInput() *GoogleComputeFirewallPolicyWithRulesRuleMatch {
	var returns *GoogleComputeFirewallPolicyWithRulesRuleMatch
	_jsii_.Get(
		j,
		"matchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFirewallPolicyWithRulesRuleOutputReference) Priority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFirewallPolicyWithRulesRuleOutputReference) PriorityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFirewallPolicyWithRulesRuleOutputReference) RuleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFirewallPolicyWithRulesRuleOutputReference) RuleNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFirewallPolicyWithRulesRuleOutputReference) SecurityProfileGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityProfileGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFirewallPolicyWithRulesRuleOutputReference) SecurityProfileGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityProfileGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFirewallPolicyWithRulesRuleOutputReference) TargetResources() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetResources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFirewallPolicyWithRulesRuleOutputReference) TargetResourcesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetResourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFirewallPolicyWithRulesRuleOutputReference) TargetServiceAccounts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetServiceAccounts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFirewallPolicyWithRulesRuleOutputReference) TargetServiceAccountsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetServiceAccountsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFirewallPolicyWithRulesRuleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFirewallPolicyWithRulesRuleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFirewallPolicyWithRulesRuleOutputReference) TlsInspect() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tlsInspect",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFirewallPolicyWithRulesRuleOutputReference) TlsInspectInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tlsInspectInput",
		&returns,
	)
	return returns
}


func NewGoogleComputeFirewallPolicyWithRulesRuleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleComputeFirewallPolicyWithRulesRuleOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleComputeFirewallPolicyWithRulesRuleOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeFirewallPolicyWithRulesRuleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleComputeFirewallPolicyWithRules.GoogleComputeFirewallPolicyWithRulesRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleComputeFirewallPolicyWithRulesRuleOutputReference_Override(g GoogleComputeFirewallPolicyWithRulesRuleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleComputeFirewallPolicyWithRules.GoogleComputeFirewallPolicyWithRulesRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeFirewallPolicyWithRulesRuleOutputReference)SetAction(val *string) {
	if err := j.validateSetActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"action",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeFirewallPolicyWithRulesRuleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeFirewallPolicyWithRulesRuleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeFirewallPolicyWithRulesRuleOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeFirewallPolicyWithRulesRuleOutputReference)SetDirection(val *string) {
	if err := j.validateSetDirectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"direction",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeFirewallPolicyWithRulesRuleOutputReference)SetDisabled(val interface{}) {
	if err := j.validateSetDisabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disabled",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeFirewallPolicyWithRulesRuleOutputReference)SetEnableLogging(val interface{}) {
	if err := j.validateSetEnableLoggingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableLogging",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeFirewallPolicyWithRulesRuleOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeFirewallPolicyWithRulesRuleOutputReference)SetPriority(val *float64) {
	if err := j.validateSetPriorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"priority",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeFirewallPolicyWithRulesRuleOutputReference)SetRuleName(val *string) {
	if err := j.validateSetRuleNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ruleName",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeFirewallPolicyWithRulesRuleOutputReference)SetSecurityProfileGroup(val *string) {
	if err := j.validateSetSecurityProfileGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityProfileGroup",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeFirewallPolicyWithRulesRuleOutputReference)SetTargetResources(val *[]*string) {
	if err := j.validateSetTargetResourcesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetResources",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeFirewallPolicyWithRulesRuleOutputReference)SetTargetServiceAccounts(val *[]*string) {
	if err := j.validateSetTargetServiceAccountsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetServiceAccounts",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeFirewallPolicyWithRulesRuleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeFirewallPolicyWithRulesRuleOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeFirewallPolicyWithRulesRuleOutputReference)SetTlsInspect(val interface{}) {
	if err := j.validateSetTlsInspectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tlsInspect",
		val,
	)
}

func (g *jsiiProxy_GoogleComputeFirewallPolicyWithRulesRuleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeFirewallPolicyWithRulesRuleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleComputeFirewallPolicyWithRulesRuleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeFirewallPolicyWithRulesRuleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComputeFirewallPolicyWithRulesRuleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComputeFirewallPolicyWithRulesRuleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComputeFirewallPolicyWithRulesRuleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComputeFirewallPolicyWithRulesRuleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComputeFirewallPolicyWithRulesRuleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComputeFirewallPolicyWithRulesRuleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComputeFirewallPolicyWithRulesRuleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeFirewallPolicyWithRulesRuleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeFirewallPolicyWithRulesRuleOutputReference) PutMatch(value *GoogleComputeFirewallPolicyWithRulesRuleMatch) {
	if err := g.validatePutMatchParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMatch",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeFirewallPolicyWithRulesRuleOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeFirewallPolicyWithRulesRuleOutputReference) ResetDirection() {
	_jsii_.InvokeVoid(
		g,
		"resetDirection",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeFirewallPolicyWithRulesRuleOutputReference) ResetDisabled() {
	_jsii_.InvokeVoid(
		g,
		"resetDisabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeFirewallPolicyWithRulesRuleOutputReference) ResetEnableLogging() {
	_jsii_.InvokeVoid(
		g,
		"resetEnableLogging",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeFirewallPolicyWithRulesRuleOutputReference) ResetRuleName() {
	_jsii_.InvokeVoid(
		g,
		"resetRuleName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeFirewallPolicyWithRulesRuleOutputReference) ResetSecurityProfileGroup() {
	_jsii_.InvokeVoid(
		g,
		"resetSecurityProfileGroup",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeFirewallPolicyWithRulesRuleOutputReference) ResetTargetResources() {
	_jsii_.InvokeVoid(
		g,
		"resetTargetResources",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeFirewallPolicyWithRulesRuleOutputReference) ResetTargetServiceAccounts() {
	_jsii_.InvokeVoid(
		g,
		"resetTargetServiceAccounts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeFirewallPolicyWithRulesRuleOutputReference) ResetTlsInspect() {
	_jsii_.InvokeVoid(
		g,
		"resetTlsInspect",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeFirewallPolicyWithRulesRuleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := g.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeFirewallPolicyWithRulesRuleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

