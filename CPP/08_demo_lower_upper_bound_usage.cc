#include <algorithm>
#include <vector>
#include <iostream>
#include <iterator>
 
using namespace std;
void p_lower_upper(vector<int>&v, int key) {
    /*
     reference: http://www.cplusplus.com/reference/algorithm/lower_bound/
     */
    std::vector<int>::iterator low,up;
    cout << "Key is: " << key << endl;
    low=std::lower_bound (v.begin(), v.end(), key);         
    up= std::upper_bound (v.begin(), v.end(), key); 
    std::cout << "lower_bound at position " << (low- v.begin()) << ", element: " << *low << '\n';
    std::cout << "upper_bound at position " << (up - v.begin()) << ", element: " << *up << '\n';
    cout << "---------------------------------------------------------------------" << '\n';
}

int main () {
    int myints[] = {10,20,30,30,20,10,10,20};
    std::vector<int> v(myints,myints+8);           // 10 20 30 30 20 10 10 20

    std::sort (v.begin(), v.end());                // 10 10 10 20 20 20 30 30

    for(auto i = v.begin(); i != v.end(); i++) {
        cout << "idx: " << distance(v.begin(), i) <<  ", val: " <<  *i << '\n';
    }
    p_lower_upper(v, 5);
    p_lower_upper(v, 10);
    p_lower_upper(v, 15);
    p_lower_upper(v, 20);
    p_lower_upper(v, 25);
    p_lower_upper(v, 30);
    p_lower_upper(v, 35);
    return 0;
}
