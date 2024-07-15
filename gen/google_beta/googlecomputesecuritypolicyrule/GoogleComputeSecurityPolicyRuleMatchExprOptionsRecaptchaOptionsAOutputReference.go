package googlecomputesecuritypolicyrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecomputesecuritypolicyrule/internal"
)

type GoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptionsAOutputReference interface {
	cdktf.ComplexObject
	ActionTokenSiteKeys() *[]*string
	SetActionTokenSiteKeys(val *[]*string)
	ActionTokenSiteKeysInput() *[]*string
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
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptionsA
	SetInternalValue(val *GoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptionsA)
	SessionTokenSiteKeys() *[]*string
	SetSessionTokenSiteKeys(val *[]*string)
	SessionTokenSiteKeysInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	ResetActionTokenSiteKeys()
	ResetSessionTokenSiteKeys()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptionsAOutputReference
type jsiiProxy_GoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptionsAOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptionsAOutputReference) ActionTokenSiteKeys() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"actionTokenSiteKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptionsAOutputReference) ActionTokenSiteKeysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"actionTokenSiteKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptionsAOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptionsAOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptionsAOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptionsAOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptionsAOutputReference) InternalValue() *GoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptionsA {
	var returns *GoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptionsA
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptionsAOutputReference) SessionTokenSiteKeys() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sessionTokenSiteKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptionsAOutputReference) SessionTokenSiteKeysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sessionTokenSiteKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptionsAOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptionsAOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptionsAOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptionsAOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptionsAOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptionsAOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleComputeSecurityPolicyRule.GoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptionsAOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptionsAOutputReference_Override(g GoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptionsAOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleComputeSecurityPolicyRule.GoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptionsAOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptionsAOutputReference)SetActionTokenSiteKeys(val *[]*string) {
	if err := j.validateSetActionTokenSiteKeysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"actionTokenSiteKeys",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptionsAOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptionsAOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptionsAOutputReference)SetInternalValue(val *GoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptionsA) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptionsAOutputReference)SetSessionTokenSiteKeys(val *[]*string) {
	if err := j.validateSetSessionTokenSiteKeysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sessionTokenSiteKeys",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptionsAOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptionsAOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptionsAOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptionsAOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptionsAOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptionsAOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptionsAOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptionsAOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptionsAOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptionsAOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptionsAOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptionsAOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptionsAOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptionsAOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptionsAOutputReference) ResetActionTokenSiteKeys() {
	_jsii_.InvokeVoid(
		g,
		"resetActionTokenSiteKeys",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptionsAOutputReference) ResetSessionTokenSiteKeys() {
	_jsii_.InvokeVoid(
		g,
		"resetSessionTokenSiteKeys",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptionsAOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptionsAOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

