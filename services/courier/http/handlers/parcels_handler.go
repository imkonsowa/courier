package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"courier/pkg/responses"
	"courier/services/courier/data/adapters"
)

type Cargo struct {
	OrderIDs    []int64 `json:"order_ids"`
	TotalWeight int     `json:"total_weight"`
}

type ParcelHandler struct {
	ParcelAdapter *adapters.ParcelAdapter
}

func NewParcelHandler(adapter *adapters.ParcelAdapter) *ParcelHandler {
	return &ParcelHandler{
		ParcelAdapter: adapter,
	}
}

func (p *ParcelHandler) GenerateCargoReport(context *gin.Context) {
	day := context.Query("day")
	if len(day) == 0 {
		responses.NewContextResponse(context).
			Error().
			Code(http.StatusBadRequest).
			Message("The day to generate report for is missing").
			Send()
		return
	}
	parsedDay, parsedDayErr := time.Parse("2006-01-02", day)
	if parsedDayErr != nil {
		responses.NewContextResponse(context).
			Error().
			Code(http.StatusBadRequest).
			Message("Invalid day value sent").
			Send()
		return
	}

	const ParcelsJsonFile = "parcels-cargo.json"

	parcels, _ := p.ParcelAdapter.PaginatedParcels(&parsedDay, 0, 1000000)

	var cargos []*Cargo
	cargo := new(Cargo)

	for _, parcel := range parcels {
		if cargo.TotalWeight+int(parcel.Weight) >= 500 {
			cargos = append(cargos, cargo)
			cargo = new(Cargo)
		}

		cargo.OrderIDs = append(cargo.OrderIDs, parcel.ParcelID)
		cargo.TotalWeight += int(parcel.Weight)
	}

	file, _ := json.MarshalIndent(cargos, "", " ")
	_ = ioutil.WriteFile("parcels-cargo.json", file, 0644)

	context.FileAttachment("parcels-cargo.json", "parcels-cargo.json")
}
