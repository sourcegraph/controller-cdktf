package discoveryenginewidgetconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/discoveryenginewidgetconfig/internal"
)

type DiscoveryEngineWidgetConfigUiSettingsOutputReference interface {
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
	DataStoreUiConfigs() DiscoveryEngineWidgetConfigUiSettingsDataStoreUiConfigsList
	DataStoreUiConfigsInput() any
	DefaultSearchRequestOrderBy() *string
	SetDefaultSearchRequestOrderBy(val *string)
	DefaultSearchRequestOrderByInput() *string
	DisableUserEventsCollection() any
	SetDisableUserEventsCollection(val any)
	DisableUserEventsCollectionInput() any
	EnableAutocomplete() any
	SetEnableAutocomplete(val any)
	EnableAutocompleteInput() any
	EnableCreateAgentButton() any
	SetEnableCreateAgentButton(val any)
	EnableCreateAgentButtonInput() any
	EnablePeopleSearch() any
	SetEnablePeopleSearch(val any)
	EnablePeopleSearchInput() any
	EnableQualityFeedback() any
	SetEnableQualityFeedback(val any)
	EnableQualityFeedbackInput() any
	EnableSafeSearch() any
	SetEnableSafeSearch(val any)
	EnableSafeSearchInput() any
	EnableSearchAsYouType() any
	SetEnableSearchAsYouType(val any)
	EnableSearchAsYouTypeInput() any
	EnableVisualContentSummary() any
	SetEnableVisualContentSummary(val any)
	EnableVisualContentSummaryInput() any
	// Experimental.
	Fqn() *string
	GenerativeAnswerConfig() DiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference
	GenerativeAnswerConfigInput() *DiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfig
	InteractionType() *string
	SetInteractionType(val *string)
	InteractionTypeInput() *string
	InternalValue() *DiscoveryEngineWidgetConfigUiSettings
	SetInternalValue(val *DiscoveryEngineWidgetConfigUiSettings)
	ResultDescriptionType() *string
	SetResultDescriptionType(val *string)
	ResultDescriptionTypeInput() *string
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
	PutDataStoreUiConfigs(value any)
	PutGenerativeAnswerConfig(value *DiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfig)
	ResetDataStoreUiConfigs()
	ResetDefaultSearchRequestOrderBy()
	ResetDisableUserEventsCollection()
	ResetEnableAutocomplete()
	ResetEnableCreateAgentButton()
	ResetEnablePeopleSearch()
	ResetEnableQualityFeedback()
	ResetEnableSafeSearch()
	ResetEnableSearchAsYouType()
	ResetEnableVisualContentSummary()
	ResetGenerativeAnswerConfig()
	ResetInteractionType()
	ResetResultDescriptionType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DiscoveryEngineWidgetConfigUiSettingsOutputReference
type jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsOutputReference) DataStoreUiConfigs() DiscoveryEngineWidgetConfigUiSettingsDataStoreUiConfigsList {
	var returns DiscoveryEngineWidgetConfigUiSettingsDataStoreUiConfigsList
	_jsii_.Get(
		j,
		"dataStoreUiConfigs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsOutputReference) DataStoreUiConfigsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"dataStoreUiConfigsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsOutputReference) DefaultSearchRequestOrderBy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultSearchRequestOrderBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsOutputReference) DefaultSearchRequestOrderByInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultSearchRequestOrderByInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsOutputReference) DisableUserEventsCollection() any {
	var returns any
	_jsii_.Get(
		j,
		"disableUserEventsCollection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsOutputReference) DisableUserEventsCollectionInput() any {
	var returns any
	_jsii_.Get(
		j,
		"disableUserEventsCollectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsOutputReference) EnableAutocomplete() any {
	var returns any
	_jsii_.Get(
		j,
		"enableAutocomplete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsOutputReference) EnableAutocompleteInput() any {
	var returns any
	_jsii_.Get(
		j,
		"enableAutocompleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsOutputReference) EnableCreateAgentButton() any {
	var returns any
	_jsii_.Get(
		j,
		"enableCreateAgentButton",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsOutputReference) EnableCreateAgentButtonInput() any {
	var returns any
	_jsii_.Get(
		j,
		"enableCreateAgentButtonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsOutputReference) EnablePeopleSearch() any {
	var returns any
	_jsii_.Get(
		j,
		"enablePeopleSearch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsOutputReference) EnablePeopleSearchInput() any {
	var returns any
	_jsii_.Get(
		j,
		"enablePeopleSearchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsOutputReference) EnableQualityFeedback() any {
	var returns any
	_jsii_.Get(
		j,
		"enableQualityFeedback",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsOutputReference) EnableQualityFeedbackInput() any {
	var returns any
	_jsii_.Get(
		j,
		"enableQualityFeedbackInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsOutputReference) EnableSafeSearch() any {
	var returns any
	_jsii_.Get(
		j,
		"enableSafeSearch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsOutputReference) EnableSafeSearchInput() any {
	var returns any
	_jsii_.Get(
		j,
		"enableSafeSearchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsOutputReference) EnableSearchAsYouType() any {
	var returns any
	_jsii_.Get(
		j,
		"enableSearchAsYouType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsOutputReference) EnableSearchAsYouTypeInput() any {
	var returns any
	_jsii_.Get(
		j,
		"enableSearchAsYouTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsOutputReference) EnableVisualContentSummary() any {
	var returns any
	_jsii_.Get(
		j,
		"enableVisualContentSummary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsOutputReference) EnableVisualContentSummaryInput() any {
	var returns any
	_jsii_.Get(
		j,
		"enableVisualContentSummaryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsOutputReference) GenerativeAnswerConfig() DiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference {
	var returns DiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference
	_jsii_.Get(
		j,
		"generativeAnswerConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsOutputReference) GenerativeAnswerConfigInput() *DiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfig {
	var returns *DiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfig
	_jsii_.Get(
		j,
		"generativeAnswerConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsOutputReference) InteractionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interactionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsOutputReference) InteractionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interactionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsOutputReference) InternalValue() *DiscoveryEngineWidgetConfigUiSettings {
	var returns *DiscoveryEngineWidgetConfigUiSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsOutputReference) ResultDescriptionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resultDescriptionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsOutputReference) ResultDescriptionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resultDescriptionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewDiscoveryEngineWidgetConfigUiSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DiscoveryEngineWidgetConfigUiSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewDiscoveryEngineWidgetConfigUiSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.discoveryEngineWidgetConfig.DiscoveryEngineWidgetConfigUiSettingsOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDiscoveryEngineWidgetConfigUiSettingsOutputReference_Override(d DiscoveryEngineWidgetConfigUiSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.discoveryEngineWidgetConfig.DiscoveryEngineWidgetConfigUiSettingsOutputReference",
		[]any{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsOutputReference) SetDefaultSearchRequestOrderBy(val *string) {
	if err := j.validateSetDefaultSearchRequestOrderByParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultSearchRequestOrderBy",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsOutputReference) SetDisableUserEventsCollection(val any) {
	if err := j.validateSetDisableUserEventsCollectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableUserEventsCollection",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsOutputReference) SetEnableAutocomplete(val any) {
	if err := j.validateSetEnableAutocompleteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableAutocomplete",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsOutputReference) SetEnableCreateAgentButton(val any) {
	if err := j.validateSetEnableCreateAgentButtonParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableCreateAgentButton",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsOutputReference) SetEnablePeopleSearch(val any) {
	if err := j.validateSetEnablePeopleSearchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enablePeopleSearch",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsOutputReference) SetEnableQualityFeedback(val any) {
	if err := j.validateSetEnableQualityFeedbackParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableQualityFeedback",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsOutputReference) SetEnableSafeSearch(val any) {
	if err := j.validateSetEnableSafeSearchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableSafeSearch",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsOutputReference) SetEnableSearchAsYouType(val any) {
	if err := j.validateSetEnableSearchAsYouTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableSearchAsYouType",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsOutputReference) SetEnableVisualContentSummary(val any) {
	if err := j.validateSetEnableVisualContentSummaryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableVisualContentSummary",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsOutputReference) SetInteractionType(val *string) {
	if err := j.validateSetInteractionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"interactionType",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsOutputReference) SetInternalValue(val *DiscoveryEngineWidgetConfigUiSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsOutputReference) SetResultDescriptionType(val *string) {
	if err := j.validateSetResultDescriptionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resultDescriptionType",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (d *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsOutputReference) PutDataStoreUiConfigs(value any) {
	if err := d.validatePutDataStoreUiConfigsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putDataStoreUiConfigs",
		[]any{value},
	)
}

func (d *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsOutputReference) PutGenerativeAnswerConfig(value *DiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfig) {
	if err := d.validatePutGenerativeAnswerConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putGenerativeAnswerConfig",
		[]any{value},
	)
}

func (d *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsOutputReference) ResetDataStoreUiConfigs() {
	_jsii_.InvokeVoid(
		d,
		"resetDataStoreUiConfigs",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsOutputReference) ResetDefaultSearchRequestOrderBy() {
	_jsii_.InvokeVoid(
		d,
		"resetDefaultSearchRequestOrderBy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsOutputReference) ResetDisableUserEventsCollection() {
	_jsii_.InvokeVoid(
		d,
		"resetDisableUserEventsCollection",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsOutputReference) ResetEnableAutocomplete() {
	_jsii_.InvokeVoid(
		d,
		"resetEnableAutocomplete",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsOutputReference) ResetEnableCreateAgentButton() {
	_jsii_.InvokeVoid(
		d,
		"resetEnableCreateAgentButton",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsOutputReference) ResetEnablePeopleSearch() {
	_jsii_.InvokeVoid(
		d,
		"resetEnablePeopleSearch",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsOutputReference) ResetEnableQualityFeedback() {
	_jsii_.InvokeVoid(
		d,
		"resetEnableQualityFeedback",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsOutputReference) ResetEnableSafeSearch() {
	_jsii_.InvokeVoid(
		d,
		"resetEnableSafeSearch",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsOutputReference) ResetEnableSearchAsYouType() {
	_jsii_.InvokeVoid(
		d,
		"resetEnableSearchAsYouType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsOutputReference) ResetEnableVisualContentSummary() {
	_jsii_.InvokeVoid(
		d,
		"resetEnableVisualContentSummary",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsOutputReference) ResetGenerativeAnswerConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetGenerativeAnswerConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsOutputReference) ResetInteractionType() {
	_jsii_.InvokeVoid(
		d,
		"resetInteractionType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsOutputReference) ResetResultDescriptionType() {
	_jsii_.InvokeVoid(
		d,
		"resetResultDescriptionType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (d *jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
