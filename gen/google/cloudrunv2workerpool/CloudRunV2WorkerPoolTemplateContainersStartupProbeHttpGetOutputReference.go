package cloudrunv2workerpool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/cloudrunv2workerpool/internal"
)

type CloudRunV2WorkerPoolTemplateContainersStartupProbeHttpGetOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() any
	// Experimental.
	SetComplexObjectIndex(val any)
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
	HttpHeaders() CloudRunV2WorkerPoolTemplateContainersStartupProbeHttpGetHttpHeadersOutputReference
	HttpHeadersInput() *CloudRunV2WorkerPoolTemplateContainersStartupProbeHttpGetHttpHeaders
	InternalValue() *CloudRunV2WorkerPoolTemplateContainersStartupProbeHttpGet
	SetInternalValue(val *CloudRunV2WorkerPoolTemplateContainersStartupProbeHttpGet)
	Path() *string
	SetPath(val *string)
	PathInput() *string
	Port() *float64
	SetPort(val *float64)
	PortInput() *float64
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
	GetAnyMapAttribute(terraformAttribute *string) *map[string]any
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
	PutHttpHeaders(value *CloudRunV2WorkerPoolTemplateContainersStartupProbeHttpGetHttpHeaders)
	ResetHttpHeaders()
	ResetPath()
	ResetPort()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CloudRunV2WorkerPoolTemplateContainersStartupProbeHttpGetOutputReference
type jsiiProxy_CloudRunV2WorkerPoolTemplateContainersStartupProbeHttpGetOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersStartupProbeHttpGetOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersStartupProbeHttpGetOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersStartupProbeHttpGetOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersStartupProbeHttpGetOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersStartupProbeHttpGetOutputReference) HttpHeaders() CloudRunV2WorkerPoolTemplateContainersStartupProbeHttpGetHttpHeadersOutputReference {
	var returns CloudRunV2WorkerPoolTemplateContainersStartupProbeHttpGetHttpHeadersOutputReference
	_jsii_.Get(
		j,
		"httpHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersStartupProbeHttpGetOutputReference) HttpHeadersInput() *CloudRunV2WorkerPoolTemplateContainersStartupProbeHttpGetHttpHeaders {
	var returns *CloudRunV2WorkerPoolTemplateContainersStartupProbeHttpGetHttpHeaders
	_jsii_.Get(
		j,
		"httpHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersStartupProbeHttpGetOutputReference) InternalValue() *CloudRunV2WorkerPoolTemplateContainersStartupProbeHttpGet {
	var returns *CloudRunV2WorkerPoolTemplateContainersStartupProbeHttpGet
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersStartupProbeHttpGetOutputReference) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersStartupProbeHttpGetOutputReference) PathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersStartupProbeHttpGetOutputReference) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersStartupProbeHttpGetOutputReference) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersStartupProbeHttpGetOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersStartupProbeHttpGetOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewCloudRunV2WorkerPoolTemplateContainersStartupProbeHttpGetOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CloudRunV2WorkerPoolTemplateContainersStartupProbeHttpGetOutputReference {
	_init_.Initialize()

	if err := validateNewCloudRunV2WorkerPoolTemplateContainersStartupProbeHttpGetOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudRunV2WorkerPoolTemplateContainersStartupProbeHttpGetOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.cloudRunV2WorkerPool.CloudRunV2WorkerPoolTemplateContainersStartupProbeHttpGetOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCloudRunV2WorkerPoolTemplateContainersStartupProbeHttpGetOutputReference_Override(c CloudRunV2WorkerPoolTemplateContainersStartupProbeHttpGetOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.cloudRunV2WorkerPool.CloudRunV2WorkerPoolTemplateContainersStartupProbeHttpGetOutputReference",
		[]any{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersStartupProbeHttpGetOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersStartupProbeHttpGetOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersStartupProbeHttpGetOutputReference) SetInternalValue(val *CloudRunV2WorkerPoolTemplateContainersStartupProbeHttpGet) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersStartupProbeHttpGetOutputReference) SetPath(val *string) {
	if err := j.validateSetPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"path",
		val,
	)
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersStartupProbeHttpGetOutputReference) SetPort(val *float64) {
	if err := j.validateSetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersStartupProbeHttpGetOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersStartupProbeHttpGetOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersStartupProbeHttpGetOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersStartupProbeHttpGetOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersStartupProbeHttpGetOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersStartupProbeHttpGetOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersStartupProbeHttpGetOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersStartupProbeHttpGetOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersStartupProbeHttpGetOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersStartupProbeHttpGetOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersStartupProbeHttpGetOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersStartupProbeHttpGetOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersStartupProbeHttpGetOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersStartupProbeHttpGetOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersStartupProbeHttpGetOutputReference) PutHttpHeaders(value *CloudRunV2WorkerPoolTemplateContainersStartupProbeHttpGetHttpHeaders) {
	if err := c.validatePutHttpHeadersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putHttpHeaders",
		[]any{value},
	)
}

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersStartupProbeHttpGetOutputReference) ResetHttpHeaders() {
	_jsii_.InvokeVoid(
		c,
		"resetHttpHeaders",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersStartupProbeHttpGetOutputReference) ResetPath() {
	_jsii_.InvokeVoid(
		c,
		"resetPath",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersStartupProbeHttpGetOutputReference) ResetPort() {
	_jsii_.InvokeVoid(
		c,
		"resetPort",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersStartupProbeHttpGetOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		c,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersStartupProbeHttpGetOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
