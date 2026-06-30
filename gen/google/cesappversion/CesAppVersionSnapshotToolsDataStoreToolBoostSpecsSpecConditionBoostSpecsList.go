package cesappversion

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/cesappversion/internal"
)

type CesAppVersionSnapshotToolsDataStoreToolBoostSpecsSpecConditionBoostSpecsList interface {
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
	Get(index *float64) CesAppVersionSnapshotToolsDataStoreToolBoostSpecsSpecConditionBoostSpecsOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CesAppVersionSnapshotToolsDataStoreToolBoostSpecsSpecConditionBoostSpecsList
type jsiiProxy_CesAppVersionSnapshotToolsDataStoreToolBoostSpecsSpecConditionBoostSpecsList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_CesAppVersionSnapshotToolsDataStoreToolBoostSpecsSpecConditionBoostSpecsList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppVersionSnapshotToolsDataStoreToolBoostSpecsSpecConditionBoostSpecsList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppVersionSnapshotToolsDataStoreToolBoostSpecsSpecConditionBoostSpecsList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppVersionSnapshotToolsDataStoreToolBoostSpecsSpecConditionBoostSpecsList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesAppVersionSnapshotToolsDataStoreToolBoostSpecsSpecConditionBoostSpecsList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}

func NewCesAppVersionSnapshotToolsDataStoreToolBoostSpecsSpecConditionBoostSpecsList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) CesAppVersionSnapshotToolsDataStoreToolBoostSpecsSpecConditionBoostSpecsList {
	_init_.Initialize()

	if err := validateNewCesAppVersionSnapshotToolsDataStoreToolBoostSpecsSpecConditionBoostSpecsListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_CesAppVersionSnapshotToolsDataStoreToolBoostSpecsSpecConditionBoostSpecsList{}

	_jsii_.Create(
		"@cdktf/provider-google.cesAppVersion.CesAppVersionSnapshotToolsDataStoreToolBoostSpecsSpecConditionBoostSpecsList",
		[]any{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewCesAppVersionSnapshotToolsDataStoreToolBoostSpecsSpecConditionBoostSpecsList_Override(c CesAppVersionSnapshotToolsDataStoreToolBoostSpecsSpecConditionBoostSpecsList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.cesAppVersion.CesAppVersionSnapshotToolsDataStoreToolBoostSpecsSpecConditionBoostSpecsList",
		[]any{terraformResource, terraformAttribute, wrapsSet},
		c,
	)
}

func (j *jsiiProxy_CesAppVersionSnapshotToolsDataStoreToolBoostSpecsSpecConditionBoostSpecsList) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CesAppVersionSnapshotToolsDataStoreToolBoostSpecsSpecConditionBoostSpecsList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CesAppVersionSnapshotToolsDataStoreToolBoostSpecsSpecConditionBoostSpecsList) SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (c *jsiiProxy_CesAppVersionSnapshotToolsDataStoreToolBoostSpecsSpecConditionBoostSpecsList) AllWithMapKey(mapKeyAttributeName *string) cdktf.DynamicListTerraformIterator {
	if err := c.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktf.DynamicListTerraformIterator

	_jsii_.Invoke(
		c,
		"allWithMapKey",
		[]any{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAppVersionSnapshotToolsDataStoreToolBoostSpecsSpecConditionBoostSpecsList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAppVersionSnapshotToolsDataStoreToolBoostSpecsSpecConditionBoostSpecsList) Get(index *float64) CesAppVersionSnapshotToolsDataStoreToolBoostSpecsSpecConditionBoostSpecsOutputReference {
	if err := c.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns CesAppVersionSnapshotToolsDataStoreToolBoostSpecsSpecConditionBoostSpecsOutputReference

	_jsii_.Invoke(
		c,
		"get",
		[]any{index},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesAppVersionSnapshotToolsDataStoreToolBoostSpecsSpecConditionBoostSpecsList) Resolve(_context cdktf.IResolveContext) any {
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

func (c *jsiiProxy_CesAppVersionSnapshotToolsDataStoreToolBoostSpecsSpecConditionBoostSpecsList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
