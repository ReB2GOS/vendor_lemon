package android
type Product_variables struct {
	Additional_gralloc_10_usage_bits struct {
		Cppflags []string
	}
	Has_legacy_camera_hal1 struct {
		Cflags []string
	}
	Needs_text_relocations struct {
		Cppflags []string
	}
	Target_init_vendor_lib struct {
		Whole_static_libs []string
	}
	Target_process_sdk_version_override struct {
		Cppflags []string
	}
	Target_shim_libs struct {
		Cppflags []string
	}
}

type ProductVariables struct {
	Additional_gralloc_10_usage_bits  *string `json:",omitempty"`
	Has_legacy_camera_hal1  *bool `json:",omitempty"`
	Needs_text_relocations  *bool `json:",omitempty"`
	Target_init_vendor_lib  *string `json:",omitempty"`
	Target_process_sdk_version_override *string `json:",omitempty"`
	Target_shim_libs  *string `json:",omitempty"`
}
