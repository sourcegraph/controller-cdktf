package googlecomputesecuritypolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/googlebeta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/googlebeta/googlecomputesecuritypolicy/internal"
)

type GoogleComputeSecurityPolicyRuleRateLimitOptionsOutputReference interface {
	cdktf.ComplexObject
	BanDurationSec() *float64
	SetBanDurationSec(val *float64)
	BanDurationSecInput() *float64
	BanThreshold() GoogleComputeSecurityPolicyRuleRateLimitOptionsBanThresholdOutputReference
	BanThresholdInput() *GoogleComputeSecurityPolicyRuleRateLimitOptionsBanThreshold
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
	ConformAction() *string
	SetConformAction(val *string)
	ConformActionInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	EnforceOnKey() *string
	SetEnforceOnKey(val *string)
	EnforceOnKeyInput() *string
	EnforceOnKeyName() *string
	SetEnforceOnKeyName(val *string)
	EnforceOnKeyNameInput() *string
	ExceedAction() *string
	SetExceedAction(val *string)
	ExceedActionInput() *string
	ExceedRedirectOptions() GoogleComputeSecurityPolicyRuleRateLimitOptionsExceedRedirectOptionsOutputReference
	ExceedRedirectOptionsInput() *GoogleComputeSecurityPolicyRuleRateLimitOptionsExceedRedirectOptions
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleComputeSecurityPolicyRuleRateLimitOptions
	SetInternalValue(val *GoogleComputeSecurityPolicyRuleRateLimitOptions)
	RateLimitThreshold() GoogleComputeSecurityPolicyRuleRateLimitOptionsRateLimitThresholdOutputReference
	RateLimitThresholdInput() *GoogleComputeSecurityPolicyRuleRateLimitOptionsRateLimitThreshold
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
	PutBanThreshold(value *GoogleComputeSecurityPolicyRuleRateLimitOptionsBanThreshold)
	PutExceedRedirectOptions(value *GoogleComputeSecurityPolicyRuleRateLimitOptionsExceedRedirectOptions)
	PutRateLimitThreshold(value *GoogleComputeSecurityPolicyRuleRateLimitOptionsRateLimitThreshold)
	ResetBanDurationSec()
	ResetBanThreshold()
	ResetEnforceOnKey()
	ResetEnforceOnKeyName()
	ResetExceedRedirectOptions()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleComputeSecurityPolicyRuleRateLimitOptionsOutputReference
type jsiiProxy_GoogleComputeSecurityPolicyRuleRateLimitOptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleRateLimitOptionsOutputReference) BanDurationSec() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"banDurationSec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleRateLimitOptionsOutputReference) BanDurationSecInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"banDurationSecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleRateLimitOptionsOutputReference) BanThreshold() GoogleComputeSecurityPolicyRuleRateLimitOptionsBanThresholdOutputReference {
	var returns GoogleComputeSecurityPolicyRuleRateLimitOptionsBanThresholdOutputReference
	_jsii_.Get(
		j,
		"banThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleRateLimitOptionsOutputReference) BanThresholdInput() *GoogleComputeSecurityPolicyRuleRateLimitOptionsBanThreshold {
	var returns *GoogleComputeSecurityPolicyRuleRateLimitOptionsBanThreshold
	_jsii_.Get(
		j,
		"banThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleRateLimitOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleRateLimitOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleRateLimitOptionsOutputReference) ConformAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"conformAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleRateLimitOptionsOutputReference) ConformActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"conformActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleRateLimitOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleRateLimitOptionsOutputReference) EnforceOnKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enforceOnKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleRateLimitOptionsOutputReference) EnforceOnKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enforceOnKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleRateLimitOptionsOutputReference) EnforceOnKeyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enforceOnKeyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleRateLimitOptionsOutputReference) EnforceOnKeyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enforceOnKeyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleRateLimitOptionsOutputReference) ExceedAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exceedAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleRateLimitOptionsOutputReference) ExceedActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exceedActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleRateLimitOptionsOutputReference) ExceedRedirectOptions() GoogleComputeSecurityPolicyRuleRateLimitOptionsExceedRedirectOptionsOutputReference {
	var returns GoogleComputeSecurityPolicyRuleRateLimitOptionsExceedRedirectOptionsOutputReference
	_jsii_.Get(
		j,
		"exceedRedirectOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleRateLimitOptionsOutputReference) ExceedRedirectOptionsInput() *GoogleComputeSecurityPolicyRuleRateLimitOptionsExceedRedirectOptions {
	var returns *GoogleComputeSecurityPolicyRuleRateLimitOptionsExceedRedirectOptions
	_jsii_.Get(
		j,
		"exceedRedirectOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleRateLimitOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleRateLimitOptionsOutputReference) InternalValue() *GoogleComputeSecurityPolicyRuleRateLimitOptions {
	var returns *GoogleComputeSecurityPolicyRuleRateLimitOptions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleRateLimitOptionsOutputReference) RateLimitThreshold() GoogleComputeSecurityPolicyRuleRateLimitOptionsRateLimitThresholdOutputReference {
	var returns GoogleComputeSecurityPolicyRuleRateLimitOptionsRateLimitThresholdOutputReference
	_jsii_.Get(
		j,
		"rateLimitThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleRateLimitOptionsOutputReference) RateLimitThresholdInput() *GoogleComputeSecurityPolicyRuleRateLimitOptionsRateLimitThreshold {
	var returns *GoogleComputeSecurityPolicyRuleRateLimitOptionsRateLimitThreshold
	_jsii_.Get(
		j,
		"rateLimitThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleRateLimitOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleRateLimitOptionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleComputeSecurityPolicyRuleRateLimitOptionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleComputeSecurityPolicyRuleRateLimitOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleComputeSecurityPolicyRuleRateLimitOptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeSecurityPolicyRuleRateLimitOptionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-googlebeta.googleComputeSecurityPolicy.GoogleComputeSecurityPolicyRuleRateLimitOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleComputeSecurityPolicyRuleRateLimitOptionsOutputReference_Override(g GoogleComputeSecurityPolicyRuleRateLimitOptionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-googlebeta.googleComputeSecurityPolicy.GoogleComputeSecurityPolicyRuleRateLimitOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleRateLimitOptionsOutputReference)SetBanDurationSec(val *float64) {
	if err := j.validateSetBanDurationSecParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"banDurationSec",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleRateLimitOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleRateLimitOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleRateLimitOptionsOutputReference)SetConformAction(val *string) {
	if err := j.validateSetConformActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"conformAction",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleRateLimitOptionsOutputReference)SetEnforceOnKey(val *string) {
	if err := j.validateSetEnforceOnKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enforceOnKey",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleRateLimitOptionsOutputReference)SetEnforceOnKeyName(val *string) {
	if err := j.validateSetEnforceOnKeyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enforceOnKeyName",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleRateLimitOptionsOutputReference)SetExceedAction(val *string) {
	if err := j.validateSetExceedActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exceedAction",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleRateLimitOptionsOutputReference)SetInternalValue(val *GoogleComputeSecurityPolicyRuleRateLimitOptions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleRateLimitOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleRateLimitOptionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleComputeSecurityPolicyRuleRateLimitOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeSecurityPolicyRuleRateLimitOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleComputeSecurityPolicyRuleRateLimitOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeSecurityPolicyRuleRateLimitOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComputeSecurityPolicyRuleRateLimitOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComputeSecurityPolicyRuleRateLimitOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComputeSecurityPolicyRuleRateLimitOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComputeSecurityPolicyRuleRateLimitOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComputeSecurityPolicyRuleRateLimitOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComputeSecurityPolicyRuleRateLimitOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComputeSecurityPolicyRuleRateLimitOptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeSecurityPolicyRuleRateLimitOptionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeSecurityPolicyRuleRateLimitOptionsOutputReference) PutBanThreshold(value *GoogleComputeSecurityPolicyRuleRateLimitOptionsBanThreshold) {
	if err := g.validatePutBanThresholdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putBanThreshold",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeSecurityPolicyRuleRateLimitOptionsOutputReference) PutExceedRedirectOptions(value *GoogleComputeSecurityPolicyRuleRateLimitOptionsExceedRedirectOptions) {
	if err := g.validatePutExceedRedirectOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putExceedRedirectOptions",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeSecurityPolicyRuleRateLimitOptionsOutputReference) PutRateLimitThreshold(value *GoogleComputeSecurityPolicyRuleRateLimitOptionsRateLimitThreshold) {
	if err := g.validatePutRateLimitThresholdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRateLimitThreshold",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeSecurityPolicyRuleRateLimitOptionsOutputReference) ResetBanDurationSec() {
	_jsii_.InvokeVoid(
		g,
		"resetBanDurationSec",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeSecurityPolicyRuleRateLimitOptionsOutputReference) ResetBanThreshold() {
	_jsii_.InvokeVoid(
		g,
		"resetBanThreshold",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeSecurityPolicyRuleRateLimitOptionsOutputReference) ResetEnforceOnKey() {
	_jsii_.InvokeVoid(
		g,
		"resetEnforceOnKey",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeSecurityPolicyRuleRateLimitOptionsOutputReference) ResetEnforceOnKeyName() {
	_jsii_.InvokeVoid(
		g,
		"resetEnforceOnKeyName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeSecurityPolicyRuleRateLimitOptionsOutputReference) ResetExceedRedirectOptions() {
	_jsii_.InvokeVoid(
		g,
		"resetExceedRedirectOptions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeSecurityPolicyRuleRateLimitOptionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleComputeSecurityPolicyRuleRateLimitOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

