package chronicledashboardchart

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/chronicledashboardchart/internal"
)

type ChronicleDashboardChartDashboardChartDrillDownConfigLeftDrillDownsCustomSettingsOutputReference interface {
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
	ExternalLink() ChronicleDashboardChartDashboardChartDrillDownConfigLeftDrillDownsCustomSettingsExternalLinkOutputReference
	ExternalLinkInput() *ChronicleDashboardChartDashboardChartDrillDownConfigLeftDrillDownsCustomSettingsExternalLink
	Filter() ChronicleDashboardChartDashboardChartDrillDownConfigLeftDrillDownsCustomSettingsFilterOutputReference
	FilterInput() *ChronicleDashboardChartDashboardChartDrillDownConfigLeftDrillDownsCustomSettingsFilter
	// Experimental.
	Fqn() *string
	InternalValue() *ChronicleDashboardChartDashboardChartDrillDownConfigLeftDrillDownsCustomSettings
	SetInternalValue(val *ChronicleDashboardChartDashboardChartDrillDownConfigLeftDrillDownsCustomSettings)
	LeftClickColumn() *string
	SetLeftClickColumn(val *string)
	LeftClickColumnInput() *string
	NewTab() any
	SetNewTab(val any)
	NewTabInput() any
	Query() ChronicleDashboardChartDashboardChartDrillDownConfigLeftDrillDownsCustomSettingsQueryOutputReference
	QueryInput() *ChronicleDashboardChartDashboardChartDrillDownConfigLeftDrillDownsCustomSettingsQuery
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
	PutExternalLink(value *ChronicleDashboardChartDashboardChartDrillDownConfigLeftDrillDownsCustomSettingsExternalLink)
	PutFilter(value *ChronicleDashboardChartDashboardChartDrillDownConfigLeftDrillDownsCustomSettingsFilter)
	PutQuery(value *ChronicleDashboardChartDashboardChartDrillDownConfigLeftDrillDownsCustomSettingsQuery)
	ResetExternalLink()
	ResetFilter()
	ResetLeftClickColumn()
	ResetQuery()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ChronicleDashboardChartDashboardChartDrillDownConfigLeftDrillDownsCustomSettingsOutputReference
type jsiiProxy_ChronicleDashboardChartDashboardChartDrillDownConfigLeftDrillDownsCustomSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartDrillDownConfigLeftDrillDownsCustomSettingsOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartDrillDownConfigLeftDrillDownsCustomSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartDrillDownConfigLeftDrillDownsCustomSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartDrillDownConfigLeftDrillDownsCustomSettingsOutputReference) ExternalLink() ChronicleDashboardChartDashboardChartDrillDownConfigLeftDrillDownsCustomSettingsExternalLinkOutputReference {
	var returns ChronicleDashboardChartDashboardChartDrillDownConfigLeftDrillDownsCustomSettingsExternalLinkOutputReference
	_jsii_.Get(
		j,
		"externalLink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartDrillDownConfigLeftDrillDownsCustomSettingsOutputReference) ExternalLinkInput() *ChronicleDashboardChartDashboardChartDrillDownConfigLeftDrillDownsCustomSettingsExternalLink {
	var returns *ChronicleDashboardChartDashboardChartDrillDownConfigLeftDrillDownsCustomSettingsExternalLink
	_jsii_.Get(
		j,
		"externalLinkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartDrillDownConfigLeftDrillDownsCustomSettingsOutputReference) Filter() ChronicleDashboardChartDashboardChartDrillDownConfigLeftDrillDownsCustomSettingsFilterOutputReference {
	var returns ChronicleDashboardChartDashboardChartDrillDownConfigLeftDrillDownsCustomSettingsFilterOutputReference
	_jsii_.Get(
		j,
		"filter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartDrillDownConfigLeftDrillDownsCustomSettingsOutputReference) FilterInput() *ChronicleDashboardChartDashboardChartDrillDownConfigLeftDrillDownsCustomSettingsFilter {
	var returns *ChronicleDashboardChartDashboardChartDrillDownConfigLeftDrillDownsCustomSettingsFilter
	_jsii_.Get(
		j,
		"filterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartDrillDownConfigLeftDrillDownsCustomSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartDrillDownConfigLeftDrillDownsCustomSettingsOutputReference) InternalValue() *ChronicleDashboardChartDashboardChartDrillDownConfigLeftDrillDownsCustomSettings {
	var returns *ChronicleDashboardChartDashboardChartDrillDownConfigLeftDrillDownsCustomSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartDrillDownConfigLeftDrillDownsCustomSettingsOutputReference) LeftClickColumn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"leftClickColumn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartDrillDownConfigLeftDrillDownsCustomSettingsOutputReference) LeftClickColumnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"leftClickColumnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartDrillDownConfigLeftDrillDownsCustomSettingsOutputReference) NewTab() any {
	var returns any
	_jsii_.Get(
		j,
		"newTab",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartDrillDownConfigLeftDrillDownsCustomSettingsOutputReference) NewTabInput() any {
	var returns any
	_jsii_.Get(
		j,
		"newTabInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartDrillDownConfigLeftDrillDownsCustomSettingsOutputReference) Query() ChronicleDashboardChartDashboardChartDrillDownConfigLeftDrillDownsCustomSettingsQueryOutputReference {
	var returns ChronicleDashboardChartDashboardChartDrillDownConfigLeftDrillDownsCustomSettingsQueryOutputReference
	_jsii_.Get(
		j,
		"query",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartDrillDownConfigLeftDrillDownsCustomSettingsOutputReference) QueryInput() *ChronicleDashboardChartDashboardChartDrillDownConfigLeftDrillDownsCustomSettingsQuery {
	var returns *ChronicleDashboardChartDashboardChartDrillDownConfigLeftDrillDownsCustomSettingsQuery
	_jsii_.Get(
		j,
		"queryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartDrillDownConfigLeftDrillDownsCustomSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartDrillDownConfigLeftDrillDownsCustomSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewChronicleDashboardChartDashboardChartDrillDownConfigLeftDrillDownsCustomSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ChronicleDashboardChartDashboardChartDrillDownConfigLeftDrillDownsCustomSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewChronicleDashboardChartDashboardChartDrillDownConfigLeftDrillDownsCustomSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ChronicleDashboardChartDashboardChartDrillDownConfigLeftDrillDownsCustomSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.chronicleDashboardChart.ChronicleDashboardChartDashboardChartDrillDownConfigLeftDrillDownsCustomSettingsOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewChronicleDashboardChartDashboardChartDrillDownConfigLeftDrillDownsCustomSettingsOutputReference_Override(c ChronicleDashboardChartDashboardChartDrillDownConfigLeftDrillDownsCustomSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.chronicleDashboardChart.ChronicleDashboardChartDashboardChartDrillDownConfigLeftDrillDownsCustomSettingsOutputReference",
		[]any{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartDrillDownConfigLeftDrillDownsCustomSettingsOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartDrillDownConfigLeftDrillDownsCustomSettingsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartDrillDownConfigLeftDrillDownsCustomSettingsOutputReference) SetInternalValue(val *ChronicleDashboardChartDashboardChartDrillDownConfigLeftDrillDownsCustomSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartDrillDownConfigLeftDrillDownsCustomSettingsOutputReference) SetLeftClickColumn(val *string) {
	if err := j.validateSetLeftClickColumnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"leftClickColumn",
		val,
	)
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartDrillDownConfigLeftDrillDownsCustomSettingsOutputReference) SetNewTab(val any) {
	if err := j.validateSetNewTabParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"newTab",
		val,
	)
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartDrillDownConfigLeftDrillDownsCustomSettingsOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartDrillDownConfigLeftDrillDownsCustomSettingsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartDrillDownConfigLeftDrillDownsCustomSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartDrillDownConfigLeftDrillDownsCustomSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartDrillDownConfigLeftDrillDownsCustomSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartDrillDownConfigLeftDrillDownsCustomSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartDrillDownConfigLeftDrillDownsCustomSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartDrillDownConfigLeftDrillDownsCustomSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartDrillDownConfigLeftDrillDownsCustomSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartDrillDownConfigLeftDrillDownsCustomSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartDrillDownConfigLeftDrillDownsCustomSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartDrillDownConfigLeftDrillDownsCustomSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartDrillDownConfigLeftDrillDownsCustomSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartDrillDownConfigLeftDrillDownsCustomSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartDrillDownConfigLeftDrillDownsCustomSettingsOutputReference) PutExternalLink(value *ChronicleDashboardChartDashboardChartDrillDownConfigLeftDrillDownsCustomSettingsExternalLink) {
	if err := c.validatePutExternalLinkParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putExternalLink",
		[]any{value},
	)
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartDrillDownConfigLeftDrillDownsCustomSettingsOutputReference) PutFilter(value *ChronicleDashboardChartDashboardChartDrillDownConfigLeftDrillDownsCustomSettingsFilter) {
	if err := c.validatePutFilterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putFilter",
		[]any{value},
	)
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartDrillDownConfigLeftDrillDownsCustomSettingsOutputReference) PutQuery(value *ChronicleDashboardChartDashboardChartDrillDownConfigLeftDrillDownsCustomSettingsQuery) {
	if err := c.validatePutQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putQuery",
		[]any{value},
	)
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartDrillDownConfigLeftDrillDownsCustomSettingsOutputReference) ResetExternalLink() {
	_jsii_.InvokeVoid(
		c,
		"resetExternalLink",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartDrillDownConfigLeftDrillDownsCustomSettingsOutputReference) ResetFilter() {
	_jsii_.InvokeVoid(
		c,
		"resetFilter",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartDrillDownConfigLeftDrillDownsCustomSettingsOutputReference) ResetLeftClickColumn() {
	_jsii_.InvokeVoid(
		c,
		"resetLeftClickColumn",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartDrillDownConfigLeftDrillDownsCustomSettingsOutputReference) ResetQuery() {
	_jsii_.InvokeVoid(
		c,
		"resetQuery",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartDrillDownConfigLeftDrillDownsCustomSettingsOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartDrillDownConfigLeftDrillDownsCustomSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
