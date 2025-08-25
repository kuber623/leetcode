package leetcode827

// https://leetcode.cn/problems/making-a-large-island/

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

func largestIsland(grid [][]int) int {
	// 获取矩阵长度
	n := len(grid)
	if n == 0 {
		return 0
	}

	// 将相邻地块合并成岛屿
	set := NewUnionSet(n * n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				continue
			}
			// 向上连接
			if i != 0 && grid[i-1][j] == 1 {
				set.Union(mappingIndex(i, j, n), mappingIndex(i-1, j, n))
			}
			// 向左连接
			if j != 0 && grid[i][j-1] == 1 {
				set.Union(mappingIndex(i, j, n), mappingIndex(i, j-1, n))
			}
		}
	}

	// 计算每块岛屿的面积
	islandSizes := make(map[int]int, n*n/2)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				continue
			}
			root := set.Find(mappingIndex(i, j, n))
			islandSizes[root]++
		}
	}

	// 获取每个海面所有相邻岛屿的大小
	maxSize := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				continue
			}

			// 如果为海面
			adjacentSet := make(map[int]int, 4)
			// 向上查询
			if i != 0 && grid[i-1][j] == 1 {
				root := set.Find(mappingIndex(i-1, j, n))
				adjacentSet[root] = islandSizes[root]
			}
			// 向下查询
			if i != n-1 && grid[i+1][j] == 1 {
				root := set.Find(mappingIndex(i+1, j, n))
				adjacentSet[root] = islandSizes[root]
			}
			// 向左查询
			if j != 0 && grid[i][j-1] == 1 {
				root := set.Find(mappingIndex(i, j-1, n))
				adjacentSet[root] = islandSizes[root]
			}
			// 向右查询
			if j != n-1 && grid[i][j+1] == 1 {
				root := set.Find(mappingIndex(i, j+1, n))
				adjacentSet[root] = islandSizes[root]
			}

			// 计算连接岛屿后的最大面积
			curSize := 1
			for _, v := range adjacentSet {
				curSize += v
			}

			if curSize > maxSize {
				maxSize = curSize
			}
		}
	}

	// 边界条件：网格中所有格子均为地块
	if maxSize == 0 {
		return n * n
	}

	return maxSize
}

func mappingIndex(x, y, n int) int {
	return x*n + y
}
