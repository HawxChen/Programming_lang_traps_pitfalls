#include <algorithm>
#include <vector>
#include <iostream>
#include <iterator>
 
using namespace std;
void p_max_minimum(vector<int>&v, int key) {
    std::vector<int>::iterator low,up;
    cout << "Key is: " << key << endl;
    low = std::lower_bound (v.begin(), v.end(), key);         
    up = std::upper_bound (v.begin(), v.end(), key); 
    std::cout << "lower_bound at position " << (low- v.begin()) << ", element: " << *low << '\n';
    int dist = distance(v.begin(), low);
    cout << "the number of items before this key: " << *low << " is " << dist << endl;
    int result;
    do {
        if(dist == 0) { //important!
            result = dist;
            if(v.front() <= key) {
                result = dist - 1;
            }
        } 
        if(dist == v.size()) {
           result = dist - 1;
           break;
        } 
        result = dist - 1;
    } while (0);
    cout << "result: " << result << endl;
    cout << "---------------------------------------------------------------------" << '\n';
}

int main () {
    int myints[] = {10,20,30,30,20,10,10,20};
    std::vector<int> v(myints,myints+8);           // 10 20 30 30 20 10 10 20

    std::sort (v.begin(), v.end());                // 10 10 10 20 20 20 30 30

    for(auto i = v.begin(); i != v.end(); i++) {
        cout << "idx: " << distance(v.begin(), i) <<  ", val: " <<  *i << '\n';
    }
    p_max_minimum(v, 5);
    p_max_minimum(v, 10);
    p_max_minimum(v, 15);
    p_max_minimum(v, 20);
    p_max_minimum(v, 25);
    p_max_minimum(v, 30);
    p_max_minimum(v, 35);
    return 0;
}
