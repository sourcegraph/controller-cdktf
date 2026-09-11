package googlecomputeregionurlmap

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecomputeregionurlmap/internal"
)

type GoogleComputeRegionUrlMapPathMatcherDefaultRouteActionCorsPolicyOutputReference interface {
	cdktn.ComplexObject
	AllowCredentials() interface{}
	SetAllowCredentials(val interface{})
	AllowCredentialsInput() interface{}
	AllowHeaders() *[]*string
	SetAllowHeaders(val *[]*string)
	AllowHeadersInput() *[]*string
	AllowMethods() *[]*string
	SetAllowMethods(val *[]*string)
	AllowMethodsInput() *[]*string
	AllowOriginRegexes() *[]*string
	SetAllowOriginRegexes(val *[]*string)
	AllowOriginRegexesInput() *[]*string
	AllowOrigins() *[]*string
	SetAllowOrigins(val *[]*string)
	AllowOriginsInput() *[]*string
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
	Disabled() interface{}
	SetDisabled(val interface{})
	DisabledInput() interface{}
	ExposeHeaders() *[]*string
	SetExposeHeaders(val *[]*string)
	ExposeHeadersInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleComputeRegionUrlMapPathMatcherDefaultRouteActionCorsPolicy
	SetInternalValue(val *GoogleComputeRegionUrlMapPathMatcherDefaultRouteActionCorsPolicy)
	MaxAge() *float64
	SetMaxAge(val *float64)
	MaxAgeInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
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
	ResetAllowCredentials()
	ResetAllowHeaders()
	ResetAllowMethods()
	ResetAllowOriginRegexes()
	ResetAllowOrigins()
	ResetDisabled()
	ResetExposeHeaders()
	ResetMaxAge()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleComputeRegionUrlMapPathMatcherDefaultRouteActionCorsPolicyOutputReference
type jsiiProxy_GoogleComputeRegionUrlMapPathMatcherDefaultRouteActionCorsPolicyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherDefaultRouteActionCorsPolicyOutputReference) AllowCredentials() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherDefaultRouteActionCorsPolicyOutputReference) AllowCredentialsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowCredentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherDefaultRouteActionCorsPolicyOutputReference) AllowHeaders() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherDefaultRouteActionCorsPolicyOutputReference) AllowHeadersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherDefaultRouteActionCorsPolicyOutputReference) AllowMethods() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowMethods",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherDefaultRouteActionCorsPolicyOutputReference) AllowMethodsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowMethodsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherDefaultRouteActionCorsPolicyOutputReference) AllowOriginRegexes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowOriginRegexes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherDefaultRouteActionCorsPolicyOutputReference) AllowOriginRegexesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowOriginRegexesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherDefaultRouteActionCorsPolicyOutputReference) AllowOrigins() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowOrigins",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherDefaultRouteActionCorsPolicyOutputReference) AllowOriginsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowOriginsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherDefaultRouteActionCorsPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherDefaultRouteActionCorsPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherDefaultRouteActionCorsPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherDefaultRouteActionCorsPolicyOutputReference) Disabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherDefaultRouteActionCorsPolicyOutputReference) DisabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherDefaultRouteActionCorsPolicyOutputReference) ExposeHeaders() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"exposeHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherDefaultRouteActionCorsPolicyOutputReference) ExposeHeadersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"exposeHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherDefaultRouteActionCorsPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherDefaultRouteActionCorsPolicyOutputReference) InternalValue() *GoogleComputeRegionUrlMapPathMatcherDefaultRouteActionCorsPolicy {
	var returns *GoogleComputeRegionUrlMapPathMatcherDefaultRouteActionCorsPolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherDefaultRouteActionCorsPolicyOutputReference) MaxAge() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherDefaultRouteActionCorsPolicyOutputReference) MaxAgeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAgeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherDefaultRouteActionCorsPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherDefaultRouteActionCorsPolicyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleComputeRegionUrlMapPathMatcherDefaultRouteActionCorsPolicyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleComputeRegionUrlMapPathMatcherDefaultRouteActionCorsPolicyOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleComputeRegionUrlMapPathMatcherDefaultRouteActionCorsPolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeRegionUrlMapPathMatcherDefaultRouteActionCorsPolicyOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleComputeRegionUrlMap.GoogleComputeRegionUrlMapPathMatcherDefaultRouteActionCorsPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleComputeRegionUrlMapPathMatcherDefaultRouteActionCorsPolicyOutputReference_Override(g GoogleComputeRegionUrlMapPathMatcherDefaultRouteActionCorsPolicyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleComputeRegionUrlMap.GoogleComputeRegionUrlMapPathMatcherDefaultRouteActionCorsPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherDefaultRouteActionCorsPolicyOutputReference)SetAllowCredentials(val interface{}) {
	if err := j.validateSetAllowCredentialsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowCredentials",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherDefaultRouteActionCorsPolicyOutputReference)SetAllowHeaders(val *[]*string) {
	if err := j.validateSetAllowHeadersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowHeaders",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherDefaultRouteActionCorsPolicyOutputReference)SetAllowMethods(val *[]*string) {
	if err := j.validateSetAllowMethodsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowMethods",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherDefaultRouteActionCorsPolicyOutputReference)SetAllowOriginRegexes(val *[]*string) {
	if err := j.validateSetAllowOriginRegexesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowOriginRegexes",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherDefaultRouteActionCorsPolicyOutputReference)SetAllowOrigins(val *[]*string) {
	if err := j.validateSetAllowOriginsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowOrigins",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherDefaultRouteActionCorsPolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherDefaultRouteActionCorsPolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherDefaultRouteActionCorsPolicyOutputReference)SetDisabled(val interface{}) {
	if err := j.validateSetDisabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disabled",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherDefaultRouteActionCorsPolicyOutputReference)SetExposeHeaders(val *[]*string) {
	if err := j.validateSetExposeHeadersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exposeHeaders",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherDefaultRouteActionCorsPolicyOutputReference)SetInternalValue(val *GoogleComputeRegionUrlMapPathMatcherDefaultRouteActionCorsPolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherDefaultRouteActionCorsPolicyOutputReference)SetMaxAge(val *float64) {
	if err := j.validateSetMaxAgeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxAge",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherDefaultRouteActionCorsPolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherDefaultRouteActionCorsPolicyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherDefaultRouteActionCorsPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherDefaultRouteActionCorsPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherDefaultRouteActionCorsPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherDefaultRouteActionCorsPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherDefaultRouteActionCorsPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherDefaultRouteActionCorsPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherDefaultRouteActionCorsPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherDefaultRouteActionCorsPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherDefaultRouteActionCorsPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherDefaultRouteActionCorsPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherDefaultRouteActionCorsPolicyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherDefaultRouteActionCorsPolicyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherDefaultRouteActionCorsPolicyOutputReference) ResetAllowCredentials() {
	_jsii_.InvokeVoid(
		g,
		"resetAllowCredentials",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherDefaultRouteActionCorsPolicyOutputReference) ResetAllowHeaders() {
	_jsii_.InvokeVoid(
		g,
		"resetAllowHeaders",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherDefaultRouteActionCorsPolicyOutputReference) ResetAllowMethods() {
	_jsii_.InvokeVoid(
		g,
		"resetAllowMethods",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherDefaultRouteActionCorsPolicyOutputReference) ResetAllowOriginRegexes() {
	_jsii_.InvokeVoid(
		g,
		"resetAllowOriginRegexes",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherDefaultRouteActionCorsPolicyOutputReference) ResetAllowOrigins() {
	_jsii_.InvokeVoid(
		g,
		"resetAllowOrigins",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherDefaultRouteActionCorsPolicyOutputReference) ResetDisabled() {
	_jsii_.InvokeVoid(
		g,
		"resetDisabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherDefaultRouteActionCorsPolicyOutputReference) ResetExposeHeaders() {
	_jsii_.InvokeVoid(
		g,
		"resetExposeHeaders",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherDefaultRouteActionCorsPolicyOutputReference) ResetMaxAge() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxAge",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherDefaultRouteActionCorsPolicyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherDefaultRouteActionCorsPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

