package thirty_days_of_code

import "fmt"

func DictionariesAndMaps() {
	var n int
	fmt.Scan(&n)
	phoneBook := make(map[string]int, n)

	for i := 0; i < n; i++ {
		var phone int
		var name string
		fmt.Scanf("%s %d", &name, &phone)
		phoneBook[name] = phone
	}

	for {
		var name string
		fmt.Scan(&name)
		if name == "" {
			break
		}
		if phone, ok := phoneBook[name]; ok {
			fmt.Printf("%s=%d\n", name, phone)
		} else {
			fmt.Printf("Not found\n")
		}
	}
}
