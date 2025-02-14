package sellersrp

import (
	"testing"

	"github.com/maxwelbm/alkemy-g6/internal/factories"
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/maxwelbm/alkemy-g6/pkg/testdb"
	"github.com/stretchr/testify/require"
)

func TestGetAll(t *testing.T) {
	db, truncate, teardown := testdb.NewConn(t)
	defer teardown()
	factory := factories.NewSellerFactory(db)
	allFixtures := []models.Seller{
		factory.Build(models.Seller{ID: 1}),
		factory.Build(models.Seller{ID: 2}),
		factory.Build(models.Seller{ID: 3}),
	}

	type arg struct {
		dtos []models.SellerDTO
	}
	type want struct {
		sellers []models.Seller
		err     error
	}
	tests := []struct {
		name  string
		setup func()
		arg
		want
	}{
		{
			name: "When find all Sellers",
			setup: func() {
				for _, fixture := range allFixtures {
					_, err := factory.Create(fixture)
					require.NoError(t, err)
				}
			},
			arg: arg{
				dtos: convertToDTOs(allFixtures),
			},
			want: want{
				sellers: allFixtures,
				err:     nil,
			},
		}, /*{
			name: "When not found Sellers",
			arg: arg{
				dtos: convertToDTOs(allFixtures),
			},
			want: want{
				err: nil,
			},
		},*/
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			t.Cleanup(truncate)
			if tt.setup != nil {
				tt.setup()
			}

			// Arrange
			rp := NewSellersRepository(db)
			// Act

			sellers, err := rp.GetAll()

			// Assert
			if tt.err != nil {
				require.Contains(t, err.Error(), tt.err.Error())
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, tt.want.sellers, sellers)
		})
	}
}

func convertToDTOs(sellers []models.Seller) []models.SellerDTO {
	dtos := make([]models.SellerDTO, len(sellers))
	for i, seller := range sellers {
		dtos[i] = models.SellerDTO{
			ID:          seller.ID,
			CID:         seller.CID,
			CompanyName: seller.CompanyName,
			Address:     seller.Address,
			Telephone:   seller.Telephone,
			LocalityID:  seller.LocalityID,
		}
	}
	return dtos
}
