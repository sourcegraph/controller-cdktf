package containercluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/containercluster/internal"
)

type ContainerClusterNodeConfigContainerdConfigOutputReference interface {
	cdktf.ComplexObject
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
	InternalValue() *ContainerClusterNodeConfigContainerdConfig
	SetInternalValue(val *ContainerClusterNodeConfigContainerdConfig)
	PrivateRegistryAccessConfig() ContainerClusterNodeConfigContainerdConfigPrivateRegistryAccessConfigOutputReference
	PrivateRegistryAccessConfigInput() *ContainerClusterNodeConfigContainerdConfigPrivateRegistryAccessConfig
	RegistryHosts() ContainerClusterNodeConfigContainerdConfigRegistryHostsList
	RegistryHostsInput() any
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WritableCgroups() ContainerClusterNodeConfigContainerdConfigWritableCgroupsOutputReference
	WritableCgroupsInput() *ContainerClusterNodeConfigContainerdConfigWritableCgroups
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
	PutPrivateRegistryAccessConfig(value *ContainerClusterNodeConfigContainerdConfigPrivateRegistryAccessConfig)
	PutRegistryHosts(value any)
	PutWritableCgroups(value *ContainerClusterNodeConfigContainerdConfigWritableCgroups)
	ResetPrivateRegistryAccessConfig()
	ResetRegistryHosts()
	ResetWritableCgroups()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ContainerClusterNodeConfigContainerdConfigOutputReference
type jsiiProxy_ContainerClusterNodeConfigContainerdConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ContainerClusterNodeConfigContainerdConfigOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigContainerdConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigContainerdConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigContainerdConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigContainerdConfigOutputReference) InternalValue() *ContainerClusterNodeConfigContainerdConfig {
	var returns *ContainerClusterNodeConfigContainerdConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigContainerdConfigOutputReference) PrivateRegistryAccessConfig() ContainerClusterNodeConfigContainerdConfigPrivateRegistryAccessConfigOutputReference {
	var returns ContainerClusterNodeConfigContainerdConfigPrivateRegistryAccessConfigOutputReference
	_jsii_.Get(
		j,
		"privateRegistryAccessConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigContainerdConfigOutputReference) PrivateRegistryAccessConfigInput() *ContainerClusterNodeConfigContainerdConfigPrivateRegistryAccessConfig {
	var returns *ContainerClusterNodeConfigContainerdConfigPrivateRegistryAccessConfig
	_jsii_.Get(
		j,
		"privateRegistryAccessConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigContainerdConfigOutputReference) RegistryHosts() ContainerClusterNodeConfigContainerdConfigRegistryHostsList {
	var returns ContainerClusterNodeConfigContainerdConfigRegistryHostsList
	_jsii_.Get(
		j,
		"registryHosts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigContainerdConfigOutputReference) RegistryHostsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"registryHostsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigContainerdConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigContainerdConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigContainerdConfigOutputReference) WritableCgroups() ContainerClusterNodeConfigContainerdConfigWritableCgroupsOutputReference {
	var returns ContainerClusterNodeConfigContainerdConfigWritableCgroupsOutputReference
	_jsii_.Get(
		j,
		"writableCgroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigContainerdConfigOutputReference) WritableCgroupsInput() *ContainerClusterNodeConfigContainerdConfigWritableCgroups {
	var returns *ContainerClusterNodeConfigContainerdConfigWritableCgroups
	_jsii_.Get(
		j,
		"writableCgroupsInput",
		&returns,
	)
	return returns
}

func NewContainerClusterNodeConfigContainerdConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ContainerClusterNodeConfigContainerdConfigOutputReference {
	_init_.Initialize()

	if err := validateNewContainerClusterNodeConfigContainerdConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ContainerClusterNodeConfigContainerdConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.containerCluster.ContainerClusterNodeConfigContainerdConfigOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewContainerClusterNodeConfigContainerdConfigOutputReference_Override(c ContainerClusterNodeConfigContainerdConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.containerCluster.ContainerClusterNodeConfigContainerdConfigOutputReference",
		[]any{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ContainerClusterNodeConfigContainerdConfigOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodeConfigContainerdConfigOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodeConfigContainerdConfigOutputReference) SetInternalValue(val *ContainerClusterNodeConfigContainerdConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodeConfigContainerdConfigOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodeConfigContainerdConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ContainerClusterNodeConfigContainerdConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterNodeConfigContainerdConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (c *jsiiProxy_ContainerClusterNodeConfigContainerdConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ContainerClusterNodeConfigContainerdConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ContainerClusterNodeConfigContainerdConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ContainerClusterNodeConfigContainerdConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ContainerClusterNodeConfigContainerdConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ContainerClusterNodeConfigContainerdConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ContainerClusterNodeConfigContainerdConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ContainerClusterNodeConfigContainerdConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ContainerClusterNodeConfigContainerdConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterNodeConfigContainerdConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ContainerClusterNodeConfigContainerdConfigOutputReference) PutPrivateRegistryAccessConfig(value *ContainerClusterNodeConfigContainerdConfigPrivateRegistryAccessConfig) {
	if err := c.validatePutPrivateRegistryAccessConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putPrivateRegistryAccessConfig",
		[]any{value},
	)
}

func (c *jsiiProxy_ContainerClusterNodeConfigContainerdConfigOutputReference) PutRegistryHosts(value any) {
	if err := c.validatePutRegistryHostsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRegistryHosts",
		[]any{value},
	)
}

func (c *jsiiProxy_ContainerClusterNodeConfigContainerdConfigOutputReference) PutWritableCgroups(value *ContainerClusterNodeConfigContainerdConfigWritableCgroups) {
	if err := c.validatePutWritableCgroupsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putWritableCgroups",
		[]any{value},
	)
}

func (c *jsiiProxy_ContainerClusterNodeConfigContainerdConfigOutputReference) ResetPrivateRegistryAccessConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetPrivateRegistryAccessConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodeConfigContainerdConfigOutputReference) ResetRegistryHosts() {
	_jsii_.InvokeVoid(
		c,
		"resetRegistryHosts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodeConfigContainerdConfigOutputReference) ResetWritableCgroups() {
	_jsii_.InvokeVoid(
		c,
		"resetWritableCgroups",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodeConfigContainerdConfigOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (c *jsiiProxy_ContainerClusterNodeConfigContainerdConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
