package googlecomputehealthcheck

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/googlebeta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/googlebeta/googlecomputehealthcheck/internal"
)

type GoogleComputeHealthCheckHttpsHealthCheckOutputReference interface {
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
	Host() *string
	SetHost(val *string)
	HostInput() *string
	InternalValue() *GoogleComputeHealthCheckHttpsHealthCheck
	SetInternalValue(val *GoogleComputeHealthCheckHttpsHealthCheck)
	Port() *float64
	SetPort(val *float64)
	PortInput() *float64
	PortName() *string
	SetPortName(val *string)
	PortNameInput() *string
	PortSpecification() *string
	SetPortSpecification(val *string)
	PortSpecificationInput() *string
	ProxyHeader() *string
	SetProxyHeader(val *string)
	ProxyHeaderInput() *string
	RequestPath() *string
	SetRequestPath(val *string)
	RequestPathInput() *string
	Response() *string
	SetResponse(val *string)
	ResponseInput() *string
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
	ResetHost()
	ResetPort()
	ResetPortName()
	ResetPortSpecification()
	ResetProxyHeader()
	ResetRequestPath()
	ResetResponse()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleComputeHealthCheckHttpsHealthCheckOutputReference
type jsiiProxy_GoogleComputeHealthCheckHttpsHealthCheckOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleComputeHealthCheckHttpsHealthCheckOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeHealthCheckHttpsHealthCheckOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeHealthCheckHttpsHealthCheckOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeHealthCheckHttpsHealthCheckOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeHealthCheckHttpsHealthCheckOutputReference) Host() *string {
	var returns *string
	_jsii_.Get(
		j,
		"host",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeHealthCheckHttpsHealthCheckOutputReference) HostInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeHealthCheckHttpsHealthCheckOutputReference) InternalValue() *GoogleComputeHealthCheckHttpsHealthCheck {
	var returns *GoogleComputeHealthCheckHttpsHealthCheck
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeHealthCheckHttpsHealthCheckOutputReference) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeHealthCheckHttpsHealthCheckOutputReference) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeHealthCheckHttpsHealthCheckOutputReference) PortName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"portName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeHealthCheckHttpsHealthCheckOutputReference) PortNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"portNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeHealthCheckHttpsHealthCheckOutputReference) PortSpecification() *string {
	var returns *string
	_jsii_.Get(
		j,
		"portSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeHealthCheckHttpsHealthCheckOutputReference) PortSpecificationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"portSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeHealthCheckHttpsHealthCheckOutputReference) ProxyHeader() *string {
	var returns *string
	_jsii_.Get(
		j,
		"proxyHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeHealthCheckHttpsHealthCheckOutputReference) ProxyHeaderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"proxyHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeHealthCheckHttpsHealthCheckOutputReference) RequestPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeHealthCheckHttpsHealthCheckOutputReference) RequestPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeHealthCheckHttpsHealthCheckOutputReference) Response() *string {
	var returns *string
	_jsii_.Get(
		j,
		"response",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeHealthCheckHttpsHealthCheckOutputReference) ResponseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeHealthCheckHttpsHealthCheckOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeHealthCheckHttpsHealthCheckOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleComputeHealthCheckHttpsHealthCheckOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleComputeHealthCheckHttpsHealthCheckOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleComputeHealthCheckHttpsHealthCheckOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeHealthCheckHttpsHealthCheckOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-googlebeta.googleComputeHealthCheck.GoogleComputeHealthCheckHttpsHealthCheckOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleComputeHealthCheckHttpsHealthCheckOutputReference_Override(g GoogleComputeHealthCheckHttpsHealthCheckOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-googlebeta.googleComputeHealthCheck.GoogleComputeHealthCheckHttpsHealthCheckOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeHealthCheckHttpsHealthCheckOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeHealthCheckHttpsHealthCheckOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeHealthCheckHttpsHealthCheckOutputReference)SetHost(val *string) {
	if err := j.validateSetHostParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"host",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeHealthCheckHttpsHealthCheckOutputReference)SetInternalValue(val *GoogleComputeHealthCheckHttpsHealthCheck) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeHealthCheckHttpsHealthCheckOutputReference)SetPort(val *float64) {
	if err := j.validateSetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeHealthCheckHttpsHealthCheckOutputReference)SetPortName(val *string) {
	if err := j.validateSetPortNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"portName",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeHealthCheckHttpsHealthCheckOutputReference)SetPortSpecification(val *string) {
	if err := j.validateSetPortSpecificationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"portSpecification",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeHealthCheckHttpsHealthCheckOutputReference)SetProxyHeader(val *string) {
	if err := j.validateSetProxyHeaderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"proxyHeader",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeHealthCheckHttpsHealthCheckOutputReference)SetRequestPath(val *string) {
	if err := j.validateSetRequestPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestPath",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeHealthCheckHttpsHealthCheckOutputReference)SetResponse(val *string) {
	if err := j.validateSetResponseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"response",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeHealthCheckHttpsHealthCheckOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeHealthCheckHttpsHealthCheckOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleComputeHealthCheckHttpsHealthCheckOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeHealthCheckHttpsHealthCheckOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleComputeHealthCheckHttpsHealthCheckOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeHealthCheckHttpsHealthCheckOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComputeHealthCheckHttpsHealthCheckOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComputeHealthCheckHttpsHealthCheckOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComputeHealthCheckHttpsHealthCheckOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComputeHealthCheckHttpsHealthCheckOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComputeHealthCheckHttpsHealthCheckOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComputeHealthCheckHttpsHealthCheckOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComputeHealthCheckHttpsHealthCheckOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeHealthCheckHttpsHealthCheckOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeHealthCheckHttpsHealthCheckOutputReference) ResetHost() {
	_jsii_.InvokeVoid(
		g,
		"resetHost",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeHealthCheckHttpsHealthCheckOutputReference) ResetPort() {
	_jsii_.InvokeVoid(
		g,
		"resetPort",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeHealthCheckHttpsHealthCheckOutputReference) ResetPortName() {
	_jsii_.InvokeVoid(
		g,
		"resetPortName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeHealthCheckHttpsHealthCheckOutputReference) ResetPortSpecification() {
	_jsii_.InvokeVoid(
		g,
		"resetPortSpecification",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeHealthCheckHttpsHealthCheckOutputReference) ResetProxyHeader() {
	_jsii_.InvokeVoid(
		g,
		"resetProxyHeader",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeHealthCheckHttpsHealthCheckOutputReference) ResetRequestPath() {
	_jsii_.InvokeVoid(
		g,
		"resetRequestPath",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeHealthCheckHttpsHealthCheckOutputReference) ResetResponse() {
	_jsii_.InvokeVoid(
		g,
		"resetResponse",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeHealthCheckHttpsHealthCheckOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleComputeHealthCheckHttpsHealthCheckOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

