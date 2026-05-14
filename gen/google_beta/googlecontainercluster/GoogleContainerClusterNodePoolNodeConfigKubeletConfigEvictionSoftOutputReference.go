package googlecontainercluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecontainercluster/internal"
)

type GoogleContainerClusterNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference interface {
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
	ImagefsAvailable() *string
	SetImagefsAvailable(val *string)
	ImagefsAvailableInput() *string
	ImagefsInodesFree() *string
	SetImagefsInodesFree(val *string)
	ImagefsInodesFreeInput() *string
	InternalValue() *GoogleContainerClusterNodePoolNodeConfigKubeletConfigEvictionSoft
	SetInternalValue(val *GoogleContainerClusterNodePoolNodeConfigKubeletConfigEvictionSoft)
	MemoryAvailable() *string
	SetMemoryAvailable(val *string)
	MemoryAvailableInput() *string
	NodefsAvailable() *string
	SetNodefsAvailable(val *string)
	NodefsAvailableInput() *string
	NodefsInodesFree() *string
	SetNodefsInodesFree(val *string)
	NodefsInodesFreeInput() *string
	PidAvailable() *string
	SetPidAvailable(val *string)
	PidAvailableInput() *string
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
	ResetImagefsAvailable()
	ResetImagefsInodesFree()
	ResetMemoryAvailable()
	ResetNodefsAvailable()
	ResetNodefsInodesFree()
	ResetPidAvailable()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleContainerClusterNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference
type jsiiProxy_GoogleContainerClusterNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference) ImagefsAvailable() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imagefsAvailable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference) ImagefsAvailableInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imagefsAvailableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference) ImagefsInodesFree() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imagefsInodesFree",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference) ImagefsInodesFreeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imagefsInodesFreeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference) InternalValue() *GoogleContainerClusterNodePoolNodeConfigKubeletConfigEvictionSoft {
	var returns *GoogleContainerClusterNodePoolNodeConfigKubeletConfigEvictionSoft
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference) MemoryAvailable() *string {
	var returns *string
	_jsii_.Get(
		j,
		"memoryAvailable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference) MemoryAvailableInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"memoryAvailableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference) NodefsAvailable() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodefsAvailable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference) NodefsAvailableInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodefsAvailableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference) NodefsInodesFree() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodefsInodesFree",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference) NodefsInodesFreeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodefsInodesFreeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference) PidAvailable() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pidAvailable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference) PidAvailableInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pidAvailableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleContainerClusterNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleContainerClusterNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleContainerClusterNodePoolNodeConfigKubeletConfigEvictionSoftOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleContainerClusterNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleContainerCluster.GoogleContainerClusterNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleContainerClusterNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference_Override(g GoogleContainerClusterNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleContainerCluster.GoogleContainerClusterNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference)SetImagefsAvailable(val *string) {
	if err := j.validateSetImagefsAvailableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imagefsAvailable",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference)SetImagefsInodesFree(val *string) {
	if err := j.validateSetImagefsInodesFreeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imagefsInodesFree",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference)SetInternalValue(val *GoogleContainerClusterNodePoolNodeConfigKubeletConfigEvictionSoft) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference)SetMemoryAvailable(val *string) {
	if err := j.validateSetMemoryAvailableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memoryAvailable",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference)SetNodefsAvailable(val *string) {
	if err := j.validateSetNodefsAvailableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodefsAvailable",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference)SetNodefsInodesFree(val *string) {
	if err := j.validateSetNodefsInodesFreeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodefsInodesFree",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference)SetPidAvailable(val *string) {
	if err := j.validateSetPidAvailableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pidAvailable",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference) ResetImagefsAvailable() {
	_jsii_.InvokeVoid(
		g,
		"resetImagefsAvailable",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference) ResetImagefsInodesFree() {
	_jsii_.InvokeVoid(
		g,
		"resetImagefsInodesFree",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference) ResetMemoryAvailable() {
	_jsii_.InvokeVoid(
		g,
		"resetMemoryAvailable",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference) ResetNodefsAvailable() {
	_jsii_.InvokeVoid(
		g,
		"resetNodefsAvailable",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference) ResetNodefsInodesFree() {
	_jsii_.InvokeVoid(
		g,
		"resetNodefsInodesFree",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference) ResetPidAvailable() {
	_jsii_.InvokeVoid(
		g,
		"resetPidAvailable",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

