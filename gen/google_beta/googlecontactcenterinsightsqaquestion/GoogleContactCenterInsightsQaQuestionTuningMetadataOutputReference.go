package googlecontactcenterinsightsqaquestion

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecontactcenterinsightsqaquestion/internal"
)

type GoogleContactCenterInsightsQaQuestionTuningMetadataOutputReference interface {
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
	DatasetValidationWarnings() *[]*string
	SetDatasetValidationWarnings(val *[]*string)
	DatasetValidationWarningsInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleContactCenterInsightsQaQuestionTuningMetadata
	SetInternalValue(val *GoogleContactCenterInsightsQaQuestionTuningMetadata)
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
	ResetDatasetValidationWarnings()
	ResetTotalValidLabelCount()
	ResetTuningError()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleContactCenterInsightsQaQuestionTuningMetadataOutputReference
type jsiiProxy_GoogleContactCenterInsightsQaQuestionTuningMetadataOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestionTuningMetadataOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestionTuningMetadataOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestionTuningMetadataOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestionTuningMetadataOutputReference) DatasetValidationWarnings() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"datasetValidationWarnings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestionTuningMetadataOutputReference) DatasetValidationWarningsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"datasetValidationWarningsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestionTuningMetadataOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestionTuningMetadataOutputReference) InternalValue() *GoogleContactCenterInsightsQaQuestionTuningMetadata {
	var returns *GoogleContactCenterInsightsQaQuestionTuningMetadata
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestionTuningMetadataOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestionTuningMetadataOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestionTuningMetadataOutputReference) TotalValidLabelCount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"totalValidLabelCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestionTuningMetadataOutputReference) TotalValidLabelCountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"totalValidLabelCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestionTuningMetadataOutputReference) TuningError() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tuningError",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestionTuningMetadataOutputReference) TuningErrorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tuningErrorInput",
		&returns,
	)
	return returns
}


func NewGoogleContactCenterInsightsQaQuestionTuningMetadataOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleContactCenterInsightsQaQuestionTuningMetadataOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleContactCenterInsightsQaQuestionTuningMetadataOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleContactCenterInsightsQaQuestionTuningMetadataOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleContactCenterInsightsQaQuestion.GoogleContactCenterInsightsQaQuestionTuningMetadataOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleContactCenterInsightsQaQuestionTuningMetadataOutputReference_Override(g GoogleContactCenterInsightsQaQuestionTuningMetadataOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleContactCenterInsightsQaQuestion.GoogleContactCenterInsightsQaQuestionTuningMetadataOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestionTuningMetadataOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestionTuningMetadataOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestionTuningMetadataOutputReference)SetDatasetValidationWarnings(val *[]*string) {
	if err := j.validateSetDatasetValidationWarningsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"datasetValidationWarnings",
		val,
	)
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestionTuningMetadataOutputReference)SetInternalValue(val *GoogleContactCenterInsightsQaQuestionTuningMetadata) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestionTuningMetadataOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestionTuningMetadataOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestionTuningMetadataOutputReference)SetTotalValidLabelCount(val *string) {
	if err := j.validateSetTotalValidLabelCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"totalValidLabelCount",
		val,
	)
}

func (j *jsiiProxy_GoogleContactCenterInsightsQaQuestionTuningMetadataOutputReference)SetTuningError(val *string) {
	if err := j.validateSetTuningErrorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tuningError",
		val,
	)
}

func (g *jsiiProxy_GoogleContactCenterInsightsQaQuestionTuningMetadataOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContactCenterInsightsQaQuestionTuningMetadataOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleContactCenterInsightsQaQuestionTuningMetadataOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleContactCenterInsightsQaQuestionTuningMetadataOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleContactCenterInsightsQaQuestionTuningMetadataOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleContactCenterInsightsQaQuestionTuningMetadataOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleContactCenterInsightsQaQuestionTuningMetadataOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleContactCenterInsightsQaQuestionTuningMetadataOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleContactCenterInsightsQaQuestionTuningMetadataOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleContactCenterInsightsQaQuestionTuningMetadataOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleContactCenterInsightsQaQuestionTuningMetadataOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContactCenterInsightsQaQuestionTuningMetadataOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleContactCenterInsightsQaQuestionTuningMetadataOutputReference) ResetDatasetValidationWarnings() {
	_jsii_.InvokeVoid(
		g,
		"resetDatasetValidationWarnings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContactCenterInsightsQaQuestionTuningMetadataOutputReference) ResetTotalValidLabelCount() {
	_jsii_.InvokeVoid(
		g,
		"resetTotalValidLabelCount",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContactCenterInsightsQaQuestionTuningMetadataOutputReference) ResetTuningError() {
	_jsii_.InvokeVoid(
		g,
		"resetTuningError",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContactCenterInsightsQaQuestionTuningMetadataOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleContactCenterInsightsQaQuestionTuningMetadataOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

