package googlecomputeinterconnect

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecomputeinterconnect/internal"
)

type GoogleComputeInterconnectApplicationAwareInterconnectOutputReference interface {
	cdktn.ComplexObject
	BandwidthPercentagePolicy() GoogleComputeInterconnectApplicationAwareInterconnectBandwidthPercentagePolicyOutputReference
	BandwidthPercentagePolicyInput() *GoogleComputeInterconnectApplicationAwareInterconnectBandwidthPercentagePolicy
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
	InternalValue() *GoogleComputeInterconnectApplicationAwareInterconnect
	SetInternalValue(val *GoogleComputeInterconnectApplicationAwareInterconnect)
	ProfileDescription() *string
	SetProfileDescription(val *string)
	ProfileDescriptionInput() *string
	ShapeAveragePercentage() GoogleComputeInterconnectApplicationAwareInterconnectShapeAveragePercentageList
	ShapeAveragePercentageInput() interface{}
	StrictPriorityPolicy() GoogleComputeInterconnectApplicationAwareInterconnectStrictPriorityPolicyOutputReference
	StrictPriorityPolicyInput() *GoogleComputeInterconnectApplicationAwareInterconnectStrictPriorityPolicy
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
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
	InterpolationAsList() cdktn.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	PutBandwidthPercentagePolicy(value *GoogleComputeInterconnectApplicationAwareInterconnectBandwidthPercentagePolicy)
	PutShapeAveragePercentage(value interface{})
	PutStrictPriorityPolicy(value *GoogleComputeInterconnectApplicationAwareInterconnectStrictPriorityPolicy)
	ResetBandwidthPercentagePolicy()
	ResetProfileDescription()
	ResetShapeAveragePercentage()
	ResetStrictPriorityPolicy()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleComputeInterconnectApplicationAwareInterconnectOutputReference
type jsiiProxy_GoogleComputeInterconnectApplicationAwareInterconnectOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleComputeInterconnectApplicationAwareInterconnectOutputReference) BandwidthPercentagePolicy() GoogleComputeInterconnectApplicationAwareInterconnectBandwidthPercentagePolicyOutputReference {
	var returns GoogleComputeInterconnectApplicationAwareInterconnectBandwidthPercentagePolicyOutputReference
	_jsii_.Get(
		j,
		"bandwidthPercentagePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnectApplicationAwareInterconnectOutputReference) BandwidthPercentagePolicyInput() *GoogleComputeInterconnectApplicationAwareInterconnectBandwidthPercentagePolicy {
	var returns *GoogleComputeInterconnectApplicationAwareInterconnectBandwidthPercentagePolicy
	_jsii_.Get(
		j,
		"bandwidthPercentagePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnectApplicationAwareInterconnectOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnectApplicationAwareInterconnectOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnectApplicationAwareInterconnectOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnectApplicationAwareInterconnectOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnectApplicationAwareInterconnectOutputReference) InternalValue() *GoogleComputeInterconnectApplicationAwareInterconnect {
	var returns *GoogleComputeInterconnectApplicationAwareInterconnect
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnectApplicationAwareInterconnectOutputReference) ProfileDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profileDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnectApplicationAwareInterconnectOutputReference) ProfileDescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profileDescriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnectApplicationAwareInterconnectOutputReference) ShapeAveragePercentage() GoogleComputeInterconnectApplicationAwareInterconnectShapeAveragePercentageList {
	var returns GoogleComputeInterconnectApplicationAwareInterconnectShapeAveragePercentageList
	_jsii_.Get(
		j,
		"shapeAveragePercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnectApplicationAwareInterconnectOutputReference) ShapeAveragePercentageInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shapeAveragePercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnectApplicationAwareInterconnectOutputReference) StrictPriorityPolicy() GoogleComputeInterconnectApplicationAwareInterconnectStrictPriorityPolicyOutputReference {
	var returns GoogleComputeInterconnectApplicationAwareInterconnectStrictPriorityPolicyOutputReference
	_jsii_.Get(
		j,
		"strictPriorityPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnectApplicationAwareInterconnectOutputReference) StrictPriorityPolicyInput() *GoogleComputeInterconnectApplicationAwareInterconnectStrictPriorityPolicy {
	var returns *GoogleComputeInterconnectApplicationAwareInterconnectStrictPriorityPolicy
	_jsii_.Get(
		j,
		"strictPriorityPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnectApplicationAwareInterconnectOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnectApplicationAwareInterconnectOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleComputeInterconnectApplicationAwareInterconnectOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleComputeInterconnectApplicationAwareInterconnectOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleComputeInterconnectApplicationAwareInterconnectOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeInterconnectApplicationAwareInterconnectOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleComputeInterconnect.GoogleComputeInterconnectApplicationAwareInterconnectOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleComputeInterconnectApplicationAwareInterconnectOutputReference_Override(g GoogleComputeInterconnectApplicationAwareInterconnectOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleComputeInterconnect.GoogleComputeInterconnectApplicationAwareInterconnectOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeInterconnectApplicationAwareInterconnectOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInterconnectApplicationAwareInterconnectOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInterconnectApplicationAwareInterconnectOutputReference)SetInternalValue(val *GoogleComputeInterconnectApplicationAwareInterconnect) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInterconnectApplicationAwareInterconnectOutputReference)SetProfileDescription(val *string) {
	if err := j.validateSetProfileDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"profileDescription",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInterconnectApplicationAwareInterconnectOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInterconnectApplicationAwareInterconnectOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleComputeInterconnectApplicationAwareInterconnectOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeInterconnectApplicationAwareInterconnectOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleComputeInterconnectApplicationAwareInterconnectOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeInterconnectApplicationAwareInterconnectOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComputeInterconnectApplicationAwareInterconnectOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComputeInterconnectApplicationAwareInterconnectOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComputeInterconnectApplicationAwareInterconnectOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComputeInterconnectApplicationAwareInterconnectOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComputeInterconnectApplicationAwareInterconnectOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComputeInterconnectApplicationAwareInterconnectOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComputeInterconnectApplicationAwareInterconnectOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeInterconnectApplicationAwareInterconnectOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeInterconnectApplicationAwareInterconnectOutputReference) PutBandwidthPercentagePolicy(value *GoogleComputeInterconnectApplicationAwareInterconnectBandwidthPercentagePolicy) {
	if err := g.validatePutBandwidthPercentagePolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putBandwidthPercentagePolicy",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeInterconnectApplicationAwareInterconnectOutputReference) PutShapeAveragePercentage(value interface{}) {
	if err := g.validatePutShapeAveragePercentageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putShapeAveragePercentage",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeInterconnectApplicationAwareInterconnectOutputReference) PutStrictPriorityPolicy(value *GoogleComputeInterconnectApplicationAwareInterconnectStrictPriorityPolicy) {
	if err := g.validatePutStrictPriorityPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putStrictPriorityPolicy",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeInterconnectApplicationAwareInterconnectOutputReference) ResetBandwidthPercentagePolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetBandwidthPercentagePolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInterconnectApplicationAwareInterconnectOutputReference) ResetProfileDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetProfileDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInterconnectApplicationAwareInterconnectOutputReference) ResetShapeAveragePercentage() {
	_jsii_.InvokeVoid(
		g,
		"resetShapeAveragePercentage",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInterconnectApplicationAwareInterconnectOutputReference) ResetStrictPriorityPolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetStrictPriorityPolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInterconnectApplicationAwareInterconnectOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := g.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeInterconnectApplicationAwareInterconnectOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

