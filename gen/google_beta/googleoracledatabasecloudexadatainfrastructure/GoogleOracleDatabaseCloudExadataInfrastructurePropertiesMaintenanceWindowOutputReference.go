package googleoracledatabasecloudexadatainfrastructure

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleoracledatabasecloudexadatainfrastructure/internal"
)

type GoogleOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference interface {
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
	CustomActionTimeoutMins() *float64
	SetCustomActionTimeoutMins(val *float64)
	CustomActionTimeoutMinsInput() *float64
	DaysOfWeek() *[]*string
	SetDaysOfWeek(val *[]*string)
	DaysOfWeekInput() *[]*string
	// Experimental.
	Fqn() *string
	HoursOfDay() *[]*float64
	SetHoursOfDay(val *[]*float64)
	HoursOfDayInput() *[]*float64
	InternalValue() *GoogleOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindow
	SetInternalValue(val *GoogleOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindow)
	IsCustomActionTimeoutEnabled() interface{}
	SetIsCustomActionTimeoutEnabled(val interface{})
	IsCustomActionTimeoutEnabledInput() interface{}
	LeadTimeWeek() *float64
	SetLeadTimeWeek(val *float64)
	LeadTimeWeekInput() *float64
	Months() *[]*string
	SetMonths(val *[]*string)
	MonthsInput() *[]*string
	PatchingMode() *string
	SetPatchingMode(val *string)
	PatchingModeInput() *string
	Preference() *string
	SetPreference(val *string)
	PreferenceInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WeeksOfMonth() *[]*float64
	SetWeeksOfMonth(val *[]*float64)
	WeeksOfMonthInput() *[]*float64
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
	ResetCustomActionTimeoutMins()
	ResetDaysOfWeek()
	ResetHoursOfDay()
	ResetIsCustomActionTimeoutEnabled()
	ResetLeadTimeWeek()
	ResetMonths()
	ResetPatchingMode()
	ResetPreference()
	ResetWeeksOfMonth()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference
type jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference) CustomActionTimeoutMins() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"customActionTimeoutMins",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference) CustomActionTimeoutMinsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"customActionTimeoutMinsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference) DaysOfWeek() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"daysOfWeek",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference) DaysOfWeekInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"daysOfWeekInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference) HoursOfDay() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"hoursOfDay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference) HoursOfDayInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"hoursOfDayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference) InternalValue() *GoogleOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindow {
	var returns *GoogleOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindow
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference) IsCustomActionTimeoutEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isCustomActionTimeoutEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference) IsCustomActionTimeoutEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isCustomActionTimeoutEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference) LeadTimeWeek() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"leadTimeWeek",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference) LeadTimeWeekInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"leadTimeWeekInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference) Months() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"months",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference) MonthsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"monthsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference) PatchingMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"patchingMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference) PatchingModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"patchingModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference) Preference() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference) PreferenceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference) WeeksOfMonth() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"weeksOfMonth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference) WeeksOfMonthInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"weeksOfMonthInput",
		&returns,
	)
	return returns
}


func NewGoogleOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleOracleDatabaseCloudExadataInfrastructure.GoogleOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference_Override(g GoogleOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleOracleDatabaseCloudExadataInfrastructure.GoogleOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference)SetCustomActionTimeoutMins(val *float64) {
	if err := j.validateSetCustomActionTimeoutMinsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customActionTimeoutMins",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference)SetDaysOfWeek(val *[]*string) {
	if err := j.validateSetDaysOfWeekParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"daysOfWeek",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference)SetHoursOfDay(val *[]*float64) {
	if err := j.validateSetHoursOfDayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hoursOfDay",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference)SetInternalValue(val *GoogleOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindow) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference)SetIsCustomActionTimeoutEnabled(val interface{}) {
	if err := j.validateSetIsCustomActionTimeoutEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isCustomActionTimeoutEnabled",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference)SetLeadTimeWeek(val *float64) {
	if err := j.validateSetLeadTimeWeekParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"leadTimeWeek",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference)SetMonths(val *[]*string) {
	if err := j.validateSetMonthsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"months",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference)SetPatchingMode(val *string) {
	if err := j.validateSetPatchingModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"patchingMode",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference)SetPreference(val *string) {
	if err := j.validateSetPreferenceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preference",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference)SetWeeksOfMonth(val *[]*float64) {
	if err := j.validateSetWeeksOfMonthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"weeksOfMonth",
		val,
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference) ResetCustomActionTimeoutMins() {
	_jsii_.InvokeVoid(
		g,
		"resetCustomActionTimeoutMins",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference) ResetDaysOfWeek() {
	_jsii_.InvokeVoid(
		g,
		"resetDaysOfWeek",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference) ResetHoursOfDay() {
	_jsii_.InvokeVoid(
		g,
		"resetHoursOfDay",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference) ResetIsCustomActionTimeoutEnabled() {
	_jsii_.InvokeVoid(
		g,
		"resetIsCustomActionTimeoutEnabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference) ResetLeadTimeWeek() {
	_jsii_.InvokeVoid(
		g,
		"resetLeadTimeWeek",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference) ResetMonths() {
	_jsii_.InvokeVoid(
		g,
		"resetMonths",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference) ResetPatchingMode() {
	_jsii_.InvokeVoid(
		g,
		"resetPatchingMode",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference) ResetPreference() {
	_jsii_.InvokeVoid(
		g,
		"resetPreference",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference) ResetWeeksOfMonth() {
	_jsii_.InvokeVoid(
		g,
		"resetWeeksOfMonth",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

