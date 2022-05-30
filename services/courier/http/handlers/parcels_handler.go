package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"courier/pkg/responses"
	"courier/pkg/utils"
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

	country := context.Query("country")
	if !utils.IsValidCountry(country) {
		responses.NewContextResponse(context).
			Error().
			Code(http.StatusBadRequest).
			Message("Invalid country!").
			Send()
		return
	}

	guid, _ := uuid.NewRandom()
	parcelsJsonFile := fmt.Sprintf("%s.json", guid)

	// TODO: utilize concurrent fetching from db aka pagination.
	total := int(p.ParcelAdapter.Count(&parsedDay))
	var cargos []*Cargo
	parcels, _ := p.ParcelAdapter.PaginatedParcels(&parsedDay, country, 0, total)
	cargo := new(Cargo)

	for _, parcel := range parcels {
		if cargo.TotalWeight+int(parcel.Weight) >= 500 {
			cargos = append(cargos, cargo)
			cargo = new(Cargo)
		}

		cargo.OrderIDs = append(cargo.OrderIDs, parcel.ParcelID)
		cargo.TotalWeight += int(parcel.Weight)
	}

	// append the last cargo set
	cargos = append(cargos, cargo)

	exportPayload := map[string][]*Cargo{
		"cargos": cargos,
	}

	file, _ := json.MarshalIndent(exportPayload, "", "  ")
	_ = ioutil.WriteFile(parcelsJsonFile, file, 0644)

	context.FileAttachment(parcelsJsonFile, parcelsJsonFile)

	defer func(name string) {
		err := os.Remove(name)
		if err != nil {
			log.Printf("failed to remove generated file: %s", parcelsJsonFile)
		}
	}(parcelsJsonFile)
}
