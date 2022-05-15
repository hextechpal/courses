import java.util.ArrayDeque;
import java.util.ArrayList;
import java.util.LinkedList;
import java.util.Queue;
import java.util.Scanner;

public class BFS {
  private static final Integer INFINITE = -1;

  private static int distance(ArrayList<Integer>[] adj, int s, int t) {
    int[] distances = new int[adj.length];
    Queue<Integer> queue = new ArrayDeque<>();

    for (int i = 0; i < adj.length; i++) {
      distances[i] = INFINITE;
    }
    distances[s] = 0;
    queue.add(s);

    while (!queue.isEmpty()) {
      int vertex = queue.remove();
      for (int i : adj[vertex]) {
        if (distances[i] == INFINITE) {
          queue.add(i);
          distances[i] = distances[vertex] + 1;
        }
      }
    }

    return distances[t];

  }

  public static void main(String[] args) {
    Scanner scanner = new Scanner(System.in);
    int n = scanner.nextInt();
    int m = scanner.nextInt();
    ArrayList<Integer>[] adj = (ArrayList<Integer>[]) new ArrayList[n];
    for (int i = 0; i < n; i++) {
      adj[i] = new ArrayList<Integer>();
    }
    for (int i = 0; i < m; i++) {
      int x, y;
      x = scanner.nextInt();
      y = scanner.nextInt();
      adj[x - 1].add(y - 1);
      adj[y - 1].add(x - 1);
    }
    int x = scanner.nextInt() - 1;
    int y = scanner.nextInt() - 1;
    System.out.println(distance(adj, x, y));
  }
}

