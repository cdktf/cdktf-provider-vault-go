// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package configuicustommessage


type ConfigUiCustomMessageLink struct {
	// The URL of the hyperlink.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.2.1/docs/resources/config_ui_custom_message#href ConfigUiCustomMessage#href}
	Href *string `field:"required" json:"href" yaml:"href"`
	// The title of the hyperlink.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.2.1/docs/resources/config_ui_custom_message#title ConfigUiCustomMessage#title}
	Title *string `field:"required" json:"title" yaml:"title"`
}

