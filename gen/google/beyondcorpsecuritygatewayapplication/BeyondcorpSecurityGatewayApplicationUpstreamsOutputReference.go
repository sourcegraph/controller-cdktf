package beyondcorpsecuritygatewayapplication

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/beyondcorpsecuritygatewayapplication/internal"
)

type BeyondcorpSecurityGatewayApplicationUpstreamsOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() any
	// Experimental.
	SetComplexObjectIndex(val any)
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
	EgressPolicy() BeyondcorpSecurityGatewayApplicationUpstreamsEgressPolicyOutputReference
	EgressPolicyInput() *BeyondcorpSecurityGatewayApplicationUpstreamsEgressPolicy
	External() BeyondcorpSecurityGatewayApplicationUpstreamsExternalOutputReference
	ExternalInput() *BeyondcorpSecurityGatewayApplicationUpstreamsExternal
	// Experimental.
	Fqn() *string
	InternalValue() any
	SetInternalValue(val any)
	Network() BeyondcorpSecurityGatewayApplicationUpstreamsNetworkOutputReference
	NetworkInput() *BeyondcorpSecurityGatewayApplicationUpstreamsNetwork
	ProxyProtocol() BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference
	ProxyProtocolInput() *BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocol
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
	GetAnyMapAttribute(terraformAttribute *string) *map[string]any
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
	PutEgressPolicy(value *BeyondcorpSecurityGatewayApplicationUpstreamsEgressPolicy)
	PutExternal(value *BeyondcorpSecurityGatewayApplicationUpstreamsExternal)
	PutNetwork(value *BeyondcorpSecurityGatewayApplicationUpstreamsNetwork)
	PutProxyProtocol(value *BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocol)
	ResetEgressPolicy()
	ResetExternal()
	ResetNetwork()
	ResetProxyProtocol()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BeyondcorpSecurityGatewayApplicationUpstreamsOutputReference
type jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsOutputReference) EgressPolicy() BeyondcorpSecurityGatewayApplicationUpstreamsEgressPolicyOutputReference {
	var returns BeyondcorpSecurityGatewayApplicationUpstreamsEgressPolicyOutputReference
	_jsii_.Get(
		j,
		"egressPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsOutputReference) EgressPolicyInput() *BeyondcorpSecurityGatewayApplicationUpstreamsEgressPolicy {
	var returns *BeyondcorpSecurityGatewayApplicationUpstreamsEgressPolicy
	_jsii_.Get(
		j,
		"egressPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsOutputReference) External() BeyondcorpSecurityGatewayApplicationUpstreamsExternalOutputReference {
	var returns BeyondcorpSecurityGatewayApplicationUpstreamsExternalOutputReference
	_jsii_.Get(
		j,
		"external",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsOutputReference) ExternalInput() *BeyondcorpSecurityGatewayApplicationUpstreamsExternal {
	var returns *BeyondcorpSecurityGatewayApplicationUpstreamsExternal
	_jsii_.Get(
		j,
		"externalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsOutputReference) InternalValue() any {
	var returns any
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsOutputReference) Network() BeyondcorpSecurityGatewayApplicationUpstreamsNetworkOutputReference {
	var returns BeyondcorpSecurityGatewayApplicationUpstreamsNetworkOutputReference
	_jsii_.Get(
		j,
		"network",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsOutputReference) NetworkInput() *BeyondcorpSecurityGatewayApplicationUpstreamsNetwork {
	var returns *BeyondcorpSecurityGatewayApplicationUpstreamsNetwork
	_jsii_.Get(
		j,
		"networkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsOutputReference) ProxyProtocol() BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference {
	var returns BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference
	_jsii_.Get(
		j,
		"proxyProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsOutputReference) ProxyProtocolInput() *BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocol {
	var returns *BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocol
	_jsii_.Get(
		j,
		"proxyProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewBeyondcorpSecurityGatewayApplicationUpstreamsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) BeyondcorpSecurityGatewayApplicationUpstreamsOutputReference {
	_init_.Initialize()

	if err := validateNewBeyondcorpSecurityGatewayApplicationUpstreamsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.beyondcorpSecurityGatewayApplication.BeyondcorpSecurityGatewayApplicationUpstreamsOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewBeyondcorpSecurityGatewayApplicationUpstreamsOutputReference_Override(b BeyondcorpSecurityGatewayApplicationUpstreamsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.beyondcorpSecurityGatewayApplication.BeyondcorpSecurityGatewayApplicationUpstreamsOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		b,
	)
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsOutputReference) SetInternalValue(val any) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := b.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := b.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := b.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := b.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := b.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := b.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := b.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := b.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := b.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsOutputReference) PutEgressPolicy(value *BeyondcorpSecurityGatewayApplicationUpstreamsEgressPolicy) {
	if err := b.validatePutEgressPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putEgressPolicy",
		[]any{value},
	)
}

func (b *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsOutputReference) PutExternal(value *BeyondcorpSecurityGatewayApplicationUpstreamsExternal) {
	if err := b.validatePutExternalParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putExternal",
		[]any{value},
	)
}

func (b *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsOutputReference) PutNetwork(value *BeyondcorpSecurityGatewayApplicationUpstreamsNetwork) {
	if err := b.validatePutNetworkParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putNetwork",
		[]any{value},
	)
}

func (b *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsOutputReference) PutProxyProtocol(value *BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocol) {
	if err := b.validatePutProxyProtocolParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putProxyProtocol",
		[]any{value},
	)
}

func (b *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsOutputReference) ResetEgressPolicy() {
	_jsii_.InvokeVoid(
		b,
		"resetEgressPolicy",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsOutputReference) ResetExternal() {
	_jsii_.InvokeVoid(
		b,
		"resetExternal",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsOutputReference) ResetNetwork() {
	_jsii_.InvokeVoid(
		b,
		"resetNetwork",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsOutputReference) ResetProxyProtocol() {
	_jsii_.InvokeVoid(
		b,
		"resetProxyProtocol",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := b.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		b,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
