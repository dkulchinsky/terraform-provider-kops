package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandDataSourceOpenstackBlockStorageConfig(t *testing.T) {
	_default := kops.OpenstackBlockStorageConfig{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.OpenstackBlockStorageConfig
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"version":              nil,
					"ignore_az":            nil,
					"override_az":          nil,
					"create_storage_class": nil,
					"csi_plugin_image":     "",
					"csi_topology_support": nil,
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandDataSourceOpenstackBlockStorageConfig(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandDataSourceOpenstackBlockStorageConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceOpenstackBlockStorageConfigInto(t *testing.T) {
	_default := map[string]interface{}{
		"version":              nil,
		"ignore_az":            nil,
		"override_az":          nil,
		"create_storage_class": nil,
		"csi_plugin_image":     "",
		"csi_topology_support": nil,
	}
	type args struct {
		in kops.OpenstackBlockStorageConfig
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.OpenstackBlockStorageConfig{},
			},
			want: _default,
		},
		{
			name: "Version - default",
			args: args{
				in: func() kops.OpenstackBlockStorageConfig {
					subject := kops.OpenstackBlockStorageConfig{}
					subject.Version = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "IgnoreAZ - default",
			args: args{
				in: func() kops.OpenstackBlockStorageConfig {
					subject := kops.OpenstackBlockStorageConfig{}
					subject.IgnoreAZ = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "OverrideAZ - default",
			args: args{
				in: func() kops.OpenstackBlockStorageConfig {
					subject := kops.OpenstackBlockStorageConfig{}
					subject.OverrideAZ = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CreateStorageClass - default",
			args: args{
				in: func() kops.OpenstackBlockStorageConfig {
					subject := kops.OpenstackBlockStorageConfig{}
					subject.CreateStorageClass = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CsiPluginImage - default",
			args: args{
				in: func() kops.OpenstackBlockStorageConfig {
					subject := kops.OpenstackBlockStorageConfig{}
					subject.CSIPluginImage = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CsiTopologySupport - default",
			args: args{
				in: func() kops.OpenstackBlockStorageConfig {
					subject := kops.OpenstackBlockStorageConfig{}
					subject.CSITopologySupport = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenDataSourceOpenstackBlockStorageConfigInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceOpenstackBlockStorageConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenDataSourceOpenstackBlockStorageConfig(t *testing.T) {
	_default := map[string]interface{}{
		"version":              nil,
		"ignore_az":            nil,
		"override_az":          nil,
		"create_storage_class": nil,
		"csi_plugin_image":     "",
		"csi_topology_support": nil,
	}
	type args struct {
		in kops.OpenstackBlockStorageConfig
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.OpenstackBlockStorageConfig{},
			},
			want: _default,
		},
		{
			name: "Version - default",
			args: args{
				in: func() kops.OpenstackBlockStorageConfig {
					subject := kops.OpenstackBlockStorageConfig{}
					subject.Version = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "IgnoreAZ - default",
			args: args{
				in: func() kops.OpenstackBlockStorageConfig {
					subject := kops.OpenstackBlockStorageConfig{}
					subject.IgnoreAZ = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "OverrideAZ - default",
			args: args{
				in: func() kops.OpenstackBlockStorageConfig {
					subject := kops.OpenstackBlockStorageConfig{}
					subject.OverrideAZ = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CreateStorageClass - default",
			args: args{
				in: func() kops.OpenstackBlockStorageConfig {
					subject := kops.OpenstackBlockStorageConfig{}
					subject.CreateStorageClass = nil
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CsiPluginImage - default",
			args: args{
				in: func() kops.OpenstackBlockStorageConfig {
					subject := kops.OpenstackBlockStorageConfig{}
					subject.CSIPluginImage = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "CsiTopologySupport - default",
			args: args{
				in: func() kops.OpenstackBlockStorageConfig {
					subject := kops.OpenstackBlockStorageConfig{}
					subject.CSITopologySupport = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenDataSourceOpenstackBlockStorageConfig(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenDataSourceOpenstackBlockStorageConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
