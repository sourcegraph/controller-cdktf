package googlegkeonpremvmwareadmincluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlegkeonpremvmwareadmincluster/internal"
)

type GoogleGkeonpremVmwareAdminClusterNetworkConfigHaControlPlaneConfigControlPlaneIpBlockOutputReference interface {
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
	Gateway() *string
	SetGateway(val *string)
	GatewayInput() *string
	InternalValue() *GoogleGkeonpremVmwareAdminClusterNetworkConfigHaControlPlaneConfigControlPlaneIpBlock
	SetInternalValue(val *GoogleGkeonpremVmwareAdminClusterNetworkConfigHaControlPlaneConfigControlPlaneIpBlock)
	Ips() GoogleGkeonpremVmwareAdminClusterNetworkConfigHaControlPlaneConfigControlPlaneIpBlockIpsList
	IpsInput() interface{}
	Netmask() *string
	SetNetmask(val *string)
	NetmaskInput() *string
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
	PutIps(value interface{})
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleGkeonpremVmwareAdminClusterNetworkConfigHaControlPlaneConfigControlPlaneIpBlockOutputReference
type jsiiProxy_GoogleGkeonpremVmwareAdminClusterNetworkConfigHaControlPlaneConfigControlPlaneIpBlockOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminClusterNetworkConfigHaControlPlaneConfigControlPlaneIpBlockOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminClusterNetworkConfigHaControlPlaneConfigControlPlaneIpBlockOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminClusterNetworkConfigHaControlPlaneConfigControlPlaneIpBlockOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminClusterNetworkConfigHaControlPlaneConfigControlPlaneIpBlockOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminClusterNetworkConfigHaControlPlaneConfigControlPlaneIpBlockOutputReference) Gateway() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminClusterNetworkConfigHaControlPlaneConfigControlPlaneIpBlockOutputReference) GatewayInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminClusterNetworkConfigHaControlPlaneConfigControlPlaneIpBlockOutputReference) InternalValue() *GoogleGkeonpremVmwareAdminClusterNetworkConfigHaControlPlaneConfigControlPlaneIpBlock {
	var returns *GoogleGkeonpremVmwareAdminClusterNetworkConfigHaControlPlaneConfigControlPlaneIpBlock
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminClusterNetworkConfigHaControlPlaneConfigControlPlaneIpBlockOutputReference) Ips() GoogleGkeonpremVmwareAdminClusterNetworkConfigHaControlPlaneConfigControlPlaneIpBlockIpsList {
	var returns GoogleGkeonpremVmwareAdminClusterNetworkConfigHaControlPlaneConfigControlPlaneIpBlockIpsList
	_jsii_.Get(
		j,
		"ips",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminClusterNetworkConfigHaControlPlaneConfigControlPlaneIpBlockOutputReference) IpsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminClusterNetworkConfigHaControlPlaneConfigControlPlaneIpBlockOutputReference) Netmask() *string {
	var returns *string
	_jsii_.Get(
		j,
		"netmask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminClusterNetworkConfigHaControlPlaneConfigControlPlaneIpBlockOutputReference) NetmaskInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"netmaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminClusterNetworkConfigHaControlPlaneConfigControlPlaneIpBlockOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminClusterNetworkConfigHaControlPlaneConfigControlPlaneIpBlockOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleGkeonpremVmwareAdminClusterNetworkConfigHaControlPlaneConfigControlPlaneIpBlockOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleGkeonpremVmwareAdminClusterNetworkConfigHaControlPlaneConfigControlPlaneIpBlockOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleGkeonpremVmwareAdminClusterNetworkConfigHaControlPlaneConfigControlPlaneIpBlockOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleGkeonpremVmwareAdminClusterNetworkConfigHaControlPlaneConfigControlPlaneIpBlockOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleGkeonpremVmwareAdminCluster.GoogleGkeonpremVmwareAdminClusterNetworkConfigHaControlPlaneConfigControlPlaneIpBlockOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleGkeonpremVmwareAdminClusterNetworkConfigHaControlPlaneConfigControlPlaneIpBlockOutputReference_Override(g GoogleGkeonpremVmwareAdminClusterNetworkConfigHaControlPlaneConfigControlPlaneIpBlockOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleGkeonpremVmwareAdminCluster.GoogleGkeonpremVmwareAdminClusterNetworkConfigHaControlPlaneConfigControlPlaneIpBlockOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminClusterNetworkConfigHaControlPlaneConfigControlPlaneIpBlockOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminClusterNetworkConfigHaControlPlaneConfigControlPlaneIpBlockOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminClusterNetworkConfigHaControlPlaneConfigControlPlaneIpBlockOutputReference)SetGateway(val *string) {
	if err := j.validateSetGatewayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gateway",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminClusterNetworkConfigHaControlPlaneConfigControlPlaneIpBlockOutputReference)SetInternalValue(val *GoogleGkeonpremVmwareAdminClusterNetworkConfigHaControlPlaneConfigControlPlaneIpBlock) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminClusterNetworkConfigHaControlPlaneConfigControlPlaneIpBlockOutputReference)SetNetmask(val *string) {
	if err := j.validateSetNetmaskParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"netmask",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminClusterNetworkConfigHaControlPlaneConfigControlPlaneIpBlockOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeonpremVmwareAdminClusterNetworkConfigHaControlPlaneConfigControlPlaneIpBlockOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminClusterNetworkConfigHaControlPlaneConfigControlPlaneIpBlockOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminClusterNetworkConfigHaControlPlaneConfigControlPlaneIpBlockOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminClusterNetworkConfigHaControlPlaneConfigControlPlaneIpBlockOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminClusterNetworkConfigHaControlPlaneConfigControlPlaneIpBlockOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminClusterNetworkConfigHaControlPlaneConfigControlPlaneIpBlockOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminClusterNetworkConfigHaControlPlaneConfigControlPlaneIpBlockOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminClusterNetworkConfigHaControlPlaneConfigControlPlaneIpBlockOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminClusterNetworkConfigHaControlPlaneConfigControlPlaneIpBlockOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminClusterNetworkConfigHaControlPlaneConfigControlPlaneIpBlockOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminClusterNetworkConfigHaControlPlaneConfigControlPlaneIpBlockOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminClusterNetworkConfigHaControlPlaneConfigControlPlaneIpBlockOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminClusterNetworkConfigHaControlPlaneConfigControlPlaneIpBlockOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminClusterNetworkConfigHaControlPlaneConfigControlPlaneIpBlockOutputReference) PutIps(value interface{}) {
	if err := g.validatePutIpsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putIps",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminClusterNetworkConfigHaControlPlaneConfigControlPlaneIpBlockOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleGkeonpremVmwareAdminClusterNetworkConfigHaControlPlaneConfigControlPlaneIpBlockOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

