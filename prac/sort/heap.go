package sort

func swap(arr []int, i int, j int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}

// Heapify 讓父節點以下的子節點變成堆(heap) 父節點>子節點
func Heapify(tree []int, n int, i int) {
	//i不能大於n
	if i >= n {
		return
	}
	//左邊子節點
	c1 := 2*i + 1
	//右邊子傑點
	c2 := 2*i + 2
	//假設第i個節點最大
	max := i
	//heapify定義 : 讓父節點要大於兩個子節點
	//做heapify
	if c1 < n && tree[c1] > tree[max] {
		max = c1
	}
	if c2 < n && tree[c2] > tree[max] {
		max = c2
	}
	//假設兩邊子節點換完，第i個還不是最大值，就繼續往下換
	if max != i {
		swap(tree, max, i)
		Heapify(tree, n, max)
	}
}

// BuildHeap 建堆
//假設整個樹都是亂的，只要從最後一個子樹的父節點往前做到最上面heapify就可以
func BuildHeap(tree []int, n int) {
	//n是陣列長度
	lastNode := n - 1
	parent := (lastNode - 1) / 2
	for i := parent; i >= 0; i-- {
		Heapify(tree, n, i)
	}
}

// HeapSort
//1. 建堆，根節點一定是最大的
//2. 只要把最後一個節點與根節點交換，最後一個節點一定是最大的
//3. 把最後一個節點拿出去(陣列長度會減少)，在根節點做heapify
//4. 重複2,3
func HeapSort(tree []int, n int) {
	BuildHeap(tree, n)
	for i := n - 1; i >= 0; i-- {
		swap(tree, i, 0)
		Heapify(tree, i, 0)
	}
}
