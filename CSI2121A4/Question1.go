package main
/*
Andreas Petersen, 6394911
CSI2120 Assignment 6
 */
import (
	"fmt"
)

type House struct{
	rooms []string
	name string
	sizeFt []roomSz
}

type roomSz struct{
	width float32
	length float32
}

type Semi struct{
	House
}

type Home interface{
	inputSqft()
	printMetric()
}

func NewHouse() *House{
	rs := []roomSz{}
	h := House{rooms: []string{"kitchen", "living", "dining", "main"}, name: "House", sizeFt: rs}
	return &h
}

func NewHouseRooms(roomList []string) *House{
	rs := []roomSz{}
	rSlice := []string{"kitchen", "living", "dining", "main"}
	h := House{rooms: append(rSlice, roomList...), name: "House", sizeFt: rs}
	return &h
}

func NewSemi() *Semi{
	rs := []roomSz{}
	h := Semi{House{rooms: []string{"kitchen", "living", "dining", "main"}, name: "Semi", sizeFt: rs}}
	return &h
}

func NewSemiRooms(roomList []string) *Semi{
	rs := []roomSz{}
	rSlice := []string{"kitchen", "living", "dining", "main"}
	h := Semi{House{rooms: append(rSlice, roomList...), name: "Semi", sizeFt: rs}}
	return &h
}

func (h *House) inputSqft(){
	fmt.Print("For the " + h.name + " please:\n")
	for i := 0; i<len(h.rooms); i++ {
		fmt.Print("Enter the width of the " + h.rooms[i] + ":")
		var width float32
		fmt.Scan(&width)
		fmt.Print("Enter the length of the " + h.rooms[i] + ":")
		var length float32
		fmt.Scan(&length)

		rs := roomSz{width: width, length: length}
		h.sizeFt = append(h.sizeFt, rs)
	}
}

func (h *House) printMetric(){
	fmt.Print(h.name + "\n")
	for i := 0; i<len(h.rooms); i++ {
		fmt.Print(h.rooms[i] + " : width x length: ")
		fmt.Print(h.sizeFt[i].width)
		fmt.Print("x")
		fmt.Print(h.sizeFt[i].length)
		fmt.Print("\n")
	}
}

func main() {
	homes := []Home{NewHouse(), NewSemi(), NewHouseRooms([]string{"bedroom1", "bedroom2"})}

	for i := 0; i<len(homes); i++{
		homes[i].inputSqft()
	}

	for i := 0; i<len(homes); i++{
		homes[i].printMetric()
	}
}