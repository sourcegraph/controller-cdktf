package googledataprocautoscalingpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledataprocautoscalingpolicy/internal"
)

type GoogleDataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference interface {
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
	GracefulDecommissionTimeout() *string
	SetGracefulDecommissionTimeout(val *string)
	GracefulDecommissionTimeoutInput() *string
	InternalValue() *GoogleDataprocAutoscalingPolicyBasicAlgorithmYarnConfig
	SetInternalValue(val *GoogleDataprocAutoscalingPolicyBasicAlgorithmYarnConfig)
	ScaleDownFactor() *float64
	SetScaleDownFactor(val *float64)
	ScaleDownFactorInput() *float64
	ScaleDownMinWorkerFraction() *float64
	SetScaleDownMinWorkerFraction(val *float64)
	ScaleDownMinWorkerFractionInput() *float64
	ScaleUpFactor() *float64
	SetScaleUpFactor(val *float64)
	ScaleUpFactorInput() *float64
	ScaleUpMinWorkerFraction() *float64
	SetScaleUpMinWorkerFraction(val *float64)
	ScaleUpMinWorkerFractionInput() *float64
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
	ResetScaleDownMinWorkerFraction()
	ResetScaleUpMinWorkerFraction()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference
type jsiiProxy_GoogleDataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference) GracefulDecommissionTimeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gracefulDecommissionTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference) GracefulDecommissionTimeoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gracefulDecommissionTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference) InternalValue() *GoogleDataprocAutoscalingPolicyBasicAlgorithmYarnConfig {
	var returns *GoogleDataprocAutoscalingPolicyBasicAlgorithmYarnConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference) ScaleDownFactor() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scaleDownFactor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference) ScaleDownFactorInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scaleDownFactorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference) ScaleDownMinWorkerFraction() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scaleDownMinWorkerFraction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference) ScaleDownMinWorkerFractionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scaleDownMinWorkerFractionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference) ScaleUpFactor() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scaleUpFactor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference) ScaleUpFactorInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scaleUpFactorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference) ScaleUpMinWorkerFraction() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scaleUpMinWorkerFraction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference) ScaleUpMinWorkerFractionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scaleUpMinWorkerFractionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleDataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleDataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDataprocAutoscalingPolicy.GoogleDataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference_Override(g GoogleDataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDataprocAutoscalingPolicy.GoogleDataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference)SetGracefulDecommissionTimeout(val *string) {
	if err := j.validateSetGracefulDecommissionTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gracefulDecommissionTimeout",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference)SetInternalValue(val *GoogleDataprocAutoscalingPolicyBasicAlgorithmYarnConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference)SetScaleDownFactor(val *float64) {
	if err := j.validateSetScaleDownFactorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scaleDownFactor",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference)SetScaleDownMinWorkerFraction(val *float64) {
	if err := j.validateSetScaleDownMinWorkerFractionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scaleDownMinWorkerFraction",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference)SetScaleUpFactor(val *float64) {
	if err := j.validateSetScaleUpFactorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scaleUpFactor",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference)SetScaleUpMinWorkerFraction(val *float64) {
	if err := j.validateSetScaleUpMinWorkerFractionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scaleUpMinWorkerFraction",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference) ResetScaleDownMinWorkerFraction() {
	_jsii_.InvokeVoid(
		g,
		"resetScaleDownMinWorkerFraction",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference) ResetScaleUpMinWorkerFraction() {
	_jsii_.InvokeVoid(
		g,
		"resetScaleUpMinWorkerFraction",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

