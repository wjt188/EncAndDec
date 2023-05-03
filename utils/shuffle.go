package utils

import "math/rand"

const seedOffset int64 = 10077

// 给定一个n值，返回【0，n）区间内整数值的一个伪随机排列
func ShuffleN(n int, seed int64) (res []int) {
	randSource := rand.New(rand.NewSource(seed))
	res = randSource.Perm(n) //该函数返回0至n-1共计n个数组成的一个随机排列
	return
}

// 给定一个int类型的切片，将切片内容的元素进行随机的重新排列并返回新的切片
func Shuffle(origin []int, seed int64) (res []int) {
	randSource := rand.New(rand.NewSource(seed))
	l := len(origin)
	res = make([]int, l, l)
	perm := randSource.Perm(l)
	for i, randIndex := range perm {
		res[i] = origin[randIndex]
	}
	return
}

func GenByteMap(seed int64) (m map[byte]byte) {
	m = make(map[byte]byte, 256)
	origin := make([]byte, 256, 256)
	for i, _ := range origin {
		origin[i] = byte(i)
	}
	permTop := rand.New(rand.NewSource(seed)).Perm(128)
	for i, _ := range permTop {
		permTop[i] += 128
	}
	permTail := rand.New(rand.NewSource(seed + seedOffset)).Perm(128)
	perm := append(permTop, permTail...)
	for i, V := range perm {
		m[origin[i]] = byte(V)
	}
	return
}

// 反转map类型的键key和value的值
func ReverseByteMap(m map[byte]byte) (n map[byte]byte) {
	n = make(map[byte]byte, len(m))
	for k, v := range m {
		n[v] = k
	}
	return
}
