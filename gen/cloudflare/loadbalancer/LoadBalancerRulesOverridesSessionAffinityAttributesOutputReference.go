package loadbalancer

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/cloudflare/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/cloudflare/loadbalancer/internal"
)

type LoadBalancerRulesOverridesSessionAffinityAttributesOutputReference interface {
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Samesite() *string
	SetSamesite(val *string)
	SamesiteInput() *string
	Secure() *string
	SetSecure(val *string)
	SecureInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	ZeroDowntimeFailover() *string
	SetZeroDowntimeFailover(val *string)
	ZeroDowntimeFailoverInput() *string
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
	ResetSamesite()
	ResetSecure()
	ResetZeroDowntimeFailover()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LoadBalancerRulesOverridesSessionAffinityAttributesOutputReference
type jsiiProxy_LoadBalancerRulesOverridesSessionAffinityAttributesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LoadBalancerRulesOverridesSessionAffinityAttributesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerRulesOverridesSessionAffinityAttributesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerRulesOverridesSessionAffinityAttributesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerRulesOverridesSessionAffinityAttributesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerRulesOverridesSessionAffinityAttributesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerRulesOverridesSessionAffinityAttributesOutputReference) Samesite() *string {
	var returns *string
	_jsii_.Get(
		j,
		"samesite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerRulesOverridesSessionAffinityAttributesOutputReference) SamesiteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"samesiteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerRulesOverridesSessionAffinityAttributesOutputReference) Secure() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerRulesOverridesSessionAffinityAttributesOutputReference) SecureInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerRulesOverridesSessionAffinityAttributesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerRulesOverridesSessionAffinityAttributesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerRulesOverridesSessionAffinityAttributesOutputReference) ZeroDowntimeFailover() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zeroDowntimeFailover",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerRulesOverridesSessionAffinityAttributesOutputReference) ZeroDowntimeFailoverInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zeroDowntimeFailoverInput",
		&returns,
	)
	return returns
}


func NewLoadBalancerRulesOverridesSessionAffinityAttributesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) LoadBalancerRulesOverridesSessionAffinityAttributesOutputReference {
	_init_.Initialize()

	if err := validateNewLoadBalancerRulesOverridesSessionAffinityAttributesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_LoadBalancerRulesOverridesSessionAffinityAttributesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.loadBalancer.LoadBalancerRulesOverridesSessionAffinityAttributesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewLoadBalancerRulesOverridesSessionAffinityAttributesOutputReference_Override(l LoadBalancerRulesOverridesSessionAffinityAttributesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.loadBalancer.LoadBalancerRulesOverridesSessionAffinityAttributesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		l,
	)
}

func (j *jsiiProxy_LoadBalancerRulesOverridesSessionAffinityAttributesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LoadBalancerRulesOverridesSessionAffinityAttributesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LoadBalancerRulesOverridesSessionAffinityAttributesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LoadBalancerRulesOverridesSessionAffinityAttributesOutputReference)SetSamesite(val *string) {
	if err := j.validateSetSamesiteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"samesite",
		val,
	)
}

func (j *jsiiProxy_LoadBalancerRulesOverridesSessionAffinityAttributesOutputReference)SetSecure(val *string) {
	if err := j.validateSetSecureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secure",
		val,
	)
}

func (j *jsiiProxy_LoadBalancerRulesOverridesSessionAffinityAttributesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LoadBalancerRulesOverridesSessionAffinityAttributesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_LoadBalancerRulesOverridesSessionAffinityAttributesOutputReference)SetZeroDowntimeFailover(val *string) {
	if err := j.validateSetZeroDowntimeFailoverParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zeroDowntimeFailover",
		val,
	)
}

func (l *jsiiProxy_LoadBalancerRulesOverridesSessionAffinityAttributesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadBalancerRulesOverridesSessionAffinityAttributesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (l *jsiiProxy_LoadBalancerRulesOverridesSessionAffinityAttributesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LoadBalancerRulesOverridesSessionAffinityAttributesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (l *jsiiProxy_LoadBalancerRulesOverridesSessionAffinityAttributesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (l *jsiiProxy_LoadBalancerRulesOverridesSessionAffinityAttributesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (l *jsiiProxy_LoadBalancerRulesOverridesSessionAffinityAttributesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (l *jsiiProxy_LoadBalancerRulesOverridesSessionAffinityAttributesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (l *jsiiProxy_LoadBalancerRulesOverridesSessionAffinityAttributesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (l *jsiiProxy_LoadBalancerRulesOverridesSessionAffinityAttributesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (l *jsiiProxy_LoadBalancerRulesOverridesSessionAffinityAttributesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadBalancerRulesOverridesSessionAffinityAttributesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LoadBalancerRulesOverridesSessionAffinityAttributesOutputReference) ResetSamesite() {
	_jsii_.InvokeVoid(
		l,
		"resetSamesite",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadBalancerRulesOverridesSessionAffinityAttributesOutputReference) ResetSecure() {
	_jsii_.InvokeVoid(
		l,
		"resetSecure",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadBalancerRulesOverridesSessionAffinityAttributesOutputReference) ResetZeroDowntimeFailover() {
	_jsii_.InvokeVoid(
		l,
		"resetZeroDowntimeFailover",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadBalancerRulesOverridesSessionAffinityAttributesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (l *jsiiProxy_LoadBalancerRulesOverridesSessionAffinityAttributesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

