#include <iostream>
#include <vector>
using namespace std;

int main() {
    int m, n;
    cin >> m >> n;

    vector<int> a(n);
    for (int i = 0; i < n; i++) {
        cin >> a[i];
    }

    vector<pair<int, int>> queries(m);
    for (int i = 0; i < m; i++) {
        cin >> queries[i].first >> queries[i].second;
    }

    cout << m << " " << n << endl;

    for (int i = 0; i < n; i++) {
        cout << a[i];
        if (i + 1 < n) cout << " ";
    }
    cout << endl;

    for (int i = 0; i < m; i++) {
        cout << queries[i].first << " " << queries[i].second << endl;
    }

    return 0;
}
