package contactcenterinsightsqaquestion

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/contactcenterinsightsqaquestion/internal"
)

type ContactCenterInsightsQaQuestionTuningMetadataOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() any
	// Experimental.
	SetComplexObjectIndex(val any)
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
	DatasetValidationWarnings() *[]*string
	SetDatasetValidationWarnings(val *[]*string)
	DatasetValidationWarningsInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *ContactCenterInsightsQaQuestionTuningMetadata
	SetInternalValue(val *ContactCenterInsightsQaQuestionTuningMetadata)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TotalValidLabelCount() *string
	SetTotalValidLabelCount(val *string)
	TotalValidLabelCountInput() *string
	TuningError() *string
	SetTuningError(val *string)
	TuningErrorInput() *string
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]any
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
	ResetDatasetValidationWarnings()
	ResetTotalValidLabelCount()
	ResetTuningError()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ContactCenterInsightsQaQuestionTuningMetadataOutputReference
type jsiiProxy_ContactCenterInsightsQaQuestionTuningMetadataOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestionTuningMetadataOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestionTuningMetadataOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestionTuningMetadataOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestionTuningMetadataOutputReference) DatasetValidationWarnings() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"datasetValidationWarnings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestionTuningMetadataOutputReference) DatasetValidationWarningsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"datasetValidationWarningsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestionTuningMetadataOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestionTuningMetadataOutputReference) InternalValue() *ContactCenterInsightsQaQuestionTuningMetadata {
	var returns *ContactCenterInsightsQaQuestionTuningMetadata
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestionTuningMetadataOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestionTuningMetadataOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestionTuningMetadataOutputReference) TotalValidLabelCount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"totalValidLabelCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestionTuningMetadataOutputReference) TotalValidLabelCountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"totalValidLabelCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestionTuningMetadataOutputReference) TuningError() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tuningError",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestionTuningMetadataOutputReference) TuningErrorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tuningErrorInput",
		&returns,
	)
	return returns
}

func NewContactCenterInsightsQaQuestionTuningMetadataOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ContactCenterInsightsQaQuestionTuningMetadataOutputReference {
	_init_.Initialize()

	if err := validateNewContactCenterInsightsQaQuestionTuningMetadataOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ContactCenterInsightsQaQuestionTuningMetadataOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.contactCenterInsightsQaQuestion.ContactCenterInsightsQaQuestionTuningMetadataOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewContactCenterInsightsQaQuestionTuningMetadataOutputReference_Override(c ContactCenterInsightsQaQuestionTuningMetadataOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.contactCenterInsightsQaQuestion.ContactCenterInsightsQaQuestionTuningMetadataOutputReference",
		[]any{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestionTuningMetadataOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestionTuningMetadataOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestionTuningMetadataOutputReference) SetDatasetValidationWarnings(val *[]*string) {
	if err := j.validateSetDatasetValidationWarningsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"datasetValidationWarnings",
		val,
	)
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestionTuningMetadataOutputReference) SetInternalValue(val *ContactCenterInsightsQaQuestionTuningMetadata) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestionTuningMetadataOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestionTuningMetadataOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestionTuningMetadataOutputReference) SetTotalValidLabelCount(val *string) {
	if err := j.validateSetTotalValidLabelCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"totalValidLabelCount",
		val,
	)
}

func (j *jsiiProxy_ContactCenterInsightsQaQuestionTuningMetadataOutputReference) SetTuningError(val *string) {
	if err := j.validateSetTuningErrorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tuningError",
		val,
	)
}

func (c *jsiiProxy_ContactCenterInsightsQaQuestionTuningMetadataOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContactCenterInsightsQaQuestionTuningMetadataOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContactCenterInsightsQaQuestionTuningMetadataOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContactCenterInsightsQaQuestionTuningMetadataOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContactCenterInsightsQaQuestionTuningMetadataOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContactCenterInsightsQaQuestionTuningMetadataOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContactCenterInsightsQaQuestionTuningMetadataOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContactCenterInsightsQaQuestionTuningMetadataOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContactCenterInsightsQaQuestionTuningMetadataOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContactCenterInsightsQaQuestionTuningMetadataOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContactCenterInsightsQaQuestionTuningMetadataOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContactCenterInsightsQaQuestionTuningMetadataOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContactCenterInsightsQaQuestionTuningMetadataOutputReference) ResetDatasetValidationWarnings() {
	_jsii_.InvokeVoid(
		c,
		"resetDatasetValidationWarnings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContactCenterInsightsQaQuestionTuningMetadataOutputReference) ResetTotalValidLabelCount() {
	_jsii_.InvokeVoid(
		c,
		"resetTotalValidLabelCount",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContactCenterInsightsQaQuestionTuningMetadataOutputReference) ResetTuningError() {
	_jsii_.InvokeVoid(
		c,
		"resetTuningError",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContactCenterInsightsQaQuestionTuningMetadataOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		c,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContactCenterInsightsQaQuestionTuningMetadataOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
