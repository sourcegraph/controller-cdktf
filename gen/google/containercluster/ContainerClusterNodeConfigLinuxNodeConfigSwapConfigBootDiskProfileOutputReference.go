package containercluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google/containercluster/internal"
)

type ContainerClusterNodeConfigLinuxNodeConfigSwapConfigBootDiskProfileOutputReference interface {
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
	InternalValue() *ContainerClusterNodeConfigLinuxNodeConfigSwapConfigBootDiskProfile
	SetInternalValue(val *ContainerClusterNodeConfigLinuxNodeConfigSwapConfigBootDiskProfile)
	SwapSizeGib() *float64
	SetSwapSizeGib(val *float64)
	SwapSizeGibInput() *float64
	SwapSizePercent() *float64
	SetSwapSizePercent(val *float64)
	SwapSizePercentInput() *float64
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
	ResetSwapSizeGib()
	ResetSwapSizePercent()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ContainerClusterNodeConfigLinuxNodeConfigSwapConfigBootDiskProfileOutputReference
type jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigSwapConfigBootDiskProfileOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigSwapConfigBootDiskProfileOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigSwapConfigBootDiskProfileOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigSwapConfigBootDiskProfileOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigSwapConfigBootDiskProfileOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigSwapConfigBootDiskProfileOutputReference) InternalValue() *ContainerClusterNodeConfigLinuxNodeConfigSwapConfigBootDiskProfile {
	var returns *ContainerClusterNodeConfigLinuxNodeConfigSwapConfigBootDiskProfile
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigSwapConfigBootDiskProfileOutputReference) SwapSizeGib() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"swapSizeGib",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigSwapConfigBootDiskProfileOutputReference) SwapSizeGibInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"swapSizeGibInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigSwapConfigBootDiskProfileOutputReference) SwapSizePercent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"swapSizePercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigSwapConfigBootDiskProfileOutputReference) SwapSizePercentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"swapSizePercentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigSwapConfigBootDiskProfileOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigSwapConfigBootDiskProfileOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewContainerClusterNodeConfigLinuxNodeConfigSwapConfigBootDiskProfileOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ContainerClusterNodeConfigLinuxNodeConfigSwapConfigBootDiskProfileOutputReference {
	_init_.Initialize()

	if err := validateNewContainerClusterNodeConfigLinuxNodeConfigSwapConfigBootDiskProfileOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigSwapConfigBootDiskProfileOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.containerCluster.ContainerClusterNodeConfigLinuxNodeConfigSwapConfigBootDiskProfileOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewContainerClusterNodeConfigLinuxNodeConfigSwapConfigBootDiskProfileOutputReference_Override(c ContainerClusterNodeConfigLinuxNodeConfigSwapConfigBootDiskProfileOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.containerCluster.ContainerClusterNodeConfigLinuxNodeConfigSwapConfigBootDiskProfileOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigSwapConfigBootDiskProfileOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigSwapConfigBootDiskProfileOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigSwapConfigBootDiskProfileOutputReference)SetInternalValue(val *ContainerClusterNodeConfigLinuxNodeConfigSwapConfigBootDiskProfile) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigSwapConfigBootDiskProfileOutputReference)SetSwapSizeGib(val *float64) {
	if err := j.validateSetSwapSizeGibParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"swapSizeGib",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigSwapConfigBootDiskProfileOutputReference)SetSwapSizePercent(val *float64) {
	if err := j.validateSetSwapSizePercentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"swapSizePercent",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigSwapConfigBootDiskProfileOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigSwapConfigBootDiskProfileOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigSwapConfigBootDiskProfileOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigSwapConfigBootDiskProfileOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigSwapConfigBootDiskProfileOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigSwapConfigBootDiskProfileOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigSwapConfigBootDiskProfileOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigSwapConfigBootDiskProfileOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigSwapConfigBootDiskProfileOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigSwapConfigBootDiskProfileOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigSwapConfigBootDiskProfileOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigSwapConfigBootDiskProfileOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigSwapConfigBootDiskProfileOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigSwapConfigBootDiskProfileOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigSwapConfigBootDiskProfileOutputReference) ResetSwapSizeGib() {
	_jsii_.InvokeVoid(
		c,
		"resetSwapSizeGib",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigSwapConfigBootDiskProfileOutputReference) ResetSwapSizePercent() {
	_jsii_.InvokeVoid(
		c,
		"resetSwapSizePercent",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigSwapConfigBootDiskProfileOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigSwapConfigBootDiskProfileOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

