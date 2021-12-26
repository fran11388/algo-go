package main

func main() {

}

func isUgly(n int) bool {
	factors := []int{2, 3, 5}
	for {
		notUgly := true
		for _, factor := range factors {
			if n%factor == 0 {
				n = n / factor
				notUgly = false
				break
			}
		}
		if n == 0 {
			return false
		}
		if n == 1 {
			return true
		}
		if notUgly {
			return false
		}

	}
}
