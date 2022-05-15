import java.util.ArrayList;
import java.util.HashMap;
import java.util.Map;
import java.util.Scanner;

public class StronglyConnected {
  private static int numberOfStronglyConnectedComponents(ArrayList<Integer>[] adj) {
    ArrayList<Integer>[] reversedAdj = reverseGraph(adj);
    boolean[] used = new boolean[reversedAdj.length];
    ArrayList<Integer> order = new ArrayList<Integer>();
    for (int i = 0; i < reversedAdj.length; i++) {
      if (!used[i]) {
        dfs(reversedAdj, used, order, i);
      }
    }

    Map<Integer, ArrayList<Integer>> scc = new HashMap<>();
    used = new boolean[adj.length];

    int count = 0;
    for (int vertex : order) {
      if (!used[vertex]) {
        scc.put(count, calculateScc(adj, used, new ArrayList<>(), vertex));
        count++;
      }
    }

    return scc.size();

  }

  private static ArrayList<Integer>[] reverseGraph(ArrayList<Integer>[] adj) {
    ArrayList<Integer>[] reverseGraph = new ArrayList[adj.length];
    for (int i = 0; i < adj.length; i++) {
      reverseGraph[i] = new ArrayList<Integer>();
    }
    for (int i = 0; i < adj.length; i++) {
      for (int e : adj[i]) {
        reverseGraph[e].add(i);
      }
    }
    return reverseGraph;
  }

  private static void dfs(ArrayList<Integer>[] adj, boolean[] used, ArrayList<Integer> order, int s) {
    explore(adj, used, order, s);
    order.add(0, s);
  }

  private static ArrayList<Integer> calculateScc(ArrayList<Integer>[] adj, boolean[] used, ArrayList<Integer> scc, int s) {
    used[s] = true;
    scc.add(s);
    for (int i : adj[s]) {
      if (!used[i]) {
        dfs(adj, used, scc, i);
      }
    }
    return scc;
  }

  private static void explore(ArrayList<Integer>[] adj, boolean[] used, ArrayList<Integer> order, int s) {
    used[s] = true;
    for (int i : adj[s]) {
      if (!used[i]) {
        dfs(adj, used, order, i);
      }
    }
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
    }
    System.out.println(numberOfStronglyConnectedComponents(adj));
  }
}
