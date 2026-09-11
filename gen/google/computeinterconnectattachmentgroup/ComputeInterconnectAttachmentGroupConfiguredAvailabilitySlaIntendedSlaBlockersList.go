package computeinterconnectattachmentgroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google/computeinterconnectattachmentgroup/internal"
)

type ComputeInterconnectAttachmentGroupConfiguredAvailabilitySlaIntendedSlaBlockersList interface {
	cdktn.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	WrapsSet() *bool
	// Experimental.
	SetWrapsSet(val *bool)
	// Creating an iterator for this complex list.
	//
	// The list will be converted into a map with the mapKeyAttributeName as the key.
	// Experimental.
	AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) ComputeInterconnectAttachmentGroupConfiguredAvailabilitySlaIntendedSlaBlockersOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComputeInterconnectAttachmentGroupConfiguredAvailabilitySlaIntendedSlaBlockersList
type jsiiProxy_ComputeInterconnectAttachmentGroupConfiguredAvailabilitySlaIntendedSlaBlockersList struct {
	internal.Type__cdktnComplexList
}

func (j *jsiiProxy_ComputeInterconnectAttachmentGroupConfiguredAvailabilitySlaIntendedSlaBlockersList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectAttachmentGroupConfiguredAvailabilitySlaIntendedSlaBlockersList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectAttachmentGroupConfiguredAvailabilitySlaIntendedSlaBlockersList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectAttachmentGroupConfiguredAvailabilitySlaIntendedSlaBlockersList) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectAttachmentGroupConfiguredAvailabilitySlaIntendedSlaBlockersList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewComputeInterconnectAttachmentGroupConfiguredAvailabilitySlaIntendedSlaBlockersList(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) ComputeInterconnectAttachmentGroupConfiguredAvailabilitySlaIntendedSlaBlockersList {
	_init_.Initialize()

	if err := validateNewComputeInterconnectAttachmentGroupConfiguredAvailabilitySlaIntendedSlaBlockersListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeInterconnectAttachmentGroupConfiguredAvailabilitySlaIntendedSlaBlockersList{}

	_jsii_.Create(
		"@cdktn/provider-google.computeInterconnectAttachmentGroup.ComputeInterconnectAttachmentGroupConfiguredAvailabilitySlaIntendedSlaBlockersList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewComputeInterconnectAttachmentGroupConfiguredAvailabilitySlaIntendedSlaBlockersList_Override(c ComputeInterconnectAttachmentGroupConfiguredAvailabilitySlaIntendedSlaBlockersList, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.computeInterconnectAttachmentGroup.ComputeInterconnectAttachmentGroupConfiguredAvailabilitySlaIntendedSlaBlockersList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		c,
	)
}

func (j *jsiiProxy_ComputeInterconnectAttachmentGroupConfiguredAvailabilitySlaIntendedSlaBlockersList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComputeInterconnectAttachmentGroupConfiguredAvailabilitySlaIntendedSlaBlockersList)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ComputeInterconnectAttachmentGroupConfiguredAvailabilitySlaIntendedSlaBlockersList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (c *jsiiProxy_ComputeInterconnectAttachmentGroupConfiguredAvailabilitySlaIntendedSlaBlockersList) AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator {
	if err := c.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktn.DynamicListTerraformIterator

	_jsii_.Invoke(
		c,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInterconnectAttachmentGroupConfiguredAvailabilitySlaIntendedSlaBlockersList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInterconnectAttachmentGroupConfiguredAvailabilitySlaIntendedSlaBlockersList) Get(index *float64) ComputeInterconnectAttachmentGroupConfiguredAvailabilitySlaIntendedSlaBlockersOutputReference {
	if err := c.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns ComputeInterconnectAttachmentGroupConfiguredAvailabilitySlaIntendedSlaBlockersOutputReference

	_jsii_.Invoke(
		c,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInterconnectAttachmentGroupConfiguredAvailabilitySlaIntendedSlaBlockersList) Resolve(context cdktn.IResolveContext) interface{} {
	if err := c.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInterconnectAttachmentGroupConfiguredAvailabilitySlaIntendedSlaBlockersList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

