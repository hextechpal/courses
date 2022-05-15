import java.util.ArrayList;
import java.util.Arrays;
import java.util.Scanner;

public class NegativeCycle {

  private static final int INFINITE = Integer.MAX_VALUE;

  private static int negativeCycle(ArrayList<Integer>[] adj,
                                   ArrayList<Integer>[] cost) {
    ArrayList<Integer>[] connectedAdj = makeGraphConnected(adj);
    ArrayList<Integer>[] connectedCosts = makeCostConnected(cost);

    int[] distances = new int[connectedAdj.length];
    int[] prev = new int[connectedAdj.length];
    for (int i = 0; i < connectedAdj.length; i++) {
      distances[i] = INFINITE;
    }
    bellmanFord(connectedAdj, connectedCosts, connectedAdj.length - 1,
                distances);

    return finalIteration(connectedAdj, connectedCosts, distances);
  }

  private static int finalIteration(ArrayList<Integer>[] adj,
                                    ArrayList<Integer>[] cost,
                                    int[] distances) {
    boolean cycleFound = false;
    for (int esource = 0; esource < adj.length; esource++) {
      ArrayList<Integer> edges = adj[esource];
      for (int k = 0; k < edges.size(); k++) {
        int destination = edges.get(k);
        if (distances[destination] >
                distances[esource] + cost[esource].get(k) &&
            distances[esource] != INFINITE) {
          cycleFound = true;
        }
      }
    }

    return cycleFound ? 1 : 0;
  }

  private static void bellmanFord(ArrayList<Integer>[] adj,
                                  ArrayList<Integer>[] cost, int source,
                                  int[] distances) {
    distances[source] = 0;
    for (int i = 1; i < adj.length; i++) {
      for (int esource = 0; esource < adj.length; esource++) {
        ArrayList<Integer> edges = adj[esource];
        for (int k = 0; k < edges.size(); k++) {
          int destination = edges.get(k);
          if (distances[destination] >
                  distances[esource] + cost[esource].get(k) &&
              distances[esource] != INFINITE) {
            distances[destination] = distances[esource] + cost[esource].get(k);
          }
        }
      }
    }
  }

  private static ArrayList<Integer>[] makeCostConnected(
      ArrayList<Integer>[] cost) {
    ArrayList<Integer>[] connectedAdj = Arrays.copyOf(cost, cost.length + 1);
    ArrayList<Integer> edges = new ArrayList<>();
    for (int i = 0; i < cost.length; i++) {
      edges.add(1);
    }
    connectedAdj[cost.length] = edges;
    return connectedAdj;
  }

  private static ArrayList<Integer>[] makeGraphConnected(
      ArrayList<Integer>[] adj) {
    ArrayList<Integer>[] connectedAdj = Arrays.copyOf(adj, adj.length + 1);
    ArrayList<Integer> edges = new ArrayList<>();
    for (int i = 0; i < adj.length; i++) {
      edges.add(i);
    }
    connectedAdj[adj.length] = edges;
    return connectedAdj;
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
    System.out.println(negativeCycle(adj, cost));
  }
}
