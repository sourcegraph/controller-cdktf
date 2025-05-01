package googlecomputenetworkfirewallpolicywithrules

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecomputenetworkfirewallpolicywithrules/internal"
)

type GoogleComputeNetworkFirewallPolicyWithRulesRuleOutputReference interface {
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
	Match() GoogleComputeNetworkFirewallPolicyWithRulesRuleMatchOutputReference
	MatchInput() *GoogleComputeNetworkFirewallPolicyWithRulesRuleMatch
	Priority() *float64
	SetPriority(val *float64)
	PriorityInput() *float64
	RuleName() *string
	SetRuleName(val *string)
	RuleNameInput() *string
	SecurityProfileGroup() *string
	SetSecurityProfileGroup(val *string)
	SecurityProfileGroupInput() *string
	TargetSecureTag() GoogleComputeNetworkFirewallPolicyWithRulesRuleTargetSecureTagList
	TargetSecureTagInput() interface{}
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
	PutMatch(value *GoogleComputeNetworkFirewallPolicyWithRulesRuleMatch)
	PutTargetSecureTag(value interface{})
	ResetDescription()
	ResetDirection()
	ResetDisabled()
	ResetEnableLogging()
	ResetRuleName()
	ResetSecurityProfileGroup()
	ResetTargetSecureTag()
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

// The jsii proxy struct for GoogleComputeNetworkFirewallPolicyWithRulesRuleOutputReference
type jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesRuleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesRuleOutputReference) Action() *string {
	var returns *string
	_jsii_.Get(
		j,
		"action",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesRuleOutputReference) ActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesRuleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesRuleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesRuleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesRuleOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesRuleOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesRuleOutputReference) Direction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"direction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesRuleOutputReference) DirectionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesRuleOutputReference) Disabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesRuleOutputReference) DisabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesRuleOutputReference) EnableLogging() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableLogging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesRuleOutputReference) EnableLoggingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableLoggingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesRuleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesRuleOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesRuleOutputReference) Match() GoogleComputeNetworkFirewallPolicyWithRulesRuleMatchOutputReference {
	var returns GoogleComputeNetworkFirewallPolicyWithRulesRuleMatchOutputReference
	_jsii_.Get(
		j,
		"match",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesRuleOutputReference) MatchInput() *GoogleComputeNetworkFirewallPolicyWithRulesRuleMatch {
	var returns *GoogleComputeNetworkFirewallPolicyWithRulesRuleMatch
	_jsii_.Get(
		j,
		"matchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesRuleOutputReference) Priority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesRuleOutputReference) PriorityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesRuleOutputReference) RuleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesRuleOutputReference) RuleNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesRuleOutputReference) SecurityProfileGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityProfileGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesRuleOutputReference) SecurityProfileGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityProfileGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesRuleOutputReference) TargetSecureTag() GoogleComputeNetworkFirewallPolicyWithRulesRuleTargetSecureTagList {
	var returns GoogleComputeNetworkFirewallPolicyWithRulesRuleTargetSecureTagList
	_jsii_.Get(
		j,
		"targetSecureTag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesRuleOutputReference) TargetSecureTagInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"targetSecureTagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesRuleOutputReference) TargetServiceAccounts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetServiceAccounts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesRuleOutputReference) TargetServiceAccountsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetServiceAccountsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesRuleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesRuleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesRuleOutputReference) TlsInspect() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tlsInspect",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesRuleOutputReference) TlsInspectInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tlsInspectInput",
		&returns,
	)
	return returns
}


func NewGoogleComputeNetworkFirewallPolicyWithRulesRuleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleComputeNetworkFirewallPolicyWithRulesRuleOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleComputeNetworkFirewallPolicyWithRulesRuleOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesRuleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleComputeNetworkFirewallPolicyWithRules.GoogleComputeNetworkFirewallPolicyWithRulesRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleComputeNetworkFirewallPolicyWithRulesRuleOutputReference_Override(g GoogleComputeNetworkFirewallPolicyWithRulesRuleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleComputeNetworkFirewallPolicyWithRules.GoogleComputeNetworkFirewallPolicyWithRulesRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesRuleOutputReference)SetAction(val *string) {
	if err := j.validateSetActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"action",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesRuleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesRuleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesRuleOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesRuleOutputReference)SetDirection(val *string) {
	if err := j.validateSetDirectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"direction",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesRuleOutputReference)SetDisabled(val interface{}) {
	if err := j.validateSetDisabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disabled",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesRuleOutputReference)SetEnableLogging(val interface{}) {
	if err := j.validateSetEnableLoggingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableLogging",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesRuleOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesRuleOutputReference)SetPriority(val *float64) {
	if err := j.validateSetPriorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"priority",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesRuleOutputReference)SetRuleName(val *string) {
	if err := j.validateSetRuleNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ruleName",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesRuleOutputReference)SetSecurityProfileGroup(val *string) {
	if err := j.validateSetSecurityProfileGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityProfileGroup",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesRuleOutputReference)SetTargetServiceAccounts(val *[]*string) {
	if err := j.validateSetTargetServiceAccountsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetServiceAccounts",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesRuleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesRuleOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesRuleOutputReference)SetTlsInspect(val interface{}) {
	if err := j.validateSetTlsInspectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tlsInspect",
		val,
	)
}

func (g *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesRuleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesRuleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesRuleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesRuleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesRuleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesRuleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesRuleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesRuleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesRuleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesRuleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesRuleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesRuleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesRuleOutputReference) PutMatch(value *GoogleComputeNetworkFirewallPolicyWithRulesRuleMatch) {
	if err := g.validatePutMatchParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMatch",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesRuleOutputReference) PutTargetSecureTag(value interface{}) {
	if err := g.validatePutTargetSecureTagParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTargetSecureTag",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesRuleOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesRuleOutputReference) ResetDirection() {
	_jsii_.InvokeVoid(
		g,
		"resetDirection",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesRuleOutputReference) ResetDisabled() {
	_jsii_.InvokeVoid(
		g,
		"resetDisabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesRuleOutputReference) ResetEnableLogging() {
	_jsii_.InvokeVoid(
		g,
		"resetEnableLogging",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesRuleOutputReference) ResetRuleName() {
	_jsii_.InvokeVoid(
		g,
		"resetRuleName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesRuleOutputReference) ResetSecurityProfileGroup() {
	_jsii_.InvokeVoid(
		g,
		"resetSecurityProfileGroup",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesRuleOutputReference) ResetTargetSecureTag() {
	_jsii_.InvokeVoid(
		g,
		"resetTargetSecureTag",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesRuleOutputReference) ResetTargetServiceAccounts() {
	_jsii_.InvokeVoid(
		g,
		"resetTargetServiceAccounts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesRuleOutputReference) ResetTlsInspect() {
	_jsii_.InvokeVoid(
		g,
		"resetTlsInspect",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesRuleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesRuleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

