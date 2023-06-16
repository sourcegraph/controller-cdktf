package pagerule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/cloudflare/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/cloudflare/pagerule/internal"
)

type PageRuleActionsCacheKeyFieldsOutputReference interface {
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
	Cookie() PageRuleActionsCacheKeyFieldsCookieOutputReference
	CookieInput() *PageRuleActionsCacheKeyFieldsCookie
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	Header() PageRuleActionsCacheKeyFieldsHeaderOutputReference
	HeaderInput() *PageRuleActionsCacheKeyFieldsHeader
	Host() PageRuleActionsCacheKeyFieldsHostOutputReference
	HostInput() *PageRuleActionsCacheKeyFieldsHost
	InternalValue() *PageRuleActionsCacheKeyFields
	SetInternalValue(val *PageRuleActionsCacheKeyFields)
	QueryString() PageRuleActionsCacheKeyFieldsQueryStringOutputReference
	QueryStringInput() *PageRuleActionsCacheKeyFieldsQueryString
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	User() PageRuleActionsCacheKeyFieldsUserOutputReference
	UserInput() *PageRuleActionsCacheKeyFieldsUser
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
	PutCookie(value *PageRuleActionsCacheKeyFieldsCookie)
	PutHeader(value *PageRuleActionsCacheKeyFieldsHeader)
	PutHost(value *PageRuleActionsCacheKeyFieldsHost)
	PutQueryString(value *PageRuleActionsCacheKeyFieldsQueryString)
	PutUser(value *PageRuleActionsCacheKeyFieldsUser)
	ResetCookie()
	ResetHeader()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PageRuleActionsCacheKeyFieldsOutputReference
type jsiiProxy_PageRuleActionsCacheKeyFieldsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PageRuleActionsCacheKeyFieldsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsCacheKeyFieldsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsCacheKeyFieldsOutputReference) Cookie() PageRuleActionsCacheKeyFieldsCookieOutputReference {
	var returns PageRuleActionsCacheKeyFieldsCookieOutputReference
	_jsii_.Get(
		j,
		"cookie",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsCacheKeyFieldsOutputReference) CookieInput() *PageRuleActionsCacheKeyFieldsCookie {
	var returns *PageRuleActionsCacheKeyFieldsCookie
	_jsii_.Get(
		j,
		"cookieInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsCacheKeyFieldsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsCacheKeyFieldsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsCacheKeyFieldsOutputReference) Header() PageRuleActionsCacheKeyFieldsHeaderOutputReference {
	var returns PageRuleActionsCacheKeyFieldsHeaderOutputReference
	_jsii_.Get(
		j,
		"header",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsCacheKeyFieldsOutputReference) HeaderInput() *PageRuleActionsCacheKeyFieldsHeader {
	var returns *PageRuleActionsCacheKeyFieldsHeader
	_jsii_.Get(
		j,
		"headerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsCacheKeyFieldsOutputReference) Host() PageRuleActionsCacheKeyFieldsHostOutputReference {
	var returns PageRuleActionsCacheKeyFieldsHostOutputReference
	_jsii_.Get(
		j,
		"host",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsCacheKeyFieldsOutputReference) HostInput() *PageRuleActionsCacheKeyFieldsHost {
	var returns *PageRuleActionsCacheKeyFieldsHost
	_jsii_.Get(
		j,
		"hostInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsCacheKeyFieldsOutputReference) InternalValue() *PageRuleActionsCacheKeyFields {
	var returns *PageRuleActionsCacheKeyFields
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsCacheKeyFieldsOutputReference) QueryString() PageRuleActionsCacheKeyFieldsQueryStringOutputReference {
	var returns PageRuleActionsCacheKeyFieldsQueryStringOutputReference
	_jsii_.Get(
		j,
		"queryString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsCacheKeyFieldsOutputReference) QueryStringInput() *PageRuleActionsCacheKeyFieldsQueryString {
	var returns *PageRuleActionsCacheKeyFieldsQueryString
	_jsii_.Get(
		j,
		"queryStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsCacheKeyFieldsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsCacheKeyFieldsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsCacheKeyFieldsOutputReference) User() PageRuleActionsCacheKeyFieldsUserOutputReference {
	var returns PageRuleActionsCacheKeyFieldsUserOutputReference
	_jsii_.Get(
		j,
		"user",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsCacheKeyFieldsOutputReference) UserInput() *PageRuleActionsCacheKeyFieldsUser {
	var returns *PageRuleActionsCacheKeyFieldsUser
	_jsii_.Get(
		j,
		"userInput",
		&returns,
	)
	return returns
}


func NewPageRuleActionsCacheKeyFieldsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PageRuleActionsCacheKeyFieldsOutputReference {
	_init_.Initialize()

	if err := validateNewPageRuleActionsCacheKeyFieldsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PageRuleActionsCacheKeyFieldsOutputReference{}

	_jsii_.Create(
		"cloudflare.pageRule.PageRuleActionsCacheKeyFieldsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPageRuleActionsCacheKeyFieldsOutputReference_Override(p PageRuleActionsCacheKeyFieldsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"cloudflare.pageRule.PageRuleActionsCacheKeyFieldsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PageRuleActionsCacheKeyFieldsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PageRuleActionsCacheKeyFieldsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PageRuleActionsCacheKeyFieldsOutputReference)SetInternalValue(val *PageRuleActionsCacheKeyFields) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PageRuleActionsCacheKeyFieldsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PageRuleActionsCacheKeyFieldsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PageRuleActionsCacheKeyFieldsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PageRuleActionsCacheKeyFieldsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PageRuleActionsCacheKeyFieldsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PageRuleActionsCacheKeyFieldsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PageRuleActionsCacheKeyFieldsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PageRuleActionsCacheKeyFieldsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PageRuleActionsCacheKeyFieldsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PageRuleActionsCacheKeyFieldsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PageRuleActionsCacheKeyFieldsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PageRuleActionsCacheKeyFieldsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PageRuleActionsCacheKeyFieldsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PageRuleActionsCacheKeyFieldsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PageRuleActionsCacheKeyFieldsOutputReference) PutCookie(value *PageRuleActionsCacheKeyFieldsCookie) {
	if err := p.validatePutCookieParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putCookie",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PageRuleActionsCacheKeyFieldsOutputReference) PutHeader(value *PageRuleActionsCacheKeyFieldsHeader) {
	if err := p.validatePutHeaderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putHeader",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PageRuleActionsCacheKeyFieldsOutputReference) PutHost(value *PageRuleActionsCacheKeyFieldsHost) {
	if err := p.validatePutHostParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putHost",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PageRuleActionsCacheKeyFieldsOutputReference) PutQueryString(value *PageRuleActionsCacheKeyFieldsQueryString) {
	if err := p.validatePutQueryStringParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putQueryString",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PageRuleActionsCacheKeyFieldsOutputReference) PutUser(value *PageRuleActionsCacheKeyFieldsUser) {
	if err := p.validatePutUserParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putUser",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PageRuleActionsCacheKeyFieldsOutputReference) ResetCookie() {
	_jsii_.InvokeVoid(
		p,
		"resetCookie",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PageRuleActionsCacheKeyFieldsOutputReference) ResetHeader() {
	_jsii_.InvokeVoid(
		p,
		"resetHeader",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PageRuleActionsCacheKeyFieldsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := p.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		p,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PageRuleActionsCacheKeyFieldsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

