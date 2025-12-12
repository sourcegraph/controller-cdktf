package dataincidentalertsources

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/incident/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/incident/dataincidentalertsources/internal"
)

type DataIncidentAlertSourcesAlertSourcesTemplateExpressionsOperationsBranchesBranchesConditionGroupsConditionsOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() *DataIncidentAlertSourcesAlertSourcesTemplateExpressionsOperationsBranchesBranchesConditionGroupsConditions
	SetInternalValue(val *DataIncidentAlertSourcesAlertSourcesTemplateExpressionsOperationsBranchesBranchesConditionGroupsConditions)
	Operation() *string
	ParamBindings() DataIncidentAlertSourcesAlertSourcesTemplateExpressionsOperationsBranchesBranchesConditionGroupsConditionsParamBindingsList
	Subject() *string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataIncidentAlertSourcesAlertSourcesTemplateExpressionsOperationsBranchesBranchesConditionGroupsConditionsOutputReference
type jsiiProxy_DataIncidentAlertSourcesAlertSourcesTemplateExpressionsOperationsBranchesBranchesConditionGroupsConditionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataIncidentAlertSourcesAlertSourcesTemplateExpressionsOperationsBranchesBranchesConditionGroupsConditionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIncidentAlertSourcesAlertSourcesTemplateExpressionsOperationsBranchesBranchesConditionGroupsConditionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIncidentAlertSourcesAlertSourcesTemplateExpressionsOperationsBranchesBranchesConditionGroupsConditionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIncidentAlertSourcesAlertSourcesTemplateExpressionsOperationsBranchesBranchesConditionGroupsConditionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIncidentAlertSourcesAlertSourcesTemplateExpressionsOperationsBranchesBranchesConditionGroupsConditionsOutputReference) InternalValue() *DataIncidentAlertSourcesAlertSourcesTemplateExpressionsOperationsBranchesBranchesConditionGroupsConditions {
	var returns *DataIncidentAlertSourcesAlertSourcesTemplateExpressionsOperationsBranchesBranchesConditionGroupsConditions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIncidentAlertSourcesAlertSourcesTemplateExpressionsOperationsBranchesBranchesConditionGroupsConditionsOutputReference) Operation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIncidentAlertSourcesAlertSourcesTemplateExpressionsOperationsBranchesBranchesConditionGroupsConditionsOutputReference) ParamBindings() DataIncidentAlertSourcesAlertSourcesTemplateExpressionsOperationsBranchesBranchesConditionGroupsConditionsParamBindingsList {
	var returns DataIncidentAlertSourcesAlertSourcesTemplateExpressionsOperationsBranchesBranchesConditionGroupsConditionsParamBindingsList
	_jsii_.Get(
		j,
		"paramBindings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIncidentAlertSourcesAlertSourcesTemplateExpressionsOperationsBranchesBranchesConditionGroupsConditionsOutputReference) Subject() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIncidentAlertSourcesAlertSourcesTemplateExpressionsOperationsBranchesBranchesConditionGroupsConditionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIncidentAlertSourcesAlertSourcesTemplateExpressionsOperationsBranchesBranchesConditionGroupsConditionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataIncidentAlertSourcesAlertSourcesTemplateExpressionsOperationsBranchesBranchesConditionGroupsConditionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataIncidentAlertSourcesAlertSourcesTemplateExpressionsOperationsBranchesBranchesConditionGroupsConditionsOutputReference {
	_init_.Initialize()

	if err := validateNewDataIncidentAlertSourcesAlertSourcesTemplateExpressionsOperationsBranchesBranchesConditionGroupsConditionsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataIncidentAlertSourcesAlertSourcesTemplateExpressionsOperationsBranchesBranchesConditionGroupsConditionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-incident.dataIncidentAlertSources.DataIncidentAlertSourcesAlertSourcesTemplateExpressionsOperationsBranchesBranchesConditionGroupsConditionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataIncidentAlertSourcesAlertSourcesTemplateExpressionsOperationsBranchesBranchesConditionGroupsConditionsOutputReference_Override(d DataIncidentAlertSourcesAlertSourcesTemplateExpressionsOperationsBranchesBranchesConditionGroupsConditionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-incident.dataIncidentAlertSources.DataIncidentAlertSourcesAlertSourcesTemplateExpressionsOperationsBranchesBranchesConditionGroupsConditionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataIncidentAlertSourcesAlertSourcesTemplateExpressionsOperationsBranchesBranchesConditionGroupsConditionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataIncidentAlertSourcesAlertSourcesTemplateExpressionsOperationsBranchesBranchesConditionGroupsConditionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataIncidentAlertSourcesAlertSourcesTemplateExpressionsOperationsBranchesBranchesConditionGroupsConditionsOutputReference)SetInternalValue(val *DataIncidentAlertSourcesAlertSourcesTemplateExpressionsOperationsBranchesBranchesConditionGroupsConditions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataIncidentAlertSourcesAlertSourcesTemplateExpressionsOperationsBranchesBranchesConditionGroupsConditionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataIncidentAlertSourcesAlertSourcesTemplateExpressionsOperationsBranchesBranchesConditionGroupsConditionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataIncidentAlertSourcesAlertSourcesTemplateExpressionsOperationsBranchesBranchesConditionGroupsConditionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataIncidentAlertSourcesAlertSourcesTemplateExpressionsOperationsBranchesBranchesConditionGroupsConditionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataIncidentAlertSourcesAlertSourcesTemplateExpressionsOperationsBranchesBranchesConditionGroupsConditionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataIncidentAlertSourcesAlertSourcesTemplateExpressionsOperationsBranchesBranchesConditionGroupsConditionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataIncidentAlertSourcesAlertSourcesTemplateExpressionsOperationsBranchesBranchesConditionGroupsConditionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataIncidentAlertSourcesAlertSourcesTemplateExpressionsOperationsBranchesBranchesConditionGroupsConditionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataIncidentAlertSourcesAlertSourcesTemplateExpressionsOperationsBranchesBranchesConditionGroupsConditionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataIncidentAlertSourcesAlertSourcesTemplateExpressionsOperationsBranchesBranchesConditionGroupsConditionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataIncidentAlertSourcesAlertSourcesTemplateExpressionsOperationsBranchesBranchesConditionGroupsConditionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataIncidentAlertSourcesAlertSourcesTemplateExpressionsOperationsBranchesBranchesConditionGroupsConditionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataIncidentAlertSourcesAlertSourcesTemplateExpressionsOperationsBranchesBranchesConditionGroupsConditionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataIncidentAlertSourcesAlertSourcesTemplateExpressionsOperationsBranchesBranchesConditionGroupsConditionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataIncidentAlertSourcesAlertSourcesTemplateExpressionsOperationsBranchesBranchesConditionGroupsConditionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataIncidentAlertSourcesAlertSourcesTemplateExpressionsOperationsBranchesBranchesConditionGroupsConditionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

