import java.io.*;
import java.util.*;

class Node
{
	public static final int Letters =  4;
	public static final int NA      = -1;
	public int next [];
	public boolean patternEnd;

	Node ()
	{
		next = new int [Letters];
		Arrays.fill (next, NA);
		patternEnd = false;
	}

	@Override
	public String toString(){
		StringBuilder sb = new StringBuilder();
		sb.append(patternEnd);
		sb.append("(");
		for (int i = 0; i < next.length; i++) {
			sb.append(next[i]);	
			sb.append(" ");	
		}
		sb.append(")");
		return sb.toString();
	}
}

public class TrieMatchingExtended implements Runnable {
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

	List<Node> buildTrie(List<String> patterns) {
        List<Node> trie = new ArrayList<>();
        trie.add(new Node());
        for(String pattern : patterns){
            Node current = trie.get(0);
            for(int i=0; i < pattern.length(); i++){
				Character c = pattern.charAt(i);
				int idx = letterToIndex(c);
                int val = current.next[idx];
                if(val != Node.NA){
					Node node = trie.get(val);
					if (!node.patternEnd){
						node.patternEnd = i == pattern.length() - 1;
					}
                    current = node;
                }else{
                    current.next[idx] = trie.size();
					Node n = new Node();
					n.patternEnd = i == pattern.length() - 1 ;
                    trie.add(n);
					current = n;
                }
            }
			// System.out.println(trie);
        }
        return trie;
    }

	List <Integer> solve (String text, int n, List <String> patterns) {
		List <Integer> result = new ArrayList <Integer> ();
		List<Node> trie = buildTrie(patterns);

		int h = 0;
		while(h < text.length()){
			if(prefixTrieMatch(text.substring(h), trie)){
				result.add(h);
			}
			h++;
		}
		return result;
	}

	boolean prefixTrieMatch(String text, List<Node> trie){
		int f = 0;
		int v = 0;

		while(true) {
			if(trie.get(v).patternEnd){
				return true;
			}else{
				if(text.length() <= f){
					return false;
				}else{
					Node node = trie.get(v);
					char c = text.charAt(f);
					if(node.next[letterToIndex(c)] != Node.NA){
						v = node.next[letterToIndex(c)];
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
		new Thread (new TrieMatchingExtended ()).start();
	}
}
