// Code generated by internal/generate/tags/main.go; DO NOT EDIT.

package macie2

import (
	"github.com/aws/aws-sdk-go/aws"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
)

// map[string]*string handling

// Tags returns macie2 service tags.
func Tags(tags tftags.KeyValueTags) map[string]*string {
	return aws.StringMap(tags.Map())
}

// KeyValueTags creates KeyValueTags from macie2 service tags.
func KeyValueTags(tags map[string]*string) tftags.KeyValueTags {
	return tftags.New(tags)
}
