package com.lusains.algorithm.ratelimit;

/**
 * @author lvshan@cestc.cn
 * @version 1.0
 * @date 2023-02-13 11:37
 * @description
 */
public class AdaptiveRateLimiter {
    private final int capacity;
    private final int rate;
    private int tokens = 0;
    private long lastRefillTime = System.currentTimeMillis();

    public AdaptiveRateLimiter(int capacity, int rate) {
        this.capacity = capacity;
        this.rate = rate;
    }

    public boolean allowRequest() {
        refillToken();
        if (tokens > 0) {
            tokens--;
            return true;
        } else {
            return false;
        }
    }

    private void refillToken() {
        long currentTime = System.currentTimeMillis();
        int elapsedTime = (int) (currentTime - lastRefillTime);
        lastRefillTime = currentTime;
        tokens = Math.min(capacity, tokens + elapsedTime * rate);
    }
}

