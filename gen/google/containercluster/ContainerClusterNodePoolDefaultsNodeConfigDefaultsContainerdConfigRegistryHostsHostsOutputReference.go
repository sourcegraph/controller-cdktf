package containercluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/containercluster/internal"
)

type ContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference interface {
	cdktf.ComplexObject
	Ca() ContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsCaList
	CaInput() any
	Capabilities() *[]*string
	SetCapabilities(val *[]*string)
	CapabilitiesInput() *[]*string
	Client() ContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsClientList
	ClientInput() any
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
	DialTimeout() *string
	SetDialTimeout(val *string)
	DialTimeoutInput() *string
	// Experimental.
	Fqn() *string
	Header() ContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsHeaderList
	HeaderInput() any
	Host() *string
	SetHost(val *string)
	HostInput() *string
	InternalValue() any
	SetInternalValue(val any)
	OverridePath() any
	SetOverridePath(val any)
	OverridePathInput() any
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
	PutCa(value any)
	PutClient(value any)
	PutHeader(value any)
	ResetCa()
	ResetCapabilities()
	ResetClient()
	ResetDialTimeout()
	ResetHeader()
	ResetOverridePath()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference
type jsiiProxy_ContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference) Ca() ContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsCaList {
	var returns ContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsCaList
	_jsii_.Get(
		j,
		"ca",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference) CaInput() any {
	var returns any
	_jsii_.Get(
		j,
		"caInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference) Capabilities() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"capabilities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference) CapabilitiesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"capabilitiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference) Client() ContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsClientList {
	var returns ContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsClientList
	_jsii_.Get(
		j,
		"client",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference) ClientInput() any {
	var returns any
	_jsii_.Get(
		j,
		"clientInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference) DialTimeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dialTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference) DialTimeoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dialTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference) Header() ContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsHeaderList {
	var returns ContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsHeaderList
	_jsii_.Get(
		j,
		"header",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference) HeaderInput() any {
	var returns any
	_jsii_.Get(
		j,
		"headerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference) Host() *string {
	var returns *string
	_jsii_.Get(
		j,
		"host",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference) HostInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference) InternalValue() any {
	var returns any
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference) OverridePath() any {
	var returns any
	_jsii_.Get(
		j,
		"overridePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference) OverridePathInput() any {
	var returns any
	_jsii_.Get(
		j,
		"overridePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference {
	_init_.Initialize()

	if err := validateNewContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.containerCluster.ContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference_Override(c ContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.containerCluster.ContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_ContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference) SetCapabilities(val *[]*string) {
	if err := j.validateSetCapabilitiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"capabilities",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference) SetDialTimeout(val *string) {
	if err := j.validateSetDialTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dialTimeout",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference) SetHost(val *string) {
	if err := j.validateSetHostParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"host",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference) SetInternalValue(val any) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference) SetOverridePath(val any) {
	if err := j.validateSetOverridePathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"overridePath",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference) PutCa(value any) {
	if err := c.validatePutCaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCa",
		[]any{value},
	)
}

func (c *jsiiProxy_ContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference) PutClient(value any) {
	if err := c.validatePutClientParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putClient",
		[]any{value},
	)
}

func (c *jsiiProxy_ContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference) PutHeader(value any) {
	if err := c.validatePutHeaderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putHeader",
		[]any{value},
	)
}

func (c *jsiiProxy_ContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference) ResetCa() {
	_jsii_.InvokeVoid(
		c,
		"resetCa",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference) ResetCapabilities() {
	_jsii_.InvokeVoid(
		c,
		"resetCapabilities",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference) ResetClient() {
	_jsii_.InvokeVoid(
		c,
		"resetClient",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference) ResetDialTimeout() {
	_jsii_.InvokeVoid(
		c,
		"resetDialTimeout",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference) ResetHeader() {
	_jsii_.InvokeVoid(
		c,
		"resetHeader",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference) ResetOverridePath() {
	_jsii_.InvokeVoid(
		c,
		"resetOverridePath",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		c,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigRegistryHostsHostsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
