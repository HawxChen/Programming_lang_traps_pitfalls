#include<iostream>
#include<stdio.h>
using namespace std;
int main() {
    string s = "123566";
    s[0] = '\0';

    printf("\"%s\"\n", s.c_str());
    //The output will be "".

    cout << "\"" << s.c_str() << "\"" << endl;
    //The output will be "".

    cout << "\"" << s << "\"" << endl;
    //The output will be "23566".
    
    return 0;
}
