package service

import (
	"context"
	. "go-micro/src/Course"
	"go-micro/src/Mapper"
)

func NewCourseModel (id int32, name string) *CourseModel {
	return &CourseModel{CourseId:id,CourseName:name}
}

type CourseServiceImpl struct {

}

func (this *CourseServiceImpl) ListForTop(ctx context.Context, req *ListRequest, rsp *ListResponse)  error {
	//ret := make([]*CourseModel,0)
	//ret = append(ret,NewCourseModel(101,"java课程"),NewCourseModel(102,"php课程"))
	//rsp.Result = ret
	//return nil

	course := make([]*CourseModel,0)
	err := Mapper.GetCourseList().Find(&course).Error
	if err != nil {
		return err
	}
	rsp.Result = course
	return nil
}

func (this *CourseServiceImpl) GetDetail(ctx context.Context, req *DetailRequest, rsp *DetailResponse) error {
	if err:=Mapper.GetCourseDetail(int(req.CourseId)).Find(rsp.Result).Error;err!=nil {
		return err
	}
	return nil
}

func NewCourseServiceImpl() *CourseServiceImpl {
	return &CourseServiceImpl{}
}
