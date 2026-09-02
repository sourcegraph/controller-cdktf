package googlecesappversion

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecesappversion/internal"
)

type GoogleCesAppVersionSnapshotToolsDataStoreToolEngineSourceDataStoreSourcesDataStoreList interface {
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
	Get(index *float64) GoogleCesAppVersionSnapshotToolsDataStoreToolEngineSourceDataStoreSourcesDataStoreOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleCesAppVersionSnapshotToolsDataStoreToolEngineSourceDataStoreSourcesDataStoreList
type jsiiProxy_GoogleCesAppVersionSnapshotToolsDataStoreToolEngineSourceDataStoreSourcesDataStoreList struct {
	internal.Type__cdktnComplexList
}

func (j *jsiiProxy_GoogleCesAppVersionSnapshotToolsDataStoreToolEngineSourceDataStoreSourcesDataStoreList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppVersionSnapshotToolsDataStoreToolEngineSourceDataStoreSourcesDataStoreList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppVersionSnapshotToolsDataStoreToolEngineSourceDataStoreSourcesDataStoreList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppVersionSnapshotToolsDataStoreToolEngineSourceDataStoreSourcesDataStoreList) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppVersionSnapshotToolsDataStoreToolEngineSourceDataStoreSourcesDataStoreList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewGoogleCesAppVersionSnapshotToolsDataStoreToolEngineSourceDataStoreSourcesDataStoreList(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) GoogleCesAppVersionSnapshotToolsDataStoreToolEngineSourceDataStoreSourcesDataStoreList {
	_init_.Initialize()

	if err := validateNewGoogleCesAppVersionSnapshotToolsDataStoreToolEngineSourceDataStoreSourcesDataStoreListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleCesAppVersionSnapshotToolsDataStoreToolEngineSourceDataStoreSourcesDataStoreList{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleCesAppVersion.GoogleCesAppVersionSnapshotToolsDataStoreToolEngineSourceDataStoreSourcesDataStoreList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewGoogleCesAppVersionSnapshotToolsDataStoreToolEngineSourceDataStoreSourcesDataStoreList_Override(g GoogleCesAppVersionSnapshotToolsDataStoreToolEngineSourceDataStoreSourcesDataStoreList, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleCesAppVersion.GoogleCesAppVersionSnapshotToolsDataStoreToolEngineSourceDataStoreSourcesDataStoreList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		g,
	)
}

func (j *jsiiProxy_GoogleCesAppVersionSnapshotToolsDataStoreToolEngineSourceDataStoreSourcesDataStoreList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAppVersionSnapshotToolsDataStoreToolEngineSourceDataStoreSourcesDataStoreList)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAppVersionSnapshotToolsDataStoreToolEngineSourceDataStoreSourcesDataStoreList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (g *jsiiProxy_GoogleCesAppVersionSnapshotToolsDataStoreToolEngineSourceDataStoreSourcesDataStoreList) AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator {
	if err := g.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktn.DynamicListTerraformIterator

	_jsii_.Invoke(
		g,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesAppVersionSnapshotToolsDataStoreToolEngineSourceDataStoreSourcesDataStoreList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesAppVersionSnapshotToolsDataStoreToolEngineSourceDataStoreSourcesDataStoreList) Get(index *float64) GoogleCesAppVersionSnapshotToolsDataStoreToolEngineSourceDataStoreSourcesDataStoreOutputReference {
	if err := g.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns GoogleCesAppVersionSnapshotToolsDataStoreToolEngineSourceDataStoreSourcesDataStoreOutputReference

	_jsii_.Invoke(
		g,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesAppVersionSnapshotToolsDataStoreToolEngineSourceDataStoreSourcesDataStoreList) Resolve(context cdktn.IResolveContext) interface{} {
	if err := g.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesAppVersionSnapshotToolsDataStoreToolEngineSourceDataStoreSourcesDataStoreList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

