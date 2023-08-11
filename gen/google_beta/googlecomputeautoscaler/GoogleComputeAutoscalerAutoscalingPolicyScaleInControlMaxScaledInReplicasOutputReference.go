package googlecomputeautoscaler

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecomputeautoscaler/internal"
)

type GoogleComputeAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicasOutputReference interface {
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
	Fixed() *float64
	SetFixed(val *float64)
	FixedInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleComputeAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas
	SetInternalValue(val *GoogleComputeAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas)
	Percent() *float64
	SetPercent(val *float64)
	PercentInput() *float64
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
	ResetFixed()
	ResetPercent()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleComputeAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicasOutputReference
type jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicasOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicasOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicasOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicasOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicasOutputReference) Fixed() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fixed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicasOutputReference) FixedInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fixedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicasOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicasOutputReference) InternalValue() *GoogleComputeAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas {
	var returns *GoogleComputeAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicasOutputReference) Percent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"percent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicasOutputReference) PercentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"percentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicasOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicasOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleComputeAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicasOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleComputeAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicasOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleComputeAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicasOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicasOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleComputeAutoscaler.GoogleComputeAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicasOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleComputeAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicasOutputReference_Override(g GoogleComputeAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicasOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleComputeAutoscaler.GoogleComputeAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicasOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicasOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicasOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicasOutputReference)SetFixed(val *float64) {
	if err := j.validateSetFixedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fixed",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicasOutputReference)SetInternalValue(val *GoogleComputeAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicasOutputReference)SetPercent(val *float64) {
	if err := j.validateSetPercentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"percent",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicasOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicasOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicasOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicasOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicasOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicasOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicasOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicasOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicasOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicasOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicasOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicasOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicasOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicasOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicasOutputReference) ResetFixed() {
	_jsii_.InvokeVoid(
		g,
		"resetFixed",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicasOutputReference) ResetPercent() {
	_jsii_.InvokeVoid(
		g,
		"resetPercent",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicasOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicasOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

