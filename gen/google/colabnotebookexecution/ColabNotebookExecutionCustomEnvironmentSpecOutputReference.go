package colabnotebookexecution

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/colabnotebookexecution/internal"
)

type ColabNotebookExecutionCustomEnvironmentSpecOutputReference interface {
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
	InternalValue() *ColabNotebookExecutionCustomEnvironmentSpec
	SetInternalValue(val *ColabNotebookExecutionCustomEnvironmentSpec)
	MachineSpec() ColabNotebookExecutionCustomEnvironmentSpecMachineSpecOutputReference
	MachineSpecInput() *ColabNotebookExecutionCustomEnvironmentSpecMachineSpec
	NetworkSpec() ColabNotebookExecutionCustomEnvironmentSpecNetworkSpecOutputReference
	NetworkSpecInput() *ColabNotebookExecutionCustomEnvironmentSpecNetworkSpec
	PersistentDiskSpec() ColabNotebookExecutionCustomEnvironmentSpecPersistentDiskSpecOutputReference
	PersistentDiskSpecInput() *ColabNotebookExecutionCustomEnvironmentSpecPersistentDiskSpec
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
	PutMachineSpec(value *ColabNotebookExecutionCustomEnvironmentSpecMachineSpec)
	PutNetworkSpec(value *ColabNotebookExecutionCustomEnvironmentSpecNetworkSpec)
	PutPersistentDiskSpec(value *ColabNotebookExecutionCustomEnvironmentSpecPersistentDiskSpec)
	ResetMachineSpec()
	ResetNetworkSpec()
	ResetPersistentDiskSpec()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ColabNotebookExecutionCustomEnvironmentSpecOutputReference
type jsiiProxy_ColabNotebookExecutionCustomEnvironmentSpecOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ColabNotebookExecutionCustomEnvironmentSpecOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabNotebookExecutionCustomEnvironmentSpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabNotebookExecutionCustomEnvironmentSpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabNotebookExecutionCustomEnvironmentSpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabNotebookExecutionCustomEnvironmentSpecOutputReference) InternalValue() *ColabNotebookExecutionCustomEnvironmentSpec {
	var returns *ColabNotebookExecutionCustomEnvironmentSpec
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabNotebookExecutionCustomEnvironmentSpecOutputReference) MachineSpec() ColabNotebookExecutionCustomEnvironmentSpecMachineSpecOutputReference {
	var returns ColabNotebookExecutionCustomEnvironmentSpecMachineSpecOutputReference
	_jsii_.Get(
		j,
		"machineSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabNotebookExecutionCustomEnvironmentSpecOutputReference) MachineSpecInput() *ColabNotebookExecutionCustomEnvironmentSpecMachineSpec {
	var returns *ColabNotebookExecutionCustomEnvironmentSpecMachineSpec
	_jsii_.Get(
		j,
		"machineSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabNotebookExecutionCustomEnvironmentSpecOutputReference) NetworkSpec() ColabNotebookExecutionCustomEnvironmentSpecNetworkSpecOutputReference {
	var returns ColabNotebookExecutionCustomEnvironmentSpecNetworkSpecOutputReference
	_jsii_.Get(
		j,
		"networkSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabNotebookExecutionCustomEnvironmentSpecOutputReference) NetworkSpecInput() *ColabNotebookExecutionCustomEnvironmentSpecNetworkSpec {
	var returns *ColabNotebookExecutionCustomEnvironmentSpecNetworkSpec
	_jsii_.Get(
		j,
		"networkSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabNotebookExecutionCustomEnvironmentSpecOutputReference) PersistentDiskSpec() ColabNotebookExecutionCustomEnvironmentSpecPersistentDiskSpecOutputReference {
	var returns ColabNotebookExecutionCustomEnvironmentSpecPersistentDiskSpecOutputReference
	_jsii_.Get(
		j,
		"persistentDiskSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabNotebookExecutionCustomEnvironmentSpecOutputReference) PersistentDiskSpecInput() *ColabNotebookExecutionCustomEnvironmentSpecPersistentDiskSpec {
	var returns *ColabNotebookExecutionCustomEnvironmentSpecPersistentDiskSpec
	_jsii_.Get(
		j,
		"persistentDiskSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabNotebookExecutionCustomEnvironmentSpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabNotebookExecutionCustomEnvironmentSpecOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewColabNotebookExecutionCustomEnvironmentSpecOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ColabNotebookExecutionCustomEnvironmentSpecOutputReference {
	_init_.Initialize()

	if err := validateNewColabNotebookExecutionCustomEnvironmentSpecOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ColabNotebookExecutionCustomEnvironmentSpecOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.colabNotebookExecution.ColabNotebookExecutionCustomEnvironmentSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewColabNotebookExecutionCustomEnvironmentSpecOutputReference_Override(c ColabNotebookExecutionCustomEnvironmentSpecOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.colabNotebookExecution.ColabNotebookExecutionCustomEnvironmentSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ColabNotebookExecutionCustomEnvironmentSpecOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ColabNotebookExecutionCustomEnvironmentSpecOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ColabNotebookExecutionCustomEnvironmentSpecOutputReference)SetInternalValue(val *ColabNotebookExecutionCustomEnvironmentSpec) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ColabNotebookExecutionCustomEnvironmentSpecOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ColabNotebookExecutionCustomEnvironmentSpecOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ColabNotebookExecutionCustomEnvironmentSpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ColabNotebookExecutionCustomEnvironmentSpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ColabNotebookExecutionCustomEnvironmentSpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ColabNotebookExecutionCustomEnvironmentSpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ColabNotebookExecutionCustomEnvironmentSpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ColabNotebookExecutionCustomEnvironmentSpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ColabNotebookExecutionCustomEnvironmentSpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ColabNotebookExecutionCustomEnvironmentSpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ColabNotebookExecutionCustomEnvironmentSpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ColabNotebookExecutionCustomEnvironmentSpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ColabNotebookExecutionCustomEnvironmentSpecOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ColabNotebookExecutionCustomEnvironmentSpecOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ColabNotebookExecutionCustomEnvironmentSpecOutputReference) PutMachineSpec(value *ColabNotebookExecutionCustomEnvironmentSpecMachineSpec) {
	if err := c.validatePutMachineSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putMachineSpec",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ColabNotebookExecutionCustomEnvironmentSpecOutputReference) PutNetworkSpec(value *ColabNotebookExecutionCustomEnvironmentSpecNetworkSpec) {
	if err := c.validatePutNetworkSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putNetworkSpec",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ColabNotebookExecutionCustomEnvironmentSpecOutputReference) PutPersistentDiskSpec(value *ColabNotebookExecutionCustomEnvironmentSpecPersistentDiskSpec) {
	if err := c.validatePutPersistentDiskSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putPersistentDiskSpec",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ColabNotebookExecutionCustomEnvironmentSpecOutputReference) ResetMachineSpec() {
	_jsii_.InvokeVoid(
		c,
		"resetMachineSpec",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ColabNotebookExecutionCustomEnvironmentSpecOutputReference) ResetNetworkSpec() {
	_jsii_.InvokeVoid(
		c,
		"resetNetworkSpec",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ColabNotebookExecutionCustomEnvironmentSpecOutputReference) ResetPersistentDiskSpec() {
	_jsii_.InvokeVoid(
		c,
		"resetPersistentDiskSpec",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ColabNotebookExecutionCustomEnvironmentSpecOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ColabNotebookExecutionCustomEnvironmentSpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

