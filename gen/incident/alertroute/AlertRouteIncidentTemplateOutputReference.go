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
	CustomFields() AlertRouteIncidentTemplateCustomFieldsList
	CustomFieldsInput() interface{}
	// Experimental.
	Fqn() *string
	IncidentMode() AlertRouteIncidentTemplateIncidentModeOutputReference
	IncidentModeInput() interface{}
	IncidentType() AlertRouteIncidentTemplateIncidentTypeOutputReference
	IncidentTypeInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() AlertRouteIncidentTemplateNameOutputReference
	NameInput() interface{}
	Severity() AlertRouteIncidentTemplateSeverityOutputReference
	SeverityInput() interface{}
	StartInTriage() AlertRouteIncidentTemplateStartInTriageOutputReference
	StartInTriageInput() interface{}
	Summary() AlertRouteIncidentTemplateSummaryOutputReference
	SummaryInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Workspace() AlertRouteIncidentTemplateWorkspaceOutputReference
	WorkspaceInput() interface{}
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
	PutCustomFields(value interface{})
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
	Resolve(_context cdktf.IResolveContext) interface{}
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

func (j *jsiiProxy_AlertRouteIncidentTemplateOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
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

func (j *jsiiProxy_AlertRouteIncidentTemplateOutputReference) CustomFieldsInput() interface{} {
	var returns interface{}
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

func (j *jsiiProxy_AlertRouteIncidentTemplateOutputReference) IncidentModeInput() interface{} {
	var returns interface{}
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

func (j *jsiiProxy_AlertRouteIncidentTemplateOutputReference) IncidentTypeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"incidentTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertRouteIncidentTemplateOutputReference) InternalValue() interface{} {
	var returns interface{}
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

func (j *jsiiProxy_AlertRouteIncidentTemplateOutputReference) NameInput() interface{} {
	var returns interface{}
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

func (j *jsiiProxy_AlertRouteIncidentTemplateOutputReference) SeverityInput() interface{} {
	var returns interface{}
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

func (j *jsiiProxy_AlertRouteIncidentTemplateOutputReference) StartInTriageInput() interface{} {
	var returns interface{}
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

func (j *jsiiProxy_AlertRouteIncidentTemplateOutputReference) SummaryInput() interface{} {
	var returns interface{}
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

func (j *jsiiProxy_AlertRouteIncidentTemplateOutputReference) WorkspaceInput() interface{} {
	var returns interface{}
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
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAlertRouteIncidentTemplateOutputReference_Override(a AlertRouteIncidentTemplateOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-incident.alertRoute.AlertRouteIncidentTemplateOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AlertRouteIncidentTemplateOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AlertRouteIncidentTemplateOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AlertRouteIncidentTemplateOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AlertRouteIncidentTemplateOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AlertRouteIncidentTemplateOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
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

func (a *jsiiProxy_AlertRouteIncidentTemplateOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
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
		[]interface{}{terraformAttribute},
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
		[]interface{}{terraformAttribute},
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
		[]interface{}{terraformAttribute},
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
		[]interface{}{terraformAttribute},
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
		[]interface{}{terraformAttribute},
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
		[]interface{}{terraformAttribute},
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
		[]interface{}{terraformAttribute},
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
		[]interface{}{terraformAttribute},
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
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertRouteIncidentTemplateOutputReference) PutCustomFields(value interface{}) {
	if err := a.validatePutCustomFieldsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCustomFields",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlertRouteIncidentTemplateOutputReference) PutIncidentMode(value *AlertRouteIncidentTemplateIncidentMode) {
	if err := a.validatePutIncidentModeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putIncidentMode",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlertRouteIncidentTemplateOutputReference) PutIncidentType(value *AlertRouteIncidentTemplateIncidentType) {
	if err := a.validatePutIncidentTypeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putIncidentType",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlertRouteIncidentTemplateOutputReference) PutName(value *AlertRouteIncidentTemplateName) {
	if err := a.validatePutNameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putName",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlertRouteIncidentTemplateOutputReference) PutSeverity(value *AlertRouteIncidentTemplateSeverity) {
	if err := a.validatePutSeverityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSeverity",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlertRouteIncidentTemplateOutputReference) PutStartInTriage(value *AlertRouteIncidentTemplateStartInTriage) {
	if err := a.validatePutStartInTriageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putStartInTriage",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlertRouteIncidentTemplateOutputReference) PutSummary(value *AlertRouteIncidentTemplateSummary) {
	if err := a.validatePutSummaryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSummary",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlertRouteIncidentTemplateOutputReference) PutWorkspace(value *AlertRouteIncidentTemplateWorkspace) {
	if err := a.validatePutWorkspaceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putWorkspace",
		[]interface{}{value},
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

func (a *jsiiProxy_AlertRouteIncidentTemplateOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
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

