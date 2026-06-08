#include <iostream>
#include <vector>
using namespace std;

int main() {
    int n;
    cin >> n;

    int m = 1 << n;
    vector<int> a(m + 1);

    for (int i = 1; i <= m; i++) {
        cin >> a[i];
    }

    int mid = m / 2;

    int left_id = 1;
    for (int i = 2; i <= mid; i++) {
        if (a[i] > a[left_id]) {
            left_id = i;
        }
    }

    int right_id = mid + 1;
    for (int i = mid + 2; i <= m; i++) {
        if (a[i] > a[right_id]) {
            right_id = i;
        }
    }

    if (a[left_id] < a[right_id]) {
        cout << left_id << endl;
    } else {
        cout << right_id << endl;
    }

    return 0;
}