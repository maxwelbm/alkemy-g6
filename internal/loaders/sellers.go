package loaders

import (
	"encoding/json"
	"os"

	modelsSeller "github.com/maxwelbm/alkemy-g6/internal/models/seller"
)

func NewSellerJSONFile(path string) *SellerJSONFile {
	return &SellerJSONFile{
		path: path,
	}
}

type SellerJSONFile struct {
	path string
}

type SellerJSON struct {
	ID          int    `json:"id,omitempty"`
	CID         int    `json:"cid,omitempty"`
	CompanyName string `json:"company_name,omitempty"`
	Address     string `json:"address,omitempty"`
	Telephone   string `json:"telephone,omitempty"`
}

func (l *SellerJSONFile) Load() (sellers map[int]modelsSeller.Seller, err error) {
	// open file
	file, err := os.Open(l.path)
	if err != nil {
		return
	}
	defer file.Close()

	// decode file
	var sellersJSON []SellerJSON
	err = json.NewDecoder(file).Decode(&sellersJSON)
	if err != nil {
		return
	}

	// serialize sections
	sellers = make(map[int]modelsSeller.Seller)
	for _, s := range sellersJSON {
		sellers[s.ID] = modelsSeller.Seller{
			ID:          s.ID,
			CID:         s.CID,
			CompanyName: s.CompanyName,
			Address:     s.Address,
			Telephone:   s.Telephone,
		}
	}

	return
}
