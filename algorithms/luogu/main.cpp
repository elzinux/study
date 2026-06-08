#include <vector>
#include <set>
#include <map>
#include <iostream>
#include <typeinfo>

using namespace std;

int main(){
    map<string,int> mp;
    cout << mp["aa"] << endl;
    mp["aa"] = 1;
    mp["b"] = 2;
    map<string,int>::iterator it1 = mp.find("aa");
    auto it2 = mp.lower_bound("aa");
    if (it1 == it2) cout << "yes" << endl; // 成功打印 yes
}

