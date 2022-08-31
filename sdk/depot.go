package sdk

import (
	"github.com/domenetwork/dome-lib/pkg/cfg"
	"github.com/domenetwork/dome-lib/pkg/common"
	"github.com/domenetwork/dome-lib/pkg/log"
)

// Download a blob of data from Depot that matches the supplied
// query data.
func (sdk *SDK) Download(blob *common.Blob) (err error) {
	log.D("sdk", "depot", "download", blob)
	var v interface{}
	url := cfg.ServiceURL("depot")
	if v, err = sdk.fetch.Post(url, "/download", blob); err != nil {
		return
	}

	err = common.MapToInterface(v.(map[string]interface{}), blob)
	return
}

// Metadata updates the metadata of the supplied blob in Depot.
func (sdk *SDK) Metadata(blob *common.Blob) (err error) {
	log.D("sdk", "depot", "metadata", blob)
	url := cfg.ServiceURL("depot")
	_, err = sdk.fetch.Put(url, "/metadata", blob)
	return
}

// Upload a new blob to Depot for the logged in user.
func (sdk *SDK) Upload(blob *common.Blob) (err error) {
	log.D("sdk", "depot", "upload", blob)
	url := cfg.ServiceURL("depot")
	_, err = sdk.fetch.Post(url, "/upload", blob)
	return
}
