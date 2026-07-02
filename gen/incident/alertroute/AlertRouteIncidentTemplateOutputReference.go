package alertroute

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/incident/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/incident/alertroute/internal"
)

type AlertRouteIncidentTemplateOutputReference interface {
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
	CustomFields() AlertRouteIncidentTemplateCustomFieldsList
	CustomFieldsInput() any
	// Experimental.
	Fqn() *string
	IncidentMode() AlertRouteIncidentTemplateIncidentModeOutputReference
	IncidentModeInput() any
	IncidentType() AlertRouteIncidentTemplateIncidentTypeOutputReference
	IncidentTypeInput() any
	InternalValue() any
	SetInternalValue(val any)
	Name() AlertRouteIncidentTemplateNameOutputReference
	NameInput() any
	Severity() AlertRouteIncidentTemplateSeverityOutputReference
	SeverityInput() any
	StartInTriage() AlertRouteIncidentTemplateStartInTriageOutputReference
	StartInTriageInput() any
	Summary() AlertRouteIncidentTemplateSummaryOutputReference
	SummaryInput() any
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Workspace() AlertRouteIncidentTemplateWorkspaceOutputReference
	WorkspaceInput() any
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
	PutCustomFields(value any)
	PutIncidentMode(value *AlertRouteIncidentTemplateIncidentMode)
	PutIncidentType(value *AlertRouteIncidentTemplateIncidentType)
	PutName(value *AlertRouteIncidentTemplateName)
	PutSeverity(value *AlertRouteIncidentTemplateSeverity)
	PutStartInTriage(value *AlertRouteIncidentTemplateStartInTriage)
	PutSummary(value *AlertRouteIncidentTemplateSummary)
	PutWorkspace(value *AlertRouteIncidentTemplateWorkspace)
	ResetCustomFields()
	ResetIncidentMode()
	ResetIncidentType()
	ResetSeverity()
	ResetStartInTriage()
	ResetWorkspace()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AlertRouteIncidentTemplateOutputReference
type jsiiProxy_AlertRouteIncidentTemplateOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AlertRouteIncidentTemplateOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertRouteIncidentTemplateOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertRouteIncidentTemplateOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertRouteIncidentTemplateOutputReference) CustomFields() AlertRouteIncidentTemplateCustomFieldsList {
	var returns AlertRouteIncidentTemplateCustomFieldsList
	_jsii_.Get(
		j,
		"customFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertRouteIncidentTemplateOutputReference) CustomFieldsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"customFieldsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertRouteIncidentTemplateOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertRouteIncidentTemplateOutputReference) IncidentMode() AlertRouteIncidentTemplateIncidentModeOutputReference {
	var returns AlertRouteIncidentTemplateIncidentModeOutputReference
	_jsii_.Get(
		j,
		"incidentMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertRouteIncidentTemplateOutputReference) IncidentModeInput() any {
	var returns any
	_jsii_.Get(
		j,
		"incidentModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertRouteIncidentTemplateOutputReference) IncidentType() AlertRouteIncidentTemplateIncidentTypeOutputReference {
	var returns AlertRouteIncidentTemplateIncidentTypeOutputReference
	_jsii_.Get(
		j,
		"incidentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertRouteIncidentTemplateOutputReference) IncidentTypeInput() any {
	var returns any
	_jsii_.Get(
		j,
		"incidentTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertRouteIncidentTemplateOutputReference) InternalValue() any {
	var returns any
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertRouteIncidentTemplateOutputReference) Name() AlertRouteIncidentTemplateNameOutputReference {
	var returns AlertRouteIncidentTemplateNameOutputReference
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertRouteIncidentTemplateOutputReference) NameInput() any {
	var returns any
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertRouteIncidentTemplateOutputReference) Severity() AlertRouteIncidentTemplateSeverityOutputReference {
	var returns AlertRouteIncidentTemplateSeverityOutputReference
	_jsii_.Get(
		j,
		"severity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertRouteIncidentTemplateOutputReference) SeverityInput() any {
	var returns any
	_jsii_.Get(
		j,
		"severityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertRouteIncidentTemplateOutputReference) StartInTriage() AlertRouteIncidentTemplateStartInTriageOutputReference {
	var returns AlertRouteIncidentTemplateStartInTriageOutputReference
	_jsii_.Get(
		j,
		"startInTriage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertRouteIncidentTemplateOutputReference) StartInTriageInput() any {
	var returns any
	_jsii_.Get(
		j,
		"startInTriageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertRouteIncidentTemplateOutputReference) Summary() AlertRouteIncidentTemplateSummaryOutputReference {
	var returns AlertRouteIncidentTemplateSummaryOutputReference
	_jsii_.Get(
		j,
		"summary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertRouteIncidentTemplateOutputReference) SummaryInput() any {
	var returns any
	_jsii_.Get(
		j,
		"summaryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertRouteIncidentTemplateOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertRouteIncidentTemplateOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertRouteIncidentTemplateOutputReference) Workspace() AlertRouteIncidentTemplateWorkspaceOutputReference {
	var returns AlertRouteIncidentTemplateWorkspaceOutputReference
	_jsii_.Get(
		j,
		"workspace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertRouteIncidentTemplateOutputReference) WorkspaceInput() any {
	var returns any
	_jsii_.Get(
		j,
		"workspaceInput",
		&returns,
	)
	return returns
}

func NewAlertRouteIncidentTemplateOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AlertRouteIncidentTemplateOutputReference {
	_init_.Initialize()

	if err := validateNewAlertRouteIncidentTemplateOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AlertRouteIncidentTemplateOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-incident.alertRoute.AlertRouteIncidentTemplateOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAlertRouteIncidentTemplateOutputReference_Override(a AlertRouteIncidentTemplateOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-incident.alertRoute.AlertRouteIncidentTemplateOutputReference",
		[]any{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AlertRouteIncidentTemplateOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AlertRouteIncidentTemplateOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AlertRouteIncidentTemplateOutputReference) SetInternalValue(val any) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AlertRouteIncidentTemplateOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AlertRouteIncidentTemplateOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AlertRouteIncidentTemplateOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertRouteIncidentTemplateOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertRouteIncidentTemplateOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertRouteIncidentTemplateOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertRouteIncidentTemplateOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertRouteIncidentTemplateOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertRouteIncidentTemplateOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertRouteIncidentTemplateOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertRouteIncidentTemplateOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertRouteIncidentTemplateOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertRouteIncidentTemplateOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertRouteIncidentTemplateOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertRouteIncidentTemplateOutputReference) PutCustomFields(value any) {
	if err := a.validatePutCustomFieldsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCustomFields",
		[]any{value},
	)
}

func (a *jsiiProxy_AlertRouteIncidentTemplateOutputReference) PutIncidentMode(value *AlertRouteIncidentTemplateIncidentMode) {
	if err := a.validatePutIncidentModeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putIncidentMode",
		[]any{value},
	)
}

func (a *jsiiProxy_AlertRouteIncidentTemplateOutputReference) PutIncidentType(value *AlertRouteIncidentTemplateIncidentType) {
	if err := a.validatePutIncidentTypeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putIncidentType",
		[]any{value},
	)
}

func (a *jsiiProxy_AlertRouteIncidentTemplateOutputReference) PutName(value *AlertRouteIncidentTemplateName) {
	if err := a.validatePutNameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putName",
		[]any{value},
	)
}

func (a *jsiiProxy_AlertRouteIncidentTemplateOutputReference) PutSeverity(value *AlertRouteIncidentTemplateSeverity) {
	if err := a.validatePutSeverityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSeverity",
		[]any{value},
	)
}

func (a *jsiiProxy_AlertRouteIncidentTemplateOutputReference) PutStartInTriage(value *AlertRouteIncidentTemplateStartInTriage) {
	if err := a.validatePutStartInTriageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putStartInTriage",
		[]any{value},
	)
}

func (a *jsiiProxy_AlertRouteIncidentTemplateOutputReference) PutSummary(value *AlertRouteIncidentTemplateSummary) {
	if err := a.validatePutSummaryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSummary",
		[]any{value},
	)
}

func (a *jsiiProxy_AlertRouteIncidentTemplateOutputReference) PutWorkspace(value *AlertRouteIncidentTemplateWorkspace) {
	if err := a.validatePutWorkspaceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putWorkspace",
		[]any{value},
	)
}

func (a *jsiiProxy_AlertRouteIncidentTemplateOutputReference) ResetCustomFields() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomFields",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertRouteIncidentTemplateOutputReference) ResetIncidentMode() {
	_jsii_.InvokeVoid(
		a,
		"resetIncidentMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertRouteIncidentTemplateOutputReference) ResetIncidentType() {
	_jsii_.InvokeVoid(
		a,
		"resetIncidentType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertRouteIncidentTemplateOutputReference) ResetSeverity() {
	_jsii_.InvokeVoid(
		a,
		"resetSeverity",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertRouteIncidentTemplateOutputReference) ResetStartInTriage() {
	_jsii_.InvokeVoid(
		a,
		"resetStartInTriage",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertRouteIncidentTemplateOutputReference) ResetWorkspace() {
	_jsii_.InvokeVoid(
		a,
		"resetWorkspace",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertRouteIncidentTemplateOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := a.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		a,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertRouteIncidentTemplateOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
