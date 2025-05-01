package googlemanagedkafkaconnectcluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlemanagedkafkaconnectcluster/internal"
)

type GoogleManagedKafkaConnectClusterGcpConfigAccessConfigNetworkConfigsOutputReference interface {
	cdktf.ComplexObject
	AdditionalSubnets() *[]*string
	SetAdditionalSubnets(val *[]*string)
	AdditionalSubnetsInput() *[]*string
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
	DnsDomainNames() *[]*string
	SetDnsDomainNames(val *[]*string)
	DnsDomainNamesInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	PrimarySubnet() *string
	SetPrimarySubnet(val *string)
	PrimarySubnetInput() *string
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
	ResetAdditionalSubnets()
	ResetDnsDomainNames()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleManagedKafkaConnectClusterGcpConfigAccessConfigNetworkConfigsOutputReference
type jsiiProxy_GoogleManagedKafkaConnectClusterGcpConfigAccessConfigNetworkConfigsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleManagedKafkaConnectClusterGcpConfigAccessConfigNetworkConfigsOutputReference) AdditionalSubnets() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"additionalSubnets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleManagedKafkaConnectClusterGcpConfigAccessConfigNetworkConfigsOutputReference) AdditionalSubnetsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"additionalSubnetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleManagedKafkaConnectClusterGcpConfigAccessConfigNetworkConfigsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleManagedKafkaConnectClusterGcpConfigAccessConfigNetworkConfigsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleManagedKafkaConnectClusterGcpConfigAccessConfigNetworkConfigsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleManagedKafkaConnectClusterGcpConfigAccessConfigNetworkConfigsOutputReference) DnsDomainNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dnsDomainNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleManagedKafkaConnectClusterGcpConfigAccessConfigNetworkConfigsOutputReference) DnsDomainNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dnsDomainNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleManagedKafkaConnectClusterGcpConfigAccessConfigNetworkConfigsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleManagedKafkaConnectClusterGcpConfigAccessConfigNetworkConfigsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleManagedKafkaConnectClusterGcpConfigAccessConfigNetworkConfigsOutputReference) PrimarySubnet() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primarySubnet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleManagedKafkaConnectClusterGcpConfigAccessConfigNetworkConfigsOutputReference) PrimarySubnetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primarySubnetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleManagedKafkaConnectClusterGcpConfigAccessConfigNetworkConfigsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleManagedKafkaConnectClusterGcpConfigAccessConfigNetworkConfigsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleManagedKafkaConnectClusterGcpConfigAccessConfigNetworkConfigsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleManagedKafkaConnectClusterGcpConfigAccessConfigNetworkConfigsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleManagedKafkaConnectClusterGcpConfigAccessConfigNetworkConfigsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleManagedKafkaConnectClusterGcpConfigAccessConfigNetworkConfigsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleManagedKafkaConnectCluster.GoogleManagedKafkaConnectClusterGcpConfigAccessConfigNetworkConfigsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleManagedKafkaConnectClusterGcpConfigAccessConfigNetworkConfigsOutputReference_Override(g GoogleManagedKafkaConnectClusterGcpConfigAccessConfigNetworkConfigsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleManagedKafkaConnectCluster.GoogleManagedKafkaConnectClusterGcpConfigAccessConfigNetworkConfigsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleManagedKafkaConnectClusterGcpConfigAccessConfigNetworkConfigsOutputReference)SetAdditionalSubnets(val *[]*string) {
	if err := j.validateSetAdditionalSubnetsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"additionalSubnets",
		val,
	)
}

func (j *jsiiProxy_GoogleManagedKafkaConnectClusterGcpConfigAccessConfigNetworkConfigsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleManagedKafkaConnectClusterGcpConfigAccessConfigNetworkConfigsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleManagedKafkaConnectClusterGcpConfigAccessConfigNetworkConfigsOutputReference)SetDnsDomainNames(val *[]*string) {
	if err := j.validateSetDnsDomainNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dnsDomainNames",
		val,
	)
}

func (j *jsiiProxy_GoogleManagedKafkaConnectClusterGcpConfigAccessConfigNetworkConfigsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleManagedKafkaConnectClusterGcpConfigAccessConfigNetworkConfigsOutputReference)SetPrimarySubnet(val *string) {
	if err := j.validateSetPrimarySubnetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"primarySubnet",
		val,
	)
}

func (j *jsiiProxy_GoogleManagedKafkaConnectClusterGcpConfigAccessConfigNetworkConfigsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleManagedKafkaConnectClusterGcpConfigAccessConfigNetworkConfigsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleManagedKafkaConnectClusterGcpConfigAccessConfigNetworkConfigsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleManagedKafkaConnectClusterGcpConfigAccessConfigNetworkConfigsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleManagedKafkaConnectClusterGcpConfigAccessConfigNetworkConfigsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleManagedKafkaConnectClusterGcpConfigAccessConfigNetworkConfigsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleManagedKafkaConnectClusterGcpConfigAccessConfigNetworkConfigsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleManagedKafkaConnectClusterGcpConfigAccessConfigNetworkConfigsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleManagedKafkaConnectClusterGcpConfigAccessConfigNetworkConfigsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleManagedKafkaConnectClusterGcpConfigAccessConfigNetworkConfigsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleManagedKafkaConnectClusterGcpConfigAccessConfigNetworkConfigsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleManagedKafkaConnectClusterGcpConfigAccessConfigNetworkConfigsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleManagedKafkaConnectClusterGcpConfigAccessConfigNetworkConfigsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleManagedKafkaConnectClusterGcpConfigAccessConfigNetworkConfigsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleManagedKafkaConnectClusterGcpConfigAccessConfigNetworkConfigsOutputReference) ResetAdditionalSubnets() {
	_jsii_.InvokeVoid(
		g,
		"resetAdditionalSubnets",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleManagedKafkaConnectClusterGcpConfigAccessConfigNetworkConfigsOutputReference) ResetDnsDomainNames() {
	_jsii_.InvokeVoid(
		g,
		"resetDnsDomainNames",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleManagedKafkaConnectClusterGcpConfigAccessConfigNetworkConfigsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleManagedKafkaConnectClusterGcpConfigAccessConfigNetworkConfigsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

