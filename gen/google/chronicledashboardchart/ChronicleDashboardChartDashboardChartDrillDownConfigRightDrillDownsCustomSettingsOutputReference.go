package chronicledashboardchart

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/chronicledashboardchart/internal"
)

type ChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettingsOutputReference interface {
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
	ExternalLink() ChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettingsExternalLinkOutputReference
	ExternalLinkInput() *ChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettingsExternalLink
	Filter() ChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettingsFilterOutputReference
	FilterInput() *ChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettingsFilter
	// Experimental.
	Fqn() *string
	InternalValue() *ChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettings
	SetInternalValue(val *ChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettings)
	NewTab() any
	SetNewTab(val any)
	NewTabInput() any
	Query() ChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettingsQueryOutputReference
	QueryInput() *ChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettingsQuery
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
	PutExternalLink(value *ChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettingsExternalLink)
	PutFilter(value *ChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettingsFilter)
	PutQuery(value *ChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettingsQuery)
	ResetExternalLink()
	ResetFilter()
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

// The jsii proxy struct for ChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettingsOutputReference
type jsiiProxy_ChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettingsOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettingsOutputReference) ExternalLink() ChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettingsExternalLinkOutputReference {
	var returns ChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettingsExternalLinkOutputReference
	_jsii_.Get(
		j,
		"externalLink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettingsOutputReference) ExternalLinkInput() *ChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettingsExternalLink {
	var returns *ChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettingsExternalLink
	_jsii_.Get(
		j,
		"externalLinkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettingsOutputReference) Filter() ChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettingsFilterOutputReference {
	var returns ChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettingsFilterOutputReference
	_jsii_.Get(
		j,
		"filter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettingsOutputReference) FilterInput() *ChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettingsFilter {
	var returns *ChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettingsFilter
	_jsii_.Get(
		j,
		"filterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettingsOutputReference) InternalValue() *ChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettings {
	var returns *ChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettingsOutputReference) NewTab() any {
	var returns any
	_jsii_.Get(
		j,
		"newTab",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettingsOutputReference) NewTabInput() any {
	var returns any
	_jsii_.Get(
		j,
		"newTabInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettingsOutputReference) Query() ChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettingsQueryOutputReference {
	var returns ChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettingsQueryOutputReference
	_jsii_.Get(
		j,
		"query",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettingsOutputReference) QueryInput() *ChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettingsQuery {
	var returns *ChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettingsQuery
	_jsii_.Get(
		j,
		"queryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.chronicleDashboardChart.ChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettingsOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettingsOutputReference_Override(c ChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.chronicleDashboardChart.ChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettingsOutputReference",
		[]any{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettingsOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettingsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettingsOutputReference) SetInternalValue(val *ChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettingsOutputReference) SetNewTab(val any) {
	if err := j.validateSetNewTabParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"newTab",
		val,
	)
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettingsOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettingsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettingsOutputReference) PutExternalLink(value *ChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettingsExternalLink) {
	if err := c.validatePutExternalLinkParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putExternalLink",
		[]any{value},
	)
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettingsOutputReference) PutFilter(value *ChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettingsFilter) {
	if err := c.validatePutFilterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putFilter",
		[]any{value},
	)
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettingsOutputReference) PutQuery(value *ChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettingsQuery) {
	if err := c.validatePutQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putQuery",
		[]any{value},
	)
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettingsOutputReference) ResetExternalLink() {
	_jsii_.InvokeVoid(
		c,
		"resetExternalLink",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettingsOutputReference) ResetFilter() {
	_jsii_.InvokeVoid(
		c,
		"resetFilter",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettingsOutputReference) ResetQuery() {
	_jsii_.InvokeVoid(
		c,
		"resetQuery",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettingsOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (c *jsiiProxy_ChronicleDashboardChartDashboardChartDrillDownConfigRightDrillDownsCustomSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
