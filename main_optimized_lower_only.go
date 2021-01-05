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
	"hash/fnv"
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

	var onlymatches, debug bool
	flag.BoolVar(&onlymatches, "m", false, "only show matches from known hardcoded hashes")
	flag.BoolVar(&debug, "d", false, "use all hashes")

	flag.Parse()

	// accept input piped to program, or from an arg
	var processes io.Reader
	processes = os.Stdin

	process := flag.Arg(0)

	if process != "" {
		processes = strings.NewReader(process)
	}

	sc := bufio.NewScanner(processes)

	var lines []uint64
	var err error
	if debug {
		lines, err = readLines("hardcoded_hashes.txt")
	} else {
		lines, err = readLines("hardcoded_hashes_left.txt")
	}


	if err != nil {
	    log.Fatalf("readLines: %s", err)
	}
	for sc.Scan() {

		currentHashStr := GetHash(sc.Text())

		//func stringInSlice(a string, list []string) bool {
		for _, b := range lines {
			if b == currentHashStr {
				fmt.Printf("%s : %s\n", strconv.FormatUint(currentHashStr, 10), sc.Text())
			}
		}
	}

}

// replicate the .net hash function fnv1a
func GetHash(t string) uint64 {
	b := []byte(t)
	num := uint64(14695981039346656037)
	for _,i := range b {
		num = num ^ uint64(i)
		num = num * 1099511628211
	}
	// here we leave the fnv1a standard, xor from sunburst
	return num ^ uint64(6605813339339102567)
}

func GetHash_golib(str string) uint64 {
	// fnv1a from golib, 4 times slower on short strings than the function above, see folder ./bench/ = not used
	hash := fnv.New64a()
	hash.Write([]byte(str))
	out := hash.Sum64()

	// here we leave the standard, xor from sunburst
	return out ^ uint64(6605813339339102567)
}


// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]uint64, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var lines []uint64
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
		// convert hashes to uint64 once to avoid converting each generated uint64 hash to string
		hash,_ := strconv.ParseUint(scanner.Text(), 10, 64)
        lines = append(lines, hash)
        //lines = append(lines, scanner.Text())
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
