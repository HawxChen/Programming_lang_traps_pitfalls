#include<stdlib.h>
#include<iostream>
using namespace std;
int main(void) {
    string minus_mod = "python -c \"print -1 % 8\"";
    cout << minus_mod << ": " << endl;
    system(minus_mod.c_str());
    string positive_mod = "python -c \"print (-1+8) % 8\"";
    cout << positive_mod << ": " << endl;
    system(positive_mod.c_str());

    cout << "C++: -1 % 8: " << - 1 % 8 << endl;
    cout << "C++: (-1+8) % 8: " << (-1+8) % 8 << endl;


}
