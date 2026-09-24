Leetcode 409 Longest Palindrome
Time Complexity: O(n)
Space Complexity: O(n)


package main

func (_ Solution) LongestPalindrome(s string) int {
    seen := make(map[rune]struct{})
    result := 0
for _, ch := range s{
    _, ok := seen[ch]
    if ok {
        delete (seen, ch)
        result = result +2
    } else {
        seen[ch] = struct{}{}
    }
}
 if len(seen) != 0{
        result = result +1
    }
    return result
}
