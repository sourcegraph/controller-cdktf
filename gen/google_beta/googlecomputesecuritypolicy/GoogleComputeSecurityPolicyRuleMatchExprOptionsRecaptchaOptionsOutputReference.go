package googlecomputesecuritypolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecomputesecuritypolicy/internal"
)

type GoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptionsOutputReference interface {
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
	InternalValue() *GoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptions
	SetInternalValue(val *GoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptions)
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

// The jsii proxy struct for GoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptionsOutputReference
type jsiiProxy_GoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptionsOutputReference) ActionTokenSiteKeys() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"actionTokenSiteKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptionsOutputReference) ActionTokenSiteKeysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"actionTokenSiteKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptionsOutputReference) InternalValue() *GoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptions {
	var returns *GoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptionsOutputReference) SessionTokenSiteKeys() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sessionTokenSiteKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptionsOutputReference) SessionTokenSiteKeysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sessionTokenSiteKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleComputeSecurityPolicy.GoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptionsOutputReference_Override(g GoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleComputeSecurityPolicy.GoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptionsOutputReference)SetActionTokenSiteKeys(val *[]*string) {
	if err := j.validateSetActionTokenSiteKeysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"actionTokenSiteKeys",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptionsOutputReference)SetInternalValue(val *GoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptionsOutputReference)SetSessionTokenSiteKeys(val *[]*string) {
	if err := j.validateSetSessionTokenSiteKeysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sessionTokenSiteKeys",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptionsOutputReference) ResetActionTokenSiteKeys() {
	_jsii_.InvokeVoid(
		g,
		"resetActionTokenSiteKeys",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptionsOutputReference) ResetSessionTokenSiteKeys() {
	_jsii_.InvokeVoid(
		g,
		"resetSessionTokenSiteKeys",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

