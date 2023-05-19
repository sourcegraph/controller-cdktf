package googlecomputeautoscaler

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecomputeautoscaler/internal"
)

type GoogleComputeAutoscalerAutoscalingPolicyScaleDownControlOutputReference interface {
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
	InternalValue() *GoogleComputeAutoscalerAutoscalingPolicyScaleDownControl
	SetInternalValue(val *GoogleComputeAutoscalerAutoscalingPolicyScaleDownControl)
	MaxScaledDownReplicas() GoogleComputeAutoscalerAutoscalingPolicyScaleDownControlMaxScaledDownReplicasOutputReference
	MaxScaledDownReplicasInput() *GoogleComputeAutoscalerAutoscalingPolicyScaleDownControlMaxScaledDownReplicas
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
	PutMaxScaledDownReplicas(value *GoogleComputeAutoscalerAutoscalingPolicyScaleDownControlMaxScaledDownReplicas)
	ResetMaxScaledDownReplicas()
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

// The jsii proxy struct for GoogleComputeAutoscalerAutoscalingPolicyScaleDownControlOutputReference
type jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleDownControlOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleDownControlOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleDownControlOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleDownControlOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleDownControlOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleDownControlOutputReference) InternalValue() *GoogleComputeAutoscalerAutoscalingPolicyScaleDownControl {
	var returns *GoogleComputeAutoscalerAutoscalingPolicyScaleDownControl
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleDownControlOutputReference) MaxScaledDownReplicas() GoogleComputeAutoscalerAutoscalingPolicyScaleDownControlMaxScaledDownReplicasOutputReference {
	var returns GoogleComputeAutoscalerAutoscalingPolicyScaleDownControlMaxScaledDownReplicasOutputReference
	_jsii_.Get(
		j,
		"maxScaledDownReplicas",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleDownControlOutputReference) MaxScaledDownReplicasInput() *GoogleComputeAutoscalerAutoscalingPolicyScaleDownControlMaxScaledDownReplicas {
	var returns *GoogleComputeAutoscalerAutoscalingPolicyScaleDownControlMaxScaledDownReplicas
	_jsii_.Get(
		j,
		"maxScaledDownReplicasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleDownControlOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleDownControlOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleDownControlOutputReference) TimeWindowSec() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeWindowSec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleDownControlOutputReference) TimeWindowSecInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeWindowSecInput",
		&returns,
	)
	return returns
}


func NewGoogleComputeAutoscalerAutoscalingPolicyScaleDownControlOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleComputeAutoscalerAutoscalingPolicyScaleDownControlOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleComputeAutoscalerAutoscalingPolicyScaleDownControlOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleDownControlOutputReference{}

	_jsii_.Create(
		"google-beta.googleComputeAutoscaler.GoogleComputeAutoscalerAutoscalingPolicyScaleDownControlOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleComputeAutoscalerAutoscalingPolicyScaleDownControlOutputReference_Override(g GoogleComputeAutoscalerAutoscalingPolicyScaleDownControlOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleComputeAutoscaler.GoogleComputeAutoscalerAutoscalingPolicyScaleDownControlOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleDownControlOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleDownControlOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleDownControlOutputReference)SetInternalValue(val *GoogleComputeAutoscalerAutoscalingPolicyScaleDownControl) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleDownControlOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleDownControlOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleDownControlOutputReference)SetTimeWindowSec(val *float64) {
	if err := j.validateSetTimeWindowSecParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeWindowSec",
		val,
	)
}

func (g *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleDownControlOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleDownControlOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleDownControlOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleDownControlOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleDownControlOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleDownControlOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleDownControlOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleDownControlOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleDownControlOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleDownControlOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleDownControlOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleDownControlOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleDownControlOutputReference) PutMaxScaledDownReplicas(value *GoogleComputeAutoscalerAutoscalingPolicyScaleDownControlMaxScaledDownReplicas) {
	if err := g.validatePutMaxScaledDownReplicasParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMaxScaledDownReplicas",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleDownControlOutputReference) ResetMaxScaledDownReplicas() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxScaledDownReplicas",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleDownControlOutputReference) ResetTimeWindowSec() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeWindowSec",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleDownControlOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleDownControlOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

