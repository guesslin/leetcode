#!/usr/bin/env python
# -*- coding: utf-8 -*-


class Solution(object):
    def getHint(self, secret, guess):
        """
        :type secret: str
        :type guess: str
        :rtype: str
        """
        a = 0
        b = 0
        secretMap = {}
        guessMap = {}
        for i in xrange(len(secret)):
            addSecret = True
            addGuess = True
            if secret[i] == guess[i]:
                a += 1
                continue

            if guessMap.get(secret[i], 0) > 0:
                b += 1
                guessMap[secret[i]] -= 1
                addSecret = False

            if secretMap.get(guess[i], 0) > 0:
                b += 1
                secretMap[guess[i]] -= 1
                addGuess = False

            if addSecret:
                secretMap[secret[i]] = secretMap.get(secret[i], 0) + 1
            if addGuess:
                guessMap[guess[i]] = guessMap.get(guess[i], 0) + 1

        return '{}A{}B'.format(a, b)


if __name__ == '__main__':
    s = Solution()
    print s.getHint('2962', '7236')
