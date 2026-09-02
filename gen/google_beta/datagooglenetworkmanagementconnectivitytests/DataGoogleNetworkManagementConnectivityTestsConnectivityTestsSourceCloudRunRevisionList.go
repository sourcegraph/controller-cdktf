package datagooglenetworkmanagementconnectivitytests

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/datagooglenetworkmanagementconnectivitytests/internal"
)

type DataGoogleNetworkManagementConnectivityTestsConnectivityTestsSourceCloudRunRevisionList interface {
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
	Get(index *float64) DataGoogleNetworkManagementConnectivityTestsConnectivityTestsSourceCloudRunRevisionOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataGoogleNetworkManagementConnectivityTestsConnectivityTestsSourceCloudRunRevisionList
type jsiiProxy_DataGoogleNetworkManagementConnectivityTestsConnectivityTestsSourceCloudRunRevisionList struct {
	internal.Type__cdktnComplexList
}

func (j *jsiiProxy_DataGoogleNetworkManagementConnectivityTestsConnectivityTestsSourceCloudRunRevisionList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleNetworkManagementConnectivityTestsConnectivityTestsSourceCloudRunRevisionList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleNetworkManagementConnectivityTestsConnectivityTestsSourceCloudRunRevisionList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleNetworkManagementConnectivityTestsConnectivityTestsSourceCloudRunRevisionList) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleNetworkManagementConnectivityTestsConnectivityTestsSourceCloudRunRevisionList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewDataGoogleNetworkManagementConnectivityTestsConnectivityTestsSourceCloudRunRevisionList(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) DataGoogleNetworkManagementConnectivityTestsConnectivityTestsSourceCloudRunRevisionList {
	_init_.Initialize()

	if err := validateNewDataGoogleNetworkManagementConnectivityTestsConnectivityTestsSourceCloudRunRevisionListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataGoogleNetworkManagementConnectivityTestsConnectivityTestsSourceCloudRunRevisionList{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.dataGoogleNetworkManagementConnectivityTests.DataGoogleNetworkManagementConnectivityTestsConnectivityTestsSourceCloudRunRevisionList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewDataGoogleNetworkManagementConnectivityTestsConnectivityTestsSourceCloudRunRevisionList_Override(d DataGoogleNetworkManagementConnectivityTestsConnectivityTestsSourceCloudRunRevisionList, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.dataGoogleNetworkManagementConnectivityTests.DataGoogleNetworkManagementConnectivityTestsConnectivityTestsSourceCloudRunRevisionList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DataGoogleNetworkManagementConnectivityTestsConnectivityTestsSourceCloudRunRevisionList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataGoogleNetworkManagementConnectivityTestsConnectivityTestsSourceCloudRunRevisionList)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataGoogleNetworkManagementConnectivityTestsConnectivityTestsSourceCloudRunRevisionList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (d *jsiiProxy_DataGoogleNetworkManagementConnectivityTestsConnectivityTestsSourceCloudRunRevisionList) AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator {
	if err := d.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktn.DynamicListTerraformIterator

	_jsii_.Invoke(
		d,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleNetworkManagementConnectivityTestsConnectivityTestsSourceCloudRunRevisionList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleNetworkManagementConnectivityTestsConnectivityTestsSourceCloudRunRevisionList) Get(index *float64) DataGoogleNetworkManagementConnectivityTestsConnectivityTestsSourceCloudRunRevisionOutputReference {
	if err := d.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns DataGoogleNetworkManagementConnectivityTestsConnectivityTestsSourceCloudRunRevisionOutputReference

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleNetworkManagementConnectivityTestsConnectivityTestsSourceCloudRunRevisionList) Resolve(context cdktn.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleNetworkManagementConnectivityTestsConnectivityTestsSourceCloudRunRevisionList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

