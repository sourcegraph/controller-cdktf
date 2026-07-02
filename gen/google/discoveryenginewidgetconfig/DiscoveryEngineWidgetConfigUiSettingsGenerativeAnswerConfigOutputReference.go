package discoveryenginewidgetconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/discoveryenginewidgetconfig/internal"
)

type DiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference interface {
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
	DisableRelatedQuestions() any
	SetDisableRelatedQuestions(val any)
	DisableRelatedQuestionsInput() any
	// Experimental.
	Fqn() *string
	IgnoreAdversarialQuery() any
	SetIgnoreAdversarialQuery(val any)
	IgnoreAdversarialQueryInput() any
	IgnoreLowRelevantContent() any
	SetIgnoreLowRelevantContent(val any)
	IgnoreLowRelevantContentInput() any
	IgnoreNonAnswerSeekingQuery() any
	SetIgnoreNonAnswerSeekingQuery(val any)
	IgnoreNonAnswerSeekingQueryInput() any
	ImageSource() *string
	SetImageSource(val *string)
	ImageSourceInput() *string
	InternalValue() *DiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfig
	SetInternalValue(val *DiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfig)
	LanguageCode() *string
	SetLanguageCode(val *string)
	LanguageCodeInput() *string
	MaxRephraseSteps() *float64
	SetMaxRephraseSteps(val *float64)
	MaxRephraseStepsInput() *float64
	ModelPromptPreamble() *string
	SetModelPromptPreamble(val *string)
	ModelPromptPreambleInput() *string
	ModelVersion() *string
	SetModelVersion(val *string)
	ModelVersionInput() *string
	ResultCount() *float64
	SetResultCount(val *float64)
	ResultCountInput() *float64
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
	ResetDisableRelatedQuestions()
	ResetIgnoreAdversarialQuery()
	ResetIgnoreLowRelevantContent()
	ResetIgnoreNonAnswerSeekingQuery()
	ResetImageSource()
	ResetLanguageCode()
	ResetMaxRephraseSteps()
	ResetModelPromptPreamble()
	ResetModelVersion()
	ResetResultCount()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference
type jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) DisableRelatedQuestions() any {
	var returns any
	_jsii_.Get(
		j,
		"disableRelatedQuestions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) DisableRelatedQuestionsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"disableRelatedQuestionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) IgnoreAdversarialQuery() any {
	var returns any
	_jsii_.Get(
		j,
		"ignoreAdversarialQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) IgnoreAdversarialQueryInput() any {
	var returns any
	_jsii_.Get(
		j,
		"ignoreAdversarialQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) IgnoreLowRelevantContent() any {
	var returns any
	_jsii_.Get(
		j,
		"ignoreLowRelevantContent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) IgnoreLowRelevantContentInput() any {
	var returns any
	_jsii_.Get(
		j,
		"ignoreLowRelevantContentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) IgnoreNonAnswerSeekingQuery() any {
	var returns any
	_jsii_.Get(
		j,
		"ignoreNonAnswerSeekingQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) IgnoreNonAnswerSeekingQueryInput() any {
	var returns any
	_jsii_.Get(
		j,
		"ignoreNonAnswerSeekingQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) ImageSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) ImageSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) InternalValue() *DiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfig {
	var returns *DiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) LanguageCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"languageCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) LanguageCodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"languageCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) MaxRephraseSteps() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRephraseSteps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) MaxRephraseStepsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRephraseStepsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) ModelPromptPreamble() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelPromptPreamble",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) ModelPromptPreambleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelPromptPreambleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) ModelVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) ModelVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) ResultCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"resultCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) ResultCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"resultCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewDiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference {
	_init_.Initialize()

	if err := validateNewDiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.discoveryEngineWidgetConfig.DiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference_Override(d DiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.discoveryEngineWidgetConfig.DiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference",
		[]any{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) SetDisableRelatedQuestions(val any) {
	if err := j.validateSetDisableRelatedQuestionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableRelatedQuestions",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) SetIgnoreAdversarialQuery(val any) {
	if err := j.validateSetIgnoreAdversarialQueryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignoreAdversarialQuery",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) SetIgnoreLowRelevantContent(val any) {
	if err := j.validateSetIgnoreLowRelevantContentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignoreLowRelevantContent",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) SetIgnoreNonAnswerSeekingQuery(val any) {
	if err := j.validateSetIgnoreNonAnswerSeekingQueryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignoreNonAnswerSeekingQuery",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) SetImageSource(val *string) {
	if err := j.validateSetImageSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageSource",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) SetInternalValue(val *DiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) SetLanguageCode(val *string) {
	if err := j.validateSetLanguageCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"languageCode",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) SetMaxRephraseSteps(val *float64) {
	if err := j.validateSetMaxRephraseStepsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxRephraseSteps",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) SetModelPromptPreamble(val *string) {
	if err := j.validateSetModelPromptPreambleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"modelPromptPreamble",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) SetModelVersion(val *string) {
	if err := j.validateSetModelVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"modelVersion",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) SetResultCount(val *float64) {
	if err := j.validateSetResultCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resultCount",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) ResetDisableRelatedQuestions() {
	_jsii_.InvokeVoid(
		d,
		"resetDisableRelatedQuestions",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) ResetIgnoreAdversarialQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetIgnoreAdversarialQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) ResetIgnoreLowRelevantContent() {
	_jsii_.InvokeVoid(
		d,
		"resetIgnoreLowRelevantContent",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) ResetIgnoreNonAnswerSeekingQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetIgnoreNonAnswerSeekingQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) ResetImageSource() {
	_jsii_.InvokeVoid(
		d,
		"resetImageSource",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) ResetLanguageCode() {
	_jsii_.InvokeVoid(
		d,
		"resetLanguageCode",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) ResetMaxRephraseSteps() {
	_jsii_.InvokeVoid(
		d,
		"resetMaxRephraseSteps",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) ResetModelPromptPreamble() {
	_jsii_.InvokeVoid(
		d,
		"resetModelPromptPreamble",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) ResetModelVersion() {
	_jsii_.InvokeVoid(
		d,
		"resetModelVersion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) ResetResultCount() {
	_jsii_.InvokeVoid(
		d,
		"resetResultCount",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		d,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
