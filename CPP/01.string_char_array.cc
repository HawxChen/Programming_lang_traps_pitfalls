#include<iostream>
#include<stdio.h>
using namespace std;
int main() {
    string s = "123566";
    s[0] = '\0';
    s.push_back('C');

    puts(s.c_str());
    //The output will be "" instead of "C";

    cout << s.c_str() << endl;
    //The output will be "" instead of "C";

    cout << s << endl;
    //The output will be "23566C" instead of "C";
    
    return 0;
}

/*
 *Thanks for yoshijava's discussion.
 */
