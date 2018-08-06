package chs

// return default value if value is empty
func StrOrDefault(value, def string) string {
	if value == "" {
		return def
	}
	return value
}

func IntMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// return default value if value is zero
func IntOrDefault(value, def int) int {
	if value == 0 {
		return def
	}
	return value
}

func I8Max(a, b int8) int8 {
	if a > b {
		return a
	}
	return b
}
func I8Min(a, b int8) int8 {
	if a < b {
		return a
	}
	return b
}

// return default value if value is zero
func I8OrDefault(value, def int8) int8 {
	if value == 0 {
		return def
	}
	return value
}

func I16Max(a, b int16) int16 {
	if a > b {
		return a
	}
	return b
}
func I16Min(a, b int16) int16 {
	if a < b {
		return a
	}
	return b
}

// return default value if value is zero
func I16OrDefault(value, def int16) int16 {
	if value == 0 {
		return def
	}
	return value
}

func I32Max(a, b int32) int32 {
	if a > b {
		return a
	}
	return b
}
func I32Min(a, b int32) int32 {
	if a < b {
		return a
	}
	return b
}

// return default value if value is zero
func I32OrDefault(value, def int32) int32 {
	if value == 0 {
		return def
	}
	return value
}

func I64Max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
func I64Min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

// return default value if value is zero
func I64OrDefault(value, def int64) int64 {
	if value == 0 {
		return def
	}
	return value
}

func UintMax(a, b uint) uint {
	if a > b {
		return a
	}
	return b
}
func UintMin(a, b uint) uint {
	if a < b {
		return a
	}
	return b
}

// return default value if value is zero
func UintOrDefault(value, def uint) uint {
	if value == 0 {
		return def
	}
	return value
}

func U8Max(a, b uint8) uint8 {
	if a > b {
		return a
	}
	return b
}
func U8Min(a, b uint8) uint8 {
	if a < b {
		return a
	}
	return b
}

// return default value if value is zero
func U8OrDefault(value, def uint8) uint8 {
	if value == 0 {
		return def
	}
	return value
}

func U16Max(a, b uint16) uint16 {
	if a > b {
		return a
	}
	return b
}
func U16Min(a, b uint16) uint16 {
	if a < b {
		return a
	}
	return b
}

// return default value if value is zero
func U16OrDefault(value, def uint16) uint16 {
	if value == 0 {
		return def
	}
	return value
}

func U32Max(a, b uint32) uint32 {
	if a > b {
		return a
	}
	return b
}
func U32Min(a, b uint32) uint32 {
	if a < b {
		return a
	}
	return b
}

// return default value if value is zero
func U32OrDefault(value, def uint32) uint32 {
	if value == 0 {
		return def
	}
	return value
}

func U64Max(a, b uint64) uint64 {
	if a > b {
		return a
	}
	return b
}
func U64Min(a, b uint64) uint64 {
	if a < b {
		return a
	}
	return b
}

// return default value if value is zero
func U64OrDefault(value, def uint64) uint64 {
	if value == 0 {
		return def
	}
	return value
}
