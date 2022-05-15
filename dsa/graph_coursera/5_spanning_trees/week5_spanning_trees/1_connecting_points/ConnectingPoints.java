import java.util.HashMap;
import java.util.Map;
import java.util.Scanner;

public class ConnectingPoints {
  private static double minimumDistance(int[] x, int[] y) {
    double result = 0.;
    // write your code here
    return result;
  }

  public static void main(String[] args) {
    Scanner scanner = new Scanner(System.in);
    int n = scanner.nextInt();
    int[] x = new int[n];
    int[] y = new int[n];
    for (int i = 0; i < n; i++) {
      x[i] = scanner.nextInt();
      y[i] = scanner.nextInt();
    }
    System.out.println(minimumDistance(x, y));
  }

  private static class DisjointSet<T> {

    private static class Node<T> {
      int rank;
      T parent;

      Node(T parent, int rank) {
        this.parent = parent;
        this.rank = rank;
      }
    }

    private final Map<T, Node<T>> objectNodeMap = new HashMap<>();

    public T find(T o) {
      Node<T> node = objectNodeMap.get(o);
      if (o == null) {
        return null;
      }

      if (node.parent != o) {
        node.parent = find(node.parent);
      }

      return node.parent;
    }

    public void makeSet(T o) { objectNodeMap.put(o, new Node<>(o, 0)); }

    public void union(T x, T y) {
      T setX = find(x);
      T setY = find(y);
      if (setX == null || setY == null || setX == setY)
        return;
      Node<T> nodeX = objectNodeMap.get(setX);
      Node<T> nodeY = objectNodeMap.get(setY);
      // join the two sets by pointing the root of one at the root of the other
      if (nodeX.rank > nodeY.rank) {
        nodeY.parent = x;
      } else {
        nodeX.parent = y;
        if (nodeX.rank == nodeY.rank)
          nodeY.rank++;
      }
    }
  }
}
