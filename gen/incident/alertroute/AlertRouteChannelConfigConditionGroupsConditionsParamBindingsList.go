package alertroute

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/incident/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/incident/alertroute/internal"
)

type AlertRouteChannelConfigConditionGroupsConditionsParamBindingsList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() any
	SetInternalValue(val any)
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
	Get(index *float64) AlertRouteChannelConfigConditionGroupsConditionsParamBindingsOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AlertRouteChannelConfigConditionGroupsConditionsParamBindingsList
type jsiiProxy_AlertRouteChannelConfigConditionGroupsConditionsParamBindingsList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_AlertRouteChannelConfigConditionGroupsConditionsParamBindingsList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertRouteChannelConfigConditionGroupsConditionsParamBindingsList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertRouteChannelConfigConditionGroupsConditionsParamBindingsList) InternalValue() any {
	var returns any
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertRouteChannelConfigConditionGroupsConditionsParamBindingsList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertRouteChannelConfigConditionGroupsConditionsParamBindingsList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertRouteChannelConfigConditionGroupsConditionsParamBindingsList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}

func NewAlertRouteChannelConfigConditionGroupsConditionsParamBindingsList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) AlertRouteChannelConfigConditionGroupsConditionsParamBindingsList {
	_init_.Initialize()

	if err := validateNewAlertRouteChannelConfigConditionGroupsConditionsParamBindingsListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AlertRouteChannelConfigConditionGroupsConditionsParamBindingsList{}

	_jsii_.Create(
		"@cdktf/provider-incident.alertRoute.AlertRouteChannelConfigConditionGroupsConditionsParamBindingsList",
		[]any{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewAlertRouteChannelConfigConditionGroupsConditionsParamBindingsList_Override(a AlertRouteChannelConfigConditionGroupsConditionsParamBindingsList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-incident.alertRoute.AlertRouteChannelConfigConditionGroupsConditionsParamBindingsList",
		[]any{terraformResource, terraformAttribute, wrapsSet},
		a,
	)
}

func (j *jsiiProxy_AlertRouteChannelConfigConditionGroupsConditionsParamBindingsList) SetInternalValue(val any) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AlertRouteChannelConfigConditionGroupsConditionsParamBindingsList) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AlertRouteChannelConfigConditionGroupsConditionsParamBindingsList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AlertRouteChannelConfigConditionGroupsConditionsParamBindingsList) SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (a *jsiiProxy_AlertRouteChannelConfigConditionGroupsConditionsParamBindingsList) AllWithMapKey(mapKeyAttributeName *string) cdktf.DynamicListTerraformIterator {
	if err := a.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktf.DynamicListTerraformIterator

	_jsii_.Invoke(
		a,
		"allWithMapKey",
		[]any{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertRouteChannelConfigConditionGroupsConditionsParamBindingsList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertRouteChannelConfigConditionGroupsConditionsParamBindingsList) Get(index *float64) AlertRouteChannelConfigConditionGroupsConditionsParamBindingsOutputReference {
	if err := a.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns AlertRouteChannelConfigConditionGroupsConditionsParamBindingsOutputReference

	_jsii_.Invoke(
		a,
		"get",
		[]any{index},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertRouteChannelConfigConditionGroupsConditionsParamBindingsList) Resolve(_context cdktf.IResolveContext) any {
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

func (a *jsiiProxy_AlertRouteChannelConfigConditionGroupsConditionsParamBindingsList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
