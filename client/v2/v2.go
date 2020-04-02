package v2

import (
	"github.com/terraform-providers/terraform-provider-nutanix/client"
)

const (
	libraryVersion = "v2.0"
	defaultBaseURL = "https://%s/"
	absolutePath   = "/PrismGateway/services/rest/" + libraryVersion
	userAgent      = "nutanix/" + libraryVersion
	mediaType      = "application/json"
)

// Client manages the V2 API
type Client struct {
	client *client.Client
	V2     Service
}

// NewV2Client return a client to operate V2 resources
func NewV2Client(credentials client.Credentials) (*Client, error) {
	c, err := client.NewClient(&credentials, &client.ApiMetadata{
		LibraryVersion: libraryVersion,
		DefaultBaseURL: defaultBaseURL,
		AbsolutePath:   absolutePath,
		UserAgent:      userAgent,
		MediaType:      mediaType,
	},
	)

	if err != nil {
		return nil, err
	}

	f := &Client{
		client: c,
		V2: Operations{
			client: c,
		},
	}

	// f.client.OnRequestCompleted(func(req *http.Request, resp *http.Response, v interface{}) {
	// 	if v != nil {
	// 		utils.PrintToJSON(v, "[Debug] FINISHED REQUEST")
	// 		// TBD: How to print responses before all requests.
	// 	}
	// })

	return f, nil
}
