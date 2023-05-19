package googlenotebooksruntime

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlenotebooksruntime/internal"
)

type GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference interface {
	cdktf.ComplexObject
	AutoDelete() cdktf.IResolvable
	Boot() cdktf.IResolvable
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
	DeviceName() *string
	// Experimental.
	Fqn() *string
	GuestOsFeatures() *[]*string
	Index() *float64
	InitializeParams() GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskInitializeParamsOutputReference
	InitializeParamsInput() *GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskInitializeParams
	Interface() *string
	SetInterface(val *string)
	InterfaceInput() *string
	InternalValue() *GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigDataDisk
	SetInternalValue(val *GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigDataDisk)
	Kind() *string
	Licenses() *[]*string
	Mode() *string
	SetMode(val *string)
	ModeInput() *string
	Source() *string
	SetSource(val *string)
	SourceInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
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
	PutInitializeParams(value *GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskInitializeParams)
	ResetInitializeParams()
	ResetInterface()
	ResetMode()
	ResetSource()
	ResetType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference
type jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference) AutoDelete() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"autoDelete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference) Boot() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"boot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference) DeviceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference) GuestOsFeatures() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"guestOsFeatures",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference) Index() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"index",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference) InitializeParams() GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskInitializeParamsOutputReference {
	var returns GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskInitializeParamsOutputReference
	_jsii_.Get(
		j,
		"initializeParams",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference) InitializeParamsInput() *GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskInitializeParams {
	var returns *GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskInitializeParams
	_jsii_.Get(
		j,
		"initializeParamsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference) Interface() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interface",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference) InterfaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interfaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference) InternalValue() *GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigDataDisk {
	var returns *GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigDataDisk
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference) Licenses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"licenses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference) Mode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference) ModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference) Source() *string {
	var returns *string
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference) SourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewGoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference{}

	_jsii_.Create(
		"google-beta.googleNotebooksRuntime.GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference_Override(g GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleNotebooksRuntime.GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference)SetInterface(val *string) {
	if err := j.validateSetInterfaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"interface",
		val,
	)
}

func (j *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference)SetInternalValue(val *GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigDataDisk) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference)SetMode(val *string) {
	if err := j.validateSetModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mode",
		val,
	)
}

func (j *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference)SetSource(val *string) {
	if err := j.validateSetSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"source",
		val,
	)
}

func (j *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (g *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference) PutInitializeParams(value *GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskInitializeParams) {
	if err := g.validatePutInitializeParamsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putInitializeParams",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference) ResetInitializeParams() {
	_jsii_.InvokeVoid(
		g,
		"resetInitializeParams",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference) ResetInterface() {
	_jsii_.InvokeVoid(
		g,
		"resetInterface",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference) ResetMode() {
	_jsii_.InvokeVoid(
		g,
		"resetMode",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference) ResetSource() {
	_jsii_.InvokeVoid(
		g,
		"resetSource",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference) ResetType() {
	_jsii_.InvokeVoid(
		g,
		"resetType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

