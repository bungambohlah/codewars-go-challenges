package buildapileofcubes

func FindNb(m int) int {
	if m > 0 && m < 9223372036854775807 {
		n := 1
		for n < 5000 {
			if n > 1 {
				cubic := (n * n * n)
				j := 1
				for j < n {
					if n-j >= 1 {
						cubic += (n - j) * (n - j) * (n - j)
					}
					j++
				}

				if cubic == m {
					return n
				}
			}
			n++
		}

		return -1
	}

	return -1
}

// NOTE: better solution
// func FindNb(m int) int {
//   for n := 1 ; m > 0 ; n++ {
//     m -= n*n*n
//     if m == 0 { return n }
//   }
//   return -1
// }
