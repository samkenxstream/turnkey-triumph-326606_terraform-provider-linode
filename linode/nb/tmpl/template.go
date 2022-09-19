package tmpl

import (
	"testing"

	"github.com/linode/terraform-provider-linode/linode/acceptance"
)

type TemplateData struct {
	Label  string
	Region string
}

func Basic(t *testing.T, nodebalancer, region string) string {
	return acceptance.ExecuteTemplate(t,
		"nodebalancer_basic", TemplateData{
			Label:  nodebalancer,
			Region: region,
		})
}

func Updates(t *testing.T, nodebalancer, region string) string {
	return acceptance.ExecuteTemplate(t,
		"nodebalancer_updates", TemplateData{
			Label:  nodebalancer,
			Region: region,
		})
}

func DataBasic(t *testing.T, nodebalancer, region string) string {
	return acceptance.ExecuteTemplate(t,
		"nodebalancer_data_basic",
		TemplateData{
			Label:  nodebalancer,
			Region: region,
		})
}
