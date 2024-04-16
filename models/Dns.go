package models

import (
	"strconv"
	"strings"
)

type HostEntry struct {
	ID       int
	IP       string
	Hostname string
	Line int
}

type ByIP []HostEntry
type ByHostname []HostEntry

func (a ByIP) Len() int           { return len(a) }
func (a ByIP) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByIP) Less(i, j int) bool { return LessIP(a[i].IP, a[j].IP) }

func (a ByHostname) Len() int           { return len(a) }
func (a ByHostname) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByHostname) Less(i, j int) bool { return a[i].Hostname < a[j].Hostname }

func LessIP(ip1, ip2 string) bool {
	parts1 := strings.Split(ip1, ".")
	parts2 := strings.Split(ip2, ".")
	for i := 0; i < 4; i++ {
		num1 := Atoi(parts1[i])
		num2 := Atoi(parts2[i])
		if num1 < num2 {
			return true
		} else if num1 > num2 {
			return false
		}
	}
	return false
}

func Atoi(s string) int {
	v, _ := strconv.Atoi(s)
	return v
}
