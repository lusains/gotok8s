package com.lusains.algorithm.hash;

import org.junit.jupiter.api.Test;

import java.util.List;

/**
 * @author lvshan@cestc.cn
 * @version 1.0
 * @date 2023-02-09 21:49
 * @description
 */
class ConsistentHashTest {
    @Test
    void get() {
        ConsistentHash consistentHash = new ConsistentHash(
                100, List.of(" node1", " node2", " node3", " node4", " node5"));
        for (int i = 0; i < 100; i++) {
            System.out.println(consistentHash.get("key" + i));
        }

    }
}