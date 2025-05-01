package googlecomputeurlmap

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecomputeurlmap/internal"
)

type GoogleComputeUrlMapPathMatcherPathRuleRouteActionOutputReference interface {
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
	CorsPolicy() GoogleComputeUrlMapPathMatcherPathRuleRouteActionCorsPolicyOutputReference
	CorsPolicyInput() *GoogleComputeUrlMapPathMatcherPathRuleRouteActionCorsPolicy
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	FaultInjectionPolicy() GoogleComputeUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyOutputReference
	FaultInjectionPolicyInput() *GoogleComputeUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleComputeUrlMapPathMatcherPathRuleRouteAction
	SetInternalValue(val *GoogleComputeUrlMapPathMatcherPathRuleRouteAction)
	MaxStreamDuration() GoogleComputeUrlMapPathMatcherPathRuleRouteActionMaxStreamDurationOutputReference
	MaxStreamDurationInput() *GoogleComputeUrlMapPathMatcherPathRuleRouteActionMaxStreamDuration
	RequestMirrorPolicy() GoogleComputeUrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicyOutputReference
	RequestMirrorPolicyInput() *GoogleComputeUrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy
	RetryPolicy() GoogleComputeUrlMapPathMatcherPathRuleRouteActionRetryPolicyOutputReference
	RetryPolicyInput() *GoogleComputeUrlMapPathMatcherPathRuleRouteActionRetryPolicy
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Timeout() GoogleComputeUrlMapPathMatcherPathRuleRouteActionTimeoutOutputReference
	TimeoutInput() *GoogleComputeUrlMapPathMatcherPathRuleRouteActionTimeout
	UrlRewrite() GoogleComputeUrlMapPathMatcherPathRuleRouteActionUrlRewriteOutputReference
	UrlRewriteInput() *GoogleComputeUrlMapPathMatcherPathRuleRouteActionUrlRewrite
	WeightedBackendServices() GoogleComputeUrlMapPathMatcherPathRuleRouteActionWeightedBackendServicesList
	WeightedBackendServicesInput() interface{}
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
	PutCorsPolicy(value *GoogleComputeUrlMapPathMatcherPathRuleRouteActionCorsPolicy)
	PutFaultInjectionPolicy(value *GoogleComputeUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy)
	PutMaxStreamDuration(value *GoogleComputeUrlMapPathMatcherPathRuleRouteActionMaxStreamDuration)
	PutRequestMirrorPolicy(value *GoogleComputeUrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy)
	PutRetryPolicy(value *GoogleComputeUrlMapPathMatcherPathRuleRouteActionRetryPolicy)
	PutTimeout(value *GoogleComputeUrlMapPathMatcherPathRuleRouteActionTimeout)
	PutUrlRewrite(value *GoogleComputeUrlMapPathMatcherPathRuleRouteActionUrlRewrite)
	PutWeightedBackendServices(value interface{})
	ResetCorsPolicy()
	ResetFaultInjectionPolicy()
	ResetMaxStreamDuration()
	ResetRequestMirrorPolicy()
	ResetRetryPolicy()
	ResetTimeout()
	ResetUrlRewrite()
	ResetWeightedBackendServices()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleComputeUrlMapPathMatcherPathRuleRouteActionOutputReference
type jsiiProxy_GoogleComputeUrlMapPathMatcherPathRuleRouteActionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherPathRuleRouteActionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherPathRuleRouteActionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherPathRuleRouteActionOutputReference) CorsPolicy() GoogleComputeUrlMapPathMatcherPathRuleRouteActionCorsPolicyOutputReference {
	var returns GoogleComputeUrlMapPathMatcherPathRuleRouteActionCorsPolicyOutputReference
	_jsii_.Get(
		j,
		"corsPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherPathRuleRouteActionOutputReference) CorsPolicyInput() *GoogleComputeUrlMapPathMatcherPathRuleRouteActionCorsPolicy {
	var returns *GoogleComputeUrlMapPathMatcherPathRuleRouteActionCorsPolicy
	_jsii_.Get(
		j,
		"corsPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherPathRuleRouteActionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherPathRuleRouteActionOutputReference) FaultInjectionPolicy() GoogleComputeUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyOutputReference {
	var returns GoogleComputeUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyOutputReference
	_jsii_.Get(
		j,
		"faultInjectionPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherPathRuleRouteActionOutputReference) FaultInjectionPolicyInput() *GoogleComputeUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy {
	var returns *GoogleComputeUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy
	_jsii_.Get(
		j,
		"faultInjectionPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherPathRuleRouteActionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherPathRuleRouteActionOutputReference) InternalValue() *GoogleComputeUrlMapPathMatcherPathRuleRouteAction {
	var returns *GoogleComputeUrlMapPathMatcherPathRuleRouteAction
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherPathRuleRouteActionOutputReference) MaxStreamDuration() GoogleComputeUrlMapPathMatcherPathRuleRouteActionMaxStreamDurationOutputReference {
	var returns GoogleComputeUrlMapPathMatcherPathRuleRouteActionMaxStreamDurationOutputReference
	_jsii_.Get(
		j,
		"maxStreamDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherPathRuleRouteActionOutputReference) MaxStreamDurationInput() *GoogleComputeUrlMapPathMatcherPathRuleRouteActionMaxStreamDuration {
	var returns *GoogleComputeUrlMapPathMatcherPathRuleRouteActionMaxStreamDuration
	_jsii_.Get(
		j,
		"maxStreamDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherPathRuleRouteActionOutputReference) RequestMirrorPolicy() GoogleComputeUrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicyOutputReference {
	var returns GoogleComputeUrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicyOutputReference
	_jsii_.Get(
		j,
		"requestMirrorPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherPathRuleRouteActionOutputReference) RequestMirrorPolicyInput() *GoogleComputeUrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy {
	var returns *GoogleComputeUrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy
	_jsii_.Get(
		j,
		"requestMirrorPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherPathRuleRouteActionOutputReference) RetryPolicy() GoogleComputeUrlMapPathMatcherPathRuleRouteActionRetryPolicyOutputReference {
	var returns GoogleComputeUrlMapPathMatcherPathRuleRouteActionRetryPolicyOutputReference
	_jsii_.Get(
		j,
		"retryPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherPathRuleRouteActionOutputReference) RetryPolicyInput() *GoogleComputeUrlMapPathMatcherPathRuleRouteActionRetryPolicy {
	var returns *GoogleComputeUrlMapPathMatcherPathRuleRouteActionRetryPolicy
	_jsii_.Get(
		j,
		"retryPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherPathRuleRouteActionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherPathRuleRouteActionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherPathRuleRouteActionOutputReference) Timeout() GoogleComputeUrlMapPathMatcherPathRuleRouteActionTimeoutOutputReference {
	var returns GoogleComputeUrlMapPathMatcherPathRuleRouteActionTimeoutOutputReference
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherPathRuleRouteActionOutputReference) TimeoutInput() *GoogleComputeUrlMapPathMatcherPathRuleRouteActionTimeout {
	var returns *GoogleComputeUrlMapPathMatcherPathRuleRouteActionTimeout
	_jsii_.Get(
		j,
		"timeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherPathRuleRouteActionOutputReference) UrlRewrite() GoogleComputeUrlMapPathMatcherPathRuleRouteActionUrlRewriteOutputReference {
	var returns GoogleComputeUrlMapPathMatcherPathRuleRouteActionUrlRewriteOutputReference
	_jsii_.Get(
		j,
		"urlRewrite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherPathRuleRouteActionOutputReference) UrlRewriteInput() *GoogleComputeUrlMapPathMatcherPathRuleRouteActionUrlRewrite {
	var returns *GoogleComputeUrlMapPathMatcherPathRuleRouteActionUrlRewrite
	_jsii_.Get(
		j,
		"urlRewriteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherPathRuleRouteActionOutputReference) WeightedBackendServices() GoogleComputeUrlMapPathMatcherPathRuleRouteActionWeightedBackendServicesList {
	var returns GoogleComputeUrlMapPathMatcherPathRuleRouteActionWeightedBackendServicesList
	_jsii_.Get(
		j,
		"weightedBackendServices",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherPathRuleRouteActionOutputReference) WeightedBackendServicesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"weightedBackendServicesInput",
		&returns,
	)
	return returns
}


func NewGoogleComputeUrlMapPathMatcherPathRuleRouteActionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleComputeUrlMapPathMatcherPathRuleRouteActionOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleComputeUrlMapPathMatcherPathRuleRouteActionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeUrlMapPathMatcherPathRuleRouteActionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleComputeUrlMap.GoogleComputeUrlMapPathMatcherPathRuleRouteActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleComputeUrlMapPathMatcherPathRuleRouteActionOutputReference_Override(g GoogleComputeUrlMapPathMatcherPathRuleRouteActionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleComputeUrlMap.GoogleComputeUrlMapPathMatcherPathRuleRouteActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherPathRuleRouteActionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherPathRuleRouteActionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherPathRuleRouteActionOutputReference)SetInternalValue(val *GoogleComputeUrlMapPathMatcherPathRuleRouteAction) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherPathRuleRouteActionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherPathRuleRouteActionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherPathRuleRouteActionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherPathRuleRouteActionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherPathRuleRouteActionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherPathRuleRouteActionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherPathRuleRouteActionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherPathRuleRouteActionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherPathRuleRouteActionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherPathRuleRouteActionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherPathRuleRouteActionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherPathRuleRouteActionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherPathRuleRouteActionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherPathRuleRouteActionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherPathRuleRouteActionOutputReference) PutCorsPolicy(value *GoogleComputeUrlMapPathMatcherPathRuleRouteActionCorsPolicy) {
	if err := g.validatePutCorsPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putCorsPolicy",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherPathRuleRouteActionOutputReference) PutFaultInjectionPolicy(value *GoogleComputeUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy) {
	if err := g.validatePutFaultInjectionPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putFaultInjectionPolicy",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherPathRuleRouteActionOutputReference) PutMaxStreamDuration(value *GoogleComputeUrlMapPathMatcherPathRuleRouteActionMaxStreamDuration) {
	if err := g.validatePutMaxStreamDurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMaxStreamDuration",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherPathRuleRouteActionOutputReference) PutRequestMirrorPolicy(value *GoogleComputeUrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy) {
	if err := g.validatePutRequestMirrorPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRequestMirrorPolicy",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherPathRuleRouteActionOutputReference) PutRetryPolicy(value *GoogleComputeUrlMapPathMatcherPathRuleRouteActionRetryPolicy) {
	if err := g.validatePutRetryPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRetryPolicy",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherPathRuleRouteActionOutputReference) PutTimeout(value *GoogleComputeUrlMapPathMatcherPathRuleRouteActionTimeout) {
	if err := g.validatePutTimeoutParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeout",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherPathRuleRouteActionOutputReference) PutUrlRewrite(value *GoogleComputeUrlMapPathMatcherPathRuleRouteActionUrlRewrite) {
	if err := g.validatePutUrlRewriteParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putUrlRewrite",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherPathRuleRouteActionOutputReference) PutWeightedBackendServices(value interface{}) {
	if err := g.validatePutWeightedBackendServicesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putWeightedBackendServices",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherPathRuleRouteActionOutputReference) ResetCorsPolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetCorsPolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherPathRuleRouteActionOutputReference) ResetFaultInjectionPolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetFaultInjectionPolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherPathRuleRouteActionOutputReference) ResetMaxStreamDuration() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxStreamDuration",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherPathRuleRouteActionOutputReference) ResetRequestMirrorPolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetRequestMirrorPolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherPathRuleRouteActionOutputReference) ResetRetryPolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetRetryPolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherPathRuleRouteActionOutputReference) ResetTimeout() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeout",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherPathRuleRouteActionOutputReference) ResetUrlRewrite() {
	_jsii_.InvokeVoid(
		g,
		"resetUrlRewrite",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherPathRuleRouteActionOutputReference) ResetWeightedBackendServices() {
	_jsii_.InvokeVoid(
		g,
		"resetWeightedBackendServices",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherPathRuleRouteActionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherPathRuleRouteActionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

