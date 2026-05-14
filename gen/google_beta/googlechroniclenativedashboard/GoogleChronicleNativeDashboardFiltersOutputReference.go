package googlechroniclenativedashboard

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlechroniclenativedashboard/internal"
)

type GoogleChronicleNativeDashboardFiltersOutputReference interface {
	cdktf.ComplexObject
	ChartIds() *[]*string
	SetChartIds(val *[]*string)
	ChartIdsInput() *[]*string
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
	DataSource() *string
	SetDataSource(val *string)
	DataSourceInput() *string
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	FieldPath() *string
	SetFieldPath(val *string)
	FieldPathInput() *string
	FilterOperatorAndFieldValues() GoogleChronicleNativeDashboardFiltersFilterOperatorAndFieldValuesList
	FilterOperatorAndFieldValuesInput() interface{}
	// Experimental.
	Fqn() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IsMandatory() interface{}
	SetIsMandatory(val interface{})
	IsMandatoryInput() interface{}
	IsStandardTimeRangeFilter() interface{}
	SetIsStandardTimeRangeFilter(val interface{})
	IsStandardTimeRangeFilterEnabled() interface{}
	SetIsStandardTimeRangeFilterEnabled(val interface{})
	IsStandardTimeRangeFilterEnabledInput() interface{}
	IsStandardTimeRangeFilterInput() interface{}
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
	PutFilterOperatorAndFieldValues(value interface{})
	ResetChartIds()
	ResetDataSource()
	ResetDisplayName()
	ResetFieldPath()
	ResetFilterOperatorAndFieldValues()
	ResetId()
	ResetIsMandatory()
	ResetIsStandardTimeRangeFilter()
	ResetIsStandardTimeRangeFilterEnabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleChronicleNativeDashboardFiltersOutputReference
type jsiiProxy_GoogleChronicleNativeDashboardFiltersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleChronicleNativeDashboardFiltersOutputReference) ChartIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"chartIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleNativeDashboardFiltersOutputReference) ChartIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"chartIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleNativeDashboardFiltersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleNativeDashboardFiltersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleNativeDashboardFiltersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleNativeDashboardFiltersOutputReference) DataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleNativeDashboardFiltersOutputReference) DataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleNativeDashboardFiltersOutputReference) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleNativeDashboardFiltersOutputReference) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleNativeDashboardFiltersOutputReference) FieldPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fieldPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleNativeDashboardFiltersOutputReference) FieldPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fieldPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleNativeDashboardFiltersOutputReference) FilterOperatorAndFieldValues() GoogleChronicleNativeDashboardFiltersFilterOperatorAndFieldValuesList {
	var returns GoogleChronicleNativeDashboardFiltersFilterOperatorAndFieldValuesList
	_jsii_.Get(
		j,
		"filterOperatorAndFieldValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleNativeDashboardFiltersOutputReference) FilterOperatorAndFieldValuesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"filterOperatorAndFieldValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleNativeDashboardFiltersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleNativeDashboardFiltersOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleNativeDashboardFiltersOutputReference) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleNativeDashboardFiltersOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleNativeDashboardFiltersOutputReference) IsMandatory() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isMandatory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleNativeDashboardFiltersOutputReference) IsMandatoryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isMandatoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleNativeDashboardFiltersOutputReference) IsStandardTimeRangeFilter() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isStandardTimeRangeFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleNativeDashboardFiltersOutputReference) IsStandardTimeRangeFilterEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isStandardTimeRangeFilterEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleNativeDashboardFiltersOutputReference) IsStandardTimeRangeFilterEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isStandardTimeRangeFilterEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleNativeDashboardFiltersOutputReference) IsStandardTimeRangeFilterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isStandardTimeRangeFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleNativeDashboardFiltersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleNativeDashboardFiltersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleChronicleNativeDashboardFiltersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleChronicleNativeDashboardFiltersOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleChronicleNativeDashboardFiltersOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleChronicleNativeDashboardFiltersOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleChronicleNativeDashboard.GoogleChronicleNativeDashboardFiltersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleChronicleNativeDashboardFiltersOutputReference_Override(g GoogleChronicleNativeDashboardFiltersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleChronicleNativeDashboard.GoogleChronicleNativeDashboardFiltersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleChronicleNativeDashboardFiltersOutputReference)SetChartIds(val *[]*string) {
	if err := j.validateSetChartIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"chartIds",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleNativeDashboardFiltersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleNativeDashboardFiltersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleNativeDashboardFiltersOutputReference)SetDataSource(val *string) {
	if err := j.validateSetDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataSource",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleNativeDashboardFiltersOutputReference)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleNativeDashboardFiltersOutputReference)SetFieldPath(val *string) {
	if err := j.validateSetFieldPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fieldPath",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleNativeDashboardFiltersOutputReference)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleNativeDashboardFiltersOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleNativeDashboardFiltersOutputReference)SetIsMandatory(val interface{}) {
	if err := j.validateSetIsMandatoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isMandatory",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleNativeDashboardFiltersOutputReference)SetIsStandardTimeRangeFilter(val interface{}) {
	if err := j.validateSetIsStandardTimeRangeFilterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isStandardTimeRangeFilter",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleNativeDashboardFiltersOutputReference)SetIsStandardTimeRangeFilterEnabled(val interface{}) {
	if err := j.validateSetIsStandardTimeRangeFilterEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isStandardTimeRangeFilterEnabled",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleNativeDashboardFiltersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleNativeDashboardFiltersOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleChronicleNativeDashboardFiltersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleChronicleNativeDashboardFiltersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleChronicleNativeDashboardFiltersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleChronicleNativeDashboardFiltersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleChronicleNativeDashboardFiltersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleChronicleNativeDashboardFiltersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleChronicleNativeDashboardFiltersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleChronicleNativeDashboardFiltersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleChronicleNativeDashboardFiltersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleChronicleNativeDashboardFiltersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleChronicleNativeDashboardFiltersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleChronicleNativeDashboardFiltersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleChronicleNativeDashboardFiltersOutputReference) PutFilterOperatorAndFieldValues(value interface{}) {
	if err := g.validatePutFilterOperatorAndFieldValuesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putFilterOperatorAndFieldValues",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleNativeDashboardFiltersOutputReference) ResetChartIds() {
	_jsii_.InvokeVoid(
		g,
		"resetChartIds",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleNativeDashboardFiltersOutputReference) ResetDataSource() {
	_jsii_.InvokeVoid(
		g,
		"resetDataSource",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleNativeDashboardFiltersOutputReference) ResetDisplayName() {
	_jsii_.InvokeVoid(
		g,
		"resetDisplayName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleNativeDashboardFiltersOutputReference) ResetFieldPath() {
	_jsii_.InvokeVoid(
		g,
		"resetFieldPath",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleNativeDashboardFiltersOutputReference) ResetFilterOperatorAndFieldValues() {
	_jsii_.InvokeVoid(
		g,
		"resetFilterOperatorAndFieldValues",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleNativeDashboardFiltersOutputReference) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleNativeDashboardFiltersOutputReference) ResetIsMandatory() {
	_jsii_.InvokeVoid(
		g,
		"resetIsMandatory",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleNativeDashboardFiltersOutputReference) ResetIsStandardTimeRangeFilter() {
	_jsii_.InvokeVoid(
		g,
		"resetIsStandardTimeRangeFilter",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleNativeDashboardFiltersOutputReference) ResetIsStandardTimeRangeFilterEnabled() {
	_jsii_.InvokeVoid(
		g,
		"resetIsStandardTimeRangeFilterEnabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleNativeDashboardFiltersOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleChronicleNativeDashboardFiltersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

