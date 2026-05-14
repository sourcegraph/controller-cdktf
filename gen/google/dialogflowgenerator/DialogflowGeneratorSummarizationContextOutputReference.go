package dialogflowgenerator

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/dialogflowgenerator/internal"
)

type DialogflowGeneratorSummarizationContextOutputReference interface {
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
	FewShotExamples() DialogflowGeneratorSummarizationContextFewShotExamplesList
	FewShotExamplesInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *DialogflowGeneratorSummarizationContext
	SetInternalValue(val *DialogflowGeneratorSummarizationContext)
	OutputLanguageCode() *string
	SetOutputLanguageCode(val *string)
	OutputLanguageCodeInput() *string
	SummarizationSections() DialogflowGeneratorSummarizationContextSummarizationSectionsList
	SummarizationSectionsInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Version() *string
	SetVersion(val *string)
	VersionInput() *string
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
	PutFewShotExamples(value interface{})
	PutSummarizationSections(value interface{})
	ResetFewShotExamples()
	ResetOutputLanguageCode()
	ResetSummarizationSections()
	ResetVersion()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DialogflowGeneratorSummarizationContextOutputReference
type jsiiProxy_DialogflowGeneratorSummarizationContextOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DialogflowGeneratorSummarizationContextOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowGeneratorSummarizationContextOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowGeneratorSummarizationContextOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowGeneratorSummarizationContextOutputReference) FewShotExamples() DialogflowGeneratorSummarizationContextFewShotExamplesList {
	var returns DialogflowGeneratorSummarizationContextFewShotExamplesList
	_jsii_.Get(
		j,
		"fewShotExamples",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowGeneratorSummarizationContextOutputReference) FewShotExamplesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fewShotExamplesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowGeneratorSummarizationContextOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowGeneratorSummarizationContextOutputReference) InternalValue() *DialogflowGeneratorSummarizationContext {
	var returns *DialogflowGeneratorSummarizationContext
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowGeneratorSummarizationContextOutputReference) OutputLanguageCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputLanguageCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowGeneratorSummarizationContextOutputReference) OutputLanguageCodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputLanguageCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowGeneratorSummarizationContextOutputReference) SummarizationSections() DialogflowGeneratorSummarizationContextSummarizationSectionsList {
	var returns DialogflowGeneratorSummarizationContextSummarizationSectionsList
	_jsii_.Get(
		j,
		"summarizationSections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowGeneratorSummarizationContextOutputReference) SummarizationSectionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"summarizationSectionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowGeneratorSummarizationContextOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowGeneratorSummarizationContextOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowGeneratorSummarizationContextOutputReference) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowGeneratorSummarizationContextOutputReference) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}


func NewDialogflowGeneratorSummarizationContextOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DialogflowGeneratorSummarizationContextOutputReference {
	_init_.Initialize()

	if err := validateNewDialogflowGeneratorSummarizationContextOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DialogflowGeneratorSummarizationContextOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dialogflowGenerator.DialogflowGeneratorSummarizationContextOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDialogflowGeneratorSummarizationContextOutputReference_Override(d DialogflowGeneratorSummarizationContextOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dialogflowGenerator.DialogflowGeneratorSummarizationContextOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DialogflowGeneratorSummarizationContextOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DialogflowGeneratorSummarizationContextOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DialogflowGeneratorSummarizationContextOutputReference)SetInternalValue(val *DialogflowGeneratorSummarizationContext) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DialogflowGeneratorSummarizationContextOutputReference)SetOutputLanguageCode(val *string) {
	if err := j.validateSetOutputLanguageCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outputLanguageCode",
		val,
	)
}

func (j *jsiiProxy_DialogflowGeneratorSummarizationContextOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DialogflowGeneratorSummarizationContextOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DialogflowGeneratorSummarizationContextOutputReference)SetVersion(val *string) {
	if err := j.validateSetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

func (d *jsiiProxy_DialogflowGeneratorSummarizationContextOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowGeneratorSummarizationContextOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DialogflowGeneratorSummarizationContextOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowGeneratorSummarizationContextOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DialogflowGeneratorSummarizationContextOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DialogflowGeneratorSummarizationContextOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DialogflowGeneratorSummarizationContextOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DialogflowGeneratorSummarizationContextOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DialogflowGeneratorSummarizationContextOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DialogflowGeneratorSummarizationContextOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DialogflowGeneratorSummarizationContextOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowGeneratorSummarizationContextOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowGeneratorSummarizationContextOutputReference) PutFewShotExamples(value interface{}) {
	if err := d.validatePutFewShotExamplesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putFewShotExamples",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DialogflowGeneratorSummarizationContextOutputReference) PutSummarizationSections(value interface{}) {
	if err := d.validatePutSummarizationSectionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSummarizationSections",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DialogflowGeneratorSummarizationContextOutputReference) ResetFewShotExamples() {
	_jsii_.InvokeVoid(
		d,
		"resetFewShotExamples",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowGeneratorSummarizationContextOutputReference) ResetOutputLanguageCode() {
	_jsii_.InvokeVoid(
		d,
		"resetOutputLanguageCode",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowGeneratorSummarizationContextOutputReference) ResetSummarizationSections() {
	_jsii_.InvokeVoid(
		d,
		"resetSummarizationSections",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowGeneratorSummarizationContextOutputReference) ResetVersion() {
	_jsii_.InvokeVoid(
		d,
		"resetVersion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowGeneratorSummarizationContextOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowGeneratorSummarizationContextOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

