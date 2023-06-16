package googlecomputeregionurlmap

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecomputeregionurlmap/internal"
)

type GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionOutputReference interface {
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
	CorsPolicy() GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionCorsPolicyOutputReference
	CorsPolicyInput() *GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionCorsPolicy
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	FaultInjectionPolicy() GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionFaultInjectionPolicyOutputReference
	FaultInjectionPolicyInput() *GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionFaultInjectionPolicy
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteAction
	SetInternalValue(val *GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteAction)
	RequestMirrorPolicy() GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionRequestMirrorPolicyOutputReference
	RequestMirrorPolicyInput() *GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionRequestMirrorPolicy
	RetryPolicy() GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionRetryPolicyOutputReference
	RetryPolicyInput() *GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionRetryPolicy
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Timeout() GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionTimeoutOutputReference
	TimeoutInput() *GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionTimeout
	UrlRewrite() GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionUrlRewriteOutputReference
	UrlRewriteInput() *GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionUrlRewrite
	WeightedBackendServices() GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionWeightedBackendServicesList
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
	PutCorsPolicy(value *GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionCorsPolicy)
	PutFaultInjectionPolicy(value *GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionFaultInjectionPolicy)
	PutRequestMirrorPolicy(value *GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionRequestMirrorPolicy)
	PutRetryPolicy(value *GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionRetryPolicy)
	PutTimeout(value *GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionTimeout)
	PutUrlRewrite(value *GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionUrlRewrite)
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

// The jsii proxy struct for GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionOutputReference
type jsiiProxy_GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionOutputReference) CorsPolicy() GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionCorsPolicyOutputReference {
	var returns GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionCorsPolicyOutputReference
	_jsii_.Get(
		j,
		"corsPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionOutputReference) CorsPolicyInput() *GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionCorsPolicy {
	var returns *GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionCorsPolicy
	_jsii_.Get(
		j,
		"corsPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionOutputReference) FaultInjectionPolicy() GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionFaultInjectionPolicyOutputReference {
	var returns GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionFaultInjectionPolicyOutputReference
	_jsii_.Get(
		j,
		"faultInjectionPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionOutputReference) FaultInjectionPolicyInput() *GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionFaultInjectionPolicy {
	var returns *GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionFaultInjectionPolicy
	_jsii_.Get(
		j,
		"faultInjectionPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionOutputReference) InternalValue() *GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteAction {
	var returns *GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteAction
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionOutputReference) RequestMirrorPolicy() GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionRequestMirrorPolicyOutputReference {
	var returns GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionRequestMirrorPolicyOutputReference
	_jsii_.Get(
		j,
		"requestMirrorPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionOutputReference) RequestMirrorPolicyInput() *GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionRequestMirrorPolicy {
	var returns *GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionRequestMirrorPolicy
	_jsii_.Get(
		j,
		"requestMirrorPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionOutputReference) RetryPolicy() GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionRetryPolicyOutputReference {
	var returns GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionRetryPolicyOutputReference
	_jsii_.Get(
		j,
		"retryPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionOutputReference) RetryPolicyInput() *GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionRetryPolicy {
	var returns *GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionRetryPolicy
	_jsii_.Get(
		j,
		"retryPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionOutputReference) Timeout() GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionTimeoutOutputReference {
	var returns GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionTimeoutOutputReference
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionOutputReference) TimeoutInput() *GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionTimeout {
	var returns *GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionTimeout
	_jsii_.Get(
		j,
		"timeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionOutputReference) UrlRewrite() GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionUrlRewriteOutputReference {
	var returns GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionUrlRewriteOutputReference
	_jsii_.Get(
		j,
		"urlRewrite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionOutputReference) UrlRewriteInput() *GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionUrlRewrite {
	var returns *GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionUrlRewrite
	_jsii_.Get(
		j,
		"urlRewriteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionOutputReference) WeightedBackendServices() GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionWeightedBackendServicesList {
	var returns GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionWeightedBackendServicesList
	_jsii_.Get(
		j,
		"weightedBackendServices",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionOutputReference) WeightedBackendServicesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"weightedBackendServicesInput",
		&returns,
	)
	return returns
}


func NewGoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionOutputReference{}

	_jsii_.Create(
		"google-beta.googleComputeRegionUrlMap.GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionOutputReference_Override(g GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleComputeRegionUrlMap.GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionOutputReference)SetInternalValue(val *GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteAction) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionOutputReference) PutCorsPolicy(value *GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionCorsPolicy) {
	if err := g.validatePutCorsPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putCorsPolicy",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionOutputReference) PutFaultInjectionPolicy(value *GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionFaultInjectionPolicy) {
	if err := g.validatePutFaultInjectionPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putFaultInjectionPolicy",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionOutputReference) PutRequestMirrorPolicy(value *GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionRequestMirrorPolicy) {
	if err := g.validatePutRequestMirrorPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRequestMirrorPolicy",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionOutputReference) PutRetryPolicy(value *GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionRetryPolicy) {
	if err := g.validatePutRetryPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRetryPolicy",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionOutputReference) PutTimeout(value *GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionTimeout) {
	if err := g.validatePutTimeoutParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeout",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionOutputReference) PutUrlRewrite(value *GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionUrlRewrite) {
	if err := g.validatePutUrlRewriteParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putUrlRewrite",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionOutputReference) PutWeightedBackendServices(value interface{}) {
	if err := g.validatePutWeightedBackendServicesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putWeightedBackendServices",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionOutputReference) ResetCorsPolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetCorsPolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionOutputReference) ResetFaultInjectionPolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetFaultInjectionPolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionOutputReference) ResetRequestMirrorPolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetRequestMirrorPolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionOutputReference) ResetRetryPolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetRetryPolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionOutputReference) ResetTimeout() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeout",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionOutputReference) ResetUrlRewrite() {
	_jsii_.InvokeVoid(
		g,
		"resetUrlRewrite",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionOutputReference) ResetWeightedBackendServices() {
	_jsii_.InvokeVoid(
		g,
		"resetWeightedBackendServices",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherRouteRulesRouteActionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

