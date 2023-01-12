package main

import "fmt"

//Create a slice that contains popular movies(at least 2 movies)

// Bonus: Create a "PcParts" struct with title, id and price and create a dynamic list of at least 2 pc parts.
// Then add a third pc part to the list of parts

type pcParts struct{
    id int
    title string
    price float64

}

func main(){
    customPcParts:=[]pcParts{
        {id:1,title:"CPU",price:32.44},
        {id:1,title:"RAM",price:34.64},
    }
    popularMovies:=[]string{"Snatch\n","Casino Royale"}
    fmt.Println(popularMovies)
    fmt.Println(customPcParts)

    // Then add a third pc part to the list of parts
    customPcParts=append(customPcParts,pcParts {id:3,title:"GPU",price:560.38})
    fmt.Println(customPcParts)


}