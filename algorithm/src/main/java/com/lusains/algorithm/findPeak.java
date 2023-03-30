package com.lusains.algorithm;

import java.util.ArrayList;
import java.util.List;

/**
 * @author lvshan@cestc.cn
 * @version 1.0
 * @date 2023-02-24 18:54
 * @description
 */
public class findPeak {

    public static int[] findPeakElements(int[] nums) {
        List<Integer> peaks = new ArrayList<>();
        int left = 0, right = nums.length - 1;
        // 分别查找左右两侧的峰值
        while (left < right) {
            int mid = left + (right - left) / 2;
            if (nums[mid] > nums[mid - 1] && nums[mid] > nums[mid + 1]) {
                peaks.add(mid);
            }
            if (nums[mid] < nums[mid + 1]) {
                left = mid + 1;
            } else {
                right = mid - 1;
            }
        }
        // 将峰值列表转换成数组返回
        int[] result = new int[peaks.size()];
        for (int i = 0; i < peaks.size(); i++) {
            result[i] = peaks.get(i);
        }
        return result;
    }

    public static void main(String[] args) {
        int[] nums = new int[]{1, 5, 4, 3, 5, 6, 4};
        int[] result = findPeakElements(nums);
        // 打印结果
        for (int i : result) {
            System.out.println(i);
        }

    }

}
