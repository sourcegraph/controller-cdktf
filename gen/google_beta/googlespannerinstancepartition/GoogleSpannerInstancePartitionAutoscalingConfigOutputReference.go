package googlespannerinstancepartition

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlespannerinstancepartition/internal"
)

type GoogleSpannerInstancePartitionAutoscalingConfigOutputReference interface {
	cdktf.ComplexObject
	AutoscalingLimits() GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingLimitsOutputReference
	AutoscalingLimitsInput() *GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingLimits
	AutoscalingTargets() GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference
	AutoscalingTargetsInput() *GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingTargets
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
	InternalValue() *GoogleSpannerInstancePartitionAutoscalingConfig
	SetInternalValue(val *GoogleSpannerInstancePartitionAutoscalingConfig)
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
	PutAutoscalingLimits(value *GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingLimits)
	PutAutoscalingTargets(value *GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingTargets)
	ResetAutoscalingLimits()
	ResetAutoscalingTargets()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleSpannerInstancePartitionAutoscalingConfigOutputReference
type jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigOutputReference) AutoscalingLimits() GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingLimitsOutputReference {
	var returns GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingLimitsOutputReference
	_jsii_.Get(
		j,
		"autoscalingLimits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigOutputReference) AutoscalingLimitsInput() *GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingLimits {
	var returns *GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingLimits
	_jsii_.Get(
		j,
		"autoscalingLimitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigOutputReference) AutoscalingTargets() GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference {
	var returns GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingTargetsOutputReference
	_jsii_.Get(
		j,
		"autoscalingTargets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigOutputReference) AutoscalingTargetsInput() *GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingTargets {
	var returns *GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingTargets
	_jsii_.Get(
		j,
		"autoscalingTargetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigOutputReference) InternalValue() *GoogleSpannerInstancePartitionAutoscalingConfig {
	var returns *GoogleSpannerInstancePartitionAutoscalingConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleSpannerInstancePartitionAutoscalingConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleSpannerInstancePartitionAutoscalingConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleSpannerInstancePartitionAutoscalingConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleSpannerInstancePartition.GoogleSpannerInstancePartitionAutoscalingConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleSpannerInstancePartitionAutoscalingConfigOutputReference_Override(g GoogleSpannerInstancePartitionAutoscalingConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleSpannerInstancePartition.GoogleSpannerInstancePartitionAutoscalingConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigOutputReference)SetInternalValue(val *GoogleSpannerInstancePartitionAutoscalingConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigOutputReference) PutAutoscalingLimits(value *GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingLimits) {
	if err := g.validatePutAutoscalingLimitsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAutoscalingLimits",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigOutputReference) PutAutoscalingTargets(value *GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingTargets) {
	if err := g.validatePutAutoscalingTargetsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAutoscalingTargets",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigOutputReference) ResetAutoscalingLimits() {
	_jsii_.InvokeVoid(
		g,
		"resetAutoscalingLimits",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigOutputReference) ResetAutoscalingTargets() {
	_jsii_.InvokeVoid(
		g,
		"resetAutoscalingTargets",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

