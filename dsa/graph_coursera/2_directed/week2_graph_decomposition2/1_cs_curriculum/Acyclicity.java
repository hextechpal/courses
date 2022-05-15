import java.util.ArrayList;
import java.util.List;
import java.util.Scanner;

public class Acyclicity {
    private static int acyclic(ArrayList<Integer>[] adj) {
        ArrayList<Integer> visited = new ArrayList<>();
        ArrayList<Integer> recStack = new ArrayList<>();

        for (int i=0; i< adj.length; i++){
            if(isCyclic(adj, i, visited, recStack)){
                return 1;
            }
        }
        return 0;
    }

    private static boolean isCyclic(ArrayList<Integer>[] adj, int from, ArrayList<Integer> visited, ArrayList<Integer> recStack){
        if(recStack.contains(from)){
            return true;
        }

        if(visited.contains(from)){
            return false;
        }

        visited.add(from);
        recStack.add(from);

        for (int i : adj[from]){
            if (isCyclic(adj, i, visited, recStack)){
                return true;
            }
        }

        recStack.remove((Integer) from);
        return false;
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
        }
        System.out.println(acyclic(adj));
    }
}

