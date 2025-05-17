package leetcode.easy;

import java.util.HashSet;
import java.util.Set;

public class JewelsAndStones {
    public int numJewelsInStones(String jewels, String Stones) {
        Set<Character> jewelSet = new HashSet<>();

        for (char c : jewels.toCharArray()) {
            jewelSet.add(c);
        }

        int count = 0;

        for (char c : Stones.toCharArray()) {
            if (jewelSet.contains(c)) {
                count++;
            }
        }

        return count;
    }
}

