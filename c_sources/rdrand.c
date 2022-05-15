#include <immintrin.h>
#include <stdint.h>

uint64_t crmdnAsUint64()
{
	uint64_t val;
	_rdrand64_step(&val);
	return val;
}

uint32_t crmdnAsUint32()
{
	uint32_t val;
	_rdrand32_step(&val);
	return val;
}

uint16_t crmdnAsUint16()
{
	uint16_t val;
	_rdrand16_step(&val);
	return val;
}

uint8_t RdseedAsuint16(uint16_t val)
{
	return _rdseed16_step(&val);
}

uint8_t RdseedAsuint32(uint32_t val)
{
	return _rdseed32_step(&val);
}

uint8_t RdseedAsuint64(uint64_t val)
{
	return _rdseed64_step(&val);
}
