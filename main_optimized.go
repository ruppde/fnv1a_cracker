package main

import (
	"bufio"
	"flag"
	"log"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// 17291806236368054941 = solarwinds.businesslayerhost

/*
	Process name, service name, and driver path listings are obtained, 
	and each value is hashed via the FNV-1a + XOR algorithm as described previously 
	and checked against hardcoded blocklists.
	Some of these hashes have been brute force reversed as part of this analysis, 
	showing that these routines are scanning for analysis tools and antivirus engine components. 
*/

func main(){
	
	var onlymatches bool
	flag.BoolVar(&onlymatches, "m", false, "only show matches from known hardcoded hashes")

	flag.Parse()
	
	// accept input piped to program, or from an arg
	var processes io.Reader
	processes = os.Stdin

	process := flag.Arg(0)

	if process != "" {
		processes = strings.NewReader(process)
	}

	sc := bufio.NewScanner(processes)

	lines, err := readLines("hardcoded_hashes_left.txt")
	if err != nil {
	    log.Fatalf("readLines: %s", err)
	}
	for sc.Scan() {
		// convert to lower and check not seen
		lower_process := strings.ToLower(sc.Text())
		// get hash and print
		currentHash := GetHash(lower_process)
		currentHashStr := strconv.FormatUint(currentHash, 10)

		//func stringInSlice(a string, list []string) bool {
		for _, b := range lines {
			if b == currentHashStr {
				fmt.Printf("%s : %s\n", currentHashStr, lower_process)
			}
		}
	}

}

// replicate the .net hash function
func GetHash(t string) uint64 {
	b := []byte(t)
	num := uint64(14695981039346656037)
	val := uint64(6605813339339102567)
	for _,i := range b {
		num = num ^ uint64(i)
		num = num * 1099511628211
	}
	return num ^ val
}

func contains(slice []string, item string) bool {
    set := make(map[string]struct{}, len(slice))
    for _, s := range slice {
        set[s] = struct{}{}
    }

    _, ok := set[item] 
    return ok
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    return lines, scanner.Err()
}

// original .NET code:

// Private Shared Function GetHash(s As String) As ULong
// 	Dim num As ULong = 14695981039346656037UL
// 	Try
// 		For Each b As Byte In Encoding.UTF8.GetBytes(s)
// 			num = num Xor CULng(b)
// 			num *= 1099511628211UL
// 		Next
// 	Catch
// 	End Try
// 	Return num Xor 6605813339339102567UL
// End Function
