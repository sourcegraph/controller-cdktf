package cloudrunv2workerpool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/cloudrunv2workerpool/internal"
)

type CloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference interface {
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
	FailureThreshold() *float64
	SetFailureThreshold(val *float64)
	FailureThresholdInput() *float64
	// Experimental.
	Fqn() *string
	Grpc() CloudRunV2WorkerPoolTemplateContainersLivenessProbeGrpcOutputReference
	GrpcInput() *CloudRunV2WorkerPoolTemplateContainersLivenessProbeGrpc
	HttpGet() CloudRunV2WorkerPoolTemplateContainersLivenessProbeHttpGetOutputReference
	HttpGetInput() *CloudRunV2WorkerPoolTemplateContainersLivenessProbeHttpGet
	InitialDelaySeconds() *float64
	SetInitialDelaySeconds(val *float64)
	InitialDelaySecondsInput() *float64
	InternalValue() *CloudRunV2WorkerPoolTemplateContainersLivenessProbe
	SetInternalValue(val *CloudRunV2WorkerPoolTemplateContainersLivenessProbe)
	PeriodSeconds() *float64
	SetPeriodSeconds(val *float64)
	PeriodSecondsInput() *float64
	TcpSocket() CloudRunV2WorkerPoolTemplateContainersLivenessProbeTcpSocketOutputReference
	TcpSocketInput() *CloudRunV2WorkerPoolTemplateContainersLivenessProbeTcpSocket
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimeoutSeconds() *float64
	SetTimeoutSeconds(val *float64)
	TimeoutSecondsInput() *float64
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
	PutGrpc(value *CloudRunV2WorkerPoolTemplateContainersLivenessProbeGrpc)
	PutHttpGet(value *CloudRunV2WorkerPoolTemplateContainersLivenessProbeHttpGet)
	PutTcpSocket(value *CloudRunV2WorkerPoolTemplateContainersLivenessProbeTcpSocket)
	ResetFailureThreshold()
	ResetGrpc()
	ResetHttpGet()
	ResetInitialDelaySeconds()
	ResetPeriodSeconds()
	ResetTcpSocket()
	ResetTimeoutSeconds()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference
type jsiiProxy_CloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference) FailureThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"failureThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference) FailureThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"failureThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference) Grpc() CloudRunV2WorkerPoolTemplateContainersLivenessProbeGrpcOutputReference {
	var returns CloudRunV2WorkerPoolTemplateContainersLivenessProbeGrpcOutputReference
	_jsii_.Get(
		j,
		"grpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference) GrpcInput() *CloudRunV2WorkerPoolTemplateContainersLivenessProbeGrpc {
	var returns *CloudRunV2WorkerPoolTemplateContainersLivenessProbeGrpc
	_jsii_.Get(
		j,
		"grpcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference) HttpGet() CloudRunV2WorkerPoolTemplateContainersLivenessProbeHttpGetOutputReference {
	var returns CloudRunV2WorkerPoolTemplateContainersLivenessProbeHttpGetOutputReference
	_jsii_.Get(
		j,
		"httpGet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference) HttpGetInput() *CloudRunV2WorkerPoolTemplateContainersLivenessProbeHttpGet {
	var returns *CloudRunV2WorkerPoolTemplateContainersLivenessProbeHttpGet
	_jsii_.Get(
		j,
		"httpGetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference) InitialDelaySeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"initialDelaySeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference) InitialDelaySecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"initialDelaySecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference) InternalValue() *CloudRunV2WorkerPoolTemplateContainersLivenessProbe {
	var returns *CloudRunV2WorkerPoolTemplateContainersLivenessProbe
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference) PeriodSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"periodSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference) PeriodSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"periodSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference) TcpSocket() CloudRunV2WorkerPoolTemplateContainersLivenessProbeTcpSocketOutputReference {
	var returns CloudRunV2WorkerPoolTemplateContainersLivenessProbeTcpSocketOutputReference
	_jsii_.Get(
		j,
		"tcpSocket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference) TcpSocketInput() *CloudRunV2WorkerPoolTemplateContainersLivenessProbeTcpSocket {
	var returns *CloudRunV2WorkerPoolTemplateContainersLivenessProbeTcpSocket
	_jsii_.Get(
		j,
		"tcpSocketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference) TimeoutSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference) TimeoutSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutSecondsInput",
		&returns,
	)
	return returns
}

func NewCloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference {
	_init_.Initialize()

	if err := validateNewCloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.cloudRunV2WorkerPool.CloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference_Override(c CloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.cloudRunV2WorkerPool.CloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference",
		[]any{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference) SetFailureThreshold(val *float64) {
	if err := j.validateSetFailureThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"failureThreshold",
		val,
	)
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference) SetInitialDelaySeconds(val *float64) {
	if err := j.validateSetInitialDelaySecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"initialDelaySeconds",
		val,
	)
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference) SetInternalValue(val *CloudRunV2WorkerPoolTemplateContainersLivenessProbe) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference) SetPeriodSeconds(val *float64) {
	if err := j.validateSetPeriodSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"periodSeconds",
		val,
	)
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference) SetTimeoutSeconds(val *float64) {
	if err := j.validateSetTimeoutSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeoutSeconds",
		val,
	)
}

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference) PutGrpc(value *CloudRunV2WorkerPoolTemplateContainersLivenessProbeGrpc) {
	if err := c.validatePutGrpcParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putGrpc",
		[]any{value},
	)
}

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference) PutHttpGet(value *CloudRunV2WorkerPoolTemplateContainersLivenessProbeHttpGet) {
	if err := c.validatePutHttpGetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putHttpGet",
		[]any{value},
	)
}

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference) PutTcpSocket(value *CloudRunV2WorkerPoolTemplateContainersLivenessProbeTcpSocket) {
	if err := c.validatePutTcpSocketParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTcpSocket",
		[]any{value},
	)
}

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference) ResetFailureThreshold() {
	_jsii_.InvokeVoid(
		c,
		"resetFailureThreshold",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference) ResetGrpc() {
	_jsii_.InvokeVoid(
		c,
		"resetGrpc",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference) ResetHttpGet() {
	_jsii_.InvokeVoid(
		c,
		"resetHttpGet",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference) ResetInitialDelaySeconds() {
	_jsii_.InvokeVoid(
		c,
		"resetInitialDelaySeconds",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference) ResetPeriodSeconds() {
	_jsii_.InvokeVoid(
		c,
		"resetPeriodSeconds",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference) ResetTcpSocket() {
	_jsii_.InvokeVoid(
		c,
		"resetTcpSocket",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference) ResetTimeoutSeconds() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeoutSeconds",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
