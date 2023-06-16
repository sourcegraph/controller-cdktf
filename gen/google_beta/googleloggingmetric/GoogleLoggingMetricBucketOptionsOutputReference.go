package googleloggingmetric

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleloggingmetric/internal"
)

type GoogleLoggingMetricBucketOptionsOutputReference interface {
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
	ExplicitBuckets() GoogleLoggingMetricBucketOptionsExplicitBucketsOutputReference
	ExplicitBucketsInput() *GoogleLoggingMetricBucketOptionsExplicitBuckets
	ExponentialBuckets() GoogleLoggingMetricBucketOptionsExponentialBucketsOutputReference
	ExponentialBucketsInput() *GoogleLoggingMetricBucketOptionsExponentialBuckets
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleLoggingMetricBucketOptions
	SetInternalValue(val *GoogleLoggingMetricBucketOptions)
	LinearBuckets() GoogleLoggingMetricBucketOptionsLinearBucketsOutputReference
	LinearBucketsInput() *GoogleLoggingMetricBucketOptionsLinearBuckets
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
	PutExplicitBuckets(value *GoogleLoggingMetricBucketOptionsExplicitBuckets)
	PutExponentialBuckets(value *GoogleLoggingMetricBucketOptionsExponentialBuckets)
	PutLinearBuckets(value *GoogleLoggingMetricBucketOptionsLinearBuckets)
	ResetExplicitBuckets()
	ResetExponentialBuckets()
	ResetLinearBuckets()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleLoggingMetricBucketOptionsOutputReference
type jsiiProxy_GoogleLoggingMetricBucketOptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleLoggingMetricBucketOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLoggingMetricBucketOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLoggingMetricBucketOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLoggingMetricBucketOptionsOutputReference) ExplicitBuckets() GoogleLoggingMetricBucketOptionsExplicitBucketsOutputReference {
	var returns GoogleLoggingMetricBucketOptionsExplicitBucketsOutputReference
	_jsii_.Get(
		j,
		"explicitBuckets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLoggingMetricBucketOptionsOutputReference) ExplicitBucketsInput() *GoogleLoggingMetricBucketOptionsExplicitBuckets {
	var returns *GoogleLoggingMetricBucketOptionsExplicitBuckets
	_jsii_.Get(
		j,
		"explicitBucketsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLoggingMetricBucketOptionsOutputReference) ExponentialBuckets() GoogleLoggingMetricBucketOptionsExponentialBucketsOutputReference {
	var returns GoogleLoggingMetricBucketOptionsExponentialBucketsOutputReference
	_jsii_.Get(
		j,
		"exponentialBuckets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLoggingMetricBucketOptionsOutputReference) ExponentialBucketsInput() *GoogleLoggingMetricBucketOptionsExponentialBuckets {
	var returns *GoogleLoggingMetricBucketOptionsExponentialBuckets
	_jsii_.Get(
		j,
		"exponentialBucketsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLoggingMetricBucketOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLoggingMetricBucketOptionsOutputReference) InternalValue() *GoogleLoggingMetricBucketOptions {
	var returns *GoogleLoggingMetricBucketOptions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLoggingMetricBucketOptionsOutputReference) LinearBuckets() GoogleLoggingMetricBucketOptionsLinearBucketsOutputReference {
	var returns GoogleLoggingMetricBucketOptionsLinearBucketsOutputReference
	_jsii_.Get(
		j,
		"linearBuckets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLoggingMetricBucketOptionsOutputReference) LinearBucketsInput() *GoogleLoggingMetricBucketOptionsLinearBuckets {
	var returns *GoogleLoggingMetricBucketOptionsLinearBuckets
	_jsii_.Get(
		j,
		"linearBucketsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLoggingMetricBucketOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLoggingMetricBucketOptionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleLoggingMetricBucketOptionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleLoggingMetricBucketOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleLoggingMetricBucketOptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleLoggingMetricBucketOptionsOutputReference{}

	_jsii_.Create(
		"google-beta.googleLoggingMetric.GoogleLoggingMetricBucketOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleLoggingMetricBucketOptionsOutputReference_Override(g GoogleLoggingMetricBucketOptionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleLoggingMetric.GoogleLoggingMetricBucketOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleLoggingMetricBucketOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleLoggingMetricBucketOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleLoggingMetricBucketOptionsOutputReference)SetInternalValue(val *GoogleLoggingMetricBucketOptions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleLoggingMetricBucketOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleLoggingMetricBucketOptionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleLoggingMetricBucketOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleLoggingMetricBucketOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleLoggingMetricBucketOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleLoggingMetricBucketOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleLoggingMetricBucketOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleLoggingMetricBucketOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleLoggingMetricBucketOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleLoggingMetricBucketOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleLoggingMetricBucketOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleLoggingMetricBucketOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleLoggingMetricBucketOptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleLoggingMetricBucketOptionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleLoggingMetricBucketOptionsOutputReference) PutExplicitBuckets(value *GoogleLoggingMetricBucketOptionsExplicitBuckets) {
	if err := g.validatePutExplicitBucketsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putExplicitBuckets",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleLoggingMetricBucketOptionsOutputReference) PutExponentialBuckets(value *GoogleLoggingMetricBucketOptionsExponentialBuckets) {
	if err := g.validatePutExponentialBucketsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putExponentialBuckets",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleLoggingMetricBucketOptionsOutputReference) PutLinearBuckets(value *GoogleLoggingMetricBucketOptionsLinearBuckets) {
	if err := g.validatePutLinearBucketsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putLinearBuckets",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleLoggingMetricBucketOptionsOutputReference) ResetExplicitBuckets() {
	_jsii_.InvokeVoid(
		g,
		"resetExplicitBuckets",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleLoggingMetricBucketOptionsOutputReference) ResetExponentialBuckets() {
	_jsii_.InvokeVoid(
		g,
		"resetExponentialBuckets",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleLoggingMetricBucketOptionsOutputReference) ResetLinearBuckets() {
	_jsii_.InvokeVoid(
		g,
		"resetLinearBuckets",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleLoggingMetricBucketOptionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleLoggingMetricBucketOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

