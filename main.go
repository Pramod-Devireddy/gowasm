/*
 * File         : main.go
 * Project      : GOWASM
 * Created Date : Sunday, May 30th 2021, 12:49:41 PM
 * Author       : Pramod Devireddy
 *
 * Last Modified: Sunday, 30th May 2021 7:35:53 pm
 * Modified By  : Pramod Devireddy
 *
 * Copyright (c)2021 - Pramod Devireddy
 * ***************************************************************
 */

package main

import (
	"log"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("."))))
}
