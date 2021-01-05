package main

import (
	"hash/fnv"
)


func GetHash(t string) uint64 {
	b := []byte(t)
	num := uint64(14695981039346656037)
	for _,i := range b {
		num = num ^ uint64(i)
		num = num * 1099511628211

	}
	return num
}


func GetHash_golib_system(str string) uint64 {
	// fnv1a from golib from system
	hash := fnv.New64a()
	hash.Write([]byte(str))
	out := hash.Sum64()

	return out
}


