package googlecolabnotebookexecution

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecolabnotebookexecution/internal"
)

type GoogleColabNotebookExecutionCustomEnvironmentSpecOutputReference interface {
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
	InternalValue() *GoogleColabNotebookExecutionCustomEnvironmentSpec
	SetInternalValue(val *GoogleColabNotebookExecutionCustomEnvironmentSpec)
	MachineSpec() GoogleColabNotebookExecutionCustomEnvironmentSpecMachineSpecOutputReference
	MachineSpecInput() *GoogleColabNotebookExecutionCustomEnvironmentSpecMachineSpec
	NetworkSpec() GoogleColabNotebookExecutionCustomEnvironmentSpecNetworkSpecOutputReference
	NetworkSpecInput() *GoogleColabNotebookExecutionCustomEnvironmentSpecNetworkSpec
	PersistentDiskSpec() GoogleColabNotebookExecutionCustomEnvironmentSpecPersistentDiskSpecOutputReference
	PersistentDiskSpecInput() *GoogleColabNotebookExecutionCustomEnvironmentSpecPersistentDiskSpec
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
	PutMachineSpec(value *GoogleColabNotebookExecutionCustomEnvironmentSpecMachineSpec)
	PutNetworkSpec(value *GoogleColabNotebookExecutionCustomEnvironmentSpecNetworkSpec)
	PutPersistentDiskSpec(value *GoogleColabNotebookExecutionCustomEnvironmentSpecPersistentDiskSpec)
	ResetMachineSpec()
	ResetNetworkSpec()
	ResetPersistentDiskSpec()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleColabNotebookExecutionCustomEnvironmentSpecOutputReference
type jsiiProxy_GoogleColabNotebookExecutionCustomEnvironmentSpecOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleColabNotebookExecutionCustomEnvironmentSpecOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabNotebookExecutionCustomEnvironmentSpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabNotebookExecutionCustomEnvironmentSpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabNotebookExecutionCustomEnvironmentSpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabNotebookExecutionCustomEnvironmentSpecOutputReference) InternalValue() *GoogleColabNotebookExecutionCustomEnvironmentSpec {
	var returns *GoogleColabNotebookExecutionCustomEnvironmentSpec
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabNotebookExecutionCustomEnvironmentSpecOutputReference) MachineSpec() GoogleColabNotebookExecutionCustomEnvironmentSpecMachineSpecOutputReference {
	var returns GoogleColabNotebookExecutionCustomEnvironmentSpecMachineSpecOutputReference
	_jsii_.Get(
		j,
		"machineSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabNotebookExecutionCustomEnvironmentSpecOutputReference) MachineSpecInput() *GoogleColabNotebookExecutionCustomEnvironmentSpecMachineSpec {
	var returns *GoogleColabNotebookExecutionCustomEnvironmentSpecMachineSpec
	_jsii_.Get(
		j,
		"machineSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabNotebookExecutionCustomEnvironmentSpecOutputReference) NetworkSpec() GoogleColabNotebookExecutionCustomEnvironmentSpecNetworkSpecOutputReference {
	var returns GoogleColabNotebookExecutionCustomEnvironmentSpecNetworkSpecOutputReference
	_jsii_.Get(
		j,
		"networkSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabNotebookExecutionCustomEnvironmentSpecOutputReference) NetworkSpecInput() *GoogleColabNotebookExecutionCustomEnvironmentSpecNetworkSpec {
	var returns *GoogleColabNotebookExecutionCustomEnvironmentSpecNetworkSpec
	_jsii_.Get(
		j,
		"networkSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabNotebookExecutionCustomEnvironmentSpecOutputReference) PersistentDiskSpec() GoogleColabNotebookExecutionCustomEnvironmentSpecPersistentDiskSpecOutputReference {
	var returns GoogleColabNotebookExecutionCustomEnvironmentSpecPersistentDiskSpecOutputReference
	_jsii_.Get(
		j,
		"persistentDiskSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabNotebookExecutionCustomEnvironmentSpecOutputReference) PersistentDiskSpecInput() *GoogleColabNotebookExecutionCustomEnvironmentSpecPersistentDiskSpec {
	var returns *GoogleColabNotebookExecutionCustomEnvironmentSpecPersistentDiskSpec
	_jsii_.Get(
		j,
		"persistentDiskSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabNotebookExecutionCustomEnvironmentSpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleColabNotebookExecutionCustomEnvironmentSpecOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleColabNotebookExecutionCustomEnvironmentSpecOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleColabNotebookExecutionCustomEnvironmentSpecOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleColabNotebookExecutionCustomEnvironmentSpecOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleColabNotebookExecutionCustomEnvironmentSpecOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleColabNotebookExecution.GoogleColabNotebookExecutionCustomEnvironmentSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleColabNotebookExecutionCustomEnvironmentSpecOutputReference_Override(g GoogleColabNotebookExecutionCustomEnvironmentSpecOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleColabNotebookExecution.GoogleColabNotebookExecutionCustomEnvironmentSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleColabNotebookExecutionCustomEnvironmentSpecOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleColabNotebookExecutionCustomEnvironmentSpecOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleColabNotebookExecutionCustomEnvironmentSpecOutputReference)SetInternalValue(val *GoogleColabNotebookExecutionCustomEnvironmentSpec) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleColabNotebookExecutionCustomEnvironmentSpecOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleColabNotebookExecutionCustomEnvironmentSpecOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleColabNotebookExecutionCustomEnvironmentSpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleColabNotebookExecutionCustomEnvironmentSpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleColabNotebookExecutionCustomEnvironmentSpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleColabNotebookExecutionCustomEnvironmentSpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleColabNotebookExecutionCustomEnvironmentSpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleColabNotebookExecutionCustomEnvironmentSpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleColabNotebookExecutionCustomEnvironmentSpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleColabNotebookExecutionCustomEnvironmentSpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleColabNotebookExecutionCustomEnvironmentSpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleColabNotebookExecutionCustomEnvironmentSpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleColabNotebookExecutionCustomEnvironmentSpecOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleColabNotebookExecutionCustomEnvironmentSpecOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleColabNotebookExecutionCustomEnvironmentSpecOutputReference) PutMachineSpec(value *GoogleColabNotebookExecutionCustomEnvironmentSpecMachineSpec) {
	if err := g.validatePutMachineSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMachineSpec",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleColabNotebookExecutionCustomEnvironmentSpecOutputReference) PutNetworkSpec(value *GoogleColabNotebookExecutionCustomEnvironmentSpecNetworkSpec) {
	if err := g.validatePutNetworkSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putNetworkSpec",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleColabNotebookExecutionCustomEnvironmentSpecOutputReference) PutPersistentDiskSpec(value *GoogleColabNotebookExecutionCustomEnvironmentSpecPersistentDiskSpec) {
	if err := g.validatePutPersistentDiskSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPersistentDiskSpec",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleColabNotebookExecutionCustomEnvironmentSpecOutputReference) ResetMachineSpec() {
	_jsii_.InvokeVoid(
		g,
		"resetMachineSpec",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleColabNotebookExecutionCustomEnvironmentSpecOutputReference) ResetNetworkSpec() {
	_jsii_.InvokeVoid(
		g,
		"resetNetworkSpec",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleColabNotebookExecutionCustomEnvironmentSpecOutputReference) ResetPersistentDiskSpec() {
	_jsii_.InvokeVoid(
		g,
		"resetPersistentDiskSpec",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleColabNotebookExecutionCustomEnvironmentSpecOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleColabNotebookExecutionCustomEnvironmentSpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

