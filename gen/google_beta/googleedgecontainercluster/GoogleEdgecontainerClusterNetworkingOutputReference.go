package googleedgecontainercluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleedgecontainercluster/internal"
)

type GoogleEdgecontainerClusterNetworkingOutputReference interface {
	cdktf.ComplexObject
	ClusterIpv4CidrBlocks() *[]*string
	SetClusterIpv4CidrBlocks(val *[]*string)
	ClusterIpv4CidrBlocksInput() *[]*string
	ClusterIpv6CidrBlocks() *[]*string
	SetClusterIpv6CidrBlocks(val *[]*string)
	ClusterIpv6CidrBlocksInput() *[]*string
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
	InternalValue() *GoogleEdgecontainerClusterNetworking
	SetInternalValue(val *GoogleEdgecontainerClusterNetworking)
	NetworkType() *string
	ServicesIpv4CidrBlocks() *[]*string
	SetServicesIpv4CidrBlocks(val *[]*string)
	ServicesIpv4CidrBlocksInput() *[]*string
	ServicesIpv6CidrBlocks() *[]*string
	SetServicesIpv6CidrBlocks(val *[]*string)
	ServicesIpv6CidrBlocksInput() *[]*string
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
	ResetClusterIpv6CidrBlocks()
	ResetServicesIpv6CidrBlocks()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleEdgecontainerClusterNetworkingOutputReference
type jsiiProxy_GoogleEdgecontainerClusterNetworkingOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleEdgecontainerClusterNetworkingOutputReference) ClusterIpv4CidrBlocks() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"clusterIpv4CidrBlocks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEdgecontainerClusterNetworkingOutputReference) ClusterIpv4CidrBlocksInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"clusterIpv4CidrBlocksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEdgecontainerClusterNetworkingOutputReference) ClusterIpv6CidrBlocks() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"clusterIpv6CidrBlocks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEdgecontainerClusterNetworkingOutputReference) ClusterIpv6CidrBlocksInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"clusterIpv6CidrBlocksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEdgecontainerClusterNetworkingOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEdgecontainerClusterNetworkingOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEdgecontainerClusterNetworkingOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEdgecontainerClusterNetworkingOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEdgecontainerClusterNetworkingOutputReference) InternalValue() *GoogleEdgecontainerClusterNetworking {
	var returns *GoogleEdgecontainerClusterNetworking
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEdgecontainerClusterNetworkingOutputReference) NetworkType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEdgecontainerClusterNetworkingOutputReference) ServicesIpv4CidrBlocks() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"servicesIpv4CidrBlocks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEdgecontainerClusterNetworkingOutputReference) ServicesIpv4CidrBlocksInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"servicesIpv4CidrBlocksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEdgecontainerClusterNetworkingOutputReference) ServicesIpv6CidrBlocks() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"servicesIpv6CidrBlocks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEdgecontainerClusterNetworkingOutputReference) ServicesIpv6CidrBlocksInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"servicesIpv6CidrBlocksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEdgecontainerClusterNetworkingOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEdgecontainerClusterNetworkingOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleEdgecontainerClusterNetworkingOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleEdgecontainerClusterNetworkingOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleEdgecontainerClusterNetworkingOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleEdgecontainerClusterNetworkingOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleEdgecontainerCluster.GoogleEdgecontainerClusterNetworkingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleEdgecontainerClusterNetworkingOutputReference_Override(g GoogleEdgecontainerClusterNetworkingOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleEdgecontainerCluster.GoogleEdgecontainerClusterNetworkingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleEdgecontainerClusterNetworkingOutputReference)SetClusterIpv4CidrBlocks(val *[]*string) {
	if err := j.validateSetClusterIpv4CidrBlocksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterIpv4CidrBlocks",
		val,
	)
}

func (j *jsiiProxy_GoogleEdgecontainerClusterNetworkingOutputReference)SetClusterIpv6CidrBlocks(val *[]*string) {
	if err := j.validateSetClusterIpv6CidrBlocksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterIpv6CidrBlocks",
		val,
	)
}

func (j *jsiiProxy_GoogleEdgecontainerClusterNetworkingOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleEdgecontainerClusterNetworkingOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleEdgecontainerClusterNetworkingOutputReference)SetInternalValue(val *GoogleEdgecontainerClusterNetworking) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleEdgecontainerClusterNetworkingOutputReference)SetServicesIpv4CidrBlocks(val *[]*string) {
	if err := j.validateSetServicesIpv4CidrBlocksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"servicesIpv4CidrBlocks",
		val,
	)
}

func (j *jsiiProxy_GoogleEdgecontainerClusterNetworkingOutputReference)SetServicesIpv6CidrBlocks(val *[]*string) {
	if err := j.validateSetServicesIpv6CidrBlocksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"servicesIpv6CidrBlocks",
		val,
	)
}

func (j *jsiiProxy_GoogleEdgecontainerClusterNetworkingOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleEdgecontainerClusterNetworkingOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleEdgecontainerClusterNetworkingOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleEdgecontainerClusterNetworkingOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleEdgecontainerClusterNetworkingOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleEdgecontainerClusterNetworkingOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleEdgecontainerClusterNetworkingOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleEdgecontainerClusterNetworkingOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleEdgecontainerClusterNetworkingOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleEdgecontainerClusterNetworkingOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleEdgecontainerClusterNetworkingOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleEdgecontainerClusterNetworkingOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleEdgecontainerClusterNetworkingOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleEdgecontainerClusterNetworkingOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleEdgecontainerClusterNetworkingOutputReference) ResetClusterIpv6CidrBlocks() {
	_jsii_.InvokeVoid(
		g,
		"resetClusterIpv6CidrBlocks",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleEdgecontainerClusterNetworkingOutputReference) ResetServicesIpv6CidrBlocks() {
	_jsii_.InvokeVoid(
		g,
		"resetServicesIpv6CidrBlocks",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleEdgecontainerClusterNetworkingOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleEdgecontainerClusterNetworkingOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

