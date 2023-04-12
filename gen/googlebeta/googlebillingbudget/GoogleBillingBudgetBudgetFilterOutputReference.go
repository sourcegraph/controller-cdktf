package googlebillingbudget

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/googlebeta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/googlebeta/googlebillingbudget/internal"
)

type GoogleBillingBudgetBudgetFilterOutputReference interface {
	cdktf.ComplexObject
	CalendarPeriod() *string
	SetCalendarPeriod(val *string)
	CalendarPeriodInput() *string
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
	CreditTypes() *[]*string
	SetCreditTypes(val *[]*string)
	CreditTypesInput() *[]*string
	CreditTypesTreatment() *string
	SetCreditTypesTreatment(val *string)
	CreditTypesTreatmentInput() *string
	CustomPeriod() GoogleBillingBudgetBudgetFilterCustomPeriodOutputReference
	CustomPeriodInput() *GoogleBillingBudgetBudgetFilterCustomPeriod
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleBillingBudgetBudgetFilter
	SetInternalValue(val *GoogleBillingBudgetBudgetFilter)
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	Projects() *[]*string
	SetProjects(val *[]*string)
	ProjectsInput() *[]*string
	Services() *[]*string
	SetServices(val *[]*string)
	ServicesInput() *[]*string
	Subaccounts() *[]*string
	SetSubaccounts(val *[]*string)
	SubaccountsInput() *[]*string
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
	PutCustomPeriod(value *GoogleBillingBudgetBudgetFilterCustomPeriod)
	ResetCalendarPeriod()
	ResetCreditTypes()
	ResetCreditTypesTreatment()
	ResetCustomPeriod()
	ResetLabels()
	ResetProjects()
	ResetServices()
	ResetSubaccounts()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleBillingBudgetBudgetFilterOutputReference
type jsiiProxy_GoogleBillingBudgetBudgetFilterOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleBillingBudgetBudgetFilterOutputReference) CalendarPeriod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"calendarPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBillingBudgetBudgetFilterOutputReference) CalendarPeriodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"calendarPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBillingBudgetBudgetFilterOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBillingBudgetBudgetFilterOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBillingBudgetBudgetFilterOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBillingBudgetBudgetFilterOutputReference) CreditTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creditTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBillingBudgetBudgetFilterOutputReference) CreditTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creditTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBillingBudgetBudgetFilterOutputReference) CreditTypesTreatment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creditTypesTreatment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBillingBudgetBudgetFilterOutputReference) CreditTypesTreatmentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creditTypesTreatmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBillingBudgetBudgetFilterOutputReference) CustomPeriod() GoogleBillingBudgetBudgetFilterCustomPeriodOutputReference {
	var returns GoogleBillingBudgetBudgetFilterCustomPeriodOutputReference
	_jsii_.Get(
		j,
		"customPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBillingBudgetBudgetFilterOutputReference) CustomPeriodInput() *GoogleBillingBudgetBudgetFilterCustomPeriod {
	var returns *GoogleBillingBudgetBudgetFilterCustomPeriod
	_jsii_.Get(
		j,
		"customPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBillingBudgetBudgetFilterOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBillingBudgetBudgetFilterOutputReference) InternalValue() *GoogleBillingBudgetBudgetFilter {
	var returns *GoogleBillingBudgetBudgetFilter
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBillingBudgetBudgetFilterOutputReference) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBillingBudgetBudgetFilterOutputReference) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBillingBudgetBudgetFilterOutputReference) Projects() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"projects",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBillingBudgetBudgetFilterOutputReference) ProjectsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"projectsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBillingBudgetBudgetFilterOutputReference) Services() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"services",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBillingBudgetBudgetFilterOutputReference) ServicesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"servicesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBillingBudgetBudgetFilterOutputReference) Subaccounts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subaccounts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBillingBudgetBudgetFilterOutputReference) SubaccountsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subaccountsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBillingBudgetBudgetFilterOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBillingBudgetBudgetFilterOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleBillingBudgetBudgetFilterOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleBillingBudgetBudgetFilterOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleBillingBudgetBudgetFilterOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleBillingBudgetBudgetFilterOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-googlebeta.googleBillingBudget.GoogleBillingBudgetBudgetFilterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleBillingBudgetBudgetFilterOutputReference_Override(g GoogleBillingBudgetBudgetFilterOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-googlebeta.googleBillingBudget.GoogleBillingBudgetBudgetFilterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleBillingBudgetBudgetFilterOutputReference)SetCalendarPeriod(val *string) {
	if err := j.validateSetCalendarPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"calendarPeriod",
		val,
	)
}

func (j *jsiiProxy_GoogleBillingBudgetBudgetFilterOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleBillingBudgetBudgetFilterOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleBillingBudgetBudgetFilterOutputReference)SetCreditTypes(val *[]*string) {
	if err := j.validateSetCreditTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"creditTypes",
		val,
	)
}

func (j *jsiiProxy_GoogleBillingBudgetBudgetFilterOutputReference)SetCreditTypesTreatment(val *string) {
	if err := j.validateSetCreditTypesTreatmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"creditTypesTreatment",
		val,
	)
}

func (j *jsiiProxy_GoogleBillingBudgetBudgetFilterOutputReference)SetInternalValue(val *GoogleBillingBudgetBudgetFilter) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleBillingBudgetBudgetFilterOutputReference)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_GoogleBillingBudgetBudgetFilterOutputReference)SetProjects(val *[]*string) {
	if err := j.validateSetProjectsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"projects",
		val,
	)
}

func (j *jsiiProxy_GoogleBillingBudgetBudgetFilterOutputReference)SetServices(val *[]*string) {
	if err := j.validateSetServicesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"services",
		val,
	)
}

func (j *jsiiProxy_GoogleBillingBudgetBudgetFilterOutputReference)SetSubaccounts(val *[]*string) {
	if err := j.validateSetSubaccountsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subaccounts",
		val,
	)
}

func (j *jsiiProxy_GoogleBillingBudgetBudgetFilterOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleBillingBudgetBudgetFilterOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleBillingBudgetBudgetFilterOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBillingBudgetBudgetFilterOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleBillingBudgetBudgetFilterOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleBillingBudgetBudgetFilterOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleBillingBudgetBudgetFilterOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleBillingBudgetBudgetFilterOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleBillingBudgetBudgetFilterOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleBillingBudgetBudgetFilterOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleBillingBudgetBudgetFilterOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleBillingBudgetBudgetFilterOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleBillingBudgetBudgetFilterOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBillingBudgetBudgetFilterOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleBillingBudgetBudgetFilterOutputReference) PutCustomPeriod(value *GoogleBillingBudgetBudgetFilterCustomPeriod) {
	if err := g.validatePutCustomPeriodParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putCustomPeriod",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBillingBudgetBudgetFilterOutputReference) ResetCalendarPeriod() {
	_jsii_.InvokeVoid(
		g,
		"resetCalendarPeriod",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBillingBudgetBudgetFilterOutputReference) ResetCreditTypes() {
	_jsii_.InvokeVoid(
		g,
		"resetCreditTypes",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBillingBudgetBudgetFilterOutputReference) ResetCreditTypesTreatment() {
	_jsii_.InvokeVoid(
		g,
		"resetCreditTypesTreatment",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBillingBudgetBudgetFilterOutputReference) ResetCustomPeriod() {
	_jsii_.InvokeVoid(
		g,
		"resetCustomPeriod",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBillingBudgetBudgetFilterOutputReference) ResetLabels() {
	_jsii_.InvokeVoid(
		g,
		"resetLabels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBillingBudgetBudgetFilterOutputReference) ResetProjects() {
	_jsii_.InvokeVoid(
		g,
		"resetProjects",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBillingBudgetBudgetFilterOutputReference) ResetServices() {
	_jsii_.InvokeVoid(
		g,
		"resetServices",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBillingBudgetBudgetFilterOutputReference) ResetSubaccounts() {
	_jsii_.InvokeVoid(
		g,
		"resetSubaccounts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBillingBudgetBudgetFilterOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleBillingBudgetBudgetFilterOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

