package main

import (
	"fmt"
)

// --- sizef of files
const (
	// this trick saves a lot of hardcoding

	_  = iota             // ignore zero value
	KB = 1 << (10 * iota) // 2^10 = 1024, move bit to 10th place
	MB                    // 2^20 = 1024 * 1024, move bit to 20th place
	GB
	TB
	PB
	EB
	ZB
	YB
)

const (
	isAdmin = 1 << iota
	isHeadquarters
	canSeeFinancials

	canSeeAfrica
	canSeeAsia
	canSeeEurope
)

/* --- fix later when I learn about functions
func (b ByteSize) String() string {
	switch {
	case b >= YB:
		return fmt.Sprintf("%.2fYB", b/YB)
	case b >= ZB:
		return fmt.Sprintf("%.2fZB", b/ZB)
	case b >= EB:
		return fmt.Sprintf("%.2fEB", b/EB)
	case b >= PB:
		return fmt.Sprintf("%.2fPB", b/PB)
	case b >= TB:
		return fmt.Sprintf("%.2fTB", b/TB)
	case b >= GB:
		return fmt.Sprintf("%.2fGB", b/GB)
	case b >= MB:
		return fmt.Sprintf("%.2fMB", b/MB)
	case b >= KB:
		return fmt.Sprintf("%.2fKB", b/KB)
	}
	return fmt.Sprintf("%.2fB", b)
}
*/
func main() {

	//ByteSize(1e13)

	var userRoles byte = isAdmin | canSeeFinancials | canSeeAsia

	fmt.Printf("user roles = %b\n", userRoles)
	fmt.Printf("Is Admin? %t\n", userRoles&isAdmin == isAdmin)
	fmt.Printf("Is HQ? %t", userRoles&isHeadquarters == isHeadquarters)

	fileSize := 2000000000.
	fmt.Printf("\n%.2f GB \n", fileSize/GB)
}
