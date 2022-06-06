import java.io.*;
import java.util.*;

class Node
{
	public static final int Letters =  4;
	public static final int NA      = -1;
	public int next [];

	Node ()
	{
		next = new int [Letters];
		Arrays.fill (next, NA);
	}
}

public class TrieMatching implements Runnable {

	int letterToIndex (char letter)
	{
		switch (letter)
		{
			case 'A': return 0;
			case 'C': return 1;
			case 'G': return 2;
			case 'T': return 3;
			default: assert (false); return Node.NA;
		}
	}

	List<Map<Character, Integer>> buildTrie(List<String> patterns) {
        List<Map<Character, Integer>> trie = new ArrayList<Map<Character, Integer>>();
        trie.add(new HashMap<>());
        for(String pattern : patterns){
            int current = 0;
            for(Character c : pattern.toCharArray()){
                Map<Character, Integer> edges = trie.get(current);
                if(edges != null && edges.containsKey(c)){
                    current = edges.get(c);
                }else{
                    if(edges == null){
                        edges = new HashMap<>();                        
                    }
                    current = trie.size();
                    edges.put(c, trie.size());
                    trie.add(new HashMap<>());
                }
            }
        }
        return trie;
    }

	List <Integer> solve (String text, int n, List <String> patterns) {
		List <Integer> result = new ArrayList <Integer> ();
		List<Map<Character, Integer>> trie = buildTrie(patterns);
		int h = 0;
		while(h < text.length()){
			if(prefixTrieMatch(text.substring(h), trie)){
				result.add(h);
			}
			h++;
		}
		return result;
	}

	boolean prefixTrieMatch(String text, List<Map<Character, Integer>> trie){
		int f = 0;
		int v = 0;

		while(true) {
			if(trie.get(v) == null || trie.get(v).isEmpty()){
				return true;
			}else{
				if(text.length() <= f){
					return false;
				}else{
					Map<Character, Integer> edges = trie.get(v);
					if(edges.containsKey(text.charAt(f))){
						v = edges.get(text.charAt(f));
						f++;
					}else{
						return false;
					}
				}
			}
		}
	}

	public void run () {
		try {
			BufferedReader in = new BufferedReader (new InputStreamReader (System.in));
			String text = in.readLine ();
		 	int n = Integer.parseInt (in.readLine ());
		 	List <String> patterns = new ArrayList <String> ();
			for (int i = 0; i < n; i++) {
				patterns.add (in.readLine ());
			}

			List <Integer> ans = solve (text, n, patterns);

			for (int j = 0; j < ans.size (); j++) {
				System.out.print ("" + ans.get (j));
				System.out.print (j + 1 < ans.size () ? " " : "\n");
			}
		}
		catch (Throwable e) {
			e.printStackTrace ();
			System.exit (1);
		}
	}

	public static void main (String [] args) {
		new Thread (new TrieMatching ()).start ();
	}
}
