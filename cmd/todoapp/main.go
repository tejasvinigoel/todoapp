package main

import (
    "todoapp/router"
)

func main() {
    r := router.SetupRouter()

    // Start the server on port 8080
    r.Run(":8080")
}