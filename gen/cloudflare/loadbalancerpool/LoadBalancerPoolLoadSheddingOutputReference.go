package loadbalancerpool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/cloudflare/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/cloudflare/loadbalancerpool/internal"
)

type LoadBalancerPoolLoadSheddingOutputReference interface {
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
	DefaultPercent() *float64
	SetDefaultPercent(val *float64)
	DefaultPercentInput() *float64
	DefaultPolicy() *string
	SetDefaultPolicy(val *string)
	DefaultPolicyInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	SessionPercent() *float64
	SetSessionPercent(val *float64)
	SessionPercentInput() *float64
	SessionPolicy() *string
	SetSessionPolicy(val *string)
	SessionPolicyInput() *string
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
	ResetDefaultPercent()
	ResetDefaultPolicy()
	ResetSessionPercent()
	ResetSessionPolicy()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LoadBalancerPoolLoadSheddingOutputReference
type jsiiProxy_LoadBalancerPoolLoadSheddingOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LoadBalancerPoolLoadSheddingOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerPoolLoadSheddingOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerPoolLoadSheddingOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerPoolLoadSheddingOutputReference) DefaultPercent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultPercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerPoolLoadSheddingOutputReference) DefaultPercentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultPercentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerPoolLoadSheddingOutputReference) DefaultPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerPoolLoadSheddingOutputReference) DefaultPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerPoolLoadSheddingOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerPoolLoadSheddingOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerPoolLoadSheddingOutputReference) SessionPercent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sessionPercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerPoolLoadSheddingOutputReference) SessionPercentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sessionPercentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerPoolLoadSheddingOutputReference) SessionPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sessionPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerPoolLoadSheddingOutputReference) SessionPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sessionPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerPoolLoadSheddingOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerPoolLoadSheddingOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewLoadBalancerPoolLoadSheddingOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) LoadBalancerPoolLoadSheddingOutputReference {
	_init_.Initialize()

	if err := validateNewLoadBalancerPoolLoadSheddingOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_LoadBalancerPoolLoadSheddingOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.loadBalancerPool.LoadBalancerPoolLoadSheddingOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewLoadBalancerPoolLoadSheddingOutputReference_Override(l LoadBalancerPoolLoadSheddingOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.loadBalancerPool.LoadBalancerPoolLoadSheddingOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		l,
	)
}

func (j *jsiiProxy_LoadBalancerPoolLoadSheddingOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LoadBalancerPoolLoadSheddingOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LoadBalancerPoolLoadSheddingOutputReference)SetDefaultPercent(val *float64) {
	if err := j.validateSetDefaultPercentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultPercent",
		val,
	)
}

func (j *jsiiProxy_LoadBalancerPoolLoadSheddingOutputReference)SetDefaultPolicy(val *string) {
	if err := j.validateSetDefaultPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultPolicy",
		val,
	)
}

func (j *jsiiProxy_LoadBalancerPoolLoadSheddingOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LoadBalancerPoolLoadSheddingOutputReference)SetSessionPercent(val *float64) {
	if err := j.validateSetSessionPercentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sessionPercent",
		val,
	)
}

func (j *jsiiProxy_LoadBalancerPoolLoadSheddingOutputReference)SetSessionPolicy(val *string) {
	if err := j.validateSetSessionPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sessionPolicy",
		val,
	)
}

func (j *jsiiProxy_LoadBalancerPoolLoadSheddingOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LoadBalancerPoolLoadSheddingOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (l *jsiiProxy_LoadBalancerPoolLoadSheddingOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadBalancerPoolLoadSheddingOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (l *jsiiProxy_LoadBalancerPoolLoadSheddingOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LoadBalancerPoolLoadSheddingOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (l *jsiiProxy_LoadBalancerPoolLoadSheddingOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (l *jsiiProxy_LoadBalancerPoolLoadSheddingOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (l *jsiiProxy_LoadBalancerPoolLoadSheddingOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (l *jsiiProxy_LoadBalancerPoolLoadSheddingOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (l *jsiiProxy_LoadBalancerPoolLoadSheddingOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (l *jsiiProxy_LoadBalancerPoolLoadSheddingOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (l *jsiiProxy_LoadBalancerPoolLoadSheddingOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadBalancerPoolLoadSheddingOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LoadBalancerPoolLoadSheddingOutputReference) ResetDefaultPercent() {
	_jsii_.InvokeVoid(
		l,
		"resetDefaultPercent",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadBalancerPoolLoadSheddingOutputReference) ResetDefaultPolicy() {
	_jsii_.InvokeVoid(
		l,
		"resetDefaultPolicy",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadBalancerPoolLoadSheddingOutputReference) ResetSessionPercent() {
	_jsii_.InvokeVoid(
		l,
		"resetSessionPercent",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadBalancerPoolLoadSheddingOutputReference) ResetSessionPolicy() {
	_jsii_.InvokeVoid(
		l,
		"resetSessionPolicy",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadBalancerPoolLoadSheddingOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (l *jsiiProxy_LoadBalancerPoolLoadSheddingOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

