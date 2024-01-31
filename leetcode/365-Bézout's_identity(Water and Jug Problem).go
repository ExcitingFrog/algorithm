package leetcode

func CanMeasureWater(jug1Capacity int, jug2Capacity int, targetCapacity int) bool {
	if jug1Capacity+jug2Capacity < targetCapacity {
		return false
	}
	if jug1Capacity == 0 || jug2Capacity == 0 {
		return targetCapacity == 0 || jug1Capacity+jug2Capacity == targetCapacity
	}
	return targetCapacity%(gcd(jug1Capacity, jug2Capacity)) == 0
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
