package disjoint

type DisjointSet struct {
    // parent[i] is the parent of i
    parent []int
    // rank[i] is the rank of i
    rank []int
    // size[i] is the size of the set whose root is i
    size []int
    // count is the number of disjoint sets
    count int
    // maxSize is the size of the largest set
    maxSize int
    // minSize is the size of the smallest set
    minSize int
    // maxSizeIndex is the index of the largest set
    maxSizeIndex int  
}

// NewDisjointSet initializes a new DisjointSet with n elements.
func NewDisjointSet(n int) *DisjointSet {
    ds := &DisjointSet{
        parent: make([]int, n),
        rank:   make([]int, n),
        size:   make([]int, n),
        count:  n,
    }
    for i := 0; i < n; i++ {
        ds.parent[i] = i
        ds.size[i] = 1
    }
    ds.maxSize = 1
    ds.minSize = 1
    return ds
}

// Find returns the root of the set containing element x with path compression.
func (ds *DisjointSet) Find(x int) int {
    if ds.parent[x] != x {
        ds.parent[x] = ds.Find(ds.parent[x]) // Path compression
    }
    return ds.parent[x]
}

// Union merges the sets containing elements x and y.
func (ds *DisjointSet) Union(x, y int) {
    rootX := ds.Find(x)
    rootY := ds.Find(y)

    if rootX != rootY {
        // Union by rank
        if ds.rank[rootX] > ds.rank[rootY] {
            ds.parent[rootY] = rootX
            ds.size[rootX] += ds.size[rootY]
            if ds.size[rootX] > ds.maxSize {
                ds.maxSize = ds.size[rootX]
                ds.maxSizeIndex = rootX
            }
        } else if ds.rank[rootX] < ds.rank[rootY] {
            ds.parent[rootX] = rootY
            ds.size[rootY] += ds.size[rootX]
            if ds.size[rootY] > ds.maxSize {
                ds.maxSize = ds.size[rootY]
                ds.maxSizeIndex = rootY
            }
        } else {
            ds.parent[rootY] = rootX
            ds.size[rootX] += ds.size[rootY]
            ds.rank[rootX]++
            if ds.size[rootX] > ds.maxSize {
                ds.maxSize = ds.size[rootX]
                ds.maxSizeIndex = rootX
            }
        }
        ds.count--
    }
}

// GetCount returns the number of disjoint sets.
func (ds *DisjointSet) GetCount() int {
    return ds.count
}

// GetSize returns the size of the set containing element x.
func (ds *DisjointSet) GetSize(x int) int {
    root := ds.Find(x)
    return ds.size[root]
}