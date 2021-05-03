package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"micro/common/config"
	"micro/common/model"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

func serviceGarage() model.GaragesClient {
	port := config.SERVICE_GARAGE_PORT
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		log.Fatal("could not connect to", port, err)
	}

	return model.NewGaragesClient(conn)
}

func serviceUser() model.UsersClient {
	port := config.SERVICE_USER_PORT
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		log.Fatal("could not connect to", port, err)
	}

	return model.NewUsersClient(conn)
}

func main() {
	user1 := model.User{
		Id:       "B00000001",
		Name:     "Fajri Illahi",
		Password: "Tech Lead",
		Gender:   model.UserGender(model.UserGender_value["MALE"]),
	}
	user2 := model.User{
		Id:       "B00000002",
		Name:     "Jin Yi Shi",
		Password: "Senior Sofware Engineer",
		Gender:   model.UserGender(model.UserGender_value["FEMALE"]),
	}

	tool1 := model.Garage{
		Id:   "P00000001",
		Name: "McBook Pro M1",
		Coordinate: &model.GarageCoordinate{
			Latitude:  45.123123123,
			Longitude: 54.1231313123,
		},
	}
	tool2 := model.Garage{
		Id:   "Q00000001",
		Name: "Dell",
		Coordinate: &model.GarageCoordinate{
			Latitude:  32.123123123,
			Longitude: 11.1231313123,
		},
	}
	tool3 := model.Garage{
		Id:   "R00000001",
		Name: "ASUS",
		Coordinate: &model.GarageCoordinate{
			Latitude:  22.123123123,
			Longitude: 123.1231313123,
		},
	}

	user := serviceUser()

	fmt.Println("\n", "===========> user test")

	// register user1
	user.Register(context.Background(), &user1)

	// register user2
	user.Register(context.Background(), &user2)

	// show all registered users
	res1, err := user.List(context.Background(), new(empty.Empty))
	if err != nil {
		log.Fatal(err.Error())
	}
	res1String, _ := json.Marshal(res1.List)
	log.Println(string(res1String))

	garage := serviceGarage()

	fmt.Println("\n", "===========> tool test A")

	// add tool1 to user1
	garage.Add(context.Background(), &model.GarageAndUserId{
		UserId: user1.Id,
		Garage: &tool1,
	})

	// add tool2 to user1
	garage.Add(context.Background(), &model.GarageAndUserId{
		UserId: user1.Id,
		Garage: &tool2,
	})

	// show all tools of user1
	res2, err := garage.List(context.Background(), &model.GarageUserId{UserId: user1.Id})
	if err != nil {
		log.Fatal(err.Error())
	}
	res2String, _ := json.Marshal(res2.List)
	log.Println(string(res2String))

	fmt.Println("\n", "===========> tool test B")

	// add tool3 to user2
	garage.Add(context.Background(), &model.GarageAndUserId{
		UserId: user2.Id,
		Garage: &tool3,
	})

	// show all tools of user2
	res3, err := garage.List(context.Background(), &model.GarageUserId{UserId: user2.Id})
	if err != nil {
		log.Fatal(err.Error())
	}
	res3String, _ := json.Marshal(res3.List)
	log.Println(string(res3String))
}
