import java.util.ArrayList;
import java.util.HashMap;
import java.util.LinkedList;
import java.util.List;
import java.util.Map;
import java.util.Queue;
import java.util.Scanner;

public class Bipartite {
    private static final Integer INFINITE = -1;

    private static int bipartite(ArrayList<Integer>[] adj) {
        if(adj.length < 2){
            return 0;
        }
        int[] distances = new int[adj.length];
        Queue<Integer> queue = new LinkedList<>();
        int s = 0;
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
                }else if (distances[i] == distances[vertex]){
                    return 0;
                }
            }
        }

        for (int distance : distances) {
            if (distance == INFINITE) {
                return 0;
            }
        }
        return 1;

    }

    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        int n = scanner.nextInt();
        int m = scanner.nextInt();
        ArrayList<Integer>[] adj = (ArrayList<Integer>[])new ArrayList[n];
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
        System.out.println(bipartite(adj));
    }
}

