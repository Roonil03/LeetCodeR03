// func deleteDuplicateFolder(paths [][]string) [][]string {
//     root := &TrieNode{Children: make(map[string]*TrieNode), Name: ""}
// 	for _, path := range paths {
// 		current := root
// 		for _, folder := range path {
// 			if _, ok := current.Children[folder]; !ok {
// 				current.Children[folder] = &TrieNode{Children: make(map[string]*TrieNode), Name: folder}
// 			}
// 			current = current.Children[folder]
// 		}
// 	}
// 	markDuplicates(root)
// 	var result []string
// 	collectRemaining(root, []string{}, &result)
//     ans := [][]string{}
// 	for _, p := range result {
// 		ans = append(ans, strings.Split(p[1:], "/"))
//     }
//     return ans
// }

// type TrieNode struct {
// 	Children map[string]*TrieNode
// 	IsMarked bool
// 	Name     string
// }

// func computeSignature(node *TrieNode) string {
// 	if node == nil {
// 		return ""
// 	}
// 	sigs := []string{}
// 	for name, child := range node.Children {
// 		sigs = append(sigs, name+":"+computeSignature(child))
// 	}
// 	sort.Strings(sigs)
// 	return "("+strings.Join(sigs, ",")+")"
// }

// func markDuplicates(root *TrieNode) {
// 	sigMap := make(map[string][]*TrieNode)
// 	dfsCollectSignatures(root, sigMap)
// 	for _, nodes := range sigMap {
// 		if len(nodes) > 1 {
// 			for _, n := range nodes {
// 				markSubtree(n)
// 			}
// 		}
// 	}
// }

// func dfsCollectSignatures(node *TrieNode, sigMap map[string][]*TrieNode) {
// 	if node == nil {
// 		return
// 	}
// 	sig := computeSignature(node)
// 	sigMap[sig] = append(sigMap[sig], node)
// 	for _, child := range node.Children {
// 		dfsCollectSignatures(child, sigMap)
// 	}
// }

// func markSubtree(node *TrieNode) {
// 	if node == nil {
// 		return
// 	}
// 	node.IsMarked = true
// 	for _, child := range node.Children {
// 		markSubtree(child)
// 	}
// }

// func collectRemaining(node *TrieNode, currentPath []string, result *[]string) {
// 	if node == nil {
// 		return
// 	}
// 	if !node.IsMarked && node.Name != "" {
// 		path := "/" + strings.Join(currentPath, "/")
// 		*result = append(*result, path)
// 	}
// 	for name, child := range node.Children {
// 		currentPath = append(currentPath, name)
// 		collectRemaining(child, currentPath, result)
// 		currentPath = currentPath[:len(currentPath)-1]
// 	}
// }

/*type Node struct {
    children map[string]*Node
    deleted  bool
}

func newNode() *Node {
    return &Node{children: make(map[string]*Node)}
}

func deleteDuplicateFolder(paths [][]string) [][]string {
    root := newNode()
    for _, path := range paths {
        cur := root
        for _, name := range path {
            if _, ok := cur.children[name]; !ok {
                cur.children[name] = newNode()
            }
            cur = cur.children[name]
        }
    }
    groups := make(map[string][]*Node)
    encode(root, groups)
    for _, nodes := range groups {
        if len(nodes) > 1 {
            for _, n := range nodes {
                n.deleted = true
            }
        }
    }
    var result [][]string
    collect(root, []string{}, &result)
    return result
}

func encode(node *Node, groups map[string][]*Node) string {
    if len(node.children) == 0 {
        return "()"
    }
    parts := make([]string, 0, len(node.children))
    for name, child := range node.children {
        parts = append(parts, name+encode(child, groups))
    }
    sort.Strings(parts)
    sign := "(" + strings.Join(parts, "") + ")"
    groups[sign] = append(groups[sign], node)
    return sign
}

func collect(node *Node, path []string, res *[][]string) {
    for name, child := range node.children {
        if child.deleted {
            continue
        }
        newPath := append(path, name)
        *res = append(*res, newPath)
        collect(child, newPath, res)
    }
}*/

/*type Node struct {
	subNodes map[string]*Node
	content  string
	remove   bool
}

func newNode() *Node {
	return &Node{
		subNodes: make(map[string]*Node),
	}
}

func (node *Node) markRemove() {
	if node.remove {
		return
	}
	node.remove = true
	for _, value := range node.subNodes {
		value.markRemove()
	}
}

func deleteDuplicateFolder(paths [][]string) [][]string {
	sort.Slice(paths, func(i, j int) bool {
		return len(paths[i]) < len(paths[j])
	})
	nodes := make([]*Node, len(paths))
	rootNode := newNode()
	for i, pathList := range paths {
		current := rootNode
		last := len(pathList) - 1
		for j := 0; j < last; j++ {
			s := pathList[j]
			if _, exists := current.subNodes[s]; !exists {
				current.subNodes[s] = newNode()
			}
			current = current.subNodes[s]
		}
		name := pathList[last]
		node := newNode()
		current.subNodes[name] = node
		nodes[i] = node
	}
	content := strings.Builder{}
	nodeByContent := make(map[string]*Node)
	for i := len(nodes) - 1; i >= 0; i-- {
		node := nodes[i]
		if len(node.subNodes) == 0 {
			continue
		}
		for key, value := range node.subNodes {
			content.WriteString(key)
			content.WriteString("{")
			content.WriteString(value.content)
			content.WriteString("}")
		}
		node.content = content.String()
		content.Reset()
		if similar, exists := nodeByContent[node.content]; exists {
			node.markRemove()
			similar.markRemove()
		} else {
			nodeByContent[node.content] = node
		}
	}
	var ans [][]string
	for i, path := range paths {
		if !nodes[i].remove {
			ans = append(ans, path)
		}
	}
	return ans
}*/

func deleteDuplicateFolder(paths [][]string) [][]string {

	t := NewTrie()
	for _, path := range paths {
		t.insert(path)
	}
	t.dedupe(t.root, make(map[string]*TrieNode))

	ans := [][]string{}
	for _, value := range t.root.children {
		path := []string{}
		t.getPath(value, &path, &ans)
	}
	return ans
}

type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	node := NewTrieNode("")
	return &Trie{node}
}

type TrieNode struct {
	name     string
	children map[string]*TrieNode
	del      bool
}

func NewTrieNode(name string) *TrieNode {
	return &TrieNode{name, make(map[string]*TrieNode), false}
}

func (t *Trie) insert(c []string) {
	node := t.root
	for i := 0; i < len(c); i++ {
		if child, ok := node.children[c[i]]; ok {
			node = child
		} else {
			child = NewTrieNode(c[i])
			node.children[c[i]] = child
			node = child
		}
	}
}

func (t *Trie) dedupe(node *TrieNode, seen map[string]*TrieNode) string {
	var subfolder string
	keys := make([]string, len(node.children))
	i := 0
	for k := range node.children {
		keys[i] = k
		i++
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	for _, key := range keys {
		subfolder += t.dedupe(node.children[key], seen)
	}
	if len(subfolder) != 0 {
		if priorNode, ok := seen[subfolder]; ok {
			priorNode.del = true
			node.del = true
		} else {
			seen[subfolder] = node
		}
	}
	return "(" + node.name + subfolder + ")"
}

func (t *Trie) getPath(node *TrieNode, path *[]string, ans *[][]string) {
	if node.del {
		return
	}
	(*path) = append(*path, node.name)
	finalPath := make([]string, len(*path))
	copy(finalPath, (*path))
	(*ans) = append(*ans, finalPath)
	for _, child := range node.children {
		t.getPath(child, path, ans)
	}
	(*path) = (*path)[:len(*path)-1]
}