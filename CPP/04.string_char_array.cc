#include<iostream>
#include<stdio.h>
using namespace std;
int main() {
    string s = "123566";
    s[0] = '\0';
    s[1] = '\0';
    s[0] = 'C';

    printf("\"%s\"\n", s.c_str());
    //The output will be "C".

    cout << "\"" << s.c_str() << "\"" << endl;
    //The output will be "C".

    cout << "\"" << s << "\"" << endl;
    //The output will be "C3566";
    
    return 0;
}
