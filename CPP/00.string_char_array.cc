#include<string>
#include<iostream>
using namespace std;
int main() {
    string x = "abcdefg";
    cout << static_cast<int>(*(x.end())) << endl;
    cout << static_cast<int>('\0') << endl;
}
