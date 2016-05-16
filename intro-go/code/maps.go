myMap := make(map[string]int)
myMap["meaning of life"] = 42
meaningOfLife := myMap["meaning of life"]

// we can check if the map exists by requesting a second value
meaningOfLife, ok := myMap["meaning of life"]
delete(myMap, "meaning of life")

var uninitializedMap map[string]int
// This will crash because it is not initialized
uninitializedMap["foo"] = 42

// We can also initialize inline
otherMap := map[rune]bool{'F': false}
