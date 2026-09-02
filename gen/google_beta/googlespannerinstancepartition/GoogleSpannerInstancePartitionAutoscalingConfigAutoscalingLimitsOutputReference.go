package googlespannerinstancepartition

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlespannerinstancepartition/internal"
)

type GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingLimitsOutputReference interface {
	cdktn.ComplexObject
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
	InternalValue() *GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingLimits
	SetInternalValue(val *GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingLimits)
	MaxNodes() *float64
	SetMaxNodes(val *float64)
	MaxNodesInput() *float64
	MaxProcessingUnits() *float64
	SetMaxProcessingUnits(val *float64)
	MaxProcessingUnitsInput() *float64
	MinNodes() *float64
	SetMinNodes(val *float64)
	MinNodesInput() *float64
	MinProcessingUnits() *float64
	SetMinProcessingUnits(val *float64)
	MinProcessingUnitsInput() *float64
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
	ResetMaxNodes()
	ResetMaxProcessingUnits()
	ResetMinNodes()
	ResetMinProcessingUnits()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingLimitsOutputReference
type jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingLimitsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingLimitsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingLimitsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingLimitsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingLimitsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingLimitsOutputReference) InternalValue() *GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingLimits {
	var returns *GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingLimits
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingLimitsOutputReference) MaxNodes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxNodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingLimitsOutputReference) MaxNodesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxNodesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingLimitsOutputReference) MaxProcessingUnits() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxProcessingUnits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingLimitsOutputReference) MaxProcessingUnitsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxProcessingUnitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingLimitsOutputReference) MinNodes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minNodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingLimitsOutputReference) MinNodesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minNodesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingLimitsOutputReference) MinProcessingUnits() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minProcessingUnits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingLimitsOutputReference) MinProcessingUnitsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minProcessingUnitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingLimitsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingLimitsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleSpannerInstancePartitionAutoscalingConfigAutoscalingLimitsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingLimitsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleSpannerInstancePartitionAutoscalingConfigAutoscalingLimitsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingLimitsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleSpannerInstancePartition.GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingLimitsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleSpannerInstancePartitionAutoscalingConfigAutoscalingLimitsOutputReference_Override(g GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingLimitsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleSpannerInstancePartition.GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingLimitsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingLimitsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingLimitsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingLimitsOutputReference)SetInternalValue(val *GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingLimits) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingLimitsOutputReference)SetMaxNodes(val *float64) {
	if err := j.validateSetMaxNodesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxNodes",
		val,
	)
}

func (j *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingLimitsOutputReference)SetMaxProcessingUnits(val *float64) {
	if err := j.validateSetMaxProcessingUnitsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxProcessingUnits",
		val,
	)
}

func (j *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingLimitsOutputReference)SetMinNodes(val *float64) {
	if err := j.validateSetMinNodesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minNodes",
		val,
	)
}

func (j *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingLimitsOutputReference)SetMinProcessingUnits(val *float64) {
	if err := j.validateSetMinProcessingUnitsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minProcessingUnits",
		val,
	)
}

func (j *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingLimitsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingLimitsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingLimitsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingLimitsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingLimitsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingLimitsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingLimitsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingLimitsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingLimitsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingLimitsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingLimitsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingLimitsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingLimitsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingLimitsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingLimitsOutputReference) ResetMaxNodes() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxNodes",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingLimitsOutputReference) ResetMaxProcessingUnits() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxProcessingUnits",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingLimitsOutputReference) ResetMinNodes() {
	_jsii_.InvokeVoid(
		g,
		"resetMinNodes",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingLimitsOutputReference) ResetMinProcessingUnits() {
	_jsii_.InvokeVoid(
		g,
		"resetMinProcessingUnits",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingLimitsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleSpannerInstancePartitionAutoscalingConfigAutoscalingLimitsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

