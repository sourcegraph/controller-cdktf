package googlecontainercluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecontainercluster/internal"
)

type GoogleContainerClusterControlPlaneEndpointsConfigOutputReference interface {
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DnsEndpointConfig() GoogleContainerClusterControlPlaneEndpointsConfigDnsEndpointConfigOutputReference
	DnsEndpointConfigInput() *GoogleContainerClusterControlPlaneEndpointsConfigDnsEndpointConfig
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleContainerClusterControlPlaneEndpointsConfig
	SetInternalValue(val *GoogleContainerClusterControlPlaneEndpointsConfig)
	IpEndpointsConfig() GoogleContainerClusterControlPlaneEndpointsConfigIpEndpointsConfigOutputReference
	IpEndpointsConfigInput() *GoogleContainerClusterControlPlaneEndpointsConfigIpEndpointsConfig
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
	PutDnsEndpointConfig(value *GoogleContainerClusterControlPlaneEndpointsConfigDnsEndpointConfig)
	PutIpEndpointsConfig(value *GoogleContainerClusterControlPlaneEndpointsConfigIpEndpointsConfig)
	ResetDnsEndpointConfig()
	ResetIpEndpointsConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleContainerClusterControlPlaneEndpointsConfigOutputReference
type jsiiProxy_GoogleContainerClusterControlPlaneEndpointsConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleContainerClusterControlPlaneEndpointsConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterControlPlaneEndpointsConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterControlPlaneEndpointsConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterControlPlaneEndpointsConfigOutputReference) DnsEndpointConfig() GoogleContainerClusterControlPlaneEndpointsConfigDnsEndpointConfigOutputReference {
	var returns GoogleContainerClusterControlPlaneEndpointsConfigDnsEndpointConfigOutputReference
	_jsii_.Get(
		j,
		"dnsEndpointConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterControlPlaneEndpointsConfigOutputReference) DnsEndpointConfigInput() *GoogleContainerClusterControlPlaneEndpointsConfigDnsEndpointConfig {
	var returns *GoogleContainerClusterControlPlaneEndpointsConfigDnsEndpointConfig
	_jsii_.Get(
		j,
		"dnsEndpointConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterControlPlaneEndpointsConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterControlPlaneEndpointsConfigOutputReference) InternalValue() *GoogleContainerClusterControlPlaneEndpointsConfig {
	var returns *GoogleContainerClusterControlPlaneEndpointsConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterControlPlaneEndpointsConfigOutputReference) IpEndpointsConfig() GoogleContainerClusterControlPlaneEndpointsConfigIpEndpointsConfigOutputReference {
	var returns GoogleContainerClusterControlPlaneEndpointsConfigIpEndpointsConfigOutputReference
	_jsii_.Get(
		j,
		"ipEndpointsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterControlPlaneEndpointsConfigOutputReference) IpEndpointsConfigInput() *GoogleContainerClusterControlPlaneEndpointsConfigIpEndpointsConfig {
	var returns *GoogleContainerClusterControlPlaneEndpointsConfigIpEndpointsConfig
	_jsii_.Get(
		j,
		"ipEndpointsConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterControlPlaneEndpointsConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterControlPlaneEndpointsConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleContainerClusterControlPlaneEndpointsConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleContainerClusterControlPlaneEndpointsConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleContainerClusterControlPlaneEndpointsConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleContainerClusterControlPlaneEndpointsConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleContainerCluster.GoogleContainerClusterControlPlaneEndpointsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleContainerClusterControlPlaneEndpointsConfigOutputReference_Override(g GoogleContainerClusterControlPlaneEndpointsConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleContainerCluster.GoogleContainerClusterControlPlaneEndpointsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleContainerClusterControlPlaneEndpointsConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterControlPlaneEndpointsConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterControlPlaneEndpointsConfigOutputReference)SetInternalValue(val *GoogleContainerClusterControlPlaneEndpointsConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterControlPlaneEndpointsConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterControlPlaneEndpointsConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleContainerClusterControlPlaneEndpointsConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerClusterControlPlaneEndpointsConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleContainerClusterControlPlaneEndpointsConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleContainerClusterControlPlaneEndpointsConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleContainerClusterControlPlaneEndpointsConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleContainerClusterControlPlaneEndpointsConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleContainerClusterControlPlaneEndpointsConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleContainerClusterControlPlaneEndpointsConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleContainerClusterControlPlaneEndpointsConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleContainerClusterControlPlaneEndpointsConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleContainerClusterControlPlaneEndpointsConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerClusterControlPlaneEndpointsConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleContainerClusterControlPlaneEndpointsConfigOutputReference) PutDnsEndpointConfig(value *GoogleContainerClusterControlPlaneEndpointsConfigDnsEndpointConfig) {
	if err := g.validatePutDnsEndpointConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDnsEndpointConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerClusterControlPlaneEndpointsConfigOutputReference) PutIpEndpointsConfig(value *GoogleContainerClusterControlPlaneEndpointsConfigIpEndpointsConfig) {
	if err := g.validatePutIpEndpointsConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putIpEndpointsConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerClusterControlPlaneEndpointsConfigOutputReference) ResetDnsEndpointConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetDnsEndpointConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterControlPlaneEndpointsConfigOutputReference) ResetIpEndpointsConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetIpEndpointsConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterControlPlaneEndpointsConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleContainerClusterControlPlaneEndpointsConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

