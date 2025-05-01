package googlecomputeregionsecuritypolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecomputeregionsecuritypolicy/internal"
)

type GoogleComputeRegionSecurityPolicyRulesRateLimitOptionsOutputReference interface {
	cdktf.ComplexObject
	BanDurationSec() *float64
	SetBanDurationSec(val *float64)
	BanDurationSecInput() *float64
	BanThreshold() GoogleComputeRegionSecurityPolicyRulesRateLimitOptionsBanThresholdOutputReference
	BanThresholdInput() *GoogleComputeRegionSecurityPolicyRulesRateLimitOptionsBanThreshold
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
	EnforceOnKeyConfigs() GoogleComputeRegionSecurityPolicyRulesRateLimitOptionsEnforceOnKeyConfigsList
	EnforceOnKeyConfigsInput() interface{}
	EnforceOnKeyInput() *string
	EnforceOnKeyName() *string
	SetEnforceOnKeyName(val *string)
	EnforceOnKeyNameInput() *string
	ExceedAction() *string
	SetExceedAction(val *string)
	ExceedActionInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleComputeRegionSecurityPolicyRulesRateLimitOptions
	SetInternalValue(val *GoogleComputeRegionSecurityPolicyRulesRateLimitOptions)
	RateLimitThreshold() GoogleComputeRegionSecurityPolicyRulesRateLimitOptionsRateLimitThresholdOutputReference
	RateLimitThresholdInput() *GoogleComputeRegionSecurityPolicyRulesRateLimitOptionsRateLimitThreshold
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
	PutBanThreshold(value *GoogleComputeRegionSecurityPolicyRulesRateLimitOptionsBanThreshold)
	PutEnforceOnKeyConfigs(value interface{})
	PutRateLimitThreshold(value *GoogleComputeRegionSecurityPolicyRulesRateLimitOptionsRateLimitThreshold)
	ResetBanDurationSec()
	ResetBanThreshold()
	ResetConformAction()
	ResetEnforceOnKey()
	ResetEnforceOnKeyConfigs()
	ResetEnforceOnKeyName()
	ResetExceedAction()
	ResetRateLimitThreshold()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleComputeRegionSecurityPolicyRulesRateLimitOptionsOutputReference
type jsiiProxy_GoogleComputeRegionSecurityPolicyRulesRateLimitOptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesRateLimitOptionsOutputReference) BanDurationSec() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"banDurationSec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesRateLimitOptionsOutputReference) BanDurationSecInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"banDurationSecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesRateLimitOptionsOutputReference) BanThreshold() GoogleComputeRegionSecurityPolicyRulesRateLimitOptionsBanThresholdOutputReference {
	var returns GoogleComputeRegionSecurityPolicyRulesRateLimitOptionsBanThresholdOutputReference
	_jsii_.Get(
		j,
		"banThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesRateLimitOptionsOutputReference) BanThresholdInput() *GoogleComputeRegionSecurityPolicyRulesRateLimitOptionsBanThreshold {
	var returns *GoogleComputeRegionSecurityPolicyRulesRateLimitOptionsBanThreshold
	_jsii_.Get(
		j,
		"banThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesRateLimitOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesRateLimitOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesRateLimitOptionsOutputReference) ConformAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"conformAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesRateLimitOptionsOutputReference) ConformActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"conformActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesRateLimitOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesRateLimitOptionsOutputReference) EnforceOnKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enforceOnKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesRateLimitOptionsOutputReference) EnforceOnKeyConfigs() GoogleComputeRegionSecurityPolicyRulesRateLimitOptionsEnforceOnKeyConfigsList {
	var returns GoogleComputeRegionSecurityPolicyRulesRateLimitOptionsEnforceOnKeyConfigsList
	_jsii_.Get(
		j,
		"enforceOnKeyConfigs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesRateLimitOptionsOutputReference) EnforceOnKeyConfigsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enforceOnKeyConfigsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesRateLimitOptionsOutputReference) EnforceOnKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enforceOnKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesRateLimitOptionsOutputReference) EnforceOnKeyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enforceOnKeyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesRateLimitOptionsOutputReference) EnforceOnKeyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enforceOnKeyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesRateLimitOptionsOutputReference) ExceedAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exceedAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesRateLimitOptionsOutputReference) ExceedActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exceedActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesRateLimitOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesRateLimitOptionsOutputReference) InternalValue() *GoogleComputeRegionSecurityPolicyRulesRateLimitOptions {
	var returns *GoogleComputeRegionSecurityPolicyRulesRateLimitOptions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesRateLimitOptionsOutputReference) RateLimitThreshold() GoogleComputeRegionSecurityPolicyRulesRateLimitOptionsRateLimitThresholdOutputReference {
	var returns GoogleComputeRegionSecurityPolicyRulesRateLimitOptionsRateLimitThresholdOutputReference
	_jsii_.Get(
		j,
		"rateLimitThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesRateLimitOptionsOutputReference) RateLimitThresholdInput() *GoogleComputeRegionSecurityPolicyRulesRateLimitOptionsRateLimitThreshold {
	var returns *GoogleComputeRegionSecurityPolicyRulesRateLimitOptionsRateLimitThreshold
	_jsii_.Get(
		j,
		"rateLimitThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesRateLimitOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesRateLimitOptionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleComputeRegionSecurityPolicyRulesRateLimitOptionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleComputeRegionSecurityPolicyRulesRateLimitOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleComputeRegionSecurityPolicyRulesRateLimitOptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeRegionSecurityPolicyRulesRateLimitOptionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleComputeRegionSecurityPolicy.GoogleComputeRegionSecurityPolicyRulesRateLimitOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleComputeRegionSecurityPolicyRulesRateLimitOptionsOutputReference_Override(g GoogleComputeRegionSecurityPolicyRulesRateLimitOptionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleComputeRegionSecurityPolicy.GoogleComputeRegionSecurityPolicyRulesRateLimitOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesRateLimitOptionsOutputReference)SetBanDurationSec(val *float64) {
	if err := j.validateSetBanDurationSecParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"banDurationSec",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesRateLimitOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesRateLimitOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesRateLimitOptionsOutputReference)SetConformAction(val *string) {
	if err := j.validateSetConformActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"conformAction",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesRateLimitOptionsOutputReference)SetEnforceOnKey(val *string) {
	if err := j.validateSetEnforceOnKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enforceOnKey",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesRateLimitOptionsOutputReference)SetEnforceOnKeyName(val *string) {
	if err := j.validateSetEnforceOnKeyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enforceOnKeyName",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesRateLimitOptionsOutputReference)SetExceedAction(val *string) {
	if err := j.validateSetExceedActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exceedAction",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesRateLimitOptionsOutputReference)SetInternalValue(val *GoogleComputeRegionSecurityPolicyRulesRateLimitOptions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesRateLimitOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesRateLimitOptionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesRateLimitOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesRateLimitOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesRateLimitOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesRateLimitOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesRateLimitOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesRateLimitOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesRateLimitOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesRateLimitOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesRateLimitOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesRateLimitOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesRateLimitOptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesRateLimitOptionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesRateLimitOptionsOutputReference) PutBanThreshold(value *GoogleComputeRegionSecurityPolicyRulesRateLimitOptionsBanThreshold) {
	if err := g.validatePutBanThresholdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putBanThreshold",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesRateLimitOptionsOutputReference) PutEnforceOnKeyConfigs(value interface{}) {
	if err := g.validatePutEnforceOnKeyConfigsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putEnforceOnKeyConfigs",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesRateLimitOptionsOutputReference) PutRateLimitThreshold(value *GoogleComputeRegionSecurityPolicyRulesRateLimitOptionsRateLimitThreshold) {
	if err := g.validatePutRateLimitThresholdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRateLimitThreshold",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesRateLimitOptionsOutputReference) ResetBanDurationSec() {
	_jsii_.InvokeVoid(
		g,
		"resetBanDurationSec",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesRateLimitOptionsOutputReference) ResetBanThreshold() {
	_jsii_.InvokeVoid(
		g,
		"resetBanThreshold",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesRateLimitOptionsOutputReference) ResetConformAction() {
	_jsii_.InvokeVoid(
		g,
		"resetConformAction",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesRateLimitOptionsOutputReference) ResetEnforceOnKey() {
	_jsii_.InvokeVoid(
		g,
		"resetEnforceOnKey",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesRateLimitOptionsOutputReference) ResetEnforceOnKeyConfigs() {
	_jsii_.InvokeVoid(
		g,
		"resetEnforceOnKeyConfigs",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesRateLimitOptionsOutputReference) ResetEnforceOnKeyName() {
	_jsii_.InvokeVoid(
		g,
		"resetEnforceOnKeyName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesRateLimitOptionsOutputReference) ResetExceedAction() {
	_jsii_.InvokeVoid(
		g,
		"resetExceedAction",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesRateLimitOptionsOutputReference) ResetRateLimitThreshold() {
	_jsii_.InvokeVoid(
		g,
		"resetRateLimitThreshold",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesRateLimitOptionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleComputeRegionSecurityPolicyRulesRateLimitOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

