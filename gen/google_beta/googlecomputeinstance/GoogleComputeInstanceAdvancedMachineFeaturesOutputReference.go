package googlecomputeinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecomputeinstance/internal"
)

type GoogleComputeInstanceAdvancedMachineFeaturesOutputReference interface {
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
	EnableNestedVirtualization() interface{}
	SetEnableNestedVirtualization(val interface{})
	EnableNestedVirtualizationInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleComputeInstanceAdvancedMachineFeatures
	SetInternalValue(val *GoogleComputeInstanceAdvancedMachineFeatures)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	ThreadsPerCore() *float64
	SetThreadsPerCore(val *float64)
	ThreadsPerCoreInput() *float64
	VisibleCoreCount() *float64
	SetVisibleCoreCount(val *float64)
	VisibleCoreCountInput() *float64
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
	ResetEnableNestedVirtualization()
	ResetThreadsPerCore()
	ResetVisibleCoreCount()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleComputeInstanceAdvancedMachineFeaturesOutputReference
type jsiiProxy_GoogleComputeInstanceAdvancedMachineFeaturesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleComputeInstanceAdvancedMachineFeaturesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceAdvancedMachineFeaturesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceAdvancedMachineFeaturesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceAdvancedMachineFeaturesOutputReference) EnableNestedVirtualization() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableNestedVirtualization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceAdvancedMachineFeaturesOutputReference) EnableNestedVirtualizationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableNestedVirtualizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceAdvancedMachineFeaturesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceAdvancedMachineFeaturesOutputReference) InternalValue() *GoogleComputeInstanceAdvancedMachineFeatures {
	var returns *GoogleComputeInstanceAdvancedMachineFeatures
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceAdvancedMachineFeaturesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceAdvancedMachineFeaturesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceAdvancedMachineFeaturesOutputReference) ThreadsPerCore() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"threadsPerCore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceAdvancedMachineFeaturesOutputReference) ThreadsPerCoreInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"threadsPerCoreInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceAdvancedMachineFeaturesOutputReference) VisibleCoreCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"visibleCoreCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceAdvancedMachineFeaturesOutputReference) VisibleCoreCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"visibleCoreCountInput",
		&returns,
	)
	return returns
}


func NewGoogleComputeInstanceAdvancedMachineFeaturesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleComputeInstanceAdvancedMachineFeaturesOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleComputeInstanceAdvancedMachineFeaturesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeInstanceAdvancedMachineFeaturesOutputReference{}

	_jsii_.Create(
		"google-beta.googleComputeInstance.GoogleComputeInstanceAdvancedMachineFeaturesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleComputeInstanceAdvancedMachineFeaturesOutputReference_Override(g GoogleComputeInstanceAdvancedMachineFeaturesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleComputeInstance.GoogleComputeInstanceAdvancedMachineFeaturesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceAdvancedMachineFeaturesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceAdvancedMachineFeaturesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceAdvancedMachineFeaturesOutputReference)SetEnableNestedVirtualization(val interface{}) {
	if err := j.validateSetEnableNestedVirtualizationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableNestedVirtualization",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceAdvancedMachineFeaturesOutputReference)SetInternalValue(val *GoogleComputeInstanceAdvancedMachineFeatures) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceAdvancedMachineFeaturesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceAdvancedMachineFeaturesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceAdvancedMachineFeaturesOutputReference)SetThreadsPerCore(val *float64) {
	if err := j.validateSetThreadsPerCoreParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"threadsPerCore",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceAdvancedMachineFeaturesOutputReference)SetVisibleCoreCount(val *float64) {
	if err := j.validateSetVisibleCoreCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"visibleCoreCount",
		val,
	)
}

func (g *jsiiProxy_GoogleComputeInstanceAdvancedMachineFeaturesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeInstanceAdvancedMachineFeaturesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleComputeInstanceAdvancedMachineFeaturesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeInstanceAdvancedMachineFeaturesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComputeInstanceAdvancedMachineFeaturesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComputeInstanceAdvancedMachineFeaturesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComputeInstanceAdvancedMachineFeaturesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComputeInstanceAdvancedMachineFeaturesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComputeInstanceAdvancedMachineFeaturesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComputeInstanceAdvancedMachineFeaturesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComputeInstanceAdvancedMachineFeaturesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeInstanceAdvancedMachineFeaturesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeInstanceAdvancedMachineFeaturesOutputReference) ResetEnableNestedVirtualization() {
	_jsii_.InvokeVoid(
		g,
		"resetEnableNestedVirtualization",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceAdvancedMachineFeaturesOutputReference) ResetThreadsPerCore() {
	_jsii_.InvokeVoid(
		g,
		"resetThreadsPerCore",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceAdvancedMachineFeaturesOutputReference) ResetVisibleCoreCount() {
	_jsii_.InvokeVoid(
		g,
		"resetVisibleCoreCount",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceAdvancedMachineFeaturesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleComputeInstanceAdvancedMachineFeaturesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

