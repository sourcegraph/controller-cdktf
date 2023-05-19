package googlecloudrunv2service

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecloudrunv2service/internal"
)

type GoogleCloudRunV2ServiceTemplateContainersStartupProbeOutputReference interface {
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
	FailureThreshold() *float64
	SetFailureThreshold(val *float64)
	FailureThresholdInput() *float64
	// Experimental.
	Fqn() *string
	Grpc() GoogleCloudRunV2ServiceTemplateContainersStartupProbeGrpcOutputReference
	GrpcInput() *GoogleCloudRunV2ServiceTemplateContainersStartupProbeGrpc
	HttpGet() GoogleCloudRunV2ServiceTemplateContainersStartupProbeHttpGetOutputReference
	HttpGetInput() *GoogleCloudRunV2ServiceTemplateContainersStartupProbeHttpGet
	InitialDelaySeconds() *float64
	SetInitialDelaySeconds(val *float64)
	InitialDelaySecondsInput() *float64
	InternalValue() *GoogleCloudRunV2ServiceTemplateContainersStartupProbe
	SetInternalValue(val *GoogleCloudRunV2ServiceTemplateContainersStartupProbe)
	PeriodSeconds() *float64
	SetPeriodSeconds(val *float64)
	PeriodSecondsInput() *float64
	TcpSocket() GoogleCloudRunV2ServiceTemplateContainersStartupProbeTcpSocketOutputReference
	TcpSocketInput() *GoogleCloudRunV2ServiceTemplateContainersStartupProbeTcpSocket
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
	PutGrpc(value *GoogleCloudRunV2ServiceTemplateContainersStartupProbeGrpc)
	PutHttpGet(value *GoogleCloudRunV2ServiceTemplateContainersStartupProbeHttpGet)
	PutTcpSocket(value *GoogleCloudRunV2ServiceTemplateContainersStartupProbeTcpSocket)
	ResetFailureThreshold()
	ResetGrpc()
	ResetHttpGet()
	ResetInitialDelaySeconds()
	ResetPeriodSeconds()
	ResetTcpSocket()
	ResetTimeoutSeconds()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleCloudRunV2ServiceTemplateContainersStartupProbeOutputReference
type jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersStartupProbeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersStartupProbeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersStartupProbeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersStartupProbeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersStartupProbeOutputReference) FailureThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"failureThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersStartupProbeOutputReference) FailureThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"failureThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersStartupProbeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersStartupProbeOutputReference) Grpc() GoogleCloudRunV2ServiceTemplateContainersStartupProbeGrpcOutputReference {
	var returns GoogleCloudRunV2ServiceTemplateContainersStartupProbeGrpcOutputReference
	_jsii_.Get(
		j,
		"grpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersStartupProbeOutputReference) GrpcInput() *GoogleCloudRunV2ServiceTemplateContainersStartupProbeGrpc {
	var returns *GoogleCloudRunV2ServiceTemplateContainersStartupProbeGrpc
	_jsii_.Get(
		j,
		"grpcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersStartupProbeOutputReference) HttpGet() GoogleCloudRunV2ServiceTemplateContainersStartupProbeHttpGetOutputReference {
	var returns GoogleCloudRunV2ServiceTemplateContainersStartupProbeHttpGetOutputReference
	_jsii_.Get(
		j,
		"httpGet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersStartupProbeOutputReference) HttpGetInput() *GoogleCloudRunV2ServiceTemplateContainersStartupProbeHttpGet {
	var returns *GoogleCloudRunV2ServiceTemplateContainersStartupProbeHttpGet
	_jsii_.Get(
		j,
		"httpGetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersStartupProbeOutputReference) InitialDelaySeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"initialDelaySeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersStartupProbeOutputReference) InitialDelaySecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"initialDelaySecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersStartupProbeOutputReference) InternalValue() *GoogleCloudRunV2ServiceTemplateContainersStartupProbe {
	var returns *GoogleCloudRunV2ServiceTemplateContainersStartupProbe
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersStartupProbeOutputReference) PeriodSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"periodSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersStartupProbeOutputReference) PeriodSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"periodSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersStartupProbeOutputReference) TcpSocket() GoogleCloudRunV2ServiceTemplateContainersStartupProbeTcpSocketOutputReference {
	var returns GoogleCloudRunV2ServiceTemplateContainersStartupProbeTcpSocketOutputReference
	_jsii_.Get(
		j,
		"tcpSocket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersStartupProbeOutputReference) TcpSocketInput() *GoogleCloudRunV2ServiceTemplateContainersStartupProbeTcpSocket {
	var returns *GoogleCloudRunV2ServiceTemplateContainersStartupProbeTcpSocket
	_jsii_.Get(
		j,
		"tcpSocketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersStartupProbeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersStartupProbeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersStartupProbeOutputReference) TimeoutSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersStartupProbeOutputReference) TimeoutSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutSecondsInput",
		&returns,
	)
	return returns
}


func NewGoogleCloudRunV2ServiceTemplateContainersStartupProbeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleCloudRunV2ServiceTemplateContainersStartupProbeOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleCloudRunV2ServiceTemplateContainersStartupProbeOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersStartupProbeOutputReference{}

	_jsii_.Create(
		"google-beta.googleCloudRunV2Service.GoogleCloudRunV2ServiceTemplateContainersStartupProbeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleCloudRunV2ServiceTemplateContainersStartupProbeOutputReference_Override(g GoogleCloudRunV2ServiceTemplateContainersStartupProbeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleCloudRunV2Service.GoogleCloudRunV2ServiceTemplateContainersStartupProbeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersStartupProbeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersStartupProbeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersStartupProbeOutputReference)SetFailureThreshold(val *float64) {
	if err := j.validateSetFailureThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"failureThreshold",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersStartupProbeOutputReference)SetInitialDelaySeconds(val *float64) {
	if err := j.validateSetInitialDelaySecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"initialDelaySeconds",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersStartupProbeOutputReference)SetInternalValue(val *GoogleCloudRunV2ServiceTemplateContainersStartupProbe) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersStartupProbeOutputReference)SetPeriodSeconds(val *float64) {
	if err := j.validateSetPeriodSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"periodSeconds",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersStartupProbeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersStartupProbeOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersStartupProbeOutputReference)SetTimeoutSeconds(val *float64) {
	if err := j.validateSetTimeoutSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeoutSeconds",
		val,
	)
}

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersStartupProbeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersStartupProbeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersStartupProbeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersStartupProbeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersStartupProbeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersStartupProbeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersStartupProbeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersStartupProbeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersStartupProbeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersStartupProbeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersStartupProbeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersStartupProbeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersStartupProbeOutputReference) PutGrpc(value *GoogleCloudRunV2ServiceTemplateContainersStartupProbeGrpc) {
	if err := g.validatePutGrpcParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGrpc",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersStartupProbeOutputReference) PutHttpGet(value *GoogleCloudRunV2ServiceTemplateContainersStartupProbeHttpGet) {
	if err := g.validatePutHttpGetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putHttpGet",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersStartupProbeOutputReference) PutTcpSocket(value *GoogleCloudRunV2ServiceTemplateContainersStartupProbeTcpSocket) {
	if err := g.validatePutTcpSocketParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTcpSocket",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersStartupProbeOutputReference) ResetFailureThreshold() {
	_jsii_.InvokeVoid(
		g,
		"resetFailureThreshold",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersStartupProbeOutputReference) ResetGrpc() {
	_jsii_.InvokeVoid(
		g,
		"resetGrpc",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersStartupProbeOutputReference) ResetHttpGet() {
	_jsii_.InvokeVoid(
		g,
		"resetHttpGet",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersStartupProbeOutputReference) ResetInitialDelaySeconds() {
	_jsii_.InvokeVoid(
		g,
		"resetInitialDelaySeconds",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersStartupProbeOutputReference) ResetPeriodSeconds() {
	_jsii_.InvokeVoid(
		g,
		"resetPeriodSeconds",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersStartupProbeOutputReference) ResetTcpSocket() {
	_jsii_.InvokeVoid(
		g,
		"resetTcpSocket",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersStartupProbeOutputReference) ResetTimeoutSeconds() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeoutSeconds",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersStartupProbeOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleCloudRunV2ServiceTemplateContainersStartupProbeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

