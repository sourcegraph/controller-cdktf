package googlecontainercluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecontainercluster/internal"
)

type GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference interface {
	cdktn.ComplexObject
	Ca() GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsCaList
	CaInput() interface{}
	Capabilities() *[]*string
	SetCapabilities(val *[]*string)
	CapabilitiesInput() *[]*string
	Client() GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsClientList
	ClientInput() interface{}
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
	DialTimeout() *string
	SetDialTimeout(val *string)
	DialTimeoutInput() *string
	// Experimental.
	Fqn() *string
	Header() GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsHeaderList
	HeaderInput() interface{}
	Host() *string
	SetHost(val *string)
	HostInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	OverridePath() interface{}
	SetOverridePath(val interface{})
	OverridePathInput() interface{}
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
	PutCa(value interface{})
	PutClient(value interface{})
	PutHeader(value interface{})
	ResetCa()
	ResetCapabilities()
	ResetClient()
	ResetDialTimeout()
	ResetHeader()
	ResetOverridePath()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference
type jsiiProxy_GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference) Ca() GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsCaList {
	var returns GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsCaList
	_jsii_.Get(
		j,
		"ca",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference) CaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"caInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference) Capabilities() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"capabilities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference) CapabilitiesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"capabilitiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference) Client() GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsClientList {
	var returns GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsClientList
	_jsii_.Get(
		j,
		"client",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference) ClientInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference) DialTimeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dialTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference) DialTimeoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dialTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference) Header() GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsHeaderList {
	var returns GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsHeaderList
	_jsii_.Get(
		j,
		"header",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference) HeaderInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"headerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference) Host() *string {
	var returns *string
	_jsii_.Get(
		j,
		"host",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference) HostInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference) OverridePath() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"overridePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference) OverridePathInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"overridePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleContainerCluster.GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference_Override(g GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleContainerCluster.GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference)SetCapabilities(val *[]*string) {
	if err := j.validateSetCapabilitiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"capabilities",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference)SetDialTimeout(val *string) {
	if err := j.validateSetDialTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dialTimeout",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference)SetHost(val *string) {
	if err := j.validateSetHostParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"host",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference)SetOverridePath(val interface{}) {
	if err := j.validateSetOverridePathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"overridePath",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference) PutCa(value interface{}) {
	if err := g.validatePutCaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putCa",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference) PutClient(value interface{}) {
	if err := g.validatePutClientParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putClient",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference) PutHeader(value interface{}) {
	if err := g.validatePutHeaderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putHeader",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference) ResetCa() {
	_jsii_.InvokeVoid(
		g,
		"resetCa",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference) ResetCapabilities() {
	_jsii_.InvokeVoid(
		g,
		"resetCapabilities",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference) ResetClient() {
	_jsii_.InvokeVoid(
		g,
		"resetClient",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference) ResetDialTimeout() {
	_jsii_.InvokeVoid(
		g,
		"resetDialTimeout",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference) ResetHeader() {
	_jsii_.InvokeVoid(
		g,
		"resetHeader",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference) ResetOverridePath() {
	_jsii_.InvokeVoid(
		g,
		"resetOverridePath",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

