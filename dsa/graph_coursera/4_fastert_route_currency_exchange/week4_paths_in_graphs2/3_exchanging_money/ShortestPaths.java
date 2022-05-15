import java.util.ArrayDeque;
import java.util.ArrayList;
import java.util.Queue;
import java.util.Scanner;

public class ShortestPaths {

  private static final int INFINITE = Integer.MAX_VALUE;

  private static void shortestPaths(ArrayList<Integer>[] adj, ArrayList<Integer>[] cost, int s, long[] distance, int[] reachable, int[] shortest) {
    for (int i = 0; i < adj.length; i++) {
      distance[i] = INFINITE;
    }
    bellmanFord(adj, cost, s, distance);
    ArrayList<Integer> allReachable = findAllReachable(adj, finalIteration(adj, cost, distance));
    for (int i = 0; i < adj.length; i++) {
      reachable[i] = distance[i] == INFINITE ? 0 : 1;
      shortest[i] = allReachable.contains(i) ? 0 : 1;
    }
  }

  private static ArrayList<Integer> findAllReachable(ArrayList<Integer>[] adj, ArrayList<Integer> modifiedFinalIteration) {
    ArrayList<Integer> reachableNodes = new ArrayList<>();
    if (modifiedFinalIteration.isEmpty()) {
      return reachableNodes;
    }
    int[] distances = new int[adj.length];
    for (int i = 0; i < adj.length; i++) {
      distances[i] = INFINITE;
    }
    distances[modifiedFinalIteration.get(0)] = 0;
    Queue<Integer> queue = new ArrayDeque<>(modifiedFinalIteration);
    while (!queue.isEmpty()) {
      int vertex = queue.remove();
      for (int i : adj[vertex]) {
        if (distances[i] == INFINITE) {
          queue.add(i);
          distances[i] = distances[vertex] + 1;
        }
      }
    }

    for (int i = 0, distancesLength = distances.length; i < distancesLength; i++) {
      int dist = distances[i];
      if (dist != INFINITE) {
        reachableNodes.add(i);
      }
    }

    return reachableNodes;
  }

  private static ArrayList<Integer> finalIteration(ArrayList<Integer>[] adj, ArrayList<Integer>[] cost, long[] distances) {
    ArrayList<Integer> queue = new ArrayList<>();
    for (int esource = 0; esource < adj.length; esource++) {
      ArrayList<Integer> edges = adj[esource];
      for (int k = 0; k < edges.size(); k++) {
        int destination = edges.get(k);
        if (distances[destination] > distances[esource] + cost[esource].get(k) && distances[esource] != INFINITE) {
          distances[destination] = distances[esource] + cost[esource].get(k);
          queue.add(destination);
        }
      }
    }
    return queue;
  }

  private static void bellmanFord(ArrayList<Integer>[] adj, ArrayList<Integer>[] cost, int source, long[] distances) {
    distances[source] = 0;
    for (int i = 1; i < adj.length; i++) {
      for (int esource = 0; esource < adj.length; esource++) {
        ArrayList<Integer> edges = adj[esource];
        for (int k = 0; k < edges.size(); k++) {
          int destination = edges.get(k);
          if (distances[destination] > distances[esource] + cost[esource].get(k) && distances[esource] != INFINITE) {
            distances[destination] = distances[esource] + cost[esource].get(k);
          }
        }
      }
    }
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
    int s = scanner.nextInt() - 1;
    long distance[] = new long[n];
    int reachable[] = new int[n];
    int shortest[] = new int[n];
    for (int i = 0; i < n; i++) {
      distance[i] = Long.MAX_VALUE;
      reachable[i] = 0;
      shortest[i] = 1;
    }
    shortestPaths(adj, cost, s, distance, reachable, shortest);
    for (int i = 0; i < n; i++) {
      if (reachable[i] == 0) {
        System.out.println('*');
      } else if (shortest[i] == 0) {
        System.out.println('-');
      } else {
        System.out.println(distance[i]);
      }
    }
  }
}
