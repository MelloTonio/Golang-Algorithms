func buildTree(root *Node, indexes [][]int32) *Node{
    queue := make([]*Node, 0)
    
    queue = append(queue, root)
    
    for _,e := range indexes{
        temp := queue[0]
        queue = queue[1:]
        
        if e[0] != -1 {
            temp.Left = &Node{Info: e[0]}
            queue = append(queue, temp.Left)
        }
        
        if e[1] != -1 {
            temp.Right = &Node{Info: e[1]}
            queue = append(queue, temp.Right)
        }
    }
    
    return root
}

func swap(root *Node,k int32, level int32, list *[]int32){
    if root == nil{
        return 
    }
    
    if level%k == 0{
        root.Left, root.Right = root.Right, root.Left
    }
    
    swap(root.Left,k,level+1,list)
    *list = append(*list, root.Info)
    swap(root.Right,k, level+1,list)
    
}

func swapNodes(indexes [][]int32, queries []int32) [][]int32 {
    root := &Node{Info: 1}
    root = buildTree(root,indexes)
    
    result := [][]int32{}
    for _,k := range queries{
        l := []int32{}
        swap(root,k,1, &l)
        result = append(result,l)
    }
    
    return result
}
