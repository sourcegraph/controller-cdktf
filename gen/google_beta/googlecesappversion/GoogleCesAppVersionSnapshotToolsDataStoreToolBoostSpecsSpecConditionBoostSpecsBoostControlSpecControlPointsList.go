package googlecesappversion

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecesappversion/internal"
)

type GoogleCesAppVersionSnapshotToolsDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecControlPointsList interface {
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
	Get(index *float64) GoogleCesAppVersionSnapshotToolsDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecControlPointsOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleCesAppVersionSnapshotToolsDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecControlPointsList
type jsiiProxy_GoogleCesAppVersionSnapshotToolsDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecControlPointsList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_GoogleCesAppVersionSnapshotToolsDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecControlPointsList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppVersionSnapshotToolsDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecControlPointsList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppVersionSnapshotToolsDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecControlPointsList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppVersionSnapshotToolsDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecControlPointsList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppVersionSnapshotToolsDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecControlPointsList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewGoogleCesAppVersionSnapshotToolsDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecControlPointsList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) GoogleCesAppVersionSnapshotToolsDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecControlPointsList {
	_init_.Initialize()

	if err := validateNewGoogleCesAppVersionSnapshotToolsDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecControlPointsListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleCesAppVersionSnapshotToolsDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecControlPointsList{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleCesAppVersion.GoogleCesAppVersionSnapshotToolsDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecControlPointsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewGoogleCesAppVersionSnapshotToolsDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecControlPointsList_Override(g GoogleCesAppVersionSnapshotToolsDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecControlPointsList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleCesAppVersion.GoogleCesAppVersionSnapshotToolsDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecControlPointsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		g,
	)
}

func (j *jsiiProxy_GoogleCesAppVersionSnapshotToolsDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecControlPointsList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAppVersionSnapshotToolsDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecControlPointsList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAppVersionSnapshotToolsDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecControlPointsList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (g *jsiiProxy_GoogleCesAppVersionSnapshotToolsDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecControlPointsList) AllWithMapKey(mapKeyAttributeName *string) cdktf.DynamicListTerraformIterator {
	if err := g.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktf.DynamicListTerraformIterator

	_jsii_.Invoke(
		g,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesAppVersionSnapshotToolsDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecControlPointsList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesAppVersionSnapshotToolsDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecControlPointsList) Get(index *float64) GoogleCesAppVersionSnapshotToolsDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecControlPointsOutputReference {
	if err := g.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns GoogleCesAppVersionSnapshotToolsDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecControlPointsOutputReference

	_jsii_.Invoke(
		g,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesAppVersionSnapshotToolsDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecControlPointsList) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleCesAppVersionSnapshotToolsDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecControlPointsList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

