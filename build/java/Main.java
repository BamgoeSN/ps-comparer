import java.util.Scanner;

public class Main {

    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        String s = sc.next();
        int crocnt = 0;
        int alphacnt = 0;

        for (int i = 0; i < s.length(); i++) {
            if (i < s.length() - 1) {
                if (s.charAt(i) == 'c' && s.charAt(i + 1) == '=') {
                    crocnt++;
                } else if (s.charAt(i) == 'c' && s.charAt(i + 1) == '-') {
                    crocnt++;
                } else if ((s.charAt(i) == 'd' && s.charAt(i + 1) == 'z' && s.charAt(i + 2) == '=')) {
                    crocnt++;
                    if (s.charAt(i + 1) == 'z' && s.charAt(i + 2) == '=') {
                        crocnt--;
                    }
                } else if (s.charAt(i) == 'd' && s.charAt(i + 1) == '-') {
                    crocnt++;
                } else if (s.charAt(i) == 'l' && s.charAt(i + 1) == 'j') {
                    crocnt++;
                } else if (s.charAt(i) == 'n' && s.charAt(i + 1) == 'j') {
                    crocnt++;
                } else if (s.charAt(i) == 's' && s.charAt(i + 1) == '=') {
                    crocnt++;
                } else if (s.charAt(i) == 'z' && s.charAt(i + 1) == '=') {
                    crocnt++;
                } else if (i > 0) {
                    if (s.charAt(i) == 'j') {
                        if (s.charAt(i - 1) != 'n' && s.charAt(i - 1) != 'l') {
                            alphacnt++;
                        }
                    } else if (s.charAt(i) >= 'a' && s.charAt(i) <= 'z' && s.charAt(i) != 'j') {
                        alphacnt++;
                    }
                } else if (i == 0) {
                    if (s.charAt(i) >= 'a' && s.charAt(i) <= 'z') {
                        alphacnt++;
                    }
                }

            } else {
                if (s.charAt(i) >= 'a' && s.charAt(i) <= 'z' && s.charAt(i) != 'j') {
                    alphacnt++;
                } else if (s.charAt(i) == 'j') {
                    if (s.charAt(i - 1) != 'n' && s.charAt(i - 1) != 'l') {
                        alphacnt++;
                    }

                }
            }

        }
        System.out.println(crocnt + alphacnt);

    }
}