package loaders

import (
	"encoding/json"
	"os"

	modelsBuyer "github.com/maxwelbm/alkemy-g6/internal/models/buyer"
)

func NewBuyerJSONFile(path string) *BuyerJSONFile {
	return &BuyerJSONFile{
		path: path,
	}
}

type BuyerJSONFile struct {
	path string
}

type BuyerJSON struct {
	Id           int    `json:"id"`
	CardNumberId string `json:"card_number_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
}

func (l *BuyerJSONFile) Load() (buyers map[int]modelsBuyer.Buyer, err error) {
	// open file
	file, err := os.Open(l.path)
	if err != nil {
		return
	}
	defer file.Close()

	// decode file
	var buyersJSON []BuyerJSON
	err = json.NewDecoder(file).Decode(&buyersJSON)
	if err != nil {
		return
	}

	// serialize sections
	buyers = make(map[int]modelsBuyer.Buyer)
	for _, s := range buyersJSON {
		buyers[s.Id] = modelsBuyer.Buyer{

			Id:           s.Id,
			CardNumberId: s.CardNumberId,
			FirstName:    s.FirstName,
			LastName:     s.LastName,
		}
	}

	return
}
