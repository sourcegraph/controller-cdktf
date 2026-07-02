package gkehubfeature

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/gkehubfeature/internal"
)

type GkeHubFeatureSpecOutputReference interface {
	cdktf.ComplexObject
	Clusterupgrade() GkeHubFeatureSpecClusterupgradeOutputReference
	ClusterupgradeInput() *GkeHubFeatureSpecClusterupgrade
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() any
	// Experimental.
	SetComplexObjectIndex(val any)
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
	Fleetobservability() GkeHubFeatureSpecFleetobservabilityOutputReference
	FleetobservabilityInput() *GkeHubFeatureSpecFleetobservability
	// Experimental.
	Fqn() *string
	InternalValue() *GkeHubFeatureSpec
	SetInternalValue(val *GkeHubFeatureSpec)
	Multiclusteringress() GkeHubFeatureSpecMulticlusteringressOutputReference
	MulticlusteringressInput() *GkeHubFeatureSpecMulticlusteringress
	Rbacrolebindingactuation() GkeHubFeatureSpecRbacrolebindingactuationOutputReference
	RbacrolebindingactuationInput() *GkeHubFeatureSpecRbacrolebindingactuation
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Workloadidentity() GkeHubFeatureSpecWorkloadidentityOutputReference
	WorkloadidentityInput() *GkeHubFeatureSpecWorkloadidentity
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]any
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
	PutClusterupgrade(value *GkeHubFeatureSpecClusterupgrade)
	PutFleetobservability(value *GkeHubFeatureSpecFleetobservability)
	PutMulticlusteringress(value *GkeHubFeatureSpecMulticlusteringress)
	PutRbacrolebindingactuation(value *GkeHubFeatureSpecRbacrolebindingactuation)
	PutWorkloadidentity(value *GkeHubFeatureSpecWorkloadidentity)
	ResetClusterupgrade()
	ResetFleetobservability()
	ResetMulticlusteringress()
	ResetRbacrolebindingactuation()
	ResetWorkloadidentity()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GkeHubFeatureSpecOutputReference
type jsiiProxy_GkeHubFeatureSpecOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GkeHubFeatureSpecOutputReference) Clusterupgrade() GkeHubFeatureSpecClusterupgradeOutputReference {
	var returns GkeHubFeatureSpecClusterupgradeOutputReference
	_jsii_.Get(
		j,
		"clusterupgrade",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureSpecOutputReference) ClusterupgradeInput() *GkeHubFeatureSpecClusterupgrade {
	var returns *GkeHubFeatureSpecClusterupgrade
	_jsii_.Get(
		j,
		"clusterupgradeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureSpecOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureSpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureSpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureSpecOutputReference) Fleetobservability() GkeHubFeatureSpecFleetobservabilityOutputReference {
	var returns GkeHubFeatureSpecFleetobservabilityOutputReference
	_jsii_.Get(
		j,
		"fleetobservability",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureSpecOutputReference) FleetobservabilityInput() *GkeHubFeatureSpecFleetobservability {
	var returns *GkeHubFeatureSpecFleetobservability
	_jsii_.Get(
		j,
		"fleetobservabilityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureSpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureSpecOutputReference) InternalValue() *GkeHubFeatureSpec {
	var returns *GkeHubFeatureSpec
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureSpecOutputReference) Multiclusteringress() GkeHubFeatureSpecMulticlusteringressOutputReference {
	var returns GkeHubFeatureSpecMulticlusteringressOutputReference
	_jsii_.Get(
		j,
		"multiclusteringress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureSpecOutputReference) MulticlusteringressInput() *GkeHubFeatureSpecMulticlusteringress {
	var returns *GkeHubFeatureSpecMulticlusteringress
	_jsii_.Get(
		j,
		"multiclusteringressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureSpecOutputReference) Rbacrolebindingactuation() GkeHubFeatureSpecRbacrolebindingactuationOutputReference {
	var returns GkeHubFeatureSpecRbacrolebindingactuationOutputReference
	_jsii_.Get(
		j,
		"rbacrolebindingactuation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureSpecOutputReference) RbacrolebindingactuationInput() *GkeHubFeatureSpecRbacrolebindingactuation {
	var returns *GkeHubFeatureSpecRbacrolebindingactuation
	_jsii_.Get(
		j,
		"rbacrolebindingactuationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureSpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureSpecOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureSpecOutputReference) Workloadidentity() GkeHubFeatureSpecWorkloadidentityOutputReference {
	var returns GkeHubFeatureSpecWorkloadidentityOutputReference
	_jsii_.Get(
		j,
		"workloadidentity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureSpecOutputReference) WorkloadidentityInput() *GkeHubFeatureSpecWorkloadidentity {
	var returns *GkeHubFeatureSpecWorkloadidentity
	_jsii_.Get(
		j,
		"workloadidentityInput",
		&returns,
	)
	return returns
}

func NewGkeHubFeatureSpecOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GkeHubFeatureSpecOutputReference {
	_init_.Initialize()

	if err := validateNewGkeHubFeatureSpecOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GkeHubFeatureSpecOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.gkeHubFeature.GkeHubFeatureSpecOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGkeHubFeatureSpecOutputReference_Override(g GkeHubFeatureSpecOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.gkeHubFeature.GkeHubFeatureSpecOutputReference",
		[]any{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GkeHubFeatureSpecOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GkeHubFeatureSpecOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GkeHubFeatureSpecOutputReference) SetInternalValue(val *GkeHubFeatureSpec) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GkeHubFeatureSpecOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GkeHubFeatureSpecOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GkeHubFeatureSpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeHubFeatureSpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeHubFeatureSpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeHubFeatureSpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeHubFeatureSpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeHubFeatureSpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeHubFeatureSpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeHubFeatureSpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeHubFeatureSpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeHubFeatureSpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeHubFeatureSpecOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeHubFeatureSpecOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeHubFeatureSpecOutputReference) PutClusterupgrade(value *GkeHubFeatureSpecClusterupgrade) {
	if err := g.validatePutClusterupgradeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putClusterupgrade",
		[]any{value},
	)
}

func (g *jsiiProxy_GkeHubFeatureSpecOutputReference) PutFleetobservability(value *GkeHubFeatureSpecFleetobservability) {
	if err := g.validatePutFleetobservabilityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putFleetobservability",
		[]any{value},
	)
}

func (g *jsiiProxy_GkeHubFeatureSpecOutputReference) PutMulticlusteringress(value *GkeHubFeatureSpecMulticlusteringress) {
	if err := g.validatePutMulticlusteringressParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMulticlusteringress",
		[]any{value},
	)
}

func (g *jsiiProxy_GkeHubFeatureSpecOutputReference) PutRbacrolebindingactuation(value *GkeHubFeatureSpecRbacrolebindingactuation) {
	if err := g.validatePutRbacrolebindingactuationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRbacrolebindingactuation",
		[]any{value},
	)
}

func (g *jsiiProxy_GkeHubFeatureSpecOutputReference) PutWorkloadidentity(value *GkeHubFeatureSpecWorkloadidentity) {
	if err := g.validatePutWorkloadidentityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putWorkloadidentity",
		[]any{value},
	)
}

func (g *jsiiProxy_GkeHubFeatureSpecOutputReference) ResetClusterupgrade() {
	_jsii_.InvokeVoid(
		g,
		"resetClusterupgrade",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeHubFeatureSpecOutputReference) ResetFleetobservability() {
	_jsii_.InvokeVoid(
		g,
		"resetFleetobservability",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeHubFeatureSpecOutputReference) ResetMulticlusteringress() {
	_jsii_.InvokeVoid(
		g,
		"resetMulticlusteringress",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeHubFeatureSpecOutputReference) ResetRbacrolebindingactuation() {
	_jsii_.InvokeVoid(
		g,
		"resetRbacrolebindingactuation",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeHubFeatureSpecOutputReference) ResetWorkloadidentity() {
	_jsii_.InvokeVoid(
		g,
		"resetWorkloadidentity",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeHubFeatureSpecOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := g.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		g,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeHubFeatureSpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
