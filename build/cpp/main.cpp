#include <stdio.h>

#define MAX_NUM 100000

int P[MAX_NUM + 10];
int M[MAX_NUM + 10];

int find_min(int num) {
    int i, min_index = num - 1;
    int min = P[num - 1];
    for (i = 0; i < num - 1; i++) {
        if (min > P[i]) {
            min = P[i];
            min_index = i;
        }
    }
    return min_index;
}

int main() {
    int i, N, min_index, pre_index;
    long long total_price = 0.0, M_temp;
    scanf("%d", &N);

    //입력
    for (i = 0; i < N - 1; i++) {
        scanf("%d", &M[i]);
    }
    for (i = 0; i < N; i++) {
        scanf("%d", &P[i]);
    }

    //주유소 제일 싼곳 찾기 A: P배열의 최소값 찾기
    pre_index = N - 1;
    while (1) {
        M_temp = 0.0;
        min_index = find_min(pre_index);
        for (i = min_index; i < pre_index; i++) {
            M_temp += (long long)M[i];
        }
        //printf("\nmin_index=%d pre_index=%d M_temp=%d",min_index, pre_index,M_temp);
        total_price += (long long)P[min_index] * M_temp;
        pre_index = min_index;
        if (min_index == 0)
            break;
    }
    printf("%lld", total_price);
    return 0;
}