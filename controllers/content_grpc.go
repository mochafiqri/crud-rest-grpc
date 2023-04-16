package controllers

import (
	"context"
	"fmt"
	"github.com/mochafiqri/simple-crud/commons/constants"
	"github.com/mochafiqri/simple-crud/commons/entities"
	"github.com/mochafiqri/simple-crud/commons/helper"
	"github.com/mochafiqri/simple-crud/commons/interfaces"
	"github.com/mochafiqri/simple-crud/proto_gen"
	"net/http"
)

type ContentGrpc struct {
	proto_gen.UnimplementedContentServiceServer
	uc interfaces.ContentUseCase
}

func NewContentGrpc(uc interfaces.ContentUseCase) *ContentGrpc {
	return &ContentGrpc{uc: uc}
}

func (c *ContentGrpc) Create(ctx context.Context, content *proto_gen.Content) (*proto_gen.ContentResp, error) {
	var tmpContent = entities.Content{
		Title: content.GetTitle(),
		Body:  content.GetBody(),
	}
	code, message, err := c.uc.Create(&tmpContent)
	fmt.Println(err)
	content.Id = tmpContent.Id
	content.CreatedAt = tmpContent.CreatedAt.Format(constants.FormatTime)
	content.UpdatedAt = ""
	return &proto_gen.ContentResp{
		Code:   int32(code),
		Status: message,
		Data:   content,
	}, err
}

func (c *ContentGrpc) Read(ctx context.Context, e *proto_gen.Empty) (*proto_gen.ContentsResp, error) {
	data, code, err := c.uc.Read()
	if err != nil {
		return nil, err
	}

	dataProto := helper.ContentsToProtoConten(data)

	return &proto_gen.ContentsResp{
		Code:   int32(code),
		Status: http.StatusText(code),
		Data:   dataProto,
	}, nil
}

func (c *ContentGrpc) Get(ctx context.Context, id *proto_gen.Id) (*proto_gen.ContentResp, error) {
	content, code, err := c.uc.Get(id.GetId())
	if err != nil {
		return nil, err
	}

	return &proto_gen.ContentResp{
		Code:   int32(code),
		Status: http.StatusText(code),
		Data:   helper.ContentToProtoContent(content),
	}, nil
}

func (c *ContentGrpc) Update(ctx context.Context, content *proto_gen.Content) (*proto_gen.ContentResp, error) {
	var tmpContent = entities.Content{
		Id:    content.Id,
		Title: content.Title,
		Body:  content.Body,
	}

	code, err := c.uc.Update(&tmpContent)
	if err != nil {
		return nil, err
	}

	return &proto_gen.ContentResp{
		Code:   int32(code),
		Status: http.StatusText(code),
		Data:   helper.ContentToProtoContent(tmpContent),
	}, nil
}

func (c *ContentGrpc) Delete(ctx context.Context, id *proto_gen.Id) (*proto_gen.Resp, error) {
	code, err := c.uc.Delete(id.GetId())
	if err != nil {
		return nil, err
	}

	return &proto_gen.Resp{
		Code:   int32(code),
		Status: http.StatusText(code),
	}, nil
}
