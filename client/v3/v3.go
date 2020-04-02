package v3

import (
	"github.com/terraform-providers/terraform-provider-nutanix/client"
)

const (
	libraryVersion = "v3"
	defaultBaseURL = "https://%s/"
	absolutePath   = "api/nutanix/" + libraryVersion
	userAgent      = "nutanix/" + libraryVersion
	mediaType      = "application/json"
)

// Client manages the V3 API
type Client struct {
	client *client.Client
	V3     Service
}

// NewV3Client return a client to operate V3 resources
func NewV3Client(credentials client.Credentials) (*Client, error) {
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
		V3: Operations{
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
