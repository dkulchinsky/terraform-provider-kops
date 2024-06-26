package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	. "github.com/terraform-kops/terraform-provider-kops/pkg/schemas"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func ResourceBastionLoadBalancerSpec() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"type": RequiredString(),
		},
	}

	return res
}

func ExpandResourceBastionLoadBalancerSpec(in map[string]interface{}) kops.BastionLoadBalancerSpec {
	if in == nil {
		panic("expand BastionLoadBalancerSpec failure, in is nil")
	}
	return kops.BastionLoadBalancerSpec{
		Type: func(in interface{}) kops.LoadBalancerType {
			return kops.LoadBalancerType(ExpandString(in))
		}(in["type"]),
	}
}

func FlattenResourceBastionLoadBalancerSpecInto(in kops.BastionLoadBalancerSpec, out map[string]interface{}) {
	out["type"] = func(in kops.LoadBalancerType) interface{} {
		return FlattenString(string(in))
	}(in.Type)
}

func FlattenResourceBastionLoadBalancerSpec(in kops.BastionLoadBalancerSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceBastionLoadBalancerSpecInto(in, out)
	return out
}
