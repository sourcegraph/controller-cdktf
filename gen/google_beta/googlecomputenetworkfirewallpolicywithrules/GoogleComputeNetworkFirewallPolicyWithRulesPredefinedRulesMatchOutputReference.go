package googlecomputenetworkfirewallpolicywithrules

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecomputenetworkfirewallpolicywithrules/internal"
)

type GoogleComputeNetworkFirewallPolicyWithRulesPredefinedRulesMatchOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
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
	DestAddressGroups() *[]*string
	DestFqdns() *[]*string
	DestIpRanges() *[]*string
	DestRegionCodes() *[]*string
	DestThreatIntelligences() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleComputeNetworkFirewallPolicyWithRulesPredefinedRulesMatch
	SetInternalValue(val *GoogleComputeNetworkFirewallPolicyWithRulesPredefinedRulesMatch)
	Layer4Config() GoogleComputeNetworkFirewallPolicyWithRulesPredefinedRulesMatchLayer4ConfigList
	SrcAddressGroups() *[]*string
	SrcFqdns() *[]*string
	SrcIpRanges() *[]*string
	SrcRegionCodes() *[]*string
	SrcSecureTag() GoogleComputeNetworkFirewallPolicyWithRulesPredefinedRulesMatchSrcSecureTagList
	SrcThreatIntelligences() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleComputeNetworkFirewallPolicyWithRulesPredefinedRulesMatchOutputReference
type jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesPredefinedRulesMatchOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesPredefinedRulesMatchOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesPredefinedRulesMatchOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesPredefinedRulesMatchOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesPredefinedRulesMatchOutputReference) DestAddressGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destAddressGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesPredefinedRulesMatchOutputReference) DestFqdns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destFqdns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesPredefinedRulesMatchOutputReference) DestIpRanges() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destIpRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesPredefinedRulesMatchOutputReference) DestRegionCodes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destRegionCodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesPredefinedRulesMatchOutputReference) DestThreatIntelligences() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destThreatIntelligences",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesPredefinedRulesMatchOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesPredefinedRulesMatchOutputReference) InternalValue() *GoogleComputeNetworkFirewallPolicyWithRulesPredefinedRulesMatch {
	var returns *GoogleComputeNetworkFirewallPolicyWithRulesPredefinedRulesMatch
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesPredefinedRulesMatchOutputReference) Layer4Config() GoogleComputeNetworkFirewallPolicyWithRulesPredefinedRulesMatchLayer4ConfigList {
	var returns GoogleComputeNetworkFirewallPolicyWithRulesPredefinedRulesMatchLayer4ConfigList
	_jsii_.Get(
		j,
		"layer4Config",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesPredefinedRulesMatchOutputReference) SrcAddressGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"srcAddressGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesPredefinedRulesMatchOutputReference) SrcFqdns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"srcFqdns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesPredefinedRulesMatchOutputReference) SrcIpRanges() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"srcIpRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesPredefinedRulesMatchOutputReference) SrcRegionCodes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"srcRegionCodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesPredefinedRulesMatchOutputReference) SrcSecureTag() GoogleComputeNetworkFirewallPolicyWithRulesPredefinedRulesMatchSrcSecureTagList {
	var returns GoogleComputeNetworkFirewallPolicyWithRulesPredefinedRulesMatchSrcSecureTagList
	_jsii_.Get(
		j,
		"srcSecureTag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesPredefinedRulesMatchOutputReference) SrcThreatIntelligences() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"srcThreatIntelligences",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesPredefinedRulesMatchOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesPredefinedRulesMatchOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleComputeNetworkFirewallPolicyWithRulesPredefinedRulesMatchOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleComputeNetworkFirewallPolicyWithRulesPredefinedRulesMatchOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleComputeNetworkFirewallPolicyWithRulesPredefinedRulesMatchOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesPredefinedRulesMatchOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleComputeNetworkFirewallPolicyWithRules.GoogleComputeNetworkFirewallPolicyWithRulesPredefinedRulesMatchOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleComputeNetworkFirewallPolicyWithRulesPredefinedRulesMatchOutputReference_Override(g GoogleComputeNetworkFirewallPolicyWithRulesPredefinedRulesMatchOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleComputeNetworkFirewallPolicyWithRules.GoogleComputeNetworkFirewallPolicyWithRulesPredefinedRulesMatchOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesPredefinedRulesMatchOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesPredefinedRulesMatchOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesPredefinedRulesMatchOutputReference)SetInternalValue(val *GoogleComputeNetworkFirewallPolicyWithRulesPredefinedRulesMatch) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesPredefinedRulesMatchOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesPredefinedRulesMatchOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesPredefinedRulesMatchOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesPredefinedRulesMatchOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesPredefinedRulesMatchOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesPredefinedRulesMatchOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesPredefinedRulesMatchOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesPredefinedRulesMatchOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesPredefinedRulesMatchOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesPredefinedRulesMatchOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesPredefinedRulesMatchOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesPredefinedRulesMatchOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesPredefinedRulesMatchOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesPredefinedRulesMatchOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesPredefinedRulesMatchOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleComputeNetworkFirewallPolicyWithRulesPredefinedRulesMatchOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

