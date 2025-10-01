package main

import (
	"fmt"
	"strings"
)

func XORDivision(dividend, divisor string) string {
	dividendBytes := []byte(dividend)
	divisorLen := len(divisor)

	for i := 0; i <= len(dividendBytes)-divisorLen; i++ {
		if dividendBytes[i] == '1' {
			for j := 0; j < divisorLen; j++ {
				if dividendBytes[i+j] == divisor[j] {
					dividendBytes[i+j] = '0'
				} else {
					dividendBytes[i+j] = '1'
				}
			}
		}
	}
	remainder := string(dividendBytes[len(dividendBytes)-(divisorLen-1):])
	return remainder
}

func validateBinaryString(s string) bool {
	for _, char := range s {
		if char != '0' && char != '1' {
			return false
		}
	}
	return true
}



func main() {
	fmt.Println("CRC Error Detection Algorithm")
	fmt.Println("=============================")

	var dataWord, divisor string

	for {
		fmt.Print("Step 1: Enter the data word (binary): ")
		fmt.Scanln(&dataWord)

		if validateBinaryString(dataWord) && len(dataWord) > 0 {
			break
		} else {
			fmt.Println("Invalid input! Please enter a binary string (only 0s and 1s).")
		}
	}

	for {
		fmt.Print("Step 2: Enter the divisor/generator polynomial (binary): ")
		fmt.Scanln(&divisor)

		if validateBinaryString(divisor) && len(divisor) > 1 && divisor[0] == '1' {
			break
		} else {
			fmt.Println("Invalid input! Divisor must be binary, length > 1, and start with '1'.")
		}
	}

	k := len(dataWord)
	r := len(divisor) - 1
	n := k + r

	fmt.Printf("\nStep 3: Calculating redundant bits:\n")
	fmt.Printf("- Length of data word (k): %d\n", k)
	fmt.Printf("- Length of divisor: %d\n", len(divisor))
	fmt.Printf("- Number of redundant bits (r): %d\n", r)
	fmt.Printf("- Total code word length (n): %d\n", n)

	newDataWord := dataWord + strings.Repeat("0", r)
	fmt.Printf("\nStep 4: Appending %d zero bits:\n", r)
	fmt.Printf("- Original data word: %s\n", dataWord)
	fmt.Printf("- New data word: %s\n", newDataWord)

	fmt.Printf("\nStep 5: Performing XOR division:\n")
	fmt.Printf("- Dividend: %s\n", newDataWord)
	fmt.Printf("- Divisor: %s\n", divisor)

	remainder := XORDivision(newDataWord, divisor)

	fmt.Printf("\nStep 6: Division remainder: %s\n", remainder)

	for len(remainder) < r {
		remainder = "0" + remainder
	}

	codeWord := dataWord + remainder
	fmt.Printf("\nStep 7: Creating final code word:\n")
	fmt.Printf("- Original data word: %s\n", dataWord)
	fmt.Printf("- CRC remainder: %s\n", remainder)
	fmt.Printf("- Final code word: %s\n", codeWord)

	fmt.Printf("\nStep 8: Final Result:\n")
	fmt.Printf("====================================\n")
	fmt.Printf("Data Word: %s\n", dataWord)
	fmt.Printf("Divisor: %s\n", divisor)
	fmt.Printf("CRC Code Word: %s\n", codeWord)
	fmt.Printf("====================================\n")

	fmt.Printf("\nVerification:\n")
	fmt.Printf("- Data bits: %d\n", len(dataWord))
	fmt.Printf("- CRC bits: %d\n", len(remainder))
	fmt.Printf("- Total bits: %d\n", len(codeWord))

	fmt.Printf("\nError Detection Test:\n")
	fmt.Println("To test error detection, divide the code word by the divisor.")
	fmt.Println("If remainder is all zeros, no error detected.")

	testRemainder := XORDivision(codeWord, divisor)

	for len(testRemainder) < r {
		testRemainder = "0" + testRemainder
	}

	fmt.Printf("Verification remainder: %s\n", testRemainder)

	allZeros := true
	for _, bit := range testRemainder {
		if bit == '1' {
			allZeros = false
			break
		}
	}

	if allZeros {
		fmt.Println("✓ No errors detected - CRC calculation is correct!")
	} else {
		fmt.Println("✗ Error detected in CRC calculation!")
	}
}
