package googledatalosspreventionjobtrigger

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledatalosspreventionjobtrigger/internal"
)

type GoogleDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference interface {
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
	ExcludedFields() GoogleDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFieldsList
	ExcludedFieldsInput() interface{}
	// Experimental.
	Fqn() *string
	IdentifyingFields() GoogleDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFieldsList
	IdentifyingFieldsInput() interface{}
	IncludedFields() GoogleDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFieldsList
	IncludedFieldsInput() interface{}
	InternalValue() *GoogleDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptions
	SetInternalValue(val *GoogleDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptions)
	RowsLimit() *float64
	SetRowsLimit(val *float64)
	RowsLimitInput() *float64
	RowsLimitPercent() *float64
	SetRowsLimitPercent(val *float64)
	RowsLimitPercentInput() *float64
	SampleMethod() *string
	SetSampleMethod(val *string)
	SampleMethodInput() *string
	TableReference() GoogleDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsTableReferenceOutputReference
	TableReferenceInput() *GoogleDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsTableReference
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
	PutExcludedFields(value interface{})
	PutIdentifyingFields(value interface{})
	PutIncludedFields(value interface{})
	PutTableReference(value *GoogleDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsTableReference)
	ResetExcludedFields()
	ResetIdentifyingFields()
	ResetIncludedFields()
	ResetRowsLimit()
	ResetRowsLimitPercent()
	ResetSampleMethod()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference
type jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference) ExcludedFields() GoogleDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFieldsList {
	var returns GoogleDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFieldsList
	_jsii_.Get(
		j,
		"excludedFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference) ExcludedFieldsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludedFieldsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference) IdentifyingFields() GoogleDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFieldsList {
	var returns GoogleDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFieldsList
	_jsii_.Get(
		j,
		"identifyingFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference) IdentifyingFieldsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"identifyingFieldsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference) IncludedFields() GoogleDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFieldsList {
	var returns GoogleDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFieldsList
	_jsii_.Get(
		j,
		"includedFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference) IncludedFieldsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includedFieldsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference) InternalValue() *GoogleDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptions {
	var returns *GoogleDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference) RowsLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rowsLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference) RowsLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rowsLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference) RowsLimitPercent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rowsLimitPercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference) RowsLimitPercentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rowsLimitPercentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference) SampleMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sampleMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference) SampleMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sampleMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference) TableReference() GoogleDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsTableReferenceOutputReference {
	var returns GoogleDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsTableReferenceOutputReference
	_jsii_.Get(
		j,
		"tableReference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference) TableReferenceInput() *GoogleDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsTableReference {
	var returns *GoogleDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsTableReference
	_jsii_.Get(
		j,
		"tableReferenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference{}

	_jsii_.Create(
		"google-beta.googleDataLossPreventionJobTrigger.GoogleDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference_Override(g GoogleDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleDataLossPreventionJobTrigger.GoogleDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference)SetInternalValue(val *GoogleDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference)SetRowsLimit(val *float64) {
	if err := j.validateSetRowsLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rowsLimit",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference)SetRowsLimitPercent(val *float64) {
	if err := j.validateSetRowsLimitPercentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rowsLimitPercent",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference)SetSampleMethod(val *string) {
	if err := j.validateSetSampleMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sampleMethod",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference) PutExcludedFields(value interface{}) {
	if err := g.validatePutExcludedFieldsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putExcludedFields",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference) PutIdentifyingFields(value interface{}) {
	if err := g.validatePutIdentifyingFieldsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putIdentifyingFields",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference) PutIncludedFields(value interface{}) {
	if err := g.validatePutIncludedFieldsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putIncludedFields",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference) PutTableReference(value *GoogleDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsTableReference) {
	if err := g.validatePutTableReferenceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTableReference",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference) ResetExcludedFields() {
	_jsii_.InvokeVoid(
		g,
		"resetExcludedFields",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference) ResetIdentifyingFields() {
	_jsii_.InvokeVoid(
		g,
		"resetIdentifyingFields",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference) ResetIncludedFields() {
	_jsii_.InvokeVoid(
		g,
		"resetIncludedFields",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference) ResetRowsLimit() {
	_jsii_.InvokeVoid(
		g,
		"resetRowsLimit",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference) ResetRowsLimitPercent() {
	_jsii_.InvokeVoid(
		g,
		"resetRowsLimitPercent",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference) ResetSampleMethod() {
	_jsii_.InvokeVoid(
		g,
		"resetSampleMethod",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

