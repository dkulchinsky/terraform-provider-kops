package schemas

import (
	"reflect"

	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func ResourceContainerdConfig() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"address":          OptionalString(),
			"config_override":  OptionalString(),
			"log_level":        OptionalString(),
			"packages":         OptionalStruct(ResourcePackagesConfig()),
			"registry_mirrors": OptionalMap(List(String())),
			"root":             OptionalString(),
			"skip_install":     OptionalBool(),
			"state":            OptionalString(),
			"version":          OptionalString(),
		},
	}

	return res
}

func ExpandResourceContainerdConfig(in map[string]interface{}) kops.ContainerdConfig {
	if in == nil {
		panic("expand ContainerdConfig failure, in is nil")
	}
	return kops.ContainerdConfig{
		Address: func(in interface{}) *string {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in string) *string {
					return &in
				}(string(ExpandString(in)))
			}(in)
		}(in["address"]),
		ConfigOverride: func(in interface{}) *string {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in string) *string {
					return &in
				}(string(ExpandString(in)))
			}(in)
		}(in["config_override"]),
		LogLevel: func(in interface{}) *string {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in string) *string {
					return &in
				}(string(ExpandString(in)))
			}(in)
		}(in["log_level"]),
		Packages: func(in interface{}) *kops.PackagesConfig {
			return func(in interface{}) *kops.PackagesConfig {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.PackagesConfig) *kops.PackagesConfig {
					return &in
				}(func(in interface{}) kops.PackagesConfig {
					if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
						return kops.PackagesConfig{}
					}
					return (ExpandResourcePackagesConfig(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["packages"]),
		RegistryMirrors: func(in interface{}) map[string][]string {
			return func(in interface{}) map[string][]string {
				if in == nil {
					return nil
				}
				if in, ok := in.(map[string]interface{}); ok {
					if len(in) > 0 {
						out := map[string][]string{}
						for key, in := range in {
							out[key] = func(in interface{}) []string {
								if in == nil {
									return nil
								}
								var out []string
								for _, in := range in.([]interface{}) {
									out = append(out, string(ExpandString(in)))
								}
								return out
							}(in)
						}
						return out
					}
				}
				return nil
			}(in)
		}(in["registry_mirrors"]),
		Root: func(in interface{}) *string {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in string) *string {
					return &in
				}(string(ExpandString(in)))
			}(in)
		}(in["root"]),
		SkipInstall: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["skip_install"]),
		State: func(in interface{}) *string {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in string) *string {
					return &in
				}(string(ExpandString(in)))
			}(in)
		}(in["state"]),
		Version: func(in interface{}) *string {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in string) *string {
					return &in
				}(string(ExpandString(in)))
			}(in)
		}(in["version"]),
	}
}

func FlattenResourceContainerdConfigInto(in kops.ContainerdConfig, out map[string]interface{}) {
	out["address"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.Address)
	out["config_override"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.ConfigOverride)
	out["log_level"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.LogLevel)
	out["packages"] = func(in *kops.PackagesConfig) interface{} {
		return func(in *kops.PackagesConfig) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.PackagesConfig) interface{} {
				return func(in kops.PackagesConfig) []interface{} {
					return []interface{}{FlattenResourcePackagesConfig(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Packages)
	out["registry_mirrors"] = func(in map[string][]string) interface{} {
		return func(in map[string][]string) map[string]interface{} {
			if in == nil {
				return nil
			}
			out := map[string]interface{}{}
			for key, in := range in {
				out[key] = func(in []string) []interface{} {
					var out []interface{}
					for _, in := range in {
						out = append(out, FlattenString(string(in)))
					}
					return out
				}(in)
			}
			return out
		}(in)
	}(in.RegistryMirrors)
	out["root"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.Root)
	out["skip_install"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.SkipInstall)
	out["state"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.State)
	out["version"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.Version)
}

func FlattenResourceContainerdConfig(in kops.ContainerdConfig) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceContainerdConfigInto(in, out)
	return out
}
