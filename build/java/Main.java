import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;

public class Main {

    private static int[] arr;
    private static BufferedReader br;

    public static void main(String[] args) throws IOException {
        makeArr();
        initializeArr();
        quickSort(0, arr.length - 1);
        printArr();
    }

    public static void makeArr() throws NumberFormatException, IOException {
        br = new BufferedReader(new InputStreamReader(System.in));
        int size = Integer.parseInt(br.readLine());
        arr = new int[size];
    }

    public static void initializeArr() throws NumberFormatException, IOException {
        for (int i = 0; i < arr.length; i++) {
            arr[i] = Integer.parseInt(br.readLine());
        }
        // suffle();
    }

    public static void quickSort(int left, int right) {
        int pLeft = left;
        int pRight = right;
        int pivot = (left + right) / 2;

        do {
            while (arr[pLeft] < arr[pivot])
                pLeft++;
            while (arr[pRight] > arr[pivot])
                pRight--;

            if (pLeft <= pRight) {
                swap(pLeft++, pRight--);
            }

        } while (pLeft <= pRight);

        if (left < pLeft - 1)
            quickSort(left, pLeft - 1);
        if (right > pLeft)
            quickSort(pLeft, right);
    }

    // public static void suffle() {
    // for (int i = 0; i < 5; i++) {
    // int value1 = (int) (Math.random() * arr.length);
    // int value2 = (int) (Math.random() * arr.length);
    // swap(value1, value2);
    // }
    // }

    public static void swap(int pLeft, int pRight) {
        int tmp = arr[pLeft];
        arr[pLeft] = arr[pRight];
        arr[pRight] = tmp;
    }

    public static void printArr() {
        for (int value : arr) {
            System.out.println(value);
        }
    }
}
