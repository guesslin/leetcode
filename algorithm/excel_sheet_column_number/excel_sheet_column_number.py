class Solution:
    # @param {string} s
    # @return {integer}
    def titleToNumber(self, s):
        table = dict(zip('ABCDEFGHIJKLMNOPQRSTUVWXYZ', [x for x in range(1, 27)]))
        result = 0
        s = s[::-1]
        for x in range(len(s)):
            result += table[s[x]]*(26**x)
        return result
