package googlecontainercluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecontainercluster/internal"
)

type GoogleContainerClusterDnsConfigOutputReference interface {
	cdktf.ComplexObject
	ClusterDns() *string
	SetClusterDns(val *string)
	ClusterDnsDomain() *string
	SetClusterDnsDomain(val *string)
	ClusterDnsDomainInput() *string
	ClusterDnsInput() *string
	ClusterDnsScope() *string
	SetClusterDnsScope(val *string)
	ClusterDnsScopeInput() *string
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
	InternalValue() *GoogleContainerClusterDnsConfig
	SetInternalValue(val *GoogleContainerClusterDnsConfig)
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
	ResetClusterDns()
	ResetClusterDnsDomain()
	ResetClusterDnsScope()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleContainerClusterDnsConfigOutputReference
type jsiiProxy_GoogleContainerClusterDnsConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleContainerClusterDnsConfigOutputReference) ClusterDns() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterDns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterDnsConfigOutputReference) ClusterDnsDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterDnsDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterDnsConfigOutputReference) ClusterDnsDomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterDnsDomainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterDnsConfigOutputReference) ClusterDnsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterDnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterDnsConfigOutputReference) ClusterDnsScope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterDnsScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterDnsConfigOutputReference) ClusterDnsScopeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterDnsScopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterDnsConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterDnsConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterDnsConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterDnsConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterDnsConfigOutputReference) InternalValue() *GoogleContainerClusterDnsConfig {
	var returns *GoogleContainerClusterDnsConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterDnsConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterDnsConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleContainerClusterDnsConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleContainerClusterDnsConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleContainerClusterDnsConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleContainerClusterDnsConfigOutputReference{}

	_jsii_.Create(
		"google-beta.googleContainerCluster.GoogleContainerClusterDnsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleContainerClusterDnsConfigOutputReference_Override(g GoogleContainerClusterDnsConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleContainerCluster.GoogleContainerClusterDnsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleContainerClusterDnsConfigOutputReference)SetClusterDns(val *string) {
	if err := j.validateSetClusterDnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterDns",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterDnsConfigOutputReference)SetClusterDnsDomain(val *string) {
	if err := j.validateSetClusterDnsDomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterDnsDomain",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterDnsConfigOutputReference)SetClusterDnsScope(val *string) {
	if err := j.validateSetClusterDnsScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterDnsScope",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterDnsConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterDnsConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterDnsConfigOutputReference)SetInternalValue(val *GoogleContainerClusterDnsConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterDnsConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterDnsConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleContainerClusterDnsConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerClusterDnsConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleContainerClusterDnsConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleContainerClusterDnsConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleContainerClusterDnsConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleContainerClusterDnsConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleContainerClusterDnsConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleContainerClusterDnsConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleContainerClusterDnsConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleContainerClusterDnsConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleContainerClusterDnsConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerClusterDnsConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleContainerClusterDnsConfigOutputReference) ResetClusterDns() {
	_jsii_.InvokeVoid(
		g,
		"resetClusterDns",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterDnsConfigOutputReference) ResetClusterDnsDomain() {
	_jsii_.InvokeVoid(
		g,
		"resetClusterDnsDomain",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterDnsConfigOutputReference) ResetClusterDnsScope() {
	_jsii_.InvokeVoid(
		g,
		"resetClusterDnsScope",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterDnsConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleContainerClusterDnsConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

