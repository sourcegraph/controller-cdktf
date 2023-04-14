package listitem

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/cloudflare/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/cloudflare/listitem/internal"
)

type ListItemRedirectOutputReference interface {
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
	IncludeSubdomains() *string
	SetIncludeSubdomains(val *string)
	IncludeSubdomainsInput() *string
	InternalValue() *ListItemRedirect
	SetInternalValue(val *ListItemRedirect)
	PreservePathSuffix() *string
	SetPreservePathSuffix(val *string)
	PreservePathSuffixInput() *string
	PreserveQueryString() *string
	SetPreserveQueryString(val *string)
	PreserveQueryStringInput() *string
	SourceUrl() *string
	SetSourceUrl(val *string)
	SourceUrlInput() *string
	StatusCode() *float64
	SetStatusCode(val *float64)
	StatusCodeInput() *float64
	SubpathMatching() *string
	SetSubpathMatching(val *string)
	SubpathMatchingInput() *string
	TargetUrl() *string
	SetTargetUrl(val *string)
	TargetUrlInput() *string
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
	ResetIncludeSubdomains()
	ResetPreservePathSuffix()
	ResetPreserveQueryString()
	ResetStatusCode()
	ResetSubpathMatching()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ListItemRedirectOutputReference
type jsiiProxy_ListItemRedirectOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ListItemRedirectOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ListItemRedirectOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ListItemRedirectOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ListItemRedirectOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ListItemRedirectOutputReference) IncludeSubdomains() *string {
	var returns *string
	_jsii_.Get(
		j,
		"includeSubdomains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ListItemRedirectOutputReference) IncludeSubdomainsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"includeSubdomainsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ListItemRedirectOutputReference) InternalValue() *ListItemRedirect {
	var returns *ListItemRedirect
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ListItemRedirectOutputReference) PreservePathSuffix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preservePathSuffix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ListItemRedirectOutputReference) PreservePathSuffixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preservePathSuffixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ListItemRedirectOutputReference) PreserveQueryString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preserveQueryString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ListItemRedirectOutputReference) PreserveQueryStringInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preserveQueryStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ListItemRedirectOutputReference) SourceUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ListItemRedirectOutputReference) SourceUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ListItemRedirectOutputReference) StatusCode() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"statusCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ListItemRedirectOutputReference) StatusCodeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"statusCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ListItemRedirectOutputReference) SubpathMatching() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subpathMatching",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ListItemRedirectOutputReference) SubpathMatchingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subpathMatchingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ListItemRedirectOutputReference) TargetUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ListItemRedirectOutputReference) TargetUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ListItemRedirectOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ListItemRedirectOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewListItemRedirectOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ListItemRedirectOutputReference {
	_init_.Initialize()

	if err := validateNewListItemRedirectOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ListItemRedirectOutputReference{}

	_jsii_.Create(
		"cloudflare.listItem.ListItemRedirectOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewListItemRedirectOutputReference_Override(l ListItemRedirectOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"cloudflare.listItem.ListItemRedirectOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		l,
	)
}

func (j *jsiiProxy_ListItemRedirectOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ListItemRedirectOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ListItemRedirectOutputReference)SetIncludeSubdomains(val *string) {
	if err := j.validateSetIncludeSubdomainsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeSubdomains",
		val,
	)
}

func (j *jsiiProxy_ListItemRedirectOutputReference)SetInternalValue(val *ListItemRedirect) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ListItemRedirectOutputReference)SetPreservePathSuffix(val *string) {
	if err := j.validateSetPreservePathSuffixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preservePathSuffix",
		val,
	)
}

func (j *jsiiProxy_ListItemRedirectOutputReference)SetPreserveQueryString(val *string) {
	if err := j.validateSetPreserveQueryStringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preserveQueryString",
		val,
	)
}

func (j *jsiiProxy_ListItemRedirectOutputReference)SetSourceUrl(val *string) {
	if err := j.validateSetSourceUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceUrl",
		val,
	)
}

func (j *jsiiProxy_ListItemRedirectOutputReference)SetStatusCode(val *float64) {
	if err := j.validateSetStatusCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"statusCode",
		val,
	)
}

func (j *jsiiProxy_ListItemRedirectOutputReference)SetSubpathMatching(val *string) {
	if err := j.validateSetSubpathMatchingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subpathMatching",
		val,
	)
}

func (j *jsiiProxy_ListItemRedirectOutputReference)SetTargetUrl(val *string) {
	if err := j.validateSetTargetUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetUrl",
		val,
	)
}

func (j *jsiiProxy_ListItemRedirectOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ListItemRedirectOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (l *jsiiProxy_ListItemRedirectOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_ListItemRedirectOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := l.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_ListItemRedirectOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_ListItemRedirectOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := l.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_ListItemRedirectOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := l.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_ListItemRedirectOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := l.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_ListItemRedirectOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := l.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_ListItemRedirectOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := l.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_ListItemRedirectOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := l.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_ListItemRedirectOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := l.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_ListItemRedirectOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_ListItemRedirectOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_ListItemRedirectOutputReference) ResetIncludeSubdomains() {
	_jsii_.InvokeVoid(
		l,
		"resetIncludeSubdomains",
		nil, // no parameters
	)
}

func (l *jsiiProxy_ListItemRedirectOutputReference) ResetPreservePathSuffix() {
	_jsii_.InvokeVoid(
		l,
		"resetPreservePathSuffix",
		nil, // no parameters
	)
}

func (l *jsiiProxy_ListItemRedirectOutputReference) ResetPreserveQueryString() {
	_jsii_.InvokeVoid(
		l,
		"resetPreserveQueryString",
		nil, // no parameters
	)
}

func (l *jsiiProxy_ListItemRedirectOutputReference) ResetStatusCode() {
	_jsii_.InvokeVoid(
		l,
		"resetStatusCode",
		nil, // no parameters
	)
}

func (l *jsiiProxy_ListItemRedirectOutputReference) ResetSubpathMatching() {
	_jsii_.InvokeVoid(
		l,
		"resetSubpathMatching",
		nil, // no parameters
	)
}

func (l *jsiiProxy_ListItemRedirectOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := l.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		l,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_ListItemRedirectOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

