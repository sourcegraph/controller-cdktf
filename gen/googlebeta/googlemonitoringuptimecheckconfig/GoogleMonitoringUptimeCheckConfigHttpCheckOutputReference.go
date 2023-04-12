package googlemonitoringuptimecheckconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/googlebeta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/googlebeta/googlemonitoringuptimecheckconfig/internal"
)

type GoogleMonitoringUptimeCheckConfigHttpCheckOutputReference interface {
	cdktf.ComplexObject
	AcceptedResponseStatusCodes() GoogleMonitoringUptimeCheckConfigHttpCheckAcceptedResponseStatusCodesList
	AcceptedResponseStatusCodesInput() interface{}
	AuthInfo() GoogleMonitoringUptimeCheckConfigHttpCheckAuthInfoOutputReference
	AuthInfoInput() *GoogleMonitoringUptimeCheckConfigHttpCheckAuthInfo
	Body() *string
	SetBody(val *string)
	BodyInput() *string
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
	ContentType() *string
	SetContentType(val *string)
	ContentTypeInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	Headers() *map[string]*string
	SetHeaders(val *map[string]*string)
	HeadersInput() *map[string]*string
	InternalValue() *GoogleMonitoringUptimeCheckConfigHttpCheck
	SetInternalValue(val *GoogleMonitoringUptimeCheckConfigHttpCheck)
	MaskHeaders() interface{}
	SetMaskHeaders(val interface{})
	MaskHeadersInput() interface{}
	Path() *string
	SetPath(val *string)
	PathInput() *string
	Port() *float64
	SetPort(val *float64)
	PortInput() *float64
	RequestMethod() *string
	SetRequestMethod(val *string)
	RequestMethodInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UseSsl() interface{}
	SetUseSsl(val interface{})
	UseSslInput() interface{}
	ValidateSsl() interface{}
	SetValidateSsl(val interface{})
	ValidateSslInput() interface{}
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
	PutAcceptedResponseStatusCodes(value interface{})
	PutAuthInfo(value *GoogleMonitoringUptimeCheckConfigHttpCheckAuthInfo)
	ResetAcceptedResponseStatusCodes()
	ResetAuthInfo()
	ResetBody()
	ResetContentType()
	ResetHeaders()
	ResetMaskHeaders()
	ResetPath()
	ResetPort()
	ResetRequestMethod()
	ResetUseSsl()
	ResetValidateSsl()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleMonitoringUptimeCheckConfigHttpCheckOutputReference
type jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckOutputReference) AcceptedResponseStatusCodes() GoogleMonitoringUptimeCheckConfigHttpCheckAcceptedResponseStatusCodesList {
	var returns GoogleMonitoringUptimeCheckConfigHttpCheckAcceptedResponseStatusCodesList
	_jsii_.Get(
		j,
		"acceptedResponseStatusCodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckOutputReference) AcceptedResponseStatusCodesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"acceptedResponseStatusCodesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckOutputReference) AuthInfo() GoogleMonitoringUptimeCheckConfigHttpCheckAuthInfoOutputReference {
	var returns GoogleMonitoringUptimeCheckConfigHttpCheckAuthInfoOutputReference
	_jsii_.Get(
		j,
		"authInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckOutputReference) AuthInfoInput() *GoogleMonitoringUptimeCheckConfigHttpCheckAuthInfo {
	var returns *GoogleMonitoringUptimeCheckConfigHttpCheckAuthInfo
	_jsii_.Get(
		j,
		"authInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckOutputReference) Body() *string {
	var returns *string
	_jsii_.Get(
		j,
		"body",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckOutputReference) BodyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bodyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckOutputReference) ContentType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckOutputReference) ContentTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckOutputReference) Headers() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"headers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckOutputReference) HeadersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"headersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckOutputReference) InternalValue() *GoogleMonitoringUptimeCheckConfigHttpCheck {
	var returns *GoogleMonitoringUptimeCheckConfigHttpCheck
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckOutputReference) MaskHeaders() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"maskHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckOutputReference) MaskHeadersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"maskHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckOutputReference) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckOutputReference) PathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckOutputReference) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckOutputReference) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckOutputReference) RequestMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckOutputReference) RequestMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckOutputReference) UseSsl() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useSsl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckOutputReference) UseSslInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useSslInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckOutputReference) ValidateSsl() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"validateSsl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckOutputReference) ValidateSslInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"validateSslInput",
		&returns,
	)
	return returns
}


func NewGoogleMonitoringUptimeCheckConfigHttpCheckOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleMonitoringUptimeCheckConfigHttpCheckOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleMonitoringUptimeCheckConfigHttpCheckOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-googlebeta.googleMonitoringUptimeCheckConfig.GoogleMonitoringUptimeCheckConfigHttpCheckOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleMonitoringUptimeCheckConfigHttpCheckOutputReference_Override(g GoogleMonitoringUptimeCheckConfigHttpCheckOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-googlebeta.googleMonitoringUptimeCheckConfig.GoogleMonitoringUptimeCheckConfigHttpCheckOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckOutputReference)SetBody(val *string) {
	if err := j.validateSetBodyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"body",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckOutputReference)SetContentType(val *string) {
	if err := j.validateSetContentTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contentType",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckOutputReference)SetHeaders(val *map[string]*string) {
	if err := j.validateSetHeadersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"headers",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckOutputReference)SetInternalValue(val *GoogleMonitoringUptimeCheckConfigHttpCheck) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckOutputReference)SetMaskHeaders(val interface{}) {
	if err := j.validateSetMaskHeadersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maskHeaders",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckOutputReference)SetPath(val *string) {
	if err := j.validateSetPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"path",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckOutputReference)SetPort(val *float64) {
	if err := j.validateSetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckOutputReference)SetRequestMethod(val *string) {
	if err := j.validateSetRequestMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestMethod",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckOutputReference)SetUseSsl(val interface{}) {
	if err := j.validateSetUseSslParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useSsl",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckOutputReference)SetValidateSsl(val interface{}) {
	if err := j.validateSetValidateSslParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"validateSsl",
		val,
	)
}

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckOutputReference) PutAcceptedResponseStatusCodes(value interface{}) {
	if err := g.validatePutAcceptedResponseStatusCodesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAcceptedResponseStatusCodes",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckOutputReference) PutAuthInfo(value *GoogleMonitoringUptimeCheckConfigHttpCheckAuthInfo) {
	if err := g.validatePutAuthInfoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAuthInfo",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckOutputReference) ResetAcceptedResponseStatusCodes() {
	_jsii_.InvokeVoid(
		g,
		"resetAcceptedResponseStatusCodes",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckOutputReference) ResetAuthInfo() {
	_jsii_.InvokeVoid(
		g,
		"resetAuthInfo",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckOutputReference) ResetBody() {
	_jsii_.InvokeVoid(
		g,
		"resetBody",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckOutputReference) ResetContentType() {
	_jsii_.InvokeVoid(
		g,
		"resetContentType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckOutputReference) ResetHeaders() {
	_jsii_.InvokeVoid(
		g,
		"resetHeaders",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckOutputReference) ResetMaskHeaders() {
	_jsii_.InvokeVoid(
		g,
		"resetMaskHeaders",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckOutputReference) ResetPath() {
	_jsii_.InvokeVoid(
		g,
		"resetPath",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckOutputReference) ResetPort() {
	_jsii_.InvokeVoid(
		g,
		"resetPort",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckOutputReference) ResetRequestMethod() {
	_jsii_.InvokeVoid(
		g,
		"resetRequestMethod",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckOutputReference) ResetUseSsl() {
	_jsii_.InvokeVoid(
		g,
		"resetUseSsl",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckOutputReference) ResetValidateSsl() {
	_jsii_.InvokeVoid(
		g,
		"resetValidateSsl",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfigHttpCheckOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

