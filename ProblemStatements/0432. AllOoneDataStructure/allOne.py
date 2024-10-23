class CountNode:
    def __init__(self, count):
        self.count = count
        self.keys = set()
        self.prev = None
        self.next = None

class AllOne:

    def __init__(self):
        self.key_count = {}
        self.count_nodes = {}
        self.head = CountNode(float('-inf'))
        self.tail = CountNode(float('inf'))
        self.head.next = self.tail
        self.tail.prev = self.head

    def _add_node_after(self, new_node, prev_node):
        new_node.prev = prev_node
        new_node.next = prev_node.next
        prev_node.next.prev = new_node
        prev_node.next = new_node

    def _remove_node(self, node):
        node.prev.next = node.next
        node.next.prev = node.prev

    def inc(self, key: str) -> None:
        current_count = self.key_count.get(key, 0)
        new_count = current_count + 1
        self.key_count[key] = new_count

        if current_count > 0:
            current_node = self.count_nodes[current_count]
            current_node.keys.remove(key)
            if not current_node.keys:
                self._remove_node(current_node)
                del self.count_nodes[current_count]

        if new_count not in self.count_nodes:
            new_node = CountNode(new_count)
            self.count_nodes[new_count] = new_node
            prev_node = self.head
            while prev_node.next.count < new_count:
                prev_node = prev_node.next
            self._add_node_after(new_node, prev_node)

        self.count_nodes[new_count].keys.add(key)

    def dec(self, key: str) -> None:
        current_count = self.key_count[key]
        new_count = current_count - 1

        current_node = self.count_nodes[current_count]
        current_node.keys.remove(key)

        if new_count > 0:
            self.key_count[key] = new_count
            if new_count not in self.count_nodes:
                new_node = CountNode(new_count)
                self.count_nodes[new_count] = new_node
                prev_node = self.head
                while prev_node.next.count < new_count:
                    prev_node = prev_node.next
                self._add_node_after(new_node, prev_node)

            self.count_nodes[new_count].keys.add(key)
        else:
            del self.key_count[key]

        if not current_node.keys:
            self._remove_node(current_node)
            del self.count_nodes[current_count]

    def getMaxKey(self) -> str:
        if self.tail.prev == self.head:
            return ""
        return next(iter(self.tail.prev.keys))

    def getMinKey(self) -> str:
        if self.head.next == self.tail:
            return ""
        return next(iter(self.head.next.keys))


# Initializing
# hello -> NULL
# hello -> hello -> NULL
# 2 hello
# 2 hello
# leet -> hello -> hello -> NULL
# 2 hello
# 1 leet

# Your AllOne object will be instantiated and called as such:
# obj = AllOne()
# obj.inc(key)
# obj.dec(key)
# param_3 = obj.getMaxKey()
# param_4 = obj.getMinKey()