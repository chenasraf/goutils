package utils

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sort"

	"github.com/rivo/uniseg"
)

func MapSlice[T, V any](ts []T, fn func(T) V) []V {
	result := make([]V, len(ts))
	for i, t := range ts {
		result[i] = fn(t)
	}
	return result
}

func HandleErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func FileExists(path string) bool {
	_, err := os.Stat(path)
	exists := !os.IsNotExist(err)
	return exists
}

func GlobExists(path string) bool {
	res, err := filepath.Glob(path)
	HandleErr(err)
	return res != nil
}

func ReadFile(path string) string {
	res, err := os.ReadFile(path)
	HandleErr(err)
	return string(res)
}

func MapKeys[K comparable, V any](mapObj map[K]V) []K {
	keys := make([]K, len(mapObj))
	i := 0
	for k := range mapObj {
		keys[i] = k
		i++
	}
	return keys
}

func SliceContains[T comparable](list []T, item T) bool {
	for _, v := range list {
		if v == item {
			return true
		}
	}
	return false
}

func RunCmd(cmd string, args ...string) (string, error) {
	res, err := exec.Command(cmd, args...).Output()
	return string(res), err
}

func Ternary[T any](cond bool, whenTrue T, whenFalse T) T {
	if cond {
		return whenTrue
	}
	return whenFalse
}

func Insert[T any](a []T, i int, item T) []T {
	return append(a[:i], append([]T{item}, a[i:]...)...)
}

func SortAlphanumeric(slice []string) []string {
	sort.SliceStable(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})
	return slice
}

func StrLen(str string) int {
	return uniseg.GraphemeClusterCount(str)
}
