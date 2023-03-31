package utils

func Crc16(data []byte) uint16 {
	var (
		crc uint16 = 0xFFFF
		b   byte
		i   int
	)
	for _, b = range data {
		crc ^= uint16(b)
		for i = 0; i < 8; i++ {
			if (crc & 0x0001) != 0 {
				crc >>= 1
				crc ^= 0xA001
			} else {
				crc >>= 1
			}
		}
	}
	return crc
}

func GetCrc16(val int64) uint16 {
	return Crc16(Int64ToBytes(val)) % 16384
}
