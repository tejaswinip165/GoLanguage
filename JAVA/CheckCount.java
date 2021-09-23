import java.util.Scanner;
public class CheckCount {
public static void main(String args[]){

String str;
int upper=0,lower=0,digit=0;
 Scanner scan=new Scanner(System.in); 
    //create a scanner object for input
    
System.out.println("Enter the String  ");
str=scan.nextLine();

for(int i=0; i<str.length(); i++){
char ch=str.charAt(i);
if(ch>='A' && ch<='Z'){
    upper++;
}
else if(ch>='a' && ch<='z'){
    lower++;
}
else if(ch>='0' && ch<='9'){
    digit++;
}
}
System.out.println("lowercase letters: "+lower);
System.out.println("uppercase letters: "+upper);
System.out.println("Numbers: "+digit);
}
}