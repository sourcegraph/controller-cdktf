package wafv2rulegroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/aws/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/aws/wafv2rulegroup/internal"
)

type Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference interface {
	cdktf.ComplexObject
	AllQueryArguments() Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchAllQueryArgumentsOutputReference
	AllQueryArgumentsInput() *Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchAllQueryArguments
	Body() Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchBodyOutputReference
	BodyInput() *Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchBody
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
	Cookies() Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookiesOutputReference
	CookiesInput() *Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookies
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	Headers() Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchHeadersList
	HeadersInput() interface{}
	InternalValue() *Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatch
	SetInternalValue(val *Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatch)
	JsonBody() Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchJsonBodyOutputReference
	JsonBodyInput() *Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchJsonBody
	Method() Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchMethodOutputReference
	MethodInput() *Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchMethod
	QueryString() Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchQueryStringOutputReference
	QueryStringInput() *Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchQueryString
	SingleHeader() Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchSingleHeaderOutputReference
	SingleHeaderInput() *Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchSingleHeader
	SingleQueryArgument() Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchSingleQueryArgumentOutputReference
	SingleQueryArgumentInput() *Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchSingleQueryArgument
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UriPath() Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchUriPathOutputReference
	UriPathInput() *Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchUriPath
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
	PutAllQueryArguments(value *Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchAllQueryArguments)
	PutBody(value *Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchBody)
	PutCookies(value *Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookies)
	PutHeaders(value interface{})
	PutJsonBody(value *Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchJsonBody)
	PutMethod(value *Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchMethod)
	PutQueryString(value *Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchQueryString)
	PutSingleHeader(value *Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchSingleHeader)
	PutSingleQueryArgument(value *Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchSingleQueryArgument)
	PutUriPath(value *Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchUriPath)
	ResetAllQueryArguments()
	ResetBody()
	ResetCookies()
	ResetHeaders()
	ResetJsonBody()
	ResetMethod()
	ResetQueryString()
	ResetSingleHeader()
	ResetSingleQueryArgument()
	ResetUriPath()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference
type jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) AllQueryArguments() Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchAllQueryArgumentsOutputReference {
	var returns Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchAllQueryArgumentsOutputReference
	_jsii_.Get(
		j,
		"allQueryArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) AllQueryArgumentsInput() *Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchAllQueryArguments {
	var returns *Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchAllQueryArguments
	_jsii_.Get(
		j,
		"allQueryArgumentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) Body() Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchBodyOutputReference {
	var returns Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchBodyOutputReference
	_jsii_.Get(
		j,
		"body",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) BodyInput() *Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchBody {
	var returns *Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchBody
	_jsii_.Get(
		j,
		"bodyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) Cookies() Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookiesOutputReference {
	var returns Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookiesOutputReference
	_jsii_.Get(
		j,
		"cookies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) CookiesInput() *Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookies {
	var returns *Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookies
	_jsii_.Get(
		j,
		"cookiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) Headers() Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchHeadersList {
	var returns Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchHeadersList
	_jsii_.Get(
		j,
		"headers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) HeadersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"headersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) InternalValue() *Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatch {
	var returns *Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatch
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) JsonBody() Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchJsonBodyOutputReference {
	var returns Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchJsonBodyOutputReference
	_jsii_.Get(
		j,
		"jsonBody",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) JsonBodyInput() *Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchJsonBody {
	var returns *Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchJsonBody
	_jsii_.Get(
		j,
		"jsonBodyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) Method() Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchMethodOutputReference {
	var returns Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchMethodOutputReference
	_jsii_.Get(
		j,
		"method",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) MethodInput() *Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchMethod {
	var returns *Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchMethod
	_jsii_.Get(
		j,
		"methodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) QueryString() Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchQueryStringOutputReference {
	var returns Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchQueryStringOutputReference
	_jsii_.Get(
		j,
		"queryString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) QueryStringInput() *Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchQueryString {
	var returns *Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchQueryString
	_jsii_.Get(
		j,
		"queryStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) SingleHeader() Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchSingleHeaderOutputReference {
	var returns Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchSingleHeaderOutputReference
	_jsii_.Get(
		j,
		"singleHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) SingleHeaderInput() *Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchSingleHeader {
	var returns *Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchSingleHeader
	_jsii_.Get(
		j,
		"singleHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) SingleQueryArgument() Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchSingleQueryArgumentOutputReference {
	var returns Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchSingleQueryArgumentOutputReference
	_jsii_.Get(
		j,
		"singleQueryArgument",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) SingleQueryArgumentInput() *Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchSingleQueryArgument {
	var returns *Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchSingleQueryArgument
	_jsii_.Get(
		j,
		"singleQueryArgumentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) UriPath() Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchUriPathOutputReference {
	var returns Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchUriPathOutputReference
	_jsii_.Get(
		j,
		"uriPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) UriPathInput() *Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchUriPath {
	var returns *Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchUriPath
	_jsii_.Get(
		j,
		"uriPathInput",
		&returns,
	)
	return returns
}


func NewWafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference {
	_init_.Initialize()

	if err := validateNewWafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.wafv2RuleGroup.Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference_Override(w Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.wafv2RuleGroup.Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference)SetInternalValue(val *Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatch) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := w.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := w.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := w.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := w.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := w.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := w.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := w.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := w.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := w.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := w.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) PutAllQueryArguments(value *Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchAllQueryArguments) {
	if err := w.validatePutAllQueryArgumentsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putAllQueryArguments",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) PutBody(value *Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchBody) {
	if err := w.validatePutBodyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putBody",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) PutCookies(value *Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookies) {
	if err := w.validatePutCookiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putCookies",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) PutHeaders(value interface{}) {
	if err := w.validatePutHeadersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putHeaders",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) PutJsonBody(value *Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchJsonBody) {
	if err := w.validatePutJsonBodyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putJsonBody",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) PutMethod(value *Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchMethod) {
	if err := w.validatePutMethodParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putMethod",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) PutQueryString(value *Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchQueryString) {
	if err := w.validatePutQueryStringParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putQueryString",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) PutSingleHeader(value *Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchSingleHeader) {
	if err := w.validatePutSingleHeaderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putSingleHeader",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) PutSingleQueryArgument(value *Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchSingleQueryArgument) {
	if err := w.validatePutSingleQueryArgumentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putSingleQueryArgument",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) PutUriPath(value *Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchUriPath) {
	if err := w.validatePutUriPathParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putUriPath",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) ResetAllQueryArguments() {
	_jsii_.InvokeVoid(
		w,
		"resetAllQueryArguments",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) ResetBody() {
	_jsii_.InvokeVoid(
		w,
		"resetBody",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) ResetCookies() {
	_jsii_.InvokeVoid(
		w,
		"resetCookies",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) ResetHeaders() {
	_jsii_.InvokeVoid(
		w,
		"resetHeaders",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) ResetJsonBody() {
	_jsii_.InvokeVoid(
		w,
		"resetJsonBody",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) ResetMethod() {
	_jsii_.InvokeVoid(
		w,
		"resetMethod",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) ResetQueryString() {
	_jsii_.InvokeVoid(
		w,
		"resetQueryString",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) ResetSingleHeader() {
	_jsii_.InvokeVoid(
		w,
		"resetSingleHeader",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) ResetSingleQueryArgument() {
	_jsii_.InvokeVoid(
		w,
		"resetSingleQueryArgument",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) ResetUriPath() {
	_jsii_.InvokeVoid(
		w,
		"resetUriPath",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := w.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

