import java.util.Scanner;

class Stack {
    final int default_capacity = 10;
    Integer[] array;
    int capacity = 0;
    Integer top_element;
    int size = 0;

    public Stack() {
        array = new Integer[default_capacity];

    }

    public Stack(int k) {
        capacity = k;
        array = new Integer[capacity];
    }

    public void push(int k) {
        if (array[array.length - 1] != null) {
            resize((array.length) * 2);
            push(k);
        } else {
            // System.out.println("k " +k);
            // System.out.println("top "+top_element);
            // System.out.println("size "+size);
            top_element = k;

            array[size] = k;
            size++;
        }

    }

    public void resize(int length) {
        Integer[] temp = new Integer[length];
        for (int i = 0; i < array.length; i++) {
            temp[i] = array[i];
        }
        array = temp;
    }

    public void pop() {
        if (size == 0) {
            System.out.println("-1");
        } else {

            // System.out.println("size : " + size);
            // System.out.println("check : " + array[size]);

            size--;

            top_element = array[size];

            System.out.println(top_element);

        }
    }

    public void size() {
        System.out.println(size);
    }

    public void empty() {
        if (size == 0) {
            System.out.println(1);
        } else {
            System.out.println(0);
        }
    }

    public void top() {
        if (size != 0) {
            if (array[size - 1] != null) {
                System.out.println(top_element);
            } else {
                System.out.println("-1");
            }
        } else {
            System.out.println("-1");
        }
    }
}

public class Main {

    public static void main(String[] args) {
        Stack k = new Stack();
        Scanner input = new Scanner(System.in);
        int j = input.nextInt();
        String m = null;
        int count = 0;
        do {
            count++;
            m = input.nextLine();
            if (m.contains("push")) {
                String[] temp1 = m.split(" ");
                int d2 = 0;
                d2 = Integer.parseInt(temp1[1]);
                k.push(d2);
            } else if (m.contains("pop")) {
                k.pop();
            } else if (m.contains("size")) {
                k.size();
            } else if (m.contains("empty")) {
                k.empty();
            } else if (m.contains("top")) {
                k.top();
            }
        } while (count <= j);
        input.close();

    }
}