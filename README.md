# micro
How to run this system :
1. Download the source code/don't change the name of this root folder
2. Delete the micro/mod.go
3. After deletion, reinitialized the source code with command line "go init mod micro"
4. command line "go tidy mod" to import depedencies
5. Make three windows of terminal
6. First terminal, command line "go run main.go" in directory micro/services/service-garage
7. Second terminal, command line "go run main.go" in directory micro/services/service-user
8. Third terminal, command line "go run main.go" in directory micro/client
9. The output will show the interaction between microservice where the user registration and ownership will be formed among the table/models
