package dialogflowcxtoolversion

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google/dialogflowcxtoolversion/internal"
)

type DialogflowCxToolVersionToolOutputReference interface {
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
	DataStoreSpec() DialogflowCxToolVersionToolDataStoreSpecOutputReference
	DataStoreSpecInput() *DialogflowCxToolVersionToolDataStoreSpec
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	// Experimental.
	Fqn() *string
	FunctionSpec() DialogflowCxToolVersionToolFunctionSpecOutputReference
	FunctionSpecInput() *DialogflowCxToolVersionToolFunctionSpec
	InternalValue() *DialogflowCxToolVersionTool
	SetInternalValue(val *DialogflowCxToolVersionTool)
	Name() *string
	OpenApiSpec() DialogflowCxToolVersionToolOpenApiSpecOutputReference
	OpenApiSpecInput() *DialogflowCxToolVersionToolOpenApiSpec
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	ToolType() *string
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
	PutDataStoreSpec(value *DialogflowCxToolVersionToolDataStoreSpec)
	PutFunctionSpec(value *DialogflowCxToolVersionToolFunctionSpec)
	PutOpenApiSpec(value *DialogflowCxToolVersionToolOpenApiSpec)
	ResetDataStoreSpec()
	ResetFunctionSpec()
	ResetOpenApiSpec()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DialogflowCxToolVersionToolOutputReference
type jsiiProxy_DialogflowCxToolVersionToolOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DialogflowCxToolVersionToolOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxToolVersionToolOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxToolVersionToolOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxToolVersionToolOutputReference) DataStoreSpec() DialogflowCxToolVersionToolDataStoreSpecOutputReference {
	var returns DialogflowCxToolVersionToolDataStoreSpecOutputReference
	_jsii_.Get(
		j,
		"dataStoreSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxToolVersionToolOutputReference) DataStoreSpecInput() *DialogflowCxToolVersionToolDataStoreSpec {
	var returns *DialogflowCxToolVersionToolDataStoreSpec
	_jsii_.Get(
		j,
		"dataStoreSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxToolVersionToolOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxToolVersionToolOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxToolVersionToolOutputReference) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxToolVersionToolOutputReference) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxToolVersionToolOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxToolVersionToolOutputReference) FunctionSpec() DialogflowCxToolVersionToolFunctionSpecOutputReference {
	var returns DialogflowCxToolVersionToolFunctionSpecOutputReference
	_jsii_.Get(
		j,
		"functionSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxToolVersionToolOutputReference) FunctionSpecInput() *DialogflowCxToolVersionToolFunctionSpec {
	var returns *DialogflowCxToolVersionToolFunctionSpec
	_jsii_.Get(
		j,
		"functionSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxToolVersionToolOutputReference) InternalValue() *DialogflowCxToolVersionTool {
	var returns *DialogflowCxToolVersionTool
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxToolVersionToolOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxToolVersionToolOutputReference) OpenApiSpec() DialogflowCxToolVersionToolOpenApiSpecOutputReference {
	var returns DialogflowCxToolVersionToolOpenApiSpecOutputReference
	_jsii_.Get(
		j,
		"openApiSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxToolVersionToolOutputReference) OpenApiSpecInput() *DialogflowCxToolVersionToolOpenApiSpec {
	var returns *DialogflowCxToolVersionToolOpenApiSpec
	_jsii_.Get(
		j,
		"openApiSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxToolVersionToolOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxToolVersionToolOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxToolVersionToolOutputReference) ToolType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"toolType",
		&returns,
	)
	return returns
}


func NewDialogflowCxToolVersionToolOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DialogflowCxToolVersionToolOutputReference {
	_init_.Initialize()

	if err := validateNewDialogflowCxToolVersionToolOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DialogflowCxToolVersionToolOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.dialogflowCxToolVersion.DialogflowCxToolVersionToolOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDialogflowCxToolVersionToolOutputReference_Override(d DialogflowCxToolVersionToolOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.dialogflowCxToolVersion.DialogflowCxToolVersionToolOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DialogflowCxToolVersionToolOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxToolVersionToolOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxToolVersionToolOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxToolVersionToolOutputReference)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxToolVersionToolOutputReference)SetInternalValue(val *DialogflowCxToolVersionTool) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxToolVersionToolOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxToolVersionToolOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DialogflowCxToolVersionToolOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxToolVersionToolOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxToolVersionToolOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxToolVersionToolOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxToolVersionToolOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxToolVersionToolOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxToolVersionToolOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxToolVersionToolOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxToolVersionToolOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxToolVersionToolOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxToolVersionToolOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxToolVersionToolOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxToolVersionToolOutputReference) PutDataStoreSpec(value *DialogflowCxToolVersionToolDataStoreSpec) {
	if err := d.validatePutDataStoreSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putDataStoreSpec",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DialogflowCxToolVersionToolOutputReference) PutFunctionSpec(value *DialogflowCxToolVersionToolFunctionSpec) {
	if err := d.validatePutFunctionSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putFunctionSpec",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DialogflowCxToolVersionToolOutputReference) PutOpenApiSpec(value *DialogflowCxToolVersionToolOpenApiSpec) {
	if err := d.validatePutOpenApiSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putOpenApiSpec",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DialogflowCxToolVersionToolOutputReference) ResetDataStoreSpec() {
	_jsii_.InvokeVoid(
		d,
		"resetDataStoreSpec",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxToolVersionToolOutputReference) ResetFunctionSpec() {
	_jsii_.InvokeVoid(
		d,
		"resetFunctionSpec",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxToolVersionToolOutputReference) ResetOpenApiSpec() {
	_jsii_.InvokeVoid(
		d,
		"resetOpenApiSpec",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxToolVersionToolOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxToolVersionToolOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

