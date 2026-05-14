package googlecontainercluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecontainercluster/internal"
)

type GoogleContainerClusterNodeConfigKubeletConfigEvictionMinimumReclaimOutputReference interface {
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
	InternalValue() *GoogleContainerClusterNodeConfigKubeletConfigEvictionMinimumReclaim
	SetInternalValue(val *GoogleContainerClusterNodeConfigKubeletConfigEvictionMinimumReclaim)
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

// The jsii proxy struct for GoogleContainerClusterNodeConfigKubeletConfigEvictionMinimumReclaimOutputReference
type jsiiProxy_GoogleContainerClusterNodeConfigKubeletConfigEvictionMinimumReclaimOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigKubeletConfigEvictionMinimumReclaimOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigKubeletConfigEvictionMinimumReclaimOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigKubeletConfigEvictionMinimumReclaimOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigKubeletConfigEvictionMinimumReclaimOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigKubeletConfigEvictionMinimumReclaimOutputReference) ImagefsAvailable() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imagefsAvailable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigKubeletConfigEvictionMinimumReclaimOutputReference) ImagefsAvailableInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imagefsAvailableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigKubeletConfigEvictionMinimumReclaimOutputReference) ImagefsInodesFree() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imagefsInodesFree",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigKubeletConfigEvictionMinimumReclaimOutputReference) ImagefsInodesFreeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imagefsInodesFreeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigKubeletConfigEvictionMinimumReclaimOutputReference) InternalValue() *GoogleContainerClusterNodeConfigKubeletConfigEvictionMinimumReclaim {
	var returns *GoogleContainerClusterNodeConfigKubeletConfigEvictionMinimumReclaim
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigKubeletConfigEvictionMinimumReclaimOutputReference) MemoryAvailable() *string {
	var returns *string
	_jsii_.Get(
		j,
		"memoryAvailable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigKubeletConfigEvictionMinimumReclaimOutputReference) MemoryAvailableInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"memoryAvailableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigKubeletConfigEvictionMinimumReclaimOutputReference) NodefsAvailable() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodefsAvailable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigKubeletConfigEvictionMinimumReclaimOutputReference) NodefsAvailableInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodefsAvailableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigKubeletConfigEvictionMinimumReclaimOutputReference) NodefsInodesFree() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodefsInodesFree",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigKubeletConfigEvictionMinimumReclaimOutputReference) NodefsInodesFreeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodefsInodesFreeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigKubeletConfigEvictionMinimumReclaimOutputReference) PidAvailable() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pidAvailable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigKubeletConfigEvictionMinimumReclaimOutputReference) PidAvailableInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pidAvailableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigKubeletConfigEvictionMinimumReclaimOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigKubeletConfigEvictionMinimumReclaimOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleContainerClusterNodeConfigKubeletConfigEvictionMinimumReclaimOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleContainerClusterNodeConfigKubeletConfigEvictionMinimumReclaimOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleContainerClusterNodeConfigKubeletConfigEvictionMinimumReclaimOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleContainerClusterNodeConfigKubeletConfigEvictionMinimumReclaimOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleContainerCluster.GoogleContainerClusterNodeConfigKubeletConfigEvictionMinimumReclaimOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleContainerClusterNodeConfigKubeletConfigEvictionMinimumReclaimOutputReference_Override(g GoogleContainerClusterNodeConfigKubeletConfigEvictionMinimumReclaimOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleContainerCluster.GoogleContainerClusterNodeConfigKubeletConfigEvictionMinimumReclaimOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigKubeletConfigEvictionMinimumReclaimOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigKubeletConfigEvictionMinimumReclaimOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigKubeletConfigEvictionMinimumReclaimOutputReference)SetImagefsAvailable(val *string) {
	if err := j.validateSetImagefsAvailableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imagefsAvailable",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigKubeletConfigEvictionMinimumReclaimOutputReference)SetImagefsInodesFree(val *string) {
	if err := j.validateSetImagefsInodesFreeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imagefsInodesFree",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigKubeletConfigEvictionMinimumReclaimOutputReference)SetInternalValue(val *GoogleContainerClusterNodeConfigKubeletConfigEvictionMinimumReclaim) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigKubeletConfigEvictionMinimumReclaimOutputReference)SetMemoryAvailable(val *string) {
	if err := j.validateSetMemoryAvailableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memoryAvailable",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigKubeletConfigEvictionMinimumReclaimOutputReference)SetNodefsAvailable(val *string) {
	if err := j.validateSetNodefsAvailableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodefsAvailable",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigKubeletConfigEvictionMinimumReclaimOutputReference)SetNodefsInodesFree(val *string) {
	if err := j.validateSetNodefsInodesFreeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodefsInodesFree",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigKubeletConfigEvictionMinimumReclaimOutputReference)SetPidAvailable(val *string) {
	if err := j.validateSetPidAvailableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pidAvailable",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigKubeletConfigEvictionMinimumReclaimOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigKubeletConfigEvictionMinimumReclaimOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodeConfigKubeletConfigEvictionMinimumReclaimOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerClusterNodeConfigKubeletConfigEvictionMinimumReclaimOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleContainerClusterNodeConfigKubeletConfigEvictionMinimumReclaimOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleContainerClusterNodeConfigKubeletConfigEvictionMinimumReclaimOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleContainerClusterNodeConfigKubeletConfigEvictionMinimumReclaimOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleContainerClusterNodeConfigKubeletConfigEvictionMinimumReclaimOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleContainerClusterNodeConfigKubeletConfigEvictionMinimumReclaimOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleContainerClusterNodeConfigKubeletConfigEvictionMinimumReclaimOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleContainerClusterNodeConfigKubeletConfigEvictionMinimumReclaimOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleContainerClusterNodeConfigKubeletConfigEvictionMinimumReclaimOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleContainerClusterNodeConfigKubeletConfigEvictionMinimumReclaimOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerClusterNodeConfigKubeletConfigEvictionMinimumReclaimOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleContainerClusterNodeConfigKubeletConfigEvictionMinimumReclaimOutputReference) ResetImagefsAvailable() {
	_jsii_.InvokeVoid(
		g,
		"resetImagefsAvailable",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodeConfigKubeletConfigEvictionMinimumReclaimOutputReference) ResetImagefsInodesFree() {
	_jsii_.InvokeVoid(
		g,
		"resetImagefsInodesFree",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodeConfigKubeletConfigEvictionMinimumReclaimOutputReference) ResetMemoryAvailable() {
	_jsii_.InvokeVoid(
		g,
		"resetMemoryAvailable",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodeConfigKubeletConfigEvictionMinimumReclaimOutputReference) ResetNodefsAvailable() {
	_jsii_.InvokeVoid(
		g,
		"resetNodefsAvailable",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodeConfigKubeletConfigEvictionMinimumReclaimOutputReference) ResetNodefsInodesFree() {
	_jsii_.InvokeVoid(
		g,
		"resetNodefsInodesFree",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodeConfigKubeletConfigEvictionMinimumReclaimOutputReference) ResetPidAvailable() {
	_jsii_.InvokeVoid(
		g,
		"resetPidAvailable",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodeConfigKubeletConfigEvictionMinimumReclaimOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleContainerClusterNodeConfigKubeletConfigEvictionMinimumReclaimOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

