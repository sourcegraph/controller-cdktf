package googlecomputeurlmap

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecomputeurlmap/internal"
)

type GoogleComputeUrlMapPathMatcherRouteRulesRouteActionOutputReference interface {
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
	CorsPolicy() GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCorsPolicyOutputReference
	CorsPolicyInput() *GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCorsPolicy
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	FaultInjectionPolicy() GoogleComputeUrlMapPathMatcherRouteRulesRouteActionFaultInjectionPolicyOutputReference
	FaultInjectionPolicyInput() *GoogleComputeUrlMapPathMatcherRouteRulesRouteActionFaultInjectionPolicy
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleComputeUrlMapPathMatcherRouteRulesRouteAction
	SetInternalValue(val *GoogleComputeUrlMapPathMatcherRouteRulesRouteAction)
	RequestMirrorPolicy() GoogleComputeUrlMapPathMatcherRouteRulesRouteActionRequestMirrorPolicyOutputReference
	RequestMirrorPolicyInput() *GoogleComputeUrlMapPathMatcherRouteRulesRouteActionRequestMirrorPolicy
	RetryPolicy() GoogleComputeUrlMapPathMatcherRouteRulesRouteActionRetryPolicyOutputReference
	RetryPolicyInput() *GoogleComputeUrlMapPathMatcherRouteRulesRouteActionRetryPolicy
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Timeout() GoogleComputeUrlMapPathMatcherRouteRulesRouteActionTimeoutOutputReference
	TimeoutInput() *GoogleComputeUrlMapPathMatcherRouteRulesRouteActionTimeout
	UrlRewrite() GoogleComputeUrlMapPathMatcherRouteRulesRouteActionUrlRewriteOutputReference
	UrlRewriteInput() *GoogleComputeUrlMapPathMatcherRouteRulesRouteActionUrlRewrite
	WeightedBackendServices() GoogleComputeUrlMapPathMatcherRouteRulesRouteActionWeightedBackendServicesList
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
	PutCorsPolicy(value *GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCorsPolicy)
	PutFaultInjectionPolicy(value *GoogleComputeUrlMapPathMatcherRouteRulesRouteActionFaultInjectionPolicy)
	PutRequestMirrorPolicy(value *GoogleComputeUrlMapPathMatcherRouteRulesRouteActionRequestMirrorPolicy)
	PutRetryPolicy(value *GoogleComputeUrlMapPathMatcherRouteRulesRouteActionRetryPolicy)
	PutTimeout(value *GoogleComputeUrlMapPathMatcherRouteRulesRouteActionTimeout)
	PutUrlRewrite(value *GoogleComputeUrlMapPathMatcherRouteRulesRouteActionUrlRewrite)
	PutWeightedBackendServices(value interface{})
	ResetCorsPolicy()
	ResetFaultInjectionPolicy()
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

// The jsii proxy struct for GoogleComputeUrlMapPathMatcherRouteRulesRouteActionOutputReference
type jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionOutputReference) CorsPolicy() GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCorsPolicyOutputReference {
	var returns GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCorsPolicyOutputReference
	_jsii_.Get(
		j,
		"corsPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionOutputReference) CorsPolicyInput() *GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCorsPolicy {
	var returns *GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCorsPolicy
	_jsii_.Get(
		j,
		"corsPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionOutputReference) FaultInjectionPolicy() GoogleComputeUrlMapPathMatcherRouteRulesRouteActionFaultInjectionPolicyOutputReference {
	var returns GoogleComputeUrlMapPathMatcherRouteRulesRouteActionFaultInjectionPolicyOutputReference
	_jsii_.Get(
		j,
		"faultInjectionPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionOutputReference) FaultInjectionPolicyInput() *GoogleComputeUrlMapPathMatcherRouteRulesRouteActionFaultInjectionPolicy {
	var returns *GoogleComputeUrlMapPathMatcherRouteRulesRouteActionFaultInjectionPolicy
	_jsii_.Get(
		j,
		"faultInjectionPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionOutputReference) InternalValue() *GoogleComputeUrlMapPathMatcherRouteRulesRouteAction {
	var returns *GoogleComputeUrlMapPathMatcherRouteRulesRouteAction
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionOutputReference) RequestMirrorPolicy() GoogleComputeUrlMapPathMatcherRouteRulesRouteActionRequestMirrorPolicyOutputReference {
	var returns GoogleComputeUrlMapPathMatcherRouteRulesRouteActionRequestMirrorPolicyOutputReference
	_jsii_.Get(
		j,
		"requestMirrorPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionOutputReference) RequestMirrorPolicyInput() *GoogleComputeUrlMapPathMatcherRouteRulesRouteActionRequestMirrorPolicy {
	var returns *GoogleComputeUrlMapPathMatcherRouteRulesRouteActionRequestMirrorPolicy
	_jsii_.Get(
		j,
		"requestMirrorPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionOutputReference) RetryPolicy() GoogleComputeUrlMapPathMatcherRouteRulesRouteActionRetryPolicyOutputReference {
	var returns GoogleComputeUrlMapPathMatcherRouteRulesRouteActionRetryPolicyOutputReference
	_jsii_.Get(
		j,
		"retryPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionOutputReference) RetryPolicyInput() *GoogleComputeUrlMapPathMatcherRouteRulesRouteActionRetryPolicy {
	var returns *GoogleComputeUrlMapPathMatcherRouteRulesRouteActionRetryPolicy
	_jsii_.Get(
		j,
		"retryPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionOutputReference) Timeout() GoogleComputeUrlMapPathMatcherRouteRulesRouteActionTimeoutOutputReference {
	var returns GoogleComputeUrlMapPathMatcherRouteRulesRouteActionTimeoutOutputReference
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionOutputReference) TimeoutInput() *GoogleComputeUrlMapPathMatcherRouteRulesRouteActionTimeout {
	var returns *GoogleComputeUrlMapPathMatcherRouteRulesRouteActionTimeout
	_jsii_.Get(
		j,
		"timeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionOutputReference) UrlRewrite() GoogleComputeUrlMapPathMatcherRouteRulesRouteActionUrlRewriteOutputReference {
	var returns GoogleComputeUrlMapPathMatcherRouteRulesRouteActionUrlRewriteOutputReference
	_jsii_.Get(
		j,
		"urlRewrite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionOutputReference) UrlRewriteInput() *GoogleComputeUrlMapPathMatcherRouteRulesRouteActionUrlRewrite {
	var returns *GoogleComputeUrlMapPathMatcherRouteRulesRouteActionUrlRewrite
	_jsii_.Get(
		j,
		"urlRewriteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionOutputReference) WeightedBackendServices() GoogleComputeUrlMapPathMatcherRouteRulesRouteActionWeightedBackendServicesList {
	var returns GoogleComputeUrlMapPathMatcherRouteRulesRouteActionWeightedBackendServicesList
	_jsii_.Get(
		j,
		"weightedBackendServices",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionOutputReference) WeightedBackendServicesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"weightedBackendServicesInput",
		&returns,
	)
	return returns
}


func NewGoogleComputeUrlMapPathMatcherRouteRulesRouteActionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleComputeUrlMapPathMatcherRouteRulesRouteActionOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleComputeUrlMapPathMatcherRouteRulesRouteActionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleComputeUrlMap.GoogleComputeUrlMapPathMatcherRouteRulesRouteActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleComputeUrlMapPathMatcherRouteRulesRouteActionOutputReference_Override(g GoogleComputeUrlMapPathMatcherRouteRulesRouteActionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleComputeUrlMap.GoogleComputeUrlMapPathMatcherRouteRulesRouteActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionOutputReference)SetInternalValue(val *GoogleComputeUrlMapPathMatcherRouteRulesRouteAction) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionOutputReference) PutCorsPolicy(value *GoogleComputeUrlMapPathMatcherRouteRulesRouteActionCorsPolicy) {
	if err := g.validatePutCorsPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putCorsPolicy",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionOutputReference) PutFaultInjectionPolicy(value *GoogleComputeUrlMapPathMatcherRouteRulesRouteActionFaultInjectionPolicy) {
	if err := g.validatePutFaultInjectionPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putFaultInjectionPolicy",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionOutputReference) PutRequestMirrorPolicy(value *GoogleComputeUrlMapPathMatcherRouteRulesRouteActionRequestMirrorPolicy) {
	if err := g.validatePutRequestMirrorPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRequestMirrorPolicy",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionOutputReference) PutRetryPolicy(value *GoogleComputeUrlMapPathMatcherRouteRulesRouteActionRetryPolicy) {
	if err := g.validatePutRetryPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRetryPolicy",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionOutputReference) PutTimeout(value *GoogleComputeUrlMapPathMatcherRouteRulesRouteActionTimeout) {
	if err := g.validatePutTimeoutParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeout",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionOutputReference) PutUrlRewrite(value *GoogleComputeUrlMapPathMatcherRouteRulesRouteActionUrlRewrite) {
	if err := g.validatePutUrlRewriteParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putUrlRewrite",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionOutputReference) PutWeightedBackendServices(value interface{}) {
	if err := g.validatePutWeightedBackendServicesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putWeightedBackendServices",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionOutputReference) ResetCorsPolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetCorsPolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionOutputReference) ResetFaultInjectionPolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetFaultInjectionPolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionOutputReference) ResetRequestMirrorPolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetRequestMirrorPolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionOutputReference) ResetRetryPolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetRetryPolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionOutputReference) ResetTimeout() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeout",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionOutputReference) ResetUrlRewrite() {
	_jsii_.InvokeVoid(
		g,
		"resetUrlRewrite",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionOutputReference) ResetWeightedBackendServices() {
	_jsii_.InvokeVoid(
		g,
		"resetWeightedBackendServices",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleComputeUrlMapPathMatcherRouteRulesRouteActionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

