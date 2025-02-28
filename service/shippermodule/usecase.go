package shippermodule

import (
	"context"
	"log"
	"time"

	m "github.com/ketekdude/devcamp-2023-snd/service/model"
)

func (p *Module) AddShipper(ctx context.Context, data m.ShipperRequest) (result m.ShipperResponse, err error) {
	data, err = SanitizeInsert(data)
	if err != nil {
		log.Println("[ShipperModule][AddShipper] bad request, err: ", err.Error())
		return
	}

	result, err = p.Storage.AddShipper(ctx, data)
	if err != nil {
		log.Println("[ShipperModule][AddShipper] problem in getting from storage, err: ", err.Error())
		return
	}

	return
}

func (p *Module) GetShipper(ctx context.Context, id int64) (result []m.ShipperResponse, err error) {
	if id != 0 {
		var res m.ShipperResponse
		res, err = p.Storage.GetShipper(ctx, id)
		if err != nil {
			log.Println("[ShipperModule][GetShipper] problem getting storage data, err: ", err.Error())
			return
		}
		result = append(result, res)
	} else {
		result, err = p.GetShipperAll(ctx)
	}
	return
}

func (p *Module) GetShipperAll(ctx context.Context) (result []m.ShipperResponse, err error) {
	result, err = p.Storage.GetShipperAll(ctx)
	if err != nil {
		log.Println("[ShipperModule][GetShipperAll] problem getting storage data, err: ", err.Error())
		return
	}

	return
}

func (p *Module) UpdateShipper(ctx context.Context, id int64, data m.ShipperRequest) (result m.ShipperResponse, err error) {
	data.UpdatedAt = time.Now()
	result, err = p.Storage.UpdateShipper(ctx, id, data)
	if err != nil {
		log.Println("[ShipperModule][UpdateShipper] problem getting storage data, err: ", err.Error())
		return
	}

	return
}
