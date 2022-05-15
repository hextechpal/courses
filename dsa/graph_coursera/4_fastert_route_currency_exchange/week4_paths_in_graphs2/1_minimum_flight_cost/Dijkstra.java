import java.util.*;

public class Dijkstra {
  private static final Integer INFINITE = Integer.MAX_VALUE;

  private static long distance(ArrayList<Integer>[] adj, ArrayList<Integer>[] cost, int s, int t) {
    int[] distances = new int[adj.length];
    boolean[] visited = new boolean[adj.length];

    for (int i = 0; i < adj.length; i++) {
      distances[i] = INFINITE;
    }
    distances[s] = 0;

    while (!areAllTrue(visited)) {
      int vertex = extractMinimum(distances, visited);
      visited[vertex] = true;
      for (int i = 0; i < adj[vertex].size(); i++) {
        int nextVertex = adj[vertex].get(i);
        if (distances[nextVertex] > distances[vertex] + cost[vertex].get(i) && distances[vertex] != INFINITE) {
          distances[nextVertex] = distances[vertex] + cost[vertex].get(i);
        }
      }
    }
    return distances[t] == INFINITE ? -1 : distances[t];
  }

  private static int extractMinimum(int[] distances, boolean[] visited) {
    int minVertex = -1;
    for (int i = 0; i < distances.length; i++) {
      if (!visited[i]) {
        if (minVertex == -1) {
          minVertex = i;
        }
        if (distances[minVertex] > distances[i]) {
          minVertex = i;
        }
      }
    }
    return minVertex;
  }

  public static boolean areAllTrue(boolean... array) {
    for (boolean b : array) {
      if (!b) {
        return false;
      }
    }
    return true;
  }

  public static void main(String[] args) {
    Scanner scanner = new Scanner(System.in);
    int n = scanner.nextInt();
    int m = scanner.nextInt();
    ArrayList<Integer>[] adj = (ArrayList<Integer>[]) new ArrayList[n];
    ArrayList<Integer>[] cost = (ArrayList<Integer>[]) new ArrayList[n];
    for (int i = 0; i < n; i++) {
      adj[i] = new ArrayList<Integer>();
      cost[i] = new ArrayList<Integer>();
    }
    for (int i = 0; i < m; i++) {
      int x, y, w;
      x = scanner.nextInt();
      y = scanner.nextInt();
      w = scanner.nextInt();
      adj[x - 1].add(y - 1);
      cost[x - 1].add(w);
    }
    int x = scanner.nextInt() - 1;
    int y = scanner.nextInt() - 1;
    System.out.println(distance(adj, cost, x, y));
  }
}

