package main
/*
Andreas Petersen, 6394911
CSI2120 Assignment 6
 */
import (
	"fmt"
	"bufio"
	"os"
	"sync"
	"time"
)

func main(){
	timer := time.NewTimer(time.Second * 2)
	ticker1 := time.NewTicker(time.Millisecond * 200)
	ticker2 := time.NewTicker(time.Millisecond * 300)

	f1, err := os.Open("input1.txt")
	if err != nil {
		panic(err)
	}

	f2, err := os.Open("input2.txt")
	if err != nil {
		panic(err)
	}

	fo, err := os.Create("output.txt")
	if err != nil {
		panic(err)
	}

	chan1 := make(chan string)
	chan2 := make(chan string)
	var wg sync.WaitGroup

	wg.Add(1)
	go func(){
		scanner1 := bufio.NewScanner(f1)
		scanner1.Split(bufio.ScanWords)
		for scanner1.Scan() {
			chan1 <- scanner1.Text()
		}
		close(chan1)
	}()

	go func(){
		scanner2 := bufio.NewScanner(f2)
		scanner2.Split(bufio.ScanWords)
		for scanner2.Scan() {
			chan2 <- scanner2.Text()
		}
		close(chan2)
	}()

	for i:=0;i<70;i++{
		select {
		case <- ticker1.C:
			temp1 := <-chan1
			fo.WriteString(temp1 + " ")
			fmt.Print(temp1 + " ")
		case <- ticker2.C:
			temp2 := <-chan2
			fo.WriteString(temp2 + " ")
			fmt.Print(temp2 + " ")
		case <- timer.C:
			fo.WriteString("\n")
			fmt.Print("\n")
			timer.Reset(time.Second * 2)
		}
	}
	wg.Wait()
}

