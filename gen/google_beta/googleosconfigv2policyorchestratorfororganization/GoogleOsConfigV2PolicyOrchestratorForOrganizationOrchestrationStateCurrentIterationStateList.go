package googleosconfigv2policyorchestratorfororganization

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleosconfigv2policyorchestratorfororganization/internal"
)

type GoogleOsConfigV2PolicyOrchestratorForOrganizationOrchestrationStateCurrentIterationStateList interface {
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
	Get(index *float64) GoogleOsConfigV2PolicyOrchestratorForOrganizationOrchestrationStateCurrentIterationStateOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleOsConfigV2PolicyOrchestratorForOrganizationOrchestrationStateCurrentIterationStateList
type jsiiProxy_GoogleOsConfigV2PolicyOrchestratorForOrganizationOrchestrationStateCurrentIterationStateList struct {
	internal.Type__cdktnComplexList
}

func (j *jsiiProxy_GoogleOsConfigV2PolicyOrchestratorForOrganizationOrchestrationStateCurrentIterationStateList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigV2PolicyOrchestratorForOrganizationOrchestrationStateCurrentIterationStateList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigV2PolicyOrchestratorForOrganizationOrchestrationStateCurrentIterationStateList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigV2PolicyOrchestratorForOrganizationOrchestrationStateCurrentIterationStateList) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigV2PolicyOrchestratorForOrganizationOrchestrationStateCurrentIterationStateList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewGoogleOsConfigV2PolicyOrchestratorForOrganizationOrchestrationStateCurrentIterationStateList(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) GoogleOsConfigV2PolicyOrchestratorForOrganizationOrchestrationStateCurrentIterationStateList {
	_init_.Initialize()

	if err := validateNewGoogleOsConfigV2PolicyOrchestratorForOrganizationOrchestrationStateCurrentIterationStateListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleOsConfigV2PolicyOrchestratorForOrganizationOrchestrationStateCurrentIterationStateList{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleOsConfigV2PolicyOrchestratorForOrganization.GoogleOsConfigV2PolicyOrchestratorForOrganizationOrchestrationStateCurrentIterationStateList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewGoogleOsConfigV2PolicyOrchestratorForOrganizationOrchestrationStateCurrentIterationStateList_Override(g GoogleOsConfigV2PolicyOrchestratorForOrganizationOrchestrationStateCurrentIterationStateList, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleOsConfigV2PolicyOrchestratorForOrganization.GoogleOsConfigV2PolicyOrchestratorForOrganizationOrchestrationStateCurrentIterationStateList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		g,
	)
}

func (j *jsiiProxy_GoogleOsConfigV2PolicyOrchestratorForOrganizationOrchestrationStateCurrentIterationStateList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleOsConfigV2PolicyOrchestratorForOrganizationOrchestrationStateCurrentIterationStateList)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleOsConfigV2PolicyOrchestratorForOrganizationOrchestrationStateCurrentIterationStateList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (g *jsiiProxy_GoogleOsConfigV2PolicyOrchestratorForOrganizationOrchestrationStateCurrentIterationStateList) AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator {
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

func (g *jsiiProxy_GoogleOsConfigV2PolicyOrchestratorForOrganizationOrchestrationStateCurrentIterationStateList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOsConfigV2PolicyOrchestratorForOrganizationOrchestrationStateCurrentIterationStateList) Get(index *float64) GoogleOsConfigV2PolicyOrchestratorForOrganizationOrchestrationStateCurrentIterationStateOutputReference {
	if err := g.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns GoogleOsConfigV2PolicyOrchestratorForOrganizationOrchestrationStateCurrentIterationStateOutputReference

	_jsii_.Invoke(
		g,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOsConfigV2PolicyOrchestratorForOrganizationOrchestrationStateCurrentIterationStateList) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleOsConfigV2PolicyOrchestratorForOrganizationOrchestrationStateCurrentIterationStateList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

