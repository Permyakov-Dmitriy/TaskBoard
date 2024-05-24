package utils

import (
	"github.com/jinzhu/copier"
)

func TransformSingleModelToResponse[ResT any, ModT any](obj *ModT) ResT {
	var responseObject ResT
	copier.Copy(&responseObject, obj)
	return responseObject
}

func TransformSliceModelToResponse[ResT any, ModT any](objs []ModT) []ResT {
	var responseObjects []ResT
	copier.Copy(&responseObjects, &objs)
	return responseObjects
}

