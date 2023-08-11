package googlegkeonprembaremetalcluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlegkeonprembaremetalcluster/internal"
)

type GoogleGkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference interface {
	cdktf.ComplexObject
	AddressPools() GoogleGkeonpremBareMetalClusterLoadBalancerBgpLbConfigAddressPoolsList
	AddressPoolsInput() interface{}
	Asn() *float64
	SetAsn(val *float64)
	AsnInput() *float64
	BgpPeerConfigs() GoogleGkeonpremBareMetalClusterLoadBalancerBgpLbConfigBgpPeerConfigsList
	BgpPeerConfigsInput() interface{}
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
	InternalValue() *GoogleGkeonpremBareMetalClusterLoadBalancerBgpLbConfig
	SetInternalValue(val *GoogleGkeonpremBareMetalClusterLoadBalancerBgpLbConfig)
	LoadBalancerNodePoolConfig() GoogleGkeonpremBareMetalClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfigOutputReference
	LoadBalancerNodePoolConfigInput() *GoogleGkeonpremBareMetalClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfig
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
	PutAddressPools(value interface{})
	PutBgpPeerConfigs(value interface{})
	PutLoadBalancerNodePoolConfig(value *GoogleGkeonpremBareMetalClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfig)
	ResetLoadBalancerNodePoolConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleGkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference
type jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference) AddressPools() GoogleGkeonpremBareMetalClusterLoadBalancerBgpLbConfigAddressPoolsList {
	var returns GoogleGkeonpremBareMetalClusterLoadBalancerBgpLbConfigAddressPoolsList
	_jsii_.Get(
		j,
		"addressPools",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference) AddressPoolsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"addressPoolsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference) Asn() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"asn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference) AsnInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"asnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference) BgpPeerConfigs() GoogleGkeonpremBareMetalClusterLoadBalancerBgpLbConfigBgpPeerConfigsList {
	var returns GoogleGkeonpremBareMetalClusterLoadBalancerBgpLbConfigBgpPeerConfigsList
	_jsii_.Get(
		j,
		"bgpPeerConfigs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference) BgpPeerConfigsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bgpPeerConfigsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference) InternalValue() *GoogleGkeonpremBareMetalClusterLoadBalancerBgpLbConfig {
	var returns *GoogleGkeonpremBareMetalClusterLoadBalancerBgpLbConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference) LoadBalancerNodePoolConfig() GoogleGkeonpremBareMetalClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfigOutputReference {
	var returns GoogleGkeonpremBareMetalClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfigOutputReference
	_jsii_.Get(
		j,
		"loadBalancerNodePoolConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference) LoadBalancerNodePoolConfigInput() *GoogleGkeonpremBareMetalClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfig {
	var returns *GoogleGkeonpremBareMetalClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfig
	_jsii_.Get(
		j,
		"loadBalancerNodePoolConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleGkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleGkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleGkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleGkeonpremBareMetalCluster.GoogleGkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleGkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference_Override(g GoogleGkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleGkeonpremBareMetalCluster.GoogleGkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference)SetAsn(val *float64) {
	if err := j.validateSetAsnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"asn",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference)SetInternalValue(val *GoogleGkeonpremBareMetalClusterLoadBalancerBgpLbConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference) PutAddressPools(value interface{}) {
	if err := g.validatePutAddressPoolsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAddressPools",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference) PutBgpPeerConfigs(value interface{}) {
	if err := g.validatePutBgpPeerConfigsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putBgpPeerConfigs",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference) PutLoadBalancerNodePoolConfig(value *GoogleGkeonpremBareMetalClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfig) {
	if err := g.validatePutLoadBalancerNodePoolConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putLoadBalancerNodePoolConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference) ResetLoadBalancerNodePoolConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetLoadBalancerNodePoolConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleGkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

