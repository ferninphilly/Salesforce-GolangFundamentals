package thing

// thing type is NOT exported // HL
type thing struct {
        val int
}

// this variable is NOT exported // HL
var privateThing thing

func ReadThing() int {
        return privateThing.val
}

func WriteThing(val int) {
        privateThing.val = val
}
