#include <iostream>
#include <vector>
#include <queue>
using namespace std;

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    int n;
    cin >> n;

    vector<int> left_child(n + 1), right_child(n + 1);

    for (int i = 1; i <= n; i++) {
        cin >> left_child[i] >> right_child[i];
    }

    queue<pair<int, int>> q;
    q.push({1, 1}); // 节点编号，当前深度

    int ans = 0;

    while (!q.empty()) {
        auto [u, depth] = q.front();
        q.pop();

        ans = max(ans, depth);

        if (left_child[u] != 0) {
            q.push({left_child[u], depth + 1});
        }

        if (right_child[u] != 0) {
            q.push({right_child[u], depth + 1});
        }
    }

    cout << ans << '\n';

    return 0;
}