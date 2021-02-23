package main

import "fmt"

func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	for i := 0; i < n; i++ {
		if gas[i] >= cost[i] {
			if help(gas, cost, i) == true {
				return i
			}
		}
	}
	return -1

}

func help(gas []int, cost []int, i int) bool {
	curr := 0
	n := len(gas)
	j := i
	for j < n {
		curr += gas[j] - cost[j]
		if curr < 0 {
			return false
		}
		j++
	}
	j = 0
	for j < i {
		curr += gas[j] - cost[j]
		if curr < 0 {
			return false
		}
		j++
	}
	return true
}

func canCompleteCircuit1(gas []int, cost []int) int {
	n := len(gas)
	total,curr:=0,0
	start:=0
	for i:=0;i<n;i++{
		total +=gas[i] - cost[i]
		curr+= gas[i]-cost[i]
		if curr<0{
			start=i+1
			curr=0
		}
	}
	if total>=0{
		return start
	}
	return -1


}


func main() {
	gas := []int{1, 2, 3, 4, 5}
	cost := []int{3, 4, 5, 1, 2}
	fmt.Println(canCompleteCircuit(gas, cost))
}
