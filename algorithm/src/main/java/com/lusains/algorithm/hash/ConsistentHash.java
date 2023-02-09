package com.lusains.algorithm.hash;

import java.util.List;
import java.util.SortedMap;
import java.util.TreeMap;

/**
 * @author lvshan@cestc.cn
 * @version 1.0
 * @date 2023-02-09 21:48
 * @description 一致性hash算法
 */
public class ConsistentHash {
    // 虚拟节点数量
    private final int numberOfReplicas;
    // 环上的节点
    private final SortedMap<Long, String> circle = new TreeMap<>();

    public ConsistentHash(int numberOfReplicas, List<String> nodes) {
        this.numberOfReplicas = numberOfReplicas;
        // 初始化环上的节点
        for (String node : nodes) {
            add(node);
        }
    }

    public void add(String node) {
        for (int i = 0; i < numberOfReplicas; i++) {
            long hash = hash(node + i);
            circle.put(hash, node);
        }
    }

    public void remove(String node) {
        for (int i = 0; i < numberOfReplicas; i++) {
            long hash = hash(node + i);
            circle.remove(hash);
        }
    }

    public String get(String key) {
        if (circle.isEmpty()) {
            return null;
        }
        long hash = hash(key);
        if (!circle.containsKey(hash)) {
            SortedMap<Long, String> tailMap = circle.tailMap(hash);
            hash = tailMap.isEmpty() ? circle.firstKey() : tailMap.firstKey();
        }
        return circle.get(hash);
    }

    private long hash(String key) {
        return (long) key.hashCode() & 0xffffffffL;
    }
}
