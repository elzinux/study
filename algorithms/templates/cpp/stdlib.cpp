/*
不能访问空容器，否则可能出现：
    - 程序崩溃
    - 输出奇怪结果
    - 看起来“没事”
    - 后续某个地方才出问题
*/
#include <iostream>
#include <string>
#include <vector>
#include <algorithm>
#include <queue>
#include <stack>
#include <deque>
#include <set>
#include <map>
#include <unordered_map>
#include <numeric>
#include <tuple>
#include <bitset>
#include <functional>
#include <cmath>
#include <typeinfo>

using namespace std;


void test_vector(){
    vector<int> a = {1,2,3};
    a.push_back(0);
    a.pop_back();
    size_t a_size = a.size();
    cout << a_size << endl;
    a.clear();

    a.resize(2); // 删除多余的元素，只保留 2 个
    a.resize(10, 0); // 加长到 10 并为增加的元素赋值 0

    if (!a.empty()){
        cout << a[0] << endl;
    }

    // 定义一个 3 行 4 列的二维数组
    vector<vector<int>> mat(3, vector<int>(4,0));
    // 定义一个 3 * 4 * 5 的三位数组
    vector<vector<vector<int>>> f(3, vector<vector<int>>(4, vector<int>(5, 0)));
}


void test_queue(){
    queue<int> q;
    q.push(1);
    q.push(0);
    cout << q.front() << endl;
    cout << q.back() << endl;
    q.pop(); // 出队：移除第一个进队的元素，不返回如何内容
    cout << q.size() << endl;
    cout << q.empty() << endl;

    deque<int> dq;
    dq.push_front(0);
    dq.push_back(1);
    dq.pop_back();
    dq.pop_front();
    dq.push_back(10);
    cout << dq.front() << endl;
    cout << dq.back() << endl;

    priority_queue<int> hp; // 默认为大根堆
    hp.push(10);
    hp.push(10);
    cout << hp.top() << endl;
    hp.pop();

    priority_queue<int, vector<int>, greater<int>> min_hp;
    min_hp.push(10);
    cout << min_hp.top() << endl;
    min_hp.pop();

}


void test_stack(){
    stack<int> st;
    st.push(1);
    st.push(2);
    st.push(3);
    int x = st.top();
    cout << x << endl;
    st.pop();
    cout << st.size() << endl;
    cout << st.empty() << endl;
}


void test_set(){
    set<int> s = {1,2,3,3};
    for (int x:s)cout << x << endl; // 自动去重
    s.insert(0);
    s.insert(10);
    cout << s.erase(3) << endl; // 删除并打印 3 的个数，0 或者 1
    if (s.count(3)) cout << "3 in s" << endl;
    s.insert(3);
    set<int>::iterator it = s.find(3); // 返回 3 出现的第一个位置，否则返回s.end()
    // 打印所有大于等于 3 的元素
    for (;it!=s.end();it++) cout << *it << endl;
    s.clear();

    s.insert(2);
    s.insert(1);
    // 当 s 中有元素 2 的时候 s.find(2) == s.lower_bound(2)
    set<int>::iterator it2 = s.lower_bound(2); 
    for (;it2!=s.end();it2++) cout << *it2 << endl;

    multiset<int> ms = {1,1,2,3,3,3}; // 允许有重复的元素
    for (int x: ms) cout << x << ' '; // 打印"1 1 2"
    cout << endl;
    // 与 set 的区别
    cout << ms.erase(3) << endl; // 删除所有 3 并且打印 3 的个数
    ms = {1,2,3,3,3};
    auto it3 = ms.lower_bound(3);
    // 只删除一个 3
    if (it3 != ms.end()) ms.erase(it3);

}


void test_map(){
    map<string,int> mp;
    mp["app"] = 10;
    mp["bna"] += 1;

    mp.erase("app");
    cout << mp["aa"] << endl; // 打印 0，只要访问了，就会创建该键并赋默认值
    cout << mp.size() << endl;
    mp.clear();

    mp["aa"] = 2;
    mp["bb"] = 3;
    map<string,int>::iterator it1 = mp.find("aa");
    auto it2 = mp.lower_bound("aa");
    if (it1 == it2) cout << "yes" << endl; // 成功打印 yes

    // 下面两个 for 打印的内容一致
    for (auto it = mp.begin();it!=mp.end();it++)
        cout << it -> first << ":" << it -> second << endl;
    for (auto pr: mp)
        cout << pr.first << ":" << pr.second << endl;

    // 基于哈希实现的字典
    unordered_map<string, int> ump;

    ump["aa"] = 1;          // 插入或修改
    ump["bb"]++;           // 不存在则自动创建，初值为 0，再加 1

    cout << ump.count("aa") << endl;       // 判断 key 是否存在，返回 0 或 1
    auto ump_it = ump.find("aa");        // 返回迭代器，找不到返回 mp.end()
    // 无序的遍历
    for(;ump_it != ump.end();ump_it++)
        cout << ump_it -> first << ":" << ump_it -> second << endl;

    cout << ump.erase("aa") << endl;       // 删除 key，存在返回 1，不存在返回 0
    cout << ump.size() << endl;            // 元素个数
    cout << ump.empty() << endl;           // 是否为空
    ump.clear();           // 清空

    for(auto &pr:ump) // 使用 & 之后不会在拷贝一份 pr 并且此时修改 pr 会影响 ump 中的内容
        cout << pr.first << ":" << pr.second << endl;

}


void test_string(){
    string s1(100,'0'); // 定义一个长度为 100，每个位置都为 ‘0’ 的字符串
    s1 = "hello cpp!"; // 为 s 重新赋值
    s1[0] = 'H'; // 为 s 的第一个位置赋值
    cout << s1 << endl; // 打印 Hello cpp!

    string s2="!!!";
    string ss = s1 + s2; // 在一个字符串后面尾接一个字符串使用 += 会快很多
    if (s1 != s2) cout << ss << endl;
    
    string subs = ss.substr(3, 4); // 从位置 3 开始，长度为 4 的字串
    cout << subs << endl;
    cout << ss.substr(2,100) << endl; // 不会报错，只会取到最后一个位置

    // 找字串，找到就返回子串出现的第一个位置，否则返回一个预定义的数字
    size_t idx = ss.find("llo");
    if (idx!= string::npos) // O(n^2)
        cout << idx << endl;

    string s = "1110.5";
    int s_i = stoi(s); // 把字符串转化为整型 stoi = string to int
    cout << s_i << endl; // 打印 1110 
    float s_f = stof(s); // 转化为 float
    cout << s_f << endl; // 打印 1110.5

    int num = 100;
    string s_num = to_string(num); // 把其他类型转化为 string
    cout << s_num << endl;
}


void test_pair(){
    pair<char, int> pr = {'1', 2};
    cout << pr.first << " " << pr.second << endl;
}


void func_sort(){
    int a = 1, b = 2;
    swap(a, b); // a = 2, b = 1

    vector<int> arr1 = {9,8,7,5,3,1};
    sort(arr1.begin(), arr1.end(), greater<int>());

    vector<pair<int,int>> arr2 = {{1,2}, {0,3}, {1,1}, {2,0}};

    auto cmp = [](pair<int,int> x, pair<int,int> y){
        if (x.first != y.first){
            return x.first < y.first;
        }
        return x.second > y.second;
    };

    sort(arr2.begin(), arr2.end(), cmp); // cmp 返回 true 则 x 排在 y 前面

}


void func_bisect(){
    vector<int> a = {0,1,1,2,3,5,5,5,6,7,8};
    auto pos = lower_bound(a.begin(), a.end(), 4);
    if (pos != a.end()){
        cout << *pos << endl; // 打印 5
        size_t idx = pos - a.begin();
        cout << a[idx] << endl; // 打印 5
    }else{
        cout << "没有找到" << endl;
    }
    
    

    vector<int>::iterator it = upper_bound(a.begin()+2, a.end(),1);
    cout << it - a.begin() << endl; // 打印 3 upper_bound
    for (;it != a.end();it++)
        cout << *it << endl;
}


void func_reverse(){
    vector<int> a = {1,2,3,4,5,6};
    reverse(a.begin(), a.end());
    for (int x: a) cout << x << " "; // “6 5 4 3 2 1”
    cout << endl;
    reverse(a.begin() + 1, a.begin() + 4);
    for (int x: a) cout << x << " "; // “6 3 4 5 2 1”
    cout << endl;
    
    cout << max({1,2,3,4,5}) << endl; // max(1,2,3) 是错误的
}


void func_unique(){
    // unique 只能对相邻元素进行去重
    // unique 去重之后数组长度不变，会把有效元素移动到前面
    // 返回无效元素的第一个位置，按照位置访问后面的元素，输出值未知的
    vector<int> a = {1,2,3,3,4,5,6,1,2};

    sort(a.begin(),a.end());
    auto it = unique(a.begin(), a.end());
    a.erase(it, a.end());

}


void func_math(){
    int a = -10;
    cout << abs(a) << endl;

    int x = 2,y = 4;
    cout << pow(x, y) << endl; // x ^ y
    cout << sqrt(y) << endl;
    cout << gcd(x,y) << endl; 
    cout << lcm(x, y) << endl;

    function<int(unsigned long long)> bit_width = [](unsigned long long x){
        int res = 0;
        while (x) {
            x >>= 1;
            res += 1;
        }
        return res;
    };
    cout << bit_width(x) << endl; // 输出 x 的二进制位长度


    float b = 10.49;
    cout << ceil(b) << endl; // 向上取整
    cout << floor(b) << endl; // 向下取整
    cout << round(b) << endl; // 4 舍 5 入为整数
}


int main(){
    string s = "10011";
    int x = stoi(s, nullptr, 2);
    cout << x << endl; // 打印 19
    auto x_2 = bitset<8>(x);
    cout << x_2 << endl;
    cout << typeid(x_2).name() << endl; // 打印 x_2 的类型
    string s_2 = x_2.to_string();

    // test_set();
    // test_string();
    // test_pair();
    // func_sort();
    // func_bisect();
    // func_reverse();
    // func_math();
    return 0;
}