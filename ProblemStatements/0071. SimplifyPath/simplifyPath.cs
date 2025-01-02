public class Solution {
    public string SimplifyPath(string path) {
        Stack<string> stack = new Stack<string>();
        string[] parts = path.Split(new char[] {'/'}, StringSplitOptions.RemoveEmptyEntries);
        foreach (string part in parts) {
            if (part == ".") continue;
            if (part == "..") {
                if (stack.Count > 0) stack.Pop();
            } else {
                stack.Push(part);
            }
        }
        if (stack.Count == 0) return "/";
        StringBuilder sb = new StringBuilder();
        foreach (string dir in stack) {
            sb.Insert(0, "/" + dir);
        }
        return sb.ToString();
    }
}