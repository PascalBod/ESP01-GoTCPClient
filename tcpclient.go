/*
 *  Copyright (C) 2015 Pascal Bodin
 *
 *  This program is free software: you can redistribute it and/or modify
 *  it under the terms of the GNU General Public License as published by
 *  the Free Software Foundation, either version 3 of the License, or
 *  (at your option) any later version.
 *
 *  This program is distributed in the hope that it will be useful,
 *  but WITHOUT ANY WARRANTY; without even the implied warranty of
 *  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *  GNU General Public License for more details.
 *
 *  You should have received a copy of the GNU General Public License
 *  along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 */

package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"strconv"
)

// Server to connect to.
const serverName = "localhost"

// Port to connect to.
const serverPort = 50000

func main() {
	fmt.Println("Connecting to server " + serverName + ":" + strconv.Itoa(serverPort))
	conn, err := net.Dial("tcp", serverName+":"+strconv.Itoa(serverPort))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close();
	
	// Send an empty frame.
	f := []byte{0, 0}
	n, errConn := conn.Write(f);
	if errConn != nil {
		log.Fatal(errConn)
	}
	if n != len(f) {
		log.Fatal(errors.New("Error: some bytes not sent"))
	}
	
	// Send a one-byte frame.
	f = []byte{1, 0, 1}
	n, errConn = conn.Write(f);
	if errConn != nil {
		log.Fatal(errConn)
	}
	if n != len(f) {
		log.Fatal(errors.New("Error: some bytes not sent"))
	}
	
	// Send a two-byte frame.
	f = []byte{2, 0, 1, 2}
	n, errConn = conn.Write(f);
	if errConn != nil {
		log.Fatal(errConn)
	}
	if n != len(f) {
		log.Fatal(errors.New("Error: some bytes not sent"))
	}
	
	// Send a frame too long, and incorrect. Connection should be closed.
	f = []byte{0xFF, 0xFF, 1, 1}
	n, errConn = conn.Write(f);
	if errConn != nil {
		log.Fatal(errConn)
	}
	if n != len(f) {
		log.Fatal(errors.New("Error: some bytes not sent"))
	}
}
