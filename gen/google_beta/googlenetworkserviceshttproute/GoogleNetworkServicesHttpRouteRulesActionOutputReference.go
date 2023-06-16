package googlenetworkserviceshttproute

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlenetworkserviceshttproute/internal"
)

type GoogleNetworkServicesHttpRouteRulesActionOutputReference interface {
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
	CorsPolicy() GoogleNetworkServicesHttpRouteRulesActionCorsPolicyOutputReference
	CorsPolicyInput() *GoogleNetworkServicesHttpRouteRulesActionCorsPolicy
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Destinations() GoogleNetworkServicesHttpRouteRulesActionDestinationsList
	DestinationsInput() interface{}
	FaultInjectionPolicy() GoogleNetworkServicesHttpRouteRulesActionFaultInjectionPolicyOutputReference
	FaultInjectionPolicyInput() *GoogleNetworkServicesHttpRouteRulesActionFaultInjectionPolicy
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleNetworkServicesHttpRouteRulesAction
	SetInternalValue(val *GoogleNetworkServicesHttpRouteRulesAction)
	Redirect() GoogleNetworkServicesHttpRouteRulesActionRedirectOutputReference
	RedirectInput() *GoogleNetworkServicesHttpRouteRulesActionRedirect
	RequestHeaderModifier() GoogleNetworkServicesHttpRouteRulesActionRequestHeaderModifierOutputReference
	RequestHeaderModifierInput() *GoogleNetworkServicesHttpRouteRulesActionRequestHeaderModifier
	RequestMirrorPolicy() GoogleNetworkServicesHttpRouteRulesActionRequestMirrorPolicyOutputReference
	RequestMirrorPolicyInput() *GoogleNetworkServicesHttpRouteRulesActionRequestMirrorPolicy
	ResponseHeaderModifier() GoogleNetworkServicesHttpRouteRulesActionResponseHeaderModifierOutputReference
	ResponseHeaderModifierInput() *GoogleNetworkServicesHttpRouteRulesActionResponseHeaderModifier
	RetryPolicy() GoogleNetworkServicesHttpRouteRulesActionRetryPolicyOutputReference
	RetryPolicyInput() *GoogleNetworkServicesHttpRouteRulesActionRetryPolicy
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Timeout() *string
	SetTimeout(val *string)
	TimeoutInput() *string
	UrlRewrite() GoogleNetworkServicesHttpRouteRulesActionUrlRewriteOutputReference
	UrlRewriteInput() *GoogleNetworkServicesHttpRouteRulesActionUrlRewrite
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
	PutCorsPolicy(value *GoogleNetworkServicesHttpRouteRulesActionCorsPolicy)
	PutDestinations(value interface{})
	PutFaultInjectionPolicy(value *GoogleNetworkServicesHttpRouteRulesActionFaultInjectionPolicy)
	PutRedirect(value *GoogleNetworkServicesHttpRouteRulesActionRedirect)
	PutRequestHeaderModifier(value *GoogleNetworkServicesHttpRouteRulesActionRequestHeaderModifier)
	PutRequestMirrorPolicy(value *GoogleNetworkServicesHttpRouteRulesActionRequestMirrorPolicy)
	PutResponseHeaderModifier(value *GoogleNetworkServicesHttpRouteRulesActionResponseHeaderModifier)
	PutRetryPolicy(value *GoogleNetworkServicesHttpRouteRulesActionRetryPolicy)
	PutUrlRewrite(value *GoogleNetworkServicesHttpRouteRulesActionUrlRewrite)
	ResetCorsPolicy()
	ResetDestinations()
	ResetFaultInjectionPolicy()
	ResetRedirect()
	ResetRequestHeaderModifier()
	ResetRequestMirrorPolicy()
	ResetResponseHeaderModifier()
	ResetRetryPolicy()
	ResetTimeout()
	ResetUrlRewrite()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleNetworkServicesHttpRouteRulesActionOutputReference
type jsiiProxy_GoogleNetworkServicesHttpRouteRulesActionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleNetworkServicesHttpRouteRulesActionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesHttpRouteRulesActionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesHttpRouteRulesActionOutputReference) CorsPolicy() GoogleNetworkServicesHttpRouteRulesActionCorsPolicyOutputReference {
	var returns GoogleNetworkServicesHttpRouteRulesActionCorsPolicyOutputReference
	_jsii_.Get(
		j,
		"corsPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesHttpRouteRulesActionOutputReference) CorsPolicyInput() *GoogleNetworkServicesHttpRouteRulesActionCorsPolicy {
	var returns *GoogleNetworkServicesHttpRouteRulesActionCorsPolicy
	_jsii_.Get(
		j,
		"corsPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesHttpRouteRulesActionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesHttpRouteRulesActionOutputReference) Destinations() GoogleNetworkServicesHttpRouteRulesActionDestinationsList {
	var returns GoogleNetworkServicesHttpRouteRulesActionDestinationsList
	_jsii_.Get(
		j,
		"destinations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesHttpRouteRulesActionOutputReference) DestinationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"destinationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesHttpRouteRulesActionOutputReference) FaultInjectionPolicy() GoogleNetworkServicesHttpRouteRulesActionFaultInjectionPolicyOutputReference {
	var returns GoogleNetworkServicesHttpRouteRulesActionFaultInjectionPolicyOutputReference
	_jsii_.Get(
		j,
		"faultInjectionPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesHttpRouteRulesActionOutputReference) FaultInjectionPolicyInput() *GoogleNetworkServicesHttpRouteRulesActionFaultInjectionPolicy {
	var returns *GoogleNetworkServicesHttpRouteRulesActionFaultInjectionPolicy
	_jsii_.Get(
		j,
		"faultInjectionPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesHttpRouteRulesActionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesHttpRouteRulesActionOutputReference) InternalValue() *GoogleNetworkServicesHttpRouteRulesAction {
	var returns *GoogleNetworkServicesHttpRouteRulesAction
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesHttpRouteRulesActionOutputReference) Redirect() GoogleNetworkServicesHttpRouteRulesActionRedirectOutputReference {
	var returns GoogleNetworkServicesHttpRouteRulesActionRedirectOutputReference
	_jsii_.Get(
		j,
		"redirect",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesHttpRouteRulesActionOutputReference) RedirectInput() *GoogleNetworkServicesHttpRouteRulesActionRedirect {
	var returns *GoogleNetworkServicesHttpRouteRulesActionRedirect
	_jsii_.Get(
		j,
		"redirectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesHttpRouteRulesActionOutputReference) RequestHeaderModifier() GoogleNetworkServicesHttpRouteRulesActionRequestHeaderModifierOutputReference {
	var returns GoogleNetworkServicesHttpRouteRulesActionRequestHeaderModifierOutputReference
	_jsii_.Get(
		j,
		"requestHeaderModifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesHttpRouteRulesActionOutputReference) RequestHeaderModifierInput() *GoogleNetworkServicesHttpRouteRulesActionRequestHeaderModifier {
	var returns *GoogleNetworkServicesHttpRouteRulesActionRequestHeaderModifier
	_jsii_.Get(
		j,
		"requestHeaderModifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesHttpRouteRulesActionOutputReference) RequestMirrorPolicy() GoogleNetworkServicesHttpRouteRulesActionRequestMirrorPolicyOutputReference {
	var returns GoogleNetworkServicesHttpRouteRulesActionRequestMirrorPolicyOutputReference
	_jsii_.Get(
		j,
		"requestMirrorPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesHttpRouteRulesActionOutputReference) RequestMirrorPolicyInput() *GoogleNetworkServicesHttpRouteRulesActionRequestMirrorPolicy {
	var returns *GoogleNetworkServicesHttpRouteRulesActionRequestMirrorPolicy
	_jsii_.Get(
		j,
		"requestMirrorPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesHttpRouteRulesActionOutputReference) ResponseHeaderModifier() GoogleNetworkServicesHttpRouteRulesActionResponseHeaderModifierOutputReference {
	var returns GoogleNetworkServicesHttpRouteRulesActionResponseHeaderModifierOutputReference
	_jsii_.Get(
		j,
		"responseHeaderModifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesHttpRouteRulesActionOutputReference) ResponseHeaderModifierInput() *GoogleNetworkServicesHttpRouteRulesActionResponseHeaderModifier {
	var returns *GoogleNetworkServicesHttpRouteRulesActionResponseHeaderModifier
	_jsii_.Get(
		j,
		"responseHeaderModifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesHttpRouteRulesActionOutputReference) RetryPolicy() GoogleNetworkServicesHttpRouteRulesActionRetryPolicyOutputReference {
	var returns GoogleNetworkServicesHttpRouteRulesActionRetryPolicyOutputReference
	_jsii_.Get(
		j,
		"retryPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesHttpRouteRulesActionOutputReference) RetryPolicyInput() *GoogleNetworkServicesHttpRouteRulesActionRetryPolicy {
	var returns *GoogleNetworkServicesHttpRouteRulesActionRetryPolicy
	_jsii_.Get(
		j,
		"retryPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesHttpRouteRulesActionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesHttpRouteRulesActionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesHttpRouteRulesActionOutputReference) Timeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesHttpRouteRulesActionOutputReference) TimeoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesHttpRouteRulesActionOutputReference) UrlRewrite() GoogleNetworkServicesHttpRouteRulesActionUrlRewriteOutputReference {
	var returns GoogleNetworkServicesHttpRouteRulesActionUrlRewriteOutputReference
	_jsii_.Get(
		j,
		"urlRewrite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesHttpRouteRulesActionOutputReference) UrlRewriteInput() *GoogleNetworkServicesHttpRouteRulesActionUrlRewrite {
	var returns *GoogleNetworkServicesHttpRouteRulesActionUrlRewrite
	_jsii_.Get(
		j,
		"urlRewriteInput",
		&returns,
	)
	return returns
}


func NewGoogleNetworkServicesHttpRouteRulesActionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleNetworkServicesHttpRouteRulesActionOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleNetworkServicesHttpRouteRulesActionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleNetworkServicesHttpRouteRulesActionOutputReference{}

	_jsii_.Create(
		"google-beta.googleNetworkServicesHttpRoute.GoogleNetworkServicesHttpRouteRulesActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleNetworkServicesHttpRouteRulesActionOutputReference_Override(g GoogleNetworkServicesHttpRouteRulesActionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleNetworkServicesHttpRoute.GoogleNetworkServicesHttpRouteRulesActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleNetworkServicesHttpRouteRulesActionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkServicesHttpRouteRulesActionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkServicesHttpRouteRulesActionOutputReference)SetInternalValue(val *GoogleNetworkServicesHttpRouteRulesAction) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkServicesHttpRouteRulesActionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkServicesHttpRouteRulesActionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkServicesHttpRouteRulesActionOutputReference)SetTimeout(val *string) {
	if err := j.validateSetTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeout",
		val,
	)
}

func (g *jsiiProxy_GoogleNetworkServicesHttpRouteRulesActionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetworkServicesHttpRouteRulesActionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleNetworkServicesHttpRouteRulesActionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleNetworkServicesHttpRouteRulesActionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleNetworkServicesHttpRouteRulesActionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleNetworkServicesHttpRouteRulesActionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleNetworkServicesHttpRouteRulesActionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleNetworkServicesHttpRouteRulesActionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleNetworkServicesHttpRouteRulesActionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleNetworkServicesHttpRouteRulesActionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleNetworkServicesHttpRouteRulesActionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetworkServicesHttpRouteRulesActionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleNetworkServicesHttpRouteRulesActionOutputReference) PutCorsPolicy(value *GoogleNetworkServicesHttpRouteRulesActionCorsPolicy) {
	if err := g.validatePutCorsPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putCorsPolicy",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleNetworkServicesHttpRouteRulesActionOutputReference) PutDestinations(value interface{}) {
	if err := g.validatePutDestinationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDestinations",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleNetworkServicesHttpRouteRulesActionOutputReference) PutFaultInjectionPolicy(value *GoogleNetworkServicesHttpRouteRulesActionFaultInjectionPolicy) {
	if err := g.validatePutFaultInjectionPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putFaultInjectionPolicy",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleNetworkServicesHttpRouteRulesActionOutputReference) PutRedirect(value *GoogleNetworkServicesHttpRouteRulesActionRedirect) {
	if err := g.validatePutRedirectParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRedirect",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleNetworkServicesHttpRouteRulesActionOutputReference) PutRequestHeaderModifier(value *GoogleNetworkServicesHttpRouteRulesActionRequestHeaderModifier) {
	if err := g.validatePutRequestHeaderModifierParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRequestHeaderModifier",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleNetworkServicesHttpRouteRulesActionOutputReference) PutRequestMirrorPolicy(value *GoogleNetworkServicesHttpRouteRulesActionRequestMirrorPolicy) {
	if err := g.validatePutRequestMirrorPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRequestMirrorPolicy",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleNetworkServicesHttpRouteRulesActionOutputReference) PutResponseHeaderModifier(value *GoogleNetworkServicesHttpRouteRulesActionResponseHeaderModifier) {
	if err := g.validatePutResponseHeaderModifierParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putResponseHeaderModifier",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleNetworkServicesHttpRouteRulesActionOutputReference) PutRetryPolicy(value *GoogleNetworkServicesHttpRouteRulesActionRetryPolicy) {
	if err := g.validatePutRetryPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRetryPolicy",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleNetworkServicesHttpRouteRulesActionOutputReference) PutUrlRewrite(value *GoogleNetworkServicesHttpRouteRulesActionUrlRewrite) {
	if err := g.validatePutUrlRewriteParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putUrlRewrite",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleNetworkServicesHttpRouteRulesActionOutputReference) ResetCorsPolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetCorsPolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkServicesHttpRouteRulesActionOutputReference) ResetDestinations() {
	_jsii_.InvokeVoid(
		g,
		"resetDestinations",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkServicesHttpRouteRulesActionOutputReference) ResetFaultInjectionPolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetFaultInjectionPolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkServicesHttpRouteRulesActionOutputReference) ResetRedirect() {
	_jsii_.InvokeVoid(
		g,
		"resetRedirect",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkServicesHttpRouteRulesActionOutputReference) ResetRequestHeaderModifier() {
	_jsii_.InvokeVoid(
		g,
		"resetRequestHeaderModifier",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkServicesHttpRouteRulesActionOutputReference) ResetRequestMirrorPolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetRequestMirrorPolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkServicesHttpRouteRulesActionOutputReference) ResetResponseHeaderModifier() {
	_jsii_.InvokeVoid(
		g,
		"resetResponseHeaderModifier",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkServicesHttpRouteRulesActionOutputReference) ResetRetryPolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetRetryPolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkServicesHttpRouteRulesActionOutputReference) ResetTimeout() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeout",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkServicesHttpRouteRulesActionOutputReference) ResetUrlRewrite() {
	_jsii_.InvokeVoid(
		g,
		"resetUrlRewrite",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkServicesHttpRouteRulesActionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleNetworkServicesHttpRouteRulesActionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

