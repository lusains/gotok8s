package com.lusains.algorithm.sort;

/**
 * @author lvshan@cestc.cn
 * @version 1.0
 * @date 2023-02-10 13:51
 * @description
 */
public class BubbleSort {
    public void bubbleSort(int[] arr) {
        if (arr == null || arr.length < 2) {
            return;
        }
        for (int i = 0; i < arr.length - 1; i++) {
            for (int j = 0; j < arr.length - 1 - i; j++) {
                if (arr[j] > arr[j + 1]) {
                    swap(arr, j, j + 1);
                }
            }
        }
    }

    public void swap(int[] arr, int i, int j) {
        int temp = arr[i];
        arr[i] = arr[j];
        arr[j] = temp;
    }

    public static void main(String[] args) {
        int[] arr = {1, 3, 2, 5, 4, 6, 7, 9, 8};
        BubbleSort bubbleSort = new BubbleSort();
        bubbleSort.bubbleSort(arr);
        for (int i : arr) {
            System.out.println(i);
        }
    }
}
