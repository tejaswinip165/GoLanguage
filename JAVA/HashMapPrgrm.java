import java.util.*;
import java.util.Map.Entry;
import java.util.ArrayList;
import java.util.HashMap;



public class HashMapPrgrm {
public static void main(String[] args) {
	HashMap<Integer,ArrayList<String>> entries = new HashMap<Integer,ArrayList<String>>();
	ArrayList<String> first = new ArrayList<String>();
	first.add("Dogs");
	first.add("cats");

	ArrayList<String> second = new ArrayList<String>();
	second.add("Sparrow");
	second.add("crow");
	entries.put(1,first);
	entries.put(2,second);
	
	for(Entry<Integer, ArrayList<String>> ent : entries.entrySet())
	{
		int key = ent.getKey();
		ArrayList<String> value = ent.getValue();
		System.out.println(key+" "+value);
		for(String refer : value)
		{
			System.out.println(refer);
		}
	}
	
}
}
