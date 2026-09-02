package googlecloudrunv2workerpool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecloudrunv2workerpool/internal"
)

type GoogleCloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference interface {
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
	FailureThreshold() *float64
	SetFailureThreshold(val *float64)
	FailureThresholdInput() *float64
	// Experimental.
	Fqn() *string
	Grpc() GoogleCloudRunV2WorkerPoolTemplateContainersLivenessProbeGrpcOutputReference
	GrpcInput() *GoogleCloudRunV2WorkerPoolTemplateContainersLivenessProbeGrpc
	HttpGet() GoogleCloudRunV2WorkerPoolTemplateContainersLivenessProbeHttpGetOutputReference
	HttpGetInput() *GoogleCloudRunV2WorkerPoolTemplateContainersLivenessProbeHttpGet
	InitialDelaySeconds() *float64
	SetInitialDelaySeconds(val *float64)
	InitialDelaySecondsInput() *float64
	InternalValue() *GoogleCloudRunV2WorkerPoolTemplateContainersLivenessProbe
	SetInternalValue(val *GoogleCloudRunV2WorkerPoolTemplateContainersLivenessProbe)
	PeriodSeconds() *float64
	SetPeriodSeconds(val *float64)
	PeriodSecondsInput() *float64
	TcpSocket() GoogleCloudRunV2WorkerPoolTemplateContainersLivenessProbeTcpSocketOutputReference
	TcpSocketInput() *GoogleCloudRunV2WorkerPoolTemplateContainersLivenessProbeTcpSocket
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TimeoutSeconds() *float64
	SetTimeoutSeconds(val *float64)
	TimeoutSecondsInput() *float64
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
	PutGrpc(value *GoogleCloudRunV2WorkerPoolTemplateContainersLivenessProbeGrpc)
	PutHttpGet(value *GoogleCloudRunV2WorkerPoolTemplateContainersLivenessProbeHttpGet)
	PutTcpSocket(value *GoogleCloudRunV2WorkerPoolTemplateContainersLivenessProbeTcpSocket)
	ResetFailureThreshold()
	ResetGrpc()
	ResetHttpGet()
	ResetInitialDelaySeconds()
	ResetPeriodSeconds()
	ResetTcpSocket()
	ResetTimeoutSeconds()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleCloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference
type jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference) FailureThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"failureThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference) FailureThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"failureThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference) Grpc() GoogleCloudRunV2WorkerPoolTemplateContainersLivenessProbeGrpcOutputReference {
	var returns GoogleCloudRunV2WorkerPoolTemplateContainersLivenessProbeGrpcOutputReference
	_jsii_.Get(
		j,
		"grpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference) GrpcInput() *GoogleCloudRunV2WorkerPoolTemplateContainersLivenessProbeGrpc {
	var returns *GoogleCloudRunV2WorkerPoolTemplateContainersLivenessProbeGrpc
	_jsii_.Get(
		j,
		"grpcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference) HttpGet() GoogleCloudRunV2WorkerPoolTemplateContainersLivenessProbeHttpGetOutputReference {
	var returns GoogleCloudRunV2WorkerPoolTemplateContainersLivenessProbeHttpGetOutputReference
	_jsii_.Get(
		j,
		"httpGet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference) HttpGetInput() *GoogleCloudRunV2WorkerPoolTemplateContainersLivenessProbeHttpGet {
	var returns *GoogleCloudRunV2WorkerPoolTemplateContainersLivenessProbeHttpGet
	_jsii_.Get(
		j,
		"httpGetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference) InitialDelaySeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"initialDelaySeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference) InitialDelaySecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"initialDelaySecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference) InternalValue() *GoogleCloudRunV2WorkerPoolTemplateContainersLivenessProbe {
	var returns *GoogleCloudRunV2WorkerPoolTemplateContainersLivenessProbe
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference) PeriodSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"periodSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference) PeriodSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"periodSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference) TcpSocket() GoogleCloudRunV2WorkerPoolTemplateContainersLivenessProbeTcpSocketOutputReference {
	var returns GoogleCloudRunV2WorkerPoolTemplateContainersLivenessProbeTcpSocketOutputReference
	_jsii_.Get(
		j,
		"tcpSocket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference) TcpSocketInput() *GoogleCloudRunV2WorkerPoolTemplateContainersLivenessProbeTcpSocket {
	var returns *GoogleCloudRunV2WorkerPoolTemplateContainersLivenessProbeTcpSocket
	_jsii_.Get(
		j,
		"tcpSocketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference) TimeoutSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference) TimeoutSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutSecondsInput",
		&returns,
	)
	return returns
}


func NewGoogleCloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleCloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleCloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleCloudRunV2WorkerPool.GoogleCloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleCloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference_Override(g GoogleCloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleCloudRunV2WorkerPool.GoogleCloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference)SetFailureThreshold(val *float64) {
	if err := j.validateSetFailureThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"failureThreshold",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference)SetInitialDelaySeconds(val *float64) {
	if err := j.validateSetInitialDelaySecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"initialDelaySeconds",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference)SetInternalValue(val *GoogleCloudRunV2WorkerPoolTemplateContainersLivenessProbe) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference)SetPeriodSeconds(val *float64) {
	if err := j.validateSetPeriodSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"periodSeconds",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference)SetTimeoutSeconds(val *float64) {
	if err := j.validateSetTimeoutSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeoutSeconds",
		val,
	)
}

func (g *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference) PutGrpc(value *GoogleCloudRunV2WorkerPoolTemplateContainersLivenessProbeGrpc) {
	if err := g.validatePutGrpcParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGrpc",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference) PutHttpGet(value *GoogleCloudRunV2WorkerPoolTemplateContainersLivenessProbeHttpGet) {
	if err := g.validatePutHttpGetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putHttpGet",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference) PutTcpSocket(value *GoogleCloudRunV2WorkerPoolTemplateContainersLivenessProbeTcpSocket) {
	if err := g.validatePutTcpSocketParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTcpSocket",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference) ResetFailureThreshold() {
	_jsii_.InvokeVoid(
		g,
		"resetFailureThreshold",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference) ResetGrpc() {
	_jsii_.InvokeVoid(
		g,
		"resetGrpc",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference) ResetHttpGet() {
	_jsii_.InvokeVoid(
		g,
		"resetHttpGet",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference) ResetInitialDelaySeconds() {
	_jsii_.InvokeVoid(
		g,
		"resetInitialDelaySeconds",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference) ResetPeriodSeconds() {
	_jsii_.InvokeVoid(
		g,
		"resetPeriodSeconds",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference) ResetTcpSocket() {
	_jsii_.InvokeVoid(
		g,
		"resetTcpSocket",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference) ResetTimeoutSeconds() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeoutSeconds",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleCloudRunV2WorkerPoolTemplateContainersLivenessProbeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

