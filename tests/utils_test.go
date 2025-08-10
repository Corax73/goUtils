package app_test

import (
	"slices"
	"testing"

	goutils "github.com/Corax73/goUtils"
)

func TestGetConfFromEnvFile(t *testing.T) {
	envData := goutils.GetConfFromEnvFile("1")
	if len(envData) == 0 {
		t.Log("Done with incorrect filename")
	} else {
		t.Errorf("Result was incorrect with incorrect filename")
	}
	envData = goutils.GetConfFromEnvFile("./env.test")
	if len(envData) > 0 {
		t.Log("Done with correct filename")
	} else {
		t.Errorf("Result was incorrect with correct filename")
	}
	envData = goutils.GetConfFromEnvFile("")
	if len(envData) > 0 {
		t.Errorf("Result was incorrect with correct filename")
	} else {
		t.Log("Done with incorrect filename")
	}
}

func TestConcatSlice(t *testing.T) {
	errStr := "321"
	correctStr := "123"
	if errStr != goutils.ConcatSlice([]string{"1", "2", "3"}) {
		t.Log("Done with incorrect comparison")
	} else {
		t.Errorf("Result was incorrect with incorrect comparison")
	}
	if correctStr == goutils.ConcatSlice([]string{"1", "2", "3"}) {
		t.Log("Done with correct comparison")
	} else {
		t.Errorf("Result was correct with incorrect comparison")
	}
}

func TestCompareMapsByStringKeys(t *testing.T) {
	map1 := map[string]string{"1": "1", "2": "2"}
	map2 := map[string]string{"1": "1", "2": "2"}
	map3 := map[string]string{"1": "1", "2": "2", "3": "3"}
	if goutils.CompareMapsByStringKeys(map1, map2) {
		t.Log("Done with identical maps")
	} else {
		t.Errorf("Result was incorrect with identical maps")
	}
	if !goutils.CompareMapsByStringKeys(map1, map3) {
		t.Log("Done with unequal maps")
	} else {
		t.Errorf("Result was incorrect with unequal maps")
	}
}

func TestGetMapKeysWithValue(t *testing.T) {
	map1 := map[string]string{"1": "1", "2": "2"}
	map2 := map[string]string{"1": "", "2": "2"}
	if len(goutils.GetMapKeysWithValue(map1)) == 2 {
		t.Log("Done with a map with two non-empty values")
	} else {
		t.Errorf("Result was incorrect with a map with two non-empty values")
	}
	if len(goutils.GetMapKeysWithValue(map2)) == 1 {
		t.Log("Done with a map with one non-empty values")
	} else {
		t.Errorf("Result was incorrect with a map with one non-empty values")
	}
}

func TestGetMapValues(t *testing.T) {
	map1 := map[string]string{"1": "1", "2": "2"}
	map2 := map[string]string{"1": "", "2": "2"}
	testSlice1 := goutils.GetMapValues(map1)
	var hasErr bool
	for _, val := range map1 {
		if !slices.Contains(testSlice1, val) {
			t.Errorf("Result was incorrect for value %s of map1", val)
			hasErr = true
		}
	}
	if !hasErr {
		t.Log("Done with a map with map1")
	}
	testSlice2 := goutils.GetMapValues(map2)
	hasErr = false
	for _, val := range map2 {
		if !slices.Contains(testSlice2, val) && val != "" {
			t.Errorf("Result was incorrect for value %s of map2", val)
			hasErr = true
		}
	}
	if !hasErr {
		t.Log("Done with a map with map2")
	}
}

func TestGetIndexByStrValue(t *testing.T) {
	testSlice := []string{"1", "2"}
	if goutils.GetIndexByStrValue(testSlice, "2") == 1 {
		t.Log("Done with a map with correct index")
	} else {
		t.Errorf("Result was incorrect with correct index")
	}
	if goutils.GetIndexByStrValue(testSlice, "2") == 0 {
		t.Errorf("Result was incorrect with incorrect index")
	} else {
		t.Log("Done with a map with incorrect index")
	}
}

func TestGetMapKeys(t *testing.T) {
	map1 := map[string]string{"1": "1", "2": "2"}
	keys1 := goutils.GetMapKeys(map1)
	map2 := map[string]string{"3": "3", "4": "4"}
	keys2 := goutils.GetMapKeys(map2)
	var hasErr bool
	for key, val := range map1 {
		if !slices.Contains(keys1, key) {
			t.Errorf("Result was incorrect for key %s of map1", val)
			hasErr = true
		}
	}
	if !hasErr {
		t.Log("Done with a map with map1")
	}
	hasErr = false
	for key, val := range map2 {
		if !slices.Contains(keys2, key) {
			t.Errorf("Result was incorrect for key %s of map2", val)
			hasErr = true
		}
	}
	if !hasErr {
		t.Log("Done with a map with map2")
	}
}

func TestPresenceMapKeysInOtherMap(t *testing.T) {
	map1 := map[string]string{"1": "1", "2": "2"}
	map2 := map[string]string{"1": "1", "2": "2"}
	map3 := map[string]string{"3": "3", "4": "4"}
	if goutils.PresenceMapKeysInOtherMap(map1, map2) {
		t.Log("Done with a map with map1 and map2")
	} else {
		t.Errorf("Result was incorrect with map1 and map2")
	}
	if goutils.PresenceMapKeysInOtherMap(map1, map3) {
		t.Errorf("Result was incorrect with map1 and map3")
	} else {
		t.Log("Done with a map with map1 and map3")
	}
}

func TestGetMapWithoutKeys(t *testing.T) {
	map1 := map[string]string{"1": "1", "2": "2", "3": "3"}
	newMap1 := goutils.GetMapWithoutKeys(map1, []string{"1"})
	var hasErr bool
	if _, ok := newMap1["1"]; ok {
		hasErr = true
	}
	if _, ok := newMap1["2"]; !ok {
		hasErr = true
	}
	if _, ok := newMap1["3"]; !ok {
		hasErr = true
	}
	if hasErr {
		t.Errorf("Result was incorrect with existing key")
	} else {
		t.Log("Done with a map with existing key")
	}
	hasErr = false
	newMap2 := goutils.GetMapWithoutKeys(map1, []string{"4"})
	if _, ok := newMap2["1"]; !ok {
		hasErr = true
	}
	if _, ok := newMap2["2"]; !ok {
		hasErr = true
	}
	if _, ok := newMap2["3"]; !ok {
		hasErr = true
	}
	if hasErr {
		t.Errorf("Result was incorrect with a non-existent key")
	} else {
		t.Log("Done with a map with a non-existent key")
	}
}

func TestIsEmail(t *testing.T) {
	validEmailStr := "123_%b@123.com"
	invalidEmailStr1 := "123_%B@123.com"
	invalidEmailStr2 := "123_%B123.com"
	invalidEmailStr3 := "123_%Я@123.com"
	if goutils.IsEmail(validEmailStr) {
		t.Log("Done with a map with with validEmailStr")
	} else {
		t.Errorf("Result was incorrect with validEmailStr")
	}
	if goutils.IsEmail(invalidEmailStr1) {
		t.Errorf("Result was incorrect with invalidEmailStr1")
	} else {
		t.Log("Done with a map with with invalidEmailStr1")
	}
	if goutils.IsEmail(invalidEmailStr2) {
		t.Errorf("Result was incorrect with invalidEmailStr2")
	} else {
		t.Log("Done with a map with with invalidEmailStr2")
	}
	if goutils.IsEmail(invalidEmailStr3) {
		t.Errorf("Result was incorrect with invalidEmailStr3")
	} else {
		t.Log("Done with a map with with invalidEmailStr3")
	}
}

func TestCheckPasswordHash(t *testing.T) {
	passwordValid := "123abc"
	passwordInvalid := "123abc1"
	passwordHash, _ := goutils.HashPassword(passwordValid)
	if goutils.CheckPasswordHash(passwordValid, passwordHash) {
		t.Log("Done with passwordValid")
	} else {
		t.Errorf("Result was incorrect with passwordValid")
	}
	if goutils.CheckPasswordHash(passwordInvalid, passwordHash) {
		t.Errorf("Result was incorrect with passwordInvalid")
	} else {
		t.Log("Done with passwordInvalid")
	}
}
