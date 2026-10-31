func handleString(s string) string {
	res := ""
	check := false
	for i := 0; i < len(s); i++ {
		str := string(s[i])
		if str == "." && !check {
			continue
		}
		if str == "+" {
			i++
			for str != "@" {
				i++
				str = string(s[i])
			}
			check = true
		}
		fmt.Println(res)
		res = res + str
	}
	return res
}

func numUniqueEmails(emails []string) int {
	m := make(map[string]int)
	for _, val := range emails {
		m[handleString(val)] += 1
	}
	return len(m)
}

