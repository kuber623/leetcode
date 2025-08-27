package union_set

// UnionSet 并查集（路径压缩 + 按秩合并）
type UnionSet struct {
	parent []int
	rank   []int
}

func NewUnionSet(size int) *UnionSet {
	set := &UnionSet{
		parent: make([]int, size),
		rank:   make([]int, size),
	}
	for i := range set.parent {
		set.parent[i] = i
	}
	return set
}

// Find 查询 x 的根节点
func (set *UnionSet) Find(x int) int {
	if set.parent[x] == x {
		return x
	}
	set.parent[x] = set.Find(set.parent[x])
	return set.parent[x]
}

// Union 合并 x 和 y 节点
func (set *UnionSet) Union(x, y int) {
	xp, yp := set.Find(x), set.Find(y)
	if xp == yp {
		return
	}

	switch {
	case set.rank[xp] < set.rank[yp]:
		set.parent[xp] = yp
	case set.rank[xp] > set.rank[yp]:
		set.parent[yp] = xp
	case set.rank[xp] == set.rank[yp]:
		set.parent[yp] = xp
		set.rank[xp]++
	default:
		panic("unreachable!!!")
	}
}
