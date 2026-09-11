package googlevmwareenginecluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlevmwareenginecluster/internal"
)

type GoogleVmwareengineClusterDatastoreMountConfigDatastoreNetworkOutputReference interface {
	cdktn.ComplexObject
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
	ConnectionCount() *float64
	SetConnectionCount(val *float64)
	ConnectionCountInput() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleVmwareengineClusterDatastoreMountConfigDatastoreNetwork
	SetInternalValue(val *GoogleVmwareengineClusterDatastoreMountConfigDatastoreNetwork)
	Mtu() *float64
	SetMtu(val *float64)
	MtuInput() *float64
	NetworkPeering() *string
	Subnet() *string
	SetSubnet(val *string)
	SubnetInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
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
	InterpolationAsList() cdktn.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	ResetConnectionCount()
	ResetMtu()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleVmwareengineClusterDatastoreMountConfigDatastoreNetworkOutputReference
type jsiiProxy_GoogleVmwareengineClusterDatastoreMountConfigDatastoreNetworkOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleVmwareengineClusterDatastoreMountConfigDatastoreNetworkOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineClusterDatastoreMountConfigDatastoreNetworkOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineClusterDatastoreMountConfigDatastoreNetworkOutputReference) ConnectionCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"connectionCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineClusterDatastoreMountConfigDatastoreNetworkOutputReference) ConnectionCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"connectionCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineClusterDatastoreMountConfigDatastoreNetworkOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineClusterDatastoreMountConfigDatastoreNetworkOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineClusterDatastoreMountConfigDatastoreNetworkOutputReference) InternalValue() *GoogleVmwareengineClusterDatastoreMountConfigDatastoreNetwork {
	var returns *GoogleVmwareengineClusterDatastoreMountConfigDatastoreNetwork
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineClusterDatastoreMountConfigDatastoreNetworkOutputReference) Mtu() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"mtu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineClusterDatastoreMountConfigDatastoreNetworkOutputReference) MtuInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"mtuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineClusterDatastoreMountConfigDatastoreNetworkOutputReference) NetworkPeering() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkPeering",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineClusterDatastoreMountConfigDatastoreNetworkOutputReference) Subnet() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineClusterDatastoreMountConfigDatastoreNetworkOutputReference) SubnetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineClusterDatastoreMountConfigDatastoreNetworkOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineClusterDatastoreMountConfigDatastoreNetworkOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleVmwareengineClusterDatastoreMountConfigDatastoreNetworkOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleVmwareengineClusterDatastoreMountConfigDatastoreNetworkOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleVmwareengineClusterDatastoreMountConfigDatastoreNetworkOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleVmwareengineClusterDatastoreMountConfigDatastoreNetworkOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleVmwareengineCluster.GoogleVmwareengineClusterDatastoreMountConfigDatastoreNetworkOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleVmwareengineClusterDatastoreMountConfigDatastoreNetworkOutputReference_Override(g GoogleVmwareengineClusterDatastoreMountConfigDatastoreNetworkOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleVmwareengineCluster.GoogleVmwareengineClusterDatastoreMountConfigDatastoreNetworkOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleVmwareengineClusterDatastoreMountConfigDatastoreNetworkOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleVmwareengineClusterDatastoreMountConfigDatastoreNetworkOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleVmwareengineClusterDatastoreMountConfigDatastoreNetworkOutputReference)SetConnectionCount(val *float64) {
	if err := j.validateSetConnectionCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionCount",
		val,
	)
}

func (j *jsiiProxy_GoogleVmwareengineClusterDatastoreMountConfigDatastoreNetworkOutputReference)SetInternalValue(val *GoogleVmwareengineClusterDatastoreMountConfigDatastoreNetwork) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleVmwareengineClusterDatastoreMountConfigDatastoreNetworkOutputReference)SetMtu(val *float64) {
	if err := j.validateSetMtuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mtu",
		val,
	)
}

func (j *jsiiProxy_GoogleVmwareengineClusterDatastoreMountConfigDatastoreNetworkOutputReference)SetSubnet(val *string) {
	if err := j.validateSetSubnetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnet",
		val,
	)
}

func (j *jsiiProxy_GoogleVmwareengineClusterDatastoreMountConfigDatastoreNetworkOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleVmwareengineClusterDatastoreMountConfigDatastoreNetworkOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleVmwareengineClusterDatastoreMountConfigDatastoreNetworkOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVmwareengineClusterDatastoreMountConfigDatastoreNetworkOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleVmwareengineClusterDatastoreMountConfigDatastoreNetworkOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVmwareengineClusterDatastoreMountConfigDatastoreNetworkOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleVmwareengineClusterDatastoreMountConfigDatastoreNetworkOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleVmwareengineClusterDatastoreMountConfigDatastoreNetworkOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleVmwareengineClusterDatastoreMountConfigDatastoreNetworkOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleVmwareengineClusterDatastoreMountConfigDatastoreNetworkOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleVmwareengineClusterDatastoreMountConfigDatastoreNetworkOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleVmwareengineClusterDatastoreMountConfigDatastoreNetworkOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleVmwareengineClusterDatastoreMountConfigDatastoreNetworkOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVmwareengineClusterDatastoreMountConfigDatastoreNetworkOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVmwareengineClusterDatastoreMountConfigDatastoreNetworkOutputReference) ResetConnectionCount() {
	_jsii_.InvokeVoid(
		g,
		"resetConnectionCount",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVmwareengineClusterDatastoreMountConfigDatastoreNetworkOutputReference) ResetMtu() {
	_jsii_.InvokeVoid(
		g,
		"resetMtu",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVmwareengineClusterDatastoreMountConfigDatastoreNetworkOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := g.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVmwareengineClusterDatastoreMountConfigDatastoreNetworkOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

