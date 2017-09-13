// Copyright (c) 2015 Elements of Programming Interviews. All rights reserved.
// Show some inconsistent bugs through API usage!

#include <cassert>
#include <vector>
#include <iostream>
#include <algorithm>

using namespace std;

// @include
int bsearch_api(int t, const vector<int>& A) {

    /* If the input vector is not sorted, then we should use find.
     * If ... is SORTED, then we can use lower_bound().
     * However, the behavior has conflicts with the well-known book: The C++ Standard Library.
     * [Conflict] In the book's page 611, All algorithms (upper_bound and lower_bound) return "end" if there is NO such value.
     *          
     * PS: from online manual: http://www.cplusplus.com/reference/algorithm/lower_bound/ 
     *     Return value:
     *       An iterator to the lower bound of val in the range.
     *       If all the element in the range compare less than val, the function returns last.
     *
     */
    auto iter = lower_bound(A.begin(), A.end(), t);

#if 0
    /*Inconsistency
     * use upper bound!
     * A = {2, 2, 2};
     * assert(-1 == bsearch(3, A));
     * the lowber bound will return iter of A[0] instead of A.end().
     */
     return iter == A.end() ? -1 : distance(A.begin(), iter);
#else
    //Corrected check.
    return iter == A.end() ? -1 : *iter != t ? -1 : distance(A.begin(), iter);
#endif
}

int bsearch(int t, const vector<int>& A) {
//    return bsearch_handmade(t, A);
    return bsearch_api(t, A);
}
// @exclude

static void SimpleTest() {
  vector<int> A = {1, 2, 3};
  assert(0 == bsearch(1, A));
  assert(1 == bsearch(2, A));
  assert(2 == bsearch(3, A));
  A = {2, 2, 2};
  assert(0 <= bsearch(2, A) && bsearch(2, A) <= 2);
  assert(-1 == bsearch(3, A));
  assert(-1 == bsearch(0, A));
}

int main(void) {
  SimpleTest();
  vector<int> A;
  A.emplace_back(1);
  A.emplace_back(2);
  A.emplace_back(2);
  A.emplace_back(2);
  A.emplace_back(2);
  A.emplace_back(3);
  A.emplace_back(3);
  A.emplace_back(3);
  A.emplace_back(5);
  A.emplace_back(6);
  A.emplace_back(10);
  A.emplace_back(100);
  assert(0 == bsearch(1, A));
  assert(1 <= bsearch(2, A) && bsearch(2, A) <= 4);
  assert(5 <= bsearch(3, A));
  assert(-1 == bsearch(4, A));
  return 0;
}
