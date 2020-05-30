package utils

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"log"
	"teacher"
	"teacher/constants"
	"teacher/models/mongo"
)

type Container struct{
	Server    *teacher.GrpcServer
	Service   *mongo.TeacherService
	DB        *mgo.Session
	Injected                 bool
}
func (c *Container) TriggerDI() error{
	if c.Injected {
		return nil
	}
	err := c.StartServer()
	if err != nil {
		c.Injected = false
		return err
	}
	c.Injected = true
	return nil
}

func (c *Container) StartServer() error{
	if c.Server == nil {
		var err error
		c.Server.Service,err = teacher.ListenGRPC(*c.GetService(),constants.ServerPort)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Container) GetService() *mongo.TeacherService {
	if c.Service == nil {
		c.Service = &mongo.TeacherService{
			Database: c.GetDB(),
		}
	}
	return c.Service
}

func (c *Container) GetDB() *mgo.Session{
	if c.DB == nil {
		dsn := fmt.Sprintf("127.0.0.1:27017")
		db,err  := MongoRepository(dsn,"","")
		if err != nil {
			log.Fatal("Error Connecting to DB",err)
		}
		c.DB = db
	}

	return c.DB
}