//go:generate protoc ./teacher.proto --go_out=plugins=grpc:./pb

package teacher

import (
	"context"
	"fmt"
	"log"
	"net"
	"teacher/constants"
	"teacher/models"
	"teacher/models/mongo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"github.com/jinzhu/copier"
	"teacher/pb"
)

type GrpcServer struct {
	Service *mongo.TeacherService
}

func (g GrpcServer) AddTeacher(ctx context.Context, request *pb.AddTeacherRequest) (*pb.AddTeacherResponse, error) {
	var input models.Teacher
	_ = copier.Copy(&input, &request.Input)
	result, err := g.Service.AddTeacher(input)
	if err != nil {
		return nil, err
	}
	respone := new(pb.Teacher)
	copier.Copy(respone,result)
	return &pb.AddTeacherResponse{
		Teacher:              respone,
	}, nil
}

func (g GrpcServer) GetTeacher(ctx context.Context, request *pb.GetTeacherRequest) (*pb.AddTeacherResponse, error) {
	result, err := g.Service.GetTeacher(int(request.ID))
	if err != nil {
		return nil, err
	}
	respone := new(pb.Teacher)
	copier.Copy(respone,result)
	return &pb.AddTeacherResponse{
		Teacher:              respone,
	}, nil
}

func (g GrpcServer) UpdateTeacher(ctx context.Context, request *pb.UpdateTeacherRequest) (*pb.AddTeacherResponse, error) {
	var input models.Teacher
	_ = copier.Copy(&input, &request.Input)
	result, err := g.Service.UpdateTeacher(int(request.ID),input)
	if err != nil {
		return nil, err
	}
	respone := new(pb.Teacher)
	copier.Copy(respone,result)
	return &pb.AddTeacherResponse{
		Teacher:              respone,
	}, nil
}

func (g GrpcServer) DeleteTeacher(ctx context.Context, request *pb.DeleteTeacherRequest) (*pb.AddTeacherResponse, error) {
	result, err := g.Service.DeleteTeacher(int(request.ID))
	if err != nil {
		return nil, err
	}
	respone := new(pb.Teacher)
	copier.Copy(respone,result)
	return &pb.AddTeacherResponse{
		Teacher:              respone,
	}, nil
}

func (g GrpcServer) GetTeachers(ctx context.Context, request *pb.GetTeachersRequest) (*pb.GetTeachersResponse, error) {
	result, err := g.Service.GetTeachers(int(request.SchoolCode))
	if err != nil {
		return nil, err
	}
	var respone []*pb.Teacher
	for _,i :=  range result {
		res := new(pb.Teacher)
		_ = copier.Copy(res,&i)
		respone= append(respone,res)
	}
	copier.Copy(respone,result)
	return &pb.GetTeachersResponse{
		Teacher:              respone,
	}, nil
}

func ListenGRPC(s mongo.TeacherService, port int) (*mongo.TeacherService,error) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))

	if err != nil {
		return nil,err
	}

	serv := grpc.NewServer()
	pb.RegisterTeacherServiceServer(serv, &GrpcServer{&s})
	reflection.Register(serv)
	log.Println("Server Started at ", constants.ServerPort)
	return &s,serv.Serve(lis)
}