package dataincidentalertsources

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/incident/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/incident/dataincidentalertsources/internal"
)

type DataIncidentAlertSourcesAlertSourcesTemplateExpressionsOperationsBranchesBranchesConditionGroupsConditionsParamBindingsList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Creating an iterator for this complex list.
	//
	// The list will be converted into a map with the mapKeyAttributeName as the key.
	// Experimental.
	AllWithMapKey(mapKeyAttributeName *string) cdktf.DynamicListTerraformIterator
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) DataIncidentAlertSourcesAlertSourcesTemplateExpressionsOperationsBranchesBranchesConditionGroupsConditionsParamBindingsOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataIncidentAlertSourcesAlertSourcesTemplateExpressionsOperationsBranchesBranchesConditionGroupsConditionsParamBindingsList
type jsiiProxy_DataIncidentAlertSourcesAlertSourcesTemplateExpressionsOperationsBranchesBranchesConditionGroupsConditionsParamBindingsList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_DataIncidentAlertSourcesAlertSourcesTemplateExpressionsOperationsBranchesBranchesConditionGroupsConditionsParamBindingsList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIncidentAlertSourcesAlertSourcesTemplateExpressionsOperationsBranchesBranchesConditionGroupsConditionsParamBindingsList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIncidentAlertSourcesAlertSourcesTemplateExpressionsOperationsBranchesBranchesConditionGroupsConditionsParamBindingsList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIncidentAlertSourcesAlertSourcesTemplateExpressionsOperationsBranchesBranchesConditionGroupsConditionsParamBindingsList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataIncidentAlertSourcesAlertSourcesTemplateExpressionsOperationsBranchesBranchesConditionGroupsConditionsParamBindingsList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}

func NewDataIncidentAlertSourcesAlertSourcesTemplateExpressionsOperationsBranchesBranchesConditionGroupsConditionsParamBindingsList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) DataIncidentAlertSourcesAlertSourcesTemplateExpressionsOperationsBranchesBranchesConditionGroupsConditionsParamBindingsList {
	_init_.Initialize()

	if err := validateNewDataIncidentAlertSourcesAlertSourcesTemplateExpressionsOperationsBranchesBranchesConditionGroupsConditionsParamBindingsListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataIncidentAlertSourcesAlertSourcesTemplateExpressionsOperationsBranchesBranchesConditionGroupsConditionsParamBindingsList{}

	_jsii_.Create(
		"@cdktf/provider-incident.dataIncidentAlertSources.DataIncidentAlertSourcesAlertSourcesTemplateExpressionsOperationsBranchesBranchesConditionGroupsConditionsParamBindingsList",
		[]any{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewDataIncidentAlertSourcesAlertSourcesTemplateExpressionsOperationsBranchesBranchesConditionGroupsConditionsParamBindingsList_Override(d DataIncidentAlertSourcesAlertSourcesTemplateExpressionsOperationsBranchesBranchesConditionGroupsConditionsParamBindingsList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-incident.dataIncidentAlertSources.DataIncidentAlertSourcesAlertSourcesTemplateExpressionsOperationsBranchesBranchesConditionGroupsConditionsParamBindingsList",
		[]any{terraformResource, terraformAttribute, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DataIncidentAlertSourcesAlertSourcesTemplateExpressionsOperationsBranchesBranchesConditionGroupsConditionsParamBindingsList) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataIncidentAlertSourcesAlertSourcesTemplateExpressionsOperationsBranchesBranchesConditionGroupsConditionsParamBindingsList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataIncidentAlertSourcesAlertSourcesTemplateExpressionsOperationsBranchesBranchesConditionGroupsConditionsParamBindingsList) SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (d *jsiiProxy_DataIncidentAlertSourcesAlertSourcesTemplateExpressionsOperationsBranchesBranchesConditionGroupsConditionsParamBindingsList) AllWithMapKey(mapKeyAttributeName *string) cdktf.DynamicListTerraformIterator {
	if err := d.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktf.DynamicListTerraformIterator

	_jsii_.Invoke(
		d,
		"allWithMapKey",
		[]any{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataIncidentAlertSourcesAlertSourcesTemplateExpressionsOperationsBranchesBranchesConditionGroupsConditionsParamBindingsList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataIncidentAlertSourcesAlertSourcesTemplateExpressionsOperationsBranchesBranchesConditionGroupsConditionsParamBindingsList) Get(index *float64) DataIncidentAlertSourcesAlertSourcesTemplateExpressionsOperationsBranchesBranchesConditionGroupsConditionsParamBindingsOutputReference {
	if err := d.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns DataIncidentAlertSourcesAlertSourcesTemplateExpressionsOperationsBranchesBranchesConditionGroupsConditionsParamBindingsOutputReference

	_jsii_.Invoke(
		d,
		"get",
		[]any{index},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataIncidentAlertSourcesAlertSourcesTemplateExpressionsOperationsBranchesBranchesConditionGroupsConditionsParamBindingsList) Resolve(_context cdktf.IResolveContext) any {
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

func (d *jsiiProxy_DataIncidentAlertSourcesAlertSourcesTemplateExpressionsOperationsBranchesBranchesConditionGroupsConditionsParamBindingsList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
