package googletranscoderjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googletranscoderjob/internal"
)

type GoogleTranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference interface {
	cdktf.ComplexObject
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
	EndTimeOffset() *string
	SetEndTimeOffset(val *string)
	EndTimeOffsetInput() *string
	FadeType() *string
	SetFadeType(val *string)
	FadeTypeInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleTranscoderJobConfigOverlaysAnimationsAnimationFade
	SetInternalValue(val *GoogleTranscoderJobConfigOverlaysAnimationsAnimationFade)
	StartTimeOffset() *string
	SetStartTimeOffset(val *string)
	StartTimeOffsetInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Xy() GoogleTranscoderJobConfigOverlaysAnimationsAnimationFadeXyOutputReference
	XyInput() *GoogleTranscoderJobConfigOverlaysAnimationsAnimationFadeXy
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
	PutXy(value *GoogleTranscoderJobConfigOverlaysAnimationsAnimationFadeXy)
	ResetEndTimeOffset()
	ResetStartTimeOffset()
	ResetXy()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleTranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference
type jsiiProxy_GoogleTranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleTranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference) EndTimeOffset() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endTimeOffset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference) EndTimeOffsetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endTimeOffsetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference) FadeType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fadeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference) FadeTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fadeTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference) InternalValue() *GoogleTranscoderJobConfigOverlaysAnimationsAnimationFade {
	var returns *GoogleTranscoderJobConfigOverlaysAnimationsAnimationFade
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference) StartTimeOffset() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startTimeOffset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference) StartTimeOffsetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startTimeOffsetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference) Xy() GoogleTranscoderJobConfigOverlaysAnimationsAnimationFadeXyOutputReference {
	var returns GoogleTranscoderJobConfigOverlaysAnimationsAnimationFadeXyOutputReference
	_jsii_.Get(
		j,
		"xy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference) XyInput() *GoogleTranscoderJobConfigOverlaysAnimationsAnimationFadeXy {
	var returns *GoogleTranscoderJobConfigOverlaysAnimationsAnimationFadeXy
	_jsii_.Get(
		j,
		"xyInput",
		&returns,
	)
	return returns
}


func NewGoogleTranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleTranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleTranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleTranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleTranscoderJob.GoogleTranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleTranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference_Override(g GoogleTranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleTranscoderJob.GoogleTranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleTranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleTranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleTranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference)SetEndTimeOffset(val *string) {
	if err := j.validateSetEndTimeOffsetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endTimeOffset",
		val,
	)
}

func (j *jsiiProxy_GoogleTranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference)SetFadeType(val *string) {
	if err := j.validateSetFadeTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fadeType",
		val,
	)
}

func (j *jsiiProxy_GoogleTranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference)SetInternalValue(val *GoogleTranscoderJobConfigOverlaysAnimationsAnimationFade) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleTranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference)SetStartTimeOffset(val *string) {
	if err := j.validateSetStartTimeOffsetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startTimeOffset",
		val,
	)
}

func (j *jsiiProxy_GoogleTranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleTranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleTranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleTranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleTranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleTranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleTranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleTranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleTranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleTranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleTranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleTranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleTranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleTranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleTranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference) PutXy(value *GoogleTranscoderJobConfigOverlaysAnimationsAnimationFadeXy) {
	if err := g.validatePutXyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putXy",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleTranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference) ResetEndTimeOffset() {
	_jsii_.InvokeVoid(
		g,
		"resetEndTimeOffset",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleTranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference) ResetStartTimeOffset() {
	_jsii_.InvokeVoid(
		g,
		"resetStartTimeOffset",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleTranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference) ResetXy() {
	_jsii_.InvokeVoid(
		g,
		"resetXy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleTranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleTranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

