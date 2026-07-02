package gkeonprembaremetaladmincluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/gkeonprembaremetaladmincluster/internal"
)

type GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigOutputReference interface {
	cdktf.ComplexObject
	AddressPools() GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigAddressPoolsList
	AddressPoolsInput() any
	Asn() *float64
	SetAsn(val *float64)
	AsnInput() *float64
	BgpPeerConfigs() GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigBgpPeerConfigsList
	BgpPeerConfigsInput() any
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
	// Experimental.
	Fqn() *string
	InternalValue() *GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfig
	SetInternalValue(val *GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfig)
	LoadBalancerNodePoolConfig() GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfigOutputReference
	LoadBalancerNodePoolConfigInput() *GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfig
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
	PutAddressPools(value any)
	PutBgpPeerConfigs(value any)
	PutLoadBalancerNodePoolConfig(value *GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfig)
	ResetAddressPools()
	ResetAsn()
	ResetBgpPeerConfigs()
	ResetLoadBalancerNodePoolConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigOutputReference
type jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigOutputReference) AddressPools() GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigAddressPoolsList {
	var returns GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigAddressPoolsList
	_jsii_.Get(
		j,
		"addressPools",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigOutputReference) AddressPoolsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"addressPoolsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigOutputReference) Asn() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"asn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigOutputReference) AsnInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"asnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigOutputReference) BgpPeerConfigs() GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigBgpPeerConfigsList {
	var returns GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigBgpPeerConfigsList
	_jsii_.Get(
		j,
		"bgpPeerConfigs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigOutputReference) BgpPeerConfigsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"bgpPeerConfigsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigOutputReference) InternalValue() *GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfig {
	var returns *GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigOutputReference) LoadBalancerNodePoolConfig() GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfigOutputReference {
	var returns GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfigOutputReference
	_jsii_.Get(
		j,
		"loadBalancerNodePoolConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigOutputReference) LoadBalancerNodePoolConfigInput() *GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfig {
	var returns *GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfig
	_jsii_.Get(
		j,
		"loadBalancerNodePoolConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewGkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.gkeonpremBareMetalAdminCluster.GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigOutputReference_Override(g GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.gkeonpremBareMetalAdminCluster.GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigOutputReference",
		[]any{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigOutputReference) SetAsn(val *float64) {
	if err := j.validateSetAsnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"asn",
		val,
	)
}

func (j *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigOutputReference) SetInternalValue(val *GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigOutputReference) PutAddressPools(value any) {
	if err := g.validatePutAddressPoolsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAddressPools",
		[]any{value},
	)
}

func (g *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigOutputReference) PutBgpPeerConfigs(value any) {
	if err := g.validatePutBgpPeerConfigsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putBgpPeerConfigs",
		[]any{value},
	)
}

func (g *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigOutputReference) PutLoadBalancerNodePoolConfig(value *GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfig) {
	if err := g.validatePutLoadBalancerNodePoolConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putLoadBalancerNodePoolConfig",
		[]any{value},
	)
}

func (g *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigOutputReference) ResetAddressPools() {
	_jsii_.InvokeVoid(
		g,
		"resetAddressPools",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigOutputReference) ResetAsn() {
	_jsii_.InvokeVoid(
		g,
		"resetAsn",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigOutputReference) ResetBgpPeerConfigs() {
	_jsii_.InvokeVoid(
		g,
		"resetBgpPeerConfigs",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigOutputReference) ResetLoadBalancerNodePoolConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetLoadBalancerNodePoolConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := g.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		g,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremBareMetalAdminClusterLoadBalancerBgpLbConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
