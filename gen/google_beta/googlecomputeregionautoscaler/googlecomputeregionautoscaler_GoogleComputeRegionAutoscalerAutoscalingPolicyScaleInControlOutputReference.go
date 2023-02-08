package googlecomputeregionautoscaler

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecomputeregionautoscaler/internal"
)

type GoogleComputeRegionAutoscalerAutoscalingPolicyScaleInControlOutputReference interface {
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
	InternalValue() *GoogleComputeRegionAutoscalerAutoscalingPolicyScaleInControl
	SetInternalValue(val *GoogleComputeRegionAutoscalerAutoscalingPolicyScaleInControl)
	MaxScaledInReplicas() GoogleComputeRegionAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicasOutputReference
	MaxScaledInReplicasInput() *GoogleComputeRegionAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimeWindowSec() *float64
	SetTimeWindowSec(val *float64)
	TimeWindowSecInput() *float64
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
	PutMaxScaledInReplicas(value *GoogleComputeRegionAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas)
	ResetMaxScaledInReplicas()
	ResetTimeWindowSec()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleComputeRegionAutoscalerAutoscalingPolicyScaleInControlOutputReference
type jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyScaleInControlOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyScaleInControlOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyScaleInControlOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyScaleInControlOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyScaleInControlOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyScaleInControlOutputReference) InternalValue() *GoogleComputeRegionAutoscalerAutoscalingPolicyScaleInControl {
	var returns *GoogleComputeRegionAutoscalerAutoscalingPolicyScaleInControl
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyScaleInControlOutputReference) MaxScaledInReplicas() GoogleComputeRegionAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicasOutputReference {
	var returns GoogleComputeRegionAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicasOutputReference
	_jsii_.Get(
		j,
		"maxScaledInReplicas",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyScaleInControlOutputReference) MaxScaledInReplicasInput() *GoogleComputeRegionAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas {
	var returns *GoogleComputeRegionAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas
	_jsii_.Get(
		j,
		"maxScaledInReplicasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyScaleInControlOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyScaleInControlOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyScaleInControlOutputReference) TimeWindowSec() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeWindowSec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyScaleInControlOutputReference) TimeWindowSecInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeWindowSecInput",
		&returns,
	)
	return returns
}


func NewGoogleComputeRegionAutoscalerAutoscalingPolicyScaleInControlOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleComputeRegionAutoscalerAutoscalingPolicyScaleInControlOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleComputeRegionAutoscalerAutoscalingPolicyScaleInControlOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyScaleInControlOutputReference{}

	_jsii_.Create(
		"google-beta.googleComputeRegionAutoscaler.GoogleComputeRegionAutoscalerAutoscalingPolicyScaleInControlOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleComputeRegionAutoscalerAutoscalingPolicyScaleInControlOutputReference_Override(g GoogleComputeRegionAutoscalerAutoscalingPolicyScaleInControlOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleComputeRegionAutoscaler.GoogleComputeRegionAutoscalerAutoscalingPolicyScaleInControlOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyScaleInControlOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyScaleInControlOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyScaleInControlOutputReference)SetInternalValue(val *GoogleComputeRegionAutoscalerAutoscalingPolicyScaleInControl) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyScaleInControlOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyScaleInControlOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyScaleInControlOutputReference)SetTimeWindowSec(val *float64) {
	if err := j.validateSetTimeWindowSecParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeWindowSec",
		val,
	)
}

func (g *jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyScaleInControlOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyScaleInControlOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyScaleInControlOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyScaleInControlOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyScaleInControlOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyScaleInControlOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyScaleInControlOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyScaleInControlOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyScaleInControlOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyScaleInControlOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyScaleInControlOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyScaleInControlOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyScaleInControlOutputReference) PutMaxScaledInReplicas(value *GoogleComputeRegionAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas) {
	if err := g.validatePutMaxScaledInReplicasParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMaxScaledInReplicas",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyScaleInControlOutputReference) ResetMaxScaledInReplicas() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxScaledInReplicas",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyScaleInControlOutputReference) ResetTimeWindowSec() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeWindowSec",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyScaleInControlOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleComputeRegionAutoscalerAutoscalingPolicyScaleInControlOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

