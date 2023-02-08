package googlecloudtasksqueue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecloudtasksqueue/internal"
)

type GoogleCloudTasksQueueRetryConfigOutputReference interface {
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
	InternalValue() *GoogleCloudTasksQueueRetryConfig
	SetInternalValue(val *GoogleCloudTasksQueueRetryConfig)
	MaxAttempts() *float64
	SetMaxAttempts(val *float64)
	MaxAttemptsInput() *float64
	MaxBackoff() *string
	SetMaxBackoff(val *string)
	MaxBackoffInput() *string
	MaxDoublings() *float64
	SetMaxDoublings(val *float64)
	MaxDoublingsInput() *float64
	MaxRetryDuration() *string
	SetMaxRetryDuration(val *string)
	MaxRetryDurationInput() *string
	MinBackoff() *string
	SetMinBackoff(val *string)
	MinBackoffInput() *string
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
	ResetMaxAttempts()
	ResetMaxBackoff()
	ResetMaxDoublings()
	ResetMaxRetryDuration()
	ResetMinBackoff()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleCloudTasksQueueRetryConfigOutputReference
type jsiiProxy_GoogleCloudTasksQueueRetryConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleCloudTasksQueueRetryConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudTasksQueueRetryConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudTasksQueueRetryConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudTasksQueueRetryConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudTasksQueueRetryConfigOutputReference) InternalValue() *GoogleCloudTasksQueueRetryConfig {
	var returns *GoogleCloudTasksQueueRetryConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudTasksQueueRetryConfigOutputReference) MaxAttempts() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAttempts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudTasksQueueRetryConfigOutputReference) MaxAttemptsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAttemptsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudTasksQueueRetryConfigOutputReference) MaxBackoff() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxBackoff",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudTasksQueueRetryConfigOutputReference) MaxBackoffInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxBackoffInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudTasksQueueRetryConfigOutputReference) MaxDoublings() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxDoublings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudTasksQueueRetryConfigOutputReference) MaxDoublingsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxDoublingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudTasksQueueRetryConfigOutputReference) MaxRetryDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxRetryDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudTasksQueueRetryConfigOutputReference) MaxRetryDurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxRetryDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudTasksQueueRetryConfigOutputReference) MinBackoff() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minBackoff",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudTasksQueueRetryConfigOutputReference) MinBackoffInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minBackoffInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudTasksQueueRetryConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudTasksQueueRetryConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleCloudTasksQueueRetryConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleCloudTasksQueueRetryConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleCloudTasksQueueRetryConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleCloudTasksQueueRetryConfigOutputReference{}

	_jsii_.Create(
		"google-beta.googleCloudTasksQueue.GoogleCloudTasksQueueRetryConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleCloudTasksQueueRetryConfigOutputReference_Override(g GoogleCloudTasksQueueRetryConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleCloudTasksQueue.GoogleCloudTasksQueueRetryConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleCloudTasksQueueRetryConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudTasksQueueRetryConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudTasksQueueRetryConfigOutputReference)SetInternalValue(val *GoogleCloudTasksQueueRetryConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudTasksQueueRetryConfigOutputReference)SetMaxAttempts(val *float64) {
	if err := j.validateSetMaxAttemptsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxAttempts",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudTasksQueueRetryConfigOutputReference)SetMaxBackoff(val *string) {
	if err := j.validateSetMaxBackoffParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxBackoff",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudTasksQueueRetryConfigOutputReference)SetMaxDoublings(val *float64) {
	if err := j.validateSetMaxDoublingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxDoublings",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudTasksQueueRetryConfigOutputReference)SetMaxRetryDuration(val *string) {
	if err := j.validateSetMaxRetryDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxRetryDuration",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudTasksQueueRetryConfigOutputReference)SetMinBackoff(val *string) {
	if err := j.validateSetMinBackoffParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minBackoff",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudTasksQueueRetryConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudTasksQueueRetryConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleCloudTasksQueueRetryConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudTasksQueueRetryConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleCloudTasksQueueRetryConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleCloudTasksQueueRetryConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleCloudTasksQueueRetryConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleCloudTasksQueueRetryConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleCloudTasksQueueRetryConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleCloudTasksQueueRetryConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleCloudTasksQueueRetryConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleCloudTasksQueueRetryConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleCloudTasksQueueRetryConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudTasksQueueRetryConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleCloudTasksQueueRetryConfigOutputReference) ResetMaxAttempts() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxAttempts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudTasksQueueRetryConfigOutputReference) ResetMaxBackoff() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxBackoff",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudTasksQueueRetryConfigOutputReference) ResetMaxDoublings() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxDoublings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudTasksQueueRetryConfigOutputReference) ResetMaxRetryDuration() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxRetryDuration",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudTasksQueueRetryConfigOutputReference) ResetMinBackoff() {
	_jsii_.InvokeVoid(
		g,
		"resetMinBackoff",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudTasksQueueRetryConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleCloudTasksQueueRetryConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

