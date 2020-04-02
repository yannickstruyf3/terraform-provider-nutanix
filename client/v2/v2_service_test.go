package v2

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"

	"github.com/terraform-providers/terraform-provider-nutanix/client"
	"github.com/terraform-providers/terraform-provider-nutanix/utils"
)

func setup() (*http.ServeMux, *client.Client, *httptest.Server) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	c, _ := client.NewClient(&client.Credentials{
		URL:      "",
		Username: "username",
		Password: "password",
		Port:     "",
		Endpoint: "",
		Insecure: true},
		&client.ApiMetadata{
			LibraryVersion: "v2.0",
			DefaultBaseURL: "https://%s/",
			AbsolutePath:   "/PrismGateway/services/rest/v2.0",
			UserAgent:      "nutanix/v2.0",
			MediaType:      "application/json"})
	c.BaseURL, _ = url.Parse(server.URL)

	return mux, c, server
}

func TestOperations_GetVM(t *testing.T) {
	mux, c, server := setup()
	defer server.Close()

	mux.HandleFunc("/PrismGateway/services/rest/v2.0/vms/cfde831a-4e87-4a75-960f-89b0148aa2cc", func(w http.ResponseWriter, r *http.Request) {
		testHTTPMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{"uuid":"cfde831a-4e87-4a75-960f-89b0148aa2cc"}`)
	})

	vmResponse := &VMResponse{}
	vmResponse.UUID = utils.StringPtr("cfde831a-4e87-4a75-960f-89b0148aa2cc")

	type fields struct {
		client *client.Client
	}

	type args struct {
		UUID string
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *VMResponse
		wantErr bool
	}{
		{
			"Test GetVM OK",
			fields{c},
			args{"cfde831a-4e87-4a75-960f-89b0148aa2cc"},
			vmResponse,
			false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			op := Operations{
				client: tt.fields.client,
			}
			got, err := op.GetVM(tt.args.UUID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Operations.GetVM() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Operations.GetVM() = %v, want %v", got, tt.want)
			}
		})
	}
}

func testHTTPMethod(t *testing.T, r *http.Request, expected string) {
	if expected != r.Method {
		t.Errorf("Request method = %v, expected %v", r.Method, expected)
	}
}
