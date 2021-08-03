#include <cstdio>

int main() {
    int num, n, count = 0;
    int arr[41] = {
        0,
    };

    for (int i = 0; i < 10; i++) {
        scanf("%d", &num);
        n = num % 42;
        arr[n]++;
    }

    for (int i = 0; i < 42; i++) {
        if (arr[i] > 0) {
            count++;
        }
    }

    printf("%d\n", count);

    return 0;
}