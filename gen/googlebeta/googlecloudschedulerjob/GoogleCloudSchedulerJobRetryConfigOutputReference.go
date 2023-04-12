package googlecloudschedulerjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/googlebeta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/googlebeta/googlecloudschedulerjob/internal"
)

type GoogleCloudSchedulerJobRetryConfigOutputReference interface {
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
	InternalValue() *GoogleCloudSchedulerJobRetryConfig
	SetInternalValue(val *GoogleCloudSchedulerJobRetryConfig)
	MaxBackoffDuration() *string
	SetMaxBackoffDuration(val *string)
	MaxBackoffDurationInput() *string
	MaxDoublings() *float64
	SetMaxDoublings(val *float64)
	MaxDoublingsInput() *float64
	MaxRetryDuration() *string
	SetMaxRetryDuration(val *string)
	MaxRetryDurationInput() *string
	MinBackoffDuration() *string
	SetMinBackoffDuration(val *string)
	MinBackoffDurationInput() *string
	RetryCount() *float64
	SetRetryCount(val *float64)
	RetryCountInput() *float64
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
	ResetMaxBackoffDuration()
	ResetMaxDoublings()
	ResetMaxRetryDuration()
	ResetMinBackoffDuration()
	ResetRetryCount()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleCloudSchedulerJobRetryConfigOutputReference
type jsiiProxy_GoogleCloudSchedulerJobRetryConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleCloudSchedulerJobRetryConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSchedulerJobRetryConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSchedulerJobRetryConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSchedulerJobRetryConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSchedulerJobRetryConfigOutputReference) InternalValue() *GoogleCloudSchedulerJobRetryConfig {
	var returns *GoogleCloudSchedulerJobRetryConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSchedulerJobRetryConfigOutputReference) MaxBackoffDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxBackoffDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSchedulerJobRetryConfigOutputReference) MaxBackoffDurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxBackoffDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSchedulerJobRetryConfigOutputReference) MaxDoublings() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxDoublings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSchedulerJobRetryConfigOutputReference) MaxDoublingsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxDoublingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSchedulerJobRetryConfigOutputReference) MaxRetryDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxRetryDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSchedulerJobRetryConfigOutputReference) MaxRetryDurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxRetryDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSchedulerJobRetryConfigOutputReference) MinBackoffDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minBackoffDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSchedulerJobRetryConfigOutputReference) MinBackoffDurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minBackoffDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSchedulerJobRetryConfigOutputReference) RetryCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retryCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSchedulerJobRetryConfigOutputReference) RetryCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retryCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSchedulerJobRetryConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSchedulerJobRetryConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleCloudSchedulerJobRetryConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleCloudSchedulerJobRetryConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleCloudSchedulerJobRetryConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleCloudSchedulerJobRetryConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-googlebeta.googleCloudSchedulerJob.GoogleCloudSchedulerJobRetryConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleCloudSchedulerJobRetryConfigOutputReference_Override(g GoogleCloudSchedulerJobRetryConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-googlebeta.googleCloudSchedulerJob.GoogleCloudSchedulerJobRetryConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleCloudSchedulerJobRetryConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudSchedulerJobRetryConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudSchedulerJobRetryConfigOutputReference)SetInternalValue(val *GoogleCloudSchedulerJobRetryConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudSchedulerJobRetryConfigOutputReference)SetMaxBackoffDuration(val *string) {
	if err := j.validateSetMaxBackoffDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxBackoffDuration",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudSchedulerJobRetryConfigOutputReference)SetMaxDoublings(val *float64) {
	if err := j.validateSetMaxDoublingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxDoublings",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudSchedulerJobRetryConfigOutputReference)SetMaxRetryDuration(val *string) {
	if err := j.validateSetMaxRetryDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxRetryDuration",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudSchedulerJobRetryConfigOutputReference)SetMinBackoffDuration(val *string) {
	if err := j.validateSetMinBackoffDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minBackoffDuration",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudSchedulerJobRetryConfigOutputReference)SetRetryCount(val *float64) {
	if err := j.validateSetRetryCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retryCount",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudSchedulerJobRetryConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudSchedulerJobRetryConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleCloudSchedulerJobRetryConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudSchedulerJobRetryConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleCloudSchedulerJobRetryConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleCloudSchedulerJobRetryConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleCloudSchedulerJobRetryConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleCloudSchedulerJobRetryConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleCloudSchedulerJobRetryConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleCloudSchedulerJobRetryConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleCloudSchedulerJobRetryConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleCloudSchedulerJobRetryConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleCloudSchedulerJobRetryConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudSchedulerJobRetryConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleCloudSchedulerJobRetryConfigOutputReference) ResetMaxBackoffDuration() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxBackoffDuration",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudSchedulerJobRetryConfigOutputReference) ResetMaxDoublings() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxDoublings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudSchedulerJobRetryConfigOutputReference) ResetMaxRetryDuration() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxRetryDuration",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudSchedulerJobRetryConfigOutputReference) ResetMinBackoffDuration() {
	_jsii_.InvokeVoid(
		g,
		"resetMinBackoffDuration",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudSchedulerJobRetryConfigOutputReference) ResetRetryCount() {
	_jsii_.InvokeVoid(
		g,
		"resetRetryCount",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudSchedulerJobRetryConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleCloudSchedulerJobRetryConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

