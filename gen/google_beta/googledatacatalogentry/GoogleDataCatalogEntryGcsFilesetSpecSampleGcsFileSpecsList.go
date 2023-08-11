package googledatacatalogentry

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledatacatalogentry/internal"
)

type GoogleDataCatalogEntryGcsFilesetSpecSampleGcsFileSpecsList interface {
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
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) GoogleDataCatalogEntryGcsFilesetSpecSampleGcsFileSpecsOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDataCatalogEntryGcsFilesetSpecSampleGcsFileSpecsList
type jsiiProxy_GoogleDataCatalogEntryGcsFilesetSpecSampleGcsFileSpecsList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_GoogleDataCatalogEntryGcsFilesetSpecSampleGcsFileSpecsList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataCatalogEntryGcsFilesetSpecSampleGcsFileSpecsList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataCatalogEntryGcsFilesetSpecSampleGcsFileSpecsList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataCatalogEntryGcsFilesetSpecSampleGcsFileSpecsList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataCatalogEntryGcsFilesetSpecSampleGcsFileSpecsList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewGoogleDataCatalogEntryGcsFilesetSpecSampleGcsFileSpecsList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) GoogleDataCatalogEntryGcsFilesetSpecSampleGcsFileSpecsList {
	_init_.Initialize()

	if err := validateNewGoogleDataCatalogEntryGcsFilesetSpecSampleGcsFileSpecsListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDataCatalogEntryGcsFilesetSpecSampleGcsFileSpecsList{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDataCatalogEntry.GoogleDataCatalogEntryGcsFilesetSpecSampleGcsFileSpecsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewGoogleDataCatalogEntryGcsFilesetSpecSampleGcsFileSpecsList_Override(g GoogleDataCatalogEntryGcsFilesetSpecSampleGcsFileSpecsList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDataCatalogEntry.GoogleDataCatalogEntryGcsFilesetSpecSampleGcsFileSpecsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		g,
	)
}

func (j *jsiiProxy_GoogleDataCatalogEntryGcsFilesetSpecSampleGcsFileSpecsList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDataCatalogEntryGcsFilesetSpecSampleGcsFileSpecsList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleDataCatalogEntryGcsFilesetSpecSampleGcsFileSpecsList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (g *jsiiProxy_GoogleDataCatalogEntryGcsFilesetSpecSampleGcsFileSpecsList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataCatalogEntryGcsFilesetSpecSampleGcsFileSpecsList) Get(index *float64) GoogleDataCatalogEntryGcsFilesetSpecSampleGcsFileSpecsOutputReference {
	if err := g.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns GoogleDataCatalogEntryGcsFilesetSpecSampleGcsFileSpecsOutputReference

	_jsii_.Invoke(
		g,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataCatalogEntryGcsFilesetSpecSampleGcsFileSpecsList) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDataCatalogEntryGcsFilesetSpecSampleGcsFileSpecsList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

