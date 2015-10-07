package main

import (
	"fmt"

	"github.com/JaegerBane/stringutil"
)

func main() {

    /* Main is full of random junk at the moment */

    const COMP int = 10

    var stuff int = COMP

    /* const declaration of type int */

    fmt.Printf("alreet\n")

    /* Printf only takes strings initially - %T shortcut for variable type? */

    fmt.Printf("stuff is of type %T\n",stuff)

    /* Println prints anything - printing an int here */

    fmt.Println(stuff)

    /* stuff from tutorial */

    fmt.Printf(stringutil.Reverse("!oG ,olleH") + "\n")

    fmt.Printf(stringutil.Reverse("Hello go") + "\n")

    

}
