package googlebigqueryroutine

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlebigqueryroutine/internal"
)

type GoogleBigqueryRoutineExternalRuntimeOptionsOutputReference interface {
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
	ContainerCpu() *float64
	SetContainerCpu(val *float64)
	ContainerCpuInput() *float64
	ContainerMemory() *string
	SetContainerMemory(val *string)
	ContainerMemoryInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleBigqueryRoutineExternalRuntimeOptions
	SetInternalValue(val *GoogleBigqueryRoutineExternalRuntimeOptions)
	MaxBatchingRows() *string
	SetMaxBatchingRows(val *string)
	MaxBatchingRowsInput() *string
	RuntimeConnection() *string
	SetRuntimeConnection(val *string)
	RuntimeConnectionInput() *string
	RuntimeVersion() *string
	SetRuntimeVersion(val *string)
	RuntimeVersionInput() *string
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
	ResetContainerCpu()
	ResetContainerMemory()
	ResetMaxBatchingRows()
	ResetRuntimeConnection()
	ResetRuntimeVersion()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleBigqueryRoutineExternalRuntimeOptionsOutputReference
type jsiiProxy_GoogleBigqueryRoutineExternalRuntimeOptionsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleBigqueryRoutineExternalRuntimeOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryRoutineExternalRuntimeOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryRoutineExternalRuntimeOptionsOutputReference) ContainerCpu() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"containerCpu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryRoutineExternalRuntimeOptionsOutputReference) ContainerCpuInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"containerCpuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryRoutineExternalRuntimeOptionsOutputReference) ContainerMemory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerMemory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryRoutineExternalRuntimeOptionsOutputReference) ContainerMemoryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerMemoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryRoutineExternalRuntimeOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryRoutineExternalRuntimeOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryRoutineExternalRuntimeOptionsOutputReference) InternalValue() *GoogleBigqueryRoutineExternalRuntimeOptions {
	var returns *GoogleBigqueryRoutineExternalRuntimeOptions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryRoutineExternalRuntimeOptionsOutputReference) MaxBatchingRows() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxBatchingRows",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryRoutineExternalRuntimeOptionsOutputReference) MaxBatchingRowsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxBatchingRowsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryRoutineExternalRuntimeOptionsOutputReference) RuntimeConnection() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeConnection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryRoutineExternalRuntimeOptionsOutputReference) RuntimeConnectionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeConnectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryRoutineExternalRuntimeOptionsOutputReference) RuntimeVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryRoutineExternalRuntimeOptionsOutputReference) RuntimeVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryRoutineExternalRuntimeOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryRoutineExternalRuntimeOptionsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleBigqueryRoutineExternalRuntimeOptionsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleBigqueryRoutineExternalRuntimeOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleBigqueryRoutineExternalRuntimeOptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleBigqueryRoutineExternalRuntimeOptionsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleBigqueryRoutine.GoogleBigqueryRoutineExternalRuntimeOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleBigqueryRoutineExternalRuntimeOptionsOutputReference_Override(g GoogleBigqueryRoutineExternalRuntimeOptionsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleBigqueryRoutine.GoogleBigqueryRoutineExternalRuntimeOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleBigqueryRoutineExternalRuntimeOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryRoutineExternalRuntimeOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryRoutineExternalRuntimeOptionsOutputReference)SetContainerCpu(val *float64) {
	if err := j.validateSetContainerCpuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerCpu",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryRoutineExternalRuntimeOptionsOutputReference)SetContainerMemory(val *string) {
	if err := j.validateSetContainerMemoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerMemory",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryRoutineExternalRuntimeOptionsOutputReference)SetInternalValue(val *GoogleBigqueryRoutineExternalRuntimeOptions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryRoutineExternalRuntimeOptionsOutputReference)SetMaxBatchingRows(val *string) {
	if err := j.validateSetMaxBatchingRowsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxBatchingRows",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryRoutineExternalRuntimeOptionsOutputReference)SetRuntimeConnection(val *string) {
	if err := j.validateSetRuntimeConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runtimeConnection",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryRoutineExternalRuntimeOptionsOutputReference)SetRuntimeVersion(val *string) {
	if err := j.validateSetRuntimeVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runtimeVersion",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryRoutineExternalRuntimeOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryRoutineExternalRuntimeOptionsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleBigqueryRoutineExternalRuntimeOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigqueryRoutineExternalRuntimeOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleBigqueryRoutineExternalRuntimeOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleBigqueryRoutineExternalRuntimeOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleBigqueryRoutineExternalRuntimeOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleBigqueryRoutineExternalRuntimeOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleBigqueryRoutineExternalRuntimeOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleBigqueryRoutineExternalRuntimeOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleBigqueryRoutineExternalRuntimeOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleBigqueryRoutineExternalRuntimeOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleBigqueryRoutineExternalRuntimeOptionsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigqueryRoutineExternalRuntimeOptionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleBigqueryRoutineExternalRuntimeOptionsOutputReference) ResetContainerCpu() {
	_jsii_.InvokeVoid(
		g,
		"resetContainerCpu",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryRoutineExternalRuntimeOptionsOutputReference) ResetContainerMemory() {
	_jsii_.InvokeVoid(
		g,
		"resetContainerMemory",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryRoutineExternalRuntimeOptionsOutputReference) ResetMaxBatchingRows() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxBatchingRows",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryRoutineExternalRuntimeOptionsOutputReference) ResetRuntimeConnection() {
	_jsii_.InvokeVoid(
		g,
		"resetRuntimeConnection",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryRoutineExternalRuntimeOptionsOutputReference) ResetRuntimeVersion() {
	_jsii_.InvokeVoid(
		g,
		"resetRuntimeVersion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryRoutineExternalRuntimeOptionsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleBigqueryRoutineExternalRuntimeOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

