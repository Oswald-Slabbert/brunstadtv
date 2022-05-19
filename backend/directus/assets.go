package directus

import (
	"fmt"
	"net/url"

	"github.com/ansel1/merry"
	"github.com/go-resty/resty/v2"
)

// Sentinel errors
var (
	ErrNotFound = merry.New("No objct was found")
)

// Asset item in the DB
type Asset struct {
	ID              int         `json:"id,omitempty"`
	Name            string      `json:"name"`
	Files           []Assetfile `json:"files"`
	Duration        int         `json:"duration"`
	MediabankenID   string      `json:"mediabanken_id"`
	EncodingVersion string      `json:"encoding_version"`
	MainStoragePath string      `json:"main_storage_path"`
}

// UID returns the id of the Asset
func (a Asset) UID() int {
	return a.ID
}

// TypeName of the item. Statically set to "asset"
func (Asset) TypeName() string {
	return "assets"
}

// FindNewestAssetByMediabankenID in directus
func FindNewestAssetByMediabankenID(c *resty.Client, mediabankenID string) (*Asset, error) {
	q := url.URL{}
	q.Path = "items/assets"

	// Just the newest one
	qq := q.Query()
	qq.Add("limit", "1")
	qq.Add("page", "1")
	qq.Add("sort", "-date_created") // Newest first

	qq.Add("fields[]", "id")
	qq.Add("fields[]", "main_storage_path")
	qq.Add("fields[]", "files.path")

	qq.Add("filter", fmt.Sprintf(`{"_and":[{"mediabanken_id":{"_eq":"%s"}}]}`, mediabankenID))

	x := struct {
		Data []Asset
	}{}

	q.RawQuery = qq.Encode()
	req := c.R().SetResult(x)

	res, err := req.Get(q.String())
	if err != nil {
		return nil, merry.Wrap(err)
	}

	assetList := res.Result().(*struct{ Data []Asset })

	if len(assetList.Data) == 0 {
		return nil, ErrNotFound.Here().WithValue("mediabankenID", mediabankenID)
	}

	return &assetList.Data[0], nil
}