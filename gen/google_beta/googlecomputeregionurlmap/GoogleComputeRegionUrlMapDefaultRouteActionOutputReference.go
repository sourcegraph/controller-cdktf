package googlecomputeregionurlmap

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecomputeregionurlmap/internal"
)

type GoogleComputeRegionUrlMapDefaultRouteActionOutputReference interface {
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
	CorsPolicy() GoogleComputeRegionUrlMapDefaultRouteActionCorsPolicyOutputReference
	CorsPolicyInput() *GoogleComputeRegionUrlMapDefaultRouteActionCorsPolicy
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	FaultInjectionPolicy() GoogleComputeRegionUrlMapDefaultRouteActionFaultInjectionPolicyOutputReference
	FaultInjectionPolicyInput() *GoogleComputeRegionUrlMapDefaultRouteActionFaultInjectionPolicy
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleComputeRegionUrlMapDefaultRouteAction
	SetInternalValue(val *GoogleComputeRegionUrlMapDefaultRouteAction)
	RequestMirrorPolicy() GoogleComputeRegionUrlMapDefaultRouteActionRequestMirrorPolicyOutputReference
	RequestMirrorPolicyInput() *GoogleComputeRegionUrlMapDefaultRouteActionRequestMirrorPolicy
	RetryPolicy() GoogleComputeRegionUrlMapDefaultRouteActionRetryPolicyOutputReference
	RetryPolicyInput() *GoogleComputeRegionUrlMapDefaultRouteActionRetryPolicy
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Timeout() GoogleComputeRegionUrlMapDefaultRouteActionTimeoutOutputReference
	TimeoutInput() *GoogleComputeRegionUrlMapDefaultRouteActionTimeout
	UrlRewrite() GoogleComputeRegionUrlMapDefaultRouteActionUrlRewriteOutputReference
	UrlRewriteInput() *GoogleComputeRegionUrlMapDefaultRouteActionUrlRewrite
	WeightedBackendServices() GoogleComputeRegionUrlMapDefaultRouteActionWeightedBackendServicesList
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
	PutCorsPolicy(value *GoogleComputeRegionUrlMapDefaultRouteActionCorsPolicy)
	PutFaultInjectionPolicy(value *GoogleComputeRegionUrlMapDefaultRouteActionFaultInjectionPolicy)
	PutRequestMirrorPolicy(value *GoogleComputeRegionUrlMapDefaultRouteActionRequestMirrorPolicy)
	PutRetryPolicy(value *GoogleComputeRegionUrlMapDefaultRouteActionRetryPolicy)
	PutTimeout(value *GoogleComputeRegionUrlMapDefaultRouteActionTimeout)
	PutUrlRewrite(value *GoogleComputeRegionUrlMapDefaultRouteActionUrlRewrite)
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

// The jsii proxy struct for GoogleComputeRegionUrlMapDefaultRouteActionOutputReference
type jsiiProxy_GoogleComputeRegionUrlMapDefaultRouteActionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapDefaultRouteActionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapDefaultRouteActionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapDefaultRouteActionOutputReference) CorsPolicy() GoogleComputeRegionUrlMapDefaultRouteActionCorsPolicyOutputReference {
	var returns GoogleComputeRegionUrlMapDefaultRouteActionCorsPolicyOutputReference
	_jsii_.Get(
		j,
		"corsPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapDefaultRouteActionOutputReference) CorsPolicyInput() *GoogleComputeRegionUrlMapDefaultRouteActionCorsPolicy {
	var returns *GoogleComputeRegionUrlMapDefaultRouteActionCorsPolicy
	_jsii_.Get(
		j,
		"corsPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapDefaultRouteActionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapDefaultRouteActionOutputReference) FaultInjectionPolicy() GoogleComputeRegionUrlMapDefaultRouteActionFaultInjectionPolicyOutputReference {
	var returns GoogleComputeRegionUrlMapDefaultRouteActionFaultInjectionPolicyOutputReference
	_jsii_.Get(
		j,
		"faultInjectionPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapDefaultRouteActionOutputReference) FaultInjectionPolicyInput() *GoogleComputeRegionUrlMapDefaultRouteActionFaultInjectionPolicy {
	var returns *GoogleComputeRegionUrlMapDefaultRouteActionFaultInjectionPolicy
	_jsii_.Get(
		j,
		"faultInjectionPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapDefaultRouteActionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapDefaultRouteActionOutputReference) InternalValue() *GoogleComputeRegionUrlMapDefaultRouteAction {
	var returns *GoogleComputeRegionUrlMapDefaultRouteAction
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapDefaultRouteActionOutputReference) RequestMirrorPolicy() GoogleComputeRegionUrlMapDefaultRouteActionRequestMirrorPolicyOutputReference {
	var returns GoogleComputeRegionUrlMapDefaultRouteActionRequestMirrorPolicyOutputReference
	_jsii_.Get(
		j,
		"requestMirrorPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapDefaultRouteActionOutputReference) RequestMirrorPolicyInput() *GoogleComputeRegionUrlMapDefaultRouteActionRequestMirrorPolicy {
	var returns *GoogleComputeRegionUrlMapDefaultRouteActionRequestMirrorPolicy
	_jsii_.Get(
		j,
		"requestMirrorPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapDefaultRouteActionOutputReference) RetryPolicy() GoogleComputeRegionUrlMapDefaultRouteActionRetryPolicyOutputReference {
	var returns GoogleComputeRegionUrlMapDefaultRouteActionRetryPolicyOutputReference
	_jsii_.Get(
		j,
		"retryPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapDefaultRouteActionOutputReference) RetryPolicyInput() *GoogleComputeRegionUrlMapDefaultRouteActionRetryPolicy {
	var returns *GoogleComputeRegionUrlMapDefaultRouteActionRetryPolicy
	_jsii_.Get(
		j,
		"retryPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapDefaultRouteActionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapDefaultRouteActionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapDefaultRouteActionOutputReference) Timeout() GoogleComputeRegionUrlMapDefaultRouteActionTimeoutOutputReference {
	var returns GoogleComputeRegionUrlMapDefaultRouteActionTimeoutOutputReference
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapDefaultRouteActionOutputReference) TimeoutInput() *GoogleComputeRegionUrlMapDefaultRouteActionTimeout {
	var returns *GoogleComputeRegionUrlMapDefaultRouteActionTimeout
	_jsii_.Get(
		j,
		"timeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapDefaultRouteActionOutputReference) UrlRewrite() GoogleComputeRegionUrlMapDefaultRouteActionUrlRewriteOutputReference {
	var returns GoogleComputeRegionUrlMapDefaultRouteActionUrlRewriteOutputReference
	_jsii_.Get(
		j,
		"urlRewrite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapDefaultRouteActionOutputReference) UrlRewriteInput() *GoogleComputeRegionUrlMapDefaultRouteActionUrlRewrite {
	var returns *GoogleComputeRegionUrlMapDefaultRouteActionUrlRewrite
	_jsii_.Get(
		j,
		"urlRewriteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapDefaultRouteActionOutputReference) WeightedBackendServices() GoogleComputeRegionUrlMapDefaultRouteActionWeightedBackendServicesList {
	var returns GoogleComputeRegionUrlMapDefaultRouteActionWeightedBackendServicesList
	_jsii_.Get(
		j,
		"weightedBackendServices",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapDefaultRouteActionOutputReference) WeightedBackendServicesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"weightedBackendServicesInput",
		&returns,
	)
	return returns
}


func NewGoogleComputeRegionUrlMapDefaultRouteActionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleComputeRegionUrlMapDefaultRouteActionOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleComputeRegionUrlMapDefaultRouteActionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeRegionUrlMapDefaultRouteActionOutputReference{}

	_jsii_.Create(
		"google-beta.googleComputeRegionUrlMap.GoogleComputeRegionUrlMapDefaultRouteActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleComputeRegionUrlMapDefaultRouteActionOutputReference_Override(g GoogleComputeRegionUrlMapDefaultRouteActionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleComputeRegionUrlMap.GoogleComputeRegionUrlMapDefaultRouteActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapDefaultRouteActionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapDefaultRouteActionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapDefaultRouteActionOutputReference)SetInternalValue(val *GoogleComputeRegionUrlMapDefaultRouteAction) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapDefaultRouteActionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapDefaultRouteActionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleComputeRegionUrlMapDefaultRouteActionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeRegionUrlMapDefaultRouteActionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleComputeRegionUrlMapDefaultRouteActionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeRegionUrlMapDefaultRouteActionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComputeRegionUrlMapDefaultRouteActionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComputeRegionUrlMapDefaultRouteActionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComputeRegionUrlMapDefaultRouteActionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComputeRegionUrlMapDefaultRouteActionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComputeRegionUrlMapDefaultRouteActionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComputeRegionUrlMapDefaultRouteActionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComputeRegionUrlMapDefaultRouteActionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeRegionUrlMapDefaultRouteActionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeRegionUrlMapDefaultRouteActionOutputReference) PutCorsPolicy(value *GoogleComputeRegionUrlMapDefaultRouteActionCorsPolicy) {
	if err := g.validatePutCorsPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putCorsPolicy",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeRegionUrlMapDefaultRouteActionOutputReference) PutFaultInjectionPolicy(value *GoogleComputeRegionUrlMapDefaultRouteActionFaultInjectionPolicy) {
	if err := g.validatePutFaultInjectionPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putFaultInjectionPolicy",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeRegionUrlMapDefaultRouteActionOutputReference) PutRequestMirrorPolicy(value *GoogleComputeRegionUrlMapDefaultRouteActionRequestMirrorPolicy) {
	if err := g.validatePutRequestMirrorPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRequestMirrorPolicy",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeRegionUrlMapDefaultRouteActionOutputReference) PutRetryPolicy(value *GoogleComputeRegionUrlMapDefaultRouteActionRetryPolicy) {
	if err := g.validatePutRetryPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRetryPolicy",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeRegionUrlMapDefaultRouteActionOutputReference) PutTimeout(value *GoogleComputeRegionUrlMapDefaultRouteActionTimeout) {
	if err := g.validatePutTimeoutParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeout",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeRegionUrlMapDefaultRouteActionOutputReference) PutUrlRewrite(value *GoogleComputeRegionUrlMapDefaultRouteActionUrlRewrite) {
	if err := g.validatePutUrlRewriteParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putUrlRewrite",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeRegionUrlMapDefaultRouteActionOutputReference) PutWeightedBackendServices(value interface{}) {
	if err := g.validatePutWeightedBackendServicesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putWeightedBackendServices",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeRegionUrlMapDefaultRouteActionOutputReference) ResetCorsPolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetCorsPolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionUrlMapDefaultRouteActionOutputReference) ResetFaultInjectionPolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetFaultInjectionPolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionUrlMapDefaultRouteActionOutputReference) ResetRequestMirrorPolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetRequestMirrorPolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionUrlMapDefaultRouteActionOutputReference) ResetRetryPolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetRetryPolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionUrlMapDefaultRouteActionOutputReference) ResetTimeout() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeout",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionUrlMapDefaultRouteActionOutputReference) ResetUrlRewrite() {
	_jsii_.InvokeVoid(
		g,
		"resetUrlRewrite",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionUrlMapDefaultRouteActionOutputReference) ResetWeightedBackendServices() {
	_jsii_.InvokeVoid(
		g,
		"resetWeightedBackendServices",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionUrlMapDefaultRouteActionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleComputeRegionUrlMapDefaultRouteActionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

