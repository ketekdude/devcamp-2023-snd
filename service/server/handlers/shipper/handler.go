package shipper

import (
	"context"
	"log"

	"github.com/ketekdude/devcamp-2023-snd/service/model"
	pb "github.com/ketekdude/devcamp-2023-snd/service/server/handlers/shipper/proto"
)

func (p *Handler) AddShipper(ctx context.Context, in *pb.AddShipperReq) (resp *pb.AddShipperResp, err error) {
	resp = &pb.AddShipperResp{}

	log.Println("Entering AddShipperHandler")
	res, err := p.shipper.AddShipper(context.Background(), model.ShipperRequest{
		Name:        in.GetName(),
		ImageURL:    in.GetImageURL(),
		Description: in.GetDescription(),
		MaxWeight:   int(in.GetMaxWeight()),
		CreatedBy:   int(in.GetCreatedBy()),
	})
	if err != nil {
		resp.Message = err.Error()
		return
	}
	resp.ID = res.ID
	return
}

func (p *Handler) GetShipper(ctx context.Context, in *pb.GetShipperReq) (resp *pb.GetShipperResp, err error) {
	resp = &pb.GetShipperResp{}

	log.Println("Entering GetShipper Handler")
	res, err := p.shipper.GetShipper(context.Background(), in.GetID())
	if err != nil {
		resp.Message = err.Error()
		return
	}

	for _, e := range res {
		resp.Data = append(resp.Data, &pb.ShipperData{
			ID:          e.ID,
			Name:        e.Name,
			ImageURL:    e.ImageURL,
			Description: e.Description,
			MaxWeight:   int32(e.MaxWeight),
			CreatedAt:   int32(e.CreatedAt.Unix()),
			CreatedBy:   int32(e.CreatedBy),
			UpdatedAt:   int32(e.UpdatedAt.Unix()),
			UpdatedBy:   int32(e.UpdatedBy),
		})
	}
	return
}
