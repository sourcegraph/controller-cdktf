package googlecomputerouterpeer

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecomputerouterpeer/internal"
)

type GoogleComputeRouterPeerBfdOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleComputeRouterPeerBfd
	SetInternalValue(val *GoogleComputeRouterPeerBfd)
	MinReceiveInterval() *float64
	SetMinReceiveInterval(val *float64)
	MinReceiveIntervalInput() *float64
	MinTransmitInterval() *float64
	SetMinTransmitInterval(val *float64)
	MinTransmitIntervalInput() *float64
	Multiplier() *float64
	SetMultiplier(val *float64)
	MultiplierInput() *float64
	SessionInitializationMode() *string
	SetSessionInitializationMode(val *string)
	SessionInitializationModeInput() *string
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
	ResetMinReceiveInterval()
	ResetMinTransmitInterval()
	ResetMultiplier()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleComputeRouterPeerBfdOutputReference
type jsiiProxy_GoogleComputeRouterPeerBfdOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleComputeRouterPeerBfdOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterPeerBfdOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterPeerBfdOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterPeerBfdOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterPeerBfdOutputReference) InternalValue() *GoogleComputeRouterPeerBfd {
	var returns *GoogleComputeRouterPeerBfd
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterPeerBfdOutputReference) MinReceiveInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minReceiveInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterPeerBfdOutputReference) MinReceiveIntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minReceiveIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterPeerBfdOutputReference) MinTransmitInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minTransmitInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterPeerBfdOutputReference) MinTransmitIntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minTransmitIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterPeerBfdOutputReference) Multiplier() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"multiplier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterPeerBfdOutputReference) MultiplierInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"multiplierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterPeerBfdOutputReference) SessionInitializationMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sessionInitializationMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterPeerBfdOutputReference) SessionInitializationModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sessionInitializationModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterPeerBfdOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterPeerBfdOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleComputeRouterPeerBfdOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleComputeRouterPeerBfdOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleComputeRouterPeerBfdOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeRouterPeerBfdOutputReference{}

	_jsii_.Create(
		"google-beta.googleComputeRouterPeer.GoogleComputeRouterPeerBfdOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleComputeRouterPeerBfdOutputReference_Override(g GoogleComputeRouterPeerBfdOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleComputeRouterPeer.GoogleComputeRouterPeerBfdOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeRouterPeerBfdOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRouterPeerBfdOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRouterPeerBfdOutputReference)SetInternalValue(val *GoogleComputeRouterPeerBfd) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRouterPeerBfdOutputReference)SetMinReceiveInterval(val *float64) {
	if err := j.validateSetMinReceiveIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minReceiveInterval",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRouterPeerBfdOutputReference)SetMinTransmitInterval(val *float64) {
	if err := j.validateSetMinTransmitIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minTransmitInterval",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRouterPeerBfdOutputReference)SetMultiplier(val *float64) {
	if err := j.validateSetMultiplierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"multiplier",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRouterPeerBfdOutputReference)SetSessionInitializationMode(val *string) {
	if err := j.validateSetSessionInitializationModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sessionInitializationMode",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRouterPeerBfdOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRouterPeerBfdOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleComputeRouterPeerBfdOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeRouterPeerBfdOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleComputeRouterPeerBfdOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeRouterPeerBfdOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComputeRouterPeerBfdOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComputeRouterPeerBfdOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComputeRouterPeerBfdOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComputeRouterPeerBfdOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComputeRouterPeerBfdOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComputeRouterPeerBfdOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComputeRouterPeerBfdOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeRouterPeerBfdOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeRouterPeerBfdOutputReference) ResetMinReceiveInterval() {
	_jsii_.InvokeVoid(
		g,
		"resetMinReceiveInterval",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRouterPeerBfdOutputReference) ResetMinTransmitInterval() {
	_jsii_.InvokeVoid(
		g,
		"resetMinTransmitInterval",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRouterPeerBfdOutputReference) ResetMultiplier() {
	_jsii_.InvokeVoid(
		g,
		"resetMultiplier",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRouterPeerBfdOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleComputeRouterPeerBfdOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

