package main

import (
	"net/http"

	"example.com/webservice/controllers"
)

func main() {
	// // println("dede")
	// // fmt.Println("deepak")

	// // array learning
	// var arr1 [3]int
	// arr1[0] = 1
	// arr1[1] = 2
	// arr1[2] = 3
	// fmt.Println(arr1)
	// fmt.Println(arr1[1])
	// // fmt.Println(arr1[4]) throws error - index out of bound

	// arr2 := [3]int{5, 6, 7}
	// fmt.Println(arr2)

	// //  slices learning
	// arr3 := [4]int{1, 2, 3, 4}
	// slice1 := arr3[:]
	// fmt.Println(slice1)
	// arr3[1] = 34
	// slice1[2] = 45
	// fmt.Println(arr3, slice1)

	// // slice continued without array
	// slice2 := []int{7, 8, 9, 0}
	// fmt.Println(slice2)

	// slice2 = append(slice2, 10, 11, 12, 13)
	// fmt.Println(slice2)

	// // slice of slices
	// s2 := slice2[1:]
	// s3 := slice2[:4]
	// s4 := slice2[1:3]
	// fmt.Println(s2)
	// fmt.Println(s3)
	// fmt.Println(s4)

	// // maps
	// fmt.Println("===== maps =====")
	// m := map[string]int{"foo": 42}
	// fmt.Println(m)
	// fmt.Println(m["foo"])

	// m["foo"] = 2
	// fmt.Println(m)

	// delete(m, "foo")
	// fmt.Println(m)

	// m1 := map[string]int{"foo": 42, "bar": 34}
	// fmt.Println(m1)

	// // struct datatypes
	// fmt.Println("===== structs =====")

	// type user1 struct {
	// 	ID        int
	// 	FirstName string
	// 	LastName  string
	// }

	// var u user1
	// u.ID = 1
	// u.FirstName = "Deepak"
	// u.LastName = "Dargade"
	// fmt.Println(u)
	// // fmt.Println(u.FirstName)

	// u2 := user1{ID: 2, FirstName: "Nima", LastName: "Dargade"}
	// fmt.Println(u2)

	// u3 := user1{
	// 	ID:        3,
	// 	FirstName: "Tanvi",
	// 	LastName:  "Dargade",
	// }
	// fmt.Println(u3)

	// fmt.Println("welcome")

	// // use User module

	// user := models.User{
	// 	ID:        2,
	// 	FirstName: "Deepak",
	// }
	// fmt.Println(user)

	// // functions demo

	// port := 3000
	// _, err := startWebServer(port, 3)
	// fmt.Println(port, err)

	controllers.RegisterController()
	http.ListenAndServe(":3000", nil)
}

// func startWebServer(port, numberOfRetires int) (int, error) {
// 	fmt.Println("Starting the Server on port:", port)
// 	fmt.Println("// loading, retries //", numberOfRetires)
// 	fmt.Println("Server started on port:", port)
// 	// errors.New("something went wrong")
// 	return port, nil
// }
