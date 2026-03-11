class LFUCache {
public:
    struct Node{
        int k, v, freq;
        Node* prev;
        Node* next;
        Node(int key, int val) : k(key), v(val), freq(1), prev(nullptr), next(nullptr) {}
    };

    struct DLL {
        Node *head, *tail;
        DLL() {
            head = new Node(-1, -1);
            tail = new Node(-1, -1);
            head->next = tail;
            tail->prev = head;
        }

        void addNode(Node* n) {
            n->next = head->next;
            n->prev = head;
            head->next->prev = n;
            head->next = n;
        }

        void removeNode(Node* n) {
            n->prev->next = n->next;
            n->next->prev = n->prev;
        }

        bool isEmpty() {
            return head->next == tail;
        }
    };

    unordered_map<int, Node*> kn;
    unordered_map<int, DLL*> freqList;
    int cap, minFreq, curSize;

    void updateFreq(Node* n) {
        int f = n->freq;
        freqList[f]->removeNode(n);
        if (f == minFreq && freqList[f]->isEmpty()) {
            minFreq++;
        }
        n->freq++;
        if (freqList.find(n->freq) == freqList.end()) {
            freqList[n->freq] = new DLL();
        }
        freqList[n->freq]->addNode(n);
    }

    LFUCache(int capacity) {
        cap = capacity;
        minFreq = 0;
        curSize = 0;
    }

    int get(int key) {
        if (kn.find(key) == kn.end()){
            return -1;
        }
        Node* n = kn[key];
        updateFreq(n);
        return n->v;
    }

    void put(int key, int value) {
        if (cap <= 0){
            return;
        }
        if (kn.find(key) != kn.end()) {
            Node* n = kn[key];
            n->v = value;
            updateFreq(n);
            return;
        }
        if (curSize == cap) {
            DLL* list = freqList[minFreq];
            Node* lruNode = list->tail->prev;
            list->removeNode(lruNode);
            kn.erase(lruNode->k);
            delete lruNode;
            curSize--;
        }
        Node* nn = new Node(key, value);
        kn[key] = nn;
        minFreq = 1;
        if (freqList.find(1) == freqList.end()){
            freqList[1] = new DLL();
        }
        freqList[1]->addNode(nn);
        curSize++;
    }
};

/**
 * Your LFUCache object will be instantiated and called as such:
 * LFUCache* obj = new LFUCache(capacity);
 * int param_1 = obj->get(key);
 * obj->put(key,value);
 */