package tmpl

import (
	"testing"

	"github.com/linode/terraform-provider-linode/linode/acceptance"
)

type TemplateData struct {
	ID string
}

func DataBasic(t *testing.T, id string) string {
	return acceptance.ExecuteTemplate(t,
		"obj_cluster_data_basic", TemplateData{ID: id})
}
