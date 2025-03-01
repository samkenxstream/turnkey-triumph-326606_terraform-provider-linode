package tmpl

import (
	"testing"

	"github.com/linode/terraform-provider-linode/linode/acceptance"
)

type InstanceTemplateData struct {
	Prefix string
	ID     string
	PubKey string
	Region string
}

type TemplateData struct {
	Instances []InstanceTemplateData

	Label string
}

func Basic(t *testing.T, label, devicePrefix, region string) string {
	return acceptance.ExecuteTemplate(t,
		"firewall_basic", TemplateData{
			Label: label,
			Instances: []InstanceTemplateData{
				{
					Prefix: devicePrefix,
					ID:     "one",
					PubKey: acceptance.PublicKeyMaterial,
					Region: region,
				},
			},
		})
}

func Updates(t *testing.T, label, devicePrefix, region string) string {
	return acceptance.ExecuteTemplate(t,
		"firewall_updates", TemplateData{
			Label: label,
			Instances: []InstanceTemplateData{
				{
					Prefix: devicePrefix,
					ID:     "one",
					PubKey: acceptance.PublicKeyMaterial,
					Region: region,
				},
				{
					Prefix: devicePrefix,
					ID:     "two",
					PubKey: acceptance.PublicKeyMaterial,
					Region: region,
				},
			},
		})
}

func Minimum(t *testing.T, label string) string {
	return acceptance.ExecuteTemplate(t,
		"firewall_minimum", TemplateData{
			Label: label,
		})
}

func MultipleRules(t *testing.T, label, devicePrefix, region string) string {
	return acceptance.ExecuteTemplate(t,
		"firewall_multiple_rules", TemplateData{
			Label: label,
			Instances: []InstanceTemplateData{
				{
					Prefix: devicePrefix,
					ID:     "one",
					PubKey: acceptance.PublicKeyMaterial,
					Region: region,
				},
			},
		})
}

func NoDevice(t *testing.T, label string) string {
	return acceptance.ExecuteTemplate(t,
		"firewall_no_device", TemplateData{
			Label: label,
		})
}

func NoIPv6(t *testing.T, label string) string {
	return acceptance.ExecuteTemplate(t,
		"firewall_no_ipv6", TemplateData{
			Label: label,
		})
}

func DataBasic(t *testing.T, label, devicePrefix, region string) string {
	return acceptance.ExecuteTemplate(t,
		"firewall_data_basic", TemplateData{
			Label: label,
			Instances: []InstanceTemplateData{
				{
					Prefix: devicePrefix,
					ID:     "one",
					PubKey: acceptance.PublicKeyMaterial,
					Region: region,
				},
			},
		})
}
