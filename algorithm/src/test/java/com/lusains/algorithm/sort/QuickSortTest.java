package com.lusains.algorithm.sort;

import org.junit.jupiter.api.Test;

/**
 * @author lvshan@cestc.cn
 * @version 1.0
 * @date 2023-02-10 10:33
 * @description
 */
class QuickSortTest {

    @Test
    void quickSort() {
        int[] arr = {1, 3, 2, 5, 4, 6, 7, 9, 8};
        QuickSort quickSort = new QuickSort();
        quickSort.quickSort(arr);
        for (int i : arr) {
            System.out.println(i);
        }
    }
}