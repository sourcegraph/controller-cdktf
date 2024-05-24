package loadbalancer

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/cloudflare/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/cloudflare/loadbalancer/internal"
)

type LoadBalancerSessionAffinityAttributesOutputReference interface {
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
	DrainDuration() *float64
	SetDrainDuration(val *float64)
	DrainDurationInput() *float64
	// Experimental.
	Fqn() *string
	Headers() *[]*string
	SetHeaders(val *[]*string)
	HeadersInput() *[]*string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	RequireAllHeaders() interface{}
	SetRequireAllHeaders(val interface{})
	RequireAllHeadersInput() interface{}
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
	ResetDrainDuration()
	ResetHeaders()
	ResetRequireAllHeaders()
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

// The jsii proxy struct for LoadBalancerSessionAffinityAttributesOutputReference
type jsiiProxy_LoadBalancerSessionAffinityAttributesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LoadBalancerSessionAffinityAttributesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerSessionAffinityAttributesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerSessionAffinityAttributesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerSessionAffinityAttributesOutputReference) DrainDuration() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"drainDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerSessionAffinityAttributesOutputReference) DrainDurationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"drainDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerSessionAffinityAttributesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerSessionAffinityAttributesOutputReference) Headers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"headers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerSessionAffinityAttributesOutputReference) HeadersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"headersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerSessionAffinityAttributesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerSessionAffinityAttributesOutputReference) RequireAllHeaders() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireAllHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerSessionAffinityAttributesOutputReference) RequireAllHeadersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireAllHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerSessionAffinityAttributesOutputReference) Samesite() *string {
	var returns *string
	_jsii_.Get(
		j,
		"samesite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerSessionAffinityAttributesOutputReference) SamesiteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"samesiteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerSessionAffinityAttributesOutputReference) Secure() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerSessionAffinityAttributesOutputReference) SecureInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerSessionAffinityAttributesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerSessionAffinityAttributesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerSessionAffinityAttributesOutputReference) ZeroDowntimeFailover() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zeroDowntimeFailover",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancerSessionAffinityAttributesOutputReference) ZeroDowntimeFailoverInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zeroDowntimeFailoverInput",
		&returns,
	)
	return returns
}


func NewLoadBalancerSessionAffinityAttributesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) LoadBalancerSessionAffinityAttributesOutputReference {
	_init_.Initialize()

	if err := validateNewLoadBalancerSessionAffinityAttributesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_LoadBalancerSessionAffinityAttributesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.loadBalancer.LoadBalancerSessionAffinityAttributesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewLoadBalancerSessionAffinityAttributesOutputReference_Override(l LoadBalancerSessionAffinityAttributesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.loadBalancer.LoadBalancerSessionAffinityAttributesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		l,
	)
}

func (j *jsiiProxy_LoadBalancerSessionAffinityAttributesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LoadBalancerSessionAffinityAttributesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LoadBalancerSessionAffinityAttributesOutputReference)SetDrainDuration(val *float64) {
	if err := j.validateSetDrainDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"drainDuration",
		val,
	)
}

func (j *jsiiProxy_LoadBalancerSessionAffinityAttributesOutputReference)SetHeaders(val *[]*string) {
	if err := j.validateSetHeadersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"headers",
		val,
	)
}

func (j *jsiiProxy_LoadBalancerSessionAffinityAttributesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LoadBalancerSessionAffinityAttributesOutputReference)SetRequireAllHeaders(val interface{}) {
	if err := j.validateSetRequireAllHeadersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireAllHeaders",
		val,
	)
}

func (j *jsiiProxy_LoadBalancerSessionAffinityAttributesOutputReference)SetSamesite(val *string) {
	if err := j.validateSetSamesiteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"samesite",
		val,
	)
}

func (j *jsiiProxy_LoadBalancerSessionAffinityAttributesOutputReference)SetSecure(val *string) {
	if err := j.validateSetSecureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secure",
		val,
	)
}

func (j *jsiiProxy_LoadBalancerSessionAffinityAttributesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LoadBalancerSessionAffinityAttributesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_LoadBalancerSessionAffinityAttributesOutputReference)SetZeroDowntimeFailover(val *string) {
	if err := j.validateSetZeroDowntimeFailoverParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zeroDowntimeFailover",
		val,
	)
}

func (l *jsiiProxy_LoadBalancerSessionAffinityAttributesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadBalancerSessionAffinityAttributesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (l *jsiiProxy_LoadBalancerSessionAffinityAttributesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LoadBalancerSessionAffinityAttributesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (l *jsiiProxy_LoadBalancerSessionAffinityAttributesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (l *jsiiProxy_LoadBalancerSessionAffinityAttributesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (l *jsiiProxy_LoadBalancerSessionAffinityAttributesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (l *jsiiProxy_LoadBalancerSessionAffinityAttributesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (l *jsiiProxy_LoadBalancerSessionAffinityAttributesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (l *jsiiProxy_LoadBalancerSessionAffinityAttributesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (l *jsiiProxy_LoadBalancerSessionAffinityAttributesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadBalancerSessionAffinityAttributesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LoadBalancerSessionAffinityAttributesOutputReference) ResetDrainDuration() {
	_jsii_.InvokeVoid(
		l,
		"resetDrainDuration",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadBalancerSessionAffinityAttributesOutputReference) ResetHeaders() {
	_jsii_.InvokeVoid(
		l,
		"resetHeaders",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadBalancerSessionAffinityAttributesOutputReference) ResetRequireAllHeaders() {
	_jsii_.InvokeVoid(
		l,
		"resetRequireAllHeaders",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadBalancerSessionAffinityAttributesOutputReference) ResetSamesite() {
	_jsii_.InvokeVoid(
		l,
		"resetSamesite",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadBalancerSessionAffinityAttributesOutputReference) ResetSecure() {
	_jsii_.InvokeVoid(
		l,
		"resetSecure",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadBalancerSessionAffinityAttributesOutputReference) ResetZeroDowntimeFailover() {
	_jsii_.InvokeVoid(
		l,
		"resetZeroDowntimeFailover",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadBalancerSessionAffinityAttributesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (l *jsiiProxy_LoadBalancerSessionAffinityAttributesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

