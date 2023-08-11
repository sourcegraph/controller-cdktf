package pagerule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/cloudflare/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/cloudflare/pagerule/internal"
)

type PageRuleActionsCacheKeyFieldsQueryStringOutputReference interface {
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
	Exclude() *[]*string
	SetExclude(val *[]*string)
	ExcludeInput() *[]*string
	// Experimental.
	Fqn() *string
	Ignore() interface{}
	SetIgnore(val interface{})
	IgnoreInput() interface{}
	Include() *[]*string
	SetInclude(val *[]*string)
	IncludeInput() *[]*string
	InternalValue() *PageRuleActionsCacheKeyFieldsQueryString
	SetInternalValue(val *PageRuleActionsCacheKeyFieldsQueryString)
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
	ResetExclude()
	ResetIgnore()
	ResetInclude()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PageRuleActionsCacheKeyFieldsQueryStringOutputReference
type jsiiProxy_PageRuleActionsCacheKeyFieldsQueryStringOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PageRuleActionsCacheKeyFieldsQueryStringOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsCacheKeyFieldsQueryStringOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsCacheKeyFieldsQueryStringOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsCacheKeyFieldsQueryStringOutputReference) Exclude() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"exclude",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsCacheKeyFieldsQueryStringOutputReference) ExcludeInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsCacheKeyFieldsQueryStringOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsCacheKeyFieldsQueryStringOutputReference) Ignore() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsCacheKeyFieldsQueryStringOutputReference) IgnoreInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsCacheKeyFieldsQueryStringOutputReference) Include() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"include",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsCacheKeyFieldsQueryStringOutputReference) IncludeInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsCacheKeyFieldsQueryStringOutputReference) InternalValue() *PageRuleActionsCacheKeyFieldsQueryString {
	var returns *PageRuleActionsCacheKeyFieldsQueryString
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsCacheKeyFieldsQueryStringOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PageRuleActionsCacheKeyFieldsQueryStringOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPageRuleActionsCacheKeyFieldsQueryStringOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PageRuleActionsCacheKeyFieldsQueryStringOutputReference {
	_init_.Initialize()

	if err := validateNewPageRuleActionsCacheKeyFieldsQueryStringOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PageRuleActionsCacheKeyFieldsQueryStringOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.pageRule.PageRuleActionsCacheKeyFieldsQueryStringOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPageRuleActionsCacheKeyFieldsQueryStringOutputReference_Override(p PageRuleActionsCacheKeyFieldsQueryStringOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.pageRule.PageRuleActionsCacheKeyFieldsQueryStringOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PageRuleActionsCacheKeyFieldsQueryStringOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PageRuleActionsCacheKeyFieldsQueryStringOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PageRuleActionsCacheKeyFieldsQueryStringOutputReference)SetExclude(val *[]*string) {
	if err := j.validateSetExcludeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exclude",
		val,
	)
}

func (j *jsiiProxy_PageRuleActionsCacheKeyFieldsQueryStringOutputReference)SetIgnore(val interface{}) {
	if err := j.validateSetIgnoreParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignore",
		val,
	)
}

func (j *jsiiProxy_PageRuleActionsCacheKeyFieldsQueryStringOutputReference)SetInclude(val *[]*string) {
	if err := j.validateSetIncludeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"include",
		val,
	)
}

func (j *jsiiProxy_PageRuleActionsCacheKeyFieldsQueryStringOutputReference)SetInternalValue(val *PageRuleActionsCacheKeyFieldsQueryString) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PageRuleActionsCacheKeyFieldsQueryStringOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PageRuleActionsCacheKeyFieldsQueryStringOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PageRuleActionsCacheKeyFieldsQueryStringOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PageRuleActionsCacheKeyFieldsQueryStringOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PageRuleActionsCacheKeyFieldsQueryStringOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PageRuleActionsCacheKeyFieldsQueryStringOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PageRuleActionsCacheKeyFieldsQueryStringOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PageRuleActionsCacheKeyFieldsQueryStringOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PageRuleActionsCacheKeyFieldsQueryStringOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PageRuleActionsCacheKeyFieldsQueryStringOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PageRuleActionsCacheKeyFieldsQueryStringOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PageRuleActionsCacheKeyFieldsQueryStringOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PageRuleActionsCacheKeyFieldsQueryStringOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PageRuleActionsCacheKeyFieldsQueryStringOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PageRuleActionsCacheKeyFieldsQueryStringOutputReference) ResetExclude() {
	_jsii_.InvokeVoid(
		p,
		"resetExclude",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PageRuleActionsCacheKeyFieldsQueryStringOutputReference) ResetIgnore() {
	_jsii_.InvokeVoid(
		p,
		"resetIgnore",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PageRuleActionsCacheKeyFieldsQueryStringOutputReference) ResetInclude() {
	_jsii_.InvokeVoid(
		p,
		"resetInclude",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PageRuleActionsCacheKeyFieldsQueryStringOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (p *jsiiProxy_PageRuleActionsCacheKeyFieldsQueryStringOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

