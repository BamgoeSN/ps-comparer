#include <iostream>
using namespace std;

struct NODE {
    int sp;
    int lowernodenum[49];
    int index;
    bool disabled;
    NODE() {
        disabled = false;
        for (index = 49; index > 0;) lowernodenum[--index] = -1;
    }
};

void circuit(NODE*& node, const int& key, const int& ex, int& result) {
    NODE& temp = node[key];

    if (temp.disabled) return;
    bool leafcheck = true;
    for (temp.index = 0; temp.lowernodenum[temp.index] != -1;) {
        if (temp.lowernodenum[temp.index] != -1) {
            leafcheck = false;
            break;
        }
    }
    if (leafcheck) {
        result += 1;
        return;
    }
    for (temp.index = 0; temp.lowernodenum[temp.index] != -1;) circuit(node, temp.lowernodenum[temp.index++], ex, result);
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(NULL);
    cout.tie(NULL);

    int N, super, result = 0, ex;
    cin >> N;
    NODE* node = new NODE[N];

    for (int i = 0; i < N; i++) {
        cin >> super;
        node[i].sp = super;
        if (super == -1) continue;
        node[super].lowernodenum[node[super].index++] = i;
    }

    for (int i = 0; i < N; i++) node[i].index = 0;

    cin >> ex;
    node[ex].disabled = true;
    for (int i = 0; i < N; i++) {
        if (node[node[ex].sp].lowernodenum[i] == ex) {
            node[node[ex].sp].lowernodenum[i] = -1;
            break;
        }
    }
    circuit(node, 0, ex, result);

    cout << result;

    delete[] node;
    return 0;
}