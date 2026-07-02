package lblistenerrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/aws/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/aws/lblistenerrule/internal"
)

type LbListenerRuleConditionOutputReference interface {
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
	HostHeader() LbListenerRuleConditionHostHeaderOutputReference
	HostHeaderInput() *LbListenerRuleConditionHostHeader
	HttpHeader() LbListenerRuleConditionHttpHeaderOutputReference
	HttpHeaderInput() *LbListenerRuleConditionHttpHeader
	HttpRequestMethod() LbListenerRuleConditionHttpRequestMethodOutputReference
	HttpRequestMethodInput() *LbListenerRuleConditionHttpRequestMethod
	InternalValue() any
	SetInternalValue(val any)
	PathPattern() LbListenerRuleConditionPathPatternOutputReference
	PathPatternInput() *LbListenerRuleConditionPathPattern
	QueryString() LbListenerRuleConditionQueryStringList
	QueryStringInput() any
	SourceIp() LbListenerRuleConditionSourceIpOutputReference
	SourceIpInput() *LbListenerRuleConditionSourceIp
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
	PutHostHeader(value *LbListenerRuleConditionHostHeader)
	PutHttpHeader(value *LbListenerRuleConditionHttpHeader)
	PutHttpRequestMethod(value *LbListenerRuleConditionHttpRequestMethod)
	PutPathPattern(value *LbListenerRuleConditionPathPattern)
	PutQueryString(value any)
	PutSourceIp(value *LbListenerRuleConditionSourceIp)
	ResetHostHeader()
	ResetHttpHeader()
	ResetHttpRequestMethod()
	ResetPathPattern()
	ResetQueryString()
	ResetSourceIp()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LbListenerRuleConditionOutputReference
type jsiiProxy_LbListenerRuleConditionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LbListenerRuleConditionOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerRuleConditionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerRuleConditionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerRuleConditionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerRuleConditionOutputReference) HostHeader() LbListenerRuleConditionHostHeaderOutputReference {
	var returns LbListenerRuleConditionHostHeaderOutputReference
	_jsii_.Get(
		j,
		"hostHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerRuleConditionOutputReference) HostHeaderInput() *LbListenerRuleConditionHostHeader {
	var returns *LbListenerRuleConditionHostHeader
	_jsii_.Get(
		j,
		"hostHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerRuleConditionOutputReference) HttpHeader() LbListenerRuleConditionHttpHeaderOutputReference {
	var returns LbListenerRuleConditionHttpHeaderOutputReference
	_jsii_.Get(
		j,
		"httpHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerRuleConditionOutputReference) HttpHeaderInput() *LbListenerRuleConditionHttpHeader {
	var returns *LbListenerRuleConditionHttpHeader
	_jsii_.Get(
		j,
		"httpHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerRuleConditionOutputReference) HttpRequestMethod() LbListenerRuleConditionHttpRequestMethodOutputReference {
	var returns LbListenerRuleConditionHttpRequestMethodOutputReference
	_jsii_.Get(
		j,
		"httpRequestMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerRuleConditionOutputReference) HttpRequestMethodInput() *LbListenerRuleConditionHttpRequestMethod {
	var returns *LbListenerRuleConditionHttpRequestMethod
	_jsii_.Get(
		j,
		"httpRequestMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerRuleConditionOutputReference) InternalValue() any {
	var returns any
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerRuleConditionOutputReference) PathPattern() LbListenerRuleConditionPathPatternOutputReference {
	var returns LbListenerRuleConditionPathPatternOutputReference
	_jsii_.Get(
		j,
		"pathPattern",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerRuleConditionOutputReference) PathPatternInput() *LbListenerRuleConditionPathPattern {
	var returns *LbListenerRuleConditionPathPattern
	_jsii_.Get(
		j,
		"pathPatternInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerRuleConditionOutputReference) QueryString() LbListenerRuleConditionQueryStringList {
	var returns LbListenerRuleConditionQueryStringList
	_jsii_.Get(
		j,
		"queryString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerRuleConditionOutputReference) QueryStringInput() any {
	var returns any
	_jsii_.Get(
		j,
		"queryStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerRuleConditionOutputReference) SourceIp() LbListenerRuleConditionSourceIpOutputReference {
	var returns LbListenerRuleConditionSourceIpOutputReference
	_jsii_.Get(
		j,
		"sourceIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerRuleConditionOutputReference) SourceIpInput() *LbListenerRuleConditionSourceIp {
	var returns *LbListenerRuleConditionSourceIp
	_jsii_.Get(
		j,
		"sourceIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerRuleConditionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerRuleConditionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewLbListenerRuleConditionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) LbListenerRuleConditionOutputReference {
	_init_.Initialize()

	if err := validateNewLbListenerRuleConditionOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_LbListenerRuleConditionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.lbListenerRule.LbListenerRuleConditionOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewLbListenerRuleConditionOutputReference_Override(l LbListenerRuleConditionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.lbListenerRule.LbListenerRuleConditionOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		l,
	)
}

func (j *jsiiProxy_LbListenerRuleConditionOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LbListenerRuleConditionOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LbListenerRuleConditionOutputReference) SetInternalValue(val any) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LbListenerRuleConditionOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LbListenerRuleConditionOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (l *jsiiProxy_LbListenerRuleConditionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbListenerRuleConditionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := l.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbListenerRuleConditionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbListenerRuleConditionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := l.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbListenerRuleConditionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := l.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbListenerRuleConditionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := l.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbListenerRuleConditionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := l.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbListenerRuleConditionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := l.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbListenerRuleConditionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := l.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbListenerRuleConditionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := l.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbListenerRuleConditionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbListenerRuleConditionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbListenerRuleConditionOutputReference) PutHostHeader(value *LbListenerRuleConditionHostHeader) {
	if err := l.validatePutHostHeaderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putHostHeader",
		[]any{value},
	)
}

func (l *jsiiProxy_LbListenerRuleConditionOutputReference) PutHttpHeader(value *LbListenerRuleConditionHttpHeader) {
	if err := l.validatePutHttpHeaderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putHttpHeader",
		[]any{value},
	)
}

func (l *jsiiProxy_LbListenerRuleConditionOutputReference) PutHttpRequestMethod(value *LbListenerRuleConditionHttpRequestMethod) {
	if err := l.validatePutHttpRequestMethodParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putHttpRequestMethod",
		[]any{value},
	)
}

func (l *jsiiProxy_LbListenerRuleConditionOutputReference) PutPathPattern(value *LbListenerRuleConditionPathPattern) {
	if err := l.validatePutPathPatternParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putPathPattern",
		[]any{value},
	)
}

func (l *jsiiProxy_LbListenerRuleConditionOutputReference) PutQueryString(value any) {
	if err := l.validatePutQueryStringParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putQueryString",
		[]any{value},
	)
}

func (l *jsiiProxy_LbListenerRuleConditionOutputReference) PutSourceIp(value *LbListenerRuleConditionSourceIp) {
	if err := l.validatePutSourceIpParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putSourceIp",
		[]any{value},
	)
}

func (l *jsiiProxy_LbListenerRuleConditionOutputReference) ResetHostHeader() {
	_jsii_.InvokeVoid(
		l,
		"resetHostHeader",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbListenerRuleConditionOutputReference) ResetHttpHeader() {
	_jsii_.InvokeVoid(
		l,
		"resetHttpHeader",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbListenerRuleConditionOutputReference) ResetHttpRequestMethod() {
	_jsii_.InvokeVoid(
		l,
		"resetHttpRequestMethod",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbListenerRuleConditionOutputReference) ResetPathPattern() {
	_jsii_.InvokeVoid(
		l,
		"resetPathPattern",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbListenerRuleConditionOutputReference) ResetQueryString() {
	_jsii_.InvokeVoid(
		l,
		"resetQueryString",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbListenerRuleConditionOutputReference) ResetSourceIp() {
	_jsii_.InvokeVoid(
		l,
		"resetSourceIp",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbListenerRuleConditionOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := l.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		l,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbListenerRuleConditionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
