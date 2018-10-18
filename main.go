/*
	wol - A simple go implementation of Wake On LAN.
    Copyright (C) 2018  hcl(HydroChLorica)

    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.

    You should have received a copy of the GNU General Public License
    along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"net"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/alexflint/go-arg"
)

type argList struct {
	Hwaddr []string `arg:"positional"`
	IP     string   `arg:"-i" help:"set the destination IP address"`
	Port   string   `arg:"-p" help:"set the destination port"`
}

type targetInfo struct {
	hwAddr string
	ip     string
	port   string
}

type magicPacket []byte

func (argList) Version() string {
	return "wol-go v1.0"
}

func (t *targetInfo) wake() {
	mp := makeMagicPacket(t.hwAddr)
	sendMagicPacket(mp, t.hwAddr, t.ip, t.port)
}

func checkHwAddr(hwaddr string) bool {
	partRe := "[0-9A-Fa-f]{1,2}"
	hwaddrReList := make([]string, 6)
	for i := range hwaddrReList {
		hwaddrReList[i] = partRe
	}
	hwaddrRe, _ := regexp.Compile(strings.Join(hwaddrReList, ":"))
	if !hwaddrRe.MatchString(hwaddr) {
		return false
	}
	return true
}

func makeMagicPacket(hwaddr string) magicPacket {
	var hwaddrByte []byte
	var mp []byte
	for _, v := range strings.Split(hwaddr, ":") {
		hex, err := strconv.ParseUint(v, 16, 0)
		if err != nil {
			log.Fatalln("Convertion Error")
		}
		hwaddrByte = append(hwaddrByte, byte(hex))
	}
	for i := 0; i < 6; i++ {
		decoded, _ := hex.DecodeString("FF")
		mp = append(mp, decoded...)
	}

	for i := 0; i < 16; i++ {
		mp = append(mp, hwaddrByte...)
	}
	return mp
}

func sendMagicPacket(mp magicPacket, hwaddr string, ip string, port string) {
	var addr strings.Builder
	addr.WriteString(ip)
	addr.WriteString(":")
	addr.WriteString(port)
	socket, err := net.Dial("udp", addr.String())
	if err != nil {
		log.Fatalln("Can't create socket.")
	}
	_, err = socket.Write(mp)
	if err != nil {
		log.Fatalln("Can't send magic packet.")
	}
	fmt.Printf("Sending magic packet to %s:%s with %s \n", ip, port, hwaddr)
	socket.Close()
}

func main() {
	var args argList
	args.IP = "255.255.255.255"
	args.Port = "9"
	arg.MustParse(&args)
	hwAddr := args.Hwaddr
	ip := args.IP
	port := args.Port

	if len(hwAddr) < 1 {
		fmt.Println("No enough argument.")
		os.Exit(1)
	}

	for k, item := range hwAddr {
		var target targetInfo
		target.ip = ip
		target.port = port
		if !checkHwAddr(item) {
			fmt.Printf("Arguments HWADDR error at position %d \n", k)
			log.Fatalf("Invalid hardware address: %s ", item)
		}
		target.hwAddr = item
		target.wake()
	}
}
