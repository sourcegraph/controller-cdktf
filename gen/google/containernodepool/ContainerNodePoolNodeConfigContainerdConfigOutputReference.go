package containernodepool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google/containernodepool/internal"
)

type ContainerNodePoolNodeConfigContainerdConfigOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() *ContainerNodePoolNodeConfigContainerdConfig
	SetInternalValue(val *ContainerNodePoolNodeConfigContainerdConfig)
	PrivateRegistryAccessConfig() ContainerNodePoolNodeConfigContainerdConfigPrivateRegistryAccessConfigOutputReference
	PrivateRegistryAccessConfigInput() *ContainerNodePoolNodeConfigContainerdConfigPrivateRegistryAccessConfig
	RegistryHosts() ContainerNodePoolNodeConfigContainerdConfigRegistryHostsList
	RegistryHostsInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	WritableCgroups() ContainerNodePoolNodeConfigContainerdConfigWritableCgroupsOutputReference
	WritableCgroupsInput() *ContainerNodePoolNodeConfigContainerdConfigWritableCgroups
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
	PutPrivateRegistryAccessConfig(value *ContainerNodePoolNodeConfigContainerdConfigPrivateRegistryAccessConfig)
	PutRegistryHosts(value interface{})
	PutWritableCgroups(value *ContainerNodePoolNodeConfigContainerdConfigWritableCgroups)
	ResetPrivateRegistryAccessConfig()
	ResetRegistryHosts()
	ResetWritableCgroups()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ContainerNodePoolNodeConfigContainerdConfigOutputReference
type jsiiProxy_ContainerNodePoolNodeConfigContainerdConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigContainerdConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigContainerdConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigContainerdConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigContainerdConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigContainerdConfigOutputReference) InternalValue() *ContainerNodePoolNodeConfigContainerdConfig {
	var returns *ContainerNodePoolNodeConfigContainerdConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigContainerdConfigOutputReference) PrivateRegistryAccessConfig() ContainerNodePoolNodeConfigContainerdConfigPrivateRegistryAccessConfigOutputReference {
	var returns ContainerNodePoolNodeConfigContainerdConfigPrivateRegistryAccessConfigOutputReference
	_jsii_.Get(
		j,
		"privateRegistryAccessConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigContainerdConfigOutputReference) PrivateRegistryAccessConfigInput() *ContainerNodePoolNodeConfigContainerdConfigPrivateRegistryAccessConfig {
	var returns *ContainerNodePoolNodeConfigContainerdConfigPrivateRegistryAccessConfig
	_jsii_.Get(
		j,
		"privateRegistryAccessConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigContainerdConfigOutputReference) RegistryHosts() ContainerNodePoolNodeConfigContainerdConfigRegistryHostsList {
	var returns ContainerNodePoolNodeConfigContainerdConfigRegistryHostsList
	_jsii_.Get(
		j,
		"registryHosts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigContainerdConfigOutputReference) RegistryHostsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"registryHostsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigContainerdConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigContainerdConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigContainerdConfigOutputReference) WritableCgroups() ContainerNodePoolNodeConfigContainerdConfigWritableCgroupsOutputReference {
	var returns ContainerNodePoolNodeConfigContainerdConfigWritableCgroupsOutputReference
	_jsii_.Get(
		j,
		"writableCgroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigContainerdConfigOutputReference) WritableCgroupsInput() *ContainerNodePoolNodeConfigContainerdConfigWritableCgroups {
	var returns *ContainerNodePoolNodeConfigContainerdConfigWritableCgroups
	_jsii_.Get(
		j,
		"writableCgroupsInput",
		&returns,
	)
	return returns
}


func NewContainerNodePoolNodeConfigContainerdConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ContainerNodePoolNodeConfigContainerdConfigOutputReference {
	_init_.Initialize()

	if err := validateNewContainerNodePoolNodeConfigContainerdConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ContainerNodePoolNodeConfigContainerdConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.containerNodePool.ContainerNodePoolNodeConfigContainerdConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewContainerNodePoolNodeConfigContainerdConfigOutputReference_Override(c ContainerNodePoolNodeConfigContainerdConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.containerNodePool.ContainerNodePoolNodeConfigContainerdConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigContainerdConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigContainerdConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigContainerdConfigOutputReference)SetInternalValue(val *ContainerNodePoolNodeConfigContainerdConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigContainerdConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigContainerdConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ContainerNodePoolNodeConfigContainerdConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerNodePoolNodeConfigContainerdConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerNodePoolNodeConfigContainerdConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerNodePoolNodeConfigContainerdConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerNodePoolNodeConfigContainerdConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerNodePoolNodeConfigContainerdConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerNodePoolNodeConfigContainerdConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerNodePoolNodeConfigContainerdConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerNodePoolNodeConfigContainerdConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerNodePoolNodeConfigContainerdConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerNodePoolNodeConfigContainerdConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerNodePoolNodeConfigContainerdConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerNodePoolNodeConfigContainerdConfigOutputReference) PutPrivateRegistryAccessConfig(value *ContainerNodePoolNodeConfigContainerdConfigPrivateRegistryAccessConfig) {
	if err := c.validatePutPrivateRegistryAccessConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putPrivateRegistryAccessConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerNodePoolNodeConfigContainerdConfigOutputReference) PutRegistryHosts(value interface{}) {
	if err := c.validatePutRegistryHostsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRegistryHosts",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerNodePoolNodeConfigContainerdConfigOutputReference) PutWritableCgroups(value *ContainerNodePoolNodeConfigContainerdConfigWritableCgroups) {
	if err := c.validatePutWritableCgroupsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putWritableCgroups",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerNodePoolNodeConfigContainerdConfigOutputReference) ResetPrivateRegistryAccessConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetPrivateRegistryAccessConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerNodePoolNodeConfigContainerdConfigOutputReference) ResetRegistryHosts() {
	_jsii_.InvokeVoid(
		c,
		"resetRegistryHosts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerNodePoolNodeConfigContainerdConfigOutputReference) ResetWritableCgroups() {
	_jsii_.InvokeVoid(
		c,
		"resetWritableCgroups",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerNodePoolNodeConfigContainerdConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := c.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerNodePoolNodeConfigContainerdConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

