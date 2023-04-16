package helper

import (
	"github.com/mochafiqri/simple-crud/commons/constants"
	"github.com/mochafiqri/simple-crud/commons/entities"
	"github.com/mochafiqri/simple-crud/proto_gen"
)

func ContentsToProtoConten(contents []entities.Content) []*proto_gen.Content {
	res := make([]*proto_gen.Content, 0)
	for _, v := range contents {
		var tmp = proto_gen.Content{
			Id:        v.Id,
			Title:     v.Title,
			Body:      v.Body,
			CreatedAt: v.CreatedAt.Format(constants.FormatTime),
			UpdatedAt: v.UpdateAt.Format(constants.FormatTime),
		}

		res = append(res, &tmp)
	}

	return res
}

func ContentToProtoContent(content entities.Content) *proto_gen.Content {
	return &proto_gen.Content{
		Id:        content.Id,
		Title:     content.Title,
		Body:      content.Body,
		CreatedAt: content.CreatedAt.Format(constants.FormatTime),
		UpdatedAt: content.UpdateAt.Format(constants.FormatTime),
	}
}
